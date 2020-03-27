package service

import (
	"context"
	"math/big"
	"testing"

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
	logger.SetLevel(logrus.DebugLevel)
	s.log = logrus.NewEntry(logger)
}

func (s *ServiceSuite) TestServiceHappyCase() {
	ctx, cancel := context.WithCancel(context.Background())

	service := NewService(s.log, s.Backend, s.Backend,
		s.ERC20, StaticSource{}, StaticSource{}, *s.FundedKeys[0], s.bridge, big.NewInt(0), big.NewInt(10))

	for i := 1; i < len(s.FundedKeys); i++ {
		opts := *s.FundedKeys[0]
		_, err := s.ERC20.Transfer(&opts, s.FundedKeys[i].From, big.NewInt(10))
		s.Require().NoError(err)
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
	cancel()
}
