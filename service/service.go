package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/homebridge"
)

type WhitelistService interface {
	GetWhitelisted(context.Context) ([]common.Address, error)
}

type Service struct {
	log *logrus.Entry

	lclient, fclient *ethclient.Client

	erc20 *erc20.ERC20

	bankAddress []common.Address
	txOpts      bind.TransactOpts
	bridge      *homebridge.HomeBridge

	whitelist WhitelistService

	blockDelay *big.Int

	scanStep uint64
}

func (s *Service) Run(ctx context.Context) error {
	start, err := s.bridge.GetLastBlock(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}
	head, err := s.fclient.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	end := head.Number
	end = end.Sub(end, s.blockDelay)
	if end.Cmp(start) == -1 {
		return nil
	}
	s.log.Debugf("running scan for erc20 transfer between %v and %v", start, end)
	whitelist, err := s.whitelist.GetWhitelisted(ctx)
	if err != nil {
		return err
	}
	filter, err := s.erc20.FilterTransfer(&bind.FilterOpts{}, whitelist, s.bankAddress)
	if err != nil {
		return err
	}
	for filter.Next() {
		opts := s.txOpts
		tx, err := s.bridge.ExecuteTransfer(&opts, filter.Event.From, filter.Event.Value, [32]byte(filter.Event.Raw.TxHash))
		if err != nil {
			return err
		}
		receipt, err := bind.WaitMined(ctx, s.lclient, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("failed to execute transfer. recipient %v. value %v", filter.Event.From, filter.Event.Value)
		}
		s.log.Infof("executed transfer to %v. value %v", filter.Event.From, filter.Event.Value)
	}
	if filter.Error() != nil {
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
