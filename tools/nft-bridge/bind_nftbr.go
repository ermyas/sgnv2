// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// NFTBridgeMetaData contains all meta data concerning the NFTBridge contract.
var NFTBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"}],\"name\":\"MessageBusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChid\",\"type\":\"uint64\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcNft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstNft\",\"type\":\"address\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChid\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dstNft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dstBridge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_backToOrigin\",\"type\":\"bool\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChid\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dstNft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dstBridge\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"destTxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChid\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chid\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setTxFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_dstChid\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NFTBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTBridgeMetaData.ABI instead.
var NFTBridgeABI = NFTBridgeMetaData.ABI

// NFTBridge is an auto generated Go binding around an Ethereum contract.
type NFTBridge struct {
	NFTBridgeCaller     // Read-only binding to the contract
	NFTBridgeTransactor // Write-only binding to the contract
	NFTBridgeFilterer   // Log filterer for contract events
}

// NFTBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTBridgeSession struct {
	Contract     *NFTBridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTBridgeCallerSession struct {
	Contract *NFTBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NFTBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTBridgeTransactorSession struct {
	Contract     *NFTBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NFTBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTBridgeRaw struct {
	Contract *NFTBridge // Generic contract binding to access the raw methods on
}

// NFTBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTBridgeCallerRaw struct {
	Contract *NFTBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// NFTBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTBridgeTransactorRaw struct {
	Contract *NFTBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTBridge creates a new instance of NFTBridge, bound to a specific deployed contract.
func NewNFTBridge(address common.Address, backend bind.ContractBackend) (*NFTBridge, error) {
	contract, err := bindNFTBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTBridge{NFTBridgeCaller: NFTBridgeCaller{contract: contract}, NFTBridgeTransactor: NFTBridgeTransactor{contract: contract}, NFTBridgeFilterer: NFTBridgeFilterer{contract: contract}}, nil
}

// NewNFTBridgeCaller creates a new read-only instance of NFTBridge, bound to a specific deployed contract.
func NewNFTBridgeCaller(address common.Address, caller bind.ContractCaller) (*NFTBridgeCaller, error) {
	contract, err := bindNFTBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTBridgeCaller{contract: contract}, nil
}

// NewNFTBridgeTransactor creates a new write-only instance of NFTBridge, bound to a specific deployed contract.
func NewNFTBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTBridgeTransactor, error) {
	contract, err := bindNFTBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTBridgeTransactor{contract: contract}, nil
}

// NewNFTBridgeFilterer creates a new log filterer instance of NFTBridge, bound to a specific deployed contract.
func NewNFTBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTBridgeFilterer, error) {
	contract, err := bindNFTBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTBridgeFilterer{contract: contract}, nil
}

// bindNFTBridge binds a generic wrapper to an already deployed contract.
func bindNFTBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTBridge *NFTBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTBridge.Contract.NFTBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTBridge *NFTBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBridge.Contract.NFTBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTBridge *NFTBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTBridge.Contract.NFTBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTBridge *NFTBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTBridge *NFTBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTBridge *NFTBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTBridge.Contract.contract.Transact(opts, method, params...)
}

