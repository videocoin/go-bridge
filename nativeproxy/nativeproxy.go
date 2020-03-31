// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativeproxy

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NativeProxyABI is the input ABI used to generate the binding from.
const NativeProxyABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferProxied\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"proxy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// NativeProxyBin is the compiled bytecode used for deploying new contracts.
var NativeProxyBin = "0x608060405234801561001057600080fd5b506101e1806100206000396000f3fe60806040526004361061001e5760003560e01c806306713c3e14610023575b600080fd5b6100656004803603602081101561003957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610067565b005b60008173ffffffffffffffffffffffffffffffffffffffff163460405180600001905060006040518083038185875af1925050503d80600081146100c7576040519150601f19603f3d011682016040523d82523d6000602084013e6100cc565b606091505b5050905080610143576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f7472616e736665722070726f7879206661696c6564000000000000000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ffb13b2ffecafe583990a6ad9a7f0ee9e37c92e5932bb33b2491e7b4617600b99346040518082815260200191505060405180910390a3505056fea265627a7a723158206f193ead2db638004521c7c0611eddedf547bb2ea7a5d169b93cc5ce0a4983bc64736f6c63430005100032"

// DeployNativeProxy deploys a new Ethereum contract, binding an instance of NativeProxy to it.
func DeployNativeProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NativeProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(NativeProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NativeProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeProxy{NativeProxyCaller: NativeProxyCaller{contract: contract}, NativeProxyTransactor: NativeProxyTransactor{contract: contract}, NativeProxyFilterer: NativeProxyFilterer{contract: contract}}, nil
}

// NativeProxy is an auto generated Go binding around an Ethereum contract.
type NativeProxy struct {
	NativeProxyCaller     // Read-only binding to the contract
	NativeProxyTransactor // Write-only binding to the contract
	NativeProxyFilterer   // Log filterer for contract events
}

// NativeProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeProxySession struct {
	Contract     *NativeProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeProxyCallerSession struct {
	Contract *NativeProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NativeProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeProxyTransactorSession struct {
	Contract     *NativeProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NativeProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeProxyRaw struct {
	Contract *NativeProxy // Generic contract binding to access the raw methods on
}

// NativeProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeProxyCallerRaw struct {
	Contract *NativeProxyCaller // Generic read-only contract binding to access the raw methods on
}

// NativeProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeProxyTransactorRaw struct {
	Contract *NativeProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeProxy creates a new instance of NativeProxy, bound to a specific deployed contract.
func NewNativeProxy(address common.Address, backend bind.ContractBackend) (*NativeProxy, error) {
	contract, err := bindNativeProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeProxy{NativeProxyCaller: NativeProxyCaller{contract: contract}, NativeProxyTransactor: NativeProxyTransactor{contract: contract}, NativeProxyFilterer: NativeProxyFilterer{contract: contract}}, nil
}

// NewNativeProxyCaller creates a new read-only instance of NativeProxy, bound to a specific deployed contract.
func NewNativeProxyCaller(address common.Address, caller bind.ContractCaller) (*NativeProxyCaller, error) {
	contract, err := bindNativeProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeProxyCaller{contract: contract}, nil
}

// NewNativeProxyTransactor creates a new write-only instance of NativeProxy, bound to a specific deployed contract.
func NewNativeProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeProxyTransactor, error) {
	contract, err := bindNativeProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeProxyTransactor{contract: contract}, nil
}

// NewNativeProxyFilterer creates a new log filterer instance of NativeProxy, bound to a specific deployed contract.
func NewNativeProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeProxyFilterer, error) {
	contract, err := bindNativeProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeProxyFilterer{contract: contract}, nil
}

// bindNativeProxy binds a generic wrapper to an already deployed contract.
func bindNativeProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NativeProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeProxy *NativeProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NativeProxy.Contract.NativeProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeProxy *NativeProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeProxy.Contract.NativeProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeProxy *NativeProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeProxy.Contract.NativeProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeProxy *NativeProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NativeProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeProxy *NativeProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeProxy *NativeProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeProxy.Contract.contract.Transact(opts, method, params...)
}

// Proxy is a paid mutator transaction binding the contract method 0x06713c3e.
//
// Solidity: function proxy(address to) returns()
func (_NativeProxy *NativeProxyTransactor) Proxy(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NativeProxy.contract.Transact(opts, "proxy", to)
}

// Proxy is a paid mutator transaction binding the contract method 0x06713c3e.
//
// Solidity: function proxy(address to) returns()
func (_NativeProxy *NativeProxySession) Proxy(to common.Address) (*types.Transaction, error) {
	return _NativeProxy.Contract.Proxy(&_NativeProxy.TransactOpts, to)
}

// Proxy is a paid mutator transaction binding the contract method 0x06713c3e.
//
// Solidity: function proxy(address to) returns()
func (_NativeProxy *NativeProxyTransactorSession) Proxy(to common.Address) (*types.Transaction, error) {
	return _NativeProxy.Contract.Proxy(&_NativeProxy.TransactOpts, to)
}

// NativeProxyTransferProxiedIterator is returned from FilterTransferProxied and is used to iterate over the raw logs and unpacked data for TransferProxied events raised by the NativeProxy contract.
type NativeProxyTransferProxiedIterator struct {
	Event *NativeProxyTransferProxied // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeProxyTransferProxiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeProxyTransferProxied)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeProxyTransferProxied)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeProxyTransferProxiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeProxyTransferProxiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeProxyTransferProxied represents a TransferProxied event raised by the NativeProxy contract.
type NativeProxyTransferProxied struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferProxied is a free log retrieval operation binding the contract event 0xfb13b2ffecafe583990a6ad9a7f0ee9e37c92e5932bb33b2491e7b4617600b99.
//
// Solidity: event TransferProxied(address indexed from, address indexed to, uint256 value)
func (_NativeProxy *NativeProxyFilterer) FilterTransferProxied(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeProxyTransferProxiedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeProxy.contract.FilterLogs(opts, "TransferProxied", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeProxyTransferProxiedIterator{contract: _NativeProxy.contract, event: "TransferProxied", logs: logs, sub: sub}, nil
}

// WatchTransferProxied is a free log subscription operation binding the contract event 0xfb13b2ffecafe583990a6ad9a7f0ee9e37c92e5932bb33b2491e7b4617600b99.
//
// Solidity: event TransferProxied(address indexed from, address indexed to, uint256 value)
func (_NativeProxy *NativeProxyFilterer) WatchTransferProxied(opts *bind.WatchOpts, sink chan<- *NativeProxyTransferProxied, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeProxy.contract.WatchLogs(opts, "TransferProxied", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeProxyTransferProxied)
				if err := _NativeProxy.contract.UnpackLog(event, "TransferProxied", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferProxied is a log parse operation binding the contract event 0xfb13b2ffecafe583990a6ad9a7f0ee9e37c92e5932bb33b2491e7b4617600b99.
//
// Solidity: event TransferProxied(address indexed from, address indexed to, uint256 value)
func (_NativeProxy *NativeProxyFilterer) ParseTransferProxied(log types.Log) (*NativeProxyTransferProxied, error) {
	event := new(NativeProxyTransferProxied)
	if err := _NativeProxy.contract.UnpackLog(event, "TransferProxied", log); err != nil {
		return nil, err
	}
	return event, nil
}
