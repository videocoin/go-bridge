package testtools

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/testtools/eventer"
)

type Suite struct {
	suite.Suite
	FundedKeys []*bind.TransactOpts
	Backend    *backends.SimulatedBackend
	ERC20      *erc20.ERC20
}

func (s *Suite) SetupTest() {
	alloc := core.GenesisAlloc{}
	for i := 0; i < 20; i++ {
		pkey, err := crypto.GenerateKey()
		s.Require().NoError(err)
		opts := bind.NewKeyedTransactor(pkey)
		s.FundedKeys = append(s.FundedKeys, opts)
		alloc[opts.From] = core.GenesisAccount{Balance: new(big.Int).SetUint64(^uint64(0))}
	}

	s.Backend = backends.NewSimulatedBackend(alloc, ^uint64(0))

	address, _, _, err := eventer.DeployEventer(s.FundedKeys[0], s.Backend)
	s.Require().NoError(err)
	s.Backend.Commit()
	s.ERC20, err = erc20.NewERC20(address, s.Backend)
	s.Require().NoError(err)
}

func (s *Suite) TearDownTest() {
	s.FundedKeys = nil
	s.NoError(s.Backend.Close()) // assert intentional to allow other destructors to execute
}
