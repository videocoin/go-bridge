package service

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

type AddressesSource interface {
	All(context.Context) ([]common.Address, error)
}

type Client interface {
	HeaderByNumber(context.Context, *big.Int) (*types.Header, error)
}

type Transfer struct {
	Hash  common.Hash
	To    common.Address
	Value *big.Int
}

type TransferEngine interface {
	Execute(context.Context, []Transfer) error
}

type LastBlockAccess interface {
	Get(context.Context) (*big.Int, error)
	Set(context.Context, *big.Int) error
}

type TransfersAccess interface {
	Transfers(ctx context.Context, to []common.Address, start, end uint64) ([]Transfer, error)
}

func NewService(
	log *logrus.Entry,
	client Client,
	engine TransferEngine,
	lastBlock LastBlockAccess,
	transfers TransfersAccess,
	banks AddressesSource,
	blockDelay, scanStep *big.Int,
) *Service {
	return &Service{
		log:        log,
		client:     client,
		engine:     engine,
		banks:      banks,
		lastBlock:  lastBlock,
		transfers:  transfers,
		blockDelay: blockDelay,
		scanStep:   scanStep,
	}
}

type Service struct {
	log *logrus.Entry

	client Client

	banks AddressesSource

	engine    TransferEngine
	lastBlock LastBlockAccess
	transfers TransfersAccess

	blockDelay, scanStep *big.Int
}

func (s *Service) getRange(ctx context.Context) (*big.Int, *big.Int, error) {
	start, err := s.lastBlock.Get(ctx)
	if err != nil {
		return nil, nil, err
	}
	head, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	safe := new(big.Int).Sub(head.Number, s.blockDelay)
	end := new(big.Int).Add(start, s.scanStep)
	// IF START=1, END=5, SAFE=10 RETURN 1, 5
	if safe.Cmp(end) >= 0 {
		return start, end, nil
	}
	// IF START=7, END=12, SAFE=10 RETURN 7, 10
	return start, safe, nil
}

func (s *Service) Run(ctx context.Context) error {
	UpGauge.Set(1)
	if err := s.run(ctx); err != nil {
		FailingGauge.Set(1)
		return err
	}
	return nil
}

func (s *Service) run(ctx context.Context) error {
	start, end, err := s.getRange(ctx)
	if err != nil {
		s.log.Errorf("failed to to get scan range. err %v", err)
		return err
	}
	if start.Cmp(end) >= 0 {
		s.log.Debugf("nothing to scan. start %v. end %v", start, end)
		return nil
	}
	s.log.Debugf("running scan for transfers. start %v. end %v", start, end)
	to, err := s.banks.All(ctx)
	if err != nil {
		return err
	}
	if len(to) == 0 {
		s.log.Debugf("no configured banks, skipping scan")
		return nil
	}

	transfers, err := s.transfers.Transfers(ctx, to, start.Uint64(), end.Uint64())
	if err != nil {
		if errors.Is(err, ErrBankOutOfBalance) {
			OutOfBalance.Set(1)
		}
		if IsErrExceedsAllowance(err) {
			GasExceedsAllowance.Set(1)
		}
		return err
	}
	if err := s.engine.Execute(ctx, transfers); err != nil {
		return err
	}
	if err := s.lastBlock.Set(ctx, end); err != nil {
		return fmt.Errorf("failed to set known block %v: %v", end, err)
	}
	return nil
}
