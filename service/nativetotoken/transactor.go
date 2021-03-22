package nativetotoken

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/videocoin/go-bridge/erc20"
)

func NewTransferERC20Transactor(client bind.ContractBackend, address common.Address) (*ERC20TransferTransactor, error) {
	parsed, err := abi.JSON(strings.NewReader(erc20.ERC20ABI))
	if err != nil {
		return nil, err
	}
	return &ERC20TransferTransactor{
		client:  client,
		address: address,
		abi:     parsed,
	}, nil
}

type ERC20TransferTransactor struct {
	client  bind.ContractBackend
	address common.Address
	abi     abi.ABI
}

func (c *ERC20TransferTransactor) Create(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	input, err := c.abi.Pack("transfer", to, amount)
	if err != nil {
		return nil, err
	}
	return c.create(opts, input)
}

func (c *ERC20TransferTransactor) create(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	var err error
	// Ensure a valid value field and resolve the account nonce
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = c.client.PendingNonceAt(opts.Context, opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}
	// Figure out the gas allowance and gas price values
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = c.client.SuggestGasPrice(opts.Context)
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas price: %v", err)
		}
	}
	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		// Gas estimation cannot succeed without code for method invocations

		if code, err := c.client.PendingCodeAt(opts.Context, c.address); err != nil {
			return nil, err
		} else if len(code) == 0 {
			return nil, bind.ErrNoCode
		}

		// If the contract surely has code (or code is not needed), estimate the transaction
		msg := ethereum.CallMsg{From: opts.From, To: &c.address, GasPrice: gasPrice, Value: value, Data: input}
		gasLimit, err = c.client.EstimateGas(opts.Context, msg)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
		}
	}
	rawTx := types.NewTransaction(nonce, c.address, value, gasLimit, gasPrice, input)
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
