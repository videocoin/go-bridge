package tokentonative

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/videocoin/go-bridge/nativebridge"
	"github.com/videocoin/go-bridge/service"
	"github.com/videocoin/go-bridge/testtools"
)

func TestTokenToNativeService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	testtools.Suite

	bridge *nativebridge.NativeBridge
	log    *logrus.Entry
}

func (s *ServiceSuite) SetupTest() {
	s.Suite.SetupTest()
	opts := *s.FundedKeys[0]
	_, _, bridge, err := nativebridge.DeployNativeBridge(&opts, s.Backend)
	s.Require().NoError(err)
	s.Backend.Commit()
	s.bridge = bridge

	opts = *s.FundedKeys[0]
	opts.Value = big.NewInt(1e16)

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	s.log = logrus.NewEntry(logger)
}

func (s *ServiceSuite) TestServiceExecuteAll() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	engine := NewTransferEngine(s.log, s.Backend, *s.FundedKeys[0], s.bridge)
	service := service.NewService(s.log, s.Backend,
		engine,
		engine,
		NewERC20Access(s.log, s.ERC20),
		service.StaticSource{s.FundedKeys[0].From},
		big.NewInt(0), big.NewInt(10),
	)

	txs := []common.Hash{}
	for i := 1; i < len(s.FundedKeys)-1; i++ {
		opts := *s.FundedKeys[i]
		tx, err := s.ERC20.Transfer(&opts, s.FundedKeys[0].From, big.NewInt(10))
		s.Require().NoError(err)
		txs = append(txs, tx.Hash())
	}
	s.Backend.Commit()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				s.Backend.Commit()
			}
		}
	}()

	s.Require().NoError(service.Run(ctx))

	for _, tx := range txs {
		executed, err := s.bridge.Transfers(&bind.CallOpts{Context: ctx}, [32]byte(tx))
		s.Require().NoError(err)
		s.Require().True(executed)
	}
}

func (s *ServiceSuite) TestServiceNothingExecuted() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	engine := NewTransferEngine(s.log, s.Backend, *s.FundedKeys[0], s.bridge)
	service := service.NewService(s.log, s.Backend,
		engine,
		engine,
		NewERC20Access(s.log, s.ERC20),
		service.StaticSource{},
		big.NewInt(0), big.NewInt(10),
	)

	txs := []common.Hash{}
	for i := 1; i < len(s.FundedKeys)-1; i++ {
		opts := *s.FundedKeys[i]
		tx, err := s.ERC20.Transfer(&opts, s.FundedKeys[0].From, big.NewInt(10))
		s.Require().NoError(err)
		txs = append(txs, tx.Hash())
	}
	s.Backend.Commit()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				s.Backend.Commit()
			}
		}
	}()

	s.Require().NoError(service.Run(ctx))

	for _, tx := range txs {
		executed, err := s.bridge.Transfers(&bind.CallOpts{Context: ctx}, [32]byte(tx))
		s.Require().NoError(err)
		s.Require().False(executed)
	}
}
