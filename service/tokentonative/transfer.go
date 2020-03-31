package tokentonative

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/go-bridge/nativebridge"
	"github.com/videocoin/go-bridge/service"
)

func NewTransferEngine(
	log *logrus.Entry,
	client Client,
	txOpts bind.TransactOpts,
	bridge *nativebridge.NativeBridge) *TransferEngine {
	return &TransferEngine{
		log:    log,
		client: client,
		txOpts: txOpts,
		bridge: bridge,
	}
}

type TransferEngine struct {
	log    *logrus.Entry
	client Client

	txOpts bind.TransactOpts
	bridge *nativebridge.NativeBridge
}

func (e *TransferEngine) Execute(ctx context.Context, transfers []service.Transfer) error {
	for i := range transfers {
		transfer := &transfers[i]
		executed, err := e.bridge.Transfers(
			&bind.CallOpts{Context: ctx},
			[32]byte(transfer.Hash))
		if err != nil {
			return err
		}
		if executed {
			e.log.Debugf("transfer 0x%x already executed", transfer.Hash)
			return nil
		}

		opts := e.txOpts
		opts.Context = ctx
		balance, err := e.client.BalanceAt(ctx, opts.From, nil)
		if err != nil {
			return err
		}
		// less or equal to account for additional gas cost
		if balance.Cmp(transfer.Value) <= 0 {
			return fmt.Errorf("not enough funds on bank 0x%x to make a transfer for %v",
				opts.From, transfer.Value,
			)
		}

		opts.Value = transfer.Value
		tx, err := e.bridge.Transfer(&opts,
			transfer.To,
			[32]byte(transfer.Hash))
		if err != nil {
			return err
		}
		receipt, err := bind.WaitMined(ctx, e.client, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("failed to execute transfer. recipient 0x%x. value %v",
				transfer.To, transfer.Value)
		}
		e.log.Infof("executed transfer. to 0x%x. value %v. gas used %d",
			transfer.To, transfer.Value, receipt.CumulativeGasUsed)
	}
	return nil
}

func (e *TransferEngine) Get(ctx context.Context) (*big.Int, error) {
	return e.bridge.GetLastBlock(&bind.CallOpts{Context: ctx})
}

func (e *TransferEngine) Set(ctx context.Context, number *big.Int) error {
	opts := e.txOpts
	tx, err := e.bridge.SetLastBlock(&opts, number)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, e.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("failed to update last block to %v", number)
	}
	return nil
}
