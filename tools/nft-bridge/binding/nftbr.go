// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"}],\"name\":\"MessageBusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChid\",\"type\":\"uint64\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcNft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstNft\",\"type\":\"address\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstNftBridge\",\"type\":\"address\"}],\"name\":\"SetDestBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcNft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstNft\",\"type\":\"address\"}],\"name\":\"SetDestNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrig\",\"type\":\"bool\"}],\"name\":\"SetOrigNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SetTxFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"}],\"name\":\"delOrigNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"destBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"destNFTAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"destTxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChid\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgBus\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"origNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_dstChid\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"sendMsg\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChid\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"sendTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChid\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstNftBridge\",\"type\":\"address\"}],\"name\":\"setDestBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"dstChid\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"dstNftBridge\",\"type\":\"address[]\"}],\"name\":\"setDestBridges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcNft\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"dstChid\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstNft\",\"type\":\"address\"}],\"name\":\"setDestNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcNft\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"dstChid\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"dstNft\",\"type\":\"address[]\"}],\"name\":\"setDestNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"}],\"name\":\"setOrigNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chid\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setTxFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_dstChid\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DestBridge is a free data retrieval call binding the contract method 0x27cbe705.
//
// Solidity: function destBridge(uint64 ) view returns(address)
func (_NFTBridge *NFTBridgeCaller) DestBridge(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "destBridge", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DestBridge is a free data retrieval call binding the contract method 0x27cbe705.
//
// Solidity: function destBridge(uint64 ) view returns(address)
func (_NFTBridge *NFTBridgeSession) DestBridge(arg0 uint64) (common.Address, error) {
	return _NFTBridge.Contract.DestBridge(&_NFTBridge.CallOpts, arg0)
}

// DestBridge is a free data retrieval call binding the contract method 0x27cbe705.
//
// Solidity: function destBridge(uint64 ) view returns(address)
func (_NFTBridge *NFTBridgeCallerSession) DestBridge(arg0 uint64) (common.Address, error) {
	return _NFTBridge.Contract.DestBridge(&_NFTBridge.CallOpts, arg0)
}

// DestNFTAddr is a free data retrieval call binding the contract method 0x9e041b9d.
//
// Solidity: function destNFTAddr(address , uint64 ) view returns(address)
func (_NFTBridge *NFTBridgeCaller) DestNFTAddr(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (common.Address, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "destNFTAddr", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DestNFTAddr is a free data retrieval call binding the contract method 0x9e041b9d.
//
// Solidity: function destNFTAddr(address , uint64 ) view returns(address)
func (_NFTBridge *NFTBridgeSession) DestNFTAddr(arg0 common.Address, arg1 uint64) (common.Address, error) {
	return _NFTBridge.Contract.DestNFTAddr(&_NFTBridge.CallOpts, arg0, arg1)
}

// DestNFTAddr is a free data retrieval call binding the contract method 0x9e041b9d.
//
// Solidity: function destNFTAddr(address , uint64 ) view returns(address)
func (_NFTBridge *NFTBridgeCallerSession) DestNFTAddr(arg0 common.Address, arg1 uint64) (common.Address, error) {
	return _NFTBridge.Contract.DestNFTAddr(&_NFTBridge.CallOpts, arg0, arg1)
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

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_NFTBridge *NFTBridgeCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_NFTBridge *NFTBridgeSession) IsPauser(account common.Address) (bool, error) {
	return _NFTBridge.Contract.IsPauser(&_NFTBridge.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_NFTBridge *NFTBridgeCallerSession) IsPauser(account common.Address) (bool, error) {
	return _NFTBridge.Contract.IsPauser(&_NFTBridge.CallOpts, account)
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

// OrigNFT is a free data retrieval call binding the contract method 0x1140c84e.
//
// Solidity: function origNFT(address ) view returns(bool)
func (_NFTBridge *NFTBridgeCaller) OrigNFT(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "origNFT", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrigNFT is a free data retrieval call binding the contract method 0x1140c84e.
//
// Solidity: function origNFT(address ) view returns(bool)
func (_NFTBridge *NFTBridgeSession) OrigNFT(arg0 common.Address) (bool, error) {
	return _NFTBridge.Contract.OrigNFT(&_NFTBridge.CallOpts, arg0)
}

// OrigNFT is a free data retrieval call binding the contract method 0x1140c84e.
//
// Solidity: function origNFT(address ) view returns(bool)
func (_NFTBridge *NFTBridgeCallerSession) OrigNFT(arg0 common.Address) (bool, error) {
	return _NFTBridge.Contract.OrigNFT(&_NFTBridge.CallOpts, arg0)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFTBridge *NFTBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFTBridge *NFTBridgeSession) Paused() (bool, error) {
	return _NFTBridge.Contract.Paused(&_NFTBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFTBridge *NFTBridgeCallerSession) Paused() (bool, error) {
	return _NFTBridge.Contract.Paused(&_NFTBridge.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_NFTBridge *NFTBridgeCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NFTBridge.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_NFTBridge *NFTBridgeSession) Pausers(arg0 common.Address) (bool, error) {
	return _NFTBridge.Contract.Pausers(&_NFTBridge.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_NFTBridge *NFTBridgeCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _NFTBridge.Contract.Pausers(&_NFTBridge.CallOpts, arg0)
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

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_NFTBridge *NFTBridgeTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_NFTBridge *NFTBridgeSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.AddPauser(&_NFTBridge.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_NFTBridge *NFTBridgeTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.AddPauser(&_NFTBridge.TransactOpts, account)
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

// DelOrigNFT is a paid mutator transaction binding the contract method 0xed99bf43.
//
// Solidity: function delOrigNFT(address _nft) returns()
func (_NFTBridge *NFTBridgeTransactor) DelOrigNFT(opts *bind.TransactOpts, _nft common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "delOrigNFT", _nft)
}

// DelOrigNFT is a paid mutator transaction binding the contract method 0xed99bf43.
//
// Solidity: function delOrigNFT(address _nft) returns()
func (_NFTBridge *NFTBridgeSession) DelOrigNFT(_nft common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.DelOrigNFT(&_NFTBridge.TransactOpts, _nft)
}

// DelOrigNFT is a paid mutator transaction binding the contract method 0xed99bf43.
//
// Solidity: function delOrigNFT(address _nft) returns()
func (_NFTBridge *NFTBridgeTransactorSession) DelOrigNFT(_nft common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.DelOrigNFT(&_NFTBridge.TransactOpts, _nft)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address sender, uint64 srcChid, bytes _message, address ) payable returns(uint8)
func (_NFTBridge *NFTBridgeTransactor) ExecuteMessage(opts *bind.TransactOpts, sender common.Address, srcChid uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "executeMessage", sender, srcChid, _message, arg3)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address sender, uint64 srcChid, bytes _message, address ) payable returns(uint8)
func (_NFTBridge *NFTBridgeSession) ExecuteMessage(sender common.Address, srcChid uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessage(&_NFTBridge.TransactOpts, sender, srcChid, _message, arg3)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address sender, uint64 srcChid, bytes _message, address ) payable returns(uint8)
func (_NFTBridge *NFTBridgeTransactorSession) ExecuteMessage(sender common.Address, srcChid uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessage(&_NFTBridge.TransactOpts, sender, srcChid, _message, arg3)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransfer(&_NFTBridge.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransfer(&_NFTBridge.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransferFallback(&_NFTBridge.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransferFallback(&_NFTBridge.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransferRefund(&_NFTBridge.TransactOpts, _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_NFTBridge *NFTBridgeTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.ExecuteMessageWithTransferRefund(&_NFTBridge.TransactOpts, _token, _amount, _message, _executor)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _msgBus) returns()
func (_NFTBridge *NFTBridgeTransactor) Init(opts *bind.TransactOpts, _msgBus common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "init", _msgBus)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _msgBus) returns()
func (_NFTBridge *NFTBridgeSession) Init(_msgBus common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.Init(&_NFTBridge.TransactOpts, _msgBus)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _msgBus) returns()
func (_NFTBridge *NFTBridgeTransactorSession) Init(_msgBus common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.Init(&_NFTBridge.TransactOpts, _msgBus)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFTBridge *NFTBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFTBridge *NFTBridgeSession) Pause() (*types.Transaction, error) {
	return _NFTBridge.Contract.Pause(&_NFTBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFTBridge *NFTBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _NFTBridge.Contract.Pause(&_NFTBridge.TransactOpts)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_NFTBridge *NFTBridgeTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_NFTBridge *NFTBridgeSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.RemovePauser(&_NFTBridge.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_NFTBridge *NFTBridgeTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.RemovePauser(&_NFTBridge.TransactOpts, account)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_NFTBridge *NFTBridgeTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_NFTBridge *NFTBridgeSession) RenouncePauser() (*types.Transaction, error) {
	return _NFTBridge.Contract.RenouncePauser(&_NFTBridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_NFTBridge *NFTBridgeTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _NFTBridge.Contract.RenouncePauser(&_NFTBridge.TransactOpts)
}

// SendMsg is a paid mutator transaction binding the contract method 0xf0cb57ce.
//
// Solidity: function sendMsg(uint64 _dstChid, address _sender, address _receiver, uint256 _id, string _uri) payable returns()
func (_NFTBridge *NFTBridgeTransactor) SendMsg(opts *bind.TransactOpts, _dstChid uint64, _sender common.Address, _receiver common.Address, _id *big.Int, _uri string) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "sendMsg", _dstChid, _sender, _receiver, _id, _uri)
}

// SendMsg is a paid mutator transaction binding the contract method 0xf0cb57ce.
//
// Solidity: function sendMsg(uint64 _dstChid, address _sender, address _receiver, uint256 _id, string _uri) payable returns()
func (_NFTBridge *NFTBridgeSession) SendMsg(_dstChid uint64, _sender common.Address, _receiver common.Address, _id *big.Int, _uri string) (*types.Transaction, error) {
	return _NFTBridge.Contract.SendMsg(&_NFTBridge.TransactOpts, _dstChid, _sender, _receiver, _id, _uri)
}

// SendMsg is a paid mutator transaction binding the contract method 0xf0cb57ce.
//
// Solidity: function sendMsg(uint64 _dstChid, address _sender, address _receiver, uint256 _id, string _uri) payable returns()
func (_NFTBridge *NFTBridgeTransactorSession) SendMsg(_dstChid uint64, _sender common.Address, _receiver common.Address, _id *big.Int, _uri string) (*types.Transaction, error) {
	return _NFTBridge.Contract.SendMsg(&_NFTBridge.TransactOpts, _dstChid, _sender, _receiver, _id, _uri)
}

// SendTo is a paid mutator transaction binding the contract method 0xb2c88775.
//
// Solidity: function sendTo(address _nft, uint256 _id, uint64 _dstChid, address _receiver) payable returns()
func (_NFTBridge *NFTBridgeTransactor) SendTo(opts *bind.TransactOpts, _nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "sendTo", _nft, _id, _dstChid, _receiver)
}

// SendTo is a paid mutator transaction binding the contract method 0xb2c88775.
//
// Solidity: function sendTo(address _nft, uint256 _id, uint64 _dstChid, address _receiver) payable returns()
func (_NFTBridge *NFTBridgeSession) SendTo(_nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SendTo(&_NFTBridge.TransactOpts, _nft, _id, _dstChid, _receiver)
}

// SendTo is a paid mutator transaction binding the contract method 0xb2c88775.
//
// Solidity: function sendTo(address _nft, uint256 _id, uint64 _dstChid, address _receiver) payable returns()
func (_NFTBridge *NFTBridgeTransactorSession) SendTo(_nft common.Address, _id *big.Int, _dstChid uint64, _receiver common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SendTo(&_NFTBridge.TransactOpts, _nft, _id, _dstChid, _receiver)
}

// SetDestBridge is a paid mutator transaction binding the contract method 0xda754f44.
//
// Solidity: function setDestBridge(uint64 dstChid, address dstNftBridge) returns()
func (_NFTBridge *NFTBridgeTransactor) SetDestBridge(opts *bind.TransactOpts, dstChid uint64, dstNftBridge common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "setDestBridge", dstChid, dstNftBridge)
}

// SetDestBridge is a paid mutator transaction binding the contract method 0xda754f44.
//
// Solidity: function setDestBridge(uint64 dstChid, address dstNftBridge) returns()
func (_NFTBridge *NFTBridgeSession) SetDestBridge(dstChid uint64, dstNftBridge common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetDestBridge(&_NFTBridge.TransactOpts, dstChid, dstNftBridge)
}

// SetDestBridge is a paid mutator transaction binding the contract method 0xda754f44.
//
// Solidity: function setDestBridge(uint64 dstChid, address dstNftBridge) returns()
func (_NFTBridge *NFTBridgeTransactorSession) SetDestBridge(dstChid uint64, dstNftBridge common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetDestBridge(&_NFTBridge.TransactOpts, dstChid, dstNftBridge)
}

// SetDestBridges is a paid mutator transaction binding the contract method 0x43a8c137.
//
// Solidity: function setDestBridges(uint64[] dstChid, address[] dstNftBridge) returns()
func (_NFTBridge *NFTBridgeTransactor) SetDestBridges(opts *bind.TransactOpts, dstChid []uint64, dstNftBridge []common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "setDestBridges", dstChid, dstNftBridge)
}

// SetDestBridges is a paid mutator transaction binding the contract method 0x43a8c137.
//
// Solidity: function setDestBridges(uint64[] dstChid, address[] dstNftBridge) returns()
func (_NFTBridge *NFTBridgeSession) SetDestBridges(dstChid []uint64, dstNftBridge []common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetDestBridges(&_NFTBridge.TransactOpts, dstChid, dstNftBridge)
}

// SetDestBridges is a paid mutator transaction binding the contract method 0x43a8c137.
//
// Solidity: function setDestBridges(uint64[] dstChid, address[] dstNftBridge) returns()
func (_NFTBridge *NFTBridgeTransactorSession) SetDestBridges(dstChid []uint64, dstNftBridge []common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetDestBridges(&_NFTBridge.TransactOpts, dstChid, dstNftBridge)
}

// SetDestNFT is a paid mutator transaction binding the contract method 0x95b6a191.
//
// Solidity: function setDestNFT(address srcNft, uint64 dstChid, address dstNft) returns()
func (_NFTBridge *NFTBridgeTransactor) SetDestNFT(opts *bind.TransactOpts, srcNft common.Address, dstChid uint64, dstNft common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "setDestNFT", srcNft, dstChid, dstNft)
}

// SetDestNFT is a paid mutator transaction binding the contract method 0x95b6a191.
//
// Solidity: function setDestNFT(address srcNft, uint64 dstChid, address dstNft) returns()
func (_NFTBridge *NFTBridgeSession) SetDestNFT(srcNft common.Address, dstChid uint64, dstNft common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetDestNFT(&_NFTBridge.TransactOpts, srcNft, dstChid, dstNft)
}

// SetDestNFT is a paid mutator transaction binding the contract method 0x95b6a191.
//
// Solidity: function setDestNFT(address srcNft, uint64 dstChid, address dstNft) returns()
func (_NFTBridge *NFTBridgeTransactorSession) SetDestNFT(srcNft common.Address, dstChid uint64, dstNft common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetDestNFT(&_NFTBridge.TransactOpts, srcNft, dstChid, dstNft)
}

// SetDestNFTs is a paid mutator transaction binding the contract method 0x2d5fa47c.
//
// Solidity: function setDestNFTs(address srcNft, uint64[] dstChid, address[] dstNft) returns()
func (_NFTBridge *NFTBridgeTransactor) SetDestNFTs(opts *bind.TransactOpts, srcNft common.Address, dstChid []uint64, dstNft []common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "setDestNFTs", srcNft, dstChid, dstNft)
}

// SetDestNFTs is a paid mutator transaction binding the contract method 0x2d5fa47c.
//
// Solidity: function setDestNFTs(address srcNft, uint64[] dstChid, address[] dstNft) returns()
func (_NFTBridge *NFTBridgeSession) SetDestNFTs(srcNft common.Address, dstChid []uint64, dstNft []common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetDestNFTs(&_NFTBridge.TransactOpts, srcNft, dstChid, dstNft)
}

// SetDestNFTs is a paid mutator transaction binding the contract method 0x2d5fa47c.
//
// Solidity: function setDestNFTs(address srcNft, uint64[] dstChid, address[] dstNft) returns()
func (_NFTBridge *NFTBridgeTransactorSession) SetDestNFTs(srcNft common.Address, dstChid []uint64, dstNft []common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetDestNFTs(&_NFTBridge.TransactOpts, srcNft, dstChid, dstNft)
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

// SetOrigNFT is a paid mutator transaction binding the contract method 0x065c38fd.
//
// Solidity: function setOrigNFT(address _nft) returns()
func (_NFTBridge *NFTBridgeTransactor) SetOrigNFT(opts *bind.TransactOpts, _nft common.Address) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "setOrigNFT", _nft)
}

// SetOrigNFT is a paid mutator transaction binding the contract method 0x065c38fd.
//
// Solidity: function setOrigNFT(address _nft) returns()
func (_NFTBridge *NFTBridgeSession) SetOrigNFT(_nft common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetOrigNFT(&_NFTBridge.TransactOpts, _nft)
}

// SetOrigNFT is a paid mutator transaction binding the contract method 0x065c38fd.
//
// Solidity: function setOrigNFT(address _nft) returns()
func (_NFTBridge *NFTBridgeTransactorSession) SetOrigNFT(_nft common.Address) (*types.Transaction, error) {
	return _NFTBridge.Contract.SetOrigNFT(&_NFTBridge.TransactOpts, _nft)
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

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFTBridge *NFTBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFTBridge *NFTBridgeSession) Unpause() (*types.Transaction, error) {
	return _NFTBridge.Contract.Unpause(&_NFTBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFTBridge *NFTBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _NFTBridge.Contract.Unpause(&_NFTBridge.TransactOpts)
}

// NFTBridgeFeeClaimedIterator is returned from FilterFeeClaimed and is used to iterate over the raw logs and unpacked data for FeeClaimed events raised by the NFTBridge contract.
type NFTBridgeFeeClaimedIterator struct {
	Event *NFTBridgeFeeClaimed // Event containing the contract specifics and raw log

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
func (it *NFTBridgeFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeFeeClaimed)
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
		it.Event = new(NFTBridgeFeeClaimed)
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
func (it *NFTBridgeFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeFeeClaimed represents a FeeClaimed event raised by the NFTBridge contract.
type NFTBridgeFeeClaimed struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeClaimed is a free log retrieval operation binding the contract event 0x62b10e3ff3d45b5ff546e740b893897facb1680285f989a64ae932d62c5388e1.
//
// Solidity: event FeeClaimed(uint256 amount)
func (_NFTBridge *NFTBridgeFilterer) FilterFeeClaimed(opts *bind.FilterOpts) (*NFTBridgeFeeClaimedIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "FeeClaimed")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeFeeClaimedIterator{contract: _NFTBridge.contract, event: "FeeClaimed", logs: logs, sub: sub}, nil
}

// WatchFeeClaimed is a free log subscription operation binding the contract event 0x62b10e3ff3d45b5ff546e740b893897facb1680285f989a64ae932d62c5388e1.
//
// Solidity: event FeeClaimed(uint256 amount)
func (_NFTBridge *NFTBridgeFilterer) WatchFeeClaimed(opts *bind.WatchOpts, sink chan<- *NFTBridgeFeeClaimed) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "FeeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeFeeClaimed)
				if err := _NFTBridge.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
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

// ParseFeeClaimed is a log parse operation binding the contract event 0x62b10e3ff3d45b5ff546e740b893897facb1680285f989a64ae932d62c5388e1.
//
// Solidity: event FeeClaimed(uint256 amount)
func (_NFTBridge *NFTBridgeFilterer) ParseFeeClaimed(log types.Log) (*NFTBridgeFeeClaimed, error) {
	event := new(NFTBridgeFeeClaimed)
	if err := _NFTBridge.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// NFTBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NFTBridge contract.
type NFTBridgePausedIterator struct {
	Event *NFTBridgePaused // Event containing the contract specifics and raw log

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
func (it *NFTBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgePaused)
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
		it.Event = new(NFTBridgePaused)
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
func (it *NFTBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgePaused represents a Paused event raised by the NFTBridge contract.
type NFTBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NFTBridge *NFTBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*NFTBridgePausedIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NFTBridgePausedIterator{contract: _NFTBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NFTBridge *NFTBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NFTBridgePaused) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgePaused)
				if err := _NFTBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_NFTBridge *NFTBridgeFilterer) ParsePaused(log types.Log) (*NFTBridgePaused, error) {
	event := new(NFTBridgePaused)
	if err := _NFTBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the NFTBridge contract.
type NFTBridgePauserAddedIterator struct {
	Event *NFTBridgePauserAdded // Event containing the contract specifics and raw log

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
func (it *NFTBridgePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgePauserAdded)
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
		it.Event = new(NFTBridgePauserAdded)
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
func (it *NFTBridgePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgePauserAdded represents a PauserAdded event raised by the NFTBridge contract.
type NFTBridgePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_NFTBridge *NFTBridgeFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*NFTBridgePauserAddedIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &NFTBridgePauserAddedIterator{contract: _NFTBridge.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_NFTBridge *NFTBridgeFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *NFTBridgePauserAdded) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgePauserAdded)
				if err := _NFTBridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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
func (_NFTBridge *NFTBridgeFilterer) ParsePauserAdded(log types.Log) (*NFTBridgePauserAdded, error) {
	event := new(NFTBridgePauserAdded)
	if err := _NFTBridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the NFTBridge contract.
type NFTBridgePauserRemovedIterator struct {
	Event *NFTBridgePauserRemoved // Event containing the contract specifics and raw log

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
func (it *NFTBridgePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgePauserRemoved)
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
		it.Event = new(NFTBridgePauserRemoved)
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
func (it *NFTBridgePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgePauserRemoved represents a PauserRemoved event raised by the NFTBridge contract.
type NFTBridgePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_NFTBridge *NFTBridgeFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*NFTBridgePauserRemovedIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &NFTBridgePauserRemovedIterator{contract: _NFTBridge.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_NFTBridge *NFTBridgeFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *NFTBridgePauserRemoved) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgePauserRemoved)
				if err := _NFTBridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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
func (_NFTBridge *NFTBridgeFilterer) ParsePauserRemoved(log types.Log) (*NFTBridgePauserRemoved, error) {
	event := new(NFTBridgePauserRemoved)
	if err := _NFTBridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// NFTBridgeSetDestBridgeIterator is returned from FilterSetDestBridge and is used to iterate over the raw logs and unpacked data for SetDestBridge events raised by the NFTBridge contract.
type NFTBridgeSetDestBridgeIterator struct {
	Event *NFTBridgeSetDestBridge // Event containing the contract specifics and raw log

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
func (it *NFTBridgeSetDestBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeSetDestBridge)
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
		it.Event = new(NFTBridgeSetDestBridge)
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
func (it *NFTBridgeSetDestBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeSetDestBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeSetDestBridge represents a SetDestBridge event raised by the NFTBridge contract.
type NFTBridgeSetDestBridge struct {
	DstChid      uint64
	DstNftBridge common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetDestBridge is a free log retrieval operation binding the contract event 0x3e776334b24c927645043308f89ac1ca734002e5a921ff384a70dcbb88c92cd4.
//
// Solidity: event SetDestBridge(uint64 dstChid, address dstNftBridge)
func (_NFTBridge *NFTBridgeFilterer) FilterSetDestBridge(opts *bind.FilterOpts) (*NFTBridgeSetDestBridgeIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "SetDestBridge")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeSetDestBridgeIterator{contract: _NFTBridge.contract, event: "SetDestBridge", logs: logs, sub: sub}, nil
}

// WatchSetDestBridge is a free log subscription operation binding the contract event 0x3e776334b24c927645043308f89ac1ca734002e5a921ff384a70dcbb88c92cd4.
//
// Solidity: event SetDestBridge(uint64 dstChid, address dstNftBridge)
func (_NFTBridge *NFTBridgeFilterer) WatchSetDestBridge(opts *bind.WatchOpts, sink chan<- *NFTBridgeSetDestBridge) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "SetDestBridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeSetDestBridge)
				if err := _NFTBridge.contract.UnpackLog(event, "SetDestBridge", log); err != nil {
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

// ParseSetDestBridge is a log parse operation binding the contract event 0x3e776334b24c927645043308f89ac1ca734002e5a921ff384a70dcbb88c92cd4.
//
// Solidity: event SetDestBridge(uint64 dstChid, address dstNftBridge)
func (_NFTBridge *NFTBridgeFilterer) ParseSetDestBridge(log types.Log) (*NFTBridgeSetDestBridge, error) {
	event := new(NFTBridgeSetDestBridge)
	if err := _NFTBridge.contract.UnpackLog(event, "SetDestBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgeSetDestNFTIterator is returned from FilterSetDestNFT and is used to iterate over the raw logs and unpacked data for SetDestNFT events raised by the NFTBridge contract.
type NFTBridgeSetDestNFTIterator struct {
	Event *NFTBridgeSetDestNFT // Event containing the contract specifics and raw log

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
func (it *NFTBridgeSetDestNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeSetDestNFT)
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
		it.Event = new(NFTBridgeSetDestNFT)
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
func (it *NFTBridgeSetDestNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeSetDestNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeSetDestNFT represents a SetDestNFT event raised by the NFTBridge contract.
type NFTBridgeSetDestNFT struct {
	SrcNft  common.Address
	DstChid uint64
	DstNft  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetDestNFT is a free log retrieval operation binding the contract event 0xa5a9b84f1b7eb437335ea919a3ff6de6e242e4733d0100a77391106173794871.
//
// Solidity: event SetDestNFT(address srcNft, uint64 dstChid, address dstNft)
func (_NFTBridge *NFTBridgeFilterer) FilterSetDestNFT(opts *bind.FilterOpts) (*NFTBridgeSetDestNFTIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "SetDestNFT")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeSetDestNFTIterator{contract: _NFTBridge.contract, event: "SetDestNFT", logs: logs, sub: sub}, nil
}

// WatchSetDestNFT is a free log subscription operation binding the contract event 0xa5a9b84f1b7eb437335ea919a3ff6de6e242e4733d0100a77391106173794871.
//
// Solidity: event SetDestNFT(address srcNft, uint64 dstChid, address dstNft)
func (_NFTBridge *NFTBridgeFilterer) WatchSetDestNFT(opts *bind.WatchOpts, sink chan<- *NFTBridgeSetDestNFT) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "SetDestNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeSetDestNFT)
				if err := _NFTBridge.contract.UnpackLog(event, "SetDestNFT", log); err != nil {
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

// ParseSetDestNFT is a log parse operation binding the contract event 0xa5a9b84f1b7eb437335ea919a3ff6de6e242e4733d0100a77391106173794871.
//
// Solidity: event SetDestNFT(address srcNft, uint64 dstChid, address dstNft)
func (_NFTBridge *NFTBridgeFilterer) ParseSetDestNFT(log types.Log) (*NFTBridgeSetDestNFT, error) {
	event := new(NFTBridgeSetDestNFT)
	if err := _NFTBridge.contract.UnpackLog(event, "SetDestNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgeSetOrigNFTIterator is returned from FilterSetOrigNFT and is used to iterate over the raw logs and unpacked data for SetOrigNFT events raised by the NFTBridge contract.
type NFTBridgeSetOrigNFTIterator struct {
	Event *NFTBridgeSetOrigNFT // Event containing the contract specifics and raw log

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
func (it *NFTBridgeSetOrigNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeSetOrigNFT)
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
		it.Event = new(NFTBridgeSetOrigNFT)
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
func (it *NFTBridgeSetOrigNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeSetOrigNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeSetOrigNFT represents a SetOrigNFT event raised by the NFTBridge contract.
type NFTBridgeSetOrigNFT struct {
	Nft    common.Address
	IsOrig bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOrigNFT is a free log retrieval operation binding the contract event 0x9800fb32bf5eb9a3b2e42c910912da10ed1881dc538475101797669146166bf8.
//
// Solidity: event SetOrigNFT(address nft, bool isOrig)
func (_NFTBridge *NFTBridgeFilterer) FilterSetOrigNFT(opts *bind.FilterOpts) (*NFTBridgeSetOrigNFTIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "SetOrigNFT")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeSetOrigNFTIterator{contract: _NFTBridge.contract, event: "SetOrigNFT", logs: logs, sub: sub}, nil
}

// WatchSetOrigNFT is a free log subscription operation binding the contract event 0x9800fb32bf5eb9a3b2e42c910912da10ed1881dc538475101797669146166bf8.
//
// Solidity: event SetOrigNFT(address nft, bool isOrig)
func (_NFTBridge *NFTBridgeFilterer) WatchSetOrigNFT(opts *bind.WatchOpts, sink chan<- *NFTBridgeSetOrigNFT) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "SetOrigNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeSetOrigNFT)
				if err := _NFTBridge.contract.UnpackLog(event, "SetOrigNFT", log); err != nil {
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

// ParseSetOrigNFT is a log parse operation binding the contract event 0x9800fb32bf5eb9a3b2e42c910912da10ed1881dc538475101797669146166bf8.
//
// Solidity: event SetOrigNFT(address nft, bool isOrig)
func (_NFTBridge *NFTBridgeFilterer) ParseSetOrigNFT(log types.Log) (*NFTBridgeSetOrigNFT, error) {
	event := new(NFTBridgeSetOrigNFT)
	if err := _NFTBridge.contract.UnpackLog(event, "SetOrigNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgeSetTxFeeIterator is returned from FilterSetTxFee and is used to iterate over the raw logs and unpacked data for SetTxFee events raised by the NFTBridge contract.
type NFTBridgeSetTxFeeIterator struct {
	Event *NFTBridgeSetTxFee // Event containing the contract specifics and raw log

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
func (it *NFTBridgeSetTxFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeSetTxFee)
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
		it.Event = new(NFTBridgeSetTxFee)
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
func (it *NFTBridgeSetTxFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeSetTxFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeSetTxFee represents a SetTxFee event raised by the NFTBridge contract.
type NFTBridgeSetTxFee struct {
	Chid uint64
	Fee  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetTxFee is a free log retrieval operation binding the contract event 0x446a287c2114fa54d0083d97ce8f6f15b2ce29fa1c2df4b5a580d581ea7c4ad3.
//
// Solidity: event SetTxFee(uint64 chid, uint256 fee)
func (_NFTBridge *NFTBridgeFilterer) FilterSetTxFee(opts *bind.FilterOpts) (*NFTBridgeSetTxFeeIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "SetTxFee")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeSetTxFeeIterator{contract: _NFTBridge.contract, event: "SetTxFee", logs: logs, sub: sub}, nil
}

// WatchSetTxFee is a free log subscription operation binding the contract event 0x446a287c2114fa54d0083d97ce8f6f15b2ce29fa1c2df4b5a580d581ea7c4ad3.
//
// Solidity: event SetTxFee(uint64 chid, uint256 fee)
func (_NFTBridge *NFTBridgeFilterer) WatchSetTxFee(opts *bind.WatchOpts, sink chan<- *NFTBridgeSetTxFee) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "SetTxFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeSetTxFee)
				if err := _NFTBridge.contract.UnpackLog(event, "SetTxFee", log); err != nil {
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

// ParseSetTxFee is a log parse operation binding the contract event 0x446a287c2114fa54d0083d97ce8f6f15b2ce29fa1c2df4b5a580d581ea7c4ad3.
//
// Solidity: event SetTxFee(uint64 chid, uint256 fee)
func (_NFTBridge *NFTBridgeFilterer) ParseSetTxFee(log types.Log) (*NFTBridgeSetTxFee, error) {
	event := new(NFTBridgeSetTxFee)
	if err := _NFTBridge.contract.UnpackLog(event, "SetTxFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NFTBridge contract.
type NFTBridgeUnpausedIterator struct {
	Event *NFTBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *NFTBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBridgeUnpaused)
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
		it.Event = new(NFTBridgeUnpaused)
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
func (it *NFTBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBridgeUnpaused represents a Unpaused event raised by the NFTBridge contract.
type NFTBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NFTBridge *NFTBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NFTBridgeUnpausedIterator, error) {

	logs, sub, err := _NFTBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NFTBridgeUnpausedIterator{contract: _NFTBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NFTBridge *NFTBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NFTBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _NFTBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBridgeUnpaused)
				if err := _NFTBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_NFTBridge *NFTBridgeFilterer) ParseUnpaused(log types.Log) (*NFTBridgeUnpaused, error) {
	event := new(NFTBridgeUnpaused)
	if err := _NFTBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
