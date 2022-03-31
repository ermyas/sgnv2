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

// MsgDataTypesBridgeTransferParams is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesBridgeTransferParams struct {
	Request []byte
	Sigs    [][]byte
	Signers []common.Address
	Powers  []*big.Int
}

// MsgDataTypesMsgWithTransferExecutionParams is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesMsgWithTransferExecutionParams struct {
	Message  []byte
	Transfer MsgDataTypesTransferInfo
	Sigs     [][]byte
	Signers  []common.Address
	Powers   []*big.Int
}

// MsgDataTypesRouteInfo is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesRouteInfo struct {
	Sender     common.Address
	Receiver   common.Address
	SrcChainId uint64
	SrcTxHash  [32]byte
}

// MsgDataTypesTransferInfo is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesTransferInfo struct {
	T          uint8
	Sender     common.Address
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Wdseq      uint64
	SrcChainId uint64
	RefId      [32]byte
	SrcTxHash  [32]byte
}

// MsgBusRecvMetaData contains all meta data concerning the MsgBusRecv contract.
var MsgBusRecvMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CallReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidityBridge\",\"type\":\"address\"}],\"name\":\"LiquidityBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"NeedRetry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridge\",\"type\":\"address\"}],\"name\":\"PegBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridgeV2\",\"type\":\"address\"}],\"name\":\"PegBridgeV2Updated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVault\",\"type\":\"address\"}],\"name\":\"PegVaultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVaultV2\",\"type\":\"address\"}],\"name\":\"PegVaultV2Updated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preExecuteMessageGasUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"refundAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridgeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVaultV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usage\",\"type\":\"uint256\"}],\"name\":\"setPreExecuteMessageGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"transferAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MsgBusRecvABI is the input ABI used to generate the binding from.
// Deprecated: Use MsgBusRecvMetaData.ABI instead.
var MsgBusRecvABI = MsgBusRecvMetaData.ABI

// MsgBusRecv is an auto generated Go binding around an Ethereum contract.
type MsgBusRecv struct {
	MsgBusRecvCaller     // Read-only binding to the contract
	MsgBusRecvTransactor // Write-only binding to the contract
	MsgBusRecvFilterer   // Log filterer for contract events
}

