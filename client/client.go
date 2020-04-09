package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/nativebridge"
	"github.com/videocoin/go-bridge/nativeproxy"
	"github.com/videocoin/go-bridge/remotebridge"
)

type Config struct {
	ProxyAddress         common.Address
	ERC20Address         common.Address
	LocalBridgeAddress   common.Address
	ForeignBridgeAddress common.Address
}

func Dial(local, foreign ETHBackend, conf Config) (*Client, error) {
	erc, err := erc20.NewERC20(conf.ERC20Address, foreign)
	if err != nil {
		return nil, err
	}
	proxy, err := nativeproxy.NewNativeProxy(conf.ProxyAddress, local)
	if err != nil {
		return nil, err
	}
	localBridge, err := nativebridge.NewNativeBridge(conf.LocalBridgeAddress, local)
	if err != nil {
		return nil, err
	}
	// foreignBridge is using local network to keep track of the remote transactions
	// somewhat confusing, but done so intentionally to avoid additional gas cost on ethereum mainnet
	foreignBridge, err := remotebridge.NewRemoteBridge(conf.ForeignBridgeAddress, local)
	if err != nil {
		return nil, err
	}
	return NewClient(local, foreign, proxy, erc, localBridge, foreignBridge), nil
}

func NewClient(
	local, foreign ETHBackend,
	proxy *nativeproxy.NativeProxy,
	erc *erc20.ERC20,
	localBridge *nativebridge.NativeBridge,
	foreignBridge *remotebridge.RemoteBridge,
) *Client {
	return &Client{
		local:         local,
		foreign:       foreign,
		proxy:         proxy,
		localBridge:   localBridge,
		foreignBridge: foreignBridge,
		erc:           erc,
	}
}

type ETHBackend interface {
	bind.ContractBackend
	TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error)
	HeaderByNumber(context.Context, *big.Int) (*types.Header, error)
	TransactionByHash(context.Context, common.Hash) (*types.Transaction, bool, error)
}

type Client struct {
	local, foreign ETHBackend

	proxy       *nativeproxy.NativeProxy
	localBridge *nativebridge.NativeBridge

	foreignBridge *remotebridge.RemoteBridge
	erc           *erc20.ERC20
}

// WaitDeposit is the same as Deposit but it additioally waits until ERC transfer is bridged on local chain.
// WaitDeposit will exit only if context deadline is reached or bridged transfer is found. Ensure that context
// deadline is set to reasonable time (5m for mainnet will be plenty).
func (c *Client) WaitDeposit(ctx context.Context, key *ecdsa.PrivateKey, bank common.Address, amount *big.Int) (info TransferInfo, err error) {
	header, err := c.local.HeaderByNumber(ctx, nil)
	if err != nil {
		return info, err
	}

	hash, err := c.Deposit(ctx, key, bank, amount)
	if err != nil {
		return info, err
	}
	info.ForeignTxHash = hash
	// bridge waits for sufficient number of confirmations (blocks on top) before creating bridged transfer on local chain.
	// waiting may take 2-3 minutes.
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return info, ctx.Err()
		case <-ticker.C:
			ok, err := c.IsDepositBridged(ctx, hash)
			if err != nil {
				continue
			}
			if !ok {
				continue
			}
			events, err := c.localBridge.FilterTransferBridged(&bind.FilterOpts{
				Context: ctx,
				Start:   header.Number.Uint64(),
			}, nil, nil, [][32]byte{hash})
			if err != nil {
				continue
			}
			for events.Next() {
				info.LocalTxHash = events.Event.Raw.TxHash
				_ = events.Close()
				return info, nil
			}
			_ = events.Close()
		}
	}
}

// Deposit makes an ERC20 transfer using configured erc20 token contract. Then client
// should wait until transaction with returned hash is created in NativeBridge.
func (c *Client) Deposit(ctx context.Context, key *ecdsa.PrivateKey, bank common.Address, amount *big.Int) (hash common.Hash, err error) {
	opts := bind.NewKeyedTransactor(key)
	opts.Context = ctx
	tx, err := c.erc.Transfer(opts, bank, amount)
	if err != nil {
		return hash, err
	}
	// wait for erc20 on foreign chain to succeed
	receipt, err := bind.WaitMined(ctx, c.foreign, tx)
	if err != nil {
		return hash, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return hash, fmt.Errorf("token transfer 0x%x failed to execute", tx.Hash())
	}
	return tx.Hash(), nil
}

// IsDepositBridged returns true if ERC transfer hash was bridged.
func (c *Client) IsDepositBridged(ctx context.Context, hash common.Hash) (bool, error) {
	return c.localBridge.Transfers(&bind.CallOpts{Context: ctx}, hash)
}

func (c *Client) Withdraw(ctx context.Context, key *ecdsa.PrivateKey, bank common.Address, amount *big.Int) (hash common.Hash, err error) {
	opts := bind.NewKeyedTransactor(key)
	opts.Context = ctx
	opts.Value = amount

	tx, err := c.proxy.Proxy(opts, bank)
	if err != nil {
		return hash, err
	}
	receipt, err := bind.WaitMined(ctx, c.local, tx)
	if err != nil {
		return hash, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return hash, fmt.Errorf("native transfer 0x%x using proxy failed", tx.Hash())
	}
	return tx.Hash(), nil
}

func (c *Client) IsWithdrawBridged(ctx context.Context, hash common.Hash) (bool, error) {
	transfer, err := c.foreignBridge.Transfers(&bind.CallOpts{Context: ctx}, [32]byte(hash))
	if err != nil {
		return false, err
	}
	_, pending, err := c.foreign.TransactionByHash(ctx, transfer.Hash)
	if err != nil {
		return false, err
	}
	return !pending, nil
}

func (c *Client) WaitWithdraw(ctx context.Context, key *ecdsa.PrivateKey, bank common.Address, amount *big.Int) (info TransferInfo, err error) {
	hash, err := c.Withdraw(ctx, key, bank, amount)
	if err != nil {
		return info, err
	}
	info.LocalTxHash = hash

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return info, ctx.Err()
		case <-ticker.C:
			transfer, err := c.foreignBridge.Transfers(&bind.CallOpts{Context: ctx}, [32]byte(hash))
			if err != nil {
				continue
			}
			tx, pending, err := c.foreign.TransactionByHash(ctx, transfer.Hash)
			if err != nil {
				continue
			}
			if !pending {
				info.ForeignTxHash = tx.Hash()
				return info, nil
			}
		}
	}
}
