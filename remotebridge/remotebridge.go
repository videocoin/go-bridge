// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package remotebridge

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

// RemoteBridgeABI is the input ABI used to generate the binding from.
const RemoteBridgeABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"TransferRegistered\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"local\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remote\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"name\":\"setLastBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"local\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remote\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RemoteBridgeBin is the compiled bytecode used for deploying new contracts.
var RemoteBridgeBin = "0x608060405234801561001057600080fd5b5060006100216100c460201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100cc565b600033905090565b610e0d806100db6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80637f2c4ca8116100665780637f2c4ca81461010d5780638da5cb5b1461012b5780638f32d59b14610149578063efcb64cb14610167578063f2fde38b1461018357610093565b806313f57c3e146100985780633c64f04b146100b45780635d974a66146100e7578063715018a614610103575b600080fd5b6100b260048036036100ad919081019061098d565b61019f565b005b6100ce60048036036100c99190810190610964565b61030b565b6040516100de9493929190610bfc565b60405180910390f35b61010160048036036100fc9190810190610a2c565b61037c565b005b61010b6103cd565b005b6101156104d3565b6040516101229190610cc1565b60405180910390f35b6101336104dd565b6040516101409190610bc6565b60405180910390f35b610151610506565b60405161015e9190610be1565b60405180910390f35b610181600480360361017c91908101906109c9565b610564565b005b61019d6004803603610198919081019061093b565b61075e565b005b6101a7610506565b6101e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101dd90610ca1565b60405180910390fd5b60026000838152602001908152602001600020600101601c9054906101000a900460ff16610249576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024090610c81565b60405180910390fd5b80600260008481526020019081526020016000206000018190555060006002600084815260200190815260200160002090508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16827ff323e3760341eb70d636a6304f7927d4ba8c42e4e8777d908098205ad3c8f49c8360010160149054906101000a900467ffffffffffffffff166040516102fe9190610cdc565b60405180910390a3505050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900467ffffffffffffffff169080600101601c9054906101000a900460ff16905084565b610384610506565b6103c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ba90610ca1565b60405180910390fd5b8060018190555050565b6103d5610506565b610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b90610ca1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600154905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166105486107b1565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b61056c610506565b6105ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a290610ca1565b60405180910390fd5b60026000858152602001908152602001600020600101601c9054906101000a900460ff161561060f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060690610c61565b60405180910390fd5b60405180608001604052808481526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018267ffffffffffffffff16815260200160011515815250600260008681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600101601c6101000a81548160ff0219169083151502179055509050508173ffffffffffffffffffffffffffffffffffffffff16837ff323e3760341eb70d636a6304f7927d4ba8c42e4e8777d908098205ad3c8f49c836040516107509190610cdc565b60405180910390a350505050565b610766610506565b6107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079c90610ca1565b60405180910390fd5b6107ae816107b9565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610829576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082090610c41565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000813590506108f681610d6e565b92915050565b60008135905061090b81610d85565b92915050565b60008135905061092081610d9c565b92915050565b60008135905061093581610db3565b92915050565b60006020828403121561094d57600080fd5b600061095b848285016108e7565b91505092915050565b60006020828403121561097657600080fd5b6000610984848285016108fc565b91505092915050565b600080604083850312156109a057600080fd5b60006109ae858286016108fc565b92505060206109bf858286016108fc565b9150509250929050565b600080600080608085870312156109df57600080fd5b60006109ed878288016108fc565b94505060206109fe878288016108fc565b9350506040610a0f878288016108e7565b9250506060610a2087828801610926565b91505092959194509250565b600060208284031215610a3e57600080fd5b6000610a4c84828501610911565b91505092915050565b610a5e81610d08565b82525050565b610a6d81610d1a565b82525050565b610a7c81610d26565b82525050565b6000610a8f602683610cf7565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610af5601b83610cf7565b91507f7472616e7366657220616c7265616479207265676973746572656400000000006000830152602082019050919050565b6000610b35601783610cf7565b91507f7472616e73666572206e6f7420726567697374657265640000000000000000006000830152602082019050919050565b6000610b75602083610cf7565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b610bb181610d50565b82525050565b610bc081610d5a565b82525050565b6000602082019050610bdb6000830184610a55565b92915050565b6000602082019050610bf66000830184610a64565b92915050565b6000608082019050610c116000830187610a73565b610c1e6020830186610a55565b610c2b6040830185610bb7565b610c386060830184610a64565b95945050505050565b60006020820190508181036000830152610c5a81610a82565b9050919050565b60006020820190508181036000830152610c7a81610ae8565b9050919050565b60006020820190508181036000830152610c9a81610b28565b9050919050565b60006020820190508181036000830152610cba81610b68565b9050919050565b6000602082019050610cd66000830184610ba8565b92915050565b6000602082019050610cf16000830184610bb7565b92915050565b600082825260208201905092915050565b6000610d1382610d30565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b610d7781610d08565b8114610d8257600080fd5b50565b610d8e81610d26565b8114610d9957600080fd5b50565b610da581610d50565b8114610db057600080fd5b50565b610dbc81610d5a565b8114610dc757600080fd5b5056fea365627a7a7231582022d6c80740eaf0d89503b502375d9f5e29a3ec8b84cb3d21f2836aa42535076f6c6578706572696d656e74616cf564736f6c63430005100040"

