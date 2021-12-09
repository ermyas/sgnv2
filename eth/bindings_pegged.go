// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DelayedTransferMetaData contains all meta data concerning the DelayedTransfer contract.
var DelayedTransferMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DelayThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedTransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelayedTransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delayThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setDelayThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DelayedTransferABI is the input ABI used to generate the binding from.
// Deprecated: Use DelayedTransferMetaData.ABI instead.
var DelayedTransferABI = DelayedTransferMetaData.ABI

// DelayedTransfer is an auto generated Go binding around an Ethereum contract.
type DelayedTransfer struct {
	DelayedTransferCaller     // Read-only binding to the contract
	DelayedTransferTransactor // Write-only binding to the contract
	DelayedTransferFilterer   // Log filterer for contract events
}

// DelayedTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayedTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayedTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelayedTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayedTransferSession struct {
	Contract     *DelayedTransfer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelayedTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayedTransferCallerSession struct {
	Contract *DelayedTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DelayedTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayedTransferTransactorSession struct {
	Contract     *DelayedTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DelayedTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayedTransferRaw struct {
	Contract *DelayedTransfer // Generic contract binding to access the raw methods on
}

// DelayedTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayedTransferCallerRaw struct {
	Contract *DelayedTransferCaller // Generic read-only contract binding to access the raw methods on
}

// DelayedTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayedTransferTransactorRaw struct {
	Contract *DelayedTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayedTransfer creates a new instance of DelayedTransfer, bound to a specific deployed contract.
func NewDelayedTransfer(address common.Address, backend bind.ContractBackend) (*DelayedTransfer, error) {
	contract, err := bindDelayedTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelayedTransfer{DelayedTransferCaller: DelayedTransferCaller{contract: contract}, DelayedTransferTransactor: DelayedTransferTransactor{contract: contract}, DelayedTransferFilterer: DelayedTransferFilterer{contract: contract}}, nil
}

// NewDelayedTransferCaller creates a new read-only instance of DelayedTransfer, bound to a specific deployed contract.
func NewDelayedTransferCaller(address common.Address, caller bind.ContractCaller) (*DelayedTransferCaller, error) {
	contract, err := bindDelayedTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedTransferCaller{contract: contract}, nil
}

// NewDelayedTransferTransactor creates a new write-only instance of DelayedTransfer, bound to a specific deployed contract.
func NewDelayedTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayedTransferTransactor, error) {
	contract, err := bindDelayedTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedTransferTransactor{contract: contract}, nil
}

// NewDelayedTransferFilterer creates a new log filterer instance of DelayedTransfer, bound to a specific deployed contract.
func NewDelayedTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*DelayedTransferFilterer, error) {
	contract, err := bindDelayedTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelayedTransferFilterer{contract: contract}, nil
}

