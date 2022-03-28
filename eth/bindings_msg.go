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

// IMessageBusMetaData contains all meta data concerning the IMessageBus contract.
var IMessageBusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageBusMetaData.ABI instead.
var IMessageBusABI = IMessageBusMetaData.ABI

// IMessageBus is an auto generated Go binding around an Ethereum contract.
type IMessageBus struct {
	IMessageBusCaller     // Read-only binding to the contract
	IMessageBusTransactor // Write-only binding to the contract
	IMessageBusFilterer   // Log filterer for contract events
}

// IMessageBusCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageBusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageBusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageBusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageBusSession struct {
	Contract     *IMessageBus      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageBusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageBusCallerSession struct {
	Contract *IMessageBusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IMessageBusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageBusTransactorSession struct {
	Contract     *IMessageBusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMessageBusRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageBusRaw struct {
	Contract *IMessageBus // Generic contract binding to access the raw methods on
}

// IMessageBusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageBusCallerRaw struct {
	Contract *IMessageBusCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageBusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageBusTransactorRaw struct {
	Contract *IMessageBusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageBus creates a new instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBus(address common.Address, backend bind.ContractBackend) (*IMessageBus, error) {
	contract, err := bindIMessageBus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageBus{IMessageBusCaller: IMessageBusCaller{contract: contract}, IMessageBusTransactor: IMessageBusTransactor{contract: contract}, IMessageBusFilterer: IMessageBusFilterer{contract: contract}}, nil
}

// NewIMessageBusCaller creates a new read-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusCaller(address common.Address, caller bind.ContractCaller) (*IMessageBusCaller, error) {
	contract, err := bindIMessageBus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusCaller{contract: contract}, nil
}

// NewIMessageBusTransactor creates a new write-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageBusTransactor, error) {
	contract, err := bindIMessageBus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusTransactor{contract: contract}, nil
}

// NewIMessageBusFilterer creates a new log filterer instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageBusFilterer, error) {
	contract, err := bindIMessageBus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageBusFilterer{contract: contract}, nil
}

