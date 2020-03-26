// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package homebridge

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

// HomeBridgeABI is the input ABI used to generate the binding from.
const HomeBridgeABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_txLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumTransfer\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferExceededLimits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"TransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"TransferRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferWithinLimits\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"checkIfTransferValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"executeTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedTransfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDailyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinimalTXLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTXLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"setDailyLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastBlock\",\"type\":\"uint256\"}],\"name\":\"setLastBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"setLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txMinLimit\",\"type\":\"uint256\"}],\"name\":\"setMinimalTXLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txLimit\",\"type\":\"uint256\"}],\"name\":\"setTXLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_day\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"totalExecutedPerDayPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"totalExecutedPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_day\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"totalTXExecutedPerDayPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"txPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HomeBridgeBin is the compiled bytecode used for deploying new contracts.
var HomeBridgeBin = "0x608060405234801561001057600080fd5b506040516114373803806114378339818101604052608081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a383600181905550826002819055508160038190555080600581905550505050506112ea8061014d6000396000f3fe6080604052600436106101665760003560e01c80638f32d59b116100d1578063b4634bbe1161008a578063d1320f7b11610064578063d1320f7b14610740578063e12a799f1461077b578063ed43f594146107e0578063f2fde38b1461082f57610166565b8063b4634bbe1461065d578063b4e2216e14610698578063b6eb1cad1461071557610166565b80638f32d59b146104d75780639a9f76da14610506578063a6895b5a14610559578063aa23ed4d14610584578063b20d30a9146105f7578063b295a00e1461063257610166565b80635d974a66116101235780635d974a661461039d57806367deef47146103d8578063715018a6146104275780637f2c4ca81461043e578063853828b6146104695780638da5cb5b1461048057610166565b80631a67e3e9146101c957806327ea6f2b146101f457806331a1bbfd1461022f5780633e6968b61461029e57806345949a99146102c95780635ccfb7de1461032e575b6000341161017357600080fd5b6000803690501461018357600080fd5b343373ffffffffffffffffffffffffffffffffffffffff167fa1dd607db049dad651d4ebbc3b099a1cc910f208c0e692ee548615911da5e26c60405160405180910390a3005b3480156101d557600080fd5b506101de610880565b6040518082815260200191505060405180910390f35b34801561020057600080fd5b5061022d6004803603602081101561021757600080fd5b810190808035906020019092919050505061088a565b005b34801561023b57600080fd5b506102886004803603604081101561025257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108a5565b6040518082815260200191505060405180910390f35b3480156102aa57600080fd5b506102b3610920565b6040518082815260200191505060405180910390f35b3480156102d557600080fd5b50610318600480360360208110156102ec57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610934565b6040518082815260200191505060405180910390f35b34801561033a57600080fd5b506103876004803603604081101561035157600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061094c565b6040518082815260200191505060405180910390f35b3480156103a957600080fd5b506103d6600480360360208110156103c057600080fd5b81019080803590602001909291905050506109c7565b005b3480156103e457600080fd5b50610411600480360360208110156103fb57600080fd5b81019080803590602001909291905050506109e2565b6040518082815260200191505060405180910390f35b34801561043357600080fd5b5061043c6109fa565b005b34801561044a57600080fd5b50610453610aca565b6040518082815260200191505060405180910390f35b34801561047557600080fd5b5061047e610ad4565b005b34801561048c57600080fd5b50610495610b34565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104e357600080fd5b506104ec610b5d565b604051808215151515815260200191505060405180910390f35b34801561051257600080fd5b5061053f6004803603602081101561052957600080fd5b8101908080359060200190929190505050610bb4565b604051808215151515815260200191505060405180910390f35b34801561056557600080fd5b5061056e610bd4565b6040518082815260200191505060405180910390f35b34801561059057600080fd5b506105dd600480360360408110156105a757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bde565b604051808215151515815260200191505060405180910390f35b34801561060357600080fd5b506106306004803603602081101561061a57600080fd5b8101908080359060200190929190505050610ce3565b005b34801561063e57600080fd5b50610647610cfe565b6040518082815260200191505060405180910390f35b34801561066957600080fd5b506106966004803603602081101561068057600080fd5b8101908080359060200190929190505050610d08565b005b3480156106a457600080fd5b506106fb600480360360608110156106bb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610d23565b604051808215151515815260200191505060405180910390f35b34801561072157600080fd5b5061072a610f3d565b6040518082815260200191505060405180910390f35b34801561074c57600080fd5b506107796004803603602081101561076357600080fd5b8101908080359060200190929190505050610f47565b005b34801561078757600080fd5b506107ca6004803603602081101561079e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f62565b6040518082815260200191505060405180910390f35b3480156107ec57600080fd5b506108196004803603602081101561080357600080fd5b8101908080359060200190929190505050610fab565b6040518082815260200191505060405180910390f35b34801561083b57600080fd5b5061087e6004803603602081101561085257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fc3565b005b6000600354905090565b610892610b5d565b61089b57600080fd5b8060018190555050565b6000600760008385604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018281526020019250505060405160208183030381529060405280519060200120815260200190815260200160002054905092915050565b600062015180428161092e57fe5b04905090565b60066020528060005260406000206000915090505481565b6000600960008385604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018281526020019250505060405160208183030381529060405280519060200120815260200190815260200160002054905092915050565b6109cf610b5d565b6109d857600080fd5b8060048190555050565b60096020528060005260406000206000915090505481565b610a02610b5d565b610a0b57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600454905090565b610adc610b5d565b610ae557600080fd5b60004790503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610b30573d6000803e3d6000fd5b5050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60086020528060005260406000206000915054906101000a900460ff1681565b6000600254905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c1957600080fd5b60008211610c2657600080fd5b6000610c4383610c3586610f62565b610fe090919063ffffffff16565b90506000610c6a84610c5c610c56610920565b886108a5565b610fe090919063ffffffff16565b90506000610c926001610c84610c7e610920565b8961094c565b610fe090919063ffffffff16565b90508260015410158015610ca857508160025410155b8015610cb657508060035410155b8015610cc457506005548510155b15610cd55760019350505050610cdd565b600093505050505b92915050565b610ceb610b5d565b610cf457600080fd5b8060028190555050565b6000600154905090565b610d10610b5d565b610d1957600080fd5b8060058190555050565b6000610d2d610b5d565b610d3657600080fd5b47831115610d4357600080fd5b6008600083815260200190815260200160002060009054906101000a900460ff1615610d6e57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610da857600080fd5b60008090506000859050610dbc8186610bde565b15610ebf57610de586610de087610dd28a610f62565b610fe090919063ffffffff16565b610fff565b610e19610df0610920565b87610e1488610e06610e00610920565b8c6108a5565b610fe090919063ffffffff16565b611047565b610e2a610e24610920565b876110c1565b8073ffffffffffffffffffffffffffffffffffffffff166108fc869081150290604051600060405180830381858888f19350505050158015610e70573d6000803e3d6000fd5b5083858773ffffffffffffffffffffffffffffffffffffffff167f6eb822c246c0021523c4209516767730f4b9260f24858135daa61b0377568fd260405160405180910390a460019150610f05565b83858773ffffffffffffffffffffffffffffffffffffffff167f8c17b1c1874fa559466f835c95f0fbce9b80e6c948a17b4dbcb8785de90fc14b60405160405180910390a45b60016008600086815260200190815260200160002060006101000a81548160ff02191690831515021790555081925050509392505050565b6000600554905090565b610f4f610b5d565b610f5857600080fd5b8060038190555050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60076020528060005260406000206000915090505481565b610fcb610b5d565b610fd457600080fd5b610fdd816111bd565b50565b600080828401905083811015610ff557600080fd5b8091505092915050565b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b80600760008486604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018281526020019250505060405160208183030381529060405280519060200120815260200190815260200160002081905550505050565b6111456001600960008486604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018281526020019250505060405160208183030381529060405280519060200120815260200190815260200160002054610fe090919063ffffffff16565b600960008385604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140182815260200192505050604051602081830303815290604052805190602001208152602001908152602001600020819055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156111f757600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fea265627a7a7231582014ee25d5ac20ca96ae57596c74a70e9a009542ce2281cf54bb59ebe32d8b240b64736f6c63430005100032"

