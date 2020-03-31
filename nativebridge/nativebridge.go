// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativebridge

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

// NativeBridgeABI is the input ABI used to generate the binding from.
const NativeBridgeABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferBridged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"name\":\"setLastBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NativeBridgeBin is the compiled bytecode used for deploying new contracts.
var NativeBridgeBin = "0x608060405234801561001057600080fd5b5060006100216100c460201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100cc565b600033905090565b6108e8806100db6000396000f3fe60806040526004361061007b5760003560e01c80637f2c4ca81161004e5780637f2c4ca8146101735780638da5cb5b1461019e5780638f32d59b146101f5578063f2fde38b146102245761007b565b80633c64f04b146100805780635d974a66146100d3578063715018a61461010e5780637d32e7bd14610125575b600080fd5b34801561008c57600080fd5b506100b9600480360360208110156100a357600080fd5b8101908080359060200190929190505050610275565b604051808215151515815260200191505060405180910390f35b3480156100df57600080fd5b5061010c600480360360208110156100f657600080fd5b8101908080359060200190929190505050610295565b005b34801561011a57600080fd5b50610123610319565b005b6101716004803603604081101561013b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610452565b005b34801561017f57600080fd5b5061018861062a565b6040518082815260200191505060405180910390f35b3480156101aa57600080fd5b506101b3610634565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561020157600080fd5b5061020a61065d565b604051808215151515815260200191505060405180910390f35b34801561023057600080fd5b506102736004803603602081101561024757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106bb565b005b60026020528060005260406000206000915054906101000a900460ff1681565b61029d61065d565b61030f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8060018190555050565b61032161065d565b610393576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6002600082815260200190815260200160002060009054906101000a900460ff161561047d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156104b757600080fd5b60008273ffffffffffffffffffffffffffffffffffffffff163460405180600001905060006040518083038185875af1925050503d8060008114610517576040519150601f19603f3d011682016040523d82523d6000602084013e61051c565b606091505b5050905080610593576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f7472616e73666572206661696c6564000000000000000000000000000000000081525060200191505060405180910390fd5b818373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5817863dc4a9bed360e0cb00ea7b9d8584993922ecfcf649e6beaf091b670215346040518082815260200191505060405180910390a460016002600084815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b6000600154905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661069f610741565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b6106c361065d565b610735576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61073e81610749565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061088e6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a265627a7a723158201de139039078990c013c69d80a01a0c329853e2a28ceb5677cc9b33be8b2f7e264736f6c63430005100032"

// DeployNativeBridge deploys a new Ethereum contract, binding an instance of NativeBridge to it.
func DeployNativeBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NativeBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(NativeBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NativeBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeBridge{NativeBridgeCaller: NativeBridgeCaller{contract: contract}, NativeBridgeTransactor: NativeBridgeTransactor{contract: contract}, NativeBridgeFilterer: NativeBridgeFilterer{contract: contract}}, nil
}

// NativeBridge is an auto generated Go binding around an Ethereum contract.
type NativeBridge struct {
	NativeBridgeCaller     // Read-only binding to the contract
	NativeBridgeTransactor // Write-only binding to the contract
	NativeBridgeFilterer   // Log filterer for contract events
}

// NativeBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeBridgeSession struct {
	Contract     *NativeBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeBridgeCallerSession struct {
	Contract *NativeBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NativeBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeBridgeTransactorSession struct {
	Contract     *NativeBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativeBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeBridgeRaw struct {
	Contract *NativeBridge // Generic contract binding to access the raw methods on
}

// NativeBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeBridgeCallerRaw struct {
	Contract *NativeBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// NativeBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeBridgeTransactorRaw struct {
	Contract *NativeBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeBridge creates a new instance of NativeBridge, bound to a specific deployed contract.
func NewNativeBridge(address common.Address, backend bind.ContractBackend) (*NativeBridge, error) {
	contract, err := bindNativeBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeBridge{NativeBridgeCaller: NativeBridgeCaller{contract: contract}, NativeBridgeTransactor: NativeBridgeTransactor{contract: contract}, NativeBridgeFilterer: NativeBridgeFilterer{contract: contract}}, nil
}

// NewNativeBridgeCaller creates a new read-only instance of NativeBridge, bound to a specific deployed contract.
func NewNativeBridgeCaller(address common.Address, caller bind.ContractCaller) (*NativeBridgeCaller, error) {
	contract, err := bindNativeBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeBridgeCaller{contract: contract}, nil
}

// NewNativeBridgeTransactor creates a new write-only instance of NativeBridge, bound to a specific deployed contract.
func NewNativeBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeBridgeTransactor, error) {
	contract, err := bindNativeBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeBridgeTransactor{contract: contract}, nil
}

// NewNativeBridgeFilterer creates a new log filterer instance of NativeBridge, bound to a specific deployed contract.
func NewNativeBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeBridgeFilterer, error) {
	contract, err := bindNativeBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeBridgeFilterer{contract: contract}, nil
}