// MsgBusRecvCaller is an auto generated read-only Go binding around an Ethereum contract.
type MsgBusRecvCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgBusRecvTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MsgBusRecvTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgBusRecvFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MsgBusRecvFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgBusRecvSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MsgBusRecvSession struct {
	Contract     *MsgBusRecv       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MsgBusRecvCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MsgBusRecvCallerSession struct {
	Contract *MsgBusRecvCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MsgBusRecvTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MsgBusRecvTransactorSession struct {
	Contract     *MsgBusRecvTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MsgBusRecvRaw is an auto generated low-level Go binding around an Ethereum contract.
type MsgBusRecvRaw struct {
	Contract *MsgBusRecv // Generic contract binding to access the raw methods on
}

// MsgBusRecvCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MsgBusRecvCallerRaw struct {
	Contract *MsgBusRecvCaller // Generic read-only contract binding to access the raw methods on
}

// MsgBusRecvTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MsgBusRecvTransactorRaw struct {
	Contract *MsgBusRecvTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMsgBusRecv creates a new instance of MsgBusRecv, bound to a specific deployed contract.
func NewMsgBusRecv(address common.Address, backend bind.ContractBackend) (*MsgBusRecv, error) {
	contract, err := bindMsgBusRecv(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MsgBusRecv{MsgBusRecvCaller: MsgBusRecvCaller{contract: contract}, MsgBusRecvTransactor: MsgBusRecvTransactor{contract: contract}, MsgBusRecvFilterer: MsgBusRecvFilterer{contract: contract}}, nil
}

// NewMsgBusRecvCaller creates a new read-only instance of MsgBusRecv, bound to a specific deployed contract.
func NewMsgBusRecvCaller(address common.Address, caller bind.ContractCaller) (*MsgBusRecvCaller, error) {
	contract, err := bindMsgBusRecv(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvCaller{contract: contract}, nil
}

// NewMsgBusRecvTransactor creates a new write-only instance of MsgBusRecv, bound to a specific deployed contract.
func NewMsgBusRecvTransactor(address common.Address, transactor bind.ContractTransactor) (*MsgBusRecvTransactor, error) {
	contract, err := bindMsgBusRecv(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvTransactor{contract: contract}, nil
}

// NewMsgBusRecvFilterer creates a new log filterer instance of MsgBusRecv, bound to a specific deployed contract.
func NewMsgBusRecvFilterer(address common.Address, filterer bind.ContractFilterer) (*MsgBusRecvFilterer, error) {
	contract, err := bindMsgBusRecv(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvFilterer{contract: contract}, nil
}

// bindMsgBusRecv binds a generic wrapper to an already deployed contract.
func bindMsgBusRecv(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MsgBusRecvABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgBusRecv *MsgBusRecvRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgBusRecv.Contract.MsgBusRecvCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgBusRecv *MsgBusRecvRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.MsgBusRecvTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgBusRecv *MsgBusRecvRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.MsgBusRecvTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgBusRecv *MsgBusRecvCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgBusRecv.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgBusRecv *MsgBusRecvTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgBusRecv *MsgBusRecvTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.contract.Transact(opts, method, params...)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MsgBusRecv *MsgBusRecvCaller) ExecutedMessages(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _MsgBusRecv.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MsgBusRecv *MsgBusRecvSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MsgBusRecv.Contract.ExecutedMessages(&_MsgBusRecv.CallOpts, arg0)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MsgBusRecv *MsgBusRecvCallerSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MsgBusRecv.Contract.ExecutedMessages(&_MsgBusRecv.CallOpts, arg0)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MsgBusRecv *MsgBusRecvCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgBusRecv.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MsgBusRecv *MsgBusRecvSession) LiquidityBridge() (common.Address, error) {
	return _MsgBusRecv.Contract.LiquidityBridge(&_MsgBusRecv.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MsgBusRecv *MsgBusRecvCallerSession) LiquidityBridge() (common.Address, error) {
	return _MsgBusRecv.Contract.LiquidityBridge(&_MsgBusRecv.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MsgBusRecv *MsgBusRecvCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgBusRecv.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MsgBusRecv *MsgBusRecvSession) Owner() (common.Address, error) {
	return _MsgBusRecv.Contract.Owner(&_MsgBusRecv.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MsgBusRecv *MsgBusRecvCallerSession) Owner() (common.Address, error) {
	return _MsgBusRecv.Contract.Owner(&_MsgBusRecv.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MsgBusRecv *MsgBusRecvCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgBusRecv.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MsgBusRecv *MsgBusRecvSession) PegBridge() (common.Address, error) {
	return _MsgBusRecv.Contract.PegBridge(&_MsgBusRecv.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MsgBusRecv *MsgBusRecvCallerSession) PegBridge() (common.Address, error) {
	return _MsgBusRecv.Contract.PegBridge(&_MsgBusRecv.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MsgBusRecv *MsgBusRecvCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgBusRecv.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MsgBusRecv *MsgBusRecvSession) PegBridgeV2() (common.Address, error) {
	return _MsgBusRecv.Contract.PegBridgeV2(&_MsgBusRecv.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MsgBusRecv *MsgBusRecvCallerSession) PegBridgeV2() (common.Address, error) {
	return _MsgBusRecv.Contract.PegBridgeV2(&_MsgBusRecv.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MsgBusRecv *MsgBusRecvCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgBusRecv.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MsgBusRecv *MsgBusRecvSession) PegVault() (common.Address, error) {
	return _MsgBusRecv.Contract.PegVault(&_MsgBusRecv.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MsgBusRecv *MsgBusRecvCallerSession) PegVault() (common.Address, error) {
	return _MsgBusRecv.Contract.PegVault(&_MsgBusRecv.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MsgBusRecv *MsgBusRecvCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgBusRecv.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MsgBusRecv *MsgBusRecvSession) PegVaultV2() (common.Address, error) {
	return _MsgBusRecv.Contract.PegVaultV2(&_MsgBusRecv.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MsgBusRecv *MsgBusRecvCallerSession) PegVaultV2() (common.Address, error) {
	return _MsgBusRecv.Contract.PegVaultV2(&_MsgBusRecv.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MsgBusRecv *MsgBusRecvCaller) PreExecuteMessageGasUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MsgBusRecv.contract.Call(opts, &out, "preExecuteMessageGasUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MsgBusRecv *MsgBusRecvSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MsgBusRecv.Contract.PreExecuteMessageGasUsage(&_MsgBusRecv.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MsgBusRecv *MsgBusRecvCallerSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MsgBusRecv.Contract.PreExecuteMessageGasUsage(&_MsgBusRecv.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.ExecuteMessage(&_MsgBusRecv.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.ExecuteMessage(&_MsgBusRecv.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.ExecuteMessageWithTransfer(&_MsgBusRecv.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.ExecuteMessageWithTransfer(&_MsgBusRecv.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.ExecuteMessageWithTransferRefund(&_MsgBusRecv.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.ExecuteMessageWithTransferRefund(&_MsgBusRecv.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) RefundAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "refundAndExecuteMsg", _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MsgBusRecv *MsgBusRecvSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.RefundAndExecuteMsg(&_MsgBusRecv.TransactOpts, _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.RefundAndExecuteMsg(&_MsgBusRecv.TransactOpts, _transferParams, _msgParams)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) SetLiquidityBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "setLiquidityBridge", _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetLiquidityBridge(&_MsgBusRecv.TransactOpts, _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetLiquidityBridge(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) SetPegBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "setPegBridge", _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPegBridge(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPegBridge(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) SetPegBridgeV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "setPegBridgeV2", _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPegBridgeV2(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPegBridgeV2(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) SetPegVault(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "setPegVault", _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPegVault(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPegVault(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) SetPegVaultV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "setPegVaultV2", _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPegVaultV2(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPegVaultV2(&_MsgBusRecv.TransactOpts, _addr)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) SetPreExecuteMessageGasUsage(opts *bind.TransactOpts, _usage *big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "setPreExecuteMessageGasUsage", _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MsgBusRecv *MsgBusRecvSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPreExecuteMessageGasUsage(&_MsgBusRecv.TransactOpts, _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.SetPreExecuteMessageGasUsage(&_MsgBusRecv.TransactOpts, _usage)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) TransferAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "transferAndExecuteMsg", _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MsgBusRecv *MsgBusRecvSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.TransferAndExecuteMsg(&_MsgBusRecv.TransactOpts, _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.TransferAndExecuteMsg(&_MsgBusRecv.TransactOpts, _transferParams, _msgParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MsgBusRecv *MsgBusRecvTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MsgBusRecv *MsgBusRecvSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.TransferOwnership(&_MsgBusRecv.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MsgBusRecv *MsgBusRecvTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MsgBusRecv.Contract.TransferOwnership(&_MsgBusRecv.TransactOpts, newOwner)
}

// MsgBusRecvCallRevertedIterator is returned from FilterCallReverted and is used to iterate over the raw logs and unpacked data for CallReverted events raised by the MsgBusRecv contract.
type MsgBusRecvCallRevertedIterator struct {
	Event *MsgBusRecvCallReverted // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvCallRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvCallReverted)
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
		it.Event = new(MsgBusRecvCallReverted)
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
func (it *MsgBusRecvCallRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvCallRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvCallReverted represents a CallReverted event raised by the MsgBusRecv contract.
type MsgBusRecvCallReverted struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallReverted is a free log retrieval operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterCallReverted(opts *bind.FilterOpts) (*MsgBusRecvCallRevertedIterator, error) {

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvCallRevertedIterator{contract: _MsgBusRecv.contract, event: "CallReverted", logs: logs, sub: sub}, nil
}

// WatchCallReverted is a free log subscription operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchCallReverted(opts *bind.WatchOpts, sink chan<- *MsgBusRecvCallReverted) (event.Subscription, error) {

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvCallReverted)
				if err := _MsgBusRecv.contract.UnpackLog(event, "CallReverted", log); err != nil {
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

// ParseCallReverted is a log parse operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MsgBusRecv *MsgBusRecvFilterer) ParseCallReverted(log types.Log) (*MsgBusRecvCallReverted, error) {
	event := new(MsgBusRecvCallReverted)
	if err := _MsgBusRecv.contract.UnpackLog(event, "CallReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgBusRecvExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MsgBusRecv contract.
type MsgBusRecvExecutedIterator struct {
	Event *MsgBusRecvExecuted // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvExecuted)
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
		it.Event = new(MsgBusRecvExecuted)
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
func (it *MsgBusRecvExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvExecuted represents a Executed event raised by the MsgBusRecv contract.
type MsgBusRecvExecuted struct {
	MsgType    uint8
	MsgId      [32]byte
	Status     uint8
	Receiver   common.Address
	SrcChainId uint64
	SrcTxHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterExecuted(opts *bind.FilterOpts, receiver []common.Address) (*MsgBusRecvExecutedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvExecutedIterator{contract: _MsgBusRecv.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MsgBusRecvExecuted, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvExecuted)
				if err := _MsgBusRecv.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MsgBusRecv *MsgBusRecvFilterer) ParseExecuted(log types.Log) (*MsgBusRecvExecuted, error) {
	event := new(MsgBusRecvExecuted)
	if err := _MsgBusRecv.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgBusRecvLiquidityBridgeUpdatedIterator is returned from FilterLiquidityBridgeUpdated and is used to iterate over the raw logs and unpacked data for LiquidityBridgeUpdated events raised by the MsgBusRecv contract.
type MsgBusRecvLiquidityBridgeUpdatedIterator struct {
	Event *MsgBusRecvLiquidityBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvLiquidityBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvLiquidityBridgeUpdated)
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
		it.Event = new(MsgBusRecvLiquidityBridgeUpdated)
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
func (it *MsgBusRecvLiquidityBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvLiquidityBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvLiquidityBridgeUpdated represents a LiquidityBridgeUpdated event raised by the MsgBusRecv contract.
type MsgBusRecvLiquidityBridgeUpdated struct {
	LiquidityBridge common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityBridgeUpdated is a free log retrieval operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterLiquidityBridgeUpdated(opts *bind.FilterOpts) (*MsgBusRecvLiquidityBridgeUpdatedIterator, error) {

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvLiquidityBridgeUpdatedIterator{contract: _MsgBusRecv.contract, event: "LiquidityBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityBridgeUpdated is a free log subscription operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchLiquidityBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MsgBusRecvLiquidityBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvLiquidityBridgeUpdated)
				if err := _MsgBusRecv.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
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

// ParseLiquidityBridgeUpdated is a log parse operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MsgBusRecv *MsgBusRecvFilterer) ParseLiquidityBridgeUpdated(log types.Log) (*MsgBusRecvLiquidityBridgeUpdated, error) {
	event := new(MsgBusRecvLiquidityBridgeUpdated)
	if err := _MsgBusRecv.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgBusRecvNeedRetryIterator is returned from FilterNeedRetry and is used to iterate over the raw logs and unpacked data for NeedRetry events raised by the MsgBusRecv contract.
type MsgBusRecvNeedRetryIterator struct {
	Event *MsgBusRecvNeedRetry // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvNeedRetryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvNeedRetry)
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
		it.Event = new(MsgBusRecvNeedRetry)
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
func (it *MsgBusRecvNeedRetryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvNeedRetryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvNeedRetry represents a NeedRetry event raised by the MsgBusRecv contract.
type MsgBusRecvNeedRetry struct {
	MsgType    uint8
	MsgId      [32]byte
	SrcChainId uint64
	SrcTxHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNeedRetry is a free log retrieval operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterNeedRetry(opts *bind.FilterOpts) (*MsgBusRecvNeedRetryIterator, error) {

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvNeedRetryIterator{contract: _MsgBusRecv.contract, event: "NeedRetry", logs: logs, sub: sub}, nil
}

// WatchNeedRetry is a free log subscription operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchNeedRetry(opts *bind.WatchOpts, sink chan<- *MsgBusRecvNeedRetry) (event.Subscription, error) {

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvNeedRetry)
				if err := _MsgBusRecv.contract.UnpackLog(event, "NeedRetry", log); err != nil {
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

// ParseNeedRetry is a log parse operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MsgBusRecv *MsgBusRecvFilterer) ParseNeedRetry(log types.Log) (*MsgBusRecvNeedRetry, error) {
	event := new(MsgBusRecvNeedRetry)
	if err := _MsgBusRecv.contract.UnpackLog(event, "NeedRetry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgBusRecvOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MsgBusRecv contract.
type MsgBusRecvOwnershipTransferredIterator struct {
	Event *MsgBusRecvOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvOwnershipTransferred)
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
		it.Event = new(MsgBusRecvOwnershipTransferred)
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
func (it *MsgBusRecvOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvOwnershipTransferred represents a OwnershipTransferred event raised by the MsgBusRecv contract.
type MsgBusRecvOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MsgBusRecvOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvOwnershipTransferredIterator{contract: _MsgBusRecv.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MsgBusRecvOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvOwnershipTransferred)
				if err := _MsgBusRecv.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MsgBusRecv *MsgBusRecvFilterer) ParseOwnershipTransferred(log types.Log) (*MsgBusRecvOwnershipTransferred, error) {
	event := new(MsgBusRecvOwnershipTransferred)
	if err := _MsgBusRecv.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgBusRecvPegBridgeUpdatedIterator is returned from FilterPegBridgeUpdated and is used to iterate over the raw logs and unpacked data for PegBridgeUpdated events raised by the MsgBusRecv contract.
type MsgBusRecvPegBridgeUpdatedIterator struct {
	Event *MsgBusRecvPegBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvPegBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvPegBridgeUpdated)
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
		it.Event = new(MsgBusRecvPegBridgeUpdated)
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
func (it *MsgBusRecvPegBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvPegBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvPegBridgeUpdated represents a PegBridgeUpdated event raised by the MsgBusRecv contract.
type MsgBusRecvPegBridgeUpdated struct {
	PegBridge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeUpdated is a free log retrieval operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterPegBridgeUpdated(opts *bind.FilterOpts) (*MsgBusRecvPegBridgeUpdatedIterator, error) {

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvPegBridgeUpdatedIterator{contract: _MsgBusRecv.contract, event: "PegBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeUpdated is a free log subscription operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchPegBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MsgBusRecvPegBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvPegBridgeUpdated)
				if err := _MsgBusRecv.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
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

// ParsePegBridgeUpdated is a log parse operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MsgBusRecv *MsgBusRecvFilterer) ParsePegBridgeUpdated(log types.Log) (*MsgBusRecvPegBridgeUpdated, error) {
	event := new(MsgBusRecvPegBridgeUpdated)
	if err := _MsgBusRecv.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgBusRecvPegBridgeV2UpdatedIterator is returned from FilterPegBridgeV2Updated and is used to iterate over the raw logs and unpacked data for PegBridgeV2Updated events raised by the MsgBusRecv contract.
type MsgBusRecvPegBridgeV2UpdatedIterator struct {
	Event *MsgBusRecvPegBridgeV2Updated // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvPegBridgeV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvPegBridgeV2Updated)
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
		it.Event = new(MsgBusRecvPegBridgeV2Updated)
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
func (it *MsgBusRecvPegBridgeV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvPegBridgeV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvPegBridgeV2Updated represents a PegBridgeV2Updated event raised by the MsgBusRecv contract.
type MsgBusRecvPegBridgeV2Updated struct {
	PegBridgeV2 common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeV2Updated is a free log retrieval operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterPegBridgeV2Updated(opts *bind.FilterOpts) (*MsgBusRecvPegBridgeV2UpdatedIterator, error) {

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvPegBridgeV2UpdatedIterator{contract: _MsgBusRecv.contract, event: "PegBridgeV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeV2Updated is a free log subscription operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchPegBridgeV2Updated(opts *bind.WatchOpts, sink chan<- *MsgBusRecvPegBridgeV2Updated) (event.Subscription, error) {

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvPegBridgeV2Updated)
				if err := _MsgBusRecv.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
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

// ParsePegBridgeV2Updated is a log parse operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MsgBusRecv *MsgBusRecvFilterer) ParsePegBridgeV2Updated(log types.Log) (*MsgBusRecvPegBridgeV2Updated, error) {
	event := new(MsgBusRecvPegBridgeV2Updated)
	if err := _MsgBusRecv.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgBusRecvPegVaultUpdatedIterator is returned from FilterPegVaultUpdated and is used to iterate over the raw logs and unpacked data for PegVaultUpdated events raised by the MsgBusRecv contract.
type MsgBusRecvPegVaultUpdatedIterator struct {
	Event *MsgBusRecvPegVaultUpdated // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvPegVaultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvPegVaultUpdated)
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
		it.Event = new(MsgBusRecvPegVaultUpdated)
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
func (it *MsgBusRecvPegVaultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvPegVaultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvPegVaultUpdated represents a PegVaultUpdated event raised by the MsgBusRecv contract.
type MsgBusRecvPegVaultUpdated struct {
	PegVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPegVaultUpdated is a free log retrieval operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterPegVaultUpdated(opts *bind.FilterOpts) (*MsgBusRecvPegVaultUpdatedIterator, error) {

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvPegVaultUpdatedIterator{contract: _MsgBusRecv.contract, event: "PegVaultUpdated", logs: logs, sub: sub}, nil
}

// WatchPegVaultUpdated is a free log subscription operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchPegVaultUpdated(opts *bind.WatchOpts, sink chan<- *MsgBusRecvPegVaultUpdated) (event.Subscription, error) {

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvPegVaultUpdated)
				if err := _MsgBusRecv.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
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

// ParsePegVaultUpdated is a log parse operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MsgBusRecv *MsgBusRecvFilterer) ParsePegVaultUpdated(log types.Log) (*MsgBusRecvPegVaultUpdated, error) {
	event := new(MsgBusRecvPegVaultUpdated)
	if err := _MsgBusRecv.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgBusRecvPegVaultV2UpdatedIterator is returned from FilterPegVaultV2Updated and is used to iterate over the raw logs and unpacked data for PegVaultV2Updated events raised by the MsgBusRecv contract.
type MsgBusRecvPegVaultV2UpdatedIterator struct {
	Event *MsgBusRecvPegVaultV2Updated // Event containing the contract specifics and raw log

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
func (it *MsgBusRecvPegVaultV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgBusRecvPegVaultV2Updated)
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
		it.Event = new(MsgBusRecvPegVaultV2Updated)
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
func (it *MsgBusRecvPegVaultV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgBusRecvPegVaultV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgBusRecvPegVaultV2Updated represents a PegVaultV2Updated event raised by the MsgBusRecv contract.
type MsgBusRecvPegVaultV2Updated struct {
	PegVaultV2 common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPegVaultV2Updated is a free log retrieval operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MsgBusRecv *MsgBusRecvFilterer) FilterPegVaultV2Updated(opts *bind.FilterOpts) (*MsgBusRecvPegVaultV2UpdatedIterator, error) {

	logs, sub, err := _MsgBusRecv.contract.FilterLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return &MsgBusRecvPegVaultV2UpdatedIterator{contract: _MsgBusRecv.contract, event: "PegVaultV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegVaultV2Updated is a free log subscription operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MsgBusRecv *MsgBusRecvFilterer) WatchPegVaultV2Updated(opts *bind.WatchOpts, sink chan<- *MsgBusRecvPegVaultV2Updated) (event.Subscription, error) {

	logs, sub, err := _MsgBusRecv.contract.WatchLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgBusRecvPegVaultV2Updated)
				if err := _MsgBusRecv.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
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

// ParsePegVaultV2Updated is a log parse operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MsgBusRecv *MsgBusRecvFilterer) ParsePegVaultV2Updated(log types.Log) (*MsgBusRecvPegVaultV2Updated, error) {
	event := new(MsgBusRecvPegVaultV2Updated)
	if err := _MsgBusRecv.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