// DestTxFee is a free data retrieval call binding the contract method 0x151ff4eb.
//
// Solidity: function destTxFee(uint64 ) view returns(uint256)
func (_NFTBridge *NFTBridgeCaller) DestTxFee(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "destTxFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestTxFee is a free data retrieval call binding the contract method 0x151ff4eb.
//
// Solidity: function destTxFee(uint64 ) view returns(uint256)
func (_NFTBridge *NFTBridgeSession) DestTxFee(arg0 uint64) (*big.Int, error) {
	return _NFTBridge.Contract.DestTxFee(&_NFTBridge.CallOpts, arg0)
}

// DestTxFee is a free data retrieval call binding the contract method 0x151ff4eb.
//
// Solidity: function destTxFee(uint64 ) view returns(uint256)
func (_NFTBridge *NFTBridgeCallerSession) DestTxFee(arg0 uint64) (*big.Int, error) {
	return _NFTBridge.Contract.DestTxFee(&_NFTBridge.CallOpts, arg0)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_NFTBridge *NFTBridgeCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_NFTBridge *NFTBridgeSession) MessageBus() (common.Address, error) {
	return _NFTBridge.Contract.MessageBus(&_NFTBridge.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_NFTBridge *NFTBridgeCallerSession) MessageBus() (common.Address, error) {
	return _NFTBridge.Contract.MessageBus(&_NFTBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTBridge *NFTBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTBridge *NFTBridgeSession) Owner() (common.Address, error) {
	return _NFTBridge.Contract.Owner(&_NFTBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTBridge *NFTBridgeCallerSession) Owner() (common.Address, error) {
	return _NFTBridge.Contract.Owner(&_NFTBridge.CallOpts)
}

// TotalFee is a free data retrieval call binding the contract method 0x9c1a65bd.
//
// Solidity: function totalFee(uint64 _dstChid, address _nft, uint256 _id) view returns(uint256)
func (_NFTBridge *NFTBridgeCaller) TotalFee(opts *bind.CallOpts, _dstChid uint64, _nft common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "totalFee", _dstChid, _nft, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFee is a free data retrieval call binding the contract method 0x9c1a65bd.
//
// Solidity: function totalFee(uint64 _dstChid, address _nft, uint256 _id) view returns(uint256)
func (_NFTBridge *NFTBridgeSession) TotalFee(_dstChid uint64, _nft common.Address, _id *big.Int) (*big.Int, error) {
	return _NFTBridge.Contract.TotalFee(&_NFTBridge.CallOpts, _dstChid, _nft, _id)
}

// TotalFee is a free data retrieval call binding the contract method 0x9c1a65bd.
//
// Solidity: function totalFee(uint64 _dstChid, address _nft, uint256 _id) view returns(uint256)
func (_NFTBridge *NFTBridgeCallerSession) TotalFee(_dstChid uint64, _nft common.Address, _id *big.Int) (*big.Int, error) {
	return _NFTBridge.Contract.TotalFee(&_NFTBridge.CallOpts, _dstChid, _nft, _id)
}

// Burn is a paid mutator transaction binding the contract method 0x46622643.
//
// Solidity: function burn(address _nft, uint256 _id, uint64 _dstChid, address _receiver, address _dstNft, address _dstBridge, bool _backToOrigin) payable returns()
func (_NFTBridge *NFTBridgeTransactor) Burn(opts *bind.TransactOpts, _nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address, _dstNft common.Address, _dstBridge common.Address, _backToOrigin bool) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "burn", _nft, _id, _dstChid, _receiver, _dstNft, _dstBridge, _backToOrigin)
}

// Burn is a paid mutator transaction binding the contract method 0x46622643.
//
// Solidity: function burn(address _nft, uint256 _id, uint64 _dstChid, address _receiver, address _dstNft, address _dstBridge, bool _backToOrigin) payable returns()
func (_NFTBridge *NFTBridgeSession) Burn(_nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address, _dstNft common.Address, _dstBridge common.Address, _backToOrigin bool) (*types.Transaction, error) {
	return _NFTBridge.Contract.Burn(&_NFTBridge.TransactOpts, _nft, _id, _dstChid, _receiver, _dstNft, _dstBridge, _backToOrigin)
}

// Burn is a paid mutator transaction binding the contract method 0x46622643.
//
// Solidity: function burn(address _nft, uint256 _id, uint64 _dstChid, address _receiver, address _dstNft, address _dstBridge, bool _backToOrigin) payable returns()
func (_NFTBridge *NFTBridgeTransactorSession) Burn(_nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address, _dstNft common.Address, _dstBridge common.Address, _backToOrigin bool) (*types.Transaction, error) {
	return _NFTBridge.Contract.Burn(&_NFTBridge.TransactOpts, _nft, _id, _dstChid, _receiver, _dstNft, _dstBridge, _backToOrigin)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_NFTBridge *NFTBridgeTransactor) ClaimFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "claimFee")
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_NFTBridge *NFTBridgeSession) ClaimFee() (*types.Transaction, error) {
	return _NFTBridge.Contract.ClaimFee(&_NFTBridge.TransactOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_NFTBridge *NFTBridgeTransactorSession) ClaimFee() (*types.Transaction, error) {
	return _NFTBridge.Contract.ClaimFee(&_NFTBridge.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x31aba041.
//
// Solidity: function deposit(address _nft, uint256 _id, uint64 _dstChid, address _receiver, address _dstNft, address _dstBridge) payable returns()
func (_NFTBridge *NFTBridgeTransactor) Deposit(opts *bind.TransactOpts, _nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address, _dstNft common.Address, _dstBridge common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "deposit", _nft, _id, _dstChid, _receiver, _dstNft, _dstBridge)
}

// Deposit is a paid mutator transaction binding the contract method 0x31aba041.
//
// Solidity: function deposit(address _nft, uint256 _id, uint64 _dstChid, address _receiver, address _dstNft, address _dstBridge) payable returns()
func (_NFTBridge *NFTBridgeSession) Deposit(_nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address, _dstNft common.Address, _dstBridge common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.Deposit(&_NFTBridge.TransactOpts, _nft, _id, _dstChid, _receiver, _dstNft, _dstBridge)
}

// Deposit is a paid mutator transaction binding the contract method 0x31aba041.
//
// Solidity: function deposit(address _nft, uint256 _id, uint64 _dstChid, address _receiver, address _dstNft, address _dstBridge) payable returns()
func (_NFTBridge *NFTBridgeTransactorSession) Deposit(_nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address, _dstNft common.Address, _dstBridge common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.Deposit(&_NFTBridge.TransactOpts, _nft, _id, _dstChid, _receiver, _dstNft, _dstBridge)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 srcChid, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeTransactor) ExecuteMessage(opts *bind.TransactOpts, arg0 common.Address, srcChid uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "executeMessage", arg0, srcChid, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 srcChid, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeSession) ExecuteMessage(arg0 common.Address, srcChid uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessage(&_NFTBridge.TransactOpts, arg0, srcChid, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 srcChid, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeTransactorSession) ExecuteMessage(arg0 common.Address, srcChid uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessage(&_NFTBridge.TransactOpts, arg0, srcChid, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransfer(&_NFTBridge.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransfer(&_NFTBridge.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransferFallback(&_NFTBridge.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransferFallback(&_NFTBridge.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransferRefund(&_NFTBridge.TransactOpts, _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_NFTBridge *NFTBridgeTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransferRefund(&_NFTBridge.TransactOpts, _token, _amount, _message)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_NFTBridge *NFTBridgeTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_NFTBridge *NFTBridgeSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetMessageBus(&_NFTBridge.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_NFTBridge *NFTBridgeTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetMessageBus(&_NFTBridge.TransactOpts, _messageBus)
}

// SetTxFee is a paid mutator transaction binding the contract method 0x9f825f07.
//
// Solidity: function setTxFee(uint64 chid, uint256 fee) returns()
func (_NFTBridge *NFTBridgeTransactor) SetTxFee(opts *bind.TransactOpts, chid uint64, fee *big.Int) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "setTxFee", chid, fee)
}

// SetTxFee is a paid mutator transaction binding the contract method 0x9f825f07.
//
// Solidity: function setTxFee(uint64 chid, uint256 fee) returns()
func (_NFTBridge *NFTBridgeSession) SetTxFee(chid uint64, fee *big.Int) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetTxFee(&_NFTBridge.TransactOpts, chid, fee)
}

// SetTxFee is a paid mutator transaction binding the contract method 0x9f825f07.
//
// Solidity: function setTxFee(uint64 chid, uint256 fee) returns()
func (_NFTBridge *NFTBridgeTransactorSession) SetTxFee(chid uint64, fee *big.Int) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetTxFee(&_NFTBridge.TransactOpts, chid, fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTBridge *NFTBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTBridge *NFTBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.TransferOwnership(&_NFTBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTBridge *NFTBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.TransferOwnership(&_NFTBridge.TransactOpts, newOwner)
}

// NFTBridgeMessageBusUpdatedIterator is returned from FilterMessageBusUpdated and is used to iterate over the raw logs and unpacked data for MessageBusUpdated events raised by the NFTBridge contract.
type NFTBridgeMessageBusUpdatedIterator struct {
	Event *NFTBridgeMessageBusUpdated // Event containing the contract specifics and raw log

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
func (it *NFTBridgeMessageBusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeMessageBusUpdated)
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
		it.Event = new(NFTBridgeMessageBusUpdated)
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
func (it *NFTBridgeMessageBusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeMessageBusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeMessageBusUpdated represents a MessageBusUpdated event raised by the NFTBridge contract.
type NFTBridgeMessageBusUpdated struct {
	MessageBus common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageBusUpdated is a free log retrieval operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_NFTBridge *NFTBridgeFilterer) FilterMessageBusUpdated(opts *bind.FilterOpts) (*NFTBridgeMessageBusUpdatedIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "MessageBusUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeMessageBusUpdatedIterator{contract: _NFTBridge.contract, event: "MessageBusUpdated", logs: logs, sub: sub}, nil
}

// WatchMessageBusUpdated is a free log subscription operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_NFTBridge *NFTBridgeFilterer) WatchMessageBusUpdated(opts *bind.WatchOpts, sink chan<- *NFTBridgeMessageBusUpdated) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "MessageBusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeMessageBusUpdated)
				if err := _NFTBridge.contract.UnpackLog(event, "MessageBusUpdated", log); err != nil {
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

// ParseMessageBusUpdated is a log parse operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_NFTBridge *NFTBridgeFilterer) ParseMessageBusUpdated(log types.Log) (*NFTBridgeMessageBusUpdated, error) {
	event := new(NFTBridgeMessageBusUpdated)
	if err := _NFTBridge.contract.UnpackLog(event, "MessageBusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTBridge contract.
type NFTBridgeOwnershipTransferredIterator struct {
	Event *NFTBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NFTBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeOwnershipTransferred)
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
		it.Event = new(NFTBridgeOwnershipTransferred)
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
func (it *NFTBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the NFTBridge contract.
type NFTBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTBridge *NFTBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTBridgeOwnershipTransferredIterator{contract: _NFTBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTBridge *NFTBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeOwnershipTransferred)
				if err := _NFTBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NFTBridge *NFTBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*NFTBridgeOwnershipTransferred, error) {
	event := new(NFTBridgeOwnershipTransferred)
	if err := _NFTBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgeReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the NFTBridge contract.
type NFTBridgeReceivedIterator struct {
	Event *NFTBridgeReceived // Event containing the contract specifics and raw log

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
func (it *NFTBridgeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeReceived)
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
		it.Event = new(NFTBridgeReceived)
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
func (it *NFTBridgeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeReceived represents a Received event raised by the NFTBridge contract.
type NFTBridgeReceived struct {
	Receiver common.Address
	Nft      common.Address
	Id       *big.Int
	SrcChid  uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x0aac355db06d21352d6b898d8e0ae1334d55f65b6c4c09e26951166a8eb4dba7.
//
// Solidity: event Received(address receiver, address nft, uint256 id, uint64 srcChid)
func (_NFTBridge *NFTBridgeFilterer) FilterReceived(opts *bind.FilterOpts) (*NFTBridgeReceivedIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeReceivedIterator{contract: _NFTBridge.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x0aac355db06d21352d6b898d8e0ae1334d55f65b6c4c09e26951166a8eb4dba7.
//
// Solidity: event Received(address receiver, address nft, uint256 id, uint64 srcChid)
func (_NFTBridge *NFTBridgeFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *NFTBridgeReceived) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeReceived)
				if err := _NFTBridge.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x0aac355db06d21352d6b898d8e0ae1334d55f65b6c4c09e26951166a8eb4dba7.
//
// Solidity: event Received(address receiver, address nft, uint256 id, uint64 srcChid)
func (_NFTBridge *NFTBridgeFilterer) ParseReceived(log types.Log) (*NFTBridgeReceived, error) {
	event := new(NFTBridgeReceived)
	if err := _NFTBridge.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgeSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the NFTBridge contract.
type NFTBridgeSentIterator struct {
	Event *NFTBridgeSent // Event containing the contract specifics and raw log

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
func (it *NFTBridgeSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeSent)
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
		it.Event = new(NFTBridgeSent)
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
func (it *NFTBridgeSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeSent represents a Sent event raised by the NFTBridge contract.
type NFTBridgeSent struct {
	Sender   common.Address
	SrcNft   common.Address
	Id       *big.Int
	DstChid  uint64
	Receiver common.Address
	DstNft   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x40143e5b72b2109d658cfa709dec6213f60364dbd08b7253cdaf5f4e0c49561c.
//
// Solidity: event Sent(address sender, address srcNft, uint256 id, uint64 dstChid, address receiver, address dstNft)
func (_NFTBridge *NFTBridgeFilterer) FilterSent(opts *bind.FilterOpts) (*NFTBridgeSentIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeSentIterator{contract: _NFTBridge.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x40143e5b72b2109d658cfa709dec6213f60364dbd08b7253cdaf5f4e0c49561c.
//
// Solidity: event Sent(address sender, address srcNft, uint256 id, uint64 dstChid, address receiver, address dstNft)
func (_NFTBridge *NFTBridgeFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *NFTBridgeSent) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeSent)
				if err := _NFTBridge.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x40143e5b72b2109d658cfa709dec6213f60364dbd08b7253cdaf5f4e0c49561c.
//
// Solidity: event Sent(address sender, address srcNft, uint256 id, uint64 dstChid, address receiver, address dstNft)
func (_NFTBridge *NFTBridgeFilterer) ParseSent(log types.Log) (*NFTBridgeSent, error) {
	event := new(NFTBridgeSent)
	if err := _NFTBridge.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
