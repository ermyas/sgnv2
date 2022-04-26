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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DelayThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedTransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelayedTransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EpochLengthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EpochVolumeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MaxSendUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinAddUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinSendUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"}],\"name\":\"Relay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resetTime\",\"type\":\"uint256\"}],\"name\":\"ResetNotification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxSlippage\",\"type\":\"uint32\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"SignersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"refid\",\"type\":\"bytes32\"}],\"name\":\"WithdrawDone\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addNativeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addseq\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delayThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumeCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"executeDelayedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"increaseNoticePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastOpTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalMaxSlippage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyResetSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_relayRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"resetSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"sendNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setDelayThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"name\":\"setEpochVolumeCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMaxSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minimalMaxSlippage\",\"type\":\"uint32\"}],\"name\":\"setMinimalMaxSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"setWrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ssHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_triggerTime\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_newSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_curSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_curPowers\",\"type\":\"uint256[]\"}],\"name\":\"updateSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_wdmsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdraws\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001d3362000048565b60016005556006805460ff19169055620000373362000098565b620000423362000162565b62000222565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526007602052604090205460ff1615620001075760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064015b60405180910390fd5b6001600160a01b038116600081815260076020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f891015b60405180910390a150565b6001600160a01b03811660009081526008602052604090205460ff1615620001cd5760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f7200000000006044820152606401620000fe565b6001600160a01b038116600081815260086020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5910162000157565b61517d80620002326000396000f3fe6080604052600436106103645760003560e01c806382dc1ec4116101c6578063ba2cb25c116100f7578063e43581b811610095578063f20c922a1161006f578063f20c922a14610aaf578063f2fde38b14610acf578063f832138314610aef578063f8b30d7d14610b1c57600080fd5b8063e43581b814610a36578063e999e5f414610a6f578063eecdac8814610a8f57600080fd5b8063d0790da9116100d1578063d0790da9146109ab578063e026049c146109c1578063e09ab428146109d6578063e3eece2614610a0657600080fd5b8063ba2cb25c1461093e578063ccde517a1461095e578063cdd1b25d1461098b57600080fd5b80639ff9001a11610164578063a7bdf45a1161013e578063a7bdf45a14610861578063adc0d57f14610881578063b1c94d94146108fb578063b5f2bc471461091157600080fd5b80639ff9001a14610801578063a21a928014610821578063a5977fbb1461084157600080fd5b806389e39127116101a057806389e39127146107735780638da5cb5b146107ad5780639b14d4c6146107cb5780639e25fc5c146107e157600080fd5b806382dc1ec41461071e5780638456cb591461073e578063878fe1ce1461075357600080fd5b806348234126116102a0578063618ee0551161023e5780636b2c0f55116102185780636b2c0f55146106a65780636ef8d66d146106c65780637044c89e146106db57806380f51c12146106ee57600080fd5b8063618ee0551461064357806365a114f114610670578063682dbc221461068657600080fd5b8063566887001161027a57806356688700146105c857806357d775f8146105e85780635c975abb146105fe57806360216b001461061657600080fd5b8063482341261461055b57806352532faa1461057b57806354eea796146105a857600080fd5b80633c64f04b1161030d5780633f4ba83a116102e75780633f4ba83a146104b5578063457bfa2f146104ca57806346fbf68e1461050257806347b16c6c1461053b57600080fd5b80633c64f04b146104425780633d572107146104825780633f2e5fc3146104a257600080fd5b80632fd1b0a41161033e5780632fd1b0a4146103c7578063370fb47b146103fe5780633c4a25d01461042257600080fd5b8063089927411461037057806317bdbae51461039257806325c38b9f146103b257600080fd5b3661036b57005b600080fd5b34801561037c57600080fd5b5061039061038b3660046148ce565b610b49565b005b34801561039e57600080fd5b506103906103ad3660046148ce565b610cec565b3480156103be57600080fd5b50610390610e83565b3480156103d357600080fd5b506017546103e49063ffffffff1681565b60405163ffffffff90911681526020015b60405180910390f35b34801561040a57600080fd5b5061041460025481565b6040519081526020016103f5565b34801561042e57600080fd5b5061039061043d366004614956565b610f22565b34801561044e57600080fd5b5061047261045d366004614971565b60146020526000908152604090205460ff1681565b60405190151581526020016103f5565b34801561048e57600080fd5b5061039061049d366004614971565b610f85565b6103906104b03660046149b6565b611019565b3480156104c157600080fd5b5061039061126f565b3480156104d657600080fd5b506013546104ea906001600160a01b031681565b6040516001600160a01b0390911681526020016103f5565b34801561050e57600080fd5b5061047261051d366004614956565b6001600160a01b031660009081526007602052604090205460ff1690565b34801561054757600080fd5b506103906105563660046148ce565b6112d8565b34801561056757600080fd5b50610390610576366004614a14565b61146f565b34801561058757600080fd5b50610414610596366004614956565b600e6020526000908152604090205481565b3480156105b457600080fd5b506103906105c3366004614971565b6114e3565b3480156105d457600080fd5b506103906105e3366004614a2f565b611570565b3480156105f457600080fd5b5061041460095481565b34801561060a57600080fd5b5060065460ff16610472565b34801561062257600080fd5b50610414610631366004614956565b600a6020526000908152604090205481565b34801561064f57600080fd5b5061041461065e366004614956565b60166020526000908152604090205481565b34801561067c57600080fd5b5061041460035481565b34801561069257600080fd5b506103906106a1366004614a6f565b611732565b3480156106b257600080fd5b506103906106c1366004614956565b61181e565b3480156106d257600080fd5b5061039061187e565b6103906106e9366004614971565b611887565b3480156106fa57600080fd5b50610472610709366004614956565b60076020526000908152604090205460ff1681565b34801561072a57600080fd5b50610390610739366004614956565b611b39565b34801561074a57600080fd5b50610390611b99565b34801561075f57600080fd5b5061039061076e3660046148ce565b611c00565b34801561077f57600080fd5b506010546107949067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016103f5565b3480156107b957600080fd5b506000546001600160a01b03166104ea565b3480156107d757600080fd5b5061041460045481565b3480156107ed57600080fd5b506103906107fc366004614971565b611d97565b34801561080d57600080fd5b5061039061081c366004614956565b611e05565b34801561082d57600080fd5b5061039061083c366004614b9d565b611e7e565b34801561084d57600080fd5b5061039061085c366004614c8c565b6121c5565b34801561086d57600080fd5b5061039061087c3660046148ce565b612313565b34801561088d57600080fd5b506108d061089c366004614971565b600d6020526000908152604090208054600182015460028301546003909301546001600160a01b0392831693919092169184565b604080516001600160a01b0395861681529490931660208501529183015260608201526080016103f5565b34801561090757600080fd5b50610414600f5481565b34801561091d57600080fd5b5061041461092c366004614956565b600b6020526000908152604090205481565b34801561094a57600080fd5b50610390610959366004614cf9565b6123d3565b34801561096a57600080fd5b50610414610979366004614956565b60116020526000908152604090205481565b34801561099757600080fd5b506103906109a6366004614b9d565b612529565b3480156109b757600080fd5b5061041460015481565b3480156109cd57600080fd5b5061039061284e565b3480156109e257600080fd5b506104726109f1366004614971565b60126020526000908152604090205460ff1681565b348015610a1257600080fd5b50610472610a21366004614956565b60086020526000908152604090205460ff1681565b348015610a4257600080fd5b50610472610a51366004614956565b6001600160a01b031660009081526008602052604090205460ff1690565b348015610a7b57600080fd5b50610390610a8a3660046148ce565b612857565b348015610a9b57600080fd5b50610390610aaa366004614956565b6129ee565b348015610abb57600080fd5b50610390610aca366004614971565b612a4e565b348015610adb57600080fd5b50610390610aea366004614956565b612b07565b348015610afb57600080fd5b50610414610b0a366004614956565b600c6020526000908152604090205481565b348015610b2857600080fd5b50610414610b37366004614956565b60156020526000908152604090205481565b3360009081526008602052604090205460ff16610ba65760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b60448201526064015b60405180910390fd5b828114610be75760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610b9d565b60005b83811015610ce557828282818110610c0457610c04614df9565b9050602002013560156000878785818110610c2157610c21614df9565b9050602002016020810190610c369190614956565b6001600160a01b031681526020810191909152604001600020557f8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9858583818110610c8357610c83614df9565b9050602002016020810190610c989190614956565b848484818110610caa57610caa614df9565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180610cdd81614e25565b915050610bea565b5050505050565b3360009081526008602052604090205460ff16610d445760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610b9d565b828114610d855760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610b9d565b60005b83811015610ce557828282818110610da257610da2614df9565b90506020020135600e6000878785818110610dbf57610dbf614df9565b9050602002016020810190610dd49190614956565b6001600160a01b031681526020810191909152604001600020557fceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce858583818110610e2157610e21614df9565b9050602002016020810190610e369190614956565b848484818110610e4857610e48614df9565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180610e7b81614e25565b915050610d88565b33610e966000546001600160a01b031690565b6001600160a01b031614610eda5760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b600454610ee79042614e40565b60038190556040519081527f68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1906020015b60405180910390a1565b33610f356000546001600160a01b031690565b6001600160a01b031614610f795760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b610f8281612be3565b50565b3360009081526008602052604090205460ff16610fdd5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610b9d565b600f8190556040518181527fc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6906020015b60405180910390a150565b6002600554141561106c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610b9d565b600260055560065460ff16156110b75760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b9d565b8334146110f85760405162461bcd60e51b815260206004820152600f60248201526e082dadeeadce840dad2e6dac2e8c6d608b1b6044820152606401610b9d565b6013546001600160a01b03166111505760405162461bcd60e51b815260206004820152601360248201527f4e61746976652077726170206e6f7420736574000000000000000000000000006044820152606401610b9d565b60135460009061116e9087906001600160a01b031687878787612ca0565b9050601360009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0866040518263ffffffff1660e01b81526004016000604051808303818588803b1580156111c057600080fd5b505af11580156111d4573d6000803e3d6000fd5b5050601354604080518681523360208201526001600160a01b03808d1692820192909252911660608201526080810189905267ffffffffffffffff80891660a0830152871660c082015263ffffffff861660e08201527f89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01935061010001915061125a9050565b60405180910390a15050600160055550505050565b3360009081526007602052604090205460ff166112ce5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f74207061757365720000000000000000000000006044820152606401610b9d565b6112d6612ed3565b565b3360009081526008602052604090205460ff166113305760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610b9d565b8281146113715760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610b9d565b60005b83811015610ce55782828281811061138e5761138e614df9565b90506020020135600b60008787858181106113ab576113ab614df9565b90506020020160208101906113c09190614956565b6001600160a01b031681526020810191909152604001600020557f608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e8985858381811061140d5761140d614df9565b90506020020160208101906114229190614956565b84848481811061143457611434614df9565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a18061146781614e25565b915050611374565b3360009081526008602052604090205460ff166114c75760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610b9d565b6017805463ffffffff191663ffffffff92909216919091179055565b3360009081526008602052604090205460ff1661153b5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610b9d565b60098190556040518181527f2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b39060200161100e565b600260055414156115c35760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610b9d565b600260055560065460ff161561160e5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b9d565b6001600160a01b03821660009081526011602052604090205481116116685760405162461bcd60e51b815260206004820152601060248201526f185b5bdd5b9d081d1bdbc81cdb585b1b60821b6044820152606401610b9d565b601080546001919060009061168890849067ffffffffffffffff16614e58565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506116ce333083856001600160a01b0316612f6a909392919063ffffffff16565b6010546040805167ffffffffffffffff90921682523360208301526001600160a01b0384168282015260608201839052517fd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce7649181900360800190a150506001600555565b60008484848460405160200161174b9493929190614ef0565b60405160208183030381529060405280519060200120905080600154146117b45760405162461bcd60e51b815260206004820152601860248201527f4d69736d617463682063757272656e74207369676e65727300000000000000006044820152606401610b9d565b87516020808a0191909120604080517f19457468657265756d205369676e6564204d6573736167653a0a33320000000081850152603c8082019390935281518082039093018352605c019052805191012061181490888888888888613002565b5050505050505050565b336118316000546001600160a01b031690565b6001600160a01b0316146118755760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b610f8281613338565b6112d633613338565b600260055414156118da5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610b9d565b600260055560065460ff16156119255760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b9d565b8034146119665760405162461bcd60e51b815260206004820152600f60248201526e082dadeeadce840dad2e6dac2e8c6d608b1b6044820152606401610b9d565b6013546001600160a01b03166119be5760405162461bcd60e51b815260206004820152601360248201527f4e61746976652077726170206e6f7420736574000000000000000000000000006044820152606401610b9d565b6013546001600160a01b03166000908152601160205260409020548111611a1a5760405162461bcd60e51b815260206004820152601060248201526f185b5bdd5b9d081d1bdbc81cdb585b1b60821b6044820152606401610b9d565b6010805460019190600090611a3a90849067ffffffffffffffff16614e58565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550601360009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b158015611ab057600080fd5b505af1158015611ac4573d6000803e3d6000fd5b50506010546013546040805167ffffffffffffffff90931683523360208401526001600160a01b0390911690820152606081018590527fd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce76493506080019150611b299050565b60405180910390a1506001600555565b33611b4c6000546001600160a01b031690565b6001600160a01b031614611b905760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b610f82816133f1565b3360009081526007602052604090205460ff16611bf85760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f74207061757365720000000000000000000000006044820152606401610b9d565b6112d66134ae565b3360009081526008602052604090205460ff16611c585760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610b9d565b828114611c995760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610b9d565b60005b83811015610ce557828282818110611cb657611cb6614df9565b9050602002013560166000878785818110611cd357611cd3614df9565b9050602002016020810190611ce89190614956565b6001600160a01b031681526020810191909152604001600020557f4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c858583818110611d3557611d35614df9565b9050602002016020810190611d4a9190614956565b848484818110611d5c57611d5c614df9565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a180611d8f81614e25565b915050611c9c565b60065460ff1615611ddd5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b9d565b6000611de882613529565b9050611e018160000151826020015183604001516136ee565b5050565b33611e186000546001600160a01b031690565b6001600160a01b031614611e5c5760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b601380546001600160a01b0319166001600160a01b0392909216919091179055565b60065460ff1615611ec45760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b9d565b60004630604051602001611f1a92919091825260601b6bffffffffffffffffffffffff191660208201527f57697468647261774d73670000000000000000000000000000000000000000006034820152603f0190565b604051602081830303815290604052805190602001209050611f64818a8a604051602001611f4a93929190614f11565b604051602081830303815290604052888888888888611732565b6000611fa58a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061382392505050565b905060008160000151826020015183604001518460600151856080015160405160200161201e95949392919060c095861b6001600160c01b031990811682529490951b9093166008850152606091821b6bffffffffffffffffffffffff199081166010860152911b166024830152603882015260580190565b60408051601f1981840301815291815281516020928301206000818152601290935291205490915060ff16156120965760405162461bcd60e51b815260206004820152601a60248201527f776974686472617720616c7265616479207375636365656465640000000000006044820152606401610b9d565b6000818152601260205260409020805460ff19166001179055606082015160808301516120c39190613983565b60608201516001600160a01b03166000908152600e602052604090205480158015906120f25750808360800151115b156121145761210f82846040015185606001518660800151613a9b565b61212b565b61212b8360400151846060015185608001516136ee565b7f48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be8284602001518560400151866060015187608001518860a001516040516121af9695949392919095865267ffffffffffffffff9490941660208601526001600160a01b03928316604086015291166060840152608083015260a082015260c00190565b60405180910390a1505050505050505050505050565b600260055414156122185760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610b9d565b600260055560065460ff16156122635760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b9d565b6000612273878787878787612ca0565b905061228a6001600160a01b038716333088612f6a565b604080518281523360208201526001600160a01b0389811682840152881660608201526080810187905267ffffffffffffffff86811660a0830152851660c082015263ffffffff841660e082015290517f89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01918190036101000190a1505060016005555050505050565b336123266000546001600160a01b031690565b6001600160a01b03161461236a5760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b60035442116123bb5760405162461bcd60e51b815260206004820152601460248201527f6e6f742072656163682072657365742074696d650000000000000000000000006044820152606401610b9d565b6000196003556123cd84848484613bae565b50505050565b6002548b116124245760405162461bcd60e51b815260206004820152601e60248201527f547269676765722074696d65206973206e6f7420696e6372656173696e6700006044820152606401610b9d565b61243042610e10614e40565b8b1061247e5760405162461bcd60e51b815260206004820152601960248201527f547269676765722074696d6520697320746f6f206c61726765000000000000006044820152606401610b9d565b600046306040516020016124d492919091825260601b6bffffffffffffffffffffffff191660208201527f5570646174655369676e65727300000000000000000000000000000000000000603482015260410190565b60405160208183030381529060405280519060200120905061250a818d8d8d8d8d604051602001611f4a96959493929190614f2b565b6125168b8b8b8b613bae565b5050506002989098555050505050505050565b60065460ff161561256f5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b9d565b600046306040516020016125c592919091825260601b6bffffffffffffffffffffffff191660208201527f52656c6179000000000000000000000000000000000000000000000000000000603482015260390190565b6040516020818303038152906040528051906020012090506125f5818a8a604051602001611f4a93929190614f11565b60006126368a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613d5892505050565b8051602080830151604080850151606080870151608088015160a089015160c0808b015187519a861b6bffffffffffffffffffffffff199081168c8c015298861b891660348c01529590941b9096166048890152605c880191909152811b6001600160c01b0319908116607c88015293901b9092166084850152608c808501929092528051808503909201825260ac909301835280519082012060008181526014909252919020549192509060ff16156127245760405162461bcd60e51b815260206004820152600f60248201526e7472616e736665722065786973747360881b6044820152606401610b9d565b60008181526014602052604090819020805460ff1916600117905582015160608301516127519190613983565b6040808301516001600160a01b03166000908152600e6020522054801580159061277e5750808360600151115b156127a05761279b82846020015185604001518660600151613a9b565b6127b7565b6127b78360200151846040015185606001516136ee565b7f79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c82846000015185602001518660400151876060015188608001518960c001516040516121af97969594939291909687526001600160a01b0395861660208801529385166040870152919093166060850152608084019290925267ffffffffffffffff9190911660a083015260c082015260e00190565b6112d633613ed0565b3360009081526008602052604090205460ff166128af5760405162461bcd60e51b815260206004820152601660248201527521b0b63632b91034b9903737ba1033b7bb32b93737b960511b6044820152606401610b9d565b8281146128f05760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610b9d565b60005b83811015610ce55782828281811061290d5761290d614df9565b905060200201356011600087878581811061292a5761292a614df9565b905060200201602081019061293f9190614956565b6001600160a01b031681526020810191909152604001600020557fc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f85858381811061298c5761298c614df9565b90506020020160208101906129a19190614956565b8484848181106129b3576129b3614df9565b604080516001600160a01b0390951685526020918202939093013590840152500160405180910390a1806129e681614e25565b9150506128f3565b33612a016000546001600160a01b031690565b6001600160a01b031614612a455760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b610f8281613ed0565b33612a616000546001600160a01b031690565b6001600160a01b031614612aa55760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b6004548111612b025760405162461bcd60e51b815260206004820152602360248201527f6e6f7469636520706572696f642063616e206f6e6c7920626520696e637265616044820152621cd95960ea1b6064820152608401610b9d565b600455565b33612b1a6000546001600160a01b031690565b6001600160a01b031614612b5e5760405162461bcd60e51b815260206004820181905260248201526000805160206151288339815191526044820152606401610b9d565b6001600160a01b038116612bda5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610b9d565b610f8281613f89565b6001600160a01b03811660009081526008602052604090205460ff1615612c4c5760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f7200000000006044820152606401610b9d565b6001600160a01b038116600081815260086020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5910161100e565b6001600160a01b0385166000908152601560205260408120548511612cfa5760405162461bcd60e51b815260206004820152601060248201526f185b5bdd5b9d081d1bdbc81cdb585b1b60821b6044820152606401610b9d565b6001600160a01b0386166000908152601660205260409020541580612d3757506001600160a01b0386166000908152601660205260409020548511155b612d835760405162461bcd60e51b815260206004820152601060248201527f616d6f756e7420746f6f206c61726765000000000000000000000000000000006044820152606401610b9d565b60175463ffffffff90811690831611612dde5760405162461bcd60e51b815260206004820152601660248201527f6d617820736c69707061676520746f6f20736d616c6c000000000000000000006044820152606401610b9d565b6040516bffffffffffffffffffffffff1933606090811b8216602084015289811b8216603484015288901b166048820152605c81018690526001600160c01b031960c086811b8216607c84015285811b8216608484015246901b16608c82015260009060940160408051601f1981840301815291815281516020928301206000818152601490935291205490915060ff1615612eae5760405162461bcd60e51b815260206004820152600f60248201526e7472616e736665722065786973747360881b6044820152606401610b9d565b6000818152601460205260409020805460ff1916600117905590509695505050505050565b60065460ff16612f255760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610b9d565b6006805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001610f18565b6040516001600160a01b03808516602483015283166044820152606481018290526123cd9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613fd9565b82811461305d5760405162461bcd60e51b815260206004820152602360248201527f7369676e65727320616e6420706f77657273206c656e677468206e6f74206d616044820152620e8c6d60eb1b6064820152608401610b9d565b6000805b848110156130a15783838281811061307b5761307b614df9565b905060200201358261308d9190614e40565b91508061309981614e25565b915050613061565b50600060036130b1836002614f53565b6130bb9190614f72565b6130c6906001614e40565b905060008080805b8a8110156132e65760006131518d8d848181106130ed576130ed614df9565b90506020028101906130ff9190614f94565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508f6140be90919063ffffffff16565b9050836001600160a01b0316816001600160a01b0316116131b45760405162461bcd60e51b815260206004820152601e60248201527f7369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610b9d565b8093505b8a8a848181106131ca576131ca614df9565b90506020020160208101906131df9190614956565b6001600160a01b0316816001600160a01b0316111561325957613203600184614e40565b92508983106132545760405162461bcd60e51b815260206004820152601060248201527f7369676e6572206e6f7420666f756e64000000000000000000000000000000006044820152606401610b9d565b6131b8565b8a8a8481811061326b5761326b614df9565b90506020020160208101906132809190614956565b6001600160a01b0316816001600160a01b031614156132c0578888848181106132ab576132ab614df9565b90506020020135856132bd9190614e40565b94505b8585106132d3575050505050505061332f565b50806132de81614e25565b9150506130ce565b5060405162461bcd60e51b815260206004820152601260248201527f71756f72756d206e6f74207265616368656400000000000000000000000000006044820152606401610b9d565b50505050505050565b6001600160a01b03811660009081526007602052604090205460ff166133a05760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f742070617573657200000000000000000000006044820152606401610b9d565b6001600160a01b038116600081815260076020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e910161100e565b6001600160a01b03811660009081526007602052604090205460ff161561345a5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c726561647920706175736572000000000000006044820152606401610b9d565b6001600160a01b038116600081815260076020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8910161100e565b60065460ff16156134f45760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b9d565b6006805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612f523390565b6040805160808101825260008082526020820181905291810182905260608101919091526000828152600d6020908152604091829020825160808101845281546001600160a01b03908116825260018301541692810192909252600281015492820192909252600390910154606082018190526135e85760405162461bcd60e51b815260206004820152601a60248201527f64656c61796564207472616e73666572206e6f742065786973740000000000006044820152606401610b9d565b600f5481606001516135fa9190614e40565b42116136485760405162461bcd60e51b815260206004820152601d60248201527f64656c61796564207472616e73666572207374696c6c206c6f636b65640000006044820152606401610b9d565b6000838152600d6020908152604080832080546001600160a01b03199081168255600182018054909116905560028101849055600301929092558251908301518383015192517f3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426936136e093889390929091909384526001600160a01b03928316602085015291166040830152606082015260800190565b60405180910390a192915050565b6013546001600160a01b038381169116141561380a57601354604051632e1a7d4d60e01b8152600481018390526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b15801561374a57600080fd5b505af115801561375e573d6000803e3d6000fd5b505050506000836001600160a01b03168261c35090604051600060405180830381858888f193505050503d80600081146137b4576040519150601f19603f3d011682016040523d82523d6000602084013e6137b9565b606091505b50509050806123cd5760405162461bcd60e51b815260206004820152601b60248201527f6661696c656420746f2073656e64206e617469766520746f6b656e00000000006044820152606401610b9d565b61381e6001600160a01b03831684836140e2565b505050565b6040805160c08101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905283518085019094528184528301849052909190805b6020830151518351101561397b5761388183614112565b909250905081600114156138a9576138988361414c565b67ffffffffffffffff16845261386a565b81600214156138cf576138bb8361414c565b67ffffffffffffffff16602085015261386a565b81600314156138fc576138e96138e4846141ce565b61428b565b6001600160a01b0316604085015261386a565b8160041415613924576139116138e4846141ce565b6001600160a01b0316606085015261386a565b81600514156139485761393e613939846141ce565b61429c565b608085015261386a565b816006141561396c5761396261395d846141ce565b6142d3565b60a085015261386a565b61397683826142eb565b61386a565b505050919050565b60095461398e575050565b6001600160a01b0382166000908152600b6020526040902054806139b157505050565b6001600160a01b0383166000908152600a6020526040812054600954909142916139db8184614f72565b6139e59190614f53565b6001600160a01b0387166000908152600c6020526040902054909150811115613a1057849250613a1d565b613a1a8584614e40565b92505b83831115613a6d5760405162461bcd60e51b815260206004820152601260248201527f766f6c756d6520657863656564732063617000000000000000000000000000006044820152606401610b9d565b506001600160a01b039094166000908152600a6020908152604080832093909355600c905220929092555050565b6000848152600d602052604090206003015415613afa5760405162461bcd60e51b815260206004820152601f60248201527f64656c61796564207472616e7366657220616c726561647920657869737473006044820152606401610b9d565b604080516080810182526001600160a01b0380861682528481166020808401918252838501868152426060860190815260008b8152600d90935291869020945185549085166001600160a01b031991821617865592516001860180549190951693169290921790925551600283015551600390910155517fcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce690613ba09086815260200190565b60405180910390a150505050565b828114613c095760405162461bcd60e51b815260206004820152602360248201527f7369676e65727320616e6420706f77657273206c656e677468206e6f74206d616044820152620e8c6d60eb1b6064820152608401610b9d565b6000805b84811015613ce257816001600160a01b0316868683818110613c3157613c31614df9565b9050602002016020810190613c469190614956565b6001600160a01b031611613ca75760405162461bcd60e51b815260206004820152602260248201527f4e6577207369676e657273206e6f7420696e20617363656e64696e67206f726460448201526132b960f11b6064820152608401610b9d565b858582818110613cb957613cb9614df9565b9050602002016020810190613cce9190614956565b915080613cda81614e25565b915050613c0d565b5084848484604051602001613cfa9493929190614ef0565b60408051601f198184030181529082905280516020909101206001557ff126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f890613d49908790879087908790614fdb565b60405180910390a15050505050565b6040805160e08101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905260c0830182905283518085019094528184528301849052909190805b6020830151518351101561397b57613dbd83614112565b90925090508160011415613de757613dd76138e4846141ce565b6001600160a01b03168452613da6565b8160021415613e0f57613dfc6138e4846141ce565b6001600160a01b03166020850152613da6565b8160031415613e3757613e246138e4846141ce565b6001600160a01b03166040850152613da6565b8160041415613e5657613e4c613939846141ce565b6060850152613da6565b8160051415613e7c57613e688361414c565b67ffffffffffffffff166080850152613da6565b8160061415613ea257613e8e8361414c565b67ffffffffffffffff1660a0850152613da6565b8160071415613ec157613eb761395d846141ce565b60c0850152613da6565b613ecb83826142eb565b613da6565b6001600160a01b03811660009081526008602052604090205460ff16613f385760405162461bcd60e51b815260206004820152601760248201527f4163636f756e74206973206e6f7420676f7665726e6f720000000000000000006044820152606401610b9d565b6001600160a01b038116600081815260086020908152604091829020805460ff1916905590519182527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b910161100e565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600061402e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661435d9092919063ffffffff16565b80519091501561381e578080602001905181019061404c919061505d565b61381e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610b9d565b60008060006140cd8585614376565b915091506140da816143e6565b509392505050565b6040516001600160a01b03831660248201526044810182905261381e90849063a9059cbb60e01b90606401612f9e565b60008060006141208461414c565b905061412d600882614f72565b92508060071660058111156141445761414461507f565b915050915091565b602080820151825181019091015160009182805b600a8110156141c85783811a9150614179816007614f53565b82607f16901b8517945081608016600014156141b65761419a816001614e40565b865187906141a9908390614e40565b9052509395945050505050565b806141c081614e25565b915050614160565b50600080fd5b606060006141db8361414c565b905060008184600001516141ef9190614e40565b905083602001515181111561420357600080fd5b8167ffffffffffffffff81111561421c5761421c614a59565b6040519080825280601f01601f191660200182016040528015614246576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015614280578181015183820152614279602082614e40565b905061425e565b505050935250919050565b6000614296826145a1565b92915050565b60006020825111156142ad57600080fd5b60208201519050815160206142c29190615095565b6142cd906008614f53565b1c919050565b600081516020146142e357600080fd5b506020015190565b60008160058111156142ff576142ff61507f565b141561430e5761381e8261414c565b60028160058111156143225761432261507f565b141561036b5760006143338361414c565b905080836000018181516143479190614e40565b9052506020830151518351111561381e57600080fd5b606061436c84846000856145c9565b90505b9392505050565b6000808251604114156143ad5760208301516040840151606085015160001a6143a187828585614711565b945094505050506143df565b8251604014156143d757602083015160408401516143cc8683836147fe565b9350935050506143df565b506000905060025b9250929050565b60008160048111156143fa576143fa61507f565b14156144035750565b60018160048111156144175761441761507f565b14156144655760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610b9d565b60028160048111156144795761447961507f565b14156144c75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610b9d565b60038160048111156144db576144db61507f565b14156145345760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610b9d565b60048160048111156145485761454861507f565b1415610f825760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610b9d565b600081516014146145b157600080fd5b50602001516c01000000000000000000000000900490565b6060824710156146415760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610b9d565b6001600160a01b0385163b6146985760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b9d565b600080866001600160a01b031685876040516146b491906150d8565b60006040518083038185875af1925050503d80600081146146f1576040519150601f19603f3d011682016040523d82523d6000602084013e6146f6565b606091505b5091509150614706828286614850565b979650505050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561474857506000905060036147f5565b8460ff16601b1415801561476057508460ff16601c14155b1561477157506000905060046147f5565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156147c5573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166147ee576000600192509250506147f5565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83168161483460ff86901c601b614e40565b905061484287828885614711565b935093505050935093915050565b6060831561485f57508161436f565b82511561486f5782518084602001fd5b8160405162461bcd60e51b8152600401610b9d91906150f4565b60008083601f84011261489b57600080fd5b50813567ffffffffffffffff8111156148b357600080fd5b6020830191508360208260051b85010111156143df57600080fd5b600080600080604085870312156148e457600080fd5b843567ffffffffffffffff808211156148fc57600080fd5b61490888838901614889565b9096509450602087013591508082111561492157600080fd5b5061492e87828801614889565b95989497509550505050565b80356001600160a01b038116811461495157600080fd5b919050565b60006020828403121561496857600080fd5b61436f8261493a565b60006020828403121561498357600080fd5b5035919050565b803567ffffffffffffffff8116811461495157600080fd5b803563ffffffff8116811461495157600080fd5b600080600080600060a086880312156149ce57600080fd5b6149d78661493a565b9450602086013593506149ec6040870161498a565b92506149fa6060870161498a565b9150614a08608087016149a2565b90509295509295909350565b600060208284031215614a2657600080fd5b61436f826149a2565b60008060408385031215614a4257600080fd5b614a4b8361493a565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b60008060008060008060006080888a031215614a8a57600080fd5b873567ffffffffffffffff80821115614aa257600080fd5b818a0191508a601f830112614ab657600080fd5b813581811115614ac857614ac8614a59565b604051601f8201601f19908116603f01168101908382118183101715614af057614af0614a59565b816040528281528d6020848701011115614b0957600080fd5b82602086016020830137600094508460208483010152809b5050505060208a013581811115614b36578283fd5b614b428c828d01614889565b90995097505060408a013581811115614b59578283fd5b614b658c828d01614889565b90975095505060608a013581811115614b7c578283fd5b614b888c828d01614889565b9a9d999c50979a509598949794955050505050565b6000806000806000806000806080898b031215614bb957600080fd5b883567ffffffffffffffff80821115614bd157600080fd5b818b0191508b601f830112614be557600080fd5b813581811115614bf457600080fd5b8c6020828501011115614c0657600080fd5b60209283019a509850908a01359080821115614c2157600080fd5b614c2d8c838d01614889565b909850965060408b0135915080821115614c4657600080fd5b614c528c838d01614889565b909650945060608b0135915080821115614c6b57600080fd5b50614c788b828c01614889565b999c989b5096995094979396929594505050565b60008060008060008060c08789031215614ca557600080fd5b614cae8761493a565b9550614cbc6020880161493a565b945060408701359350614cd16060880161498a565b9250614cdf6080880161498a565b9150614ced60a088016149a2565b90509295509295509295565b600080600080600080600080600080600060c08c8e031215614d1a57600080fd5b8b359a5067ffffffffffffffff8060208e01351115614d3857600080fd5b614d488e60208f01358f01614889565b909b50995060408d0135811015614d5e57600080fd5b614d6e8e60408f01358f01614889565b909950975060608d0135811015614d8457600080fd5b614d948e60608f01358f01614889565b909750955060808d0135811015614daa57600080fd5b614dba8e60808f01358f01614889565b909550935060a08d0135811015614dd057600080fd5b50614de18d60a08e01358e01614889565b81935080925050509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415614e3957614e39614e0f565b5060010190565b60008219821115614e5357614e53614e0f565b500190565b600067ffffffffffffffff808316818516808303821115614e7b57614e7b614e0f565b01949350505050565b60008160005b84811015614eb9576001600160a01b03614ea38361493a565b1686526020958601959190910190600101614e8a565b5093949350505050565b60006001600160fb1b03831115614ed957600080fd5b8260051b8083863760009401938452509192915050565b6000614f07614f00838789614e84565b8486614ec3565b9695505050505050565b838152818360208301376000910160200190815292915050565b8681528560208201526000614f47614f00604084018789614e84565b98975050505050505050565b6000816000190483118215151615614f6d57614f6d614e0f565b500290565b600082614f8f57634e487b7160e01b600052601260045260246000fd5b500490565b6000808335601e19843603018112614fab57600080fd5b83018035915067ffffffffffffffff821115614fc657600080fd5b6020019150368190038213156143df57600080fd5b6040808252810184905260008560608301825b8781101561501c576001600160a01b036150078461493a565b16825260209283019290910190600101614fee565b5083810360208501528481526001600160fb1b0385111561503c57600080fd5b8460051b915081866020830137600091016020019081529695505050505050565b60006020828403121561506f57600080fd5b8151801515811461436f57600080fd5b634e487b7160e01b600052602160045260246000fd5b6000828210156150a7576150a7614e0f565b500390565b60005b838110156150c75781810151838201526020016150af565b838111156123cd5750506000910152565b600082516150ea8184602087016150ac565b9190910192915050565b60208152600082518060208401526151138160408501602087016150ac565b601f01601f1916919091016040019291505056fe4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a26469706673582212206d60a51d5002002f894e83a39519759341e663ecfa76f95b11dd8e5f9891512864736f6c63430008090033",
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
