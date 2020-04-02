package nativetotoken

import (
	"context"
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

const erc20TransferGasLimit uint64 = 50000

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
) *TransferEngine {
	return &TransferEngine{
		log:     log,
		lclient: lclient,
		rclient: rclient,
		lopts:   lopts,
		txOpts:  ropts,
		remote:  rbridge,
		erc20:   erc,
	}
}

type TransferEngine struct {
	log *logrus.Entry

	// lclient is a client from where we transfer funds (videocoin)
	lclient Client
	lopts   bind.TransactOpts
	remote  *remotebridge.RemoteBridge

	// rclient is a client to where we transfer funds (ethereum)
	rclient Client
	txOpts  bind.TransactOpts
	erc20   *erc20.ERC20
}

func (e *TransferEngine) Execute(ctx context.Context, transfers []service.Transfer) error {
	for i := range transfers {
		var (
			transfer = &transfers[i]
			opts     = e.txOpts
		)
		opts.GasLimit = erc20TransferGasLimit
		opts.Context = ctx

		registered, err := e.remote.Transfers(&bind.CallOpts{Context: ctx}, transfer.Hash)
		if err != nil {
			return err
		}
		// TODO handle condition when signer changes but tx is missing
		if !registered.Exist {
			nonce, err := e.rclient.PendingNonceAt(ctx, opts.From)
			if err != nil {
				return err
			}
			e.log.Debugf("transfer 0x%x observed for the first time. nonce %d", transfer.Hash, nonce)
			opts.Nonce = new(big.Int).SetUint64(nonce)
		} else {
			known, pending, err := e.rclient.TransactionByHash(ctx, common.Hash(registered.Hash))
			// TODO if pending => bind.WaitMined
			if known != nil || pending {
				e.log.Debugf("transfer 0x%x is known by remote blockchain", registered.Hash)
				continue
			}
			if err != nil && err != ethereum.NotFound {
				return err
			}
			e.log.Debugf("transfer 0x%x is missing in remote blockchain. will be resubmitted with nonce %d",
				registered.Hash, registered.Nonce)
			// we failed to send transaction, need to retry
			opts.Nonce = new(big.Int).SetUint64(registered.Nonce)

			// TODO handle condition when transfer is known, mined but failed
			// in such case we should retry transaction
		}

		balance, err := e.erc20.BalanceOf(&bind.CallOpts{Context: ctx}, opts.From)
		if err != nil {
			return err
		}
		// less or equal to account for additional gas cost
		if balance.Cmp(transfer.Value) <= 0 {
			return fmt.Errorf("%w: not enough funds on bank 0x%x to make a transfer for %v",
				service.ErrBankOutOfBalance, opts.From, transfer.Value,
			)
		}

		tx, err := e.erc20.Transfer(&opts, transfer.To, transfer.Value)
		if err != nil {
			return err
		}

		// TODO this should be done before sending transfer
		// requires exposing transaction object before it is sent to the network
		if err := e.register(ctx, transfer.Hash, tx, registered.Exist); err != nil {
			return err
		}

		e.log.Debugf("waiting for 0x%x to get mined as 0x%x", transfer.Hash, tx.Hash())
		receipt, err := bind.WaitMined(ctx, e.rclient, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("failed to execute transfer. recipient 0x%x. value %v",
				transfer.To, transfer.Value)
		}
		e.log.Infof("executed transfer. to 0x%x. value %v. gas used %d",
			transfer.To, transfer.Value, receipt.GasUsed)
	}
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