// bindNativeBridge binds a generic wrapper to an already deployed contract.
func bindNativeBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NativeBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeBridge *NativeBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NativeBridge.Contract.NativeBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeBridge *NativeBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeBridge.Contract.NativeBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeBridge *NativeBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeBridge.Contract.NativeBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeBridge *NativeBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NativeBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeBridge *NativeBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeBridge *NativeBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeBridge.Contract.contract.Transact(opts, method, params...)
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_NativeBridge *NativeBridgeCaller) GetLastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NativeBridge.contract.Call(opts, out, "getLastBlock")
	return *ret0, err
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_NativeBridge *NativeBridgeSession) GetLastBlock() (*big.Int, error) {
	return _NativeBridge.Contract.GetLastBlock(&_NativeBridge.CallOpts)
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_NativeBridge *NativeBridgeCallerSession) GetLastBlock() (*big.Int, error) {
	return _NativeBridge.Contract.GetLastBlock(&_NativeBridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_NativeBridge *NativeBridgeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NativeBridge.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_NativeBridge *NativeBridgeSession) IsOwner() (bool, error) {
	return _NativeBridge.Contract.IsOwner(&_NativeBridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_NativeBridge *NativeBridgeCallerSession) IsOwner() (bool, error) {
	return _NativeBridge.Contract.IsOwner(&_NativeBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NativeBridge *NativeBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NativeBridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NativeBridge *NativeBridgeSession) Owner() (common.Address, error) {
	return _NativeBridge.Contract.Owner(&_NativeBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NativeBridge *NativeBridgeCallerSession) Owner() (common.Address, error) {
	return _NativeBridge.Contract.Owner(&_NativeBridge.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bool)
func (_NativeBridge *NativeBridgeCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NativeBridge.contract.Call(opts, out, "transfers", arg0)
	return *ret0, err
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bool)
func (_NativeBridge *NativeBridgeSession) Transfers(arg0 [32]byte) (bool, error) {
	return _NativeBridge.Contract.Transfers(&_NativeBridge.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bool)
func (_NativeBridge *NativeBridgeCallerSession) Transfers(arg0 [32]byte) (bool, error) {
	return _NativeBridge.Contract.Transfers(&_NativeBridge.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeBridge *NativeBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeBridge *NativeBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeBridge.Contract.RenounceOwnership(&_NativeBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeBridge *NativeBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeBridge.Contract.RenounceOwnership(&_NativeBridge.TransactOpts)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 lastBlock) returns()
func (_NativeBridge *NativeBridgeTransactor) SetLastBlock(opts *bind.TransactOpts, lastBlock *big.Int) (*types.Transaction, error) {
	return _NativeBridge.contract.Transact(opts, "setLastBlock", lastBlock)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 lastBlock) returns()
func (_NativeBridge *NativeBridgeSession) SetLastBlock(lastBlock *big.Int) (*types.Transaction, error) {
	return _NativeBridge.Contract.SetLastBlock(&_NativeBridge.TransactOpts, lastBlock)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 lastBlock) returns()
func (_NativeBridge *NativeBridgeTransactorSession) SetLastBlock(lastBlock *big.Int) (*types.Transaction, error) {
	return _NativeBridge.Contract.SetLastBlock(&_NativeBridge.TransactOpts, lastBlock)
}

// Transfer is a paid mutator transaction binding the contract method 0x7d32e7bd.
//
// Solidity: function transfer(address to, bytes32 txHash) returns()
func (_NativeBridge *NativeBridgeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, txHash [32]byte) (*types.Transaction, error) {
	return _NativeBridge.contract.Transact(opts, "transfer", to, txHash)
}

// Transfer is a paid mutator transaction binding the contract method 0x7d32e7bd.
//
// Solidity: function transfer(address to, bytes32 txHash) returns()
func (_NativeBridge *NativeBridgeSession) Transfer(to common.Address, txHash [32]byte) (*types.Transaction, error) {
	return _NativeBridge.Contract.Transfer(&_NativeBridge.TransactOpts, to, txHash)
}

// Transfer is a paid mutator transaction binding the contract method 0x7d32e7bd.
//
// Solidity: function transfer(address to, bytes32 txHash) returns()
func (_NativeBridge *NativeBridgeTransactorSession) Transfer(to common.Address, txHash [32]byte) (*types.Transaction, error) {
	return _NativeBridge.Contract.Transfer(&_NativeBridge.TransactOpts, to, txHash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeBridge *NativeBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeBridge *NativeBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeBridge.Contract.TransferOwnership(&_NativeBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeBridge *NativeBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeBridge.Contract.TransferOwnership(&_NativeBridge.TransactOpts, newOwner)
}

// NativeBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeBridge contract.
type NativeBridgeOwnershipTransferredIterator struct {
	Event *NativeBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeBridgeOwnershipTransferred)
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
		it.Event = new(NativeBridgeOwnershipTransferred)
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
func (it *NativeBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the NativeBridge contract.
type NativeBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeBridge *NativeBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeBridgeOwnershipTransferredIterator{contract: _NativeBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeBridge *NativeBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeBridgeOwnershipTransferred)
				if err := _NativeBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeBridge *NativeBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*NativeBridgeOwnershipTransferred, error) {
	event := new(NativeBridgeOwnershipTransferred)
	if err := _NativeBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NativeBridgeTransferBridgedIterator is returned from FilterTransferBridged and is used to iterate over the raw logs and unpacked data for TransferBridged events raised by the NativeBridge contract.
type NativeBridgeTransferBridgedIterator struct {
	Event *NativeBridgeTransferBridged // Event containing the contract specifics and raw log

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
func (it *NativeBridgeTransferBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeBridgeTransferBridged)
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
		it.Event = new(NativeBridgeTransferBridged)
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
func (it *NativeBridgeTransferBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeBridgeTransferBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeBridgeTransferBridged represents a TransferBridged event raised by the NativeBridge contract.
type NativeBridgeTransferBridged struct {
	From   common.Address
	To     common.Address
	TxHash [32]byte
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferBridged is a free log retrieval operation binding the contract event 0x5817863dc4a9bed360e0cb00ea7b9d8584993922ecfcf649e6beaf091b670215.
//
// Solidity: event TransferBridged(address indexed from, address indexed to, bytes32 indexed txHash, uint256 value)
func (_NativeBridge *NativeBridgeFilterer) FilterTransferBridged(opts *bind.FilterOpts, from []common.Address, to []common.Address, txHash [][32]byte) (*NativeBridgeTransferBridgedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _NativeBridge.contract.FilterLogs(opts, "TransferBridged", fromRule, toRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &NativeBridgeTransferBridgedIterator{contract: _NativeBridge.contract, event: "TransferBridged", logs: logs, sub: sub}, nil
}

// WatchTransferBridged is a free log subscription operation binding the contract event 0x5817863dc4a9bed360e0cb00ea7b9d8584993922ecfcf649e6beaf091b670215.
//
// Solidity: event TransferBridged(address indexed from, address indexed to, bytes32 indexed txHash, uint256 value)
func (_NativeBridge *NativeBridgeFilterer) WatchTransferBridged(opts *bind.WatchOpts, sink chan<- *NativeBridgeTransferBridged, from []common.Address, to []common.Address, txHash [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _NativeBridge.contract.WatchLogs(opts, "TransferBridged", fromRule, toRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeBridgeTransferBridged)
				if err := _NativeBridge.contract.UnpackLog(event, "TransferBridged", log); err != nil {
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

// ParseTransferBridged is a log parse operation binding the contract event 0x5817863dc4a9bed360e0cb00ea7b9d8584993922ecfcf649e6beaf091b670215.
//
// Solidity: event TransferBridged(address indexed from, address indexed to, bytes32 indexed txHash, uint256 value)
func (_NativeBridge *NativeBridgeFilterer) ParseTransferBridged(log types.Log) (*NativeBridgeTransferBridged, error) {
	event := new(NativeBridgeTransferBridged)
	if err := _NativeBridge.contract.UnpackLog(event, "TransferBridged", log); err != nil {
		return nil, err
	}
	return event, nil
}