// DeployHomeBridge deploys a new Ethereum contract, binding an instance of HomeBridge to it.
func DeployHomeBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _maxLimit *big.Int, _dailyLimit *big.Int, _txLimit *big.Int, _minimumTransfer *big.Int) (common.Address, *types.Transaction, *HomeBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(HomeBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HomeBridgeBin), backend, _maxLimit, _dailyLimit, _txLimit, _minimumTransfer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HomeBridge{HomeBridgeCaller: HomeBridgeCaller{contract: contract}, HomeBridgeTransactor: HomeBridgeTransactor{contract: contract}, HomeBridgeFilterer: HomeBridgeFilterer{contract: contract}}, nil
}

// HomeBridge is an auto generated Go binding around an Ethereum contract.
type HomeBridge struct {
	HomeBridgeCaller     // Read-only binding to the contract
	HomeBridgeTransactor // Write-only binding to the contract
	HomeBridgeFilterer   // Log filterer for contract events
}

// HomeBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HomeBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HomeBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HomeBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HomeBridgeSession struct {
	Contract     *HomeBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HomeBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HomeBridgeCallerSession struct {
	Contract *HomeBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HomeBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HomeBridgeTransactorSession struct {
	Contract     *HomeBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HomeBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HomeBridgeRaw struct {
	Contract *HomeBridge // Generic contract binding to access the raw methods on
}

// HomeBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HomeBridgeCallerRaw struct {
	Contract *HomeBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// HomeBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HomeBridgeTransactorRaw struct {
	Contract *HomeBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHomeBridge creates a new instance of HomeBridge, bound to a specific deployed contract.
func NewHomeBridge(address common.Address, backend bind.ContractBackend) (*HomeBridge, error) {
	contract, err := bindHomeBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HomeBridge{HomeBridgeCaller: HomeBridgeCaller{contract: contract}, HomeBridgeTransactor: HomeBridgeTransactor{contract: contract}, HomeBridgeFilterer: HomeBridgeFilterer{contract: contract}}, nil
}

// NewHomeBridgeCaller creates a new read-only instance of HomeBridge, bound to a specific deployed contract.
func NewHomeBridgeCaller(address common.Address, caller bind.ContractCaller) (*HomeBridgeCaller, error) {
	contract, err := bindHomeBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeCaller{contract: contract}, nil
}

// NewHomeBridgeTransactor creates a new write-only instance of HomeBridge, bound to a specific deployed contract.
func NewHomeBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*HomeBridgeTransactor, error) {
	contract, err := bindHomeBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeTransactor{contract: contract}, nil
}

// NewHomeBridgeFilterer creates a new log filterer instance of HomeBridge, bound to a specific deployed contract.
func NewHomeBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*HomeBridgeFilterer, error) {
	contract, err := bindHomeBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeFilterer{contract: contract}, nil
}

