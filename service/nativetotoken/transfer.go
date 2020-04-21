package nativetotoken

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/remotebridge"
	"github.com/videocoin/go-bridge/service"
)

const erc20TransferGasLimit uint64 = 100000

type Client interface {
	bind.ContractBackend
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
}

func NewTransferEngine(
	log *logrus.Entry,
	lclient, rclient Client,
	lopts, ropts bind.TransactOpts,
	rbridge *remotebridge.RemoteBridge,
	erc *erc20.ERC20,
	erctransfer *ERC20TransferTransactor,
) *TransferEngine {
	return &TransferEngine{
		log:         log,
		lclient:     lclient,
		rclient:     rclient,
		lopts:       lopts,
		txOpts:      ropts,
		remote:      rbridge,
		erc20:       erc,
		erctransfer: erctransfer,
	}
}

type TransferEngine struct {
	log *logrus.Entry

	// lclient is a client from where we transfer funds (videocoin)
	lclient Client
	lopts   bind.TransactOpts
	remote  *remotebridge.RemoteBridge

	// rclient is a client to where we transfer funds (ethereum)
	rclient     Client
	txOpts      bind.TransactOpts
	erc20       *erc20.ERC20
	erctransfer *ERC20TransferTransactor
}

func (e *TransferEngine) Execute(ctx context.Context, transfers []service.Transfer) error {
	for i := range transfers {
		var (
			transfer = &transfers[i]
			opts     = e.txOpts
		)
		opts.GasLimit = erc20TransferGasLimit
		opts.Context = ctx

		if err := e.execute(&opts, transfer); err != nil {
			return err
		}
	}
	return nil
}

func (e *TransferEngine) execute(opts *bind.TransactOpts, transfer *service.Transfer) error {
	registered, err := e.remote.Transfers(&bind.CallOpts{Context: opts.Context}, transfer.Hash)
	if err != nil {
		return err
	}
	if registered.Exist {
		known, pending, err := e.rclient.TransactionByHash(opts.Context, common.Hash(registered.Hash))
		if known != nil || pending {
			e.log.Debugf("transfer 0x%x is known by remote blockchain", registered.Hash)
			err := e.waitMined(opts.Context, transfer, known)
			if err == nil {
				return nil
			}
			if !errors.Is(err, service.ErrTransactionReverted) {
				return nil
			}
			e.log.Debugf("transfer 0x%x is reverted. will be resubmitted.", registered.Hash)
		} else {
			if err != nil && err != ethereum.NotFound {
				return err
			}
			e.log.Debugf("transfer 0x%x is missing in remote blockchain. will be resubmitted",
				registered.Hash)
		}
	}

	balance, err := e.erc20.BalanceOf(&bind.CallOpts{Context: opts.Context}, opts.From)
	if err != nil {
		return err
	}
	metric, _ := new(big.Float).SetInt(balance).Float64()
	service.TokenBankBalanceGauge.Set(metric)
	// less or equal to account for additional gas cost
	if balance.Cmp(transfer.Value) <= 0 {
		return fmt.Errorf("%w: not enough funds on bank 0x%x to make a transfer for %v",
			service.ErrBankOutOfBalance, opts.From, transfer.Value,
		)
	}

	tx, err := e.erctransfer.Create(opts, transfer.To, transfer.Value)
	if err != nil {
		return err
	}

	if err := e.register(opts.Context, transfer.Hash, tx, registered.Exist); err != nil {
		return err
	}

	if err := e.rclient.SendTransaction(opts.Context, tx); err != nil {
		return err
	}
	if err := e.waitMined(opts.Context, transfer, tx); err != nil {
		return err
	}
	return nil
}

func (e *TransferEngine) waitMined(ctx context.Context, transfer *service.Transfer, tx *types.Transaction) error {
	e.log.Debugf("waiting for 0x%x to get mined as 0x%x", transfer.Hash, tx.Hash())
	receipt, err := bind.WaitMined(ctx, e.rclient, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("%w: failed to execute transfer. recipient 0x%x. value %v",
			service.ErrTransactionReverted, transfer.To, transfer.Value)
	}
	e.log.Infof("transfer was completed. to 0x%x. value %v. gas used %d",
		transfer.To, transfer.Value, receipt.GasUsed)
	return nil
}

func (e *TransferEngine) register(ctx context.Context, hash common.Hash, rtx *types.Transaction, exists bool) error {
	var (
		tx  *types.Transaction
		err error
	)
	e.log.Debugf("registering transfer 0x%x -> 0x%x", hash, rtx.Hash())
	if !exists {
		opts := e.lopts
		tx, err = e.remote.Register(&opts,
			[32]byte(hash),
			[32]byte(rtx.Hash()),
			e.txOpts.From, rtx.Nonce(),
		)
		if err != nil {
			return err
		}
	} else {
		opts := e.lopts
		tx, err = e.remote.Update(&opts, [32]byte(hash), [32]byte(rtx.Hash()))
		if err != nil {
			return err
		}
	}
	_, err = bind.WaitMined(ctx, e.lclient, tx)
	return err
}

func (e *TransferEngine) Get(ctx context.Context) (*big.Int, error) {
	return e.remote.GetLastBlock(&bind.CallOpts{Context: ctx})
}

func (e *TransferEngine) Set(ctx context.Context, number *big.Int) error {
	opts := e.lopts
	tx, err := e.remote.SetLastBlock(&opts, number)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, e.lclient, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("failed to update last block to %v", number)
	}
	return nil
}