// bindIMessageBus binds a generic wrapper to an already deployed contract.
func bindIMessageBus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageBusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.IMessageBusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusSession) CalcFee(_message []byte) (*big.Int, error) {
	return _IMessageBus.Contract.CalcFee(&_IMessageBus.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _IMessageBus.Contract.CalcFee(&_IMessageBus.CallOpts, _message)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusSession) LiquidityBridge() (common.Address, error) {
	return _IMessageBus.Contract.LiquidityBridge(&_IMessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) LiquidityBridge() (common.Address, error) {
	return _IMessageBus.Contract.LiquidityBridge(&_IMessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusSession) PegBridge() (common.Address, error) {
	return _IMessageBus.Contract.PegBridge(&_IMessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegBridge() (common.Address, error) {
	return _IMessageBus.Contract.PegBridge(&_IMessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusSession) PegBridgeV2() (common.Address, error) {
	return _IMessageBus.Contract.PegBridgeV2(&_IMessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegBridgeV2() (common.Address, error) {
	return _IMessageBus.Contract.PegBridgeV2(&_IMessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusSession) PegVault() (common.Address, error) {
	return _IMessageBus.Contract.PegVault(&_IMessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegVault() (common.Address, error) {
	return _IMessageBus.Contract.PegVault(&_IMessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusSession) PegVaultV2() (common.Address, error) {
	return _IMessageBus.Contract.PegVaultV2(&_IMessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegVaultV2() (common.Address, error) {
	return _IMessageBus.Contract.PegVaultV2(&_IMessageBus.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransfer(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransfer(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransferRefund(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransferRefund(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactor) SendMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessageWithTransfer(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessageWithTransfer(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// IMessageReceiverAppMetaData contains all meta data concerning the IMessageReceiverApp contract.
var IMessageReceiverAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecuctionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecuctionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecuctionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecuctionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IMessageReceiverAppABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageReceiverAppMetaData.ABI instead.
var IMessageReceiverAppABI = IMessageReceiverAppMetaData.ABI

// IMessageReceiverApp is an auto generated Go binding around an Ethereum contract.
type IMessageReceiverApp struct {
	IMessageReceiverAppCaller     // Read-only binding to the contract
	IMessageReceiverAppTransactor // Write-only binding to the contract
	IMessageReceiverAppFilterer   // Log filterer for contract events
}

// IMessageReceiverAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageReceiverAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageReceiverAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageReceiverAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageReceiverAppSession struct {
	Contract     *IMessageReceiverApp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMessageReceiverAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageReceiverAppCallerSession struct {
	Contract *IMessageReceiverAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMessageReceiverAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageReceiverAppTransactorSession struct {
	Contract     *IMessageReceiverAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMessageReceiverAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageReceiverAppRaw struct {
	Contract *IMessageReceiverApp // Generic contract binding to access the raw methods on
}

// IMessageReceiverAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageReceiverAppCallerRaw struct {
	Contract *IMessageReceiverAppCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageReceiverAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageReceiverAppTransactorRaw struct {
	Contract *IMessageReceiverAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageReceiverApp creates a new instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverApp(address common.Address, backend bind.ContractBackend) (*IMessageReceiverApp, error) {
	contract, err := bindIMessageReceiverApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverApp{IMessageReceiverAppCaller: IMessageReceiverAppCaller{contract: contract}, IMessageReceiverAppTransactor: IMessageReceiverAppTransactor{contract: contract}, IMessageReceiverAppFilterer: IMessageReceiverAppFilterer{contract: contract}}, nil
}

// NewIMessageReceiverAppCaller creates a new read-only instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppCaller(address common.Address, caller bind.ContractCaller) (*IMessageReceiverAppCaller, error) {
	contract, err := bindIMessageReceiverApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppCaller{contract: contract}, nil
}

// NewIMessageReceiverAppTransactor creates a new write-only instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageReceiverAppTransactor, error) {
	contract, err := bindIMessageReceiverApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppTransactor{contract: contract}, nil
}

// NewIMessageReceiverAppFilterer creates a new log filterer instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageReceiverAppFilterer, error) {
	contract, err := bindIMessageReceiverApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppFilterer{contract: contract}, nil
}

// bindIMessageReceiverApp binds a generic wrapper to an already deployed contract.
func bindIMessageReceiverApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageReceiverAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageReceiverApp *IMessageReceiverAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageReceiverApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageReceiverApp *IMessageReceiverAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageReceiverApp *IMessageReceiverAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.contract.Transact(opts, method, params...)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessage(&_IMessageReceiverApp.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessage(&_IMessageReceiverApp.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransfer(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransfer(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferFallback(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferFallback(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferRefund(&_IMessageReceiverApp.TransactOpts, _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferRefund(&_IMessageReceiverApp.TransactOpts, _token, _amount, _message, _executor)
}

// MessageBusMetaData contains all meta data concerning the MessageBus contract.
var MessageBusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CallReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBase\",\"type\":\"uint256\"}],\"name\":\"FeeBaseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"}],\"name\":\"FeePerByteUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidityBridge\",\"type\":\"address\"}],\"name\":\"LiquidityBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageWithTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"NeedRetry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridge\",\"type\":\"address\"}],\"name\":\"PegBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridgeV2\",\"type\":\"address\"}],\"name\":\"PegBridgeV2Updated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVault\",\"type\":\"address\"}],\"name\":\"PegVaultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVaultV2\",\"type\":\"address\"}],\"name\":\"PegVaultV2Updated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preExecuteMessageGasUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"refundAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeeBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeePerByte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridgeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVaultV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usage\",\"type\":\"uint256\"}],\"name\":\"setPreExecuteMessageGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"transferAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620039be380380620039be83398101604081905262000034916200011d565b84848484848a6200004533620000b4565b6001600160a01b03908116608052600580546001600160a01b0319908116978316979097179055600680548716958216959095179094556007805486169385169390931790925560088054851691841691909117905560098054909316911617905550620001b1945050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146200011a57600080fd5b50565b60008060008060008060c087890312156200013757600080fd5b8651620001448162000104565b6020880151909650620001578162000104565b60408801519095506200016a8162000104565b60608801519094506200017d8162000104565b6080880151909350620001908162000104565b60a0880151909250620001a38162000104565b809150509295509295509295565b6080516137ea620001d4600039600081816104b201526107f701526137ea6000f3fe6080604052600436106101d85760003560e01c806382980dc411610102578063ccf2683b11610095578063e2c1ed2511610064578063e2c1ed2514610551578063f2fde38b14610571578063f60bbe2a14610591578063f83b0fb9146105a757600080fd5b8063ccf2683b146104a0578063cd2abd66146104d4578063d8257d1714610511578063dfa2dbaf1461053157600080fd5b806395e911a8116100d157806395e911a8146104375780639b05a7751461044d5780639f3ce55a1461046d578063c66a9c5a1461048057600080fd5b806382980dc4146103a157806382efd502146103d95780638da5cb5b146103f957806395b12c271461041757600080fd5b80634586f3311161017a578063588be02b11610149578063588be02b146103215780635b3e5f5014610341578063723d0a9d1461036e5780637b80ab201461038e57600080fd5b80634586f331146102a5578063468a2d04146102c55780635335dca2146102d8578063584e45e11461030b57600080fd5b8063359ef75b116101b6578063359ef75b1461023f5780633f395aff1461025f57806340d0d026146102725780634289fbb31461029257600080fd5b806303cbfe66146101dd57806306c28bd6146101ff5780632ff4c4111461021f575b600080fd5b3480156101e957600080fd5b506101fd6101f8366004612a6b565b6105c7565b005b34801561020b57600080fd5b506101fd61021a366004612a86565b6106c0565b34801561022b57600080fd5b506101fd61023a366004612aeb565b61074c565b34801561024b57600080fd5b506101fd61025a366004612b9f565b6109a5565b6101fd61026d366004612c46565b6109c1565b34801561027e57600080fd5b506101fd61028d366004612d4f565b610cb4565b6101fd6102a0366004612dbb565b610d0c565b3480156102b157600080fd5b506101fd6102c0366004612a86565b610df5565b6101fd6102d3366004612e33565b610e51565b3480156102e457600080fd5b506102f86102f3366004612ef8565b6110cd565b6040519081526020015b60405180910390f35b34801561031757600080fd5b506102f8600a5481565b34801561032d57600080fd5b506101fd61033c366004612a6b565b6110f3565b34801561034d57600080fd5b506102f861035c366004612a6b565b60036020526000908152604090205481565b34801561037a57600080fd5b506101fd610389366004612d4f565b6111e0565b6101fd61039c366004612c46565b61122e565b3480156103ad57600080fd5b506005546103c1906001600160a01b031681565b6040516001600160a01b039091168152602001610302565b3480156103e557600080fd5b506101fd6103f4366004612a6b565b61143d565b34801561040557600080fd5b506000546001600160a01b03166103c1565b34801561042357600080fd5b506008546103c1906001600160a01b031681565b34801561044357600080fd5b506102f860015481565b34801561045957600080fd5b506101fd610468366004612a6b565b61152a565b6101fd61047b366004612f3a565b611617565b34801561048c57600080fd5b506009546103c1906001600160a01b031681565b3480156104ac57600080fd5b506103c17f000000000000000000000000000000000000000000000000000000000000000081565b3480156104e057600080fd5b506105046104ef366004612a86565b60046020526000908152604090205460ff1681565b6040516103029190612fbe565b34801561051d57600080fd5b506007546103c1906001600160a01b031681565b34801561053d57600080fd5b506006546103c1906001600160a01b031681565b34801561055d57600080fd5b506101fd61056c366004612a86565b6116fa565b34801561057d57600080fd5b506101fd61058c366004612a6b565b611786565b34801561059d57600080fd5b506102f860025481565b3480156105b357600080fd5b506101fd6105c2366004612a6b565b611865565b336105da6000546001600160a01b031690565b6001600160a01b0316146106235760405162461bcd60e51b8152602060048201819052602482015260008051602061379583398151915260448201526064015b60405180910390fd5b6001600160a01b03811661066b5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161061a565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527fd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d906020015b60405180910390a150565b336106d36000546001600160a01b031690565b6001600160a01b0316146107175760405162461bcd60e51b81526020600482018190526024820152600080516020613795833981519152604482015260640161061a565b60018190556040518181527f892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2906020016106b5565b6000463060405160200161079d92919091825260601b6001600160601b03191660208201527f77697468647261774665650000000000000000000000000000000000000000006034820152603f0190565b60408051808303601f19018152828252805160209182012090830181905260608c901b6001600160601b0319168383015260548084018c9052825180850390910181526074840192839052633416de1160e11b90925292507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163682dbc229161083e918b908b908b908b908b908b90607801613178565b60006040518083038186803b15801561085657600080fd5b505afa15801561086a573d6000803e3d6000fd5b505050506001600160a01b038916600090815260036020526040812054610891908a6131ec565b9050600081116108e35760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f20776974686472617700000000000000604482015260640161061a565b6001600160a01b038a166000818152600360205260408082208c90555190919061c35090849084818181858888f193505050503d8060008114610942576040519150601f19603f3d011682016040523d82523d6000602084013e610947565b606091505b50509050806109985760405162461bcd60e51b815260206004820152601660248201527f6661696c656420746f2077697468647261772066656500000000000000000000604482015260640161061a565b5050505050505050505050565b6109ad611952565b6109ba85858585856119b6565b5050505050565b60006109cc88611a6e565b90506000808281526004602081905260409091205460ff16908111156109f4576109f4612f94565b14610a415760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c726561647920657865637574656400000000000000604482015260640161061a565b6000818152600460208181526040808420805460ff1916909317909255815146918101919091526001600160601b03193060601b16918101919091527f4d657373616765576974685472616e73666572000000000000000000000000006054820152606701604051602081830303815290604052805190602001209050600560009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e8e6101000135604051602001610b00959493929190613203565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401610b379796959493929190613178565b60006040518083038186803b158015610b4f57600080fd5b505afa158015610b63573d6000803e3d6000fd5b50505050600080610b758b8e8e612327565b90506001816002811115610b8b57610b8b612f94565b1415610b9a5760019150610c64565b6002816002811115610bae57610bae612f94565b1415610c2f576000848152600460205260408120805460ff19166001835b02179055507fe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c6000858d60c0016020810190610c089190613225565b8e6101000135604051610c1e949392919061325f565b60405180910390a150505050610ca9565b610c3a8b8e8e61248f565b90506001816002811115610c5057610c50612f94565b1415610c5f5760039150610c64565b600291505b60008481526004602081905260409091208054849260ff19909116906001908490811115610c9457610c94612f94565b0217905550610ca484838d6124ca565b505050505b505050505050505050565b610ccd610cc76040830160208401613292565b8361253c565b610d08610cda82806132b3565b60208401610cec6101408601866132fa565b610cfa6101608801886132fa565b61039c6101808a018a6132fa565b5050565b46851415610d4e5760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9259608a1b604482015260640161061a565b6000610d5a83836110cd565b905080341015610d9f5760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b604482015260640161061a565b336001600160a01b03167f172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f6688888888888834604051610de49796959493929190613344565b60405180910390a250505050505050565b33610e086000546001600160a01b031690565b6001600160a01b031614610e4c5760405162461bcd60e51b81526020600482018190526024820152600080516020613795833981519152604482015260640161061a565b600a55565b6000610e5e888b8b61279b565b90506000808281526004602081905260409091205460ff1690811115610e8657610e86612f94565b14610ed35760405162461bcd60e51b815260206004820152601860248201527f6d65737361676520616c72656164792065786563757465640000000000000000604482015260640161061a565b600081815260046020818152604092839020805460ff1916909217909155815146818301523060601b6001600160601b031916818401527f4d6573736167650000000000000000000000000000000000000000000000000060548201528251603b818303018152605b820184528051920191909120600554607b8301829052609b8084018690528451808503909101815260bb840194859052633416de1160e11b90945290926001600160a01b039091169163682dbc2291610fa3918c908c908c908c908c908c9060bf01613178565b60006040518083038186803b158015610fbb57600080fd5b505afa158015610fcf573d6000803e3d6000fd5b50505050600080610fe18b8e8e61280a565b90506001816002811115610ff757610ff7612f94565b1415611006576001915061108d565b600281600281111561101a5761101a612f94565b1415611088576000848152600460205260408120805460ff191660018302179055507fe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c6001858d60400160208101906110739190613225565b8e60600135604051610c1e949392919061325f565b600291505b60008481526004602081905260409091208054849260ff199091169060019084908111156110bd576110bd612f94565b0217905550610ca484838d612869565b6002546000906110dd9083613391565b6001546110ea91906133b0565b90505b92915050565b336111066000546001600160a01b031690565b6001600160a01b03161461114a5760405162461bcd60e51b81526020600482018190526024820152600080516020613795833981519152604482015260640161061a565b6001600160a01b0381166111925760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161061a565b600580546001600160a01b0319166001600160a01b0383169081179091556040519081527fbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058906020016106b5565b6111f3610cc76040830160208401613292565b610d0861120082806132b3565b602084016112126101408601866132fa565b6112206101608801886132fa565b61026d6101808a018a6132fa565b600061123988611a6e565b90506000808281526004602081905260409091205460ff169081111561126157611261612f94565b146112ae5760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c726561647920657865637574656400000000000000604482015260640161061a565b6000818152600460208181526040808420805460ff1916909317909255815146918101919091526001600160601b03193060601b16918101919091527f4d657373616765576974685472616e73666572526566756e64000000000000006054820152606d01604051602081830303815290604052805190602001209050600560009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e8e610100013560405160200161136d959493929190613203565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b81526004016113a49796959493929190613178565b60006040518083038186803b1580156113bc57600080fd5b505afa1580156113d0573d6000803e3d6000fd5b505050506000806113e28b8e8e6128cd565b905060018160028111156113f8576113f8612f94565b14156114075760019150610c64565b600281600281111561141b5761141b612f94565b1415610c5f576000848152600460205260408120805460ff1916600183610bcc565b336114506000546001600160a01b031690565b6001600160a01b0316146114945760405162461bcd60e51b81526020600482018190526024820152600080516020613795833981519152604482015260640161061a565b6001600160a01b0381166114dc5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161061a565b600880546001600160a01b0319166001600160a01b0383169081179091556040519081527ffb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8906020016106b5565b3361153d6000546001600160a01b031690565b6001600160a01b0316146115815760405162461bcd60e51b81526020600482018190526024820152600080516020613795833981519152604482015260640161061a565b6001600160a01b0381166115c95760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161061a565b600780546001600160a01b0319166001600160a01b0383169081179091556040519081527fa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad906020016106b5565b468314156116595760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9259608a1b604482015260640161061a565b600061166583836110cd565b9050803410156116aa5760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b604482015260640161061a565b336001600160a01b03167fce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e486868686346040516116eb9594939291906133c8565b60405180910390a25050505050565b3361170d6000546001600160a01b031690565b6001600160a01b0316146117515760405162461bcd60e51b81526020600482018190526024820152600080516020613795833981519152604482015260640161061a565b60028190556040518181527f210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c906020016106b5565b336117996000546001600160a01b031690565b6001600160a01b0316146117dd5760405162461bcd60e51b81526020600482018190526024820152600080516020613795833981519152604482015260640161061a565b6001600160a01b0381166118595760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161061a565b61186281612924565b50565b336118786000546001600160a01b031690565b6001600160a01b0316146118bc5760405162461bcd60e51b81526020600482018190526024820152600080516020613795833981519152604482015260640161061a565b6001600160a01b0381166119045760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161061a565b600980546001600160a01b0319166001600160a01b0383169081179091556040519081527f918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5906020016106b5565b6000546001600160a01b0316156119ab5760405162461bcd60e51b815260206004820152601160248201527f6f776e657220616c726561647920736574000000000000000000000000000000604482015260640161061a565b6119b433612924565b565b6005546001600160a01b031615611a0f5760405162461bcd60e51b815260206004820152601b60248201527f6c697175696469747942726964676520616c7265616479207365740000000000604482015260640161061a565b600580546001600160a01b03199081166001600160a01b03978816179091556006805482169587169590951790945560078054851693861693909317909255600880548416918516919091179055600980549092169216919091179055565b600080806001611a816020860186613292565b6006811115611a9257611a92612f94565b1415611c2d57611aa86040850160208601612a6b565b611ab86060860160408701612a6b565b611ac86080870160608801612a6b565b6080870135611add60e0890160c08a01613225565b6040516001600160601b0319606096871b8116602083015294861b851660348201529290941b9092166048820152605c8101919091526001600160c01b031960c092831b8116607c8301524690921b909116608482015260e0850135608c82015260ac0160408051808303601f19018152908290528051602090910120600554633c64f04b60e01b8352600483018290529093506001600160a01b031691508190633c64f04b9060240160206040518083038186803b158015611b9f57600080fd5b505afa158015611bb3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd79190613403565b1515600114611c285760405162461bcd60e51b815260206004820152601660248201527f6272696467652072656c6179206e6f7420657869737400000000000000000000604482015260640161061a565b6122f2565b6002611c3c6020860186613292565b6006811115611c4d57611c4d612f94565b1415611dba5746611c6460c0860160a08701613225565b611c746060870160408801612a6b565b611c846080880160608901612a6b565b6040516001600160c01b031960c095861b811660208301529390941b90921660288401526001600160601b0319606091821b8116603085015291901b1660448201526080850135605882015260780160408051808303601f19018152908290528051602090910120600554631c13568560e31b8352600483018290529093506001600160a01b03169150819063e09ab4289060240160206040518083038186803b158015611d3157600080fd5b505afa158015611d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d699190613403565b1515600114611c285760405162461bcd60e51b815260206004820152601960248201527f627269646765207769746864726177206e6f7420657869737400000000000000604482015260640161061a565b6003611dc96020860186613292565b6006811115611dda57611dda612f94565b1480611e0357506004611df06020860186613292565b6006811115611e0157611e01612f94565b145b1561208557611e186060850160408601612a6b565b611e286080860160608701612a6b565b6080860135611e3d6040880160208901612a6b565b611e4d60e0890160c08a01613225565b604051606095861b6001600160601b0319908116602083015294861b851660348201526048810193909352931b909116606882015260c09190911b6001600160c01b031916607c82015260e0850135608482015260a40160408051601f19818403018152919052805160209091012091506003611ecd6020860186613292565b6006811115611ede57611ede612f94565b1415611fb557506006546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e64725906024015b60206040518083038186803b158015611f2c57600080fd5b505afa158015611f40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f649190613403565b1515600114611c285760405162461bcd60e51b815260206004820152601560248201527f6d696e74207265636f7264206e6f742065786973740000000000000000000000604482015260640161061a565b506007546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e647259060240160206040518083038186803b158015611ffc57600080fd5b505afa158015612010573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120349190613403565b1515600114611c285760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f7420657869737400000000000000604482015260640161061a565b60056120946020860186613292565b60068111156120a5576120a5612f94565b14806120ce575060066120bb6020860186613292565b60068111156120cc576120cc612f94565b145b156122f25760056120e26020860186613292565b60068111156120f3576120f3612f94565b141561210b57506008546001600160a01b0316612119565b506009546001600160a01b03165b6121296060850160408601612a6b565b6121396080860160608701612a6b565b608086013561214e6040880160208901612a6b565b61215e60e0890160c08a01613225565b604051606095861b6001600160601b0319908116602083015294861b85166034820152604881019390935290841b8316606883015260c01b6001600160c01b031916607c82015260e087013560848201529183901b1660a482015260b80160408051601f198184030181529190528051602090910120915060056121e56020860186613292565b60068111156121f6576121f6612f94565b1415612229576040516301e6472560e01b8152600481018390526001600160a01b038216906301e6472590602401611f14565b6040516301e6472560e01b8152600481018390526001600160a01b038216906301e647259060240160206040518083038186803b15801561226957600080fd5b505afa15801561227d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122a19190613403565b15156001146122f25760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f7420657869737400000000000000604482015260640161061a565b600081836040516020016123089392919061343c565b6040516020818303038152906040528051906020012092505050919050565b6000805a90506000806123406060880160408901612a6b565b6001600160a01b031634631f34afff60e21b61236260408b0160208c01612a6b565b61237260808c0160608d01612a6b565b60808c013561238760e08e0160c08f01613225565b8c8c336040516024016123a09796959493929190613468565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252905161240b91906134c1565b60006040518083038185875af1925050503d8060008114612448576040519150601f19603f3d011682016040523d82523d6000602084013e61244d565b606091505b50915091508115612476578080602001905181019061246c91906134dd565b9350505050612488565b6124808382612974565b600093505050505b9392505050565b6000805a90506000806124a86060880160408901612a6b565b6001600160a01b031634632d5bd7e360e11b61236260408b0160208c01612a6b565b6124da6060820160408301612a6b565b6001600160a01b03167fa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d6000858561251860e0870160c08801613225565b86610100013560405161252f9594939291906134fe565b60405180910390a2505050565b600182600681111561255057612550612f94565b14156125f1576005546001600160a01b031663cdd1b25d61257183806132b3565b61257e60208601866132fa565b61258b60408801886132fa565b61259860608a018a6132fa565b6040518963ffffffff1660e01b81526004016125bb98979695949392919061353c565b600060405180830381600087803b1580156125d557600080fd5b505af11580156125e9573d6000803e3d6000fd5b505050505050565b600282600681111561260557612605612f94565b1415612626576005546001600160a01b031663a21a928061257183806132b3565b600382600681111561263a5761263a612f94565b141561265b576006546001600160a01b031663f873430261257183806132b3565b600582600681111561266f5761266f612f94565b1415612731576008546001600160a01b031663f873430261269083806132b3565b61269d60208601866132fa565b6126aa60408801886132fa565b6126b760608a018a6132fa565b6040518963ffffffff1660e01b81526004016126da98979695949392919061353c565b602060405180830381600087803b1580156126f457600080fd5b505af1158015612708573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061272c919061359c565b505050565b600482600681111561274557612745612f94565b1415612766576007546001600160a01b031663a21a928061257183806132b3565b600682600681111561277a5761277a612f94565b1415610d08576009546001600160a01b031663a21a928061269083806132b3565b600060016127ac6020860186612a6b565b6127bc6040870160208801612a6b565b6127cc6060880160408901613225565b87606001354688886040516020016127eb9897969594939291906135b5565b6040516020818303038152906040528051906020012090509392505050565b6000805a90506000806128236040880160208901612a6b565b6001600160a01b031634639c649fdf60e01b61284260208b018b612a6b565b61285260608c0160408d01613225565b8a8a336040516024016123a0959493929190613629565b6128796040820160208301612a6b565b6001600160a01b03167fa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d600185856128b76060870160408801613225565b866060013560405161252f9594939291906134fe565b6000805a90506000806128e66060880160408901612a6b565b6001600160a01b0316346305e5a4c160e11b61290860808b0160608c01612a6b565b8a608001358a8a336040516024016123a0959493929190613672565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005a90506000600a544561298991906131ec565b905080841080156129a457506129a060408561369c565b8211155b156129ab57fe5b7fffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f6129d5846129f0565b6040516129e291906136be565b60405180910390a150505050565b6060604482511015612a3557505060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c79000000602082015290565b600482019150818060200190518101906110ed91906136e7565b80356001600160a01b0381168114612a6657600080fd5b919050565b600060208284031215612a7d57600080fd5b6110ea82612a4f565b600060208284031215612a9857600080fd5b5035919050565b60008083601f840112612ab157600080fd5b50813567ffffffffffffffff811115612ac957600080fd5b6020830191508360208260051b8501011115612ae457600080fd5b9250929050565b60008060008060008060008060a0898b031215612b0757600080fd5b612b1089612a4f565b975060208901359650604089013567ffffffffffffffff80821115612b3457600080fd5b612b408c838d01612a9f565b909850965060608b0135915080821115612b5957600080fd5b612b658c838d01612a9f565b909650945060808b0135915080821115612b7e57600080fd5b50612b8b8b828c01612a9f565b999c989b5096995094979396929594505050565b600080600080600060a08688031215612bb757600080fd5b612bc086612a4f565b9450612bce60208701612a4f565b9350612bdc60408701612a4f565b9250612bea60608701612a4f565b9150612bf860808701612a4f565b90509295509295909350565b60008083601f840112612c1657600080fd5b50813567ffffffffffffffff811115612c2e57600080fd5b602083019150836020828501011115612ae457600080fd5b6000806000806000806000806000898b036101a0811215612c6657600080fd5b8a3567ffffffffffffffff80821115612c7e57600080fd5b612c8a8e838f01612c04565b909c509a508a9150610120601f1984011215612ca557600080fd5b60208d0199506101408d0135925080831115612cc057600080fd5b612ccc8e848f01612a9f565b90995097506101608d0135925088915080831115612ce957600080fd5b612cf58e848f01612a9f565b90975095506101808d0135925086915080831115612d1257600080fd5b5050612d208c828d01612a9f565b915080935050809150509295985092959850929598565b600060808284031215612d4957600080fd5b50919050565b60008060408385031215612d6257600080fd5b823567ffffffffffffffff80821115612d7a57600080fd5b612d8686838701612d37565b93506020850135915080821115612d9c57600080fd5b5083016101a08186031215612db057600080fd5b809150509250929050565b60008060008060008060a08789031215612dd457600080fd5b612ddd87612a4f565b955060208701359450612df260408801612a4f565b935060608701359250608087013567ffffffffffffffff811115612e1557600080fd5b612e2189828a01612c04565b979a9699509497509295939492505050565b60008060008060008060008060006101008a8c031215612e5257600080fd5b893567ffffffffffffffff80821115612e6a57600080fd5b612e768d838e01612c04565b909b509950899150612e8b8d60208e01612d37565b985060a08c0135915080821115612ea157600080fd5b612ead8d838e01612a9f565b909850965060c08c0135915080821115612ec657600080fd5b612ed28d838e01612a9f565b909650945060e08c0135915080821115612eeb57600080fd5b50612d208c828d01612a9f565b60008060208385031215612f0b57600080fd5b823567ffffffffffffffff811115612f2257600080fd5b612f2e85828601612c04565b90969095509350505050565b60008060008060608587031215612f5057600080fd5b612f5985612a4f565b935060208501359250604085013567ffffffffffffffff811115612f7c57600080fd5b612f8887828801612c04565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b60058110612fba57612fba612f94565b9052565b602081016110ed8284612faa565b60005b83811015612fe7578181015183820152602001612fcf565b83811115612ff6576000848401525b50505050565b60008151808452613014816020860160208601612fcc565b601f01601f19169290920160200192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260006020808501808196508560051b810191508460005b878110156130d55782840389528135601e1988360301811261308c57600080fd5b8701803567ffffffffffffffff8111156130a557600080fd5b8036038913156130b457600080fd5b6130c18682898501613028565b9a87019a955050509084019060010161306b565b5091979650505050505050565b8183526000602080850194508260005b8581101561311e576001600160a01b0361310b83612a4f565b16875295820195908201906001016130f2565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561315b57600080fd5b8260051b8083602087013760009401602001938452509192915050565b60808152600061318b608083018a612ffc565b828103602084015261319e81898b613051565b905082810360408401526131b38187896130e2565b905082810360608401526131c8818587613129565b9a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156131fe576131fe6131d6565b500390565b8581528460208201528284604083013760409201918201526060019392505050565b60006020828403121561323757600080fd5b813567ffffffffffffffff8116811461248857600080fd5b60028110612fba57612fba612f94565b6080810161326d828761324f565b84602083015267ffffffffffffffff8416604083015282606083015295945050505050565b6000602082840312156132a457600080fd5b81356007811061248857600080fd5b6000808335601e198436030181126132ca57600080fd5b83018035915067ffffffffffffffff8211156132e557600080fd5b602001915036819003821315612ae457600080fd5b6000808335601e1984360301811261331157600080fd5b83018035915067ffffffffffffffff82111561332c57600080fd5b6020019150600581901b3603821315612ae457600080fd5b60006001600160a01b03808a16835288602084015280881660408401525085606083015260c0608083015261337d60c083018587613028565b90508260a083015298975050505050505050565b60008160001904831182151516156133ab576133ab6131d6565b500290565b600082198211156133c3576133c36131d6565b500190565b6001600160a01b03861681528460208201526080604082015260006133f1608083018587613028565b90508260608301529695505050505050565b60006020828403121561341557600080fd5b8151801515811461248857600080fd5b6002811061343557613435612f94565b60f81b9052565b6134468185613425565b60609290921b6001600160601b03191660018301526015820152603501919050565b60006001600160a01b03808a168352808916602084015287604084015267ffffffffffffffff8716606084015260c060808401526134aa60c084018688613028565b915080841660a08401525098975050505050505050565b600082516134d3818460208701612fcc565b9190910192915050565b6000602082840312156134ef57600080fd5b81516003811061248857600080fd5b60a0810161350c828861324f565b85602083015261351f6040830186612faa565b67ffffffffffffffff939093166060820152608001529392505050565b608081526000613550608083018a8c613028565b828103602084015261356381898b613051565b905082810360408401526135788187896130e2565b9050828103606084015261358d818587613129565b9b9a5050505050505050505050565b6000602082840312156135ae57600080fd5b5051919050565b6135bf818a613425565b60006bffffffffffffffffffffffff19808a60601b166001840152808960601b166015840152506001600160c01b0319808860c01b166029840152866031840152808660c01b16605184015250828460598401375060009101605901908152979650505050505050565b60006001600160a01b03808816835267ffffffffffffffff871660208401526080604084015261365d608084018688613028565b91508084166060840152509695505050505050565b60006001600160a01b0380881683528660208401526080604084015261365d608084018688613028565b6000826136b957634e487b7160e01b600052601260045260246000fd5b500490565b6020815260006110ea6020830184612ffc565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156136f957600080fd5b815167ffffffffffffffff8082111561371157600080fd5b818401915084601f83011261372557600080fd5b815181811115613737576137376136d1565b604051601f8201601f19908116603f0116810190838211818310171561375f5761375f6136d1565b8160405282815287602084870101111561377857600080fd5b613789836020830160208801612fcc565b97965050505050505056fe4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a2646970667358221220ff63be3785443446ca0ff125d755b0fc4dc68e77dc42b653eb98a62a1f47b12a64736f6c63430008090033",
}

// MessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusMetaData.ABI instead.
var MessageBusABI = MessageBusMetaData.ABI

// MessageBusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusMetaData.Bin instead.
var MessageBusBin = MessageBusMetaData.Bin

// DeployMessageBus deploys a new Ethereum contract, binding an instance of MessageBus to it.
func DeployMessageBus(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (common.Address, *types.Transaction, *MessageBus, error) {
	parsed, err := MessageBusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusBin), backend, _sigsVerifier, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBus{MessageBusCaller: MessageBusCaller{contract: contract}, MessageBusTransactor: MessageBusTransactor{contract: contract}, MessageBusFilterer: MessageBusFilterer{contract: contract}}, nil
}

// MessageBus is an auto generated Go binding around an Ethereum contract.
type MessageBus struct {
	MessageBusCaller     // Read-only binding to the contract
	MessageBusTransactor // Write-only binding to the contract
	MessageBusFilterer   // Log filterer for contract events
}

// MessageBusCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusSession struct {
	Contract     *MessageBus       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageBusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusCallerSession struct {
	Contract *MessageBusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MessageBusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusTransactorSession struct {
	Contract     *MessageBusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MessageBusRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusRaw struct {
	Contract *MessageBus // Generic contract binding to access the raw methods on
}

// MessageBusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusCallerRaw struct {
	Contract *MessageBusCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusTransactorRaw struct {
	Contract *MessageBusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBus creates a new instance of MessageBus, bound to a specific deployed contract.
func NewMessageBus(address common.Address, backend bind.ContractBackend) (*MessageBus, error) {
	contract, err := bindMessageBus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBus{MessageBusCaller: MessageBusCaller{contract: contract}, MessageBusTransactor: MessageBusTransactor{contract: contract}, MessageBusFilterer: MessageBusFilterer{contract: contract}}, nil
}

// NewMessageBusCaller creates a new read-only instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusCaller(address common.Address, caller bind.ContractCaller) (*MessageBusCaller, error) {
	contract, err := bindMessageBus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusCaller{contract: contract}, nil
}

// NewMessageBusTransactor creates a new write-only instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusTransactor, error) {
	contract, err := bindMessageBus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusTransactor{contract: contract}, nil
}

// NewMessageBusFilterer creates a new log filterer instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusFilterer, error) {
	contract, err := bindMessageBus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusFilterer{contract: contract}, nil
}

// bindMessageBus binds a generic wrapper to an already deployed contract.
func bindMessageBus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBus *MessageBusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBus.Contract.MessageBusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBus *MessageBusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBus.Contract.MessageBusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBus *MessageBusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBus.Contract.MessageBusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBus *MessageBusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBus *MessageBusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBus *MessageBusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBus.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBus.Contract.CalcFee(&_MessageBus.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBus.Contract.CalcFee(&_MessageBus.CallOpts, _message)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusCaller) ExecutedMessages(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBus.Contract.ExecutedMessages(&_MessageBus.CallOpts, arg0)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusCallerSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBus.Contract.ExecutedMessages(&_MessageBus.CallOpts, arg0)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusSession) FeeBase() (*big.Int, error) {
	return _MessageBus.Contract.FeeBase(&_MessageBus.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusCallerSession) FeeBase() (*big.Int, error) {
	return _MessageBus.Contract.FeeBase(&_MessageBus.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusCaller) FeePerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "feePerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusSession) FeePerByte() (*big.Int, error) {
	return _MessageBus.Contract.FeePerByte(&_MessageBus.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusCallerSession) FeePerByte() (*big.Int, error) {
	return _MessageBus.Contract.FeePerByte(&_MessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusSession) LiquidityBridge() (common.Address, error) {
	return _MessageBus.Contract.LiquidityBridge(&_MessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusCallerSession) LiquidityBridge() (common.Address, error) {
	return _MessageBus.Contract.LiquidityBridge(&_MessageBus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusSession) Owner() (common.Address, error) {
	return _MessageBus.Contract.Owner(&_MessageBus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusCallerSession) Owner() (common.Address, error) {
	return _MessageBus.Contract.Owner(&_MessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusSession) PegBridge() (common.Address, error) {
	return _MessageBus.Contract.PegBridge(&_MessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegBridge() (common.Address, error) {
	return _MessageBus.Contract.PegBridge(&_MessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusSession) PegBridgeV2() (common.Address, error) {
	return _MessageBus.Contract.PegBridgeV2(&_MessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegBridgeV2() (common.Address, error) {
	return _MessageBus.Contract.PegBridgeV2(&_MessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusSession) PegVault() (common.Address, error) {
	return _MessageBus.Contract.PegVault(&_MessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegVault() (common.Address, error) {
	return _MessageBus.Contract.PegVault(&_MessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusSession) PegVaultV2() (common.Address, error) {
	return _MessageBus.Contract.PegVaultV2(&_MessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegVaultV2() (common.Address, error) {
	return _MessageBus.Contract.PegVaultV2(&_MessageBus.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBus *MessageBusCaller) PreExecuteMessageGasUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "preExecuteMessageGasUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBus *MessageBusSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MessageBus.Contract.PreExecuteMessageGasUsage(&_MessageBus.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBus *MessageBusCallerSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MessageBus.Contract.PreExecuteMessageGasUsage(&_MessageBus.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusSession) SigsVerifier() (common.Address, error) {
	return _MessageBus.Contract.SigsVerifier(&_MessageBus.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusCallerSession) SigsVerifier() (common.Address, error) {
	return _MessageBus.Contract.SigsVerifier(&_MessageBus.CallOpts)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusCaller) WithdrawnFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "withdrawnFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBus.Contract.WithdrawnFees(&_MessageBus.CallOpts, arg0)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusCallerSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBus.Contract.WithdrawnFees(&_MessageBus.CallOpts, arg0)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessage(&_MessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessage(&_MessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransfer(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransfer(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransferRefund(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransferRefund(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusTransactor) Init(opts *bind.TransactOpts, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "init", _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusSession) Init(_liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.Init(&_MessageBus.TransactOpts, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusTransactorSession) Init(_liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.Init(&_MessageBus.TransactOpts, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusTransactor) RefundAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "refundAndExecuteMsg", _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.Contract.RefundAndExecuteMsg(&_MessageBus.TransactOpts, _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusTransactorSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.Contract.RefundAndExecuteMsg(&_MessageBus.TransactOpts, _transferParams, _msgParams)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactor) SendMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessage(&_MessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactorSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessage(&_MessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessageWithTransfer(&_MessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessageWithTransfer(&_MessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactor) SetFeeBase(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setFeeBase", _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeeBase(&_MessageBus.TransactOpts, _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactorSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeeBase(&_MessageBus.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactor) SetFeePerByte(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setFeePerByte", _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeePerByte(&_MessageBus.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactorSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeePerByte(&_MessageBus.TransactOpts, _fee)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetLiquidityBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setLiquidityBridge", _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetLiquidityBridge(&_MessageBus.TransactOpts, _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetLiquidityBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegBridge", _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegBridgeV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegBridgeV2", _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridgeV2(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridgeV2(&_MessageBus.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegVault(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegVault", _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVault(&_MessageBus.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVault(&_MessageBus.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegVaultV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegVaultV2", _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVaultV2(&_MessageBus.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVaultV2(&_MessageBus.TransactOpts, _addr)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBus *MessageBusTransactor) SetPreExecuteMessageGasUsage(opts *bind.TransactOpts, _usage *big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPreExecuteMessageGasUsage", _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBus *MessageBusSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPreExecuteMessageGasUsage(&_MessageBus.TransactOpts, _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBus *MessageBusTransactorSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPreExecuteMessageGasUsage(&_MessageBus.TransactOpts, _usage)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusTransactor) TransferAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "transferAndExecuteMsg", _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferAndExecuteMsg(&_MessageBus.TransactOpts, _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusTransactorSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferAndExecuteMsg(&_MessageBus.TransactOpts, _transferParams, _msgParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferOwnership(&_MessageBus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferOwnership(&_MessageBus.TransactOpts, newOwner)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.WithdrawFee(&_MessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.WithdrawFee(&_MessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// MessageBusCallRevertedIterator is returned from FilterCallReverted and is used to iterate over the raw logs and unpacked data for CallReverted events raised by the MessageBus contract.
type MessageBusCallRevertedIterator struct {
	Event *MessageBusCallReverted // Event containing the contract specifics and raw log

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
func (it *MessageBusCallRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusCallReverted)
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
		it.Event = new(MessageBusCallReverted)
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
func (it *MessageBusCallRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusCallRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusCallReverted represents a CallReverted event raised by the MessageBus contract.
type MessageBusCallReverted struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallReverted is a free log retrieval operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBus *MessageBusFilterer) FilterCallReverted(opts *bind.FilterOpts) (*MessageBusCallRevertedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return &MessageBusCallRevertedIterator{contract: _MessageBus.contract, event: "CallReverted", logs: logs, sub: sub}, nil
}

// WatchCallReverted is a free log subscription operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBus *MessageBusFilterer) WatchCallReverted(opts *bind.WatchOpts, sink chan<- *MessageBusCallReverted) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusCallReverted)
				if err := _MessageBus.contract.UnpackLog(event, "CallReverted", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParseCallReverted(log types.Log) (*MessageBusCallReverted, error) {
	event := new(MessageBusCallReverted)
	if err := _MessageBus.contract.UnpackLog(event, "CallReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MessageBus contract.
type MessageBusExecutedIterator struct {
	Event *MessageBusExecuted // Event containing the contract specifics and raw log

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
func (it *MessageBusExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusExecuted)
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
		it.Event = new(MessageBusExecuted)
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
func (it *MessageBusExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusExecuted represents a Executed event raised by the MessageBus contract.
type MessageBusExecuted struct {
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
func (_MessageBus *MessageBusFilterer) FilterExecuted(opts *bind.FilterOpts, receiver []common.Address) (*MessageBusExecutedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusExecutedIterator{contract: _MessageBus.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MessageBusExecuted, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusExecuted)
				if err := _MessageBus.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParseExecuted(log types.Log) (*MessageBusExecuted, error) {
	event := new(MessageBusExecuted)
	if err := _MessageBus.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusFeeBaseUpdatedIterator is returned from FilterFeeBaseUpdated and is used to iterate over the raw logs and unpacked data for FeeBaseUpdated events raised by the MessageBus contract.
type MessageBusFeeBaseUpdatedIterator struct {
	Event *MessageBusFeeBaseUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusFeeBaseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusFeeBaseUpdated)
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
		it.Event = new(MessageBusFeeBaseUpdated)
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
func (it *MessageBusFeeBaseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusFeeBaseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusFeeBaseUpdated represents a FeeBaseUpdated event raised by the MessageBus contract.
type MessageBusFeeBaseUpdated struct {
	FeeBase *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeBaseUpdated is a free log retrieval operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBus *MessageBusFilterer) FilterFeeBaseUpdated(opts *bind.FilterOpts) (*MessageBusFeeBaseUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "FeeBaseUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusFeeBaseUpdatedIterator{contract: _MessageBus.contract, event: "FeeBaseUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeBaseUpdated is a free log subscription operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBus *MessageBusFilterer) WatchFeeBaseUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusFeeBaseUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "FeeBaseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusFeeBaseUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "FeeBaseUpdated", log); err != nil {
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

// ParseFeeBaseUpdated is a log parse operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBus *MessageBusFilterer) ParseFeeBaseUpdated(log types.Log) (*MessageBusFeeBaseUpdated, error) {
	event := new(MessageBusFeeBaseUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "FeeBaseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusFeePerByteUpdatedIterator is returned from FilterFeePerByteUpdated and is used to iterate over the raw logs and unpacked data for FeePerByteUpdated events raised by the MessageBus contract.
type MessageBusFeePerByteUpdatedIterator struct {
	Event *MessageBusFeePerByteUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusFeePerByteUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusFeePerByteUpdated)
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
		it.Event = new(MessageBusFeePerByteUpdated)
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
func (it *MessageBusFeePerByteUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusFeePerByteUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusFeePerByteUpdated represents a FeePerByteUpdated event raised by the MessageBus contract.
type MessageBusFeePerByteUpdated struct {
	FeePerByte *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeePerByteUpdated is a free log retrieval operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBus *MessageBusFilterer) FilterFeePerByteUpdated(opts *bind.FilterOpts) (*MessageBusFeePerByteUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "FeePerByteUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusFeePerByteUpdatedIterator{contract: _MessageBus.contract, event: "FeePerByteUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePerByteUpdated is a free log subscription operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBus *MessageBusFilterer) WatchFeePerByteUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusFeePerByteUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "FeePerByteUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusFeePerByteUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "FeePerByteUpdated", log); err != nil {
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

// ParseFeePerByteUpdated is a log parse operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBus *MessageBusFilterer) ParseFeePerByteUpdated(log types.Log) (*MessageBusFeePerByteUpdated, error) {
	event := new(MessageBusFeePerByteUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "FeePerByteUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusLiquidityBridgeUpdatedIterator is returned from FilterLiquidityBridgeUpdated and is used to iterate over the raw logs and unpacked data for LiquidityBridgeUpdated events raised by the MessageBus contract.
type MessageBusLiquidityBridgeUpdatedIterator struct {
	Event *MessageBusLiquidityBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusLiquidityBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusLiquidityBridgeUpdated)
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
		it.Event = new(MessageBusLiquidityBridgeUpdated)
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
func (it *MessageBusLiquidityBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusLiquidityBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusLiquidityBridgeUpdated represents a LiquidityBridgeUpdated event raised by the MessageBus contract.
type MessageBusLiquidityBridgeUpdated struct {
	LiquidityBridge common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityBridgeUpdated is a free log retrieval operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBus *MessageBusFilterer) FilterLiquidityBridgeUpdated(opts *bind.FilterOpts) (*MessageBusLiquidityBridgeUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusLiquidityBridgeUpdatedIterator{contract: _MessageBus.contract, event: "LiquidityBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityBridgeUpdated is a free log subscription operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBus *MessageBusFilterer) WatchLiquidityBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusLiquidityBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusLiquidityBridgeUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParseLiquidityBridgeUpdated(log types.Log) (*MessageBusLiquidityBridgeUpdated, error) {
	event := new(MessageBusLiquidityBridgeUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the MessageBus contract.
type MessageBusMessageIterator struct {
	Event *MessageBusMessage // Event containing the contract specifics and raw log

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
func (it *MessageBusMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusMessage)
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
		it.Event = new(MessageBusMessage)
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
func (it *MessageBusMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusMessage represents a Message event raised by the MessageBus contract.
type MessageBusMessage struct {
	Sender     common.Address
	Receiver   common.Address
	DstChainId *big.Int
	Message    []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) FilterMessage(opts *bind.FilterOpts, sender []common.Address) (*MessageBusMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusMessageIterator{contract: _MessageBus.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *MessageBusMessage, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusMessage)
				if err := _MessageBus.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) ParseMessage(log types.Log) (*MessageBusMessage, error) {
	event := new(MessageBusMessage)
	if err := _MessageBus.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusMessageWithTransferIterator is returned from FilterMessageWithTransfer and is used to iterate over the raw logs and unpacked data for MessageWithTransfer events raised by the MessageBus contract.
type MessageBusMessageWithTransferIterator struct {
	Event *MessageBusMessageWithTransfer // Event containing the contract specifics and raw log

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
func (it *MessageBusMessageWithTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusMessageWithTransfer)
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
		it.Event = new(MessageBusMessageWithTransfer)
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
func (it *MessageBusMessageWithTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusMessageWithTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusMessageWithTransfer represents a MessageWithTransfer event raised by the MessageBus contract.
type MessageBusMessageWithTransfer struct {
	Sender        common.Address
	Receiver      common.Address
	DstChainId    *big.Int
	Bridge        common.Address
	SrcTransferId [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageWithTransfer is a free log retrieval operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) FilterMessageWithTransfer(opts *bind.FilterOpts, sender []common.Address) (*MessageBusMessageWithTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusMessageWithTransferIterator{contract: _MessageBus.contract, event: "MessageWithTransfer", logs: logs, sub: sub}, nil
}

// WatchMessageWithTransfer is a free log subscription operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) WatchMessageWithTransfer(opts *bind.WatchOpts, sink chan<- *MessageBusMessageWithTransfer, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusMessageWithTransfer)
				if err := _MessageBus.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
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

// ParseMessageWithTransfer is a log parse operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) ParseMessageWithTransfer(log types.Log) (*MessageBusMessageWithTransfer, error) {
	event := new(MessageBusMessageWithTransfer)
	if err := _MessageBus.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusNeedRetryIterator is returned from FilterNeedRetry and is used to iterate over the raw logs and unpacked data for NeedRetry events raised by the MessageBus contract.
type MessageBusNeedRetryIterator struct {
	Event *MessageBusNeedRetry // Event containing the contract specifics and raw log

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
func (it *MessageBusNeedRetryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusNeedRetry)
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
		it.Event = new(MessageBusNeedRetry)
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
func (it *MessageBusNeedRetryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusNeedRetryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusNeedRetry represents a NeedRetry event raised by the MessageBus contract.
type MessageBusNeedRetry struct {
	MsgType    uint8
	MsgId      [32]byte
	SrcChainId uint64
	SrcTxHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNeedRetry is a free log retrieval operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) FilterNeedRetry(opts *bind.FilterOpts) (*MessageBusNeedRetryIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return &MessageBusNeedRetryIterator{contract: _MessageBus.contract, event: "NeedRetry", logs: logs, sub: sub}, nil
}

// WatchNeedRetry is a free log subscription operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) WatchNeedRetry(opts *bind.WatchOpts, sink chan<- *MessageBusNeedRetry) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusNeedRetry)
				if err := _MessageBus.contract.UnpackLog(event, "NeedRetry", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParseNeedRetry(log types.Log) (*MessageBusNeedRetry, error) {
	event := new(MessageBusNeedRetry)
	if err := _MessageBus.contract.UnpackLog(event, "NeedRetry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBus contract.
type MessageBusOwnershipTransferredIterator struct {
	Event *MessageBusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusOwnershipTransferred)
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
		it.Event = new(MessageBusOwnershipTransferred)
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
func (it *MessageBusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBus contract.
type MessageBusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBus *MessageBusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusOwnershipTransferredIterator{contract: _MessageBus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBus *MessageBusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusOwnershipTransferred)
				if err := _MessageBus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusOwnershipTransferred, error) {
	event := new(MessageBusOwnershipTransferred)
	if err := _MessageBus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusPegBridgeUpdatedIterator is returned from FilterPegBridgeUpdated and is used to iterate over the raw logs and unpacked data for PegBridgeUpdated events raised by the MessageBus contract.
type MessageBusPegBridgeUpdatedIterator struct {
	Event *MessageBusPegBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusPegBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusPegBridgeUpdated)
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
		it.Event = new(MessageBusPegBridgeUpdated)
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
func (it *MessageBusPegBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusPegBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusPegBridgeUpdated represents a PegBridgeUpdated event raised by the MessageBus contract.
type MessageBusPegBridgeUpdated struct {
	PegBridge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeUpdated is a free log retrieval operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBus *MessageBusFilterer) FilterPegBridgeUpdated(opts *bind.FilterOpts) (*MessageBusPegBridgeUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusPegBridgeUpdatedIterator{contract: _MessageBus.contract, event: "PegBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeUpdated is a free log subscription operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBus *MessageBusFilterer) WatchPegBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusPegBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusPegBridgeUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParsePegBridgeUpdated(log types.Log) (*MessageBusPegBridgeUpdated, error) {
	event := new(MessageBusPegBridgeUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusPegBridgeV2UpdatedIterator is returned from FilterPegBridgeV2Updated and is used to iterate over the raw logs and unpacked data for PegBridgeV2Updated events raised by the MessageBus contract.
type MessageBusPegBridgeV2UpdatedIterator struct {
	Event *MessageBusPegBridgeV2Updated // Event containing the contract specifics and raw log

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
func (it *MessageBusPegBridgeV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusPegBridgeV2Updated)
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
		it.Event = new(MessageBusPegBridgeV2Updated)
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
func (it *MessageBusPegBridgeV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusPegBridgeV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusPegBridgeV2Updated represents a PegBridgeV2Updated event raised by the MessageBus contract.
type MessageBusPegBridgeV2Updated struct {
	PegBridgeV2 common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeV2Updated is a free log retrieval operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBus *MessageBusFilterer) FilterPegBridgeV2Updated(opts *bind.FilterOpts) (*MessageBusPegBridgeV2UpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return &MessageBusPegBridgeV2UpdatedIterator{contract: _MessageBus.contract, event: "PegBridgeV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeV2Updated is a free log subscription operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBus *MessageBusFilterer) WatchPegBridgeV2Updated(opts *bind.WatchOpts, sink chan<- *MessageBusPegBridgeV2Updated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusPegBridgeV2Updated)
				if err := _MessageBus.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParsePegBridgeV2Updated(log types.Log) (*MessageBusPegBridgeV2Updated, error) {
	event := new(MessageBusPegBridgeV2Updated)
	if err := _MessageBus.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusPegVaultUpdatedIterator is returned from FilterPegVaultUpdated and is used to iterate over the raw logs and unpacked data for PegVaultUpdated events raised by the MessageBus contract.
type MessageBusPegVaultUpdatedIterator struct {
	Event *MessageBusPegVaultUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusPegVaultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusPegVaultUpdated)
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
		it.Event = new(MessageBusPegVaultUpdated)
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
func (it *MessageBusPegVaultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusPegVaultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusPegVaultUpdated represents a PegVaultUpdated event raised by the MessageBus contract.
type MessageBusPegVaultUpdated struct {
	PegVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPegVaultUpdated is a free log retrieval operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBus *MessageBusFilterer) FilterPegVaultUpdated(opts *bind.FilterOpts) (*MessageBusPegVaultUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusPegVaultUpdatedIterator{contract: _MessageBus.contract, event: "PegVaultUpdated", logs: logs, sub: sub}, nil
}

// WatchPegVaultUpdated is a free log subscription operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBus *MessageBusFilterer) WatchPegVaultUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusPegVaultUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusPegVaultUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParsePegVaultUpdated(log types.Log) (*MessageBusPegVaultUpdated, error) {
	event := new(MessageBusPegVaultUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusPegVaultV2UpdatedIterator is returned from FilterPegVaultV2Updated and is used to iterate over the raw logs and unpacked data for PegVaultV2Updated events raised by the MessageBus contract.
type MessageBusPegVaultV2UpdatedIterator struct {
	Event *MessageBusPegVaultV2Updated // Event containing the contract specifics and raw log

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
func (it *MessageBusPegVaultV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusPegVaultV2Updated)
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
		it.Event = new(MessageBusPegVaultV2Updated)
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
func (it *MessageBusPegVaultV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusPegVaultV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusPegVaultV2Updated represents a PegVaultV2Updated event raised by the MessageBus contract.
type MessageBusPegVaultV2Updated struct {
	PegVaultV2 common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPegVaultV2Updated is a free log retrieval operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBus *MessageBusFilterer) FilterPegVaultV2Updated(opts *bind.FilterOpts) (*MessageBusPegVaultV2UpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return &MessageBusPegVaultV2UpdatedIterator{contract: _MessageBus.contract, event: "PegVaultV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegVaultV2Updated is a free log subscription operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBus *MessageBusFilterer) WatchPegVaultV2Updated(opts *bind.WatchOpts, sink chan<- *MessageBusPegVaultV2Updated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusPegVaultV2Updated)
				if err := _MessageBus.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParsePegVaultV2Updated(log types.Log) (*MessageBusPegVaultV2Updated, error) {
	event := new(MessageBusPegVaultV2Updated)
	if err := _MessageBus.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverMetaData contains all meta data concerning the MessageBusReceiver contract.
var MessageBusReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CallReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidityBridge\",\"type\":\"address\"}],\"name\":\"LiquidityBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"NeedRetry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridge\",\"type\":\"address\"}],\"name\":\"PegBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridgeV2\",\"type\":\"address\"}],\"name\":\"PegBridgeV2Updated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVault\",\"type\":\"address\"}],\"name\":\"PegVaultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVaultV2\",\"type\":\"address\"}],\"name\":\"PegVaultV2Updated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preExecuteMessageGasUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"refundAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridgeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVaultV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usage\",\"type\":\"uint256\"}],\"name\":\"setPreExecuteMessageGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"transferAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002ea738038062002ea783398101604081905262000034916200010f565b6200003f33620000a2565b600280546001600160a01b03199081166001600160a01b039788161790915560038054821695871695909517909455600480548516938616939093179092556005805484169185169190911790556006805490921692169190911790556200017f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146200010a57600080fd5b919050565b600080600080600060a086880312156200012857600080fd5b6200013386620000f2565b94506200014360208701620000f2565b93506200015360408701620000f2565b92506200016360608701620000f2565b91506200017360808701620000f2565b90509295509295909350565b612d18806200018f6000396000f3fe60806040526004361061015f5760003560e01c806382efd502116100c0578063cd2abd6611610074578063dfa2dbaf11610059578063dfa2dbaf1461039b578063f2fde38b146103bb578063f83b0fb9146103db57600080fd5b8063cd2abd661461033e578063d8257d171461037b57600080fd5b806395b12c27116100a557806395b12c27146102de5780639b05a775146102fe578063c66a9c5a1461031e57600080fd5b806382efd502146102a05780638da5cb5b146102c057600080fd5b8063584e45e111610117578063723d0a9d116100fc578063723d0a9d146102355780637b80ab201461025557806382980dc41461026857600080fd5b8063584e45e1146101ec578063588be02b1461021557600080fd5b806340d0d0261161014857806340d0d026146101995780634586f331146101b9578063468a2d04146101d957600080fd5b806303cbfe66146101645780633f395aff14610186575b600080fd5b34801561017057600080fd5b5061018461017f3660046122a4565b6103fb565b005b61018461019436600461234d565b610506565b3480156101a557600080fd5b506101846101b4366004612456565b6107fb565b3480156101c557600080fd5b506101846101d43660046124c2565b610853565b6101846101e73660046124db565b6108c1565b3480156101f857600080fd5b5061020260075481565b6040519081526020015b60405180910390f35b34801561022157600080fd5b506101846102303660046122a4565b610b41565b34801561024157600080fd5b50610184610250366004612456565b610c40565b61018461026336600461234d565b610c8e565b34801561027457600080fd5b50600254610288906001600160a01b031681565b6040516001600160a01b03909116815260200161020c565b3480156102ac57600080fd5b506101846102bb3660046122a4565b610ea1565b3480156102cc57600080fd5b506000546001600160a01b0316610288565b3480156102ea57600080fd5b50600554610288906001600160a01b031681565b34801561030a57600080fd5b506101846103193660046122a4565b610fa0565b34801561032a57600080fd5b50600654610288906001600160a01b031681565b34801561034a57600080fd5b5061036e6103593660046124c2565b60016020526000908152604090205460ff1681565b60405161020c91906125ca565b34801561038757600080fd5b50600454610288906001600160a01b031681565b3480156103a757600080fd5b50600354610288906001600160a01b031681565b3480156103c757600080fd5b506101846103d63660046122a4565b61109f565b3480156103e757600080fd5b506101846103f63660046122a4565b611190565b3361040e6000546001600160a01b031690565b6001600160a01b0316146104695760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b0381166104b15760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610460565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527fd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d906020015b60405180910390a150565b60006105118861128f565b90506000808281526001602052604090205460ff166004811115610537576105376125a0565b146105845760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c7265616479206578656375746564000000000000006044820152606401610460565b6000818152600160209081526040808320805460ff19166004179055805146928101929092526bffffffffffffffffffffffff193060601b16908201527f4d657373616765576974685472616e73666572000000000000000000000000006054820152606701604051602081830303815290604052805190602001209050600260009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e8e61010001356040516020016106449594939291906125d8565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b815260040161067b97969594939291906127aa565b60006040518083038186803b15801561069357600080fd5b505afa1580156106a7573d6000803e3d6000fd5b505050506000806106b98b8e8e611b5a565b905060018160028111156106cf576106cf6125a0565b14156106de57600191506107ad565b60028160028111156106f2576106f26125a0565b141561077857600084815260016020819052604082208054909160ff1990911690835b02179055507fe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c6000858d60c00160208101906107519190612808565b8e61010001356040516107679493929190612842565b60405180910390a1505050506107f0565b6107838b8e8e611cc2565b90506001816002811115610799576107996125a0565b14156107a857600391506107ad565b600291505b60008481526001602081905260409091208054849260ff19909116908360048111156107db576107db6125a0565b02179055506107eb84838d611cfd565b505050505b505050505050505050565b61081461080e6040830160208401612875565b83611d6f565b61084f6108218280612896565b602084016108336101408601866128dd565b6108416101608801886128dd565b6102636101808a018a6128dd565b5050565b336108666000546001600160a01b031690565b6001600160a01b0316146108bc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b600755565b60006108ce888b8b611fce565b90506000808281526001602052604090205460ff1660048111156108f4576108f46125a0565b146109415760405162461bcd60e51b815260206004820152601860248201527f6d65737361676520616c726561647920657865637574656400000000000000006044820152606401610460565b6000818152600160209081526040808320805460ff19166004179055805146928101929092526bffffffffffffffffffffffff193060601b16908201527f4d657373616765000000000000000000000000000000000000000000000000006054820152605b0160408051601f1981840301815282825280516020918201206002549184018190528383018690528251808503840181526060850193849052633416de1160e11b90935293506001600160a01b03169163682dbc2291610a14918c908c908c908c908c908c906064016127aa565b60006040518083038186803b158015610a2c57600080fd5b505afa158015610a40573d6000803e3d6000fd5b50505050600080610a528b8e8e61203d565b90506001816002811115610a6857610a686125a0565b1415610a775760019150610b03565b6002816002811115610a8b57610a8b6125a0565b1415610afe57600084815260016020819052604082208054909160ff19909116908302179055507fe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c6001858d6040016020810190610ae99190612808565b8e606001356040516107679493929190612842565b600291505b60008481526001602081905260409091208054849260ff1990911690836004811115610b3157610b316125a0565b02179055506107eb84838d61209c565b33610b546000546001600160a01b031690565b6001600160a01b031614610baa5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b6001600160a01b038116610bf25760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610460565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527fbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058906020016104fb565b610c5361080e6040830160208401612875565b61084f610c608280612896565b60208401610c726101408601866128dd565b610c806101608801886128dd565b6101946101808a018a6128dd565b6000610c998861128f565b90506000808281526001602052604090205460ff166004811115610cbf57610cbf6125a0565b14610d0c5760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c7265616479206578656375746564000000000000006044820152606401610460565b6000818152600160209081526040808320805460ff19166004179055805146928101929092526bffffffffffffffffffffffff193060601b16908201527f4d657373616765576974685472616e73666572526566756e64000000000000006054820152606d01604051602081830303815290604052805190602001209050600260009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e8e6101000135604051602001610dcc9594939291906125d8565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401610e0397969594939291906127aa565b60006040518083038186803b158015610e1b57600080fd5b505afa158015610e2f573d6000803e3d6000fd5b50505050600080610e418b8e8e612100565b90506001816002811115610e5757610e576125a0565b1415610e6657600191506107ad565b6002816002811115610e7a57610e7a6125a0565b14156107a857600084815260016020819052604082208054909160ff199091169083610715565b33610eb46000546001600160a01b031690565b6001600160a01b031614610f0a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b6001600160a01b038116610f525760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610460565b600580546001600160a01b0319166001600160a01b0383169081179091556040519081527ffb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8906020016104fb565b33610fb36000546001600160a01b031690565b6001600160a01b0316146110095760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b6001600160a01b0381166110515760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610460565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527fa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad906020016104fb565b336110b26000546001600160a01b031690565b6001600160a01b0316146111085760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b6001600160a01b0381166111845760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610460565b61118d81612157565b50565b336111a36000546001600160a01b031690565b6001600160a01b0316146111f95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b6001600160a01b0381166112415760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610460565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527f918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5906020016104fb565b6000808060016112a26020860186612875565b60068111156112b3576112b36125a0565b1415611453576112c960408501602086016122a4565b6112d960608601604087016122a4565b6112e960808701606088016122a4565b60808701356112fe60e0890160c08a01612808565b6040516bffffffffffffffffffffffff19606096871b8116602083015294861b851660348201529290941b9092166048820152605c8101919091526001600160c01b031960c092831b8116607c8301524690921b909116608482015260e0850135608c82015260ac0160408051808303601f19018152908290528051602090910120600254633c64f04b60e01b8352600483018290529093506001600160a01b031691508190633c64f04b9060240160206040518083038186803b1580156113c557600080fd5b505afa1580156113d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113fd9190612927565b151560011461144e5760405162461bcd60e51b815260206004820152601660248201527f6272696467652072656c6179206e6f74206578697374000000000000000000006044820152606401610460565b611b25565b60026114626020860186612875565b6006811115611473576114736125a0565b14156115e5574661148a60c0860160a08701612808565b61149a60608701604088016122a4565b6114aa60808801606089016122a4565b6040516001600160c01b031960c095861b811660208301529390941b90921660288401526bffffffffffffffffffffffff19606091821b8116603085015291901b1660448201526080850135605882015260780160408051808303601f19018152908290528051602090910120600254631c13568560e31b8352600483018290529093506001600160a01b03169150819063e09ab4289060240160206040518083038186803b15801561155c57600080fd5b505afa158015611570573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115949190612927565b151560011461144e5760405162461bcd60e51b815260206004820152601960248201527f627269646765207769746864726177206e6f74206578697374000000000000006044820152606401610460565b60036115f46020860186612875565b6006811115611605576116056125a0565b148061162e5750600461161b6020860186612875565b600681111561162c5761162c6125a0565b145b156118b35761164360608501604086016122a4565b61165360808601606087016122a4565b608086013561166860408801602089016122a4565b61167860e0890160c08a01612808565b604051606095861b6bffffffffffffffffffffffff19908116602083015294861b851660348201526048810193909352931b909116606882015260c09190911b6001600160c01b031916607c82015260e0850135608482015260a40160408051601f198184030181529190528051602090910120915060036116fd6020860186612875565b600681111561170e5761170e6125a0565b14156117e557506003546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e64725906024015b60206040518083038186803b15801561175c57600080fd5b505afa158015611770573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117949190612927565b151560011461144e5760405162461bcd60e51b815260206004820152601560248201527f6d696e74207265636f7264206e6f7420657869737400000000000000000000006044820152606401610460565b50600480546040516301e6472560e01b81529182018390526001600160a01b03169081906301e647259060240160206040518083038186803b15801561182a57600080fd5b505afa15801561183e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118629190612927565b151560011461144e5760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f74206578697374000000000000006044820152606401610460565b60056118c26020860186612875565b60068111156118d3576118d36125a0565b14806118fc575060066118e96020860186612875565b60068111156118fa576118fa6125a0565b145b15611b255760056119106020860186612875565b6006811115611921576119216125a0565b141561193957506005546001600160a01b0316611947565b506006546001600160a01b03165b61195760608501604086016122a4565b61196760808601606087016122a4565b608086013561197c60408801602089016122a4565b61198c60e0890160c08a01612808565b604051606095861b6bffffffffffffffffffffffff19908116602083015294861b85166034820152604881019390935290841b8316606883015260c01b6001600160c01b031916607c82015260e087013560848201529183901b1660a482015260b80160408051601f19818403018152919052805160209091012091506005611a186020860186612875565b6006811115611a2957611a296125a0565b1415611a5c576040516301e6472560e01b8152600481018390526001600160a01b038216906301e6472590602401611744565b6040516301e6472560e01b8152600481018390526001600160a01b038216906301e647259060240160206040518083038186803b158015611a9c57600080fd5b505afa158015611ab0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ad49190612927565b1515600114611b255760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f74206578697374000000000000006044820152606401610460565b60008183604051602001611b3b93929190612960565b6040516020818303038152906040528051906020012092505050919050565b6000805a9050600080611b7360608801604089016122a4565b6001600160a01b031634631f34afff60e21b611b9560408b0160208c016122a4565b611ba560808c0160608d016122a4565b60808c0135611bba60e08e0160c08f01612808565b8c8c33604051602401611bd39796959493929190612991565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909416939093179092529051611c3e91906129ea565b60006040518083038185875af1925050503d8060008114611c7b576040519150601f19603f3d011682016040523d82523d6000602084013e611c80565b606091505b50915091508115611ca95780806020019051810190611c9f9190612a06565b9350505050611cbb565b611cb383826121a7565b600093505050505b9392505050565b6000805a9050600080611cdb60608801604089016122a4565b6001600160a01b031634632d5bd7e360e11b611b9560408b0160208c016122a4565b611d0d60608201604083016122a4565b6001600160a01b03167fa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d60008585611d4b60e0870160c08801612808565b866101000135604051611d62959493929190612a27565b60405180910390a2505050565b6001826006811115611d8357611d836125a0565b1415611e24576002546001600160a01b031663cdd1b25d611da48380612896565b611db160208601866128dd565b611dbe60408801886128dd565b611dcb60608a018a6128dd565b6040518963ffffffff1660e01b8152600401611dee989796959493929190612a65565b600060405180830381600087803b158015611e0857600080fd5b505af1158015611e1c573d6000803e3d6000fd5b505050505050565b6002826006811115611e3857611e386125a0565b1415611e59576002546001600160a01b031663a21a9280611da48380612896565b6003826006811115611e6d57611e6d6125a0565b1415611e8e576003546001600160a01b031663f8734302611da48380612896565b6005826006811115611ea257611ea26125a0565b1415611f64576005546001600160a01b031663f8734302611ec38380612896565b611ed060208601866128dd565b611edd60408801886128dd565b611eea60608a018a6128dd565b6040518963ffffffff1660e01b8152600401611f0d989796959493929190612a65565b602060405180830381600087803b158015611f2757600080fd5b505af1158015611f3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f5f9190612ac5565b505050565b6004826006811115611f7857611f786125a0565b1415611f99576004546001600160a01b031663a21a9280611da48380612896565b6006826006811115611fad57611fad6125a0565b141561084f576006546001600160a01b031663a21a9280611ec38380612896565b60006001611fdf60208601866122a4565b611fef60408701602088016122a4565b611fff6060880160408901612808565b876060013546888860405160200161201e989796959493929190612ade565b6040516020818303038152906040528051906020012090509392505050565b6000805a905060008061205660408801602089016122a4565b6001600160a01b031634639c649fdf60e01b61207560208b018b6122a4565b61208560608c0160408d01612808565b8a8a33604051602401611bd3959493929190612b52565b6120ac60408201602083016122a4565b6001600160a01b03167fa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d600185856120ea6060870160408801612808565b8660600135604051611d62959493929190612a27565b6000805a905060008061211960608801604089016122a4565b6001600160a01b0316346305e5a4c160e11b61213b60808b0160608c016122a4565b8a608001358a8a33604051602401611bd3959493929190612b9b565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005a90506000600754456121bc9190612bc5565b905080841080156121d757506121d3604085612bea565b8211155b156121de57fe5b7fffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f61220884612223565b6040516122159190612c0c565b60405180910390a150505050565b606060448251101561226857505060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c79000000602082015290565b600482019150818060200190518101906122829190612c35565b92915050565b80356001600160a01b038116811461229f57600080fd5b919050565b6000602082840312156122b657600080fd5b611cbb82612288565b60008083601f8401126122d157600080fd5b50813567ffffffffffffffff8111156122e957600080fd5b60208301915083602082850101111561230157600080fd5b9250929050565b60008083601f84011261231a57600080fd5b50813567ffffffffffffffff81111561233257600080fd5b6020830191508360208260051b850101111561230157600080fd5b6000806000806000806000806000898b036101a081121561236d57600080fd5b8a3567ffffffffffffffff8082111561238557600080fd5b6123918e838f016122bf565b909c509a508a9150610120601f19840112156123ac57600080fd5b60208d0199506101408d01359250808311156123c757600080fd5b6123d38e848f01612308565b90995097506101608d01359250889150808311156123f057600080fd5b6123fc8e848f01612308565b90975095506101808d013592508691508083111561241957600080fd5b50506124278c828d01612308565b915080935050809150509295985092959850929598565b60006080828403121561245057600080fd5b50919050565b6000806040838503121561246957600080fd5b823567ffffffffffffffff8082111561248157600080fd5b61248d8683870161243e565b935060208501359150808211156124a357600080fd5b5083016101a081860312156124b757600080fd5b809150509250929050565b6000602082840312156124d457600080fd5b5035919050565b60008060008060008060008060006101008a8c0312156124fa57600080fd5b893567ffffffffffffffff8082111561251257600080fd5b61251e8d838e016122bf565b909b5099508991506125338d60208e0161243e565b985060a08c013591508082111561254957600080fd5b6125558d838e01612308565b909850965060c08c013591508082111561256e57600080fd5b61257a8d838e01612308565b909650945060e08c013591508082111561259357600080fd5b506124278c828d01612308565b634e487b7160e01b600052602160045260246000fd5b600581106125c6576125c66125a0565b9052565b6020810161228282846125b6565b8581528460208201528284604083013760409201918201526060019392505050565b60005b838110156126155781810151838201526020016125fd565b83811115612624576000848401525b50505050565b600081518084526126428160208601602086016125fa565b601f01601f19169290920160200192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008383855260208086019550808560051b8301018460005b8781101561270757848303601f19018952813536889003601e190181126126be57600080fd5b8701803567ffffffffffffffff8111156126d757600080fd5b8036038913156126e657600080fd5b6126f38582888501612656565b9a86019a9450505090830190600101612698565b5090979650505050505050565b8183526000602080850194508260005b85811015612750576001600160a01b0361273d83612288565b1687529582019590820190600101612724565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561278d57600080fd5b8260051b8083602087013760009401602001938452509192915050565b6080815260006127bd608083018a61262a565b82810360208401526127d081898b61267f565b905082810360408401526127e5818789612714565b905082810360608401526127fa81858761275b565b9a9950505050505050505050565b60006020828403121561281a57600080fd5b813567ffffffffffffffff81168114611cbb57600080fd5b600281106125c6576125c66125a0565b608081016128508287612832565b84602083015267ffffffffffffffff8416604083015282606083015295945050505050565b60006020828403121561288757600080fd5b813560078110611cbb57600080fd5b6000808335601e198436030181126128ad57600080fd5b83018035915067ffffffffffffffff8211156128c857600080fd5b60200191503681900382131561230157600080fd5b6000808335601e198436030181126128f457600080fd5b83018035915067ffffffffffffffff82111561290f57600080fd5b6020019150600581901b360382131561230157600080fd5b60006020828403121561293957600080fd5b81518015158114611cbb57600080fd5b60028110612959576129596125a0565b60f81b9052565b61296a8185612949565b60609290921b6bffffffffffffffffffffffff191660018301526015820152603501919050565b60006001600160a01b03808a168352808916602084015287604084015267ffffffffffffffff8716606084015260c060808401526129d360c084018688612656565b915080841660a08401525098975050505050505050565b600082516129fc8184602087016125fa565b9190910192915050565b600060208284031215612a1857600080fd5b815160038110611cbb57600080fd5b60a08101612a358288612832565b856020830152612a4860408301866125b6565b67ffffffffffffffff939093166060820152608001529392505050565b608081526000612a79608083018a8c612656565b8281036020840152612a8c81898b61267f565b90508281036040840152612aa1818789612714565b90508281036060840152612ab681858761275b565b9b9a5050505050505050505050565b600060208284031215612ad757600080fd5b5051919050565b612ae8818a612949565b60006bffffffffffffffffffffffff19808a60601b166001840152808960601b166015840152506001600160c01b0319808860c01b166029840152866031840152808660c01b16605184015250828460598401375060009101605901908152979650505050505050565b60006001600160a01b03808816835267ffffffffffffffff8716602084015260806040840152612b86608084018688612656565b91508084166060840152509695505050505050565b60006001600160a01b03808816835286602084015260806040840152612b86608084018688612656565b600082821015612be557634e487b7160e01b600052601160045260246000fd5b500390565b600082612c0757634e487b7160e01b600052601260045260246000fd5b500490565b602081526000611cbb602083018461262a565b634e487b7160e01b600052604160045260246000fd5b600060208284031215612c4757600080fd5b815167ffffffffffffffff80821115612c5f57600080fd5b818401915084601f830112612c7357600080fd5b815181811115612c8557612c85612c1f565b604051601f8201601f19908116603f01168101908382118183101715612cad57612cad612c1f565b81604052828152876020848701011115612cc657600080fd5b612cd78360208301602088016125fa565b97965050505050505056fea26469706673582212200031b7f62e7cdd537c9f1c6a7bc54690204627ffe646ae01693546e7e16aa5f864736f6c63430008090033",
}

// MessageBusReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusReceiverMetaData.ABI instead.
var MessageBusReceiverABI = MessageBusReceiverMetaData.ABI

// MessageBusReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusReceiverMetaData.Bin instead.
var MessageBusReceiverBin = MessageBusReceiverMetaData.Bin

// DeployMessageBusReceiver deploys a new Ethereum contract, binding an instance of MessageBusReceiver to it.
func DeployMessageBusReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (common.Address, *types.Transaction, *MessageBusReceiver, error) {
	parsed, err := MessageBusReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusReceiverBin), backend, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusReceiver{MessageBusReceiverCaller: MessageBusReceiverCaller{contract: contract}, MessageBusReceiverTransactor: MessageBusReceiverTransactor{contract: contract}, MessageBusReceiverFilterer: MessageBusReceiverFilterer{contract: contract}}, nil
}

// MessageBusReceiver is an auto generated Go binding around an Ethereum contract.
type MessageBusReceiver struct {
	MessageBusReceiverCaller     // Read-only binding to the contract
	MessageBusReceiverTransactor // Write-only binding to the contract
	MessageBusReceiverFilterer   // Log filterer for contract events
}

// MessageBusReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusReceiverSession struct {
	Contract     *MessageBusReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessageBusReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusReceiverCallerSession struct {
	Contract *MessageBusReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MessageBusReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusReceiverTransactorSession struct {
	Contract     *MessageBusReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MessageBusReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusReceiverRaw struct {
	Contract *MessageBusReceiver // Generic contract binding to access the raw methods on
}

// MessageBusReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusReceiverCallerRaw struct {
	Contract *MessageBusReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusReceiverTransactorRaw struct {
	Contract *MessageBusReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusReceiver creates a new instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiver(address common.Address, backend bind.ContractBackend) (*MessageBusReceiver, error) {
	contract, err := bindMessageBusReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiver{MessageBusReceiverCaller: MessageBusReceiverCaller{contract: contract}, MessageBusReceiverTransactor: MessageBusReceiverTransactor{contract: contract}, MessageBusReceiverFilterer: MessageBusReceiverFilterer{contract: contract}}, nil
}

// NewMessageBusReceiverCaller creates a new read-only instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverCaller(address common.Address, caller bind.ContractCaller) (*MessageBusReceiverCaller, error) {
	contract, err := bindMessageBusReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverCaller{contract: contract}, nil
}

// NewMessageBusReceiverTransactor creates a new write-only instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusReceiverTransactor, error) {
	contract, err := bindMessageBusReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverTransactor{contract: contract}, nil
}

// NewMessageBusReceiverFilterer creates a new log filterer instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusReceiverFilterer, error) {
	contract, err := bindMessageBusReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverFilterer{contract: contract}, nil
}

// bindMessageBusReceiver binds a generic wrapper to an already deployed contract.
func bindMessageBusReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusReceiver *MessageBusReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusReceiver.Contract.MessageBusReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusReceiver *MessageBusReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.MessageBusReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusReceiver *MessageBusReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.MessageBusReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusReceiver *MessageBusReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusReceiver *MessageBusReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusReceiver *MessageBusReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.contract.Transact(opts, method, params...)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverCaller) ExecutedMessages(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBusReceiver.Contract.ExecutedMessages(&_MessageBusReceiver.CallOpts, arg0)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBusReceiver.Contract.ExecutedMessages(&_MessageBusReceiver.CallOpts, arg0)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) LiquidityBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.LiquidityBridge(&_MessageBusReceiver.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) LiquidityBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.LiquidityBridge(&_MessageBusReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) Owner() (common.Address, error) {
	return _MessageBusReceiver.Contract.Owner(&_MessageBusReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) Owner() (common.Address, error) {
	return _MessageBusReceiver.Contract.Owner(&_MessageBusReceiver.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridge(&_MessageBusReceiver.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridge(&_MessageBusReceiver.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegBridgeV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridgeV2(&_MessageBusReceiver.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegBridgeV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridgeV2(&_MessageBusReceiver.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegVault() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVault(&_MessageBusReceiver.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegVault() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVault(&_MessageBusReceiver.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegVaultV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVaultV2(&_MessageBusReceiver.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegVaultV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVaultV2(&_MessageBusReceiver.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBusReceiver *MessageBusReceiverCaller) PreExecuteMessageGasUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "preExecuteMessageGasUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBusReceiver *MessageBusReceiverSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MessageBusReceiver.Contract.PreExecuteMessageGasUsage(&_MessageBusReceiver.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MessageBusReceiver.Contract.PreExecuteMessageGasUsage(&_MessageBusReceiver.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessage(&_MessageBusReceiver.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessage(&_MessageBusReceiver.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransfer(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransfer(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransferRefund(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransferRefund(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) RefundAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "refundAndExecuteMsg", _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.RefundAndExecuteMsg(&_MessageBusReceiver.TransactOpts, _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.RefundAndExecuteMsg(&_MessageBusReceiver.TransactOpts, _transferParams, _msgParams)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetLiquidityBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setLiquidityBridge", _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetLiquidityBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetLiquidityBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegBridge", _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegBridgeV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegBridgeV2", _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridgeV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridgeV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegVault(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegVault", _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVault(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVault(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegVaultV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegVaultV2", _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVaultV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVaultV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPreExecuteMessageGasUsage(opts *bind.TransactOpts, _usage *big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPreExecuteMessageGasUsage", _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPreExecuteMessageGasUsage(&_MessageBusReceiver.TransactOpts, _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPreExecuteMessageGasUsage(&_MessageBusReceiver.TransactOpts, _usage)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) TransferAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "transferAndExecuteMsg", _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferAndExecuteMsg(&_MessageBusReceiver.TransactOpts, _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferAndExecuteMsg(&_MessageBusReceiver.TransactOpts, _transferParams, _msgParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferOwnership(&_MessageBusReceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferOwnership(&_MessageBusReceiver.TransactOpts, newOwner)
}

// MessageBusReceiverCallRevertedIterator is returned from FilterCallReverted and is used to iterate over the raw logs and unpacked data for CallReverted events raised by the MessageBusReceiver contract.
type MessageBusReceiverCallRevertedIterator struct {
	Event *MessageBusReceiverCallReverted // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverCallRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverCallReverted)
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
		it.Event = new(MessageBusReceiverCallReverted)
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
func (it *MessageBusReceiverCallRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverCallRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverCallReverted represents a CallReverted event raised by the MessageBusReceiver contract.
type MessageBusReceiverCallReverted struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallReverted is a free log retrieval operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterCallReverted(opts *bind.FilterOpts) (*MessageBusReceiverCallRevertedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverCallRevertedIterator{contract: _MessageBusReceiver.contract, event: "CallReverted", logs: logs, sub: sub}, nil
}

// WatchCallReverted is a free log subscription operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchCallReverted(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverCallReverted) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverCallReverted)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "CallReverted", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseCallReverted(log types.Log) (*MessageBusReceiverCallReverted, error) {
	event := new(MessageBusReceiverCallReverted)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "CallReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MessageBusReceiver contract.
type MessageBusReceiverExecutedIterator struct {
	Event *MessageBusReceiverExecuted // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverExecuted)
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
		it.Event = new(MessageBusReceiverExecuted)
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
func (it *MessageBusReceiverExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverExecuted represents a Executed event raised by the MessageBusReceiver contract.
type MessageBusReceiverExecuted struct {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterExecuted(opts *bind.FilterOpts, receiver []common.Address) (*MessageBusReceiverExecutedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverExecutedIterator{contract: _MessageBusReceiver.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverExecuted, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverExecuted)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseExecuted(log types.Log) (*MessageBusReceiverExecuted, error) {
	event := new(MessageBusReceiverExecuted)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverLiquidityBridgeUpdatedIterator is returned from FilterLiquidityBridgeUpdated and is used to iterate over the raw logs and unpacked data for LiquidityBridgeUpdated events raised by the MessageBusReceiver contract.
type MessageBusReceiverLiquidityBridgeUpdatedIterator struct {
	Event *MessageBusReceiverLiquidityBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverLiquidityBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverLiquidityBridgeUpdated)
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
		it.Event = new(MessageBusReceiverLiquidityBridgeUpdated)
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
func (it *MessageBusReceiverLiquidityBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverLiquidityBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverLiquidityBridgeUpdated represents a LiquidityBridgeUpdated event raised by the MessageBusReceiver contract.
type MessageBusReceiverLiquidityBridgeUpdated struct {
	LiquidityBridge common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityBridgeUpdated is a free log retrieval operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterLiquidityBridgeUpdated(opts *bind.FilterOpts) (*MessageBusReceiverLiquidityBridgeUpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverLiquidityBridgeUpdatedIterator{contract: _MessageBusReceiver.contract, event: "LiquidityBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityBridgeUpdated is a free log subscription operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchLiquidityBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverLiquidityBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverLiquidityBridgeUpdated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseLiquidityBridgeUpdated(log types.Log) (*MessageBusReceiverLiquidityBridgeUpdated, error) {
	event := new(MessageBusReceiverLiquidityBridgeUpdated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverNeedRetryIterator is returned from FilterNeedRetry and is used to iterate over the raw logs and unpacked data for NeedRetry events raised by the MessageBusReceiver contract.
type MessageBusReceiverNeedRetryIterator struct {
	Event *MessageBusReceiverNeedRetry // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverNeedRetryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverNeedRetry)
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
		it.Event = new(MessageBusReceiverNeedRetry)
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
func (it *MessageBusReceiverNeedRetryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverNeedRetryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverNeedRetry represents a NeedRetry event raised by the MessageBusReceiver contract.
type MessageBusReceiverNeedRetry struct {
	MsgType    uint8
	MsgId      [32]byte
	SrcChainId uint64
	SrcTxHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNeedRetry is a free log retrieval operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterNeedRetry(opts *bind.FilterOpts) (*MessageBusReceiverNeedRetryIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverNeedRetryIterator{contract: _MessageBusReceiver.contract, event: "NeedRetry", logs: logs, sub: sub}, nil
}

// WatchNeedRetry is a free log subscription operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchNeedRetry(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverNeedRetry) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverNeedRetry)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "NeedRetry", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseNeedRetry(log types.Log) (*MessageBusReceiverNeedRetry, error) {
	event := new(MessageBusReceiverNeedRetry)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "NeedRetry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusReceiver contract.
type MessageBusReceiverOwnershipTransferredIterator struct {
	Event *MessageBusReceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverOwnershipTransferred)
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
		it.Event = new(MessageBusReceiverOwnershipTransferred)
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
func (it *MessageBusReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusReceiver contract.
type MessageBusReceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusReceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverOwnershipTransferredIterator{contract: _MessageBusReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverOwnershipTransferred)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusReceiverOwnershipTransferred, error) {
	event := new(MessageBusReceiverOwnershipTransferred)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverPegBridgeUpdatedIterator is returned from FilterPegBridgeUpdated and is used to iterate over the raw logs and unpacked data for PegBridgeUpdated events raised by the MessageBusReceiver contract.
type MessageBusReceiverPegBridgeUpdatedIterator struct {
	Event *MessageBusReceiverPegBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverPegBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverPegBridgeUpdated)
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
		it.Event = new(MessageBusReceiverPegBridgeUpdated)
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
func (it *MessageBusReceiverPegBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverPegBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverPegBridgeUpdated represents a PegBridgeUpdated event raised by the MessageBusReceiver contract.
type MessageBusReceiverPegBridgeUpdated struct {
	PegBridge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeUpdated is a free log retrieval operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterPegBridgeUpdated(opts *bind.FilterOpts) (*MessageBusReceiverPegBridgeUpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverPegBridgeUpdatedIterator{contract: _MessageBusReceiver.contract, event: "PegBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeUpdated is a free log subscription operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchPegBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverPegBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverPegBridgeUpdated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParsePegBridgeUpdated(log types.Log) (*MessageBusReceiverPegBridgeUpdated, error) {
	event := new(MessageBusReceiverPegBridgeUpdated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverPegBridgeV2UpdatedIterator is returned from FilterPegBridgeV2Updated and is used to iterate over the raw logs and unpacked data for PegBridgeV2Updated events raised by the MessageBusReceiver contract.
type MessageBusReceiverPegBridgeV2UpdatedIterator struct {
	Event *MessageBusReceiverPegBridgeV2Updated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverPegBridgeV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverPegBridgeV2Updated)
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
		it.Event = new(MessageBusReceiverPegBridgeV2Updated)
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
func (it *MessageBusReceiverPegBridgeV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverPegBridgeV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverPegBridgeV2Updated represents a PegBridgeV2Updated event raised by the MessageBusReceiver contract.
type MessageBusReceiverPegBridgeV2Updated struct {
	PegBridgeV2 common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeV2Updated is a free log retrieval operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterPegBridgeV2Updated(opts *bind.FilterOpts) (*MessageBusReceiverPegBridgeV2UpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverPegBridgeV2UpdatedIterator{contract: _MessageBusReceiver.contract, event: "PegBridgeV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeV2Updated is a free log subscription operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchPegBridgeV2Updated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverPegBridgeV2Updated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverPegBridgeV2Updated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParsePegBridgeV2Updated(log types.Log) (*MessageBusReceiverPegBridgeV2Updated, error) {
	event := new(MessageBusReceiverPegBridgeV2Updated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverPegVaultUpdatedIterator is returned from FilterPegVaultUpdated and is used to iterate over the raw logs and unpacked data for PegVaultUpdated events raised by the MessageBusReceiver contract.
type MessageBusReceiverPegVaultUpdatedIterator struct {
	Event *MessageBusReceiverPegVaultUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverPegVaultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverPegVaultUpdated)
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
		it.Event = new(MessageBusReceiverPegVaultUpdated)
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
func (it *MessageBusReceiverPegVaultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverPegVaultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverPegVaultUpdated represents a PegVaultUpdated event raised by the MessageBusReceiver contract.
type MessageBusReceiverPegVaultUpdated struct {
	PegVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPegVaultUpdated is a free log retrieval operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterPegVaultUpdated(opts *bind.FilterOpts) (*MessageBusReceiverPegVaultUpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverPegVaultUpdatedIterator{contract: _MessageBusReceiver.contract, event: "PegVaultUpdated", logs: logs, sub: sub}, nil
}

// WatchPegVaultUpdated is a free log subscription operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchPegVaultUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverPegVaultUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverPegVaultUpdated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParsePegVaultUpdated(log types.Log) (*MessageBusReceiverPegVaultUpdated, error) {
	event := new(MessageBusReceiverPegVaultUpdated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverPegVaultV2UpdatedIterator is returned from FilterPegVaultV2Updated and is used to iterate over the raw logs and unpacked data for PegVaultV2Updated events raised by the MessageBusReceiver contract.
type MessageBusReceiverPegVaultV2UpdatedIterator struct {
	Event *MessageBusReceiverPegVaultV2Updated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverPegVaultV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverPegVaultV2Updated)
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
		it.Event = new(MessageBusReceiverPegVaultV2Updated)
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
func (it *MessageBusReceiverPegVaultV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverPegVaultV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverPegVaultV2Updated represents a PegVaultV2Updated event raised by the MessageBusReceiver contract.
type MessageBusReceiverPegVaultV2Updated struct {
	PegVaultV2 common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPegVaultV2Updated is a free log retrieval operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterPegVaultV2Updated(opts *bind.FilterOpts) (*MessageBusReceiverPegVaultV2UpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverPegVaultV2UpdatedIterator{contract: _MessageBusReceiver.contract, event: "PegVaultV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegVaultV2Updated is a free log subscription operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchPegVaultV2Updated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverPegVaultV2Updated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverPegVaultV2Updated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParsePegVaultV2Updated(log types.Log) (*MessageBusReceiverPegVaultV2Updated, error) {
	event := new(MessageBusReceiverPegVaultV2Updated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMetaData contains all meta data concerning the MessageBusSender contract.
var MessageBusSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBase\",\"type\":\"uint256\"}],\"name\":\"FeeBaseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"}],\"name\":\"FeePerByteUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageWithTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeeBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeePerByte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610ff9380380610ff983398101604081905261002f91610099565b61003833610049565b6001600160a01b03166080526100c9565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100ab57600080fd5b81516001600160a01b03811681146100c257600080fd5b9392505050565b608051610f0e6100eb600039600081816101ee01526103c50152610f0e6000f3fe6080604052600436106100c75760003560e01c806395e911a811610074578063e2c1ed251161004e578063e2c1ed2514610210578063f2fde38b14610230578063f60bbe2a1461025057600080fd5b806395e911a8146101b35780639f3ce55a146101c9578063ccf2683b146101dc57600080fd5b80635335dca2116100a55780635335dca2146101215780635b3e5f50146101545780638da5cb5b1461018157600080fd5b806306c28bd6146100cc5780632ff4c411146100ee5780634289fbb31461010e575b600080fd5b3480156100d857600080fd5b506100ec6100e736600461095a565b610266565b005b3480156100fa57600080fd5b506100ec6101093660046109db565b610310565b6100ec61011c366004610ad1565b610573565b34801561012d57600080fd5b5061014161013c366004610b49565b61065c565b6040519081526020015b60405180910390f35b34801561016057600080fd5b5061014161016f366004610b8b565b60036020526000908152604090205481565b34801561018d57600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161014b565b3480156101bf57600080fd5b5061014160015481565b6100ec6101d7366004610ba6565b610680565b3480156101e857600080fd5b5061019b7f000000000000000000000000000000000000000000000000000000000000000081565b34801561021c57600080fd5b506100ec61022b36600461095a565b610763565b34801561023c57600080fd5b506100ec61024b366004610b8b565b610801565b34801561025c57600080fd5b5061014160025481565b336102796000546001600160a01b031690565b6001600160a01b0316146102d45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b60018190556040518181527f892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2906020015b60405180910390a150565b6000463060405160200161036692919091825260601b6bffffffffffffffffffffffff191660208201527f77697468647261774665650000000000000000000000000000000000000000006034820152603f0190565b60408051808303601f19018152828252805160209182012090830181905260608c901b6bffffffffffffffffffffffff19168383015260548084018c9052825180850390910181526074840192839052633416de1160e11b90925292507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163682dbc229161040c918b908b908b908b908b908b90607801610d50565b60006040518083038186803b15801561042457600080fd5b505afa158015610438573d6000803e3d6000fd5b505050506001600160a01b03891660009081526003602052604081205461045f908a610e02565b9050600081116104b15760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f2077697468647261770000000000000060448201526064016102cb565b6001600160a01b038a166000818152600360205260408082208c90555190919061c35090849084818181858888f193505050503d8060008114610510576040519150601f19603f3d011682016040523d82523d6000602084013e610515565b606091505b50509050806105665760405162461bcd60e51b815260206004820152601660248201527f6661696c656420746f207769746864726177206665650000000000000000000060448201526064016102cb565b5050505050505050505050565b468514156105b55760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9259608a1b60448201526064016102cb565b60006105c1838361065c565b9050803410156106065760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016102cb565b336001600160a01b03167f172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f668888888888883460405161064b9796959493929190610e19565b60405180910390a250505050505050565b60025460009061066c9083610e66565b6001546106799190610e85565b9392505050565b468314156106c25760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9259608a1b60448201526064016102cb565b60006106ce838361065c565b9050803410156107135760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016102cb565b336001600160a01b03167fce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e48686868634604051610754959493929190610e9d565b60405180910390a25050505050565b336107766000546001600160a01b031690565b6001600160a01b0316146107cc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102cb565b60028190556040518181527f210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c90602001610305565b336108146000546001600160a01b031690565b6001600160a01b03161461086a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102cb565b6001600160a01b0381166108e65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102cb565b6108ef816108f2565b50565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561096c57600080fd5b5035919050565b80356001600160a01b038116811461098a57600080fd5b919050565b60008083601f8401126109a157600080fd5b50813567ffffffffffffffff8111156109b957600080fd5b6020830191508360208260051b85010111156109d457600080fd5b9250929050565b60008060008060008060008060a0898b0312156109f757600080fd5b610a0089610973565b975060208901359650604089013567ffffffffffffffff80821115610a2457600080fd5b610a308c838d0161098f565b909850965060608b0135915080821115610a4957600080fd5b610a558c838d0161098f565b909650945060808b0135915080821115610a6e57600080fd5b50610a7b8b828c0161098f565b999c989b5096995094979396929594505050565b60008083601f840112610aa157600080fd5b50813567ffffffffffffffff811115610ab957600080fd5b6020830191508360208285010111156109d457600080fd5b60008060008060008060a08789031215610aea57600080fd5b610af387610973565b955060208701359450610b0860408801610973565b935060608701359250608087013567ffffffffffffffff811115610b2b57600080fd5b610b3789828a01610a8f565b979a9699509497509295939492505050565b60008060208385031215610b5c57600080fd5b823567ffffffffffffffff811115610b7357600080fd5b610b7f85828601610a8f565b90969095509350505050565b600060208284031215610b9d57600080fd5b61067982610973565b60008060008060608587031215610bbc57600080fd5b610bc585610973565b935060208501359250604085013567ffffffffffffffff811115610be857600080fd5b610bf487828801610a8f565b95989497509550505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260006020808501808196508560051b810191508460005b87811015610cad5782840389528135601e19883603018112610c6457600080fd5b8701803567ffffffffffffffff811115610c7d57600080fd5b803603891315610c8c57600080fd5b610c998682898501610c00565b9a87019a9550505090840190600101610c43565b5091979650505050505050565b8183526000602080850194508260005b85811015610cf6576001600160a01b03610ce383610973565b1687529582019590820190600101610cca565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610d3357600080fd5b8260051b8083602087013760009401602001938452509192915050565b608081526000885180608084015260005b81811015610d7e576020818c0181015160a0868401015201610d61565b81811115610d9057600060a083860101525b50601f01601f1916820182810360a09081016020850152610db4908201898b610c29565b90508281036040840152610dc9818789610cba565b90508281036060840152610dde818587610d01565b9a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610e1457610e14610dec565b500390565b60006001600160a01b03808a16835288602084015280881660408401525085606083015260c06080830152610e5260c083018587610c00565b90508260a083015298975050505050505050565b6000816000190483118215151615610e8057610e80610dec565b500290565b60008219821115610e9857610e98610dec565b500190565b6001600160a01b0386168152846020820152608060408201526000610ec6608083018587610c00565b9050826060830152969550505050505056fea26469706673582212201b96b7ada5c44a077b369fa6b5bb8231951d92e50ab8cebc945188ba0fbd691c64736f6c63430008090033",
}

// MessageBusSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusSenderMetaData.ABI instead.
var MessageBusSenderABI = MessageBusSenderMetaData.ABI

// MessageBusSenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusSenderMetaData.Bin instead.
var MessageBusSenderBin = MessageBusSenderMetaData.Bin

// DeployMessageBusSender deploys a new Ethereum contract, binding an instance of MessageBusSender to it.
func DeployMessageBusSender(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address) (common.Address, *types.Transaction, *MessageBusSender, error) {
	parsed, err := MessageBusSenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusSenderBin), backend, _sigsVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusSender{MessageBusSenderCaller: MessageBusSenderCaller{contract: contract}, MessageBusSenderTransactor: MessageBusSenderTransactor{contract: contract}, MessageBusSenderFilterer: MessageBusSenderFilterer{contract: contract}}, nil
}

// MessageBusSender is an auto generated Go binding around an Ethereum contract.
type MessageBusSender struct {
	MessageBusSenderCaller     // Read-only binding to the contract
	MessageBusSenderTransactor // Write-only binding to the contract
	MessageBusSenderFilterer   // Log filterer for contract events
}

// MessageBusSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusSenderSession struct {
	Contract     *MessageBusSender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageBusSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusSenderCallerSession struct {
	Contract *MessageBusSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MessageBusSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusSenderTransactorSession struct {
	Contract     *MessageBusSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MessageBusSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusSenderRaw struct {
	Contract *MessageBusSender // Generic contract binding to access the raw methods on
}

// MessageBusSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusSenderCallerRaw struct {
	Contract *MessageBusSenderCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusSenderTransactorRaw struct {
	Contract *MessageBusSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusSender creates a new instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSender(address common.Address, backend bind.ContractBackend) (*MessageBusSender, error) {
	contract, err := bindMessageBusSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusSender{MessageBusSenderCaller: MessageBusSenderCaller{contract: contract}, MessageBusSenderTransactor: MessageBusSenderTransactor{contract: contract}, MessageBusSenderFilterer: MessageBusSenderFilterer{contract: contract}}, nil
}

// NewMessageBusSenderCaller creates a new read-only instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderCaller(address common.Address, caller bind.ContractCaller) (*MessageBusSenderCaller, error) {
	contract, err := bindMessageBusSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderCaller{contract: contract}, nil
}

// NewMessageBusSenderTransactor creates a new write-only instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusSenderTransactor, error) {
	contract, err := bindMessageBusSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderTransactor{contract: contract}, nil
}

// NewMessageBusSenderFilterer creates a new log filterer instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusSenderFilterer, error) {
	contract, err := bindMessageBusSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderFilterer{contract: contract}, nil
}

// bindMessageBusSender binds a generic wrapper to an already deployed contract.
func bindMessageBusSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusSenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusSender *MessageBusSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusSender.Contract.MessageBusSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusSender *MessageBusSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSender.Contract.MessageBusSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusSender *MessageBusSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusSender.Contract.MessageBusSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusSender *MessageBusSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusSender *MessageBusSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusSender *MessageBusSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusSender.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBusSender.Contract.CalcFee(&_MessageBusSender.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBusSender.Contract.CalcFee(&_MessageBusSender.CallOpts, _message)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) FeeBase() (*big.Int, error) {
	return _MessageBusSender.Contract.FeeBase(&_MessageBusSender.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) FeeBase() (*big.Int, error) {
	return _MessageBusSender.Contract.FeeBase(&_MessageBusSender.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) FeePerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "feePerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) FeePerByte() (*big.Int, error) {
	return _MessageBusSender.Contract.FeePerByte(&_MessageBusSender.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) FeePerByte() (*big.Int, error) {
	return _MessageBusSender.Contract.FeePerByte(&_MessageBusSender.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderSession) Owner() (common.Address, error) {
	return _MessageBusSender.Contract.Owner(&_MessageBusSender.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderCallerSession) Owner() (common.Address, error) {
	return _MessageBusSender.Contract.Owner(&_MessageBusSender.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderSession) SigsVerifier() (common.Address, error) {
	return _MessageBusSender.Contract.SigsVerifier(&_MessageBusSender.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderCallerSession) SigsVerifier() (common.Address, error) {
	return _MessageBusSender.Contract.SigsVerifier(&_MessageBusSender.CallOpts)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) WithdrawnFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "withdrawnFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBusSender.Contract.WithdrawnFees(&_MessageBusSender.CallOpts, arg0)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBusSender.Contract.WithdrawnFees(&_MessageBusSender.CallOpts, arg0)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactor) SendMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessage(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessage(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessageWithTransfer(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessageWithTransfer(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactor) SetFeeBase(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "setFeeBase", _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeeBase(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeeBase(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactor) SetFeePerByte(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "setFeePerByte", _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeePerByte(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeePerByte(&_MessageBusSender.TransactOpts, _fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.Contract.TransferOwnership(&_MessageBusSender.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.Contract.TransferOwnership(&_MessageBusSender.TransactOpts, newOwner)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.WithdrawFee(&_MessageBusSender.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.WithdrawFee(&_MessageBusSender.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// MessageBusSenderFeeBaseUpdatedIterator is returned from FilterFeeBaseUpdated and is used to iterate over the raw logs and unpacked data for FeeBaseUpdated events raised by the MessageBusSender contract.
type MessageBusSenderFeeBaseUpdatedIterator struct {
	Event *MessageBusSenderFeeBaseUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderFeeBaseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderFeeBaseUpdated)
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
		it.Event = new(MessageBusSenderFeeBaseUpdated)
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
func (it *MessageBusSenderFeeBaseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderFeeBaseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderFeeBaseUpdated represents a FeeBaseUpdated event raised by the MessageBusSender contract.
type MessageBusSenderFeeBaseUpdated struct {
	FeeBase *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeBaseUpdated is a free log retrieval operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBusSender *MessageBusSenderFilterer) FilterFeeBaseUpdated(opts *bind.FilterOpts) (*MessageBusSenderFeeBaseUpdatedIterator, error) {

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "FeeBaseUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderFeeBaseUpdatedIterator{contract: _MessageBusSender.contract, event: "FeeBaseUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeBaseUpdated is a free log subscription operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBusSender *MessageBusSenderFilterer) WatchFeeBaseUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusSenderFeeBaseUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "FeeBaseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderFeeBaseUpdated)
				if err := _MessageBusSender.contract.UnpackLog(event, "FeeBaseUpdated", log); err != nil {
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

// ParseFeeBaseUpdated is a log parse operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBusSender *MessageBusSenderFilterer) ParseFeeBaseUpdated(log types.Log) (*MessageBusSenderFeeBaseUpdated, error) {
	event := new(MessageBusSenderFeeBaseUpdated)
	if err := _MessageBusSender.contract.UnpackLog(event, "FeeBaseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderFeePerByteUpdatedIterator is returned from FilterFeePerByteUpdated and is used to iterate over the raw logs and unpacked data for FeePerByteUpdated events raised by the MessageBusSender contract.
type MessageBusSenderFeePerByteUpdatedIterator struct {
	Event *MessageBusSenderFeePerByteUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderFeePerByteUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderFeePerByteUpdated)
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
		it.Event = new(MessageBusSenderFeePerByteUpdated)
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
func (it *MessageBusSenderFeePerByteUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderFeePerByteUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderFeePerByteUpdated represents a FeePerByteUpdated event raised by the MessageBusSender contract.
type MessageBusSenderFeePerByteUpdated struct {
	FeePerByte *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeePerByteUpdated is a free log retrieval operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBusSender *MessageBusSenderFilterer) FilterFeePerByteUpdated(opts *bind.FilterOpts) (*MessageBusSenderFeePerByteUpdatedIterator, error) {

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "FeePerByteUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderFeePerByteUpdatedIterator{contract: _MessageBusSender.contract, event: "FeePerByteUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePerByteUpdated is a free log subscription operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBusSender *MessageBusSenderFilterer) WatchFeePerByteUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusSenderFeePerByteUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "FeePerByteUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderFeePerByteUpdated)
				if err := _MessageBusSender.contract.UnpackLog(event, "FeePerByteUpdated", log); err != nil {
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

// ParseFeePerByteUpdated is a log parse operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBusSender *MessageBusSenderFilterer) ParseFeePerByteUpdated(log types.Log) (*MessageBusSenderFeePerByteUpdated, error) {
	event := new(MessageBusSenderFeePerByteUpdated)
	if err := _MessageBusSender.contract.UnpackLog(event, "FeePerByteUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the MessageBusSender contract.
type MessageBusSenderMessageIterator struct {
	Event *MessageBusSenderMessage // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderMessage)
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
		it.Event = new(MessageBusSenderMessage)
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
func (it *MessageBusSenderMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderMessage represents a Message event raised by the MessageBusSender contract.
type MessageBusSenderMessage struct {
	Sender     common.Address
	Receiver   common.Address
	DstChainId *big.Int
	Message    []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) FilterMessage(opts *bind.FilterOpts, sender []common.Address) (*MessageBusSenderMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderMessageIterator{contract: _MessageBusSender.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *MessageBusSenderMessage, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderMessage)
				if err := _MessageBusSender.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) ParseMessage(log types.Log) (*MessageBusSenderMessage, error) {
	event := new(MessageBusSenderMessage)
	if err := _MessageBusSender.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMessageWithTransferIterator is returned from FilterMessageWithTransfer and is used to iterate over the raw logs and unpacked data for MessageWithTransfer events raised by the MessageBusSender contract.
type MessageBusSenderMessageWithTransferIterator struct {
	Event *MessageBusSenderMessageWithTransfer // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderMessageWithTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderMessageWithTransfer)
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
		it.Event = new(MessageBusSenderMessageWithTransfer)
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
func (it *MessageBusSenderMessageWithTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderMessageWithTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderMessageWithTransfer represents a MessageWithTransfer event raised by the MessageBusSender contract.
type MessageBusSenderMessageWithTransfer struct {
	Sender        common.Address
	Receiver      common.Address
	DstChainId    *big.Int
	Bridge        common.Address
	SrcTransferId [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageWithTransfer is a free log retrieval operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) FilterMessageWithTransfer(opts *bind.FilterOpts, sender []common.Address) (*MessageBusSenderMessageWithTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderMessageWithTransferIterator{contract: _MessageBusSender.contract, event: "MessageWithTransfer", logs: logs, sub: sub}, nil
}

// WatchMessageWithTransfer is a free log subscription operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) WatchMessageWithTransfer(opts *bind.WatchOpts, sink chan<- *MessageBusSenderMessageWithTransfer, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderMessageWithTransfer)
				if err := _MessageBusSender.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
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

// ParseMessageWithTransfer is a log parse operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) ParseMessageWithTransfer(log types.Log) (*MessageBusSenderMessageWithTransfer, error) {
	event := new(MessageBusSenderMessageWithTransfer)
	if err := _MessageBusSender.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusSender contract.
type MessageBusSenderOwnershipTransferredIterator struct {
	Event *MessageBusSenderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderOwnershipTransferred)
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
		it.Event = new(MessageBusSenderOwnershipTransferred)
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
func (it *MessageBusSenderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusSender contract.
type MessageBusSenderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSender *MessageBusSenderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusSenderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderOwnershipTransferredIterator{contract: _MessageBusSender.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSender *MessageBusSenderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusSenderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderOwnershipTransferred)
				if err := _MessageBusSender.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBusSender *MessageBusSenderFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusSenderOwnershipTransferred, error) {
	event := new(MessageBusSenderOwnershipTransferred)
	if err := _MessageBusSender.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgDataTypesMetaData contains all meta data concerning the MsgDataTypes contract.
var MsgDataTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122074cf73930ba179cb6e15925e2ed77dd8f4eca8cf22384a7c1ab87d7e28fa616264736f6c63430008090033",
}

// MsgDataTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use MsgDataTypesMetaData.ABI instead.
var MsgDataTypesABI = MsgDataTypesMetaData.ABI

// MsgDataTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MsgDataTypesMetaData.Bin instead.
var MsgDataTypesBin = MsgDataTypesMetaData.Bin

// DeployMsgDataTypes deploys a new Ethereum contract, binding an instance of MsgDataTypes to it.
func DeployMsgDataTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MsgDataTypes, error) {
	parsed, err := MsgDataTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MsgDataTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MsgDataTypes{MsgDataTypesCaller: MsgDataTypesCaller{contract: contract}, MsgDataTypesTransactor: MsgDataTypesTransactor{contract: contract}, MsgDataTypesFilterer: MsgDataTypesFilterer{contract: contract}}, nil
}

// MsgDataTypes is an auto generated Go binding around an Ethereum contract.
type MsgDataTypes struct {
	MsgDataTypesCaller     // Read-only binding to the contract
	MsgDataTypesTransactor // Write-only binding to the contract
	MsgDataTypesFilterer   // Log filterer for contract events
}

// MsgDataTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MsgDataTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MsgDataTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MsgDataTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MsgDataTypesSession struct {
	Contract     *MsgDataTypes     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MsgDataTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MsgDataTypesCallerSession struct {
	Contract *MsgDataTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MsgDataTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MsgDataTypesTransactorSession struct {
	Contract     *MsgDataTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MsgDataTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MsgDataTypesRaw struct {
	Contract *MsgDataTypes // Generic contract binding to access the raw methods on
}

// MsgDataTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MsgDataTypesCallerRaw struct {
	Contract *MsgDataTypesCaller // Generic read-only contract binding to access the raw methods on
}

// MsgDataTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MsgDataTypesTransactorRaw struct {
	Contract *MsgDataTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMsgDataTypes creates a new instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypes(address common.Address, backend bind.ContractBackend) (*MsgDataTypes, error) {
	contract, err := bindMsgDataTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypes{MsgDataTypesCaller: MsgDataTypesCaller{contract: contract}, MsgDataTypesTransactor: MsgDataTypesTransactor{contract: contract}, MsgDataTypesFilterer: MsgDataTypesFilterer{contract: contract}}, nil
}

// NewMsgDataTypesCaller creates a new read-only instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesCaller(address common.Address, caller bind.ContractCaller) (*MsgDataTypesCaller, error) {
	contract, err := bindMsgDataTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesCaller{contract: contract}, nil
}

// NewMsgDataTypesTransactor creates a new write-only instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*MsgDataTypesTransactor, error) {
	contract, err := bindMsgDataTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesTransactor{contract: contract}, nil
}

// NewMsgDataTypesFilterer creates a new log filterer instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*MsgDataTypesFilterer, error) {
	contract, err := bindMsgDataTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesFilterer{contract: contract}, nil
}

// bindMsgDataTypes binds a generic wrapper to an already deployed contract.
func bindMsgDataTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MsgDataTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgDataTypes *MsgDataTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgDataTypes.Contract.MsgDataTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgDataTypes *MsgDataTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.MsgDataTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgDataTypes *MsgDataTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.MsgDataTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgDataTypes *MsgDataTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgDataTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgDataTypes *MsgDataTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgDataTypes *MsgDataTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.contract.Transact(opts, method, params...)
}