// DeployRemoteBridge deploys a new Ethereum contract, binding an instance of RemoteBridge to it.
func DeployRemoteBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RemoteBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(RemoteBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RemoteBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RemoteBridge{RemoteBridgeCaller: RemoteBridgeCaller{contract: contract}, RemoteBridgeTransactor: RemoteBridgeTransactor{contract: contract}, RemoteBridgeFilterer: RemoteBridgeFilterer{contract: contract}}, nil
}

// RemoteBridge is an auto generated Go binding around an Ethereum contract.
type RemoteBridge struct {
	RemoteBridgeCaller     // Read-only binding to the contract
	RemoteBridgeTransactor // Write-only binding to the contract
	RemoteBridgeFilterer   // Log filterer for contract events
}

// RemoteBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RemoteBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RemoteBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RemoteBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RemoteBridgeSession struct {
	Contract     *RemoteBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RemoteBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RemoteBridgeCallerSession struct {
	Contract *RemoteBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RemoteBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RemoteBridgeTransactorSession struct {
	Contract     *RemoteBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RemoteBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RemoteBridgeRaw struct {
	Contract *RemoteBridge // Generic contract binding to access the raw methods on
}

// RemoteBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RemoteBridgeCallerRaw struct {
	Contract *RemoteBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// RemoteBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RemoteBridgeTransactorRaw struct {
	Contract *RemoteBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRemoteBridge creates a new instance of RemoteBridge, bound to a specific deployed contract.
func NewRemoteBridge(address common.Address, backend bind.ContractBackend) (*RemoteBridge, error) {
	contract, err := bindRemoteBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RemoteBridge{RemoteBridgeCaller: RemoteBridgeCaller{contract: contract}, RemoteBridgeTransactor: RemoteBridgeTransactor{contract: contract}, RemoteBridgeFilterer: RemoteBridgeFilterer{contract: contract}}, nil
}

// NewRemoteBridgeCaller creates a new read-only instance of RemoteBridge, bound to a specific deployed contract.
func NewRemoteBridgeCaller(address common.Address, caller bind.ContractCaller) (*RemoteBridgeCaller, error) {
	contract, err := bindRemoteBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeCaller{contract: contract}, nil
}

// NewRemoteBridgeTransactor creates a new write-only instance of RemoteBridge, bound to a specific deployed contract.
func NewRemoteBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*RemoteBridgeTransactor, error) {
	contract, err := bindRemoteBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeTransactor{contract: contract}, nil
}

// NewRemoteBridgeFilterer creates a new log filterer instance of RemoteBridge, bound to a specific deployed contract.
func NewRemoteBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*RemoteBridgeFilterer, error) {
	contract, err := bindRemoteBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeFilterer{contract: contract}, nil
}