// bindHomeBridge binds a generic wrapper to an already deployed contract.
func bindHomeBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HomeBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HomeBridge *HomeBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HomeBridge.Contract.HomeBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HomeBridge *HomeBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HomeBridge.Contract.HomeBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HomeBridge *HomeBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HomeBridge.Contract.HomeBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HomeBridge *HomeBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HomeBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HomeBridge *HomeBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HomeBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HomeBridge *HomeBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HomeBridge.Contract.contract.Transact(opts, method, params...)
}

// CheckIfTransferValid is a free data retrieval call binding the contract method 0xaa23ed4d.
//
// Solidity: function checkIfTransferValid(address _recipient, uint256 _value) constant returns(bool)
func (_HomeBridge *HomeBridgeCaller) CheckIfTransferValid(opts *bind.CallOpts, _recipient common.Address, _value *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "checkIfTransferValid", _recipient, _value)
	return *ret0, err
}

// CheckIfTransferValid is a free data retrieval call binding the contract method 0xaa23ed4d.
//
// Solidity: function checkIfTransferValid(address _recipient, uint256 _value) constant returns(bool)
func (_HomeBridge *HomeBridgeSession) CheckIfTransferValid(_recipient common.Address, _value *big.Int) (bool, error) {
	return _HomeBridge.Contract.CheckIfTransferValid(&_HomeBridge.CallOpts, _recipient, _value)
}

// CheckIfTransferValid is a free data retrieval call binding the contract method 0xaa23ed4d.
//
// Solidity: function checkIfTransferValid(address _recipient, uint256 _value) constant returns(bool)
func (_HomeBridge *HomeBridgeCallerSession) CheckIfTransferValid(_recipient common.Address, _value *big.Int) (bool, error) {
	return _HomeBridge.Contract.CheckIfTransferValid(&_HomeBridge.CallOpts, _recipient, _value)
}

// Executed is a free data retrieval call binding the contract method 0x45949a99.
//
// Solidity: function executed(address ) constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) Executed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "executed", arg0)
	return *ret0, err
}

// Executed is a free data retrieval call binding the contract method 0x45949a99.
//
// Solidity: function executed(address ) constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) Executed(arg0 common.Address) (*big.Int, error) {
	return _HomeBridge.Contract.Executed(&_HomeBridge.CallOpts, arg0)
}

// Executed is a free data retrieval call binding the contract method 0x45949a99.
//
// Solidity: function executed(address ) constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) Executed(arg0 common.Address) (*big.Int, error) {
	return _HomeBridge.Contract.Executed(&_HomeBridge.CallOpts, arg0)
}

// ExecutedPerDay is a free data retrieval call binding the contract method 0xed43f594.
//
// Solidity: function executedPerDay(bytes32 ) constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) ExecutedPerDay(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "executedPerDay", arg0)
	return *ret0, err
}

// ExecutedPerDay is a free data retrieval call binding the contract method 0xed43f594.
//
// Solidity: function executedPerDay(bytes32 ) constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) ExecutedPerDay(arg0 [32]byte) (*big.Int, error) {
	return _HomeBridge.Contract.ExecutedPerDay(&_HomeBridge.CallOpts, arg0)
}

// ExecutedPerDay is a free data retrieval call binding the contract method 0xed43f594.
//
// Solidity: function executedPerDay(bytes32 ) constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) ExecutedPerDay(arg0 [32]byte) (*big.Int, error) {
	return _HomeBridge.Contract.ExecutedPerDay(&_HomeBridge.CallOpts, arg0)
}

// ExecutedTransfers is a free data retrieval call binding the contract method 0x9a9f76da.
//
// Solidity: function executedTransfers(bytes32 ) constant returns(bool)
func (_HomeBridge *HomeBridgeCaller) ExecutedTransfers(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "executedTransfers", arg0)
	return *ret0, err
}

// ExecutedTransfers is a free data retrieval call binding the contract method 0x9a9f76da.
//
// Solidity: function executedTransfers(bytes32 ) constant returns(bool)
func (_HomeBridge *HomeBridgeSession) ExecutedTransfers(arg0 [32]byte) (bool, error) {
	return _HomeBridge.Contract.ExecutedTransfers(&_HomeBridge.CallOpts, arg0)
}

// ExecutedTransfers is a free data retrieval call binding the contract method 0x9a9f76da.
//
// Solidity: function executedTransfers(bytes32 ) constant returns(bool)
func (_HomeBridge *HomeBridgeCallerSession) ExecutedTransfers(arg0 [32]byte) (bool, error) {
	return _HomeBridge.Contract.ExecutedTransfers(&_HomeBridge.CallOpts, arg0)
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) GetCurrentDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "getCurrentDay")
	return *ret0, err
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) GetCurrentDay() (*big.Int, error) {
	return _HomeBridge.Contract.GetCurrentDay(&_HomeBridge.CallOpts)
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) GetCurrentDay() (*big.Int, error) {
	return _HomeBridge.Contract.GetCurrentDay(&_HomeBridge.CallOpts)
}

// GetDailyLimit is a free data retrieval call binding the contract method 0xa6895b5a.
//
// Solidity: function getDailyLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) GetDailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "getDailyLimit")
	return *ret0, err
}

// GetDailyLimit is a free data retrieval call binding the contract method 0xa6895b5a.
//
// Solidity: function getDailyLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) GetDailyLimit() (*big.Int, error) {
	return _HomeBridge.Contract.GetDailyLimit(&_HomeBridge.CallOpts)
}

