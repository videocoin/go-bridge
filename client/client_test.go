package client

import (
	"context"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/videocoin/go-bridge/nativebridge"
	"github.com/videocoin/go-bridge/nativeproxy"
	"github.com/videocoin/go-bridge/remotebridge"
	"github.com/videocoin/go-bridge/service"
	"github.com/videocoin/go-bridge/service/nativetotoken"
	"github.com/videocoin/go-bridge/service/tokentonative"
	"github.com/videocoin/go-bridge/testtools"
)

func TestClient(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}

type ClientSuite struct {
	testtools.Suite

	client     *Client
	nativeBank common.Address
	tokenBank  common.Address

	wg     sync.WaitGroup
	ctx    context.Context
	cancel func()
}

func (s *ClientSuite) SetupTest() {
	s.Suite.SetupTest()

	s.ctx, s.cancel = context.WithCancel(context.Background())

	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	log := logrus.NewEntry(logger)

	_, _, remote, err := remotebridge.DeployRemoteBridge(s.FundedKeys[0], s.Backend)
	s.Require().NoError(err)

	_, _, proxy, err := nativeproxy.DeployNativeProxy(s.FundedKeys[0], s.Backend)
	s.Require().NoError(err)

	_, _, native, err := nativebridge.DeployNativeBridge(s.FundedKeys[0], s.Backend)
	s.Require().NoError(err)

	s.Backend.Commit()

	s.nativeBank = s.FundedKeys[2].From
	s.tokenBank = s.FundedKeys[3].From

	s.client = NewClient(s.Backend, s.Backend, proxy, s.ERC20, native, remote)

	_, err = s.ERC20.Transfer(s.FundedKeys[0], s.FundedKeys[2].From, big.NewInt(1e18))
	s.Require().NoError(err)
	s.Backend.Commit()
	erctransfer, err := nativetotoken.NewTransferERC20Transactor(s.Backend, s.ERC20Address)
	ntengine := nativetotoken.NewTransferEngine(log,
		s.Backend, s.Backend,
		*s.FundedKeys[0], *s.FundedKeys[2],
		remote, s.ERC20, erctransfer)
	ntsvc := service.NewService(log,
		s.Backend,
		ntengine,
		ntengine,
		nativetotoken.NewNativeTransferAccess(log, proxy),
		service.StaticSource{s.nativeBank},
		big.NewInt(0), big.NewInt(100),
	)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		_ = service.PollForever(s.ctx, time.Second, 10*time.Minute, func(ctx context.Context) {
			err := ntsvc.Run(ctx)
			if err != nil {
				log.Debugf("poll failed with %v", err)
			}
		})
	}()

	tnengine := tokentonative.NewTransferEngine(log, s.Backend, *s.FundedKeys[1], native)
	tnsvc := service.NewService(log, s.Backend,
		tnengine,
		tnengine,
		tokentonative.NewERC20Access(log, s.ERC20),
		service.StaticSource{s.tokenBank},
		big.NewInt(3), big.NewInt(100),
	)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		_ = service.PollForever(s.ctx, time.Second, 10*time.Minute, func(ctx context.Context) {
			err := tnsvc.Run(ctx)
			if err != nil {
				log.Debugf("poll failed with %v", err)
			}
		})
	}()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				s.Backend.Commit()
			}
		}
	}()
}

func (s *ClientSuite) TearDownTest() {
	s.cancel()
	s.wg.Wait()
	s.Suite.TearDownTest()
}

func (s *ClientSuite) TestDeposit() {
	ctx := context.Background()
	amount := big.NewInt(100)
	info, err := s.client.WaitDeposit(ctx, s.FundedPrivateKeys[4], s.tokenBank, amount)
	s.Require().NoError(err)

	bridged, _, err := s.Backend.TransactionByHash(ctx, info.LocalTxHash)
	s.Require().NoError(err)
	s.Require().Equal(amount.Int64(), bridged.Value().Int64())
}

func (s *ClientSuite) TestWithdraw() {
	ctx := context.Background()
	amount := big.NewInt(100)
	_, err := s.client.WaitWithdraw(ctx, s.FundedPrivateKeys[4], s.nativeBank, amount)
	s.Require().NoError(err)

	balance, err := s.ERC20.BalanceOf(&bind.CallOpts{Context: ctx}, s.FundedKeys[4].From)
	s.Require().NoError(err)
	s.Require().Equal(amount.Int64(), balance.Int64())
}
