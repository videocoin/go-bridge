package nativetotoken

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/videocoin/go-bridge/nativeproxy"
	"github.com/videocoin/go-bridge/testtools"
)

func TestNativeTransferAccess(t *testing.T) {
	suite.Run(t, new(TransfersAccessSuite))
}

type TransfersAccessSuite struct {
	testtools.Suite

	access *NativeTransfersAccess
	proxy  *nativeproxy.NativeProxy
}

func (s *TransfersAccessSuite) SetupTest() {
	s.Suite.SetupTest()

	_, _, proxy, err := nativeproxy.DeployNativeProxy(s.FundedKeys[0], s.Backend)
	s.Require().NoError(err)
	s.Backend.Commit()
	s.proxy = proxy

	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	log := logrus.NewEntry(logger)
	s.access = NewNativeTransferAccess(log, proxy)
}

func (s *TransfersAccessSuite) TestTransfers() {
	ctx := context.Background()

	to := common.Address{5, 5, 5}
	opts := *s.FundedKeys[1]
	opts.Value = big.NewInt(100)
	opts.Context = ctx
	tx, err := s.proxy.Proxy(&opts, to)
	s.Require().NoError(err)
	s.Backend.Commit()

	balance, err := s.Backend.BalanceAt(ctx, to, nil)
	s.Require().NoError(err)
	s.Require().Equal(opts.Value.Int64(), balance.Int64())

	transfers, err := s.access.Transfers(ctx, []common.Address{to}, 0, 10)
	s.Require().NoError(err)
	s.Require().Len(transfers, 1)
	s.Require().Equal(opts.Value.Int64(), transfers[0].Value.Int64())
	s.Require().Equal(tx.Hash(), transfers[0].Hash)
}