// bindRemoteBridge binds a generic wrapper to an already deployed contract.
func bindRemoteBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RemoteBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RemoteBridge *RemoteBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RemoteBridge.Contract.RemoteBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RemoteBridge *RemoteBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RemoteBridge.Contract.RemoteBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RemoteBridge *RemoteBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RemoteBridge.Contract.RemoteBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RemoteBridge *RemoteBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RemoteBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RemoteBridge *RemoteBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RemoteBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RemoteBridge *RemoteBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RemoteBridge.Contract.contract.Transact(opts, method, params...)
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_RemoteBridge *RemoteBridgeCaller) GetLastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RemoteBridge.contract.Call(opts, out, "getLastBlock")
	return *ret0, err
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_RemoteBridge *RemoteBridgeSession) GetLastBlock() (*big.Int, error) {
	return _RemoteBridge.Contract.GetLastBlock(&_RemoteBridge.CallOpts)
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_RemoteBridge *RemoteBridgeCallerSession) GetLastBlock() (*big.Int, error) {
	return _RemoteBridge.Contract.GetLastBlock(&_RemoteBridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RemoteBridge *RemoteBridgeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RemoteBridge.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RemoteBridge *RemoteBridgeSession) IsOwner() (bool, error) {
	return _RemoteBridge.Contract.IsOwner(&_RemoteBridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RemoteBridge *RemoteBridgeCallerSession) IsOwner() (bool, error) {
	return _RemoteBridge.Contract.IsOwner(&_RemoteBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RemoteBridge *RemoteBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RemoteBridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RemoteBridge *RemoteBridgeSession) Owner() (common.Address, error) {
	return _RemoteBridge.Contract.Owner(&_RemoteBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RemoteBridge *RemoteBridgeCallerSession) Owner() (common.Address, error) {
	return _RemoteBridge.Contract.Owner(&_RemoteBridge.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 hash, address signer, uint64 nonce, bool exist)
func (_RemoteBridge *RemoteBridgeCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Hash   [32]byte
	Signer common.Address
	Nonce  uint64
	Exist  bool
}, error) {
	ret := new(struct {
		Hash   [32]byte
		Signer common.Address
		Nonce  uint64
		Exist  bool
	})
	out := ret
	err := _RemoteBridge.contract.Call(opts, out, "transfers", arg0)
	return *ret, err
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 hash, address signer, uint64 nonce, bool exist)
func (_RemoteBridge *RemoteBridgeSession) Transfers(arg0 [32]byte) (struct {
	Hash   [32]byte
	Signer common.Address
	Nonce  uint64
	Exist  bool
}, error) {
	return _RemoteBridge.Contract.Transfers(&_RemoteBridge.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 hash, address signer, uint64 nonce, bool exist)
func (_RemoteBridge *RemoteBridgeCallerSession) Transfers(arg0 [32]byte) (struct {
	Hash   [32]byte
	Signer common.Address
	Nonce  uint64
	Exist  bool
}, error) {
	return _RemoteBridge.Contract.Transfers(&_RemoteBridge.CallOpts, arg0)
}

// Register is a paid mutator transaction binding the contract method 0xefcb64cb.
//
// Solidity: function register(bytes32 local, bytes32 remote, address signer, uint64 nonce) returns()
func (_RemoteBridge *RemoteBridgeTransactor) Register(opts *bind.TransactOpts, local [32]byte, remote [32]byte, signer common.Address, nonce uint64) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "register", local, remote, signer, nonce)
}

// Register is a paid mutator transaction binding the contract method 0xefcb64cb.
//
// Solidity: function register(bytes32 local, bytes32 remote, address signer, uint64 nonce) returns()
func (_RemoteBridge *RemoteBridgeSession) Register(local [32]byte, remote [32]byte, signer common.Address, nonce uint64) (*types.Transaction, error) {
	return _RemoteBridge.Contract.Register(&_RemoteBridge.TransactOpts, local, remote, signer, nonce)
}

// Register is a paid mutator transaction binding the contract method 0xefcb64cb.
//
// Solidity: function register(bytes32 local, bytes32 remote, address signer, uint64 nonce) returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) Register(local [32]byte, remote [32]byte, signer common.Address, nonce uint64) (*types.Transaction, error) {
	return _RemoteBridge.Contract.Register(&_RemoteBridge.TransactOpts, local, remote, signer, nonce)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RemoteBridge *RemoteBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RemoteBridge *RemoteBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _RemoteBridge.Contract.RenounceOwnership(&_RemoteBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RemoteBridge.Contract.RenounceOwnership(&_RemoteBridge.TransactOpts)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 lastBlock) returns()
func (_RemoteBridge *RemoteBridgeTransactor) SetLastBlock(opts *bind.TransactOpts, lastBlock *big.Int) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "setLastBlock", lastBlock)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 lastBlock) returns()
func (_RemoteBridge *RemoteBridgeSession) SetLastBlock(lastBlock *big.Int) (*types.Transaction, error) {
	return _RemoteBridge.Contract.SetLastBlock(&_RemoteBridge.TransactOpts, lastBlock)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 lastBlock) returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) SetLastBlock(lastBlock *big.Int) (*types.Transaction, error) {
	return _RemoteBridge.Contract.SetLastBlock(&_RemoteBridge.TransactOpts, lastBlock)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RemoteBridge *RemoteBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RemoteBridge *RemoteBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RemoteBridge.Contract.TransferOwnership(&_RemoteBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RemoteBridge.Contract.TransferOwnership(&_RemoteBridge.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 local, bytes32 remote) returns()
func (_RemoteBridge *RemoteBridgeTransactor) Update(opts *bind.TransactOpts, local [32]byte, remote [32]byte) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "update", local, remote)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 local, bytes32 remote) returns()
func (_RemoteBridge *RemoteBridgeSession) Update(local [32]byte, remote [32]byte) (*types.Transaction, error) {
	return _RemoteBridge.Contract.Update(&_RemoteBridge.TransactOpts, local, remote)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 local, bytes32 remote) returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) Update(local [32]byte, remote [32]byte) (*types.Transaction, error) {
	return _RemoteBridge.Contract.Update(&_RemoteBridge.TransactOpts, local, remote)
}

// RemoteBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RemoteBridge contract.
type RemoteBridgeOwnershipTransferredIterator struct {
	Event *RemoteBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RemoteBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RemoteBridgeOwnershipTransferred)
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
		it.Event = new(RemoteBridgeOwnershipTransferred)
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
func (it *RemoteBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RemoteBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RemoteBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the RemoteBridge contract.
type RemoteBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RemoteBridge *RemoteBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RemoteBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RemoteBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeOwnershipTransferredIterator{contract: _RemoteBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RemoteBridge *RemoteBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RemoteBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RemoteBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RemoteBridgeOwnershipTransferred)
				if err := _RemoteBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RemoteBridge *RemoteBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*RemoteBridgeOwnershipTransferred, error) {
	event := new(RemoteBridgeOwnershipTransferred)
	if err := _RemoteBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RemoteBridgeTransferRegisteredIterator is returned from FilterTransferRegistered and is used to iterate over the raw logs and unpacked data for TransferRegistered events raised by the RemoteBridge contract.
type RemoteBridgeTransferRegisteredIterator struct {
	Event *RemoteBridgeTransferRegistered // Event containing the contract specifics and raw log

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
func (it *RemoteBridgeTransferRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RemoteBridgeTransferRegistered)
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
		it.Event = new(RemoteBridgeTransferRegistered)
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
func (it *RemoteBridgeTransferRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RemoteBridgeTransferRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RemoteBridgeTransferRegistered represents a TransferRegistered event raised by the RemoteBridge contract.
type RemoteBridgeTransferRegistered struct {
	Hash   [32]byte
	Signer common.Address
	Nonce  uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferRegistered is a free log retrieval operation binding the contract event 0xf323e3760341eb70d636a6304f7927d4ba8c42e4e8777d908098205ad3c8f49c.
//
// Solidity: event TransferRegistered(bytes32 indexed hash, address indexed signer, uint64 nonce)
func (_RemoteBridge *RemoteBridgeFilterer) FilterTransferRegistered(opts *bind.FilterOpts, hash [][32]byte, signer []common.Address) (*RemoteBridgeTransferRegisteredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _RemoteBridge.contract.FilterLogs(opts, "TransferRegistered", hashRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeTransferRegisteredIterator{contract: _RemoteBridge.contract, event: "TransferRegistered", logs: logs, sub: sub}, nil
}

// WatchTransferRegistered is a free log subscription operation binding the contract event 0xf323e3760341eb70d636a6304f7927d4ba8c42e4e8777d908098205ad3c8f49c.
//
// Solidity: event TransferRegistered(bytes32 indexed hash, address indexed signer, uint64 nonce)
func (_RemoteBridge *RemoteBridgeFilterer) WatchTransferRegistered(opts *bind.WatchOpts, sink chan<- *RemoteBridgeTransferRegistered, hash [][32]byte, signer []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _RemoteBridge.contract.WatchLogs(opts, "TransferRegistered", hashRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RemoteBridgeTransferRegistered)
				if err := _RemoteBridge.contract.UnpackLog(event, "TransferRegistered", log); err != nil {
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

// ParseTransferRegistered is a log parse operation binding the contract event 0xf323e3760341eb70d636a6304f7927d4ba8c42e4e8777d908098205ad3c8f49c.
//
// Solidity: event TransferRegistered(bytes32 indexed hash, address indexed signer, uint64 nonce)
func (_RemoteBridge *RemoteBridgeFilterer) ParseTransferRegistered(log types.Log) (*RemoteBridgeTransferRegistered, error) {
	event := new(RemoteBridgeTransferRegistered)
	if err := _RemoteBridge.contract.UnpackLog(event, "TransferRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
