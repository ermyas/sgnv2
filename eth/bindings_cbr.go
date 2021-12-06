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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DelayThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedTransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelayedTransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EpochLengthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EpochVolumeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MaxSendUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinAddUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinSendUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"}],\"name\":\"Relay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resetTime\",\"type\":\"uint256\"}],\"name\":\"ResetNotification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxSlippage\",\"type\":\"uint32\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"SignersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"refid\",\"type\":\"bytes32\"}],\"name\":\"WithdrawDone\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addNativeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addseq\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delayThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumeCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"executeDelayedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"increaseNoticePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastOpTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalMaxSlippage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyResetSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_relayRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"resetSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"sendNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setDelayThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"name\":\"setEpochVolumeCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMaxSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minimalMaxSlippage\",\"type\":\"uint32\"}],\"name\":\"setMinimalMaxSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"setWrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ssHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_triggerTime\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_newSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_curSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_curPowers\",\"type\":\"uint256[]\"}],\"name\":\"updateSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_wdmsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdraws\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001d3362000048565b60016005556006805460ff19169055620000373362000098565b620000423362000162565b62000222565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526007602052604090205460ff1615620001075760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064015b60405180910390fd5b6001600160a01b038116600081815260076020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f891015b60405180910390a150565b6001600160a01b03811660009081526008602052604090205460ff1615620001cd5760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f7200000000006044820152606401620000fe565b6001600160a01b038116600081815260086020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5910162000157565b61505e80620002326000396000f3fe60806040526004361061036f5760003560e01c806382dc1ec4116101c6578063ba2cb25c116100f7578063e43581b811610095578063f20c922a1161006f578063f20c922a14610acf578063f2fde38b14610aef578063f832138314610b0f578063f8b30d7d14610b3c57600080fd5b8063e43581b814610a56578063e999e5f414610a8f578063eecdac8814610aaf57600080fd5b8063d0790da9116100d1578063d0790da9146109cb578063e026049c146109e1578063e09ab428146109f6578063e3eece2614610a2657600080fd5b8063ba2cb25c1461095e578063ccde517a1461097e578063cdd1b25d146109ab57600080fd5b80639ff9001a11610164578063a7bdf45a1161013e578063a7bdf45a14610881578063adc0d57f146108a1578063b1c94d941461091b578063b5f2bc471461093157600080fd5b80639ff9001a14610821578063a21a928014610841578063a5977fbb1461086157600080fd5b806389e39127116101a057806389e39127146107935780638da5cb5b146107cd5780639b14d4c6146107eb5780639e25fc5c1461080157600080fd5b806382dc1ec41461073e5780638456cb591461075e578063878fe1ce1461077357600080fd5b806352532faa116102a057806365a114f11161023e5780636ef8d66d116102185780636ef8d66d146106d15780637044c89e146106e6578063715018a6146106f957806380f51c121461070e57600080fd5b806365a114f11461067b578063682dbc22146106915780636b2c0f55146106b157600080fd5b806357d775f81161027a57806357d775f8146105f35780635c975abb1461060957806360216b0014610621578063618ee0551461064e57600080fd5b806352532faa1461058657806354eea796146105b357806356688700146105d357600080fd5b80633d5721071161030d578063457bfa2f116102e7578063457bfa2f146104d557806346fbf68e1461050d57806347b16c6c14610546578063482341261461056657600080fd5b80633d5721071461048d5780633f2e5fc3146104ad5780633f4ba83a146104c057600080fd5b80632fd1b0a4116103495780632fd1b0a4146103d2578063370fb47b146104095780633c4a25d01461042d5780633c64f04b1461044d57600080fd5b8063089927411461037b57806317bdbae51461039d57806325c38b9f146103bd57600080fd5b3661037657005b600080fd5b34801561038757600080fd5b5061039b6103963660046147b9565b610b69565b005b3480156103a957600080fd5b5061039b6103b83660046147b9565b610d0c565b3480156103c957600080fd5b5061039b610ea3565b3480156103de57600080fd5b506017546103ef9063ffffffff1681565b60405163ffffffff90911681526020015b60405180910390f35b34801561041557600080fd5b5061041f60025481565b604051908152602001610400565b34801561043957600080fd5b5061039b610448366004614841565b610f33565b34801561045957600080fd5b5061047d61046836600461485c565b60146020526000908152604090205460ff1681565b6040519015158152602001610400565b34801561049957600080fd5b5061039b6104a836600461485c565b610f87565b61039b6104bb3660046148a1565b61101b565b3480156104cc57600080fd5b5061039b611271565b3480156104e157600080fd5b506013546104f5906001600160a01b031681565b6040516001600160a01b039091168152602001610400565b34801561051957600080fd5b5061047d610528366004614841565b6001600160a01b031660009081526007602052604090205460ff1690565b34801561055257600080fd5b5061039b6105613660046147b9565b6112da565b34801561057257600080fd5b5061039b6105813660046148ff565b611471565b34801561059257600080fd5b5061041f6105a1366004614841565b600e6020526000908152604090205481565b3480156105bf57600080fd5b5061039b6105ce36600461485c565b6114e5565b3480156105df57600080fd5b5061039b6105ee36600461491a565b611572565b3480156105ff57600080fd5b5061041f60095481565b34801561061557600080fd5b5060065460ff1661047d565b34801561062d57600080fd5b5061041f61063c366004614841565b600a6020526000908152604090205481565b34801561065a57600080fd5b5061041f610669366004614841565b60166020526000908152604090205481565b34801561068757600080fd5b5061041f60035481565b34801561069d57600080fd5b5061039b6106ac36600461495a565b611734565b3480156106bd57600080fd5b5061039b6106cc366004614841565b611820565b3480156106dd57600080fd5b5061039b611871565b61039b6106f436600461485c565b61187a565b34801561070557600080fd5b5061039b611b2c565b34801561071a57600080fd5b5061047d610729366004614841565b60076020526000908152604090205460ff1681565b34801561074a57600080fd5b5061039b610759366004614841565b611b7e565b34801561076a57600080fd5b5061039b611bcf565b34801561077f57600080fd5b5061039b61078e3660046147b9565b611c36565b34801561079f57600080fd5b506010546107b49067ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610400565b3480156107d957600080fd5b506000546001600160a01b03166104f5565b3480156107f757600080fd5b5061041f60045481565b34801561080d57600080fd5b5061039b61081c36600461485c565b611dcd565b34801561082d57600080fd5b5061039b61083c366004614841565b611e3b565b34801561084d57600080fd5b5061039b61085c366004614a88565b611ea5565b34801561086d57600080fd5b5061039b61087c366004614b77565b6121ec565b34801561088d57600080fd5b5061039b61089c3660046147b9565b61233a565b3480156108ad57600080fd5b506108f06108bc36600461485c565b600d6020526000908152604090208054600182015460028301546003909301546001600160a01b0392831693919092169184565b604080516001600160a01b039586168152949093166020850152918301526060820152608001610400565b34801561092757600080fd5b5061041f600f5481565b34801561093d57600080fd5b5061041f61094c366004614841565b600b6020526000908152604090205481565b34801561096a57600080fd5b5061039b610979366004614be4565b6123eb565b34801561098a57600080fd5b5061041f610999366004614841565b60116020526000908152604090205481565b3480156109b757600080fd5b5061039b6109c6366004614a88565b612541565b3480156109d757600080fd5b5061041f60015481565b3480156109ed57600080fd5b5061039b612866565b348015610a0257600080fd5b5061047d610a1136600461485c565b60126020526000908152604090205460ff1681565b348015610a3257600080fd5b5061047d610a41366004614841565b60086020526000908152604090205460ff1681565b348015610a6257600080fd5b5061047d610a71366004614841565b6001600160a01b031660009081526008602052604090205460ff1690565b348015610a9b57600080fd5b5061039b610aaa3660046147b9565b61286f565b348015610abb57600080fd5b5061039b610aca366004614841565b612a06565b348015610adb57600080fd5b5061039b610aea36600461485c565b612a57565b348015610afb57600080fd5b5061039b610b0a366004614841565b612b01565b348015610b1b57600080fd5b5061041f610b2a366004614841565b600c6020526000908152604090205481565b348015610b4857600080fd5b5061041f610b57366004614841565b60156020526000908152604090205481565b3360009081526008602052604090205460ff16610bc65760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b60448201526064015b60405180910390fd5b828114610c075760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610bbd565b60005b83811015610d0557828282818110610c2457610c24614ce4565b9050602002013560156000878785818110610c4157610c41614ce4565b9050602002016020810190610c569190614841565b6001600160a01b031681526020810191909152604001600020557f8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9858583818110610ca357610ca3614ce4565b9050602002016020810190610cb89190614841565b848484818110610cca57610cca614ce4565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180610cfd81614d10565b915050610c0a565b5050505050565b3360009081526008602052604090205460ff16610d645760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610bbd565b828114610da55760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610bbd565b60005b83811015610d0557828282818110610dc257610dc2614ce4565b90506020020135600e6000878785818110610ddf57610ddf614ce4565b9050602002016020810190610df49190614841565b6001600160a01b031681526020810191909152604001600020557fceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce858583818110610e4157610e41614ce4565b9050602002016020810190610e569190614841565b848484818110610e6857610e68614ce4565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180610e9b81614d10565b915050610da8565b6000546001600160a01b03163314610eeb5760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b600454610ef89042614d2b565b60038190556040519081527f68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1906020015b60405180910390a1565b6000546001600160a01b03163314610f7b5760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b610f8481612bce565b50565b3360009081526008602052604090205460ff16610fdf5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610bbd565b600f8190556040518181527fc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6906020015b60405180910390a150565b6002600554141561106e5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610bbd565b600260055560065460ff16156110b95760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610bbd565b8334146110fa5760405162461bcd60e51b815260206004820152600f60248201526e082dadeeadce840dad2e6dac2e8c6d608b1b6044820152606401610bbd565b6013546001600160a01b03166111525760405162461bcd60e51b815260206004820152601360248201527f4e61746976652077726170206e6f7420736574000000000000000000000000006044820152606401610bbd565b6013546000906111709087906001600160a01b031687878787612c8b565b9050601360009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0866040518263ffffffff1660e01b81526004016000604051808303818588803b1580156111c257600080fd5b505af11580156111d6573d6000803e3d6000fd5b5050601354604080518681523360208201526001600160a01b03808d1692820192909252911660608201526080810189905267ffffffffffffffff80891660a0830152871660c082015263ffffffff861660e08201527f89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01935061010001915061125c9050565b60405180910390a15050600160055550505050565b3360009081526007602052604090205460ff166112d05760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f74207061757365720000000000000000000000006044820152606401610bbd565b6112d8612ebe565b565b3360009081526008602052604090205460ff166113325760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610bbd565b8281146113735760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610bbd565b60005b83811015610d055782828281811061139057611390614ce4565b90506020020135600b60008787858181106113ad576113ad614ce4565b90506020020160208101906113c29190614841565b6001600160a01b031681526020810191909152604001600020557f608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e8985858381811061140f5761140f614ce4565b90506020020160208101906114249190614841565b84848481811061143657611436614ce4565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a18061146981614d10565b915050611376565b3360009081526008602052604090205460ff166114c95760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610bbd565b6017805463ffffffff191663ffffffff92909216919091179055565b3360009081526008602052604090205460ff1661153d5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610bbd565b60098190556040518181527f2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b390602001611010565b600260055414156115c55760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610bbd565b600260055560065460ff16156116105760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610bbd565b6001600160a01b038216600090815260116020526040902054811161166a5760405162461bcd60e51b815260206004820152601060248201526f185b5bdd5b9d081d1bdbc81cdb585b1b60821b6044820152606401610bbd565b601080546001919060009061168a90849067ffffffffffffffff16614d43565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506116d0333083856001600160a01b0316612f55909392919063ffffffff16565b6010546040805167ffffffffffffffff90921682523360208301526001600160a01b0384168282015260608201839052517fd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce7649181900360800190a150506001600555565b60008484848460405160200161174d9493929190614ddb565b60405160208183030381529060405280519060200120905080600154146117b65760405162461bcd60e51b815260206004820152601860248201527f4d69736d617463682063757272656e74207369676e65727300000000000000006044820152606401610bbd565b87516020808a0191909120604080517f19457468657265756d205369676e6564204d6573736167653a0a33320000000081850152603c8082019390935281518082039093018352605c019052805191012061181690888888888888612fed565b5050505050505050565b6000546001600160a01b031633146118685760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b610f8481613323565b6112d833613323565b600260055414156118cd5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610bbd565b600260055560065460ff16156119185760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610bbd565b8034146119595760405162461bcd60e51b815260206004820152600f60248201526e082dadeeadce840dad2e6dac2e8c6d608b1b6044820152606401610bbd565b6013546001600160a01b03166119b15760405162461bcd60e51b815260206004820152601360248201527f4e61746976652077726170206e6f7420736574000000000000000000000000006044820152606401610bbd565b6013546001600160a01b03166000908152601160205260409020548111611a0d5760405162461bcd60e51b815260206004820152601060248201526f185b5bdd5b9d081d1bdbc81cdb585b1b60821b6044820152606401610bbd565b6010805460019190600090611a2d90849067ffffffffffffffff16614d43565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550601360009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b158015611aa357600080fd5b505af1158015611ab7573d6000803e3d6000fd5b50506010546013546040805167ffffffffffffffff90931683523360208401526001600160a01b0390911690820152606081018590527fd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce76493506080019150611b1c9050565b60405180910390a1506001600555565b6000546001600160a01b03163314611b745760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b6112d860006133dc565b6000546001600160a01b03163314611bc65760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b610f848161342c565b3360009081526007602052604090205460ff16611c2e5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f74207061757365720000000000000000000000006044820152606401610bbd565b6112d86134e9565b3360009081526008602052604090205460ff16611c8e5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610bbd565b828114611ccf5760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610bbd565b60005b83811015610d0557828282818110611cec57611cec614ce4565b9050602002013560166000878785818110611d0957611d09614ce4565b9050602002016020810190611d1e9190614841565b6001600160a01b031681526020810191909152604001600020557f4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c858583818110611d6b57611d6b614ce4565b9050602002016020810190611d809190614841565b848484818110611d9257611d92614ce4565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180611dc581614d10565b915050611cd2565b60065460ff1615611e135760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610bbd565b6000611e1e82613564565b9050611e37816000015182602001518360400151613729565b5050565b6000546001600160a01b03163314611e835760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b601380546001600160a01b0319166001600160a01b0392909216919091179055565b60065460ff1615611eeb5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610bbd565b60004630604051602001611f4192919091825260601b6bffffffffffffffffffffffff191660208201527f57697468647261774d73670000000000000000000000000000000000000000006034820152603f0190565b604051602081830303815290604052805190602001209050611f8b818a8a604051602001611f7193929190614df2565b604051602081830303815290604052888888888888611734565b6000611fcc8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061385e92505050565b905060008160000151826020015183604001518460600151856080015160405160200161204595949392919060c095861b6001600160c01b031990811682529490951b9093166008850152606091821b6bffffffffffffffffffffffff199081166010860152911b166024830152603882015260580190565b60408051601f1981840301815291815281516020928301206000818152601290935291205490915060ff16156120bd5760405162461bcd60e51b815260206004820152601a60248201527f776974686472617720616c7265616479207375636365656465640000000000006044820152606401610bbd565b6000818152601260205260409020805460ff19166001179055606082015160808301516120ea91906139be565b60608201516001600160a01b03166000908152600e602052604090205480158015906121195750808360800151115b1561213b5761213682846040015185606001518660800151613ad6565b612152565b612152836040015184606001518560800151613729565b7f48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be8284602001518560400151866060015187608001518860a001516040516121d69695949392919095865267ffffffffffffffff9490941660208601526001600160a01b03928316604086015291166060840152608083015260a082015260c00190565b60405180910390a1505050505050505050505050565b6002600554141561223f5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610bbd565b600260055560065460ff161561228a5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610bbd565b600061229a878787878787612c8b565b90506122b16001600160a01b038716333088612f55565b604080518281523360208201526001600160a01b0389811682840152881660608201526080810187905267ffffffffffffffff86811660a0830152851660c082015263ffffffff841660e082015290517f89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01918190036101000190a1505060016005555050505050565b6000546001600160a01b031633146123825760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b60035442116123d35760405162461bcd60e51b815260206004820152601460248201527f6e6f742072656163682072657365742074696d650000000000000000000000006044820152606401610bbd565b6000196003556123e584848484613be9565b50505050565b6002548b1161243c5760405162461bcd60e51b815260206004820152601e60248201527f547269676765722074696d65206973206e6f7420696e6372656173696e6700006044820152606401610bbd565b61244842610e10614d2b565b8b106124965760405162461bcd60e51b815260206004820152601960248201527f547269676765722074696d6520697320746f6f206c61726765000000000000006044820152606401610bbd565b600046306040516020016124ec92919091825260601b6bffffffffffffffffffffffff191660208201527f5570646174655369676e65727300000000000000000000000000000000000000603482015260410190565b604051602081830303815290604052805190602001209050612522818d8d8d8d8d604051602001611f7196959493929190614e0c565b61252e8b8b8b8b613be9565b5050506002989098555050505050505050565b60065460ff16156125875760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610bbd565b600046306040516020016125dd92919091825260601b6bffffffffffffffffffffffff191660208201527f52656c6179000000000000000000000000000000000000000000000000000000603482015260390190565b60405160208183030381529060405280519060200120905061260d818a8a604051602001611f7193929190614df2565b600061264e8a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613d9392505050565b8051602080830151604080850151606080870151608088015160a089015160c0808b015187519a861b6bffffffffffffffffffffffff199081168c8c015298861b891660348c01529590941b9096166048890152605c880191909152811b6001600160c01b0319908116607c88015293901b9092166084850152608c808501929092528051808503909201825260ac909301835280519082012060008181526014909252919020549192509060ff161561273c5760405162461bcd60e51b815260206004820152600f60248201526e7472616e736665722065786973747360881b6044820152606401610bbd565b60008181526014602052604090819020805460ff19166001179055820151606083015161276991906139be565b6040808301516001600160a01b03166000908152600e602052205480158015906127965750808360600151115b156127b8576127b382846020015185604001518660600151613ad6565b6127cf565b6127cf836020015184604001518560600151613729565b7f79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c82846000015185602001518660400151876060015188608001518960c001516040516121d697969594939291909687526001600160a01b0395861660208801529385166040870152919093166060850152608084019290925267ffffffffffffffff9190911660a083015260c082015260e00190565b6112d833613f0b565b3360009081526008602052604090205460ff166128c75760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610bbd565b8281146129085760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610bbd565b60005b83811015610d055782828281811061292557612925614ce4565b905060200201356011600087878581811061294257612942614ce4565b90506020020160208101906129579190614841565b6001600160a01b031681526020810191909152604001600020557fc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f8585838181106129a4576129a4614ce4565b90506020020160208101906129b99190614841565b8484848181106129cb576129cb614ce4565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a1806129fe81614d10565b91505061290b565b6000546001600160a01b03163314612a4e5760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b610f8481613f0b565b6000546001600160a01b03163314612a9f5760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b6004548111612afc5760405162461bcd60e51b815260206004820152602360248201527f6e6f7469636520706572696f642063616e206f6e6c7920626520696e637265616044820152621cd95960ea1b6064820152608401610bbd565b600455565b6000546001600160a01b03163314612b495760405162461bcd60e51b815260206004820181905260248201526000805160206150098339815191526044820152606401610bbd565b6001600160a01b038116612bc55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610bbd565b610f84816133dc565b6001600160a01b03811660009081526008602052604090205460ff1615612c375760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f7200000000006044820152606401610bbd565b6001600160a01b038116600081815260086020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b59101611010565b6001600160a01b0385166000908152601560205260408120548511612ce55760405162461bcd60e51b815260206004820152601060248201526f185b5bdd5b9d081d1bdbc81cdb585b1b60821b6044820152606401610bbd565b6001600160a01b0386166000908152601660205260409020541580612d2257506001600160a01b0386166000908152601660205260409020548511155b612d6e5760405162461bcd60e51b815260206004820152601060248201527f616d6f756e7420746f6f206c61726765000000000000000000000000000000006044820152606401610bbd565b60175463ffffffff90811690831611612dc95760405162461bcd60e51b815260206004820152601660248201527f6d617820736c69707061676520746f6f20736d616c6c000000000000000000006044820152606401610bbd565b6040516bffffffffffffffffffffffff1933606090811b8216602084015289811b8216603484015288901b166048820152605c81018690526001600160c01b031960c086811b8216607c84015285811b8216608484015246901b16608c82015260009060940160408051601f1981840301815291815281516020928301206000818152601490935291205490915060ff1615612e995760405162461bcd60e51b815260206004820152600f60248201526e7472616e736665722065786973747360881b6044820152606401610bbd565b6000818152601460205260409020805460ff1916600117905590509695505050505050565b60065460ff16612f105760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610bbd565b6006805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001610f29565b6040516001600160a01b03808516602483015283166044820152606481018290526123e59085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613fc4565b8281146130485760405162461bcd60e51b815260206004820152602360248201527f7369676e65727320616e6420706f77657273206c656e677468206e6f74206d616044820152620e8c6d60eb1b6064820152608401610bbd565b6000805b8481101561308c5783838281811061306657613066614ce4565b90506020020135826130789190614d2b565b91508061308481614d10565b91505061304c565b506000600361309c836002614e34565b6130a69190614e53565b6130b1906001614d2b565b905060008080805b8a8110156132d157600061313c8d8d848181106130d8576130d8614ce4565b90506020028101906130ea9190614e75565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508f6140a990919063ffffffff16565b9050836001600160a01b0316816001600160a01b03161161319f5760405162461bcd60e51b815260206004820152601e60248201527f7369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610bbd565b8093505b8a8a848181106131b5576131b5614ce4565b90506020020160208101906131ca9190614841565b6001600160a01b0316816001600160a01b03161115613244576131ee600184614d2b565b925089831061323f5760405162461bcd60e51b815260206004820152601060248201527f7369676e6572206e6f7420666f756e64000000000000000000000000000000006044820152606401610bbd565b6131a3565b8a8a8481811061325657613256614ce4565b905060200201602081019061326b9190614841565b6001600160a01b0316816001600160a01b031614156132ab5788888481811061329657613296614ce4565b90506020020135856132a89190614d2b565b94505b8585106132be575050505050505061331a565b50806132c981614d10565b9150506130b9565b5060405162461bcd60e51b815260206004820152601260248201527f71756f72756d206e6f74207265616368656400000000000000000000000000006044820152606401610bbd565b50505050505050565b6001600160a01b03811660009081526007602052604090205460ff1661338b5760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f742070617573657200000000000000000000006044820152606401610bbd565b6001600160a01b038116600081815260076020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e9101611010565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526007602052604090205460ff16156134955760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c726561647920706175736572000000000000006044820152606401610bbd565b6001600160a01b038116600081815260076020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101611010565b60065460ff161561352f5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610bbd565b6006805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612f3d3390565b6040805160808101825260008082526020820181905291810182905260608101919091526000828152600d6020908152604091829020825160808101845281546001600160a01b03908116825260018301541692810192909252600281015492820192909252600390910154606082018190526136235760405162461bcd60e51b815260206004820152601a60248201527f64656c61796564207472616e73666572206e6f742065786973740000000000006044820152606401610bbd565b600f5481606001516136359190614d2b565b42116136835760405162461bcd60e51b815260206004820152601d60248201527f64656c61796564207472616e73666572207374696c6c206c6f636b65640000006044820152606401610bbd565b6000838152600d6020908152604080832080546001600160a01b03199081168255600182018054909116905560028101849055600301929092558251908301518383015192517f3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d44269361371b93889390929091909384526001600160a01b03928316602085015291166040830152606082015260800190565b60405180910390a192915050565b6013546001600160a01b038381169116141561384557601354604051632e1a7d4d60e01b8152600481018390526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b15801561378557600080fd5b505af1158015613799573d6000803e3d6000fd5b505050506000836001600160a01b03168261c35090604051600060405180830381858888f193505050503d80600081146137ef576040519150601f19603f3d011682016040523d82523d6000602084013e6137f4565b606091505b50509050806123e55760405162461bcd60e51b815260206004820152601b60248201527f6661696c656420746f2073656e64206e617469766520746f6b656e00000000006044820152606401610bbd565b6138596001600160a01b0383168483614153565b505050565b6040805160c08101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905283518085019094528184528301849052909190805b602083015151835110156139b6576138bc83614183565b909250905081600114156138e4576138d3836141bd565b67ffffffffffffffff1684526138a5565b816002141561390a576138f6836141bd565b67ffffffffffffffff1660208501526138a5565b81600314156139375761392461391f8461423f565b6142fc565b6001600160a01b031660408501526138a5565b816004141561395f5761394c61391f8461423f565b6001600160a01b031660608501526138a5565b8160051415613983576139796139748461423f565b614307565b60808501526138a5565b81600614156139a75761399d6139988461423f565b61433e565b60a08501526138a5565b6139b18382614356565b6138a5565b505050919050565b6009546139c9575050565b6001600160a01b0382166000908152600b6020526040902054806139ec57505050565b6001600160a01b0383166000908152600a602052604081205460095490914291613a168184614e53565b613a209190614e34565b6001600160a01b0387166000908152600c6020526040902054909150811115613a4b57849250613a58565b613a558584614d2b565b92505b83831115613aa85760405162461bcd60e51b815260206004820152601260248201527f766f6c756d6520657863656564732063617000000000000000000000000000006044820152606401610bbd565b506001600160a01b039094166000908152600a6020908152604080832093909355600c905220929092555050565b6000848152600d602052604090206003015415613b355760405162461bcd60e51b815260206004820152601f60248201527f64656c61796564207472616e7366657220616c726561647920657869737473006044820152606401610bbd565b604080516080810182526001600160a01b0380861682528481166020808401918252838501868152426060860190815260008b8152600d90935291869020945185549085166001600160a01b031991821617865592516001860180549190951693169290921790925551600283015551600390910155517fcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce690613bdb9086815260200190565b60405180910390a150505050565b828114613c445760405162461bcd60e51b815260206004820152602360248201527f7369676e65727320616e6420706f77657273206c656e677468206e6f74206d616044820152620e8c6d60eb1b6064820152608401610bbd565b6000805b84811015613d1d57816001600160a01b0316868683818110613c6c57613c6c614ce4565b9050602002016020810190613c819190614841565b6001600160a01b031611613ce25760405162461bcd60e51b815260206004820152602260248201527f4e6577207369676e657273206e6f7420696e20617363656e64696e67206f726460448201526132b960f11b6064820152608401610bbd565b858582818110613cf457613cf4614ce4565b9050602002016020810190613d099190614841565b915080613d1581614d10565b915050613c48565b5084848484604051602001613d359493929190614ddb565b60408051601f198184030181529082905280516020909101206001557ff126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f890613d84908790879087908790614ebc565b60405180910390a15050505050565b6040805160e08101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905260c0830182905283518085019094528184528301849052909190805b602083015151835110156139b657613df883614183565b90925090508160011415613e2257613e1261391f8461423f565b6001600160a01b03168452613de1565b8160021415613e4a57613e3761391f8461423f565b6001600160a01b03166020850152613de1565b8160031415613e7257613e5f61391f8461423f565b6001600160a01b03166040850152613de1565b8160041415613e9157613e876139748461423f565b6060850152613de1565b8160051415613eb757613ea3836141bd565b67ffffffffffffffff166080850152613de1565b8160061415613edd57613ec9836141bd565b67ffffffffffffffff1660a0850152613de1565b8160071415613efc57613ef26139988461423f565b60c0850152613de1565b613f068382614356565b613de1565b6001600160a01b03811660009081526008602052604090205460ff16613f735760405162461bcd60e51b815260206004820152601760248201527f4163636f756e74206973206e6f7420676f7665726e6f720000000000000000006044820152606401610bbd565b6001600160a01b038116600081815260086020908152604091829020805460ff1916905590519182527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b9101611010565b6000614019826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166143c89092919063ffffffff16565b80519091501561385957808060200190518101906140379190614f3e565b6138595760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610bbd565b60008151604114156140dd5760208201516040830151606084015160001a6140d3868285856143e1565b935050505061414d565b81516040141561410557602082015160408301516140fc85838361458a565b9250505061414d565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610bbd565b92915050565b6040516001600160a01b03831660248201526044810182905261385990849063a9059cbb60e01b90606401612f89565b6000806000614191846141bd565b905061419e600882614e53565b92508060071660058111156141b5576141b5614f60565b915050915091565b602080820151825181019091015160009182805b600a8110156142395783811a91506141ea816007614e34565b82607f16901b8517945081608016600014156142275761420b816001614d2b565b8651879061421a908390614d2b565b9052509395945050505050565b8061423181614d10565b9150506141d1565b50600080fd5b6060600061424c836141bd565b905060008184600001516142609190614d2b565b905083602001515181111561427457600080fd5b8167ffffffffffffffff81111561428d5761428d614944565b6040519080825280601f01601f1916602001820160405280156142b7576020820181803683370190505b50602080860151865192955091818601919083010160005b858110156142f15781810151838201526142ea602082614d2b565b90506142cf565b505050935250919050565b600061414d826145cd565b600060208251111561431857600080fd5b602082015190508151602061432d9190614f76565b614338906008614e34565b1c919050565b6000815160201461434e57600080fd5b506020015190565b600081600581111561436a5761436a614f60565b141561437957613859826141bd565b600281600581111561438d5761438d614f60565b141561037657600061439e836141bd565b905080836000018181516143b29190614d2b565b9052506020830151518351111561385957600080fd5b60606143d784846000856145f5565b90505b9392505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561445e5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610bbd565b8360ff16601b148061447357508360ff16601c145b6144ca5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610bbd565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa15801561451e573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166145815760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610bbd565b95945050505050565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821660ff83901c601b016145c3868287856143e1565b9695505050505050565b600081516014146145dd57600080fd5b50602001516c01000000000000000000000000900490565b60608247101561466d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610bbd565b843b6146bb5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610bbd565b600080866001600160a01b031685876040516146d79190614fb9565b60006040518083038185875af1925050503d8060008114614714576040519150601f19603f3d011682016040523d82523d6000602084013e614719565b606091505b5091509150614729828286614734565b979650505050505050565b606083156147435750816143da565b8251156147535782518084602001fd5b8160405162461bcd60e51b8152600401610bbd9190614fd5565b60008083601f84011261477f57600080fd5b50813567ffffffffffffffff81111561479757600080fd5b6020830191508360208260051b85010111156147b257600080fd5b9250929050565b600080600080604085870312156147cf57600080fd5b843567ffffffffffffffff808211156147e757600080fd5b6147f38883890161476d565b9096509450602087013591508082111561480c57600080fd5b506148198782880161476d565b95989497509550505050565b80356001600160a01b038116811461483c57600080fd5b919050565b60006020828403121561485357600080fd5b6143da82614825565b60006020828403121561486e57600080fd5b5035919050565b803567ffffffffffffffff8116811461483c57600080fd5b803563ffffffff8116811461483c57600080fd5b600080600080600060a086880312156148b957600080fd5b6148c286614825565b9450602086013593506148d760408701614875565b92506148e560608701614875565b91506148f36080870161488d565b90509295509295909350565b60006020828403121561491157600080fd5b6143da8261488d565b6000806040838503121561492d57600080fd5b61493683614825565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b60008060008060008060006080888a03121561497557600080fd5b873567ffffffffffffffff8082111561498d57600080fd5b818a0191508a601f8301126149a157600080fd5b8135818111156149b3576149b3614944565b604051601f8201601f19908116603f011681019083821181831017156149db576149db614944565b816040528281528d60208487010111156149f457600080fd5b82602086016020830137600094508460208483010152809b5050505060208a013581811115614a21578283fd5b614a2d8c828d0161476d565b90995097505060408a013581811115614a44578283fd5b614a508c828d0161476d565b90975095505060608a013581811115614a67578283fd5b614a738c828d0161476d565b9a9d999c50979a509598949794955050505050565b6000806000806000806000806080898b031215614aa457600080fd5b883567ffffffffffffffff80821115614abc57600080fd5b818b0191508b601f830112614ad057600080fd5b813581811115614adf57600080fd5b8c6020828501011115614af157600080fd5b60209283019a509850908a01359080821115614b0c57600080fd5b614b188c838d0161476d565b909850965060408b0135915080821115614b3157600080fd5b614b3d8c838d0161476d565b909650945060608b0135915080821115614b5657600080fd5b50614b638b828c0161476d565b999c989b5096995094979396929594505050565b60008060008060008060c08789031215614b9057600080fd5b614b9987614825565b9550614ba760208801614825565b945060408701359350614bbc60608801614875565b9250614bca60808801614875565b9150614bd860a0880161488d565b90509295509295509295565b600080600080600080600080600080600060c08c8e031215614c0557600080fd5b8b359a5067ffffffffffffffff8060208e01351115614c2357600080fd5b614c338e60208f01358f0161476d565b909b50995060408d0135811015614c4957600080fd5b614c598e60408f01358f0161476d565b909950975060608d0135811015614c6f57600080fd5b614c7f8e60608f01358f0161476d565b909750955060808d0135811015614c9557600080fd5b614ca58e60808f01358f0161476d565b909550935060a08d0135811015614cbb57600080fd5b50614ccc8d60a08e01358e0161476d565b81935080925050509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415614d2457614d24614cfa565b5060010190565b60008219821115614d3e57614d3e614cfa565b500190565b600067ffffffffffffffff808316818516808303821115614d6657614d66614cfa565b01949350505050565b60008160005b84811015614da4576001600160a01b03614d8e83614825565b1686526020958601959190910190600101614d75565b5093949350505050565b60006001600160fb1b03831115614dc457600080fd5b8260051b8083863760009401938452509192915050565b60006145c3614deb838789614d6f565b8486614dae565b838152818360208301376000910160200190815292915050565b8681528560208201526000614e28614deb604084018789614d6f565b98975050505050505050565b6000816000190483118215151615614e4e57614e4e614cfa565b500290565b600082614e7057634e487b7160e01b600052601260045260246000fd5b500490565b6000808335601e19843603018112614e8c57600080fd5b83018035915067ffffffffffffffff821115614ea757600080fd5b6020019150368190038213156147b257600080fd5b6040808252810184905260008560608301825b87811015614efd576001600160a01b03614ee884614825565b16825260209283019290910190600101614ecf565b5083810360208501528481526001600160fb1b03851115614f1d57600080fd5b8460051b915081866020830137600091016020019081529695505050505050565b600060208284031215614f5057600080fd5b815180151581146143da57600080fd5b634e487b7160e01b600052602160045260246000fd5b600082821015614f8857614f88614cfa565b500390565b60005b83811015614fa8578181015183820152602001614f90565b838111156123e55750506000910152565b60008251614fcb818460208701614f8d565b9190910192915050565b6020815260008251806020840152614ff4816040850160208701614f8d565b601f01601f1916919091016040019291505056fe4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a26469706673582212205382b8d96ac91efd2817a3bab16c7eabe2435603f76e49f893c090661dce3bb864736f6c63430008090033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Bridge *BridgeCaller) Addseq(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "addseq")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Bridge *BridgeSession) Addseq() (uint64, error) {
	return _Bridge.Contract.Addseq(&_Bridge.CallOpts)
}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Bridge *BridgeCallerSession) Addseq() (uint64, error) {
	return _Bridge.Contract.Addseq(&_Bridge.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Bridge *BridgeCaller) DelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "delayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Bridge *BridgeSession) DelayPeriod() (*big.Int, error) {
	return _Bridge.Contract.DelayPeriod(&_Bridge.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Bridge *BridgeCallerSession) DelayPeriod() (*big.Int, error) {
	return _Bridge.Contract.DelayPeriod(&_Bridge.CallOpts)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Bridge *BridgeCaller) DelayThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "delayThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Bridge *BridgeSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.DelayThresholds(&_Bridge.CallOpts, arg0)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.DelayThresholds(&_Bridge.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Bridge *BridgeCaller) DelayedTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "delayedTransfers", arg0)

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
func (_Bridge *BridgeSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Bridge.Contract.DelayedTransfers(&_Bridge.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Bridge *BridgeCallerSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Bridge.Contract.DelayedTransfers(&_Bridge.CallOpts, arg0)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Bridge *BridgeCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Bridge *BridgeSession) EpochLength() (*big.Int, error) {
	return _Bridge.Contract.EpochLength(&_Bridge.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Bridge *BridgeCallerSession) EpochLength() (*big.Int, error) {
	return _Bridge.Contract.EpochLength(&_Bridge.CallOpts)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Bridge *BridgeCaller) EpochVolumeCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "epochVolumeCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Bridge *BridgeSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.EpochVolumeCaps(&_Bridge.CallOpts, arg0)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.EpochVolumeCaps(&_Bridge.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Bridge *BridgeCaller) EpochVolumes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "epochVolumes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Bridge *BridgeSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.EpochVolumes(&_Bridge.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.EpochVolumes(&_Bridge.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Bridge *BridgeCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Bridge *BridgeSession) Governors(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Governors(&_Bridge.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Governors(&_Bridge.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Bridge *BridgeCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Bridge *BridgeSession) IsGovernor(_account common.Address) (bool, error) {
	return _Bridge.Contract.IsGovernor(&_Bridge.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Bridge *BridgeCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _Bridge.Contract.IsGovernor(&_Bridge.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Bridge *BridgeCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Bridge *BridgeSession) IsPauser(account common.Address) (bool, error) {
	return _Bridge.Contract.IsPauser(&_Bridge.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Bridge *BridgeCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Bridge.Contract.IsPauser(&_Bridge.CallOpts, account)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Bridge *BridgeCaller) LastOpTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "lastOpTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Bridge *BridgeSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.LastOpTimestamps(&_Bridge.CallOpts, arg0)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.LastOpTimestamps(&_Bridge.CallOpts, arg0)
}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Bridge *BridgeCaller) MaxSend(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "maxSend", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Bridge *BridgeSession) MaxSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MaxSend(&_Bridge.CallOpts, arg0)
}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) MaxSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MaxSend(&_Bridge.CallOpts, arg0)
}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Bridge *BridgeCaller) MinAdd(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minAdd", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Bridge *BridgeSession) MinAdd(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinAdd(&_Bridge.CallOpts, arg0)
}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) MinAdd(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinAdd(&_Bridge.CallOpts, arg0)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeCaller) MinSend(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minSend", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeSession) MinSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinSend(&_Bridge.CallOpts, arg0)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) MinSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinSend(&_Bridge.CallOpts, arg0)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Bridge *BridgeCaller) MinimalMaxSlippage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minimalMaxSlippage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Bridge *BridgeSession) MinimalMaxSlippage() (uint32, error) {
	return _Bridge.Contract.MinimalMaxSlippage(&_Bridge.CallOpts)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Bridge *BridgeCallerSession) MinimalMaxSlippage() (uint32, error) {
	return _Bridge.Contract.MinimalMaxSlippage(&_Bridge.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Bridge *BridgeCaller) NativeWrap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nativeWrap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Bridge *BridgeSession) NativeWrap() (common.Address, error) {
	return _Bridge.Contract.NativeWrap(&_Bridge.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Bridge *BridgeCallerSession) NativeWrap() (common.Address, error) {
	return _Bridge.Contract.NativeWrap(&_Bridge.CallOpts)
}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Bridge *BridgeCaller) NoticePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "noticePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Bridge *BridgeSession) NoticePeriod() (*big.Int, error) {
	return _Bridge.Contract.NoticePeriod(&_Bridge.CallOpts)
}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Bridge *BridgeCallerSession) NoticePeriod() (*big.Int, error) {
	return _Bridge.Contract.NoticePeriod(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Bridge *BridgeCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Bridge *BridgeSession) Pausers(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Pausers(&_Bridge.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Pausers(&_Bridge.CallOpts, arg0)
}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Bridge *BridgeCaller) ResetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "resetTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Bridge *BridgeSession) ResetTime() (*big.Int, error) {
	return _Bridge.Contract.ResetTime(&_Bridge.CallOpts)
}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Bridge *BridgeCallerSession) ResetTime() (*big.Int, error) {
	return _Bridge.Contract.ResetTime(&_Bridge.CallOpts)
}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Bridge *BridgeCaller) SsHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "ssHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Bridge *BridgeSession) SsHash() ([32]byte, error) {
	return _Bridge.Contract.SsHash(&_Bridge.CallOpts)
}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Bridge *BridgeCallerSession) SsHash() ([32]byte, error) {
	return _Bridge.Contract.SsHash(&_Bridge.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "transfers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) Transfers(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Transfers(&_Bridge.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Transfers(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Transfers(&_Bridge.CallOpts, arg0)
}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Bridge *BridgeCaller) TriggerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "triggerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Bridge *BridgeSession) TriggerTime() (*big.Int, error) {
	return _Bridge.Contract.TriggerTime(&_Bridge.CallOpts)
}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Bridge *BridgeCallerSession) TriggerTime() (*big.Int, error) {
	return _Bridge.Contract.TriggerTime(&_Bridge.CallOpts)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Bridge *BridgeCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "verifySigs", _msg, _sigs, _signers, _powers)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Bridge *BridgeSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _Bridge.Contract.VerifySigs(&_Bridge.CallOpts, _msg, _sigs, _signers, _powers)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Bridge *BridgeCallerSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _Bridge.Contract.VerifySigs(&_Bridge.CallOpts, _msg, _sigs, _signers, _powers)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) Withdraws(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "withdraws", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) Withdraws(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Withdraws(&_Bridge.CallOpts, arg0)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Withdraws(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Withdraws(&_Bridge.CallOpts, arg0)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Bridge *BridgeTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Bridge *BridgeSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddGovernor(&_Bridge.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Bridge *BridgeTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddGovernor(&_Bridge.TransactOpts, _account)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addLiquidity", _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddLiquidity(&_Bridge.TransactOpts, _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddLiquidity(&_Bridge.TransactOpts, _token, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Bridge *BridgeTransactor) AddNativeLiquidity(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addNativeLiquidity", _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Bridge *BridgeSession) AddNativeLiquidity(_amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddNativeLiquidity(&_Bridge.TransactOpts, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Bridge *BridgeTransactorSession) AddNativeLiquidity(_amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddNativeLiquidity(&_Bridge.TransactOpts, _amount)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Bridge *BridgeTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Bridge *BridgeSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddPauser(&_Bridge.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Bridge *BridgeTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddPauser(&_Bridge.TransactOpts, account)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Bridge *BridgeTransactor) ExecuteDelayedTransfer(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeDelayedTransfer", id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Bridge *BridgeSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteDelayedTransfer(&_Bridge.TransactOpts, id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Bridge *BridgeTransactorSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteDelayedTransfer(&_Bridge.TransactOpts, id)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Bridge *BridgeTransactor) IncreaseNoticePeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "increaseNoticePeriod", period)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Bridge *BridgeSession) IncreaseNoticePeriod(period *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.IncreaseNoticePeriod(&_Bridge.TransactOpts, period)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Bridge *BridgeTransactorSession) IncreaseNoticePeriod(period *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.IncreaseNoticePeriod(&_Bridge.TransactOpts, period)
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Bridge *BridgeTransactor) NotifyResetSigners(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "notifyResetSigners")
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Bridge *BridgeSession) NotifyResetSigners() (*types.Transaction, error) {
	return _Bridge.Contract.NotifyResetSigners(&_Bridge.TransactOpts)
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Bridge *BridgeTransactorSession) NotifyResetSigners() (*types.Transaction, error) {
	return _Bridge.Contract.NotifyResetSigners(&_Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactor) Relay(opts *bind.TransactOpts, _relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "relay", _relayRequest, _sigs, _signers, _powers)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeSession) Relay(_relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Relay(&_Bridge.TransactOpts, _relayRequest, _sigs, _signers, _powers)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactorSession) Relay(_relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Relay(&_Bridge.TransactOpts, _relayRequest, _sigs, _signers, _powers)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Bridge *BridgeTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Bridge *BridgeSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveGovernor(&_Bridge.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Bridge *BridgeTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveGovernor(&_Bridge.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Bridge *BridgeTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Bridge *BridgeSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemovePauser(&_Bridge.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Bridge *BridgeTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemovePauser(&_Bridge.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Bridge *BridgeTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Bridge *BridgeSession) RenounceGovernor() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceGovernor(&_Bridge.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Bridge *BridgeTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceGovernor(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Bridge *BridgeTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Bridge *BridgeSession) RenouncePauser() (*types.Transaction, error) {
	return _Bridge.Contract.RenouncePauser(&_Bridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Bridge *BridgeTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Bridge.Contract.RenouncePauser(&_Bridge.TransactOpts)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactor) ResetSigners(opts *bind.TransactOpts, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "resetSigners", _signers, _powers)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeSession) ResetSigners(_signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ResetSigners(&_Bridge.TransactOpts, _signers, _powers)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactorSession) ResetSigners(_signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ResetSigners(&_Bridge.TransactOpts, _signers, _powers)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeTransactor) Send(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "send", _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.Send(&_Bridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeTransactorSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.Send(&_Bridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Bridge *BridgeTransactor) SendNative(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "sendNative", _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Bridge *BridgeSession) SendNative(_receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SendNative(&_Bridge.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Bridge *BridgeTransactorSession) SendNative(_receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SendNative(&_Bridge.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Bridge *BridgeTransactor) SetDelayPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setDelayPeriod", _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Bridge *BridgeSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayPeriod(&_Bridge.TransactOpts, _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Bridge *BridgeTransactorSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayPeriod(&_Bridge.TransactOpts, _period)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Bridge *BridgeTransactor) SetDelayThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setDelayThresholds", _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Bridge *BridgeSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayThresholds(&_Bridge.TransactOpts, _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Bridge *BridgeTransactorSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayThresholds(&_Bridge.TransactOpts, _tokens, _thresholds)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Bridge *BridgeTransactor) SetEpochLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setEpochLength", _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Bridge *BridgeSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetEpochLength(&_Bridge.TransactOpts, _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Bridge *BridgeTransactorSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetEpochLength(&_Bridge.TransactOpts, _length)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Bridge *BridgeTransactor) SetEpochVolumeCaps(opts *bind.TransactOpts, _tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setEpochVolumeCaps", _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Bridge *BridgeSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetEpochVolumeCaps(&_Bridge.TransactOpts, _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Bridge *BridgeTransactorSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetEpochVolumeCaps(&_Bridge.TransactOpts, _tokens, _caps)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactor) SetMaxSend(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMaxSend", _tokens, _amounts)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeSession) SetMaxSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMaxSend(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactorSession) SetMaxSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMaxSend(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactor) SetMinAdd(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMinAdd", _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeSession) SetMinAdd(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinAdd(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactorSession) SetMinAdd(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinAdd(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactor) SetMinSend(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMinSend", _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeSession) SetMinSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinSend(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactorSession) SetMinSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinSend(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Bridge *BridgeTransactor) SetMinimalMaxSlippage(opts *bind.TransactOpts, _minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMinimalMaxSlippage", _minimalMaxSlippage)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Bridge *BridgeSession) SetMinimalMaxSlippage(_minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinimalMaxSlippage(&_Bridge.TransactOpts, _minimalMaxSlippage)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Bridge *BridgeTransactorSession) SetMinimalMaxSlippage(_minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinimalMaxSlippage(&_Bridge.TransactOpts, _minimalMaxSlippage)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Bridge *BridgeTransactor) SetWrap(opts *bind.TransactOpts, _weth common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setWrap", _weth)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Bridge *BridgeSession) SetWrap(_weth common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetWrap(&_Bridge.TransactOpts, _weth)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Bridge *BridgeTransactorSession) SetWrap(_weth common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetWrap(&_Bridge.TransactOpts, _weth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Bridge *BridgeTransactor) UpdateSigners(opts *bind.TransactOpts, _triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateSigners", _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Bridge *BridgeSession) UpdateSigners(_triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateSigners(&_Bridge.TransactOpts, _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Bridge *BridgeTransactorSession) UpdateSigners(_triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateSigners(&_Bridge.TransactOpts, _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, _wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", _wdmsg, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeSession) Withdraw(_wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _wdmsg, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(_wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _wdmsg, _sigs, _signers, _powers)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeDelayPeriodUpdatedIterator is returned from FilterDelayPeriodUpdated and is used to iterate over the raw logs and unpacked data for DelayPeriodUpdated events raised by the Bridge contract.
type BridgeDelayPeriodUpdatedIterator struct {
	Event *BridgeDelayPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeDelayPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelayPeriodUpdated)
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
		it.Event = new(BridgeDelayPeriodUpdated)
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
func (it *BridgeDelayPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelayPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelayPeriodUpdated represents a DelayPeriodUpdated event raised by the Bridge contract.
type BridgeDelayPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelayPeriodUpdated is a free log retrieval operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Bridge *BridgeFilterer) FilterDelayPeriodUpdated(opts *bind.FilterOpts) (*BridgeDelayPeriodUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeDelayPeriodUpdatedIterator{contract: _Bridge.contract, event: "DelayPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayPeriodUpdated is a free log subscription operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Bridge *BridgeFilterer) WatchDelayPeriodUpdated(opts *bind.WatchOpts, sink chan<- *BridgeDelayPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelayPeriodUpdated)
				if err := _Bridge.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDelayPeriodUpdated(log types.Log) (*BridgeDelayPeriodUpdated, error) {
	event := new(BridgeDelayPeriodUpdated)
	if err := _Bridge.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDelayThresholdUpdatedIterator is returned from FilterDelayThresholdUpdated and is used to iterate over the raw logs and unpacked data for DelayThresholdUpdated events raised by the Bridge contract.
type BridgeDelayThresholdUpdatedIterator struct {
	Event *BridgeDelayThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeDelayThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelayThresholdUpdated)
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
		it.Event = new(BridgeDelayThresholdUpdated)
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
func (it *BridgeDelayThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelayThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelayThresholdUpdated represents a DelayThresholdUpdated event raised by the Bridge contract.
type BridgeDelayThresholdUpdated struct {
	Token     common.Address
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayThresholdUpdated is a free log retrieval operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Bridge *BridgeFilterer) FilterDelayThresholdUpdated(opts *bind.FilterOpts) (*BridgeDelayThresholdUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeDelayThresholdUpdatedIterator{contract: _Bridge.contract, event: "DelayThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayThresholdUpdated is a free log subscription operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Bridge *BridgeFilterer) WatchDelayThresholdUpdated(opts *bind.WatchOpts, sink chan<- *BridgeDelayThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelayThresholdUpdated)
				if err := _Bridge.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDelayThresholdUpdated(log types.Log) (*BridgeDelayThresholdUpdated, error) {
	event := new(BridgeDelayThresholdUpdated)
	if err := _Bridge.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDelayedTransferAddedIterator is returned from FilterDelayedTransferAdded and is used to iterate over the raw logs and unpacked data for DelayedTransferAdded events raised by the Bridge contract.
type BridgeDelayedTransferAddedIterator struct {
	Event *BridgeDelayedTransferAdded // Event containing the contract specifics and raw log

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
func (it *BridgeDelayedTransferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelayedTransferAdded)
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
		it.Event = new(BridgeDelayedTransferAdded)
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
func (it *BridgeDelayedTransferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelayedTransferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelayedTransferAdded represents a DelayedTransferAdded event raised by the Bridge contract.
type BridgeDelayedTransferAdded struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferAdded is a free log retrieval operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Bridge *BridgeFilterer) FilterDelayedTransferAdded(opts *bind.FilterOpts) (*BridgeDelayedTransferAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeDelayedTransferAddedIterator{contract: _Bridge.contract, event: "DelayedTransferAdded", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferAdded is a free log subscription operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Bridge *BridgeFilterer) WatchDelayedTransferAdded(opts *bind.WatchOpts, sink chan<- *BridgeDelayedTransferAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelayedTransferAdded)
				if err := _Bridge.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDelayedTransferAdded(log types.Log) (*BridgeDelayedTransferAdded, error) {
	event := new(BridgeDelayedTransferAdded)
	if err := _Bridge.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDelayedTransferExecutedIterator is returned from FilterDelayedTransferExecuted and is used to iterate over the raw logs and unpacked data for DelayedTransferExecuted events raised by the Bridge contract.
type BridgeDelayedTransferExecutedIterator struct {
	Event *BridgeDelayedTransferExecuted // Event containing the contract specifics and raw log

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
func (it *BridgeDelayedTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelayedTransferExecuted)
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
		it.Event = new(BridgeDelayedTransferExecuted)
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
func (it *BridgeDelayedTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelayedTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelayedTransferExecuted represents a DelayedTransferExecuted event raised by the Bridge contract.
type BridgeDelayedTransferExecuted struct {
	Id       [32]byte
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferExecuted is a free log retrieval operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterDelayedTransferExecuted(opts *bind.FilterOpts) (*BridgeDelayedTransferExecutedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return &BridgeDelayedTransferExecutedIterator{contract: _Bridge.contract, event: "DelayedTransferExecuted", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferExecuted is a free log subscription operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchDelayedTransferExecuted(opts *bind.WatchOpts, sink chan<- *BridgeDelayedTransferExecuted) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelayedTransferExecuted)
				if err := _Bridge.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDelayedTransferExecuted(log types.Log) (*BridgeDelayedTransferExecuted, error) {
	event := new(BridgeDelayedTransferExecuted)
	if err := _Bridge.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEpochLengthUpdatedIterator is returned from FilterEpochLengthUpdated and is used to iterate over the raw logs and unpacked data for EpochLengthUpdated events raised by the Bridge contract.
type BridgeEpochLengthUpdatedIterator struct {
	Event *BridgeEpochLengthUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeEpochLengthUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEpochLengthUpdated)
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
		it.Event = new(BridgeEpochLengthUpdated)
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
func (it *BridgeEpochLengthUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEpochLengthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEpochLengthUpdated represents a EpochLengthUpdated event raised by the Bridge contract.
type BridgeEpochLengthUpdated struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEpochLengthUpdated is a free log retrieval operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Bridge *BridgeFilterer) FilterEpochLengthUpdated(opts *bind.FilterOpts) (*BridgeEpochLengthUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeEpochLengthUpdatedIterator{contract: _Bridge.contract, event: "EpochLengthUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochLengthUpdated is a free log subscription operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Bridge *BridgeFilterer) WatchEpochLengthUpdated(opts *bind.WatchOpts, sink chan<- *BridgeEpochLengthUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEpochLengthUpdated)
				if err := _Bridge.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseEpochLengthUpdated(log types.Log) (*BridgeEpochLengthUpdated, error) {
	event := new(BridgeEpochLengthUpdated)
	if err := _Bridge.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEpochVolumeUpdatedIterator is returned from FilterEpochVolumeUpdated and is used to iterate over the raw logs and unpacked data for EpochVolumeUpdated events raised by the Bridge contract.
type BridgeEpochVolumeUpdatedIterator struct {
	Event *BridgeEpochVolumeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeEpochVolumeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEpochVolumeUpdated)
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
		it.Event = new(BridgeEpochVolumeUpdated)
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
func (it *BridgeEpochVolumeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEpochVolumeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEpochVolumeUpdated represents a EpochVolumeUpdated event raised by the Bridge contract.
type BridgeEpochVolumeUpdated struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEpochVolumeUpdated is a free log retrieval operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Bridge *BridgeFilterer) FilterEpochVolumeUpdated(opts *bind.FilterOpts) (*BridgeEpochVolumeUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeEpochVolumeUpdatedIterator{contract: _Bridge.contract, event: "EpochVolumeUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochVolumeUpdated is a free log subscription operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Bridge *BridgeFilterer) WatchEpochVolumeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeEpochVolumeUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEpochVolumeUpdated)
				if err := _Bridge.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseEpochVolumeUpdated(log types.Log) (*BridgeEpochVolumeUpdated, error) {
	event := new(BridgeEpochVolumeUpdated)
	if err := _Bridge.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the Bridge contract.
type BridgeGovernorAddedIterator struct {
	Event *BridgeGovernorAdded // Event containing the contract specifics and raw log

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
func (it *BridgeGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeGovernorAdded)
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
		it.Event = new(BridgeGovernorAdded)
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
func (it *BridgeGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeGovernorAdded represents a GovernorAdded event raised by the Bridge contract.
type BridgeGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Bridge *BridgeFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*BridgeGovernorAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeGovernorAddedIterator{contract: _Bridge.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Bridge *BridgeFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *BridgeGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeGovernorAdded)
				if err := _Bridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseGovernorAdded(log types.Log) (*BridgeGovernorAdded, error) {
	event := new(BridgeGovernorAdded)
	if err := _Bridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the Bridge contract.
type BridgeGovernorRemovedIterator struct {
	Event *BridgeGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeGovernorRemoved)
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
		it.Event = new(BridgeGovernorRemoved)
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
func (it *BridgeGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeGovernorRemoved represents a GovernorRemoved event raised by the Bridge contract.
type BridgeGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Bridge *BridgeFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*BridgeGovernorRemovedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &BridgeGovernorRemovedIterator{contract: _Bridge.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Bridge *BridgeFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *BridgeGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeGovernorRemoved)
				if err := _Bridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseGovernorRemoved(log types.Log) (*BridgeGovernorRemoved, error) {
	event := new(BridgeGovernorRemoved)
	if err := _Bridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Bridge contract.
type BridgeLiquidityAddedIterator struct {
	Event *BridgeLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *BridgeLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLiquidityAdded)
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
		it.Event = new(BridgeLiquidityAdded)
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
func (it *BridgeLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLiquidityAdded represents a LiquidityAdded event raised by the Bridge contract.
type BridgeLiquidityAdded struct {
	Seqnum   uint64
	Provider common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*BridgeLiquidityAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeLiquidityAddedIterator{contract: _Bridge.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *BridgeLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLiquidityAdded)
				if err := _Bridge.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseLiquidityAdded(log types.Log) (*BridgeLiquidityAdded, error) {
	event := new(BridgeLiquidityAdded)
	if err := _Bridge.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMaxSendUpdatedIterator is returned from FilterMaxSendUpdated and is used to iterate over the raw logs and unpacked data for MaxSendUpdated events raised by the Bridge contract.
type BridgeMaxSendUpdatedIterator struct {
	Event *BridgeMaxSendUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeMaxSendUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMaxSendUpdated)
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
		it.Event = new(BridgeMaxSendUpdated)
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
func (it *BridgeMaxSendUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMaxSendUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMaxSendUpdated represents a MaxSendUpdated event raised by the Bridge contract.
type BridgeMaxSendUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxSendUpdated is a free log retrieval operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterMaxSendUpdated(opts *bind.FilterOpts) (*BridgeMaxSendUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MaxSendUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeMaxSendUpdatedIterator{contract: _Bridge.contract, event: "MaxSendUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSendUpdated is a free log subscription operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchMaxSendUpdated(opts *bind.WatchOpts, sink chan<- *BridgeMaxSendUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MaxSendUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMaxSendUpdated)
				if err := _Bridge.contract.UnpackLog(event, "MaxSendUpdated", log); err != nil {
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

// ParseMaxSendUpdated is a log parse operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseMaxSendUpdated(log types.Log) (*BridgeMaxSendUpdated, error) {
	event := new(BridgeMaxSendUpdated)
	if err := _Bridge.contract.UnpackLog(event, "MaxSendUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMinAddUpdatedIterator is returned from FilterMinAddUpdated and is used to iterate over the raw logs and unpacked data for MinAddUpdated events raised by the Bridge contract.
type BridgeMinAddUpdatedIterator struct {
	Event *BridgeMinAddUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeMinAddUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMinAddUpdated)
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
		it.Event = new(BridgeMinAddUpdated)
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
func (it *BridgeMinAddUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMinAddUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMinAddUpdated represents a MinAddUpdated event raised by the Bridge contract.
type BridgeMinAddUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinAddUpdated is a free log retrieval operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterMinAddUpdated(opts *bind.FilterOpts) (*BridgeMinAddUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MinAddUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeMinAddUpdatedIterator{contract: _Bridge.contract, event: "MinAddUpdated", logs: logs, sub: sub}, nil
}

// WatchMinAddUpdated is a free log subscription operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchMinAddUpdated(opts *bind.WatchOpts, sink chan<- *BridgeMinAddUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MinAddUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMinAddUpdated)
				if err := _Bridge.contract.UnpackLog(event, "MinAddUpdated", log); err != nil {
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

// ParseMinAddUpdated is a log parse operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseMinAddUpdated(log types.Log) (*BridgeMinAddUpdated, error) {
	event := new(BridgeMinAddUpdated)
	if err := _Bridge.contract.UnpackLog(event, "MinAddUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMinSendUpdatedIterator is returned from FilterMinSendUpdated and is used to iterate over the raw logs and unpacked data for MinSendUpdated events raised by the Bridge contract.
type BridgeMinSendUpdatedIterator struct {
	Event *BridgeMinSendUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeMinSendUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMinSendUpdated)
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
		it.Event = new(BridgeMinSendUpdated)
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
func (it *BridgeMinSendUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMinSendUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMinSendUpdated represents a MinSendUpdated event raised by the Bridge contract.
type BridgeMinSendUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinSendUpdated is a free log retrieval operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterMinSendUpdated(opts *bind.FilterOpts) (*BridgeMinSendUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MinSendUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeMinSendUpdatedIterator{contract: _Bridge.contract, event: "MinSendUpdated", logs: logs, sub: sub}, nil
}

// WatchMinSendUpdated is a free log subscription operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchMinSendUpdated(opts *bind.WatchOpts, sink chan<- *BridgeMinSendUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MinSendUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMinSendUpdated)
				if err := _Bridge.contract.UnpackLog(event, "MinSendUpdated", log); err != nil {
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

// ParseMinSendUpdated is a log parse operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseMinSendUpdated(log types.Log) (*BridgeMinSendUpdated, error) {
	event := new(BridgeMinSendUpdated)
	if err := _Bridge.contract.UnpackLog(event, "MinSendUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Bridge contract.
type BridgePauserAddedIterator struct {
	Event *BridgePauserAdded // Event containing the contract specifics and raw log

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
func (it *BridgePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePauserAdded)
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
		it.Event = new(BridgePauserAdded)
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
func (it *BridgePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePauserAdded represents a PauserAdded event raised by the Bridge contract.
type BridgePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Bridge *BridgeFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*BridgePauserAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &BridgePauserAddedIterator{contract: _Bridge.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Bridge *BridgeFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *BridgePauserAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePauserAdded)
				if err := _Bridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParsePauserAdded(log types.Log) (*BridgePauserAdded, error) {
	event := new(BridgePauserAdded)
	if err := _Bridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Bridge contract.
type BridgePauserRemovedIterator struct {
	Event *BridgePauserRemoved // Event containing the contract specifics and raw log

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
func (it *BridgePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePauserRemoved)
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
		it.Event = new(BridgePauserRemoved)
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
func (it *BridgePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePauserRemoved represents a PauserRemoved event raised by the Bridge contract.
type BridgePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Bridge *BridgeFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*BridgePauserRemovedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &BridgePauserRemovedIterator{contract: _Bridge.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Bridge *BridgeFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *BridgePauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePauserRemoved)
				if err := _Bridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParsePauserRemoved(log types.Log) (*BridgePauserRemoved, error) {
	event := new(BridgePauserRemoved)
	if err := _Bridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayIterator is returned from FilterRelay and is used to iterate over the raw logs and unpacked data for Relay events raised by the Bridge contract.
type BridgeRelayIterator struct {
	Event *BridgeRelay // Event containing the contract specifics and raw log

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
func (it *BridgeRelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelay)
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
		it.Event = new(BridgeRelay)
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
func (it *BridgeRelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelay represents a Relay event raised by the Bridge contract.
type BridgeRelay struct {
	TransferId    [32]byte
	Sender        common.Address
	Receiver      common.Address
	Token         common.Address
	Amount        *big.Int
	SrcChainId    uint64
	SrcTransferId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRelay is a free log retrieval operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) FilterRelay(opts *bind.FilterOpts) (*BridgeRelayIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Relay")
	if err != nil {
		return nil, err
	}
	return &BridgeRelayIterator{contract: _Bridge.contract, event: "Relay", logs: logs, sub: sub}, nil
}

// WatchRelay is a free log subscription operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) WatchRelay(opts *bind.WatchOpts, sink chan<- *BridgeRelay) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Relay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelay)
				if err := _Bridge.contract.UnpackLog(event, "Relay", log); err != nil {
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

// ParseRelay is a log parse operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) ParseRelay(log types.Log) (*BridgeRelay, error) {
	event := new(BridgeRelay)
	if err := _Bridge.contract.UnpackLog(event, "Relay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeResetNotificationIterator is returned from FilterResetNotification and is used to iterate over the raw logs and unpacked data for ResetNotification events raised by the Bridge contract.
type BridgeResetNotificationIterator struct {
	Event *BridgeResetNotification // Event containing the contract specifics and raw log

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
func (it *BridgeResetNotificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeResetNotification)
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
		it.Event = new(BridgeResetNotification)
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
func (it *BridgeResetNotificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeResetNotificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeResetNotification represents a ResetNotification event raised by the Bridge contract.
type BridgeResetNotification struct {
	ResetTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResetNotification is a free log retrieval operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Bridge *BridgeFilterer) FilterResetNotification(opts *bind.FilterOpts) (*BridgeResetNotificationIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ResetNotification")
	if err != nil {
		return nil, err
	}
	return &BridgeResetNotificationIterator{contract: _Bridge.contract, event: "ResetNotification", logs: logs, sub: sub}, nil
}

// WatchResetNotification is a free log subscription operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Bridge *BridgeFilterer) WatchResetNotification(opts *bind.WatchOpts, sink chan<- *BridgeResetNotification) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ResetNotification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeResetNotification)
				if err := _Bridge.contract.UnpackLog(event, "ResetNotification", log); err != nil {
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

// ParseResetNotification is a log parse operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Bridge *BridgeFilterer) ParseResetNotification(log types.Log) (*BridgeResetNotification, error) {
	event := new(BridgeResetNotification)
	if err := _Bridge.contract.UnpackLog(event, "ResetNotification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the Bridge contract.
type BridgeSendIterator struct {
	Event *BridgeSend // Event containing the contract specifics and raw log

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
func (it *BridgeSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSend)
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
		it.Event = new(BridgeSend)
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
func (it *BridgeSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSend represents a Send event raised by the Bridge contract.
type BridgeSend struct {
	TransferId  [32]byte
	Sender      common.Address
	Receiver    common.Address
	Token       common.Address
	Amount      *big.Int
	DstChainId  uint64
	Nonce       uint64
	MaxSlippage uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) FilterSend(opts *bind.FilterOpts) (*BridgeSendIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return &BridgeSendIterator{contract: _Bridge.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *BridgeSend) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSend)
				if err := _Bridge.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) ParseSend(log types.Log) (*BridgeSend, error) {
	event := new(BridgeSend)
	if err := _Bridge.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSignersUpdatedIterator is returned from FilterSignersUpdated and is used to iterate over the raw logs and unpacked data for SignersUpdated events raised by the Bridge contract.
type BridgeSignersUpdatedIterator struct {
	Event *BridgeSignersUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeSignersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSignersUpdated)
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
		it.Event = new(BridgeSignersUpdated)
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
func (it *BridgeSignersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSignersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSignersUpdated represents a SignersUpdated event raised by the Bridge contract.
type BridgeSignersUpdated struct {
	Signers []common.Address
	Powers  []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignersUpdated is a free log retrieval operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Bridge *BridgeFilterer) FilterSignersUpdated(opts *bind.FilterOpts) (*BridgeSignersUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SignersUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeSignersUpdatedIterator{contract: _Bridge.contract, event: "SignersUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersUpdated is a free log subscription operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Bridge *BridgeFilterer) WatchSignersUpdated(opts *bind.WatchOpts, sink chan<- *BridgeSignersUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SignersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSignersUpdated)
				if err := _Bridge.contract.UnpackLog(event, "SignersUpdated", log); err != nil {
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

// ParseSignersUpdated is a log parse operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Bridge *BridgeFilterer) ParseSignersUpdated(log types.Log) (*BridgeSignersUpdated, error) {
	event := new(BridgeSignersUpdated)
	if err := _Bridge.contract.UnpackLog(event, "SignersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawDoneIterator is returned from FilterWithdrawDone and is used to iterate over the raw logs and unpacked data for WithdrawDone events raised by the Bridge contract.
type BridgeWithdrawDoneIterator struct {
	Event *BridgeWithdrawDone // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawDone)
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
		it.Event = new(BridgeWithdrawDone)
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
func (it *BridgeWithdrawDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawDone represents a WithdrawDone event raised by the Bridge contract.
type BridgeWithdrawDone struct {
	WithdrawId [32]byte
	Seqnum     uint64
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Refid      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawDone is a free log retrieval operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Bridge *BridgeFilterer) FilterWithdrawDone(opts *bind.FilterOpts) (*BridgeWithdrawDoneIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawDoneIterator{contract: _Bridge.contract, event: "WithdrawDone", logs: logs, sub: sub}, nil
}

// WatchWithdrawDone is a free log subscription operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Bridge *BridgeFilterer) WatchWithdrawDone(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawDone) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawDone)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
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

// ParseWithdrawDone is a log parse operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Bridge *BridgeFilterer) ParseWithdrawDone(log types.Log) (*BridgeWithdrawDone, error) {
	event := new(BridgeWithdrawDone)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