// GetDailyLimit is a free data retrieval call binding the contract method 0xa6895b5a.
//
// Solidity: function getDailyLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) GetDailyLimit() (*big.Int, error) {
	return _HomeBridge.Contract.GetDailyLimit(&_HomeBridge.CallOpts)
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) GetLastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "getLastBlock")
	return *ret0, err
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) GetLastBlock() (*big.Int, error) {
	return _HomeBridge.Contract.GetLastBlock(&_HomeBridge.CallOpts)
}

// GetLastBlock is a free data retrieval call binding the contract method 0x7f2c4ca8.
//
// Solidity: function getLastBlock() constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) GetLastBlock() (*big.Int, error) {
	return _HomeBridge.Contract.GetLastBlock(&_HomeBridge.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) GetLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "getLimit")
	return *ret0, err
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) GetLimit() (*big.Int, error) {
	return _HomeBridge.Contract.GetLimit(&_HomeBridge.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) GetLimit() (*big.Int, error) {
	return _HomeBridge.Contract.GetLimit(&_HomeBridge.CallOpts)
}

// GetMinimalTXLimit is a free data retrieval call binding the contract method 0xb6eb1cad.
//
// Solidity: function getMinimalTXLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) GetMinimalTXLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "getMinimalTXLimit")
	return *ret0, err
}

// GetMinimalTXLimit is a free data retrieval call binding the contract method 0xb6eb1cad.
//
// Solidity: function getMinimalTXLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) GetMinimalTXLimit() (*big.Int, error) {
	return _HomeBridge.Contract.GetMinimalTXLimit(&_HomeBridge.CallOpts)
}

// GetMinimalTXLimit is a free data retrieval call binding the contract method 0xb6eb1cad.
//
// Solidity: function getMinimalTXLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) GetMinimalTXLimit() (*big.Int, error) {
	return _HomeBridge.Contract.GetMinimalTXLimit(&_HomeBridge.CallOpts)
}

// GetTXLimit is a free data retrieval call binding the contract method 0x1a67e3e9.
//
// Solidity: function getTXLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) GetTXLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "getTXLimit")
	return *ret0, err
}

// GetTXLimit is a free data retrieval call binding the contract method 0x1a67e3e9.
//
// Solidity: function getTXLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) GetTXLimit() (*big.Int, error) {
	return _HomeBridge.Contract.GetTXLimit(&_HomeBridge.CallOpts)
}

// GetTXLimit is a free data retrieval call binding the contract method 0x1a67e3e9.
//
// Solidity: function getTXLimit() constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) GetTXLimit() (*big.Int, error) {
	return _HomeBridge.Contract.GetTXLimit(&_HomeBridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_HomeBridge *HomeBridgeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_HomeBridge *HomeBridgeSession) IsOwner() (bool, error) {
	return _HomeBridge.Contract.IsOwner(&_HomeBridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_HomeBridge *HomeBridgeCallerSession) IsOwner() (bool, error) {
	return _HomeBridge.Contract.IsOwner(&_HomeBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HomeBridge *HomeBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HomeBridge *HomeBridgeSession) Owner() (common.Address, error) {
	return _HomeBridge.Contract.Owner(&_HomeBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HomeBridge *HomeBridgeCallerSession) Owner() (common.Address, error) {
	return _HomeBridge.Contract.Owner(&_HomeBridge.CallOpts)
}

// TotalExecutedPerDayPerUser is a free data retrieval call binding the contract method 0x31a1bbfd.
//
// Solidity: function totalExecutedPerDayPerUser(uint256 _day, address _addr) constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) TotalExecutedPerDayPerUser(opts *bind.CallOpts, _day *big.Int, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "totalExecutedPerDayPerUser", _day, _addr)
	return *ret0, err
}

// TotalExecutedPerDayPerUser is a free data retrieval call binding the contract method 0x31a1bbfd.
//
// Solidity: function totalExecutedPerDayPerUser(uint256 _day, address _addr) constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) TotalExecutedPerDayPerUser(_day *big.Int, _addr common.Address) (*big.Int, error) {
	return _HomeBridge.Contract.TotalExecutedPerDayPerUser(&_HomeBridge.CallOpts, _day, _addr)
}

// TotalExecutedPerDayPerUser is a free data retrieval call binding the contract method 0x31a1bbfd.
//
// Solidity: function totalExecutedPerDayPerUser(uint256 _day, address _addr) constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) TotalExecutedPerDayPerUser(_day *big.Int, _addr common.Address) (*big.Int, error) {
	return _HomeBridge.Contract.TotalExecutedPerDayPerUser(&_HomeBridge.CallOpts, _day, _addr)
}

// TotalExecutedPerUser is a free data retrieval call binding the contract method 0xe12a799f.
//
// Solidity: function totalExecutedPerUser(address _recipient) constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) TotalExecutedPerUser(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "totalExecutedPerUser", _recipient)
	return *ret0, err
}

// TotalExecutedPerUser is a free data retrieval call binding the contract method 0xe12a799f.
//
// Solidity: function totalExecutedPerUser(address _recipient) constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) TotalExecutedPerUser(_recipient common.Address) (*big.Int, error) {
	return _HomeBridge.Contract.TotalExecutedPerUser(&_HomeBridge.CallOpts, _recipient)
}

// TotalExecutedPerUser is a free data retrieval call binding the contract method 0xe12a799f.
//
// Solidity: function totalExecutedPerUser(address _recipient) constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) TotalExecutedPerUser(_recipient common.Address) (*big.Int, error) {
	return _HomeBridge.Contract.TotalExecutedPerUser(&_HomeBridge.CallOpts, _recipient)
}

// TotalTXExecutedPerDayPerUser is a free data retrieval call binding the contract method 0x5ccfb7de.
//
// Solidity: function totalTXExecutedPerDayPerUser(uint256 _day, address _addr) constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) TotalTXExecutedPerDayPerUser(opts *bind.CallOpts, _day *big.Int, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "totalTXExecutedPerDayPerUser", _day, _addr)
	return *ret0, err
}

