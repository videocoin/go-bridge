package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/homebridge"
)

type AddressesSource interface {
	All(context.Context) ([]common.Address, error)
}

type LocalClient interface {
	bind.DeployBackend
}

type ForeignClient interface {
	HeaderByNumber(context.Context, *big.Int) (*types.Header, error)
}

func NewService(
	log *logrus.Entry,
	lclient LocalClient, fclient ForeignClient,
	erc20 *erc20.ERC20,
	banks, workers AddressesSource,
	bridgeOwner bind.TransactOpts,
	bridge *homebridge.HomeBridge,
	blockDelay, scanStep *big.Int,
) *Service {
	return &Service{
		log:        log,
		lclient:    lclient,
		fclient:    fclient,
		erc20:      erc20,
		banks:      banks,
		workers:    workers,
		txOpts:     bridgeOwner,
		bridge:     bridge,
		blockDelay: blockDelay,
		scanStep:   scanStep,
	}
}

type Service struct {
	log *logrus.Entry

	lclient LocalClient
	fclient ForeignClient

	erc20 *erc20.ERC20

	banks, workers AddressesSource

	txOpts bind.TransactOpts
	bridge *homebridge.HomeBridge

	blockDelay, scanStep *big.Int
}

func (s *Service) getRange(ctx context.Context) (*big.Int, *big.Int, error) {
	start, err := s.bridge.GetLastBlock(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, err
	}
	head, err := s.fclient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	safe := head.Number.Sub(head.Number, s.blockDelay)
	end := new(big.Int).Add(start, s.scanStep)
	// IF START=1, END=5, SAFE=10 RETURN 1, 5
	if safe.Cmp(end) >= 0 {
		return start, end, nil
	}
	// IF START=7, END=12, SAFE=10 RETURN 7, 10
	return start, safe, nil
}

func (s *Service) Run(ctx context.Context) error {
	start, end, err := s.getRange(ctx)
	if err != nil {
		s.log.Errorf("failed to to get scan range. err %v", err)
		return err
	}
	if start.Cmp(end) >= 0 {
		s.log.Debugf("nothing to scan. start %v. end %v", start, end)
		return nil
	}
	if err := s.scan(ctx, start, end); err != nil {
		s.log.Errorf("scan failed. start %v. end %v. err %v", start, end, err)
		return err
	}
	return nil
}

func (s *Service) scan(ctx context.Context, start, end *big.Int) error {
	s.log.Debugf("running scan for erc20 transfers. start %v. end %v", start, end)
	from, err := s.workers.All(ctx)
	if err != nil {
		return err
	}
	to, err := s.banks.All(ctx)
	if err != nil {
		return err
	}
	if len(to) == 0 {
		s.log.Debugf("no configured banks, skipping scan")
		return nil
	}
	if from == nil {
		s.log.Debugf("list of senders is nil. skipping scan")
		return nil
	}
	if len(from) == 0 {
		s.log.Infof("transfers from all senders are accepted for current scan. start %v. end %v.", start, end)
	}

	endu := end.Uint64()
	filter, err := s.erc20.FilterTransfer(
		&bind.FilterOpts{
			Start: start.Uint64(),
			End:   &endu,
		},
		from, to,
	)
	if err != nil {
		return err
	}
	for filter.Next() {
		s.log.Debugf("found transfer 0x%X. from 0x%x. to 0x%x. value %v",
			filter.Event.Raw.TxHash, filter.Event.From, filter.Event.To, filter.Event.Value)
		executed, err := s.bridge.ExecutedTransfers(
			&bind.CallOpts{Context: ctx},
			[32]byte(filter.Event.Raw.TxHash))
		if err != nil {
			return err
		}
		if executed {
			s.log.Debugf("transfer 0x%x already executed", filter.Event.Raw.TxHash)
			continue
		}
		opts := s.txOpts
		tx, err := s.bridge.ExecuteTransfer(&opts,
			filter.Event.From,
			filter.Event.Value,
			[32]byte(filter.Event.Raw.TxHash))
		if err != nil {
			return err
		}
		receipt, err := bind.WaitMined(ctx, s.lclient, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("failed to execute transfer. recipient 0x%x. value %v",
				filter.Event.From, filter.Event.Value)
		}
		s.log.Infof("executed transfer. to 0x%x. value %v",
			filter.Event.From, filter.Event.Value)
	}
	if filter.Error() != nil {
		s.log.Errorf("filtering failed. err %v", filter.Error())
		return filter.Error()
	}
	opts := s.txOpts
	tx, err := s.bridge.SetLastBlock(&opts, end)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, s.lclient, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("failed to update last block to %v", end)
	}
	return err
}
