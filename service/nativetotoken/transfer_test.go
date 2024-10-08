package nativetotoken

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/videocoin/go-bridge/remotebridge"
	"github.com/videocoin/go-bridge/service"
	"github.com/videocoin/go-bridge/testtools"
)

func TestTransferEngine(t *testing.T) {
	suite.Run(t, new(EngineSuite))
}

type EngineSuite struct {
	testtools.Suite

	log *logrus.Entry

	bridge *remotebridge.RemoteBridge

	engine      *TransferEngine
	erctransfer *ERC20TransferTransactor
}

func (s *EngineSuite) SetupTest() {
	s.Suite.SetupTest()

	opts := *s.FundedKeys[0]
	_, _, bridge, err := remotebridge.DeployRemoteBridge(&opts, s.Backend)
	s.Require().NoError(err)
	s.Backend.Commit()
	s.bridge = bridge

	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	s.log = logrus.NewEntry(logger)

	erctransfer, err := NewTransferERC20Transactor(s.Backend, s.ERC20Address)
	s.Require().NoError(err)
	s.erctransfer = erctransfer
	s.engine = NewTransferEngine(
		s.log,
		s.Backend, s.Backend,
		*s.FundedKeys[0], *s.FundedKeys[1],
		s.bridge, s.ERC20, erctransfer,
	)
}

func (s *EngineSuite) TestOutOfBalance() {
	ctx := context.Background()
	err := s.engine.Execute(ctx, []service.Transfer{
		{
			Hash:  common.Hash{1},
			To:    common.Address{1, 1},
			Value: big.NewInt(100),
		},
	})
	s.True(errors.Is(err, service.ErrBankOutOfBalance))
}

func (s *EngineSuite) TestExceedsAllowance() {
	engine := NewTransferEngine(
		s.log,
		s.Backend, s.Backend,
		*s.FundedKeys[7], *s.FundedKeys[0],
		s.bridge, s.ERC20, s.erctransfer,
	)
	ctx := context.Background()
	err := engine.Execute(ctx, []service.Transfer{
		{
			Hash:  common.Hash{1},
			To:    common.Address{1, 1},
			Value: big.NewInt(100),
		},
	})
	s.True(service.IsErrExceedsAllowance(err))
}

func (s *EngineSuite) TestExecutedOnce() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := s.ERC20.Transfer(s.FundedKeys[1], s.FundedKeys[1].From, big.NewInt(1e18))
	s.Require().NoError(err)
	s.Backend.Commit()

	transfers := []service.Transfer{
		{
			Hash:  common.Hash{1},
			To:    common.Address{1, 1},
			Value: big.NewInt(100),
		},
		{
			Hash:  common.Hash{2},
			To:    common.Address{2, 2},
			Value: big.NewInt(100),
		},
		{
			Hash:  common.Hash{2},
			To:    common.Address{2, 2},
			Value: big.NewInt(100),
		},
	}
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

	err = s.engine.Execute(ctx, transfers)
	s.Require().NoError(err)

	for _, transfer := range transfers {
		balance, err := s.ERC20.BalanceOf(&bind.CallOpts{Context: ctx}, transfer.To)
		s.Require().NoError(err)
		s.Require().Equal(transfer.Value.Int64(), balance.Int64())

		registered, err := s.bridge.Transfers(&bind.CallOpts{Context: ctx}, transfer.Hash)
		s.Require().NoError(err)
		s.Require().True(registered.Exist)
	}
}

func (s *EngineSuite) TestReplaceNotFound() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := s.ERC20.Transfer(s.FundedKeys[1], s.FundedKeys[1].From, big.NewInt(1e18))
	s.Require().NoError(err)
	s.Backend.Commit()

	localHash := common.Hash{1}
	notfound := common.Hash{7, 7, 7}
	transfers := []service.Transfer{
		{
			Hash:  localHash,
			To:    common.Address{1, 1},
			Value: big.NewInt(100),
		},
	}

	_, err = s.bridge.Register(s.FundedKeys[0], localHash, notfound, common.Address{8}, 1)
	s.Require().NoError(err)
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

	err = s.engine.Execute(ctx, transfers)
	s.Require().NoError(err)

	for _, transfer := range transfers {
		balance, err := s.ERC20.BalanceOf(&bind.CallOpts{Context: ctx}, transfer.To)
		s.Require().NoError(err)
		s.Require().Equal(transfer.Value.Int64(), balance.Int64())

		registered, err := s.bridge.Transfers(&bind.CallOpts{Context: ctx}, transfer.Hash)
		s.Require().NoError(err)
		s.Require().True(registered.Exist)
	}
}