// TotalTXExecutedPerDayPerUser is a free data retrieval call binding the contract method 0x5ccfb7de.
//
// Solidity: function totalTXExecutedPerDayPerUser(uint256 _day, address _addr) constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) TotalTXExecutedPerDayPerUser(_day *big.Int, _addr common.Address) (*big.Int, error) {
	return _HomeBridge.Contract.TotalTXExecutedPerDayPerUser(&_HomeBridge.CallOpts, _day, _addr)
}

// TotalTXExecutedPerDayPerUser is a free data retrieval call binding the contract method 0x5ccfb7de.
//
// Solidity: function totalTXExecutedPerDayPerUser(uint256 _day, address _addr) constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) TotalTXExecutedPerDayPerUser(_day *big.Int, _addr common.Address) (*big.Int, error) {
	return _HomeBridge.Contract.TotalTXExecutedPerDayPerUser(&_HomeBridge.CallOpts, _day, _addr)
}

// TxPerDay is a free data retrieval call binding the contract method 0x67deef47.
//
// Solidity: function txPerDay(bytes32 ) constant returns(uint256)
func (_HomeBridge *HomeBridgeCaller) TxPerDay(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HomeBridge.contract.Call(opts, out, "txPerDay", arg0)
	return *ret0, err
}

// TxPerDay is a free data retrieval call binding the contract method 0x67deef47.
//
// Solidity: function txPerDay(bytes32 ) constant returns(uint256)
func (_HomeBridge *HomeBridgeSession) TxPerDay(arg0 [32]byte) (*big.Int, error) {
	return _HomeBridge.Contract.TxPerDay(&_HomeBridge.CallOpts, arg0)
}