// bindDelayedTransfer binds a generic wrapper to an already deployed contract.
func bindDelayedTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelayedTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedTransfer *DelayedTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedTransfer.Contract.DelayedTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedTransfer *DelayedTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.DelayedTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedTransfer *DelayedTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.DelayedTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedTransfer *DelayedTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedTransfer *DelayedTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedTransfer *DelayedTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.contract.Transact(opts, method, params...)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_DelayedTransfer *DelayedTransferCaller) DelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedTransfer.contract.Call(opts, &out, "delayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_DelayedTransfer *DelayedTransferSession) DelayPeriod() (*big.Int, error) {
	return _DelayedTransfer.Contract.DelayPeriod(&_DelayedTransfer.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_DelayedTransfer *DelayedTransferCallerSession) DelayPeriod() (*big.Int, error) {
	return _DelayedTransfer.Contract.DelayPeriod(&_DelayedTransfer.CallOpts)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_DelayedTransfer *DelayedTransferCaller) DelayThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelayedTransfer.contract.Call(opts, &out, "delayThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_DelayedTransfer *DelayedTransferSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _DelayedTransfer.Contract.DelayThresholds(&_DelayedTransfer.CallOpts, arg0)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_DelayedTransfer *DelayedTransferCallerSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _DelayedTransfer.Contract.DelayThresholds(&_DelayedTransfer.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_DelayedTransfer *DelayedTransferCaller) DelayedTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _DelayedTransfer.contract.Call(opts, &out, "delayedTransfers", arg0)

	outstruct := new(struct {
		Receiver  common.Address
		Token     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_DelayedTransfer *DelayedTransferSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _DelayedTransfer.Contract.DelayedTransfers(&_DelayedTransfer.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_DelayedTransfer *DelayedTransferCallerSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _DelayedTransfer.Contract.DelayedTransfers(&_DelayedTransfer.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_DelayedTransfer *DelayedTransferCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DelayedTransfer.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_DelayedTransfer *DelayedTransferSession) Governors(arg0 common.Address) (bool, error) {
	return _DelayedTransfer.Contract.Governors(&_DelayedTransfer.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_DelayedTransfer *DelayedTransferCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _DelayedTransfer.Contract.Governors(&_DelayedTransfer.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_DelayedTransfer *DelayedTransferCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _DelayedTransfer.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_DelayedTransfer *DelayedTransferSession) IsGovernor(_account common.Address) (bool, error) {
	return _DelayedTransfer.Contract.IsGovernor(&_DelayedTransfer.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_DelayedTransfer *DelayedTransferCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _DelayedTransfer.Contract.IsGovernor(&_DelayedTransfer.CallOpts, _account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedTransfer *DelayedTransferCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedTransfer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedTransfer *DelayedTransferSession) Owner() (common.Address, error) {
	return _DelayedTransfer.Contract.Owner(&_DelayedTransfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedTransfer *DelayedTransferCallerSession) Owner() (common.Address, error) {
	return _DelayedTransfer.Contract.Owner(&_DelayedTransfer.CallOpts)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_DelayedTransfer *DelayedTransferTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_DelayedTransfer *DelayedTransferSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.AddGovernor(&_DelayedTransfer.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_DelayedTransfer *DelayedTransferTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.AddGovernor(&_DelayedTransfer.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_DelayedTransfer *DelayedTransferTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_DelayedTransfer *DelayedTransferSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.RemoveGovernor(&_DelayedTransfer.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_DelayedTransfer *DelayedTransferTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.RemoveGovernor(&_DelayedTransfer.TransactOpts, _account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_DelayedTransfer *DelayedTransferTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedTransfer.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_DelayedTransfer *DelayedTransferSession) RenounceGovernor() (*types.Transaction, error) {
	return _DelayedTransfer.Contract.RenounceGovernor(&_DelayedTransfer.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_DelayedTransfer *DelayedTransferTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _DelayedTransfer.Contract.RenounceGovernor(&_DelayedTransfer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedTransfer *DelayedTransferTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedTransfer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedTransfer *DelayedTransferSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelayedTransfer.Contract.RenounceOwnership(&_DelayedTransfer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedTransfer *DelayedTransferTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelayedTransfer.Contract.RenounceOwnership(&_DelayedTransfer.TransactOpts)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_DelayedTransfer *DelayedTransferTransactor) SetDelayPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _DelayedTransfer.contract.Transact(opts, "setDelayPeriod", _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_DelayedTransfer *DelayedTransferSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.SetDelayPeriod(&_DelayedTransfer.TransactOpts, _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_DelayedTransfer *DelayedTransferTransactorSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.SetDelayPeriod(&_DelayedTransfer.TransactOpts, _period)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_DelayedTransfer *DelayedTransferTransactor) SetDelayThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _DelayedTransfer.contract.Transact(opts, "setDelayThresholds", _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_DelayedTransfer *DelayedTransferSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.SetDelayThresholds(&_DelayedTransfer.TransactOpts, _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_DelayedTransfer *DelayedTransferTransactorSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.SetDelayThresholds(&_DelayedTransfer.TransactOpts, _tokens, _thresholds)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedTransfer *DelayedTransferTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedTransfer *DelayedTransferSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.TransferOwnership(&_DelayedTransfer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedTransfer *DelayedTransferTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedTransfer.Contract.TransferOwnership(&_DelayedTransfer.TransactOpts, newOwner)
}

// DelayedTransferDelayPeriodUpdatedIterator is returned from FilterDelayPeriodUpdated and is used to iterate over the raw logs and unpacked data for DelayPeriodUpdated events raised by the DelayedTransfer contract.
type DelayedTransferDelayPeriodUpdatedIterator struct {
	Event *DelayedTransferDelayPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *DelayedTransferDelayPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedTransferDelayPeriodUpdated)
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
		it.Event = new(DelayedTransferDelayPeriodUpdated)
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
func (it *DelayedTransferDelayPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedTransferDelayPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedTransferDelayPeriodUpdated represents a DelayPeriodUpdated event raised by the DelayedTransfer contract.
type DelayedTransferDelayPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelayPeriodUpdated is a free log retrieval operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_DelayedTransfer *DelayedTransferFilterer) FilterDelayPeriodUpdated(opts *bind.FilterOpts) (*DelayedTransferDelayPeriodUpdatedIterator, error) {

	logs, sub, err := _DelayedTransfer.contract.FilterLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &DelayedTransferDelayPeriodUpdatedIterator{contract: _DelayedTransfer.contract, event: "DelayPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayPeriodUpdated is a free log subscription operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_DelayedTransfer *DelayedTransferFilterer) WatchDelayPeriodUpdated(opts *bind.WatchOpts, sink chan<- *DelayedTransferDelayPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _DelayedTransfer.contract.WatchLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedTransferDelayPeriodUpdated)
				if err := _DelayedTransfer.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
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

// ParseDelayPeriodUpdated is a log parse operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_DelayedTransfer *DelayedTransferFilterer) ParseDelayPeriodUpdated(log types.Log) (*DelayedTransferDelayPeriodUpdated, error) {
	event := new(DelayedTransferDelayPeriodUpdated)
	if err := _DelayedTransfer.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedTransferDelayThresholdUpdatedIterator is returned from FilterDelayThresholdUpdated and is used to iterate over the raw logs and unpacked data for DelayThresholdUpdated events raised by the DelayedTransfer contract.
type DelayedTransferDelayThresholdUpdatedIterator struct {
	Event *DelayedTransferDelayThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *DelayedTransferDelayThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedTransferDelayThresholdUpdated)
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
		it.Event = new(DelayedTransferDelayThresholdUpdated)
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
func (it *DelayedTransferDelayThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedTransferDelayThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedTransferDelayThresholdUpdated represents a DelayThresholdUpdated event raised by the DelayedTransfer contract.
type DelayedTransferDelayThresholdUpdated struct {
	Token     common.Address
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayThresholdUpdated is a free log retrieval operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_DelayedTransfer *DelayedTransferFilterer) FilterDelayThresholdUpdated(opts *bind.FilterOpts) (*DelayedTransferDelayThresholdUpdatedIterator, error) {

	logs, sub, err := _DelayedTransfer.contract.FilterLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &DelayedTransferDelayThresholdUpdatedIterator{contract: _DelayedTransfer.contract, event: "DelayThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayThresholdUpdated is a free log subscription operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_DelayedTransfer *DelayedTransferFilterer) WatchDelayThresholdUpdated(opts *bind.WatchOpts, sink chan<- *DelayedTransferDelayThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _DelayedTransfer.contract.WatchLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedTransferDelayThresholdUpdated)
				if err := _DelayedTransfer.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
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

// ParseDelayThresholdUpdated is a log parse operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_DelayedTransfer *DelayedTransferFilterer) ParseDelayThresholdUpdated(log types.Log) (*DelayedTransferDelayThresholdUpdated, error) {
	event := new(DelayedTransferDelayThresholdUpdated)
	if err := _DelayedTransfer.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedTransferDelayedTransferAddedIterator is returned from FilterDelayedTransferAdded and is used to iterate over the raw logs and unpacked data for DelayedTransferAdded events raised by the DelayedTransfer contract.
type DelayedTransferDelayedTransferAddedIterator struct {
	Event *DelayedTransferDelayedTransferAdded // Event containing the contract specifics and raw log

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
func (it *DelayedTransferDelayedTransferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedTransferDelayedTransferAdded)
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
		it.Event = new(DelayedTransferDelayedTransferAdded)
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
func (it *DelayedTransferDelayedTransferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedTransferDelayedTransferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedTransferDelayedTransferAdded represents a DelayedTransferAdded event raised by the DelayedTransfer contract.
type DelayedTransferDelayedTransferAdded struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferAdded is a free log retrieval operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_DelayedTransfer *DelayedTransferFilterer) FilterDelayedTransferAdded(opts *bind.FilterOpts) (*DelayedTransferDelayedTransferAddedIterator, error) {

	logs, sub, err := _DelayedTransfer.contract.FilterLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return &DelayedTransferDelayedTransferAddedIterator{contract: _DelayedTransfer.contract, event: "DelayedTransferAdded", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferAdded is a free log subscription operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_DelayedTransfer *DelayedTransferFilterer) WatchDelayedTransferAdded(opts *bind.WatchOpts, sink chan<- *DelayedTransferDelayedTransferAdded) (event.Subscription, error) {

	logs, sub, err := _DelayedTransfer.contract.WatchLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedTransferDelayedTransferAdded)
				if err := _DelayedTransfer.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
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

// ParseDelayedTransferAdded is a log parse operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_DelayedTransfer *DelayedTransferFilterer) ParseDelayedTransferAdded(log types.Log) (*DelayedTransferDelayedTransferAdded, error) {
	event := new(DelayedTransferDelayedTransferAdded)
	if err := _DelayedTransfer.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedTransferDelayedTransferExecutedIterator is returned from FilterDelayedTransferExecuted and is used to iterate over the raw logs and unpacked data for DelayedTransferExecuted events raised by the DelayedTransfer contract.
type DelayedTransferDelayedTransferExecutedIterator struct {
	Event *DelayedTransferDelayedTransferExecuted // Event containing the contract specifics and raw log

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
func (it *DelayedTransferDelayedTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedTransferDelayedTransferExecuted)
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
		it.Event = new(DelayedTransferDelayedTransferExecuted)
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
func (it *DelayedTransferDelayedTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedTransferDelayedTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedTransferDelayedTransferExecuted represents a DelayedTransferExecuted event raised by the DelayedTransfer contract.
type DelayedTransferDelayedTransferExecuted struct {
	Id       [32]byte
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferExecuted is a free log retrieval operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_DelayedTransfer *DelayedTransferFilterer) FilterDelayedTransferExecuted(opts *bind.FilterOpts) (*DelayedTransferDelayedTransferExecutedIterator, error) {

	logs, sub, err := _DelayedTransfer.contract.FilterLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return &DelayedTransferDelayedTransferExecutedIterator{contract: _DelayedTransfer.contract, event: "DelayedTransferExecuted", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferExecuted is a free log subscription operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_DelayedTransfer *DelayedTransferFilterer) WatchDelayedTransferExecuted(opts *bind.WatchOpts, sink chan<- *DelayedTransferDelayedTransferExecuted) (event.Subscription, error) {

	logs, sub, err := _DelayedTransfer.contract.WatchLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedTransferDelayedTransferExecuted)
				if err := _DelayedTransfer.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
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

// ParseDelayedTransferExecuted is a log parse operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_DelayedTransfer *DelayedTransferFilterer) ParseDelayedTransferExecuted(log types.Log) (*DelayedTransferDelayedTransferExecuted, error) {
	event := new(DelayedTransferDelayedTransferExecuted)
	if err := _DelayedTransfer.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedTransferGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the DelayedTransfer contract.
type DelayedTransferGovernorAddedIterator struct {
	Event *DelayedTransferGovernorAdded // Event containing the contract specifics and raw log

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
func (it *DelayedTransferGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedTransferGovernorAdded)
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
		it.Event = new(DelayedTransferGovernorAdded)
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
func (it *DelayedTransferGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedTransferGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedTransferGovernorAdded represents a GovernorAdded event raised by the DelayedTransfer contract.
type DelayedTransferGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_DelayedTransfer *DelayedTransferFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*DelayedTransferGovernorAddedIterator, error) {

	logs, sub, err := _DelayedTransfer.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &DelayedTransferGovernorAddedIterator{contract: _DelayedTransfer.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_DelayedTransfer *DelayedTransferFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *DelayedTransferGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _DelayedTransfer.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedTransferGovernorAdded)
				if err := _DelayedTransfer.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_DelayedTransfer *DelayedTransferFilterer) ParseGovernorAdded(log types.Log) (*DelayedTransferGovernorAdded, error) {
	event := new(DelayedTransferGovernorAdded)
	if err := _DelayedTransfer.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedTransferGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the DelayedTransfer contract.
type DelayedTransferGovernorRemovedIterator struct {
	Event *DelayedTransferGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *DelayedTransferGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedTransferGovernorRemoved)
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
		it.Event = new(DelayedTransferGovernorRemoved)
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
func (it *DelayedTransferGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedTransferGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedTransferGovernorRemoved represents a GovernorRemoved event raised by the DelayedTransfer contract.
type DelayedTransferGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_DelayedTransfer *DelayedTransferFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*DelayedTransferGovernorRemovedIterator, error) {

	logs, sub, err := _DelayedTransfer.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &DelayedTransferGovernorRemovedIterator{contract: _DelayedTransfer.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_DelayedTransfer *DelayedTransferFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *DelayedTransferGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _DelayedTransfer.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedTransferGovernorRemoved)
				if err := _DelayedTransfer.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_DelayedTransfer *DelayedTransferFilterer) ParseGovernorRemoved(log types.Log) (*DelayedTransferGovernorRemoved, error) {
	event := new(DelayedTransferGovernorRemoved)
	if err := _DelayedTransfer.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedTransferOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DelayedTransfer contract.
type DelayedTransferOwnershipTransferredIterator struct {
	Event *DelayedTransferOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DelayedTransferOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedTransferOwnershipTransferred)
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
		it.Event = new(DelayedTransferOwnershipTransferred)
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
func (it *DelayedTransferOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedTransferOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedTransferOwnershipTransferred represents a OwnershipTransferred event raised by the DelayedTransfer contract.
type DelayedTransferOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedTransfer *DelayedTransferFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DelayedTransferOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelayedTransfer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DelayedTransferOwnershipTransferredIterator{contract: _DelayedTransfer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedTransfer *DelayedTransferFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DelayedTransferOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelayedTransfer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedTransferOwnershipTransferred)
				if err := _DelayedTransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DelayedTransfer *DelayedTransferFilterer) ParseOwnershipTransferred(log types.Log) (*DelayedTransferOwnershipTransferred, error) {
	event := new(DelayedTransferOwnershipTransferred)
	if err := _DelayedTransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorMetaData contains all meta data concerning the Governor contract.
var GovernorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernorMetaData.ABI instead.
var GovernorABI = GovernorMetaData.ABI

// Governor is an auto generated Go binding around an Ethereum contract.
type Governor struct {
	GovernorCaller     // Read-only binding to the contract
	GovernorTransactor // Write-only binding to the contract
	GovernorFilterer   // Log filterer for contract events
}

// GovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernorSession struct {
	Contract     *Governor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernorCallerSession struct {
	Contract *GovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernorTransactorSession struct {
	Contract     *GovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernorRaw struct {
	Contract *Governor // Generic contract binding to access the raw methods on
}

// GovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernorCallerRaw struct {
	Contract *GovernorCaller // Generic read-only contract binding to access the raw methods on
}

// GovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernorTransactorRaw struct {
	Contract *GovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernor creates a new instance of Governor, bound to a specific deployed contract.
func NewGovernor(address common.Address, backend bind.ContractBackend) (*Governor, error) {
	contract, err := bindGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governor{GovernorCaller: GovernorCaller{contract: contract}, GovernorTransactor: GovernorTransactor{contract: contract}, GovernorFilterer: GovernorFilterer{contract: contract}}, nil
}

// NewGovernorCaller creates a new read-only instance of Governor, bound to a specific deployed contract.
func NewGovernorCaller(address common.Address, caller bind.ContractCaller) (*GovernorCaller, error) {
	contract, err := bindGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernorCaller{contract: contract}, nil
}

// NewGovernorTransactor creates a new write-only instance of Governor, bound to a specific deployed contract.
func NewGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernorTransactor, error) {
	contract, err := bindGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernorTransactor{contract: contract}, nil
}

// NewGovernorFilterer creates a new log filterer instance of Governor, bound to a specific deployed contract.
func NewGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernorFilterer, error) {
	contract, err := bindGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernorFilterer{contract: contract}, nil
}

// bindGovernor binds a generic wrapper to an already deployed contract.
func bindGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governor *GovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governor.Contract.GovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governor *GovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governor.Contract.GovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governor *GovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governor.Contract.GovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governor *GovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governor *GovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governor *GovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governor.Contract.contract.Transact(opts, method, params...)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Governor *GovernorCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governor.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Governor *GovernorSession) Governors(arg0 common.Address) (bool, error) {
	return _Governor.Contract.Governors(&_Governor.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Governor *GovernorCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _Governor.Contract.Governors(&_Governor.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Governor *GovernorCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Governor.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Governor *GovernorSession) IsGovernor(_account common.Address) (bool, error) {
	return _Governor.Contract.IsGovernor(&_Governor.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Governor *GovernorCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _Governor.Contract.IsGovernor(&_Governor.CallOpts, _account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governor *GovernorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governor *GovernorSession) Owner() (common.Address, error) {
	return _Governor.Contract.Owner(&_Governor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governor *GovernorCallerSession) Owner() (common.Address, error) {
	return _Governor.Contract.Owner(&_Governor.CallOpts)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Governor *GovernorTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Governor.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Governor *GovernorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Governor.Contract.AddGovernor(&_Governor.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Governor *GovernorTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Governor.Contract.AddGovernor(&_Governor.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Governor *GovernorTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Governor.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Governor *GovernorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Governor.Contract.RemoveGovernor(&_Governor.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Governor *GovernorTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Governor.Contract.RemoveGovernor(&_Governor.TransactOpts, _account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Governor *GovernorTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governor.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Governor *GovernorSession) RenounceGovernor() (*types.Transaction, error) {
	return _Governor.Contract.RenounceGovernor(&_Governor.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Governor *GovernorTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _Governor.Contract.RenounceGovernor(&_Governor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governor *GovernorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governor *GovernorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Governor.Contract.RenounceOwnership(&_Governor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governor *GovernorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Governor.Contract.RenounceOwnership(&_Governor.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governor *GovernorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Governor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governor *GovernorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governor.Contract.TransferOwnership(&_Governor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governor *GovernorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governor.Contract.TransferOwnership(&_Governor.TransactOpts, newOwner)
}

// GovernorGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the Governor contract.
type GovernorGovernorAddedIterator struct {
	Event *GovernorGovernorAdded // Event containing the contract specifics and raw log

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
func (it *GovernorGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorGovernorAdded)
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
		it.Event = new(GovernorGovernorAdded)
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
func (it *GovernorGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorGovernorAdded represents a GovernorAdded event raised by the Governor contract.
type GovernorGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Governor *GovernorFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*GovernorGovernorAddedIterator, error) {

	logs, sub, err := _Governor.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &GovernorGovernorAddedIterator{contract: _Governor.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Governor *GovernorFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *GovernorGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _Governor.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorGovernorAdded)
				if err := _Governor.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Governor *GovernorFilterer) ParseGovernorAdded(log types.Log) (*GovernorGovernorAdded, error) {
	event := new(GovernorGovernorAdded)
	if err := _Governor.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the Governor contract.
type GovernorGovernorRemovedIterator struct {
	Event *GovernorGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *GovernorGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorGovernorRemoved)
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
		it.Event = new(GovernorGovernorRemoved)
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
func (it *GovernorGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorGovernorRemoved represents a GovernorRemoved event raised by the Governor contract.
type GovernorGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Governor *GovernorFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*GovernorGovernorRemovedIterator, error) {

	logs, sub, err := _Governor.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &GovernorGovernorRemovedIterator{contract: _Governor.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Governor *GovernorFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *GovernorGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _Governor.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorGovernorRemoved)
				if err := _Governor.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Governor *GovernorFilterer) ParseGovernorRemoved(log types.Log) (*GovernorGovernorRemoved, error) {
	event := new(GovernorGovernorRemoved)
	if err := _Governor.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Governor contract.
type GovernorOwnershipTransferredIterator struct {
	Event *GovernorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovernorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorOwnershipTransferred)
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
		it.Event = new(GovernorOwnershipTransferred)
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
func (it *GovernorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorOwnershipTransferred represents a OwnershipTransferred event raised by the Governor contract.
type GovernorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governor *GovernorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovernorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovernorOwnershipTransferredIterator{contract: _Governor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governor *GovernorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorOwnershipTransferred)
				if err := _Governor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Governor *GovernorFilterer) ParseOwnershipTransferred(log types.Log) (*GovernorOwnershipTransferred, error) {
	event := new(GovernorOwnershipTransferred)
	if err := _Governor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultMetaData contains all meta data concerning the OriginalTokenVault contract.
var OriginalTokenVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DelayThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedTransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelayedTransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"depositId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"mintChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintAccount\",\"type\":\"address\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EpochLengthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EpochVolumeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MaxDepositUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinDepositUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"refChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"burnAccount\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delayThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_mintChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_mintAccount\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumeCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"executeDelayedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastOpTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setDelayThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"name\":\"setEpochVolumeCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMaxDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620030e3380380620030e3833981016040819052620000349162000255565b6001600055620000443362000079565b6001805460ff60a01b191690556200005c33620000cb565b620000673362000195565b6001600160a01b031660805262000287565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03811660009081526002602052604090205460ff16156200013a5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064015b60405180910390fd5b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f891015b60405180910390a150565b6001600160a01b03811660009081526003602052604090205460ff1615620002005760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f720000000000604482015260640162000131565b6001600160a01b038116600081815260036020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b591016200018a565b6000602082840312156200026857600080fd5b81516001600160a01b03811681146200028057600080fd5b9392505050565b608051612e39620002aa60003960008181610557015261145c0152612e396000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c80636ef8d66d11610145578063b1c94d94116100bd578063e3eece261161008c578063eecdac8811610071578063eecdac88146105d0578063f2fde38b146105e3578063f8321383146105f657600080fd5b8063e3eece2614610581578063e43581b8146105a457600080fd5b8063b1c94d9414610529578063b5f2bc4714610532578063ccf2683b14610552578063e026049c1461057957600080fd5b80638456cb59116101145780639e25fc5c116100f95780639e25fc5c14610496578063a21a9280146104a9578063adc0d57f146104bc57600080fd5b80638456cb59146104695780638da5cb5b1461047157600080fd5b80636ef8d66d14610423578063715018a61461042b57806380f51c121461043357806382dc1ec41461045657600080fd5b806346fbf68e116101d857806357d775f8116101a75780635ec2fa261161018c5780635ec2fa26146103dd57806360216b00146103f05780636b2c0f551461041057600080fd5b806357d775f8146103c25780635c975abb146103cb57600080fd5b806346fbf68e1461035057806347b16c6c1461037c57806352532faa1461038f57806354eea796146103af57600080fd5b80633c29f8391161022f5780633d572107116102145780633d572107146103155780633f4ba83a14610328578063402d267d1461033057600080fd5b80633c29f839146102d45780633c4a25d01461030257600080fd5b806301e647251461026157806317bdbae51461029957806323463624146102ae578063303b6442146102c1575b600080fd5b61028461026f366004612851565b600b6020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6102ac6102a73660046128b6565b610616565b005b6102ac6102bc366004612956565b6107b9565b6102ac6102cf3660046128b6565b610ad1565b6102f46102e23660046129b4565b600c6020526000908152604090205481565b604051908152602001610290565b6102ac6103103660046129b4565b610c68565b6102ac610323366004612851565b610cce565b6102ac610d62565b6102f461033e3660046129b4565b600d6020526000908152604090205481565b61028461035e3660046129b4565b6001600160a01b031660009081526002602052604090205460ff1690565b6102ac61038a3660046128b6565b610dcb565b6102f461039d3660046129b4565b60096020526000908152604090205481565b6102ac6103bd366004612851565b610f62565b6102f460045481565b600154600160a01b900460ff16610284565b6102ac6103eb3660046128b6565b610fef565b6102f46103fe3660046129b4565b60056020526000908152604090205481565b6102ac61041e3660046129b4565b611186565b6102ac6111e9565b6102ac6111f2565b6102846104413660046129b4565b60026020526000908152604090205460ff1681565b6102ac6104643660046129b4565b611256565b6102ac6112b9565b6001546001600160a01b03165b6040516001600160a01b039091168152602001610290565b6102ac6104a4366004612851565b611320565b6102ac6104b73660046129cf565b61139f565b6104fe6104ca366004612851565b60086020526000908152604090208054600182015460028301546003909301546001600160a01b0392831693919092169184565b604080516001600160a01b039586168152949093166020850152918301526060820152608001610290565b6102f4600a5481565b6102f46105403660046129b4565b60066020526000908152604090205481565b61047e7f000000000000000000000000000000000000000000000000000000000000000081565b6102ac611765565b61028461058f3660046129b4565b60036020526000908152604090205460ff1681565b6102846105b23660046129b4565b6001600160a01b031660009081526003602052604090205460ff1690565b6102ac6105de3660046129b4565b61176e565b6102ac6105f13660046129b4565b6117d1565b6102f46106043660046129b4565b60076020526000908152604090205481565b3360009081526003602052604090205460ff166106735760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b60448201526064015b60405180910390fd5b8281146106b45760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260640161066a565b60005b838110156107b2578282828181106106d1576106d1612abe565b90506020020135600960008787858181106106ee576106ee612abe565b905060200201602081019061070391906129b4565b6001600160a01b031681526020810191909152604001600020557fceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce85858381811061075057610750612abe565b905060200201602081019061076591906129b4565b84848481811061077757610777612abe565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a1806107aa81612aea565b9150506106b7565b5050505050565b6002600054141561080c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161066a565b6002600055600154600160a01b900460ff161561085e5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161066a565b6001600160a01b0385166000908152600c602052604090205484116108c55760405162461bcd60e51b815260206004820152601060248201527f616d6f756e7420746f6f20736d616c6c00000000000000000000000000000000604482015260640161066a565b6001600160a01b0385166000908152600d6020526040902054158061090257506001600160a01b0385166000908152600d60205260409020548411155b61094e5760405162461bcd60e51b815260206004820152601060248201527f616d6f756e7420746f6f206c6172676500000000000000000000000000000000604482015260640161066a565b6040516bffffffffffffffffffffffff1933606090811b8216602084015287811b821660348401526048830187905277ffffffffffffffffffffffffffffffffffffffffffffffff1960c087811b821660688601529186901b909216607084015283811b8216608484015246901b16608c82015260009060940160408051601f1981840301815291815281516020928301206000818152600b90935291205490915060ff1615610a305760405162461bcd60e51b815260206004820152600d60248201526c7265636f72642065786973747360981b604482015260640161066a565b6000818152600b60205260409020805460ff19166001179055610a5e6001600160a01b0387163330886118b0565b604080518281523360208201526001600160a01b03888116828401526060820188905267ffffffffffffffff87166080830152851660a082015290517f15d2eeefbe4963b5b2178f239ddcc730dda55f1c23c22efb79ded0eb854ac7899181900360c00190a15050600160005550505050565b3360009081526003602052604090205460ff16610b295760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b828114610b6a5760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260640161066a565b60005b838110156107b257828282818110610b8757610b87612abe565b90506020020135600d6000878785818110610ba457610ba4612abe565b9050602002016020810190610bb991906129b4565b6001600160a01b031681526020810191909152604001600020557f0e5d348f9737ccc8b4cf0eea0ccf3670af071af8bea5d64664f10e700c08de72858583818110610c0657610c06612abe565b9050602002016020810190610c1b91906129b4565b848484818110610c2d57610c2d612abe565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180610c6081612aea565b915050610b6d565b6001546001600160a01b03163314610cc25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b610ccb8161194e565b50565b3360009081526003602052604090205460ff16610d265760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b600a8190556040518181527fc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6906020015b60405180910390a150565b3360009081526002602052604090205460ff16610dc15760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f7420706175736572000000000000000000000000604482015260640161066a565b610dc9611a0b565b565b3360009081526003602052604090205460ff16610e235760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b828114610e645760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260640161066a565b60005b838110156107b257828282818110610e8157610e81612abe565b9050602002013560066000878785818110610e9e57610e9e612abe565b9050602002016020810190610eb391906129b4565b6001600160a01b031681526020810191909152604001600020557f608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89858583818110610f0057610f00612abe565b9050602002016020810190610f1591906129b4565b848484818110610f2757610f27612abe565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180610f5a81612aea565b915050610e67565b3360009081526003602052604090205460ff16610fba5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b60048190556040518181527f2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b390602001610d57565b3360009081526003602052604090205460ff166110475760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b8281146110885760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260640161066a565b60005b838110156107b2578282828181106110a5576110a5612abe565b90506020020135600c60008787858181106110c2576110c2612abe565b90506020020160208101906110d791906129b4565b6001600160a01b031681526020810191909152604001600020557f0f48d517989455cd80ed52427e80553e66f9b69fd5cee8e26bd1a1f9c364fba685858381811061112457611124612abe565b905060200201602081019061113991906129b4565b84848481811061114b5761114b612abe565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a18061117e81612aea565b91505061108b565b6001546001600160a01b031633146111e05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b610ccb81611ab1565b610dc933611ab1565b6001546001600160a01b0316331461124c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b610dc96000611b6a565b6001546001600160a01b031633146112b05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b610ccb81611bc9565b3360009081526002602052604090205460ff166113185760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f7420706175736572000000000000000000000000604482015260640161066a565b610dc9611c86565b600154600160a01b900460ff161561136d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161066a565b600061137882611d0e565b80516040820151602083015192935061139b926001600160a01b03169190611ee0565b5050565b600154600160a01b900460ff16156113ec5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161066a565b6000463060405160200161144292919091825260601b6bffffffffffffffffffffffff191660208201527f57697468647261770000000000000000000000000000000000000000000000006034820152603c0190565b6040516020818303038152906040528051906020012090507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663682dbc22828b8b60405160200161149e93929190612b05565b6040516020818303038152906040528989898989896040518863ffffffff1660e01b81526004016114d59796959493929190612c36565b60006040518083038186803b1580156114ed57600080fd5b505afa158015611501573d6000803e3d6000fd5b5050505060006115468a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611f1592505050565b6020818101518251604080850151606080870151608088015160a0890151855197841b6bffffffffffffffffffffffff19908116898b015296841b871660348901526048880194909452911b909316606885015260c09290921b77ffffffffffffffffffffffffffffffffffffffffffffffff1916607c8401526084808401929092528051808403909201825260a490920182528051908301206000818152600b9093529120549192509060ff16156116315760405162461bcd60e51b815260206004820152600d60248201526c7265636f72642065786973747360981b604482015260640161066a565b6000818152600b602052604090819020805460ff1916600117905582519083015161165c9190612077565b81516001600160a01b031660009081526009602052604090205480158015906116885750808360400151115b156116aa576116a58284602001518560000151866040015161218f565b6116cb565b6020830151604084015184516116cb926001600160a01b0390911691611ee0565b602080840151845160408087015160808089015160a0808b01516060808d015187518d81526001600160a01b039a8b169b81019b909b52978916968a01969096529488019390935267ffffffffffffffff16908601528401521660c08201527f296a629c5265cb4e5319803d016902eb70a9079b89655fe2b7737821ed88beeb9060e00160405180910390a1505050505050505050505050565b610dc9336122af565b6001546001600160a01b031633146117c85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b610ccb816122af565b6001546001600160a01b0316331461182b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b6001600160a01b0381166118a75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161066a565b610ccb81611b6a565b6040516001600160a01b03808516602483015283166044820152606481018290526119489085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152612368565b50505050565b6001600160a01b03811660009081526003602052604090205460ff16156119b75760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f720000000000604482015260640161066a565b6001600160a01b038116600081815260036020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b59101610d57565b600154600160a01b900460ff16611a645760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161066a565b6001805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b03811660009081526002602052604090205460ff16611b195760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f74207061757365720000000000000000000000604482015260640161066a565b6001600160a01b038116600081815260026020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e9101610d57565b600180546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03811660009081526002602052604090205460ff1615611c325760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640161066a565b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101610d57565b600154600160a01b900460ff1615611cd35760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161066a565b6001805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611a943390565b604080516080810182526000808252602082018190529181018290526060810191909152600082815260086020908152604091829020825160808101845281546001600160a01b0390811682526001830154169281019290925260028101549282019290925260039091015460608201819052611dcd5760405162461bcd60e51b815260206004820152601a60248201527f64656c61796564207472616e73666572206e6f74206578697374000000000000604482015260640161066a565b600a548160600151611ddf9190612d16565b4211611e2d5760405162461bcd60e51b815260206004820152601d60248201527f64656c61796564207472616e73666572207374696c6c206c6f636b6564000000604482015260640161066a565b6000838152600860209081526040808320805473ffffffffffffffffffffffffffffffffffffffff199081168255600182018054909116905560028101849055600301929092558251908301518383015192517f3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d442693611ed293889390929091909384526001600160a01b03928316602085015291166040830152606082015260800190565b60405180910390a192915050565b6040516001600160a01b038316602482015260448101829052611f1090849063a9059cbb60e01b906064016118e4565b505050565b6040805160c08101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905283518085019094528184528301849052909190805b6020830151518351101561206f57611f738361244d565b90925090508160011415611fa257611f92611f8d84612487565b612544565b6001600160a01b03168452611f5c565b8160021415611fca57611fb7611f8d84612487565b6001600160a01b03166020850152611f5c565b8160031415611fee57611fe4611fdf84612487565b612555565b6040850152611f5c565b816004141561201657612003611f8d84612487565b6001600160a01b03166060850152611f5c565b816005141561203c576120288361258c565b67ffffffffffffffff166080850152611f5c565b81600614156120605761205661205184612487565b61260e565b60a0850152611f5c565b61206a8382612626565b611f5c565b505050919050565b600454612082575050565b6001600160a01b038216600090815260066020526040902054806120a557505050565b6001600160a01b038316600090815260056020526040812054600454909142916120cf8184612d2e565b6120d99190612d50565b6001600160a01b03871660009081526007602052604090205490915081111561210457849250612111565b61210e8584612d16565b92505b838311156121615760405162461bcd60e51b815260206004820152601260248201527f766f6c756d652065786365656473206361700000000000000000000000000000604482015260640161066a565b506001600160a01b039094166000908152600560209081526040808320939093556007905220929092555050565b600084815260086020526040902060030154156121ee5760405162461bcd60e51b815260206004820152601f60248201527f64656c61796564207472616e7366657220616c72656164792065786973747300604482015260640161066a565b604080516080810182526001600160a01b0380861682528481166020808401918252838501868152426060860190815260008b81526008909352918690209451855490851673ffffffffffffffffffffffffffffffffffffffff1991821617865592516001860180549190951693169290921790925551600283015551600390910155517fcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6906122a19086815260200190565b60405180910390a150505050565b6001600160a01b03811660009081526003602052604090205460ff166123175760405162461bcd60e51b815260206004820152601760248201527f4163636f756e74206973206e6f7420676f7665726e6f72000000000000000000604482015260640161066a565b6001600160a01b038116600081815260036020908152604091829020805460ff1916905590519182527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b9101610d57565b60006123bd826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166126989092919063ffffffff16565b805190915015611f1057808060200190518101906123db9190612d6f565b611f105760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161066a565b600080600061245b8461258c565b9050612468600882612d2e565b925080600716600581111561247f5761247f612d91565b915050915091565b606060006124948361258c565b905060008184600001516124a89190612d16565b90508360200151518111156124bc57600080fd5b8167ffffffffffffffff8111156124d5576124d5612da7565b6040519080825280601f01601f1916602001820160405280156124ff576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015612539578181015183820152612532602082612d16565b9050612517565b505050935250919050565b600061254f826126b1565b92915050565b600060208251111561256657600080fd5b602082015190508151602061257b9190612dbd565b612586906008612d50565b1c919050565b602080820151825181019091015160009182805b600a8110156126085783811a91506125b9816007612d50565b82607f16901b8517945081608016600014156125f6576125da816001612d16565b865187906125e9908390612d16565b9052509395945050505050565b8061260081612aea565b9150506125a0565b50600080fd5b6000815160201461261e57600080fd5b506020015190565b600081600581111561263a5761263a612d91565b141561264957611f108261258c565b600281600581111561265d5761265d612d91565b141561025c57600061266e8361258c565b905080836000018181516126829190612d16565b90525060208301515183511115611f1057600080fd5b60606126a784846000856126d9565b90505b9392505050565b600081516014146126c157600080fd5b50602001516c01000000000000000000000000900490565b6060824710156127515760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161066a565b843b61279f5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161066a565b600080866001600160a01b031685876040516127bb9190612dd4565b60006040518083038185875af1925050503d80600081146127f8576040519150601f19603f3d011682016040523d82523d6000602084013e6127fd565b606091505b509150915061280d828286612818565b979650505050505050565b606083156128275750816126aa565b8251156128375782518084602001fd5b8160405162461bcd60e51b815260040161066a9190612df0565b60006020828403121561286357600080fd5b5035919050565b60008083601f84011261287c57600080fd5b50813567ffffffffffffffff81111561289457600080fd5b6020830191508360208260051b85010111156128af57600080fd5b9250929050565b600080600080604085870312156128cc57600080fd5b843567ffffffffffffffff808211156128e457600080fd5b6128f08883890161286a565b9096509450602087013591508082111561290957600080fd5b506129168782880161286a565b95989497509550505050565b80356001600160a01b038116811461293957600080fd5b919050565b803567ffffffffffffffff8116811461293957600080fd5b600080600080600060a0868803121561296e57600080fd5b61297786612922565b94506020860135935061298c6040870161293e565b925061299a60608701612922565b91506129a86080870161293e565b90509295509295909350565b6000602082840312156129c657600080fd5b6126aa82612922565b6000806000806000806000806080898b0312156129eb57600080fd5b883567ffffffffffffffff80821115612a0357600080fd5b818b0191508b601f830112612a1757600080fd5b813581811115612a2657600080fd5b8c6020828501011115612a3857600080fd5b60209283019a509850908a01359080821115612a5357600080fd5b612a5f8c838d0161286a565b909850965060408b0135915080821115612a7857600080fd5b612a848c838d0161286a565b909650945060608b0135915080821115612a9d57600080fd5b50612aaa8b828c0161286a565b999c989b5096995094979396929594505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415612afe57612afe612ad4565b5060010190565b838152818360208301376000910160200190815292915050565b60005b83811015612b3a578181015183820152602001612b22565b838111156119485750506000910152565b60008151808452612b63816020860160208601612b1f565b601f01601f19169290920160200192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8183526000602080850194508260005b85811015612bdc576001600160a01b03612bc983612922565b1687529582019590820190600101612bb0565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115612c1957600080fd5b8260051b8083602087013760009401602001938452509192915050565b608081526000612c49608083018a612b4b565b82810360208401528088825260208201905060208960051b8301018a60005b8b811015612cdc57848303601f190184528135368e9003601e19018112612c8e57600080fd5b8d01803567ffffffffffffffff811115612ca757600080fd5b8036038f1315612cb657600080fd5b612cc4858260208501612b77565b60209687019690955093909301925050600101612c68565b50508481036040860152612cf181898b612ba0565b925050508281036060840152612d08818587612be7565b9a9950505050505050505050565b60008219821115612d2957612d29612ad4565b500190565b600082612d4b57634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615612d6a57612d6a612ad4565b500290565b600060208284031215612d8157600080fd5b815180151581146126aa57600080fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600082821015612dcf57612dcf612ad4565b500390565b60008251612de6818460208701612b1f565b9190910192915050565b6020815260006126aa6020830184612b4b56fea2646970667358221220a55f5a8d176c0aefd7dbc4e822152cae82e678824d20c4989acd83c3c946c1dd64736f6c63430008090033",
}

// OriginalTokenVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use OriginalTokenVaultMetaData.ABI instead.
var OriginalTokenVaultABI = OriginalTokenVaultMetaData.ABI

// OriginalTokenVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OriginalTokenVaultMetaData.Bin instead.
var OriginalTokenVaultBin = OriginalTokenVaultMetaData.Bin

// DeployOriginalTokenVault deploys a new Ethereum contract, binding an instance of OriginalTokenVault to it.
func DeployOriginalTokenVault(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address) (common.Address, *types.Transaction, *OriginalTokenVault, error) {
	parsed, err := OriginalTokenVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OriginalTokenVaultBin), backend, _sigsVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OriginalTokenVault{OriginalTokenVaultCaller: OriginalTokenVaultCaller{contract: contract}, OriginalTokenVaultTransactor: OriginalTokenVaultTransactor{contract: contract}, OriginalTokenVaultFilterer: OriginalTokenVaultFilterer{contract: contract}}, nil
}

// OriginalTokenVault is an auto generated Go binding around an Ethereum contract.
type OriginalTokenVault struct {
	OriginalTokenVaultCaller     // Read-only binding to the contract
	OriginalTokenVaultTransactor // Write-only binding to the contract
	OriginalTokenVaultFilterer   // Log filterer for contract events
}

// OriginalTokenVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type OriginalTokenVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginalTokenVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OriginalTokenVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginalTokenVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OriginalTokenVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginalTokenVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OriginalTokenVaultSession struct {
	Contract     *OriginalTokenVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OriginalTokenVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OriginalTokenVaultCallerSession struct {
	Contract *OriginalTokenVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OriginalTokenVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OriginalTokenVaultTransactorSession struct {
	Contract     *OriginalTokenVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OriginalTokenVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type OriginalTokenVaultRaw struct {
	Contract *OriginalTokenVault // Generic contract binding to access the raw methods on
}

// OriginalTokenVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OriginalTokenVaultCallerRaw struct {
	Contract *OriginalTokenVaultCaller // Generic read-only contract binding to access the raw methods on
}

// OriginalTokenVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OriginalTokenVaultTransactorRaw struct {
	Contract *OriginalTokenVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOriginalTokenVault creates a new instance of OriginalTokenVault, bound to a specific deployed contract.
func NewOriginalTokenVault(address common.Address, backend bind.ContractBackend) (*OriginalTokenVault, error) {
	contract, err := bindOriginalTokenVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVault{OriginalTokenVaultCaller: OriginalTokenVaultCaller{contract: contract}, OriginalTokenVaultTransactor: OriginalTokenVaultTransactor{contract: contract}, OriginalTokenVaultFilterer: OriginalTokenVaultFilterer{contract: contract}}, nil
}

// NewOriginalTokenVaultCaller creates a new read-only instance of OriginalTokenVault, bound to a specific deployed contract.
func NewOriginalTokenVaultCaller(address common.Address, caller bind.ContractCaller) (*OriginalTokenVaultCaller, error) {
	contract, err := bindOriginalTokenVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultCaller{contract: contract}, nil
}

// NewOriginalTokenVaultTransactor creates a new write-only instance of OriginalTokenVault, bound to a specific deployed contract.
func NewOriginalTokenVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*OriginalTokenVaultTransactor, error) {
	contract, err := bindOriginalTokenVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultTransactor{contract: contract}, nil
}

// NewOriginalTokenVaultFilterer creates a new log filterer instance of OriginalTokenVault, bound to a specific deployed contract.
func NewOriginalTokenVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*OriginalTokenVaultFilterer, error) {
	contract, err := bindOriginalTokenVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultFilterer{contract: contract}, nil
}

// bindOriginalTokenVault binds a generic wrapper to an already deployed contract.
func bindOriginalTokenVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OriginalTokenVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OriginalTokenVault *OriginalTokenVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OriginalTokenVault.Contract.OriginalTokenVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OriginalTokenVault *OriginalTokenVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.OriginalTokenVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OriginalTokenVault *OriginalTokenVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.OriginalTokenVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OriginalTokenVault *OriginalTokenVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OriginalTokenVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OriginalTokenVault *OriginalTokenVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OriginalTokenVault *OriginalTokenVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.contract.Transact(opts, method, params...)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCaller) DelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "delayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultSession) DelayPeriod() (*big.Int, error) {
	return _OriginalTokenVault.Contract.DelayPeriod(&_OriginalTokenVault.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) DelayPeriod() (*big.Int, error) {
	return _OriginalTokenVault.Contract.DelayPeriod(&_OriginalTokenVault.CallOpts)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCaller) DelayThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "delayThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.DelayThresholds(&_OriginalTokenVault.CallOpts, arg0)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.DelayThresholds(&_OriginalTokenVault.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_OriginalTokenVault *OriginalTokenVaultCaller) DelayedTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "delayedTransfers", arg0)

	outstruct := new(struct {
		Receiver  common.Address
		Token     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_OriginalTokenVault *OriginalTokenVaultSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _OriginalTokenVault.Contract.DelayedTransfers(&_OriginalTokenVault.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _OriginalTokenVault.Contract.DelayedTransfers(&_OriginalTokenVault.CallOpts, arg0)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultSession) EpochLength() (*big.Int, error) {
	return _OriginalTokenVault.Contract.EpochLength(&_OriginalTokenVault.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) EpochLength() (*big.Int, error) {
	return _OriginalTokenVault.Contract.EpochLength(&_OriginalTokenVault.CallOpts)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCaller) EpochVolumeCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "epochVolumeCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.EpochVolumeCaps(&_OriginalTokenVault.CallOpts, arg0)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.EpochVolumeCaps(&_OriginalTokenVault.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCaller) EpochVolumes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "epochVolumes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.EpochVolumes(&_OriginalTokenVault.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.EpochVolumes(&_OriginalTokenVault.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultSession) Governors(arg0 common.Address) (bool, error) {
	return _OriginalTokenVault.Contract.Governors(&_OriginalTokenVault.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _OriginalTokenVault.Contract.Governors(&_OriginalTokenVault.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultSession) IsGovernor(_account common.Address) (bool, error) {
	return _OriginalTokenVault.Contract.IsGovernor(&_OriginalTokenVault.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _OriginalTokenVault.Contract.IsGovernor(&_OriginalTokenVault.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultSession) IsPauser(account common.Address) (bool, error) {
	return _OriginalTokenVault.Contract.IsPauser(&_OriginalTokenVault.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) IsPauser(account common.Address) (bool, error) {
	return _OriginalTokenVault.Contract.IsPauser(&_OriginalTokenVault.CallOpts, account)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCaller) LastOpTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "lastOpTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.LastOpTimestamps(&_OriginalTokenVault.CallOpts, arg0)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.LastOpTimestamps(&_OriginalTokenVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.MaxDeposit(&_OriginalTokenVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.MaxDeposit(&_OriginalTokenVault.CallOpts, arg0)
}

// MinDeposit is a free data retrieval call binding the contract method 0x3c29f839.
//
// Solidity: function minDeposit(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCaller) MinDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "minDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x3c29f839.
//
// Solidity: function minDeposit(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultSession) MinDeposit(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.MinDeposit(&_OriginalTokenVault.CallOpts, arg0)
}

// MinDeposit is a free data retrieval call binding the contract method 0x3c29f839.
//
// Solidity: function minDeposit(address ) view returns(uint256)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) MinDeposit(arg0 common.Address) (*big.Int, error) {
	return _OriginalTokenVault.Contract.MinDeposit(&_OriginalTokenVault.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OriginalTokenVault *OriginalTokenVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OriginalTokenVault *OriginalTokenVaultSession) Owner() (common.Address, error) {
	return _OriginalTokenVault.Contract.Owner(&_OriginalTokenVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) Owner() (common.Address, error) {
	return _OriginalTokenVault.Contract.Owner(&_OriginalTokenVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultSession) Paused() (bool, error) {
	return _OriginalTokenVault.Contract.Paused(&_OriginalTokenVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) Paused() (bool, error) {
	return _OriginalTokenVault.Contract.Paused(&_OriginalTokenVault.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultSession) Pausers(arg0 common.Address) (bool, error) {
	return _OriginalTokenVault.Contract.Pausers(&_OriginalTokenVault.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _OriginalTokenVault.Contract.Pausers(&_OriginalTokenVault.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCaller) Records(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "records", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultSession) Records(arg0 [32]byte) (bool, error) {
	return _OriginalTokenVault.Contract.Records(&_OriginalTokenVault.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) Records(arg0 [32]byte) (bool, error) {
	return _OriginalTokenVault.Contract.Records(&_OriginalTokenVault.CallOpts, arg0)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_OriginalTokenVault *OriginalTokenVaultCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OriginalTokenVault.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_OriginalTokenVault *OriginalTokenVaultSession) SigsVerifier() (common.Address, error) {
	return _OriginalTokenVault.Contract.SigsVerifier(&_OriginalTokenVault.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_OriginalTokenVault *OriginalTokenVaultCallerSession) SigsVerifier() (common.Address, error) {
	return _OriginalTokenVault.Contract.SigsVerifier(&_OriginalTokenVault.CallOpts)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.AddGovernor(&_OriginalTokenVault.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.AddGovernor(&_OriginalTokenVault.TransactOpts, _account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.AddPauser(&_OriginalTokenVault.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.AddPauser(&_OriginalTokenVault.TransactOpts, account)
}

// Deposit is a paid mutator transaction binding the contract method 0x23463624.
//
// Solidity: function deposit(address _token, uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "deposit", _token, _amount, _mintChainId, _mintAccount, _nonce)
}

// Deposit is a paid mutator transaction binding the contract method 0x23463624.
//
// Solidity: function deposit(address _token, uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) Deposit(_token common.Address, _amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.Deposit(&_OriginalTokenVault.TransactOpts, _token, _amount, _mintChainId, _mintAccount, _nonce)
}

// Deposit is a paid mutator transaction binding the contract method 0x23463624.
//
// Solidity: function deposit(address _token, uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) Deposit(_token common.Address, _amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.Deposit(&_OriginalTokenVault.TransactOpts, _token, _amount, _mintChainId, _mintAccount, _nonce)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) ExecuteDelayedTransfer(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "executeDelayedTransfer", id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.ExecuteDelayedTransfer(&_OriginalTokenVault.TransactOpts, id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.ExecuteDelayedTransfer(&_OriginalTokenVault.TransactOpts, id)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) Pause() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.Pause(&_OriginalTokenVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.Pause(&_OriginalTokenVault.TransactOpts)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RemoveGovernor(&_OriginalTokenVault.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RemoveGovernor(&_OriginalTokenVault.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RemovePauser(&_OriginalTokenVault.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RemovePauser(&_OriginalTokenVault.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) RenounceGovernor() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RenounceGovernor(&_OriginalTokenVault.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RenounceGovernor(&_OriginalTokenVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RenounceOwnership(&_OriginalTokenVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RenounceOwnership(&_OriginalTokenVault.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) RenouncePauser() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RenouncePauser(&_OriginalTokenVault.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.RenouncePauser(&_OriginalTokenVault.TransactOpts)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) SetDelayPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "setDelayPeriod", _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetDelayPeriod(&_OriginalTokenVault.TransactOpts, _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetDelayPeriod(&_OriginalTokenVault.TransactOpts, _period)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) SetDelayThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "setDelayThresholds", _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetDelayThresholds(&_OriginalTokenVault.TransactOpts, _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetDelayThresholds(&_OriginalTokenVault.TransactOpts, _tokens, _thresholds)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) SetEpochLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "setEpochLength", _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetEpochLength(&_OriginalTokenVault.TransactOpts, _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetEpochLength(&_OriginalTokenVault.TransactOpts, _length)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) SetEpochVolumeCaps(opts *bind.TransactOpts, _tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "setEpochVolumeCaps", _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetEpochVolumeCaps(&_OriginalTokenVault.TransactOpts, _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetEpochVolumeCaps(&_OriginalTokenVault.TransactOpts, _tokens, _caps)
}

// SetMaxDeposit is a paid mutator transaction binding the contract method 0x303b6442.
//
// Solidity: function setMaxDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) SetMaxDeposit(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "setMaxDeposit", _tokens, _amounts)
}

// SetMaxDeposit is a paid mutator transaction binding the contract method 0x303b6442.
//
// Solidity: function setMaxDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) SetMaxDeposit(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetMaxDeposit(&_OriginalTokenVault.TransactOpts, _tokens, _amounts)
}

// SetMaxDeposit is a paid mutator transaction binding the contract method 0x303b6442.
//
// Solidity: function setMaxDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) SetMaxDeposit(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetMaxDeposit(&_OriginalTokenVault.TransactOpts, _tokens, _amounts)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x5ec2fa26.
//
// Solidity: function setMinDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) SetMinDeposit(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "setMinDeposit", _tokens, _amounts)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x5ec2fa26.
//
// Solidity: function setMinDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) SetMinDeposit(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetMinDeposit(&_OriginalTokenVault.TransactOpts, _tokens, _amounts)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x5ec2fa26.
//
// Solidity: function setMinDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) SetMinDeposit(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.SetMinDeposit(&_OriginalTokenVault.TransactOpts, _tokens, _amounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.TransferOwnership(&_OriginalTokenVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.TransferOwnership(&_OriginalTokenVault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) Unpause() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.Unpause(&_OriginalTokenVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.Unpause(&_OriginalTokenVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactor) Withdraw(opts *bind.TransactOpts, _request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.contract.Transact(opts, "withdraw", _request, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_OriginalTokenVault *OriginalTokenVaultSession) Withdraw(_request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.Withdraw(&_OriginalTokenVault.TransactOpts, _request, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_OriginalTokenVault *OriginalTokenVaultTransactorSession) Withdraw(_request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _OriginalTokenVault.Contract.Withdraw(&_OriginalTokenVault.TransactOpts, _request, _sigs, _signers, _powers)
}

// OriginalTokenVaultDelayPeriodUpdatedIterator is returned from FilterDelayPeriodUpdated and is used to iterate over the raw logs and unpacked data for DelayPeriodUpdated events raised by the OriginalTokenVault contract.
type OriginalTokenVaultDelayPeriodUpdatedIterator struct {
	Event *OriginalTokenVaultDelayPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultDelayPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultDelayPeriodUpdated)
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
		it.Event = new(OriginalTokenVaultDelayPeriodUpdated)
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
func (it *OriginalTokenVaultDelayPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultDelayPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultDelayPeriodUpdated represents a DelayPeriodUpdated event raised by the OriginalTokenVault contract.
type OriginalTokenVaultDelayPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelayPeriodUpdated is a free log retrieval operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterDelayPeriodUpdated(opts *bind.FilterOpts) (*OriginalTokenVaultDelayPeriodUpdatedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultDelayPeriodUpdatedIterator{contract: _OriginalTokenVault.contract, event: "DelayPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayPeriodUpdated is a free log subscription operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchDelayPeriodUpdated(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultDelayPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultDelayPeriodUpdated)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
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

// ParseDelayPeriodUpdated is a log parse operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseDelayPeriodUpdated(log types.Log) (*OriginalTokenVaultDelayPeriodUpdated, error) {
	event := new(OriginalTokenVaultDelayPeriodUpdated)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultDelayThresholdUpdatedIterator is returned from FilterDelayThresholdUpdated and is used to iterate over the raw logs and unpacked data for DelayThresholdUpdated events raised by the OriginalTokenVault contract.
type OriginalTokenVaultDelayThresholdUpdatedIterator struct {
	Event *OriginalTokenVaultDelayThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultDelayThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultDelayThresholdUpdated)
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
		it.Event = new(OriginalTokenVaultDelayThresholdUpdated)
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
func (it *OriginalTokenVaultDelayThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultDelayThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultDelayThresholdUpdated represents a DelayThresholdUpdated event raised by the OriginalTokenVault contract.
type OriginalTokenVaultDelayThresholdUpdated struct {
	Token     common.Address
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayThresholdUpdated is a free log retrieval operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterDelayThresholdUpdated(opts *bind.FilterOpts) (*OriginalTokenVaultDelayThresholdUpdatedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultDelayThresholdUpdatedIterator{contract: _OriginalTokenVault.contract, event: "DelayThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayThresholdUpdated is a free log subscription operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchDelayThresholdUpdated(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultDelayThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultDelayThresholdUpdated)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
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

// ParseDelayThresholdUpdated is a log parse operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseDelayThresholdUpdated(log types.Log) (*OriginalTokenVaultDelayThresholdUpdated, error) {
	event := new(OriginalTokenVaultDelayThresholdUpdated)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultDelayedTransferAddedIterator is returned from FilterDelayedTransferAdded and is used to iterate over the raw logs and unpacked data for DelayedTransferAdded events raised by the OriginalTokenVault contract.
type OriginalTokenVaultDelayedTransferAddedIterator struct {
	Event *OriginalTokenVaultDelayedTransferAdded // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultDelayedTransferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultDelayedTransferAdded)
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
		it.Event = new(OriginalTokenVaultDelayedTransferAdded)
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
func (it *OriginalTokenVaultDelayedTransferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultDelayedTransferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultDelayedTransferAdded represents a DelayedTransferAdded event raised by the OriginalTokenVault contract.
type OriginalTokenVaultDelayedTransferAdded struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferAdded is a free log retrieval operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterDelayedTransferAdded(opts *bind.FilterOpts) (*OriginalTokenVaultDelayedTransferAddedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultDelayedTransferAddedIterator{contract: _OriginalTokenVault.contract, event: "DelayedTransferAdded", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferAdded is a free log subscription operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchDelayedTransferAdded(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultDelayedTransferAdded) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultDelayedTransferAdded)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
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

// ParseDelayedTransferAdded is a log parse operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseDelayedTransferAdded(log types.Log) (*OriginalTokenVaultDelayedTransferAdded, error) {
	event := new(OriginalTokenVaultDelayedTransferAdded)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultDelayedTransferExecutedIterator is returned from FilterDelayedTransferExecuted and is used to iterate over the raw logs and unpacked data for DelayedTransferExecuted events raised by the OriginalTokenVault contract.
type OriginalTokenVaultDelayedTransferExecutedIterator struct {
	Event *OriginalTokenVaultDelayedTransferExecuted // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultDelayedTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultDelayedTransferExecuted)
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
		it.Event = new(OriginalTokenVaultDelayedTransferExecuted)
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
func (it *OriginalTokenVaultDelayedTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultDelayedTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultDelayedTransferExecuted represents a DelayedTransferExecuted event raised by the OriginalTokenVault contract.
type OriginalTokenVaultDelayedTransferExecuted struct {
	Id       [32]byte
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferExecuted is a free log retrieval operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterDelayedTransferExecuted(opts *bind.FilterOpts) (*OriginalTokenVaultDelayedTransferExecutedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultDelayedTransferExecutedIterator{contract: _OriginalTokenVault.contract, event: "DelayedTransferExecuted", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferExecuted is a free log subscription operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchDelayedTransferExecuted(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultDelayedTransferExecuted) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultDelayedTransferExecuted)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
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

// ParseDelayedTransferExecuted is a log parse operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseDelayedTransferExecuted(log types.Log) (*OriginalTokenVaultDelayedTransferExecuted, error) {
	event := new(OriginalTokenVaultDelayedTransferExecuted)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the OriginalTokenVault contract.
type OriginalTokenVaultDepositedIterator struct {
	Event *OriginalTokenVaultDeposited // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultDeposited)
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
		it.Event = new(OriginalTokenVaultDeposited)
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
func (it *OriginalTokenVaultDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultDeposited represents a Deposited event raised by the OriginalTokenVault contract.
type OriginalTokenVaultDeposited struct {
	DepositId   [32]byte
	Depositor   common.Address
	Token       common.Address
	Amount      *big.Int
	MintChainId uint64
	MintAccount common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x15d2eeefbe4963b5b2178f239ddcc730dda55f1c23c22efb79ded0eb854ac789.
//
// Solidity: event Deposited(bytes32 depositId, address depositor, address token, uint256 amount, uint64 mintChainId, address mintAccount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterDeposited(opts *bind.FilterOpts) (*OriginalTokenVaultDepositedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultDepositedIterator{contract: _OriginalTokenVault.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x15d2eeefbe4963b5b2178f239ddcc730dda55f1c23c22efb79ded0eb854ac789.
//
// Solidity: event Deposited(bytes32 depositId, address depositor, address token, uint256 amount, uint64 mintChainId, address mintAccount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultDeposited) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultDeposited)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x15d2eeefbe4963b5b2178f239ddcc730dda55f1c23c22efb79ded0eb854ac789.
//
// Solidity: event Deposited(bytes32 depositId, address depositor, address token, uint256 amount, uint64 mintChainId, address mintAccount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseDeposited(log types.Log) (*OriginalTokenVaultDeposited, error) {
	event := new(OriginalTokenVaultDeposited)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultEpochLengthUpdatedIterator is returned from FilterEpochLengthUpdated and is used to iterate over the raw logs and unpacked data for EpochLengthUpdated events raised by the OriginalTokenVault contract.
type OriginalTokenVaultEpochLengthUpdatedIterator struct {
	Event *OriginalTokenVaultEpochLengthUpdated // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultEpochLengthUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultEpochLengthUpdated)
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
		it.Event = new(OriginalTokenVaultEpochLengthUpdated)
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
func (it *OriginalTokenVaultEpochLengthUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultEpochLengthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultEpochLengthUpdated represents a EpochLengthUpdated event raised by the OriginalTokenVault contract.
type OriginalTokenVaultEpochLengthUpdated struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEpochLengthUpdated is a free log retrieval operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterEpochLengthUpdated(opts *bind.FilterOpts) (*OriginalTokenVaultEpochLengthUpdatedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultEpochLengthUpdatedIterator{contract: _OriginalTokenVault.contract, event: "EpochLengthUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochLengthUpdated is a free log subscription operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchEpochLengthUpdated(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultEpochLengthUpdated) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultEpochLengthUpdated)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
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

// ParseEpochLengthUpdated is a log parse operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseEpochLengthUpdated(log types.Log) (*OriginalTokenVaultEpochLengthUpdated, error) {
	event := new(OriginalTokenVaultEpochLengthUpdated)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultEpochVolumeUpdatedIterator is returned from FilterEpochVolumeUpdated and is used to iterate over the raw logs and unpacked data for EpochVolumeUpdated events raised by the OriginalTokenVault contract.
type OriginalTokenVaultEpochVolumeUpdatedIterator struct {
	Event *OriginalTokenVaultEpochVolumeUpdated // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultEpochVolumeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultEpochVolumeUpdated)
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
		it.Event = new(OriginalTokenVaultEpochVolumeUpdated)
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
func (it *OriginalTokenVaultEpochVolumeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultEpochVolumeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultEpochVolumeUpdated represents a EpochVolumeUpdated event raised by the OriginalTokenVault contract.
type OriginalTokenVaultEpochVolumeUpdated struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEpochVolumeUpdated is a free log retrieval operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterEpochVolumeUpdated(opts *bind.FilterOpts) (*OriginalTokenVaultEpochVolumeUpdatedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultEpochVolumeUpdatedIterator{contract: _OriginalTokenVault.contract, event: "EpochVolumeUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochVolumeUpdated is a free log subscription operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchEpochVolumeUpdated(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultEpochVolumeUpdated) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultEpochVolumeUpdated)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
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

// ParseEpochVolumeUpdated is a log parse operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseEpochVolumeUpdated(log types.Log) (*OriginalTokenVaultEpochVolumeUpdated, error) {
	event := new(OriginalTokenVaultEpochVolumeUpdated)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the OriginalTokenVault contract.
type OriginalTokenVaultGovernorAddedIterator struct {
	Event *OriginalTokenVaultGovernorAdded // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultGovernorAdded)
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
		it.Event = new(OriginalTokenVaultGovernorAdded)
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
func (it *OriginalTokenVaultGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultGovernorAdded represents a GovernorAdded event raised by the OriginalTokenVault contract.
type OriginalTokenVaultGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*OriginalTokenVaultGovernorAddedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultGovernorAddedIterator{contract: _OriginalTokenVault.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultGovernorAdded)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseGovernorAdded(log types.Log) (*OriginalTokenVaultGovernorAdded, error) {
	event := new(OriginalTokenVaultGovernorAdded)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the OriginalTokenVault contract.
type OriginalTokenVaultGovernorRemovedIterator struct {
	Event *OriginalTokenVaultGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultGovernorRemoved)
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
		it.Event = new(OriginalTokenVaultGovernorRemoved)
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
func (it *OriginalTokenVaultGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultGovernorRemoved represents a GovernorRemoved event raised by the OriginalTokenVault contract.
type OriginalTokenVaultGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*OriginalTokenVaultGovernorRemovedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultGovernorRemovedIterator{contract: _OriginalTokenVault.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultGovernorRemoved)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseGovernorRemoved(log types.Log) (*OriginalTokenVaultGovernorRemoved, error) {
	event := new(OriginalTokenVaultGovernorRemoved)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultMaxDepositUpdatedIterator is returned from FilterMaxDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxDepositUpdated events raised by the OriginalTokenVault contract.
type OriginalTokenVaultMaxDepositUpdatedIterator struct {
	Event *OriginalTokenVaultMaxDepositUpdated // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultMaxDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultMaxDepositUpdated)
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
		it.Event = new(OriginalTokenVaultMaxDepositUpdated)
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
func (it *OriginalTokenVaultMaxDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultMaxDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultMaxDepositUpdated represents a MaxDepositUpdated event raised by the OriginalTokenVault contract.
type OriginalTokenVaultMaxDepositUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxDepositUpdated is a free log retrieval operation binding the contract event 0x0e5d348f9737ccc8b4cf0eea0ccf3670af071af8bea5d64664f10e700c08de72.
//
// Solidity: event MaxDepositUpdated(address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterMaxDepositUpdated(opts *bind.FilterOpts) (*OriginalTokenVaultMaxDepositUpdatedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "MaxDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultMaxDepositUpdatedIterator{contract: _OriginalTokenVault.contract, event: "MaxDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxDepositUpdated is a free log subscription operation binding the contract event 0x0e5d348f9737ccc8b4cf0eea0ccf3670af071af8bea5d64664f10e700c08de72.
//
// Solidity: event MaxDepositUpdated(address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchMaxDepositUpdated(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultMaxDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "MaxDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultMaxDepositUpdated)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "MaxDepositUpdated", log); err != nil {
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

// ParseMaxDepositUpdated is a log parse operation binding the contract event 0x0e5d348f9737ccc8b4cf0eea0ccf3670af071af8bea5d64664f10e700c08de72.
//
// Solidity: event MaxDepositUpdated(address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseMaxDepositUpdated(log types.Log) (*OriginalTokenVaultMaxDepositUpdated, error) {
	event := new(OriginalTokenVaultMaxDepositUpdated)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "MaxDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultMinDepositUpdatedIterator is returned from FilterMinDepositUpdated and is used to iterate over the raw logs and unpacked data for MinDepositUpdated events raised by the OriginalTokenVault contract.
type OriginalTokenVaultMinDepositUpdatedIterator struct {
	Event *OriginalTokenVaultMinDepositUpdated // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultMinDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultMinDepositUpdated)
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
		it.Event = new(OriginalTokenVaultMinDepositUpdated)
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
func (it *OriginalTokenVaultMinDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultMinDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultMinDepositUpdated represents a MinDepositUpdated event raised by the OriginalTokenVault contract.
type OriginalTokenVaultMinDepositUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinDepositUpdated is a free log retrieval operation binding the contract event 0x0f48d517989455cd80ed52427e80553e66f9b69fd5cee8e26bd1a1f9c364fba6.
//
// Solidity: event MinDepositUpdated(address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterMinDepositUpdated(opts *bind.FilterOpts) (*OriginalTokenVaultMinDepositUpdatedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "MinDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultMinDepositUpdatedIterator{contract: _OriginalTokenVault.contract, event: "MinDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMinDepositUpdated is a free log subscription operation binding the contract event 0x0f48d517989455cd80ed52427e80553e66f9b69fd5cee8e26bd1a1f9c364fba6.
//
// Solidity: event MinDepositUpdated(address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchMinDepositUpdated(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultMinDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "MinDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultMinDepositUpdated)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "MinDepositUpdated", log); err != nil {
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

// ParseMinDepositUpdated is a log parse operation binding the contract event 0x0f48d517989455cd80ed52427e80553e66f9b69fd5cee8e26bd1a1f9c364fba6.
//
// Solidity: event MinDepositUpdated(address token, uint256 amount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseMinDepositUpdated(log types.Log) (*OriginalTokenVaultMinDepositUpdated, error) {
	event := new(OriginalTokenVaultMinDepositUpdated)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "MinDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OriginalTokenVault contract.
type OriginalTokenVaultOwnershipTransferredIterator struct {
	Event *OriginalTokenVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultOwnershipTransferred)
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
		it.Event = new(OriginalTokenVaultOwnershipTransferred)
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
func (it *OriginalTokenVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultOwnershipTransferred represents a OwnershipTransferred event raised by the OriginalTokenVault contract.
type OriginalTokenVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OriginalTokenVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultOwnershipTransferredIterator{contract: _OriginalTokenVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultOwnershipTransferred)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseOwnershipTransferred(log types.Log) (*OriginalTokenVaultOwnershipTransferred, error) {
	event := new(OriginalTokenVaultOwnershipTransferred)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OriginalTokenVault contract.
type OriginalTokenVaultPausedIterator struct {
	Event *OriginalTokenVaultPaused // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultPaused)
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
		it.Event = new(OriginalTokenVaultPaused)
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
func (it *OriginalTokenVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultPaused represents a Paused event raised by the OriginalTokenVault contract.
type OriginalTokenVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*OriginalTokenVaultPausedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultPausedIterator{contract: _OriginalTokenVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultPaused) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultPaused)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParsePaused(log types.Log) (*OriginalTokenVaultPaused, error) {
	event := new(OriginalTokenVaultPaused)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the OriginalTokenVault contract.
type OriginalTokenVaultPauserAddedIterator struct {
	Event *OriginalTokenVaultPauserAdded // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultPauserAdded)
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
		it.Event = new(OriginalTokenVaultPauserAdded)
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
func (it *OriginalTokenVaultPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultPauserAdded represents a PauserAdded event raised by the OriginalTokenVault contract.
type OriginalTokenVaultPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*OriginalTokenVaultPauserAddedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultPauserAddedIterator{contract: _OriginalTokenVault.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultPauserAdded) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultPauserAdded)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParsePauserAdded(log types.Log) (*OriginalTokenVaultPauserAdded, error) {
	event := new(OriginalTokenVaultPauserAdded)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the OriginalTokenVault contract.
type OriginalTokenVaultPauserRemovedIterator struct {
	Event *OriginalTokenVaultPauserRemoved // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultPauserRemoved)
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
		it.Event = new(OriginalTokenVaultPauserRemoved)
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
func (it *OriginalTokenVaultPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultPauserRemoved represents a PauserRemoved event raised by the OriginalTokenVault contract.
type OriginalTokenVaultPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*OriginalTokenVaultPauserRemovedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultPauserRemovedIterator{contract: _OriginalTokenVault.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultPauserRemoved)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParsePauserRemoved(log types.Log) (*OriginalTokenVaultPauserRemoved, error) {
	event := new(OriginalTokenVaultPauserRemoved)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OriginalTokenVault contract.
type OriginalTokenVaultUnpausedIterator struct {
	Event *OriginalTokenVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultUnpaused)
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
		it.Event = new(OriginalTokenVaultUnpaused)
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
func (it *OriginalTokenVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultUnpaused represents a Unpaused event raised by the OriginalTokenVault contract.
type OriginalTokenVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OriginalTokenVaultUnpausedIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultUnpausedIterator{contract: _OriginalTokenVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultUnpaused)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseUnpaused(log types.Log) (*OriginalTokenVaultUnpaused, error) {
	event := new(OriginalTokenVaultUnpaused)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginalTokenVaultWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the OriginalTokenVault contract.
type OriginalTokenVaultWithdrawnIterator struct {
	Event *OriginalTokenVaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *OriginalTokenVaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginalTokenVaultWithdrawn)
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
		it.Event = new(OriginalTokenVaultWithdrawn)
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
func (it *OriginalTokenVaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginalTokenVaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginalTokenVaultWithdrawn represents a Withdrawn event raised by the OriginalTokenVault contract.
type OriginalTokenVaultWithdrawn struct {
	WithdrawId  [32]byte
	Receiver    common.Address
	Token       common.Address
	Amount      *big.Int
	RefChainId  uint64
	RefId       [32]byte
	BurnAccount common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x296a629c5265cb4e5319803d016902eb70a9079b89655fe2b7737821ed88beeb.
//
// Solidity: event Withdrawn(bytes32 withdrawId, address receiver, address token, uint256 amount, uint64 refChainId, bytes32 refId, address burnAccount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*OriginalTokenVaultWithdrawnIterator, error) {

	logs, sub, err := _OriginalTokenVault.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &OriginalTokenVaultWithdrawnIterator{contract: _OriginalTokenVault.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x296a629c5265cb4e5319803d016902eb70a9079b89655fe2b7737821ed88beeb.
//
// Solidity: event Withdrawn(bytes32 withdrawId, address receiver, address token, uint256 amount, uint64 refChainId, bytes32 refId, address burnAccount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *OriginalTokenVaultWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OriginalTokenVault.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginalTokenVaultWithdrawn)
				if err := _OriginalTokenVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x296a629c5265cb4e5319803d016902eb70a9079b89655fe2b7737821ed88beeb.
//
// Solidity: event Withdrawn(bytes32 withdrawId, address receiver, address token, uint256 amount, uint64 refChainId, bytes32 refId, address burnAccount)
func (_OriginalTokenVault *OriginalTokenVaultFilterer) ParseWithdrawn(log types.Log) (*OriginalTokenVaultWithdrawn, error) {
	event := new(OriginalTokenVaultWithdrawn)
	if err := _OriginalTokenVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenMetaData contains all meta data concerning the PeggedToken contract.
var PeggedTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001042380380620010428339810160408190526200003491620001f6565b8351849084906200004d90600390602085019062000083565b5080516200006390600490602084019062000083565b50505060ff90911660a0526001600160a01b031660805250620002d79050565b82805462000091906200029a565b90600052602060002090601f016020900481019282620000b5576000855562000100565b82601f10620000d057805160ff191683800117855562000100565b8280016001018555821562000100579182015b8281111562000100578251825591602001919060010190620000e3565b506200010e92915062000112565b5090565b5b808211156200010e576000815560010162000113565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200015157600080fd5b81516001600160401b03808211156200016e576200016e62000129565b604051601f8301601f19908116603f0116810190828211818310171562000199576200019962000129565b81604052838152602092508683858801011115620001b657600080fd5b600091505b83821015620001da5785820183015181830184015290820190620001bb565b83821115620001ec5760008385830101525b9695505050505050565b600080600080608085870312156200020d57600080fd5b84516001600160401b03808211156200022557600080fd5b62000233888389016200013f565b955060208701519150808211156200024a57600080fd5b5062000259878288016200013f565b935050604085015160ff811681146200027157600080fd5b60608601519092506001600160a01b03811681146200028f57600080fd5b939692955090935050565b600181811c90821680620002af57607f821691505b60208210811415620002d157634e487b7160e01b600052602260045260246000fd5b50919050565b60805160a051610d376200030b600039600061015c0152600081816102560152818161044301526104d80152610d376000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806370a082311161008c578063a457c2d711610066578063a457c2d7146101f2578063a9059cbb14610205578063dd62ed3e14610218578063f77c47911461025157600080fd5b806370a08231146101ae57806395d89b41146101d75780639dc29fac146101df57600080fd5b806323b872dd116100c857806323b872dd14610142578063313ce56714610155578063395093511461018657806340c10f191461019957600080fd5b806306fdde03146100ef578063095ea7b31461010d57806318160ddd14610130575b600080fd5b6100f7610290565b6040516101049190610b55565b60405180910390f35b61012061011b366004610bc6565b610322565b6040519015158152602001610104565b6002545b604051908152602001610104565b610120610150366004610bf0565b610338565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610104565b610120610194366004610bc6565b6103fc565b6101ac6101a7366004610bc6565b610438565b005b6101346101bc366004610c2c565b6001600160a01b031660009081526020819052604090205490565b6100f76104be565b6101ac6101ed366004610bc6565b6104cd565b610120610200366004610bc6565b61054f565b610120610213366004610bc6565b610600565b610134610226366004610c4e565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6102787f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610104565b60606003805461029f90610c81565b80601f01602080910402602001604051908101604052809291908181526020018280546102cb90610c81565b80156103185780601f106102ed57610100808354040283529160200191610318565b820191906000526020600020905b8154815290600101906020018083116102fb57829003601f168201915b5050505050905090565b600061032f33848461060d565b50600192915050565b6000610345848484610732565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156103e45760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206160448201527f6c6c6f77616e636500000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6103f1853385840361060d565b506001949350505050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161032f918590610433908690610cd2565b61060d565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104b05760405162461bcd60e51b815260206004820152601860248201527f63616c6c6572206973206e6f7420636f6e74726f6c6c6572000000000000000060448201526064016103db565b6104ba8282610930565b5050565b60606004805461029f90610c81565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105455760405162461bcd60e51b815260206004820152601860248201527f63616c6c6572206973206e6f7420636f6e74726f6c6c6572000000000000000060448201526064016103db565b6104ba8282610a0f565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156105e95760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016103db565b6105f6338585840361060d565b5060019392505050565b600061032f338484610732565b6001600160a01b03831661066f5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103db565b6001600160a01b0382166106d05760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103db565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b0383166107ae5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016103db565b6001600160a01b0382166108105760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103db565b6001600160a01b0383166000908152602081905260409020548181101561089f5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016103db565b6001600160a01b038085166000908152602081905260408082208585039055918516815290812080548492906108d6908490610cd2565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161092291815260200190565b60405180910390a350505050565b6001600160a01b0382166109865760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103db565b80600260008282546109989190610cd2565b90915550506001600160a01b038216600090815260208190526040812080548392906109c5908490610cd2565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6001600160a01b038216610a6f5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016103db565b6001600160a01b03821660009081526020819052604090205481811015610ae35760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016103db565b6001600160a01b0383166000908152602081905260408120838303905560028054849290610b12908490610cea565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610725565b600060208083528351808285015260005b81811015610b8257858101830151858201604001528201610b66565b81811115610b94576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b0381168114610bc157600080fd5b919050565b60008060408385031215610bd957600080fd5b610be283610baa565b946020939093013593505050565b600080600060608486031215610c0557600080fd5b610c0e84610baa565b9250610c1c60208501610baa565b9150604084013590509250925092565b600060208284031215610c3e57600080fd5b610c4782610baa565b9392505050565b60008060408385031215610c6157600080fd5b610c6a83610baa565b9150610c7860208401610baa565b90509250929050565b600181811c90821680610c9557607f821691505b60208210811415610cb657634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610ce557610ce5610cbc565b500190565b600082821015610cfc57610cfc610cbc565b50039056fea2646970667358221220bbcd28009d808ff7cc07cdb3c67bbae01993448435f11a618a4e2e509e2331ae64736f6c63430008090033",
}

// PeggedTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PeggedTokenMetaData.ABI instead.
var PeggedTokenABI = PeggedTokenMetaData.ABI

// PeggedTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PeggedTokenMetaData.Bin instead.
var PeggedTokenBin = PeggedTokenMetaData.Bin

// DeployPeggedToken deploys a new Ethereum contract, binding an instance of PeggedToken to it.
func DeployPeggedToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, _controller common.Address) (common.Address, *types.Transaction, *PeggedToken, error) {
	parsed, err := PeggedTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PeggedTokenBin), backend, name_, symbol_, decimals_, _controller)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PeggedToken{PeggedTokenCaller: PeggedTokenCaller{contract: contract}, PeggedTokenTransactor: PeggedTokenTransactor{contract: contract}, PeggedTokenFilterer: PeggedTokenFilterer{contract: contract}}, nil
}

// PeggedToken is an auto generated Go binding around an Ethereum contract.
type PeggedToken struct {
	PeggedTokenCaller     // Read-only binding to the contract
	PeggedTokenTransactor // Write-only binding to the contract
	PeggedTokenFilterer   // Log filterer for contract events
}

// PeggedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeggedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeggedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeggedTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeggedTokenSession struct {
	Contract     *PeggedToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeggedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeggedTokenCallerSession struct {
	Contract *PeggedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PeggedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeggedTokenTransactorSession struct {
	Contract     *PeggedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PeggedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeggedTokenRaw struct {
	Contract *PeggedToken // Generic contract binding to access the raw methods on
}

// PeggedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeggedTokenCallerRaw struct {
	Contract *PeggedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PeggedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeggedTokenTransactorRaw struct {
	Contract *PeggedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeggedToken creates a new instance of PeggedToken, bound to a specific deployed contract.
func NewPeggedToken(address common.Address, backend bind.ContractBackend) (*PeggedToken, error) {
	contract, err := bindPeggedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PeggedToken{PeggedTokenCaller: PeggedTokenCaller{contract: contract}, PeggedTokenTransactor: PeggedTokenTransactor{contract: contract}, PeggedTokenFilterer: PeggedTokenFilterer{contract: contract}}, nil
}

// NewPeggedTokenCaller creates a new read-only instance of PeggedToken, bound to a specific deployed contract.
func NewPeggedTokenCaller(address common.Address, caller bind.ContractCaller) (*PeggedTokenCaller, error) {
	contract, err := bindPeggedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenCaller{contract: contract}, nil
}

// NewPeggedTokenTransactor creates a new write-only instance of PeggedToken, bound to a specific deployed contract.
func NewPeggedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PeggedTokenTransactor, error) {
	contract, err := bindPeggedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenTransactor{contract: contract}, nil
}

// NewPeggedTokenFilterer creates a new log filterer instance of PeggedToken, bound to a specific deployed contract.
func NewPeggedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PeggedTokenFilterer, error) {
	contract, err := bindPeggedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenFilterer{contract: contract}, nil
}

// bindPeggedToken binds a generic wrapper to an already deployed contract.
func bindPeggedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeggedTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeggedToken *PeggedTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeggedToken.Contract.PeggedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeggedToken *PeggedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedToken.Contract.PeggedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeggedToken *PeggedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeggedToken.Contract.PeggedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeggedToken *PeggedTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeggedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeggedToken *PeggedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeggedToken *PeggedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeggedToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PeggedToken *PeggedTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PeggedToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PeggedToken *PeggedTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PeggedToken.Contract.Allowance(&_PeggedToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PeggedToken *PeggedTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PeggedToken.Contract.Allowance(&_PeggedToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PeggedToken *PeggedTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PeggedToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PeggedToken *PeggedTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PeggedToken.Contract.BalanceOf(&_PeggedToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PeggedToken *PeggedTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PeggedToken.Contract.BalanceOf(&_PeggedToken.CallOpts, account)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PeggedToken *PeggedTokenCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeggedToken.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PeggedToken *PeggedTokenSession) Controller() (common.Address, error) {
	return _PeggedToken.Contract.Controller(&_PeggedToken.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PeggedToken *PeggedTokenCallerSession) Controller() (common.Address, error) {
	return _PeggedToken.Contract.Controller(&_PeggedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PeggedToken *PeggedTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PeggedToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PeggedToken *PeggedTokenSession) Decimals() (uint8, error) {
	return _PeggedToken.Contract.Decimals(&_PeggedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PeggedToken *PeggedTokenCallerSession) Decimals() (uint8, error) {
	return _PeggedToken.Contract.Decimals(&_PeggedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PeggedToken *PeggedTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PeggedToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PeggedToken *PeggedTokenSession) Name() (string, error) {
	return _PeggedToken.Contract.Name(&_PeggedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PeggedToken *PeggedTokenCallerSession) Name() (string, error) {
	return _PeggedToken.Contract.Name(&_PeggedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PeggedToken *PeggedTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PeggedToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PeggedToken *PeggedTokenSession) Symbol() (string, error) {
	return _PeggedToken.Contract.Symbol(&_PeggedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PeggedToken *PeggedTokenCallerSession) Symbol() (string, error) {
	return _PeggedToken.Contract.Symbol(&_PeggedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PeggedToken *PeggedTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PeggedToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PeggedToken *PeggedTokenSession) TotalSupply() (*big.Int, error) {
	return _PeggedToken.Contract.TotalSupply(&_PeggedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PeggedToken *PeggedTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PeggedToken.Contract.TotalSupply(&_PeggedToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.Approve(&_PeggedToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.Approve(&_PeggedToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_PeggedToken *PeggedTokenTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_PeggedToken *PeggedTokenSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.Burn(&_PeggedToken.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_PeggedToken *PeggedTokenTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.Burn(&_PeggedToken.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PeggedToken *PeggedTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PeggedToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PeggedToken *PeggedTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.DecreaseAllowance(&_PeggedToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PeggedToken *PeggedTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.DecreaseAllowance(&_PeggedToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PeggedToken *PeggedTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PeggedToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PeggedToken *PeggedTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.IncreaseAllowance(&_PeggedToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PeggedToken *PeggedTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.IncreaseAllowance(&_PeggedToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_PeggedToken *PeggedTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_PeggedToken *PeggedTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.Mint(&_PeggedToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_PeggedToken *PeggedTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.Mint(&_PeggedToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.Transfer(&_PeggedToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.Transfer(&_PeggedToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.TransferFrom(&_PeggedToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PeggedToken *PeggedTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PeggedToken.Contract.TransferFrom(&_PeggedToken.TransactOpts, sender, recipient, amount)
}

// PeggedTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PeggedToken contract.
type PeggedTokenApprovalIterator struct {
	Event *PeggedTokenApproval // Event containing the contract specifics and raw log

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
func (it *PeggedTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenApproval)
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
		it.Event = new(PeggedTokenApproval)
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
func (it *PeggedTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenApproval represents a Approval event raised by the PeggedToken contract.
type PeggedTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PeggedToken *PeggedTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PeggedTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PeggedToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenApprovalIterator{contract: _PeggedToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PeggedToken *PeggedTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PeggedTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PeggedToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenApproval)
				if err := _PeggedToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PeggedToken *PeggedTokenFilterer) ParseApproval(log types.Log) (*PeggedTokenApproval, error) {
	event := new(PeggedTokenApproval)
	if err := _PeggedToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PeggedToken contract.
type PeggedTokenTransferIterator struct {
	Event *PeggedTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PeggedTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenTransfer)
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
		it.Event = new(PeggedTokenTransfer)
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
func (it *PeggedTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenTransfer represents a Transfer event raised by the PeggedToken contract.
type PeggedTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PeggedToken *PeggedTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PeggedTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PeggedToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenTransferIterator{contract: _PeggedToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PeggedToken *PeggedTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PeggedTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PeggedToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenTransfer)
				if err := _PeggedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PeggedToken *PeggedTokenFilterer) ParseTransfer(log types.Log) (*PeggedTokenTransfer, error) {
	event := new(PeggedTokenTransfer)
	if err := _PeggedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeMetaData contains all meta data concerning the PeggedTokenBridge contract.
var PeggedTokenBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"burnId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAccount\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DelayThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedTransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelayedTransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EpochLengthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EpochVolumeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MaxBurnUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinBurnUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mintId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"refChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_withdrawAccount\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delayThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumeCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"executeDelayedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastOpTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setDelayThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"name\":\"setEpochVolumeCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMaxBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002d8b38038062002d8b83398101604081905262000034916200024f565b6200003f3362000074565b6000805460ff60a01b191690556200005733620000c4565b62000062336200018f565b6001600160a01b031660805262000281565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff1615620001335760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064015b60405180910390fd5b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f891015b60405180910390a150565b6001600160a01b03811660009081526002602052604090205460ff1615620001fa5760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f72000000000060448201526064016200012a565b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5910162000184565b6000602082840312156200026257600080fd5b81516001600160a01b03811681146200027a57600080fd5b9392505050565b608051612ae7620002a46000396000818161051e01526114160152612ae76000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c806382dc1ec411610145578063de790c7e116100bd578063eecdac881161008c578063f832138311610071578063f8321383146105d0578063f8734302146105f0578063f9a8ea081461060357600080fd5b8063eecdac88146105aa578063f2fde38b146105bd57600080fd5b8063de790c7e14610540578063e026049c14610553578063e3eece261461055b578063e43581b81461057e57600080fd5b8063adc0d57f11610114578063b5f2bc47116100f9578063b5f2bc47146104e6578063bf4816f014610506578063ccf2683b1461051957600080fd5b8063adc0d57f14610470578063b1c94d94146104dd57600080fd5b806382dc1ec41461041d5780638456cb59146104305780638da5cb5b146104385780639e25fc5c1461045d57600080fd5b806354eea796116101d85780636b2c0f55116101a7578063715018a61161018c578063715018a6146103d25780637f856013146103da57806380f51c12146103fa57600080fd5b80636b2c0f55146103b75780636ef8d66d146103ca57600080fd5b806354eea7961461036957806357d775f81461037c5780635c975abb1461038557806360216b001461039757600080fd5b80633f4ba83a1161022f57806347b16c6c1161021457806347b16c6c14610308578063497bf3b21461031b57806352532faa1461034957600080fd5b80633f4ba83a146102d457806346fbf68e146102dc57600080fd5b806301e647251461026157806317bdbae5146102995780633c4a25d0146102ae5780633d572107146102c1575b600080fd5b61028461026f36600461256e565b600a6020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6102ac6102a73660046125d3565b610616565b005b6102ac6102bc36600461265b565b6107b9565b6102ac6102cf36600461256e565b61081f565b6102ac6108b3565b6102846102ea36600461265b565b6001600160a01b031660009081526001602052604090205460ff1690565b6102ac6103163660046125d3565b61091c565b61033b61032936600461265b565b600c6020526000908152604090205481565b604051908152602001610290565b61033b61035736600461265b565b60086020526000908152604090205481565b6102ac61037736600461256e565b610ab3565b61033b60035481565b600054600160a01b900460ff16610284565b61033b6103a536600461265b565b60046020526000908152604090205481565b6102ac6103c536600461265b565b610b40565b6102ac610ba3565b6102ac610bac565b61033b6103e836600461265b565b600b6020526000908152604090205481565b61028461040836600461265b565b60016020526000908152604090205460ff1681565b6102ac61042b36600461265b565b610c10565b6102ac610c73565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610290565b6102ac61046b36600461256e565b610cda565b6104b261047e36600461256e565b60076020526000908152604090208054600182015460028301546003909301546001600160a01b0392831693919092169184565b604080516001600160a01b039586168152949093166020850152918301526060820152608001610290565b61033b60095481565b61033b6104f436600461265b565b60056020526000908152604090205481565b6102ac6105143660046125d3565b610da7565b6104457f000000000000000000000000000000000000000000000000000000000000000081565b6102ac61054e36600461267d565b610f3e565b6102ac611227565b61028461056936600461265b565b60026020526000908152604090205460ff1681565b61028461058c36600461265b565b6001600160a01b031660009081526002602052604090205460ff1690565b6102ac6105b836600461265b565b611230565b6102ac6105cb36600461265b565b611293565b61033b6105de36600461265b565b60066020526000908152604090205481565b6102ac6105fe3660046126d9565b611372565b6102ac6106113660046125d3565b61177b565b3360009081526002602052604090205460ff166106735760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b60448201526064015b60405180910390fd5b8281146106b45760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260640161066a565b60005b838110156107b2578282828181106106d1576106d16127c8565b90506020020135600860008787858181106106ee576106ee6127c8565b9050602002016020810190610703919061265b565b6001600160a01b031681526020810191909152604001600020557fceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce858583818110610750576107506127c8565b9050602002016020810190610765919061265b565b848484818110610777576107776127c8565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a1806107aa816127f4565b9150506106b7565b5050505050565b6000546001600160a01b031633146108135760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b61081c81611912565b50565b3360009081526002602052604090205460ff166108775760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b60098190556040518181527fc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6906020015b60405180910390a150565b3360009081526001602052604090205460ff166109125760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f7420706175736572000000000000000000000000604482015260640161066a565b61091a6119cf565b565b3360009081526002602052604090205460ff166109745760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b8281146109b55760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260640161066a565b60005b838110156107b2578282828181106109d2576109d26127c8565b90506020020135600560008787858181106109ef576109ef6127c8565b9050602002016020810190610a04919061265b565b6001600160a01b031681526020810191909152604001600020557f608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89858583818110610a5157610a516127c8565b9050602002016020810190610a66919061265b565b848484818110610a7857610a786127c8565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180610aab816127f4565b9150506109b8565b3360009081526002602052604090205460ff16610b0b5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b60038190556040518181527f2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3906020016108a8565b6000546001600160a01b03163314610b9a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b61081c81611a75565b61091a33611a75565b6000546001600160a01b03163314610c065760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b61091a6000611b2e565b6000546001600160a01b03163314610c6a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b61081c81611b8b565b3360009081526001602052604090205460ff16610cd25760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f7420706175736572000000000000000000000000604482015260640161066a565b61091a611c49565b600054600160a01b900460ff1615610d275760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161066a565b6000610d3282611cd1565b6020810151815160408084015190516340c10f1960e01b81526001600160a01b039283166004820152602481019190915292935016906340c10f1990604401600060405180830381600087803b158015610d8b57600080fd5b505af1158015610d9f573d6000803e3d6000fd5b505050505050565b3360009081526002602052604090205460ff16610dff5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b828114610e405760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260640161066a565b60005b838110156107b257828282818110610e5d57610e5d6127c8565b90506020020135600b6000878785818110610e7a57610e7a6127c8565b9050602002016020810190610e8f919061265b565b6001600160a01b031681526020810191909152604001600020557f3796cd0b17a8734f8da819920625598e9a18be490f686725282e5383f1d06683858583818110610edc57610edc6127c8565b9050602002016020810190610ef1919061265b565b848484818110610f0357610f036127c8565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180610f36816127f4565b915050610e43565b600054600160a01b900460ff1615610f8b5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161066a565b6001600160a01b0384166000908152600b60205260409020548311610ff25760405162461bcd60e51b815260206004820152601060248201527f616d6f756e7420746f6f20736d616c6c00000000000000000000000000000000604482015260640161066a565b6001600160a01b0384166000908152600c6020526040902054158061102f57506001600160a01b0384166000908152600c60205260409020548311155b61107b5760405162461bcd60e51b815260206004820152601060248201527f616d6f756e7420746f6f206c6172676500000000000000000000000000000000604482015260640161066a565b6040516bffffffffffffffffffffffff1933606090811b8216602084015286811b821660348401526048830186905284901b16606882015277ffffffffffffffffffffffffffffffffffffffffffffffff1960c083811b8216607c84015246901b166084820152600090608c0160408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff16156111505760405162461bcd60e51b815260206004820152600d60248201526c7265636f72642065786973747360981b604482015260640161066a565b6000818152600a602052604090819020805460ff1916600117905551632770a7eb60e21b8152336004820152602481018590526001600160a01b03861690639dc29fac90604401600060405180830381600087803b1580156111b157600080fd5b505af11580156111c5573d6000803e3d6000fd5b5050604080518481526001600160a01b0389811660208301523382840152606082018990528716608082015290517f75f1bf55bb1de41b63a775dc7d4500f01114ee62b688a6b11d34f4692c1f3d4393509081900360a0019150a15050505050565b61091a33611ea3565b6000546001600160a01b0316331461128a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b61081c81611ea3565b6000546001600160a01b031633146112ed5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161066a565b6001600160a01b0381166113695760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161066a565b61081c81611b2e565b600054600160a01b900460ff16156113bf5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161066a565b600046306040516020016113fc92919091825260601b6bffffffffffffffffffffffff1916602082015263135a5b9d60e21b603482015260380190565b6040516020818303038152906040528051906020012090507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663682dbc22828b8b6040516020016114589392919061280f565b6040516020818303038152906040528989898989896040518863ffffffff1660e01b815260040161148f9796959493929190612979565b60006040518083038186803b1580156114a757600080fd5b505afa1580156114bb573d6000803e3d6000fd5b5050505060006115008a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611f5c92505050565b6020818101518251604080850151606080870151608088015160a0890151855197841b6bffffffffffffffffffffffff19908116898b015296841b871660348901526048880194909452911b909316606885015260c09290921b77ffffffffffffffffffffffffffffffffffffffffffffffff1916607c8401526084808401929092528051808403909201825260a490920182528051908301206000818152600a9093529120549192509060ff16156115eb5760405162461bcd60e51b815260206004820152600d60248201526c7265636f72642065786973747360981b604482015260640161066a565b6000818152600a602052604090819020805460ff1916600117905582519083015161161691906120be565b81516001600160a01b031660009081526008602052604090205480158015906116425750808360400151115b156116645761165f828460200151856000015186604001516121d6565b6116d4565b8251602084015160408086015190516340c10f1960e01b81526001600160a01b03928316600482015260248101919091529116906340c10f1990604401600060405180830381600087803b1580156116bb57600080fd5b505af11580156116cf573d6000803e3d6000fd5b505050505b7f5bc84ecccfced5bb04bfc7f3efcdbe7f5cd21949ef146811b4d1967fe41f777a8284600001518560200151866040015187608001518860a00151896060015160405161176597969594939291909687526001600160a01b0395861660208801529385166040870152606086019290925267ffffffffffffffff16608085015260a08401521660c082015260e00190565b60405180910390a1505050505050505050505050565b3360009081526002602052604090205460ff166117d35760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b604482015260640161066a565b8281146118145760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260640161066a565b60005b838110156107b257828282818110611831576118316127c8565b90506020020135600c600087878581811061184e5761184e6127c8565b9050602002016020810190611863919061265b565b6001600160a01b031681526020810191909152604001600020557fa3181379f6db47d9037efc6b6e8e3efe8c55ddb090b4f0512c152f97c4e47da58585838181106118b0576118b06127c8565b90506020020160208101906118c5919061265b565b8484848181106118d7576118d76127c8565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a18061190a816127f4565b915050611817565b6001600160a01b03811660009081526002602052604090205460ff161561197b5760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f720000000000604482015260640161066a565b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b591016108a8565b600054600160a01b900460ff16611a285760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161066a565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b03811660009081526001602052604090205460ff16611add5760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f74207061757365720000000000000000000000604482015260640161066a565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e91016108a8565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff1615611bf45760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640161066a565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f891016108a8565b600054600160a01b900460ff1615611c965760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161066a565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611a583390565b604080516080810182526000808252602082018190529181018290526060810191909152600082815260076020908152604091829020825160808101845281546001600160a01b0390811682526001830154169281019290925260028101549282019290925260039091015460608201819052611d905760405162461bcd60e51b815260206004820152601a60248201527f64656c61796564207472616e73666572206e6f74206578697374000000000000604482015260640161066a565b6009548160600151611da29190612a15565b4211611df05760405162461bcd60e51b815260206004820152601d60248201527f64656c61796564207472616e73666572207374696c6c206c6f636b6564000000604482015260640161066a565b6000838152600760209081526040808320805473ffffffffffffffffffffffffffffffffffffffff199081168255600182018054909116905560028101849055600301929092558251908301518383015192517f3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d442693611e9593889390929091909384526001600160a01b03928316602085015291166040830152606082015260800190565b60405180910390a192915050565b6001600160a01b03811660009081526002602052604090205460ff16611f0b5760405162461bcd60e51b815260206004820152601760248201527f4163636f756e74206973206e6f7420676f7665726e6f72000000000000000000604482015260640161066a565b6001600160a01b038116600081815260026020908152604091829020805460ff1916905590519182527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b91016108a8565b6040805160c08101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905283518085019094528184528301849052909190805b602083015151835110156120b657611fba836122f6565b90925090508160011415611fe957611fd9611fd484612330565b6123ed565b6001600160a01b03168452611fa3565b816002141561201157611ffe611fd484612330565b6001600160a01b03166020850152611fa3565b81600314156120355761202b61202684612330565b6123fe565b6040850152611fa3565b816004141561205d5761204a611fd484612330565b6001600160a01b03166060850152611fa3565b81600514156120835761206f83612435565b67ffffffffffffffff166080850152611fa3565b81600614156120a75761209d61209884612330565b6124b7565b60a0850152611fa3565b6120b183826124cf565b611fa3565b505050919050565b6003546120c9575050565b6001600160a01b038216600090815260056020526040902054806120ec57505050565b6001600160a01b038316600090815260046020526040812054600354909142916121168184612a2d565b6121209190612a4f565b6001600160a01b03871660009081526006602052604090205490915081111561214b57849250612158565b6121558584612a15565b92505b838311156121a85760405162461bcd60e51b815260206004820152601260248201527f766f6c756d652065786365656473206361700000000000000000000000000000604482015260640161066a565b506001600160a01b039094166000908152600460209081526040808320939093556006905220929092555050565b600084815260076020526040902060030154156122355760405162461bcd60e51b815260206004820152601f60248201527f64656c61796564207472616e7366657220616c72656164792065786973747300604482015260640161066a565b604080516080810182526001600160a01b0380861682528481166020808401918252838501868152426060860190815260008b81526007909352918690209451855490851673ffffffffffffffffffffffffffffffffffffffff1991821617865592516001860180549190951693169290921790925551600283015551600390910155517fcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6906122e89086815260200190565b60405180910390a150505050565b600080600061230484612435565b9050612311600882612a2d565b925080600716600581111561232857612328612a6e565b915050915091565b6060600061233d83612435565b905060008184600001516123519190612a15565b905083602001515181111561236557600080fd5b8167ffffffffffffffff81111561237e5761237e612a84565b6040519080825280601f01601f1916602001820160405280156123a8576020820181803683370190505b50602080860151865192955091818601919083010160005b858110156123e25781810151838201526123db602082612a15565b90506123c0565b505050935250919050565b60006123f882612546565b92915050565b600060208251111561240f57600080fd5b60208201519050815160206124249190612a9a565b61242f906008612a4f565b1c919050565b602080820151825181019091015160009182805b600a8110156124b15783811a9150612462816007612a4f565b82607f16901b85179450816080166000141561249f57612483816001612a15565b86518790612492908390612a15565b9052509395945050505050565b806124a9816127f4565b915050612449565b50600080fd5b600081516020146124c757600080fd5b506020015190565b60008160058111156124e3576124e3612a6e565b14156124f7576124f282612435565b505050565b600281600581111561250b5761250b612a6e565b141561025c57600061251c83612435565b905080836000018181516125309190612a15565b905250602083015151835111156124f257600080fd5b6000815160141461255657600080fd5b50602001516c01000000000000000000000000900490565b60006020828403121561258057600080fd5b5035919050565b60008083601f84011261259957600080fd5b50813567ffffffffffffffff8111156125b157600080fd5b6020830191508360208260051b85010111156125cc57600080fd5b9250929050565b600080600080604085870312156125e957600080fd5b843567ffffffffffffffff8082111561260157600080fd5b61260d88838901612587565b9096509450602087013591508082111561262657600080fd5b5061263387828801612587565b95989497509550505050565b80356001600160a01b038116811461265657600080fd5b919050565b60006020828403121561266d57600080fd5b6126768261263f565b9392505050565b6000806000806080858703121561269357600080fd5b61269c8561263f565b9350602085013592506126b16040860161263f565b9150606085013567ffffffffffffffff811681146126ce57600080fd5b939692955090935050565b6000806000806000806000806080898b0312156126f557600080fd5b883567ffffffffffffffff8082111561270d57600080fd5b818b0191508b601f83011261272157600080fd5b81358181111561273057600080fd5b8c602082850101111561274257600080fd5b60209283019a509850908a0135908082111561275d57600080fd5b6127698c838d01612587565b909850965060408b013591508082111561278257600080fd5b61278e8c838d01612587565b909650945060608b01359150808211156127a757600080fd5b506127b48b828c01612587565b999c989b5096995094979396929594505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415612808576128086127de565b5060010190565b838152818360208301376000910160200190815292915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260006020808501808196508560051b810191508460005b878110156128d65782840389528135601e1988360301811261288d57600080fd5b8701803567ffffffffffffffff8111156128a657600080fd5b8036038913156128b557600080fd5b6128c28682898501612829565b9a87019a955050509084019060010161286c565b5091979650505050505050565b8183526000602080850194508260005b8581101561291f576001600160a01b0361290c8361263f565b16875295820195908201906001016128f3565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561295c57600080fd5b8260051b8083602087013760009401602001938452509192915050565b608081526000885180608084015260005b818110156129a7576020818c0181015160a086840101520161298a565b818111156129b957600060a083860101525b50601f01601f1916820182810360a090810160208501526129dd908201898b612852565b905082810360408401526129f28187896128e3565b90508281036060840152612a0781858761292a565b9a9950505050505050505050565b60008219821115612a2857612a286127de565b500190565b600082612a4a57634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615612a6957612a696127de565b500290565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600082821015612aac57612aac6127de565b50039056fea2646970667358221220080c2cf9b742cdee8abd8019f6a3719058c4a90ba7451eb42c52bace9f7e192a64736f6c63430008090033",
}

// PeggedTokenBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use PeggedTokenBridgeMetaData.ABI instead.
var PeggedTokenBridgeABI = PeggedTokenBridgeMetaData.ABI

// PeggedTokenBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PeggedTokenBridgeMetaData.Bin instead.
var PeggedTokenBridgeBin = PeggedTokenBridgeMetaData.Bin

// DeployPeggedTokenBridge deploys a new Ethereum contract, binding an instance of PeggedTokenBridge to it.
func DeployPeggedTokenBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address) (common.Address, *types.Transaction, *PeggedTokenBridge, error) {
	parsed, err := PeggedTokenBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PeggedTokenBridgeBin), backend, _sigsVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PeggedTokenBridge{PeggedTokenBridgeCaller: PeggedTokenBridgeCaller{contract: contract}, PeggedTokenBridgeTransactor: PeggedTokenBridgeTransactor{contract: contract}, PeggedTokenBridgeFilterer: PeggedTokenBridgeFilterer{contract: contract}}, nil
}

// PeggedTokenBridge is an auto generated Go binding around an Ethereum contract.
type PeggedTokenBridge struct {
	PeggedTokenBridgeCaller     // Read-only binding to the contract
	PeggedTokenBridgeTransactor // Write-only binding to the contract
	PeggedTokenBridgeFilterer   // Log filterer for contract events
}

// PeggedTokenBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeggedTokenBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggedTokenBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeggedTokenBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggedTokenBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeggedTokenBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggedTokenBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeggedTokenBridgeSession struct {
	Contract     *PeggedTokenBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PeggedTokenBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeggedTokenBridgeCallerSession struct {
	Contract *PeggedTokenBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PeggedTokenBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeggedTokenBridgeTransactorSession struct {
	Contract     *PeggedTokenBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PeggedTokenBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeggedTokenBridgeRaw struct {
	Contract *PeggedTokenBridge // Generic contract binding to access the raw methods on
}

// PeggedTokenBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeggedTokenBridgeCallerRaw struct {
	Contract *PeggedTokenBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// PeggedTokenBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeggedTokenBridgeTransactorRaw struct {
	Contract *PeggedTokenBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeggedTokenBridge creates a new instance of PeggedTokenBridge, bound to a specific deployed contract.
func NewPeggedTokenBridge(address common.Address, backend bind.ContractBackend) (*PeggedTokenBridge, error) {
	contract, err := bindPeggedTokenBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridge{PeggedTokenBridgeCaller: PeggedTokenBridgeCaller{contract: contract}, PeggedTokenBridgeTransactor: PeggedTokenBridgeTransactor{contract: contract}, PeggedTokenBridgeFilterer: PeggedTokenBridgeFilterer{contract: contract}}, nil
}

// NewPeggedTokenBridgeCaller creates a new read-only instance of PeggedTokenBridge, bound to a specific deployed contract.
func NewPeggedTokenBridgeCaller(address common.Address, caller bind.ContractCaller) (*PeggedTokenBridgeCaller, error) {
	contract, err := bindPeggedTokenBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeCaller{contract: contract}, nil
}

// NewPeggedTokenBridgeTransactor creates a new write-only instance of PeggedTokenBridge, bound to a specific deployed contract.
func NewPeggedTokenBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PeggedTokenBridgeTransactor, error) {
	contract, err := bindPeggedTokenBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeTransactor{contract: contract}, nil
}

// NewPeggedTokenBridgeFilterer creates a new log filterer instance of PeggedTokenBridge, bound to a specific deployed contract.
func NewPeggedTokenBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PeggedTokenBridgeFilterer, error) {
	contract, err := bindPeggedTokenBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeFilterer{contract: contract}, nil
}

// bindPeggedTokenBridge binds a generic wrapper to an already deployed contract.
func bindPeggedTokenBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeggedTokenBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeggedTokenBridge *PeggedTokenBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeggedTokenBridge.Contract.PeggedTokenBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeggedTokenBridge *PeggedTokenBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.PeggedTokenBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeggedTokenBridge *PeggedTokenBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.PeggedTokenBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeggedTokenBridge *PeggedTokenBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeggedTokenBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.contract.Transact(opts, method, params...)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) DelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "delayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) DelayPeriod() (*big.Int, error) {
	return _PeggedTokenBridge.Contract.DelayPeriod(&_PeggedTokenBridge.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) DelayPeriod() (*big.Int, error) {
	return _PeggedTokenBridge.Contract.DelayPeriod(&_PeggedTokenBridge.CallOpts)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) DelayThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "delayThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.DelayThresholds(&_PeggedTokenBridge.CallOpts, arg0)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.DelayThresholds(&_PeggedTokenBridge.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) DelayedTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "delayedTransfers", arg0)

	outstruct := new(struct {
		Receiver  common.Address
		Token     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _PeggedTokenBridge.Contract.DelayedTransfers(&_PeggedTokenBridge.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _PeggedTokenBridge.Contract.DelayedTransfers(&_PeggedTokenBridge.CallOpts, arg0)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) EpochLength() (*big.Int, error) {
	return _PeggedTokenBridge.Contract.EpochLength(&_PeggedTokenBridge.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) EpochLength() (*big.Int, error) {
	return _PeggedTokenBridge.Contract.EpochLength(&_PeggedTokenBridge.CallOpts)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) EpochVolumeCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "epochVolumeCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.EpochVolumeCaps(&_PeggedTokenBridge.CallOpts, arg0)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.EpochVolumeCaps(&_PeggedTokenBridge.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) EpochVolumes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "epochVolumes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.EpochVolumes(&_PeggedTokenBridge.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.EpochVolumes(&_PeggedTokenBridge.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Governors(arg0 common.Address) (bool, error) {
	return _PeggedTokenBridge.Contract.Governors(&_PeggedTokenBridge.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _PeggedTokenBridge.Contract.Governors(&_PeggedTokenBridge.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) IsGovernor(_account common.Address) (bool, error) {
	return _PeggedTokenBridge.Contract.IsGovernor(&_PeggedTokenBridge.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _PeggedTokenBridge.Contract.IsGovernor(&_PeggedTokenBridge.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) IsPauser(account common.Address) (bool, error) {
	return _PeggedTokenBridge.Contract.IsPauser(&_PeggedTokenBridge.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) IsPauser(account common.Address) (bool, error) {
	return _PeggedTokenBridge.Contract.IsPauser(&_PeggedTokenBridge.CallOpts, account)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) LastOpTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "lastOpTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.LastOpTimestamps(&_PeggedTokenBridge.CallOpts, arg0)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.LastOpTimestamps(&_PeggedTokenBridge.CallOpts, arg0)
}

// MaxBurn is a free data retrieval call binding the contract method 0x497bf3b2.
//
// Solidity: function maxBurn(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) MaxBurn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "maxBurn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBurn is a free data retrieval call binding the contract method 0x497bf3b2.
//
// Solidity: function maxBurn(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) MaxBurn(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.MaxBurn(&_PeggedTokenBridge.CallOpts, arg0)
}

// MaxBurn is a free data retrieval call binding the contract method 0x497bf3b2.
//
// Solidity: function maxBurn(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) MaxBurn(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.MaxBurn(&_PeggedTokenBridge.CallOpts, arg0)
}

// MinBurn is a free data retrieval call binding the contract method 0x7f856013.
//
// Solidity: function minBurn(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) MinBurn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "minBurn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBurn is a free data retrieval call binding the contract method 0x7f856013.
//
// Solidity: function minBurn(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) MinBurn(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.MinBurn(&_PeggedTokenBridge.CallOpts, arg0)
}

// MinBurn is a free data retrieval call binding the contract method 0x7f856013.
//
// Solidity: function minBurn(address ) view returns(uint256)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) MinBurn(arg0 common.Address) (*big.Int, error) {
	return _PeggedTokenBridge.Contract.MinBurn(&_PeggedTokenBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Owner() (common.Address, error) {
	return _PeggedTokenBridge.Contract.Owner(&_PeggedTokenBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) Owner() (common.Address, error) {
	return _PeggedTokenBridge.Contract.Owner(&_PeggedTokenBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Paused() (bool, error) {
	return _PeggedTokenBridge.Contract.Paused(&_PeggedTokenBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) Paused() (bool, error) {
	return _PeggedTokenBridge.Contract.Paused(&_PeggedTokenBridge.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Pausers(arg0 common.Address) (bool, error) {
	return _PeggedTokenBridge.Contract.Pausers(&_PeggedTokenBridge.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _PeggedTokenBridge.Contract.Pausers(&_PeggedTokenBridge.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) Records(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "records", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Records(arg0 [32]byte) (bool, error) {
	return _PeggedTokenBridge.Contract.Records(&_PeggedTokenBridge.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) Records(arg0 [32]byte) (bool, error) {
	return _PeggedTokenBridge.Contract.Records(&_PeggedTokenBridge.CallOpts, arg0)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_PeggedTokenBridge *PeggedTokenBridgeCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeggedTokenBridge.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_PeggedTokenBridge *PeggedTokenBridgeSession) SigsVerifier() (common.Address, error) {
	return _PeggedTokenBridge.Contract.SigsVerifier(&_PeggedTokenBridge.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_PeggedTokenBridge *PeggedTokenBridgeCallerSession) SigsVerifier() (common.Address, error) {
	return _PeggedTokenBridge.Contract.SigsVerifier(&_PeggedTokenBridge.CallOpts)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.AddGovernor(&_PeggedTokenBridge.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.AddGovernor(&_PeggedTokenBridge.TransactOpts, _account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.AddPauser(&_PeggedTokenBridge.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.AddPauser(&_PeggedTokenBridge.TransactOpts, account)
}

// Burn is a paid mutator transaction binding the contract method 0xde790c7e.
//
// Solidity: function burn(address _token, uint256 _amount, address _withdrawAccount, uint64 _nonce) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) Burn(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _withdrawAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "burn", _token, _amount, _withdrawAccount, _nonce)
}

// Burn is a paid mutator transaction binding the contract method 0xde790c7e.
//
// Solidity: function burn(address _token, uint256 _amount, address _withdrawAccount, uint64 _nonce) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Burn(_token common.Address, _amount *big.Int, _withdrawAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.Burn(&_PeggedTokenBridge.TransactOpts, _token, _amount, _withdrawAccount, _nonce)
}

// Burn is a paid mutator transaction binding the contract method 0xde790c7e.
//
// Solidity: function burn(address _token, uint256 _amount, address _withdrawAccount, uint64 _nonce) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) Burn(_token common.Address, _amount *big.Int, _withdrawAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.Burn(&_PeggedTokenBridge.TransactOpts, _token, _amount, _withdrawAccount, _nonce)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) ExecuteDelayedTransfer(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "executeDelayedTransfer", id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.ExecuteDelayedTransfer(&_PeggedTokenBridge.TransactOpts, id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.ExecuteDelayedTransfer(&_PeggedTokenBridge.TransactOpts, id)
}

// Mint is a paid mutator transaction binding the contract method 0xf8734302.
//
// Solidity: function mint(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) Mint(opts *bind.TransactOpts, _request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "mint", _request, _sigs, _signers, _powers)
}

// Mint is a paid mutator transaction binding the contract method 0xf8734302.
//
// Solidity: function mint(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Mint(_request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.Mint(&_PeggedTokenBridge.TransactOpts, _request, _sigs, _signers, _powers)
}

// Mint is a paid mutator transaction binding the contract method 0xf8734302.
//
// Solidity: function mint(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) Mint(_request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.Mint(&_PeggedTokenBridge.TransactOpts, _request, _sigs, _signers, _powers)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Pause() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.Pause(&_PeggedTokenBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.Pause(&_PeggedTokenBridge.TransactOpts)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RemoveGovernor(&_PeggedTokenBridge.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RemoveGovernor(&_PeggedTokenBridge.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RemovePauser(&_PeggedTokenBridge.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RemovePauser(&_PeggedTokenBridge.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) RenounceGovernor() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RenounceGovernor(&_PeggedTokenBridge.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RenounceGovernor(&_PeggedTokenBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RenounceOwnership(&_PeggedTokenBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RenounceOwnership(&_PeggedTokenBridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) RenouncePauser() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RenouncePauser(&_PeggedTokenBridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.RenouncePauser(&_PeggedTokenBridge.TransactOpts)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) SetDelayPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "setDelayPeriod", _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetDelayPeriod(&_PeggedTokenBridge.TransactOpts, _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetDelayPeriod(&_PeggedTokenBridge.TransactOpts, _period)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) SetDelayThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "setDelayThresholds", _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetDelayThresholds(&_PeggedTokenBridge.TransactOpts, _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetDelayThresholds(&_PeggedTokenBridge.TransactOpts, _tokens, _thresholds)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) SetEpochLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "setEpochLength", _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetEpochLength(&_PeggedTokenBridge.TransactOpts, _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetEpochLength(&_PeggedTokenBridge.TransactOpts, _length)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) SetEpochVolumeCaps(opts *bind.TransactOpts, _tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "setEpochVolumeCaps", _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetEpochVolumeCaps(&_PeggedTokenBridge.TransactOpts, _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetEpochVolumeCaps(&_PeggedTokenBridge.TransactOpts, _tokens, _caps)
}

// SetMaxBurn is a paid mutator transaction binding the contract method 0xf9a8ea08.
//
// Solidity: function setMaxBurn(address[] _tokens, uint256[] _amounts) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) SetMaxBurn(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "setMaxBurn", _tokens, _amounts)
}

// SetMaxBurn is a paid mutator transaction binding the contract method 0xf9a8ea08.
//
// Solidity: function setMaxBurn(address[] _tokens, uint256[] _amounts) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) SetMaxBurn(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetMaxBurn(&_PeggedTokenBridge.TransactOpts, _tokens, _amounts)
}

// SetMaxBurn is a paid mutator transaction binding the contract method 0xf9a8ea08.
//
// Solidity: function setMaxBurn(address[] _tokens, uint256[] _amounts) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) SetMaxBurn(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetMaxBurn(&_PeggedTokenBridge.TransactOpts, _tokens, _amounts)
}

// SetMinBurn is a paid mutator transaction binding the contract method 0xbf4816f0.
//
// Solidity: function setMinBurn(address[] _tokens, uint256[] _amounts) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) SetMinBurn(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "setMinBurn", _tokens, _amounts)
}

// SetMinBurn is a paid mutator transaction binding the contract method 0xbf4816f0.
//
// Solidity: function setMinBurn(address[] _tokens, uint256[] _amounts) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) SetMinBurn(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetMinBurn(&_PeggedTokenBridge.TransactOpts, _tokens, _amounts)
}

// SetMinBurn is a paid mutator transaction binding the contract method 0xbf4816f0.
//
// Solidity: function setMinBurn(address[] _tokens, uint256[] _amounts) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) SetMinBurn(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.SetMinBurn(&_PeggedTokenBridge.TransactOpts, _tokens, _amounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.TransferOwnership(&_PeggedTokenBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.TransferOwnership(&_PeggedTokenBridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeggedTokenBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeSession) Unpause() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.Unpause(&_PeggedTokenBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PeggedTokenBridge *PeggedTokenBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _PeggedTokenBridge.Contract.Unpause(&_PeggedTokenBridge.TransactOpts)
}

// PeggedTokenBridgeBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeBurnIterator struct {
	Event *PeggedTokenBridgeBurn // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeBurn)
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
		it.Event = new(PeggedTokenBridgeBurn)
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
func (it *PeggedTokenBridgeBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeBurn represents a Burn event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeBurn struct {
	BurnId          [32]byte
	Token           common.Address
	Account         common.Address
	Amount          *big.Int
	WithdrawAccount common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x75f1bf55bb1de41b63a775dc7d4500f01114ee62b688a6b11d34f4692c1f3d43.
//
// Solidity: event Burn(bytes32 burnId, address token, address account, uint256 amount, address withdrawAccount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterBurn(opts *bind.FilterOpts) (*PeggedTokenBridgeBurnIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeBurnIterator{contract: _PeggedTokenBridge.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x75f1bf55bb1de41b63a775dc7d4500f01114ee62b688a6b11d34f4692c1f3d43.
//
// Solidity: event Burn(bytes32 burnId, address token, address account, uint256 amount, address withdrawAccount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeBurn) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeBurn)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x75f1bf55bb1de41b63a775dc7d4500f01114ee62b688a6b11d34f4692c1f3d43.
//
// Solidity: event Burn(bytes32 burnId, address token, address account, uint256 amount, address withdrawAccount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseBurn(log types.Log) (*PeggedTokenBridgeBurn, error) {
	event := new(PeggedTokenBridgeBurn)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeDelayPeriodUpdatedIterator is returned from FilterDelayPeriodUpdated and is used to iterate over the raw logs and unpacked data for DelayPeriodUpdated events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeDelayPeriodUpdatedIterator struct {
	Event *PeggedTokenBridgeDelayPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeDelayPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeDelayPeriodUpdated)
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
		it.Event = new(PeggedTokenBridgeDelayPeriodUpdated)
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
func (it *PeggedTokenBridgeDelayPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeDelayPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeDelayPeriodUpdated represents a DelayPeriodUpdated event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeDelayPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelayPeriodUpdated is a free log retrieval operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterDelayPeriodUpdated(opts *bind.FilterOpts) (*PeggedTokenBridgeDelayPeriodUpdatedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeDelayPeriodUpdatedIterator{contract: _PeggedTokenBridge.contract, event: "DelayPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayPeriodUpdated is a free log subscription operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchDelayPeriodUpdated(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeDelayPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeDelayPeriodUpdated)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
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

// ParseDelayPeriodUpdated is a log parse operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseDelayPeriodUpdated(log types.Log) (*PeggedTokenBridgeDelayPeriodUpdated, error) {
	event := new(PeggedTokenBridgeDelayPeriodUpdated)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeDelayThresholdUpdatedIterator is returned from FilterDelayThresholdUpdated and is used to iterate over the raw logs and unpacked data for DelayThresholdUpdated events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeDelayThresholdUpdatedIterator struct {
	Event *PeggedTokenBridgeDelayThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeDelayThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeDelayThresholdUpdated)
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
		it.Event = new(PeggedTokenBridgeDelayThresholdUpdated)
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
func (it *PeggedTokenBridgeDelayThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeDelayThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeDelayThresholdUpdated represents a DelayThresholdUpdated event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeDelayThresholdUpdated struct {
	Token     common.Address
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayThresholdUpdated is a free log retrieval operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterDelayThresholdUpdated(opts *bind.FilterOpts) (*PeggedTokenBridgeDelayThresholdUpdatedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeDelayThresholdUpdatedIterator{contract: _PeggedTokenBridge.contract, event: "DelayThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayThresholdUpdated is a free log subscription operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchDelayThresholdUpdated(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeDelayThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeDelayThresholdUpdated)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
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

// ParseDelayThresholdUpdated is a log parse operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseDelayThresholdUpdated(log types.Log) (*PeggedTokenBridgeDelayThresholdUpdated, error) {
	event := new(PeggedTokenBridgeDelayThresholdUpdated)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeDelayedTransferAddedIterator is returned from FilterDelayedTransferAdded and is used to iterate over the raw logs and unpacked data for DelayedTransferAdded events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeDelayedTransferAddedIterator struct {
	Event *PeggedTokenBridgeDelayedTransferAdded // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeDelayedTransferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeDelayedTransferAdded)
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
		it.Event = new(PeggedTokenBridgeDelayedTransferAdded)
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
func (it *PeggedTokenBridgeDelayedTransferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeDelayedTransferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeDelayedTransferAdded represents a DelayedTransferAdded event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeDelayedTransferAdded struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferAdded is a free log retrieval operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterDelayedTransferAdded(opts *bind.FilterOpts) (*PeggedTokenBridgeDelayedTransferAddedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeDelayedTransferAddedIterator{contract: _PeggedTokenBridge.contract, event: "DelayedTransferAdded", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferAdded is a free log subscription operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchDelayedTransferAdded(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeDelayedTransferAdded) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeDelayedTransferAdded)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
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

// ParseDelayedTransferAdded is a log parse operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseDelayedTransferAdded(log types.Log) (*PeggedTokenBridgeDelayedTransferAdded, error) {
	event := new(PeggedTokenBridgeDelayedTransferAdded)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeDelayedTransferExecutedIterator is returned from FilterDelayedTransferExecuted and is used to iterate over the raw logs and unpacked data for DelayedTransferExecuted events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeDelayedTransferExecutedIterator struct {
	Event *PeggedTokenBridgeDelayedTransferExecuted // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeDelayedTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeDelayedTransferExecuted)
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
		it.Event = new(PeggedTokenBridgeDelayedTransferExecuted)
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
func (it *PeggedTokenBridgeDelayedTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeDelayedTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeDelayedTransferExecuted represents a DelayedTransferExecuted event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeDelayedTransferExecuted struct {
	Id       [32]byte
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferExecuted is a free log retrieval operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterDelayedTransferExecuted(opts *bind.FilterOpts) (*PeggedTokenBridgeDelayedTransferExecutedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeDelayedTransferExecutedIterator{contract: _PeggedTokenBridge.contract, event: "DelayedTransferExecuted", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferExecuted is a free log subscription operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchDelayedTransferExecuted(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeDelayedTransferExecuted) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeDelayedTransferExecuted)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
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

// ParseDelayedTransferExecuted is a log parse operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseDelayedTransferExecuted(log types.Log) (*PeggedTokenBridgeDelayedTransferExecuted, error) {
	event := new(PeggedTokenBridgeDelayedTransferExecuted)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeEpochLengthUpdatedIterator is returned from FilterEpochLengthUpdated and is used to iterate over the raw logs and unpacked data for EpochLengthUpdated events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeEpochLengthUpdatedIterator struct {
	Event *PeggedTokenBridgeEpochLengthUpdated // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeEpochLengthUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeEpochLengthUpdated)
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
		it.Event = new(PeggedTokenBridgeEpochLengthUpdated)
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
func (it *PeggedTokenBridgeEpochLengthUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeEpochLengthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeEpochLengthUpdated represents a EpochLengthUpdated event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeEpochLengthUpdated struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEpochLengthUpdated is a free log retrieval operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterEpochLengthUpdated(opts *bind.FilterOpts) (*PeggedTokenBridgeEpochLengthUpdatedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeEpochLengthUpdatedIterator{contract: _PeggedTokenBridge.contract, event: "EpochLengthUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochLengthUpdated is a free log subscription operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchEpochLengthUpdated(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeEpochLengthUpdated) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeEpochLengthUpdated)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
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

// ParseEpochLengthUpdated is a log parse operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseEpochLengthUpdated(log types.Log) (*PeggedTokenBridgeEpochLengthUpdated, error) {
	event := new(PeggedTokenBridgeEpochLengthUpdated)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeEpochVolumeUpdatedIterator is returned from FilterEpochVolumeUpdated and is used to iterate over the raw logs and unpacked data for EpochVolumeUpdated events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeEpochVolumeUpdatedIterator struct {
	Event *PeggedTokenBridgeEpochVolumeUpdated // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeEpochVolumeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeEpochVolumeUpdated)
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
		it.Event = new(PeggedTokenBridgeEpochVolumeUpdated)
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
func (it *PeggedTokenBridgeEpochVolumeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeEpochVolumeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeEpochVolumeUpdated represents a EpochVolumeUpdated event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeEpochVolumeUpdated struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEpochVolumeUpdated is a free log retrieval operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterEpochVolumeUpdated(opts *bind.FilterOpts) (*PeggedTokenBridgeEpochVolumeUpdatedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeEpochVolumeUpdatedIterator{contract: _PeggedTokenBridge.contract, event: "EpochVolumeUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochVolumeUpdated is a free log subscription operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchEpochVolumeUpdated(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeEpochVolumeUpdated) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeEpochVolumeUpdated)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
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

// ParseEpochVolumeUpdated is a log parse operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseEpochVolumeUpdated(log types.Log) (*PeggedTokenBridgeEpochVolumeUpdated, error) {
	event := new(PeggedTokenBridgeEpochVolumeUpdated)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeGovernorAddedIterator struct {
	Event *PeggedTokenBridgeGovernorAdded // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeGovernorAdded)
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
		it.Event = new(PeggedTokenBridgeGovernorAdded)
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
func (it *PeggedTokenBridgeGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeGovernorAdded represents a GovernorAdded event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*PeggedTokenBridgeGovernorAddedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeGovernorAddedIterator{contract: _PeggedTokenBridge.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeGovernorAdded)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseGovernorAdded(log types.Log) (*PeggedTokenBridgeGovernorAdded, error) {
	event := new(PeggedTokenBridgeGovernorAdded)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeGovernorRemovedIterator struct {
	Event *PeggedTokenBridgeGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeGovernorRemoved)
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
		it.Event = new(PeggedTokenBridgeGovernorRemoved)
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
func (it *PeggedTokenBridgeGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeGovernorRemoved represents a GovernorRemoved event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*PeggedTokenBridgeGovernorRemovedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeGovernorRemovedIterator{contract: _PeggedTokenBridge.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeGovernorRemoved)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseGovernorRemoved(log types.Log) (*PeggedTokenBridgeGovernorRemoved, error) {
	event := new(PeggedTokenBridgeGovernorRemoved)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeMaxBurnUpdatedIterator is returned from FilterMaxBurnUpdated and is used to iterate over the raw logs and unpacked data for MaxBurnUpdated events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeMaxBurnUpdatedIterator struct {
	Event *PeggedTokenBridgeMaxBurnUpdated // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeMaxBurnUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeMaxBurnUpdated)
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
		it.Event = new(PeggedTokenBridgeMaxBurnUpdated)
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
func (it *PeggedTokenBridgeMaxBurnUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeMaxBurnUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeMaxBurnUpdated represents a MaxBurnUpdated event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeMaxBurnUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxBurnUpdated is a free log retrieval operation binding the contract event 0xa3181379f6db47d9037efc6b6e8e3efe8c55ddb090b4f0512c152f97c4e47da5.
//
// Solidity: event MaxBurnUpdated(address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterMaxBurnUpdated(opts *bind.FilterOpts) (*PeggedTokenBridgeMaxBurnUpdatedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "MaxBurnUpdated")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeMaxBurnUpdatedIterator{contract: _PeggedTokenBridge.contract, event: "MaxBurnUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxBurnUpdated is a free log subscription operation binding the contract event 0xa3181379f6db47d9037efc6b6e8e3efe8c55ddb090b4f0512c152f97c4e47da5.
//
// Solidity: event MaxBurnUpdated(address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchMaxBurnUpdated(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeMaxBurnUpdated) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "MaxBurnUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeMaxBurnUpdated)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "MaxBurnUpdated", log); err != nil {
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

// ParseMaxBurnUpdated is a log parse operation binding the contract event 0xa3181379f6db47d9037efc6b6e8e3efe8c55ddb090b4f0512c152f97c4e47da5.
//
// Solidity: event MaxBurnUpdated(address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseMaxBurnUpdated(log types.Log) (*PeggedTokenBridgeMaxBurnUpdated, error) {
	event := new(PeggedTokenBridgeMaxBurnUpdated)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "MaxBurnUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeMinBurnUpdatedIterator is returned from FilterMinBurnUpdated and is used to iterate over the raw logs and unpacked data for MinBurnUpdated events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeMinBurnUpdatedIterator struct {
	Event *PeggedTokenBridgeMinBurnUpdated // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeMinBurnUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeMinBurnUpdated)
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
		it.Event = new(PeggedTokenBridgeMinBurnUpdated)
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
func (it *PeggedTokenBridgeMinBurnUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeMinBurnUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeMinBurnUpdated represents a MinBurnUpdated event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeMinBurnUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinBurnUpdated is a free log retrieval operation binding the contract event 0x3796cd0b17a8734f8da819920625598e9a18be490f686725282e5383f1d06683.
//
// Solidity: event MinBurnUpdated(address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterMinBurnUpdated(opts *bind.FilterOpts) (*PeggedTokenBridgeMinBurnUpdatedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "MinBurnUpdated")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeMinBurnUpdatedIterator{contract: _PeggedTokenBridge.contract, event: "MinBurnUpdated", logs: logs, sub: sub}, nil
}

// WatchMinBurnUpdated is a free log subscription operation binding the contract event 0x3796cd0b17a8734f8da819920625598e9a18be490f686725282e5383f1d06683.
//
// Solidity: event MinBurnUpdated(address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchMinBurnUpdated(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeMinBurnUpdated) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "MinBurnUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeMinBurnUpdated)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "MinBurnUpdated", log); err != nil {
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

// ParseMinBurnUpdated is a log parse operation binding the contract event 0x3796cd0b17a8734f8da819920625598e9a18be490f686725282e5383f1d06683.
//
// Solidity: event MinBurnUpdated(address token, uint256 amount)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseMinBurnUpdated(log types.Log) (*PeggedTokenBridgeMinBurnUpdated, error) {
	event := new(PeggedTokenBridgeMinBurnUpdated)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "MinBurnUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeMintIterator struct {
	Event *PeggedTokenBridgeMint // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeMint)
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
		it.Event = new(PeggedTokenBridgeMint)
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
func (it *PeggedTokenBridgeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeMint represents a Mint event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeMint struct {
	MintId     [32]byte
	Token      common.Address
	Account    common.Address
	Amount     *big.Int
	RefChainId uint64
	RefId      [32]byte
	Depositor  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x5bc84ecccfced5bb04bfc7f3efcdbe7f5cd21949ef146811b4d1967fe41f777a.
//
// Solidity: event Mint(bytes32 mintId, address token, address account, uint256 amount, uint64 refChainId, bytes32 refId, address depositor)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterMint(opts *bind.FilterOpts) (*PeggedTokenBridgeMintIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeMintIterator{contract: _PeggedTokenBridge.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x5bc84ecccfced5bb04bfc7f3efcdbe7f5cd21949ef146811b4d1967fe41f777a.
//
// Solidity: event Mint(bytes32 mintId, address token, address account, uint256 amount, uint64 refChainId, bytes32 refId, address depositor)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeMint) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeMint)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x5bc84ecccfced5bb04bfc7f3efcdbe7f5cd21949ef146811b4d1967fe41f777a.
//
// Solidity: event Mint(bytes32 mintId, address token, address account, uint256 amount, uint64 refChainId, bytes32 refId, address depositor)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseMint(log types.Log) (*PeggedTokenBridgeMint, error) {
	event := new(PeggedTokenBridgeMint)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeOwnershipTransferredIterator struct {
	Event *PeggedTokenBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeOwnershipTransferred)
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
		it.Event = new(PeggedTokenBridgeOwnershipTransferred)
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
func (it *PeggedTokenBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PeggedTokenBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeOwnershipTransferredIterator{contract: _PeggedTokenBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeOwnershipTransferred)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*PeggedTokenBridgeOwnershipTransferred, error) {
	event := new(PeggedTokenBridgeOwnershipTransferred)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgePausedIterator struct {
	Event *PeggedTokenBridgePaused // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgePaused)
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
		it.Event = new(PeggedTokenBridgePaused)
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
func (it *PeggedTokenBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgePaused represents a Paused event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*PeggedTokenBridgePausedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgePausedIterator{contract: _PeggedTokenBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgePaused) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgePaused)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParsePaused(log types.Log) (*PeggedTokenBridgePaused, error) {
	event := new(PeggedTokenBridgePaused)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgePauserAddedIterator struct {
	Event *PeggedTokenBridgePauserAdded // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgePauserAdded)
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
		it.Event = new(PeggedTokenBridgePauserAdded)
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
func (it *PeggedTokenBridgePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgePauserAdded represents a PauserAdded event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*PeggedTokenBridgePauserAddedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgePauserAddedIterator{contract: _PeggedTokenBridge.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgePauserAdded) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgePauserAdded)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParsePauserAdded(log types.Log) (*PeggedTokenBridgePauserAdded, error) {
	event := new(PeggedTokenBridgePauserAdded)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgePauserRemovedIterator struct {
	Event *PeggedTokenBridgePauserRemoved // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgePauserRemoved)
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
		it.Event = new(PeggedTokenBridgePauserRemoved)
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
func (it *PeggedTokenBridgePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgePauserRemoved represents a PauserRemoved event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*PeggedTokenBridgePauserRemovedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgePauserRemovedIterator{contract: _PeggedTokenBridge.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgePauserRemoved) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgePauserRemoved)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParsePauserRemoved(log types.Log) (*PeggedTokenBridgePauserRemoved, error) {
	event := new(PeggedTokenBridgePauserRemoved)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggedTokenBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeUnpausedIterator struct {
	Event *PeggedTokenBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *PeggedTokenBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggedTokenBridgeUnpaused)
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
		it.Event = new(PeggedTokenBridgeUnpaused)
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
func (it *PeggedTokenBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggedTokenBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggedTokenBridgeUnpaused represents a Unpaused event raised by the PeggedTokenBridge contract.
type PeggedTokenBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PeggedTokenBridgeUnpausedIterator, error) {

	logs, sub, err := _PeggedTokenBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PeggedTokenBridgeUnpausedIterator{contract: _PeggedTokenBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PeggedTokenBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _PeggedTokenBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggedTokenBridgeUnpaused)
				if err := _PeggedTokenBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PeggedTokenBridge *PeggedTokenBridgeFilterer) ParseUnpaused(log types.Log) (*PeggedTokenBridgeUnpaused, error) {
	event := new(PeggedTokenBridgeUnpaused)
	if err := _PeggedTokenBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolumeControlMetaData contains all meta data concerning the VolumeControl contract.
var VolumeControlMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EpochLengthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EpochVolumeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumeCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastOpTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"name\":\"setEpochVolumeCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolumeControlABI is the input ABI used to generate the binding from.
// Deprecated: Use VolumeControlMetaData.ABI instead.
var VolumeControlABI = VolumeControlMetaData.ABI

// VolumeControl is an auto generated Go binding around an Ethereum contract.
type VolumeControl struct {
	VolumeControlCaller     // Read-only binding to the contract
	VolumeControlTransactor // Write-only binding to the contract
	VolumeControlFilterer   // Log filterer for contract events
}

// VolumeControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolumeControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolumeControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolumeControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolumeControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolumeControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolumeControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolumeControlSession struct {
	Contract     *VolumeControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VolumeControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolumeControlCallerSession struct {
	Contract *VolumeControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VolumeControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolumeControlTransactorSession struct {
	Contract     *VolumeControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VolumeControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolumeControlRaw struct {
	Contract *VolumeControl // Generic contract binding to access the raw methods on
}

// VolumeControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolumeControlCallerRaw struct {
	Contract *VolumeControlCaller // Generic read-only contract binding to access the raw methods on
}

// VolumeControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolumeControlTransactorRaw struct {
	Contract *VolumeControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolumeControl creates a new instance of VolumeControl, bound to a specific deployed contract.
func NewVolumeControl(address common.Address, backend bind.ContractBackend) (*VolumeControl, error) {
	contract, err := bindVolumeControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolumeControl{VolumeControlCaller: VolumeControlCaller{contract: contract}, VolumeControlTransactor: VolumeControlTransactor{contract: contract}, VolumeControlFilterer: VolumeControlFilterer{contract: contract}}, nil
}

// NewVolumeControlCaller creates a new read-only instance of VolumeControl, bound to a specific deployed contract.
func NewVolumeControlCaller(address common.Address, caller bind.ContractCaller) (*VolumeControlCaller, error) {
	contract, err := bindVolumeControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolumeControlCaller{contract: contract}, nil
}

// NewVolumeControlTransactor creates a new write-only instance of VolumeControl, bound to a specific deployed contract.
func NewVolumeControlTransactor(address common.Address, transactor bind.ContractTransactor) (*VolumeControlTransactor, error) {
	contract, err := bindVolumeControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolumeControlTransactor{contract: contract}, nil
}

// NewVolumeControlFilterer creates a new log filterer instance of VolumeControl, bound to a specific deployed contract.
func NewVolumeControlFilterer(address common.Address, filterer bind.ContractFilterer) (*VolumeControlFilterer, error) {
	contract, err := bindVolumeControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolumeControlFilterer{contract: contract}, nil
}

// bindVolumeControl binds a generic wrapper to an already deployed contract.
func bindVolumeControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VolumeControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolumeControl *VolumeControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolumeControl.Contract.VolumeControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolumeControl *VolumeControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolumeControl.Contract.VolumeControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolumeControl *VolumeControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolumeControl.Contract.VolumeControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolumeControl *VolumeControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolumeControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolumeControl *VolumeControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolumeControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolumeControl *VolumeControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolumeControl.Contract.contract.Transact(opts, method, params...)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_VolumeControl *VolumeControlCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolumeControl.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_VolumeControl *VolumeControlSession) EpochLength() (*big.Int, error) {
	return _VolumeControl.Contract.EpochLength(&_VolumeControl.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_VolumeControl *VolumeControlCallerSession) EpochLength() (*big.Int, error) {
	return _VolumeControl.Contract.EpochLength(&_VolumeControl.CallOpts)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_VolumeControl *VolumeControlCaller) EpochVolumeCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolumeControl.contract.Call(opts, &out, "epochVolumeCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_VolumeControl *VolumeControlSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _VolumeControl.Contract.EpochVolumeCaps(&_VolumeControl.CallOpts, arg0)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_VolumeControl *VolumeControlCallerSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _VolumeControl.Contract.EpochVolumeCaps(&_VolumeControl.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_VolumeControl *VolumeControlCaller) EpochVolumes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolumeControl.contract.Call(opts, &out, "epochVolumes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_VolumeControl *VolumeControlSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _VolumeControl.Contract.EpochVolumes(&_VolumeControl.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_VolumeControl *VolumeControlCallerSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _VolumeControl.Contract.EpochVolumes(&_VolumeControl.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_VolumeControl *VolumeControlCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VolumeControl.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_VolumeControl *VolumeControlSession) Governors(arg0 common.Address) (bool, error) {
	return _VolumeControl.Contract.Governors(&_VolumeControl.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_VolumeControl *VolumeControlCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _VolumeControl.Contract.Governors(&_VolumeControl.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_VolumeControl *VolumeControlCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _VolumeControl.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_VolumeControl *VolumeControlSession) IsGovernor(_account common.Address) (bool, error) {
	return _VolumeControl.Contract.IsGovernor(&_VolumeControl.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_VolumeControl *VolumeControlCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _VolumeControl.Contract.IsGovernor(&_VolumeControl.CallOpts, _account)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_VolumeControl *VolumeControlCaller) LastOpTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolumeControl.contract.Call(opts, &out, "lastOpTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_VolumeControl *VolumeControlSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _VolumeControl.Contract.LastOpTimestamps(&_VolumeControl.CallOpts, arg0)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_VolumeControl *VolumeControlCallerSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _VolumeControl.Contract.LastOpTimestamps(&_VolumeControl.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolumeControl *VolumeControlCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolumeControl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolumeControl *VolumeControlSession) Owner() (common.Address, error) {
	return _VolumeControl.Contract.Owner(&_VolumeControl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolumeControl *VolumeControlCallerSession) Owner() (common.Address, error) {
	return _VolumeControl.Contract.Owner(&_VolumeControl.CallOpts)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_VolumeControl *VolumeControlTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _VolumeControl.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_VolumeControl *VolumeControlSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _VolumeControl.Contract.AddGovernor(&_VolumeControl.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_VolumeControl *VolumeControlTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _VolumeControl.Contract.AddGovernor(&_VolumeControl.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_VolumeControl *VolumeControlTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _VolumeControl.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_VolumeControl *VolumeControlSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _VolumeControl.Contract.RemoveGovernor(&_VolumeControl.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_VolumeControl *VolumeControlTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _VolumeControl.Contract.RemoveGovernor(&_VolumeControl.TransactOpts, _account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_VolumeControl *VolumeControlTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolumeControl.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_VolumeControl *VolumeControlSession) RenounceGovernor() (*types.Transaction, error) {
	return _VolumeControl.Contract.RenounceGovernor(&_VolumeControl.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_VolumeControl *VolumeControlTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _VolumeControl.Contract.RenounceGovernor(&_VolumeControl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolumeControl *VolumeControlTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolumeControl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolumeControl *VolumeControlSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolumeControl.Contract.RenounceOwnership(&_VolumeControl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolumeControl *VolumeControlTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolumeControl.Contract.RenounceOwnership(&_VolumeControl.TransactOpts)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_VolumeControl *VolumeControlTransactor) SetEpochLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _VolumeControl.contract.Transact(opts, "setEpochLength", _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_VolumeControl *VolumeControlSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _VolumeControl.Contract.SetEpochLength(&_VolumeControl.TransactOpts, _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_VolumeControl *VolumeControlTransactorSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _VolumeControl.Contract.SetEpochLength(&_VolumeControl.TransactOpts, _length)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_VolumeControl *VolumeControlTransactor) SetEpochVolumeCaps(opts *bind.TransactOpts, _tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _VolumeControl.contract.Transact(opts, "setEpochVolumeCaps", _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_VolumeControl *VolumeControlSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _VolumeControl.Contract.SetEpochVolumeCaps(&_VolumeControl.TransactOpts, _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_VolumeControl *VolumeControlTransactorSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _VolumeControl.Contract.SetEpochVolumeCaps(&_VolumeControl.TransactOpts, _tokens, _caps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolumeControl *VolumeControlTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolumeControl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolumeControl *VolumeControlSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolumeControl.Contract.TransferOwnership(&_VolumeControl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolumeControl *VolumeControlTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolumeControl.Contract.TransferOwnership(&_VolumeControl.TransactOpts, newOwner)
}

// VolumeControlEpochLengthUpdatedIterator is returned from FilterEpochLengthUpdated and is used to iterate over the raw logs and unpacked data for EpochLengthUpdated events raised by the VolumeControl contract.
type VolumeControlEpochLengthUpdatedIterator struct {
	Event *VolumeControlEpochLengthUpdated // Event containing the contract specifics and raw log

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
func (it *VolumeControlEpochLengthUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolumeControlEpochLengthUpdated)
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
		it.Event = new(VolumeControlEpochLengthUpdated)
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
func (it *VolumeControlEpochLengthUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolumeControlEpochLengthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolumeControlEpochLengthUpdated represents a EpochLengthUpdated event raised by the VolumeControl contract.
type VolumeControlEpochLengthUpdated struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEpochLengthUpdated is a free log retrieval operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_VolumeControl *VolumeControlFilterer) FilterEpochLengthUpdated(opts *bind.FilterOpts) (*VolumeControlEpochLengthUpdatedIterator, error) {

	logs, sub, err := _VolumeControl.contract.FilterLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return &VolumeControlEpochLengthUpdatedIterator{contract: _VolumeControl.contract, event: "EpochLengthUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochLengthUpdated is a free log subscription operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_VolumeControl *VolumeControlFilterer) WatchEpochLengthUpdated(opts *bind.WatchOpts, sink chan<- *VolumeControlEpochLengthUpdated) (event.Subscription, error) {

	logs, sub, err := _VolumeControl.contract.WatchLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolumeControlEpochLengthUpdated)
				if err := _VolumeControl.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
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

// ParseEpochLengthUpdated is a log parse operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_VolumeControl *VolumeControlFilterer) ParseEpochLengthUpdated(log types.Log) (*VolumeControlEpochLengthUpdated, error) {
	event := new(VolumeControlEpochLengthUpdated)
	if err := _VolumeControl.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolumeControlEpochVolumeUpdatedIterator is returned from FilterEpochVolumeUpdated and is used to iterate over the raw logs and unpacked data for EpochVolumeUpdated events raised by the VolumeControl contract.
type VolumeControlEpochVolumeUpdatedIterator struct {
	Event *VolumeControlEpochVolumeUpdated // Event containing the contract specifics and raw log

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
func (it *VolumeControlEpochVolumeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolumeControlEpochVolumeUpdated)
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
		it.Event = new(VolumeControlEpochVolumeUpdated)
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
func (it *VolumeControlEpochVolumeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolumeControlEpochVolumeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolumeControlEpochVolumeUpdated represents a EpochVolumeUpdated event raised by the VolumeControl contract.
type VolumeControlEpochVolumeUpdated struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEpochVolumeUpdated is a free log retrieval operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_VolumeControl *VolumeControlFilterer) FilterEpochVolumeUpdated(opts *bind.FilterOpts) (*VolumeControlEpochVolumeUpdatedIterator, error) {

	logs, sub, err := _VolumeControl.contract.FilterLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return &VolumeControlEpochVolumeUpdatedIterator{contract: _VolumeControl.contract, event: "EpochVolumeUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochVolumeUpdated is a free log subscription operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_VolumeControl *VolumeControlFilterer) WatchEpochVolumeUpdated(opts *bind.WatchOpts, sink chan<- *VolumeControlEpochVolumeUpdated) (event.Subscription, error) {

	logs, sub, err := _VolumeControl.contract.WatchLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolumeControlEpochVolumeUpdated)
				if err := _VolumeControl.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
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

// ParseEpochVolumeUpdated is a log parse operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_VolumeControl *VolumeControlFilterer) ParseEpochVolumeUpdated(log types.Log) (*VolumeControlEpochVolumeUpdated, error) {
	event := new(VolumeControlEpochVolumeUpdated)
	if err := _VolumeControl.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolumeControlGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the VolumeControl contract.
type VolumeControlGovernorAddedIterator struct {
	Event *VolumeControlGovernorAdded // Event containing the contract specifics and raw log

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
func (it *VolumeControlGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolumeControlGovernorAdded)
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
		it.Event = new(VolumeControlGovernorAdded)
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
func (it *VolumeControlGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolumeControlGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolumeControlGovernorAdded represents a GovernorAdded event raised by the VolumeControl contract.
type VolumeControlGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_VolumeControl *VolumeControlFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*VolumeControlGovernorAddedIterator, error) {

	logs, sub, err := _VolumeControl.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &VolumeControlGovernorAddedIterator{contract: _VolumeControl.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_VolumeControl *VolumeControlFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *VolumeControlGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _VolumeControl.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolumeControlGovernorAdded)
				if err := _VolumeControl.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_VolumeControl *VolumeControlFilterer) ParseGovernorAdded(log types.Log) (*VolumeControlGovernorAdded, error) {
	event := new(VolumeControlGovernorAdded)
	if err := _VolumeControl.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolumeControlGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the VolumeControl contract.
type VolumeControlGovernorRemovedIterator struct {
	Event *VolumeControlGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *VolumeControlGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolumeControlGovernorRemoved)
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
		it.Event = new(VolumeControlGovernorRemoved)
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
func (it *VolumeControlGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolumeControlGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolumeControlGovernorRemoved represents a GovernorRemoved event raised by the VolumeControl contract.
type VolumeControlGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_VolumeControl *VolumeControlFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*VolumeControlGovernorRemovedIterator, error) {

	logs, sub, err := _VolumeControl.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &VolumeControlGovernorRemovedIterator{contract: _VolumeControl.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_VolumeControl *VolumeControlFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *VolumeControlGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _VolumeControl.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolumeControlGovernorRemoved)
				if err := _VolumeControl.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_VolumeControl *VolumeControlFilterer) ParseGovernorRemoved(log types.Log) (*VolumeControlGovernorRemoved, error) {
	event := new(VolumeControlGovernorRemoved)
	if err := _VolumeControl.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolumeControlOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolumeControl contract.
type VolumeControlOwnershipTransferredIterator struct {
	Event *VolumeControlOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolumeControlOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolumeControlOwnershipTransferred)
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
		it.Event = new(VolumeControlOwnershipTransferred)
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
func (it *VolumeControlOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolumeControlOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolumeControlOwnershipTransferred represents a OwnershipTransferred event raised by the VolumeControl contract.
type VolumeControlOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolumeControl *VolumeControlFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolumeControlOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolumeControl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolumeControlOwnershipTransferredIterator{contract: _VolumeControl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolumeControl *VolumeControlFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolumeControlOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolumeControl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolumeControlOwnershipTransferred)
				if err := _VolumeControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolumeControl *VolumeControlFilterer) ParseOwnershipTransferred(log types.Log) (*VolumeControlOwnershipTransferred, error) {
	event := new(VolumeControlOwnershipTransferred)
	if err := _VolumeControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
