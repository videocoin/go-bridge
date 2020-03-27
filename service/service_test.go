package service

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/videocoin/go-bridge/homebridge"
	"github.com/videocoin/go-bridge/testtools"
)

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	testtools.Suite

	bridge *homebridge.HomeBridge
	log    *logrus.Entry
}

func (s *ServiceSuite) SetupTest() {
	s.Suite.SetupTest()
	opts := *s.FundedKeys[0]
	_, _, bridge, err := homebridge.DeployHomeBridge(
		&opts, s.Backend,
		big.NewInt(1e18),
		big.NewInt(1e18),
		big.NewInt(1e18),
		big.NewInt(0),
	)
	s.Require().NoError(err)
	s.Backend.Commit()
	s.bridge = bridge

	opts = *s.FundedKeys[0]
	opts.Value = big.NewInt(1e16)

	raw := homebridge.HomeBridgeRaw{Contract: s.bridge}
	_, err = raw.Transfer(&opts)
	s.Require().NoError(err)

	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	s.log = logrus.NewEntry(logger)
}

func (s *ServiceSuite) TestServiceExecuteAll() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := NewService(s.log, s.Backend, s.Backend,
		s.ERC20, StaticSource{s.FundedKeys[0].From}, NilSource{},
		*s.FundedKeys[0], s.bridge, big.NewInt(0), big.NewInt(10),
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
		executed, err := s.bridge.ExecutedTransfers(&bind.CallOpts{Context: ctx}, [32]byte(tx))
		s.Require().NoError(err)
		s.Require().True(executed)
	}
}

func (s *ServiceSuite) TestServiceExecuteWhitelisted() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := NewService(s.log, s.Backend, s.Backend,
		s.ERC20, StaticSource{s.FundedKeys[0].From}, StaticSource{s.FundedKeys[1].From},
		*s.FundedKeys[0], s.bridge, big.NewInt(0), big.NewInt(10),
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

	executed, err := s.bridge.ExecutedTransfers(&bind.CallOpts{Context: ctx}, [32]byte(txs[0]))
	s.Require().NoError(err)
	s.Require().True(executed)
	for _, tx := range txs[1:] {
		executed, err := s.bridge.ExecutedTransfers(&bind.CallOpts{Context: ctx}, [32]byte(tx))
		s.Require().NoError(err)
		s.Require().False(executed)
	}
}

func (s *ServiceSuite) TestServiceNothingExecuted() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := NewService(s.log, s.Backend, s.Backend,
		s.ERC20, StaticSource{}, StaticSource{},
		*s.FundedKeys[0], s.bridge, big.NewInt(0), big.NewInt(10),
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
		executed, err := s.bridge.ExecutedTransfers(&bind.CallOpts{Context: ctx}, [32]byte(tx))
		s.Require().NoError(err)
		s.Require().False(executed)
	}
}