// TxPerDay is a free data retrieval call binding the contract method 0x67deef47.
//
// Solidity: function txPerDay(bytes32 ) constant returns(uint256)
func (_HomeBridge *HomeBridgeCallerSession) TxPerDay(arg0 [32]byte) (*big.Int, error) {
	return _HomeBridge.Contract.TxPerDay(&_HomeBridge.CallOpts, arg0)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0xb4e2216e.
//
// Solidity: function executeTransfer(address _recipient, uint256 _value, bytes32 _txHash) returns(bool)
func (_HomeBridge *HomeBridgeTransactor) ExecuteTransfer(opts *bind.TransactOpts, _recipient common.Address, _value *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "executeTransfer", _recipient, _value, _txHash)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0xb4e2216e.
//
// Solidity: function executeTransfer(address _recipient, uint256 _value, bytes32 _txHash) returns(bool)
func (_HomeBridge *HomeBridgeSession) ExecuteTransfer(_recipient common.Address, _value *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _HomeBridge.Contract.ExecuteTransfer(&_HomeBridge.TransactOpts, _recipient, _value, _txHash)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0xb4e2216e.
//
// Solidity: function executeTransfer(address _recipient, uint256 _value, bytes32 _txHash) returns(bool)
func (_HomeBridge *HomeBridgeTransactorSession) ExecuteTransfer(_recipient common.Address, _value *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _HomeBridge.Contract.ExecuteTransfer(&_HomeBridge.TransactOpts, _recipient, _value, _txHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HomeBridge *HomeBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HomeBridge *HomeBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _HomeBridge.Contract.RenounceOwnership(&_HomeBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HomeBridge *HomeBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HomeBridge.Contract.RenounceOwnership(&_HomeBridge.TransactOpts)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _dailyLimit) returns()
func (_HomeBridge *HomeBridgeTransactor) SetDailyLimit(opts *bind.TransactOpts, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "setDailyLimit", _dailyLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _dailyLimit) returns()
func (_HomeBridge *HomeBridgeSession) SetDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetDailyLimit(&_HomeBridge.TransactOpts, _dailyLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _dailyLimit) returns()
func (_HomeBridge *HomeBridgeTransactorSession) SetDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetDailyLimit(&_HomeBridge.TransactOpts, _dailyLimit)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 _lastBlock) returns()
func (_HomeBridge *HomeBridgeTransactor) SetLastBlock(opts *bind.TransactOpts, _lastBlock *big.Int) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "setLastBlock", _lastBlock)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 _lastBlock) returns()
func (_HomeBridge *HomeBridgeSession) SetLastBlock(_lastBlock *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetLastBlock(&_HomeBridge.TransactOpts, _lastBlock)
}

// SetLastBlock is a paid mutator transaction binding the contract method 0x5d974a66.
//
// Solidity: function setLastBlock(uint256 _lastBlock) returns()
func (_HomeBridge *HomeBridgeTransactorSession) SetLastBlock(_lastBlock *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetLastBlock(&_HomeBridge.TransactOpts, _lastBlock)
}

// SetLimit is a paid mutator transaction binding the contract method 0x27ea6f2b.
//
// Solidity: function setLimit(uint256 _limit) returns()
func (_HomeBridge *HomeBridgeTransactor) SetLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "setLimit", _limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0x27ea6f2b.
//
// Solidity: function setLimit(uint256 _limit) returns()
func (_HomeBridge *HomeBridgeSession) SetLimit(_limit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetLimit(&_HomeBridge.TransactOpts, _limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0x27ea6f2b.
//
// Solidity: function setLimit(uint256 _limit) returns()
func (_HomeBridge *HomeBridgeTransactorSession) SetLimit(_limit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetLimit(&_HomeBridge.TransactOpts, _limit)
}

// SetMinimalTXLimit is a paid mutator transaction binding the contract method 0xb4634bbe.
//
// Solidity: function setMinimalTXLimit(uint256 _txMinLimit) returns()
func (_HomeBridge *HomeBridgeTransactor) SetMinimalTXLimit(opts *bind.TransactOpts, _txMinLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "setMinimalTXLimit", _txMinLimit)
}

// SetMinimalTXLimit is a paid mutator transaction binding the contract method 0xb4634bbe.
//
// Solidity: function setMinimalTXLimit(uint256 _txMinLimit) returns()
func (_HomeBridge *HomeBridgeSession) SetMinimalTXLimit(_txMinLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetMinimalTXLimit(&_HomeBridge.TransactOpts, _txMinLimit)
}

// SetMinimalTXLimit is a paid mutator transaction binding the contract method 0xb4634bbe.
//
// Solidity: function setMinimalTXLimit(uint256 _txMinLimit) returns()
func (_HomeBridge *HomeBridgeTransactorSession) SetMinimalTXLimit(_txMinLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetMinimalTXLimit(&_HomeBridge.TransactOpts, _txMinLimit)
}

// SetTXLimit is a paid mutator transaction binding the contract method 0xd1320f7b.
//
// Solidity: function setTXLimit(uint256 _txLimit) returns()
func (_HomeBridge *HomeBridgeTransactor) SetTXLimit(opts *bind.TransactOpts, _txLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "setTXLimit", _txLimit)
}

// SetTXLimit is a paid mutator transaction binding the contract method 0xd1320f7b.
//
// Solidity: function setTXLimit(uint256 _txLimit) returns()
func (_HomeBridge *HomeBridgeSession) SetTXLimit(_txLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetTXLimit(&_HomeBridge.TransactOpts, _txLimit)
}

// SetTXLimit is a paid mutator transaction binding the contract method 0xd1320f7b.
//
// Solidity: function setTXLimit(uint256 _txLimit) returns()
func (_HomeBridge *HomeBridgeTransactorSession) SetTXLimit(_txLimit *big.Int) (*types.Transaction, error) {
	return _HomeBridge.Contract.SetTXLimit(&_HomeBridge.TransactOpts, _txLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HomeBridge *HomeBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HomeBridge *HomeBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HomeBridge.Contract.TransferOwnership(&_HomeBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HomeBridge *HomeBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HomeBridge.Contract.TransferOwnership(&_HomeBridge.TransactOpts, newOwner)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_HomeBridge *HomeBridgeTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HomeBridge.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_HomeBridge *HomeBridgeSession) WithdrawAll() (*types.Transaction, error) {
	return _HomeBridge.Contract.WithdrawAll(&_HomeBridge.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_HomeBridge *HomeBridgeTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _HomeBridge.Contract.WithdrawAll(&_HomeBridge.TransactOpts)
}

// HomeBridgeBridgeFundedIterator is returned from FilterBridgeFunded and is used to iterate over the raw logs and unpacked data for BridgeFunded events raised by the HomeBridge contract.
type HomeBridgeBridgeFundedIterator struct {
	Event *HomeBridgeBridgeFunded // Event containing the contract specifics and raw log

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
func (it *HomeBridgeBridgeFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeBridgeBridgeFunded)
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
		it.Event = new(HomeBridgeBridgeFunded)
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
func (it *HomeBridgeBridgeFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeBridgeBridgeFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeBridgeBridgeFunded represents a BridgeFunded event raised by the HomeBridge contract.
type HomeBridgeBridgeFunded struct {
	Funder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeFunded is a free log retrieval operation binding the contract event 0xa1dd607db049dad651d4ebbc3b099a1cc910f208c0e692ee548615911da5e26c.
//
// Solidity: event BridgeFunded(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) FilterBridgeFunded(opts *bind.FilterOpts, funder []common.Address, value []*big.Int) (*HomeBridgeBridgeFundedIterator, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _HomeBridge.contract.FilterLogs(opts, "BridgeFunded", funderRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeBridgeFundedIterator{contract: _HomeBridge.contract, event: "BridgeFunded", logs: logs, sub: sub}, nil
}

// WatchBridgeFunded is a free log subscription operation binding the contract event 0xa1dd607db049dad651d4ebbc3b099a1cc910f208c0e692ee548615911da5e26c.
//
// Solidity: event BridgeFunded(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) WatchBridgeFunded(opts *bind.WatchOpts, sink chan<- *HomeBridgeBridgeFunded, funder []common.Address, value []*big.Int) (event.Subscription, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _HomeBridge.contract.WatchLogs(opts, "BridgeFunded", funderRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeBridgeBridgeFunded)
				if err := _HomeBridge.contract.UnpackLog(event, "BridgeFunded", log); err != nil {
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

// ParseBridgeFunded is a log parse operation binding the contract event 0xa1dd607db049dad651d4ebbc3b099a1cc910f208c0e692ee548615911da5e26c.
//
// Solidity: event BridgeFunded(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) ParseBridgeFunded(log types.Log) (*HomeBridgeBridgeFunded, error) {
	event := new(HomeBridgeBridgeFunded)
	if err := _HomeBridge.contract.UnpackLog(event, "BridgeFunded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HomeBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HomeBridge contract.
type HomeBridgeOwnershipTransferredIterator struct {
	Event *HomeBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HomeBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeBridgeOwnershipTransferred)
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
		it.Event = new(HomeBridgeOwnershipTransferred)
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
func (it *HomeBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the HomeBridge contract.
type HomeBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HomeBridge *HomeBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HomeBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HomeBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeOwnershipTransferredIterator{contract: _HomeBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HomeBridge *HomeBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HomeBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HomeBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeBridgeOwnershipTransferred)
				if err := _HomeBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HomeBridge *HomeBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*HomeBridgeOwnershipTransferred, error) {
	event := new(HomeBridgeOwnershipTransferred)
	if err := _HomeBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HomeBridgeTransferExceededLimitsIterator is returned from FilterTransferExceededLimits and is used to iterate over the raw logs and unpacked data for TransferExceededLimits events raised by the HomeBridge contract.
type HomeBridgeTransferExceededLimitsIterator struct {
	Event *HomeBridgeTransferExceededLimits // Event containing the contract specifics and raw log

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
func (it *HomeBridgeTransferExceededLimitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeBridgeTransferExceededLimits)
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
		it.Event = new(HomeBridgeTransferExceededLimits)
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
func (it *HomeBridgeTransferExceededLimitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeBridgeTransferExceededLimitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeBridgeTransferExceededLimits represents a TransferExceededLimits event raised by the HomeBridge contract.
type HomeBridgeTransferExceededLimits struct {
	Funder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferExceededLimits is a free log retrieval operation binding the contract event 0x459606ba395da8841ce1bcd671f89a8ff0c01e27d7ee62af68b834b433c4e535.
//
// Solidity: event TransferExceededLimits(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) FilterTransferExceededLimits(opts *bind.FilterOpts, funder []common.Address, value []*big.Int) (*HomeBridgeTransferExceededLimitsIterator, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _HomeBridge.contract.FilterLogs(opts, "TransferExceededLimits", funderRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeTransferExceededLimitsIterator{contract: _HomeBridge.contract, event: "TransferExceededLimits", logs: logs, sub: sub}, nil
}

// WatchTransferExceededLimits is a free log subscription operation binding the contract event 0x459606ba395da8841ce1bcd671f89a8ff0c01e27d7ee62af68b834b433c4e535.
//
// Solidity: event TransferExceededLimits(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) WatchTransferExceededLimits(opts *bind.WatchOpts, sink chan<- *HomeBridgeTransferExceededLimits, funder []common.Address, value []*big.Int) (event.Subscription, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _HomeBridge.contract.WatchLogs(opts, "TransferExceededLimits", funderRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeBridgeTransferExceededLimits)
				if err := _HomeBridge.contract.UnpackLog(event, "TransferExceededLimits", log); err != nil {
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

// ParseTransferExceededLimits is a log parse operation binding the contract event 0x459606ba395da8841ce1bcd671f89a8ff0c01e27d7ee62af68b834b433c4e535.
//
// Solidity: event TransferExceededLimits(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) ParseTransferExceededLimits(log types.Log) (*HomeBridgeTransferExceededLimits, error) {
	event := new(HomeBridgeTransferExceededLimits)
	if err := _HomeBridge.contract.UnpackLog(event, "TransferExceededLimits", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HomeBridgeTransferExecutedIterator is returned from FilterTransferExecuted and is used to iterate over the raw logs and unpacked data for TransferExecuted events raised by the HomeBridge contract.
type HomeBridgeTransferExecutedIterator struct {
	Event *HomeBridgeTransferExecuted // Event containing the contract specifics and raw log

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
func (it *HomeBridgeTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeBridgeTransferExecuted)
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
		it.Event = new(HomeBridgeTransferExecuted)
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
func (it *HomeBridgeTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeBridgeTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeBridgeTransferExecuted represents a TransferExecuted event raised by the HomeBridge contract.
type HomeBridgeTransferExecuted struct {
	Recipient common.Address
	Value     *big.Int
	TxHash    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferExecuted is a free log retrieval operation binding the contract event 0x6eb822c246c0021523c4209516767730f4b9260f24858135daa61b0377568fd2.
//
// Solidity: event TransferExecuted(address indexed recipient, uint256 indexed value, bytes32 indexed txHash)
func (_HomeBridge *HomeBridgeFilterer) FilterTransferExecuted(opts *bind.FilterOpts, recipient []common.Address, value []*big.Int, txHash [][32]byte) (*HomeBridgeTransferExecutedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _HomeBridge.contract.FilterLogs(opts, "TransferExecuted", recipientRule, valueRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeTransferExecutedIterator{contract: _HomeBridge.contract, event: "TransferExecuted", logs: logs, sub: sub}, nil
}

// WatchTransferExecuted is a free log subscription operation binding the contract event 0x6eb822c246c0021523c4209516767730f4b9260f24858135daa61b0377568fd2.
//
// Solidity: event TransferExecuted(address indexed recipient, uint256 indexed value, bytes32 indexed txHash)
func (_HomeBridge *HomeBridgeFilterer) WatchTransferExecuted(opts *bind.WatchOpts, sink chan<- *HomeBridgeTransferExecuted, recipient []common.Address, value []*big.Int, txHash [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _HomeBridge.contract.WatchLogs(opts, "TransferExecuted", recipientRule, valueRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeBridgeTransferExecuted)
				if err := _HomeBridge.contract.UnpackLog(event, "TransferExecuted", log); err != nil {
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

// ParseTransferExecuted is a log parse operation binding the contract event 0x6eb822c246c0021523c4209516767730f4b9260f24858135daa61b0377568fd2.
//
// Solidity: event TransferExecuted(address indexed recipient, uint256 indexed value, bytes32 indexed txHash)
func (_HomeBridge *HomeBridgeFilterer) ParseTransferExecuted(log types.Log) (*HomeBridgeTransferExecuted, error) {
	event := new(HomeBridgeTransferExecuted)
	if err := _HomeBridge.contract.UnpackLog(event, "TransferExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HomeBridgeTransferRejectedIterator is returned from FilterTransferRejected and is used to iterate over the raw logs and unpacked data for TransferRejected events raised by the HomeBridge contract.
type HomeBridgeTransferRejectedIterator struct {
	Event *HomeBridgeTransferRejected // Event containing the contract specifics and raw log

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
func (it *HomeBridgeTransferRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeBridgeTransferRejected)
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
		it.Event = new(HomeBridgeTransferRejected)
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
func (it *HomeBridgeTransferRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeBridgeTransferRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeBridgeTransferRejected represents a TransferRejected event raised by the HomeBridge contract.
type HomeBridgeTransferRejected struct {
	Recipient common.Address
	Value     *big.Int
	TxHash    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferRejected is a free log retrieval operation binding the contract event 0x8c17b1c1874fa559466f835c95f0fbce9b80e6c948a17b4dbcb8785de90fc14b.
//
// Solidity: event TransferRejected(address indexed recipient, uint256 indexed value, bytes32 indexed txHash)
func (_HomeBridge *HomeBridgeFilterer) FilterTransferRejected(opts *bind.FilterOpts, recipient []common.Address, value []*big.Int, txHash [][32]byte) (*HomeBridgeTransferRejectedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _HomeBridge.contract.FilterLogs(opts, "TransferRejected", recipientRule, valueRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeTransferRejectedIterator{contract: _HomeBridge.contract, event: "TransferRejected", logs: logs, sub: sub}, nil
}

// WatchTransferRejected is a free log subscription operation binding the contract event 0x8c17b1c1874fa559466f835c95f0fbce9b80e6c948a17b4dbcb8785de90fc14b.
//
// Solidity: event TransferRejected(address indexed recipient, uint256 indexed value, bytes32 indexed txHash)
func (_HomeBridge *HomeBridgeFilterer) WatchTransferRejected(opts *bind.WatchOpts, sink chan<- *HomeBridgeTransferRejected, recipient []common.Address, value []*big.Int, txHash [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _HomeBridge.contract.WatchLogs(opts, "TransferRejected", recipientRule, valueRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeBridgeTransferRejected)
				if err := _HomeBridge.contract.UnpackLog(event, "TransferRejected", log); err != nil {
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

// ParseTransferRejected is a log parse operation binding the contract event 0x8c17b1c1874fa559466f835c95f0fbce9b80e6c948a17b4dbcb8785de90fc14b.
//
// Solidity: event TransferRejected(address indexed recipient, uint256 indexed value, bytes32 indexed txHash)
func (_HomeBridge *HomeBridgeFilterer) ParseTransferRejected(log types.Log) (*HomeBridgeTransferRejected, error) {
	event := new(HomeBridgeTransferRejected)
	if err := _HomeBridge.contract.UnpackLog(event, "TransferRejected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HomeBridgeTransferWithinLimitsIterator is returned from FilterTransferWithinLimits and is used to iterate over the raw logs and unpacked data for TransferWithinLimits events raised by the HomeBridge contract.
type HomeBridgeTransferWithinLimitsIterator struct {
	Event *HomeBridgeTransferWithinLimits // Event containing the contract specifics and raw log

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
func (it *HomeBridgeTransferWithinLimitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeBridgeTransferWithinLimits)
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
		it.Event = new(HomeBridgeTransferWithinLimits)
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
func (it *HomeBridgeTransferWithinLimitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeBridgeTransferWithinLimitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeBridgeTransferWithinLimits represents a TransferWithinLimits event raised by the HomeBridge contract.
type HomeBridgeTransferWithinLimits struct {
	Funder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferWithinLimits is a free log retrieval operation binding the contract event 0x9db0cf87f9e921b7bc49faa093f96209fe0bee2fadd2ef50f65fee2cbaa78be3.
//
// Solidity: event TransferWithinLimits(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) FilterTransferWithinLimits(opts *bind.FilterOpts, funder []common.Address, value []*big.Int) (*HomeBridgeTransferWithinLimitsIterator, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _HomeBridge.contract.FilterLogs(opts, "TransferWithinLimits", funderRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &HomeBridgeTransferWithinLimitsIterator{contract: _HomeBridge.contract, event: "TransferWithinLimits", logs: logs, sub: sub}, nil
}

// WatchTransferWithinLimits is a free log subscription operation binding the contract event 0x9db0cf87f9e921b7bc49faa093f96209fe0bee2fadd2ef50f65fee2cbaa78be3.
//
// Solidity: event TransferWithinLimits(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) WatchTransferWithinLimits(opts *bind.WatchOpts, sink chan<- *HomeBridgeTransferWithinLimits, funder []common.Address, value []*big.Int) (event.Subscription, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _HomeBridge.contract.WatchLogs(opts, "TransferWithinLimits", funderRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeBridgeTransferWithinLimits)
				if err := _HomeBridge.contract.UnpackLog(event, "TransferWithinLimits", log); err != nil {
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

// ParseTransferWithinLimits is a log parse operation binding the contract event 0x9db0cf87f9e921b7bc49faa093f96209fe0bee2fadd2ef50f65fee2cbaa78be3.
//
// Solidity: event TransferWithinLimits(address indexed funder, uint256 indexed value)
func (_HomeBridge *HomeBridgeFilterer) ParseTransferWithinLimits(log types.Log) (*HomeBridgeTransferWithinLimits, error) {
	event := new(HomeBridgeTransferWithinLimits)
	if err := _HomeBridge.contract.UnpackLog(event, "TransferWithinLimits", log); err != nil {
		return nil, err
	}
	return event, nil
}
