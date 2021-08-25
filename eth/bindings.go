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

// StakingDelegatorInfo is an auto generated low-level Go binding around an user-defined struct.
type StakingDelegatorInfo struct {
	ValAddr            common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	Undelegations      []StakingUndelegation
}

// StakingUndelegation is an auto generated low-level Go binding around an user-defined struct.
type StakingUndelegation struct {
	Shares        *big.Int
	CreationBlock *big.Int
}

// GovernMetaData contains all meta data concerning the Govern contract.
var GovernMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_governProposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_governVoteTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorBondInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ConfirmParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CreateParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGovern.VoteOption\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"VoteParam\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"celerToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"createParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getParamProposalVote\",\"outputs\":[{\"internalType\":\"enumGovern.VoteOption\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"}],\"name\":\"getParamValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextParamProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paramProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610abf380380610abf83398101604081905261002f9161019c565b600080546001600160a01b0319166001600160a01b039b909b169a909a178a5560016020527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49989098557fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f969096557fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f949094557f7dfe757ecd65cbd7922a9c0161e935dd7fdbcc0e999689c7d31633896b1fc60b929092557fedc95719e9a3b28dd8e80877cb5880a9be7de1a13fc8b05e7999683b6b567643557fe2689cd4a84e23ad2f564004f1c9013e9589d260bde6380aba3ca7e09e4df40c557f8f331abe73332f95a25873e8b430885974c0409691f89d643119a11623a7924a557fdc686ec4a0ff239c70e7c7c36e8f853eced3bc8618f48d2b816da2a74311237e5560089091527f4db623e5c4870b62d3fc9b4e8f893a1a77627d75ab45d9ff7e56ba19564af99b55610227565b6000806000806000806000806000806101408b8d0312156101bc57600080fd5b8a516001600160a01b03811681146101d357600080fd5b809a505060208b0151985060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935060e08b015192506101008b015191506101208b015190509295989b9194979a5092959850565b610889806102366000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637e5fb8f31161005b5780637e5fb8f3146100fb578063c6c21e9d14610161578063e478ed9d1461018c578063eb505dd5146101a157600080fd5b806322da792714610082578063410ae02c1461009e578063581c53c5146100b1575b600080fd5b61008b60035481565b6040519081526020015b60405180910390f35b61008b6100ac366004610632565b6101c1565b6100ee6100bf366004610690565b60008281526002602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b604051610095919061074e565b61014f610109366004610677565b60026020819052600091825260409091208054600182015492820154600383015460048401546005909401546001600160a01b039093169493919260ff91821692911686565b604051610095969594939291906106fc565b600054610174906001600160a01b031681565b6040516001600160a01b039091168152602001610095565b61019f61019a36600461064d565b610200565b005b61008b6101af366004610632565b60016020526000908152604090205481565b6000600160008360088111156101d9576101d961083d565b60088111156101ea576101ea61083d565b8152602001908152602001600020549050919050565b60035460008181526002602052604090209061021d9060016107cc565b600355600160208190527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb495482546001600160a01b0319163390811784558383018290556000929092527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f5461029390436107cc565b600284015560038301805486919060ff191660018360088111156102b9576102b961083d565b02179055506004830184905560058301805460ff191660011790556000546102ec906001600160a01b0316833084610346565b7f4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160600160035461031c91906107e4565b8383866002015489896040516103379695949392919061079b565b60405180910390a15050505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526103a09085906103a6565b50505050565b60006103fb826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166104829092919063ffffffff16565b80519091501561047d57808060200190518101906104199190610610565b61047d5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084015b60405180910390fd5b505050565b6060610491848460008561049b565b90505b9392505050565b6060824710156104fc5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610474565b843b61054a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610474565b600080866001600160a01b0316858760405161056691906106e0565b60006040518083038185875af1925050503d80600081146105a3576040519150601f19603f3d011682016040523d82523d6000602084013e6105a8565b606091505b50915091506105b88282866105c3565b979650505050505050565b606083156105d2575081610494565b8251156105e25782518084602001fd5b8160405162461bcd60e51b81526004016104749190610768565b80356009811061060b57600080fd5b919050565b60006020828403121561062257600080fd5b8151801515811461049457600080fd5b60006020828403121561064457600080fd5b610494826105fc565b6000806040838503121561066057600080fd5b610669836105fc565b946020939093013593505050565b60006020828403121561068957600080fd5b5035919050565b600080604083850312156106a357600080fd5b8235915060208301356001600160a01b03811681146106c157600080fd5b809150509250929050565b600981106106dc576106dc61083d565b9052565b600082516106f28184602087016107fb565b9190910192915050565b6001600160a01b0387168152602081018690526040810185905260c0810161072760608301866106cc565b8360808301526003831061073d5761073d61083d565b8260a0830152979650505050505050565b60208101600483106107625761076261083d565b91905290565b60208152600082518060208401526107878160408501602087016107fb565b601f01601f19169190910160400192915050565b8681526001600160a01b0386166020820152604081018590526060810184905260c0810161073d60808301856106cc565b600082198211156107df576107df610827565b500190565b6000828210156107f6576107f6610827565b500390565b60005b838110156108165781810151838201526020016107fe565b838111156103a05750506000910152565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fdfea2646970667358221220170013b7dba46934fe2a6c044e9224266e594e3f016533a4212bb2e76035bb7664736f6c63430008070033",
}

// GovernABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernMetaData.ABI instead.
var GovernABI = GovernMetaData.ABI

// GovernBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernMetaData.Bin instead.
var GovernBin = GovernMetaData.Bin

// DeployGovern deploys a new Ethereum contract, binding an instance of Govern to it.
func DeployGovern(auth *bind.TransactOpts, backend bind.ContractBackend, _celerTokenAddress common.Address, _governProposalDeposit *big.Int, _governVoteTimeout *big.Int, _slashTimeout *big.Int, _maxBondedValidators *big.Int, _minValidatorTokens *big.Int, _minSelfDelegation *big.Int, _advanceNoticePeriod *big.Int, _validatorBondInterval *big.Int, _maxSlashFactor *big.Int) (common.Address, *types.Transaction, *Govern, error) {
	parsed, err := GovernMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernBin), backend, _celerTokenAddress, _governProposalDeposit, _governVoteTimeout, _slashTimeout, _maxBondedValidators, _minValidatorTokens, _minSelfDelegation, _advanceNoticePeriod, _validatorBondInterval, _maxSlashFactor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Govern{GovernCaller: GovernCaller{contract: contract}, GovernTransactor: GovernTransactor{contract: contract}, GovernFilterer: GovernFilterer{contract: contract}}, nil
}

// Govern is an auto generated Go binding around an Ethereum contract.
type Govern struct {
	GovernCaller     // Read-only binding to the contract
	GovernTransactor // Write-only binding to the contract
	GovernFilterer   // Log filterer for contract events
}

// GovernCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernSession struct {
	Contract     *Govern           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernCallerSession struct {
	Contract *GovernCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovernTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernTransactorSession struct {
	Contract     *GovernTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernRaw struct {
	Contract *Govern // Generic contract binding to access the raw methods on
}

// GovernCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernCallerRaw struct {
	Contract *GovernCaller // Generic read-only contract binding to access the raw methods on
}

// GovernTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernTransactorRaw struct {
	Contract *GovernTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovern creates a new instance of Govern, bound to a specific deployed contract.
func NewGovern(address common.Address, backend bind.ContractBackend) (*Govern, error) {
	contract, err := bindGovern(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Govern{GovernCaller: GovernCaller{contract: contract}, GovernTransactor: GovernTransactor{contract: contract}, GovernFilterer: GovernFilterer{contract: contract}}, nil
}

// NewGovernCaller creates a new read-only instance of Govern, bound to a specific deployed contract.
func NewGovernCaller(address common.Address, caller bind.ContractCaller) (*GovernCaller, error) {
	contract, err := bindGovern(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernCaller{contract: contract}, nil
}

// NewGovernTransactor creates a new write-only instance of Govern, bound to a specific deployed contract.
func NewGovernTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernTransactor, error) {
	contract, err := bindGovern(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernTransactor{contract: contract}, nil
}

// NewGovernFilterer creates a new log filterer instance of Govern, bound to a specific deployed contract.
func NewGovernFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernFilterer, error) {
	contract, err := bindGovern(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernFilterer{contract: contract}, nil
}

// bindGovern binds a generic wrapper to an already deployed contract.
func bindGovern(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Govern *GovernRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Govern.Contract.GovernCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Govern *GovernRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Govern.Contract.GovernTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Govern *GovernRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Govern.Contract.GovernTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Govern *GovernCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Govern.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Govern *GovernTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Govern.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Govern *GovernTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Govern.Contract.contract.Transact(opts, method, params...)
}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Govern *GovernCaller) CelerToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "celerToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Govern *GovernSession) CelerToken() (common.Address, error) {
	return _Govern.Contract.CelerToken(&_Govern.CallOpts)
}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Govern *GovernCallerSession) CelerToken() (common.Address, error) {
	return _Govern.Contract.CelerToken(&_Govern.CallOpts)
}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Govern *GovernCaller) GetParamProposalVote(opts *bind.CallOpts, _proposalId *big.Int, _voter common.Address) (uint8, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "getParamProposalVote", _proposalId, _voter)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Govern *GovernSession) GetParamProposalVote(_proposalId *big.Int, _voter common.Address) (uint8, error) {
	return _Govern.Contract.GetParamProposalVote(&_Govern.CallOpts, _proposalId, _voter)
}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Govern *GovernCallerSession) GetParamProposalVote(_proposalId *big.Int, _voter common.Address) (uint8, error) {
	return _Govern.Contract.GetParamProposalVote(&_Govern.CallOpts, _proposalId, _voter)
}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Govern *GovernCaller) GetParamValue(opts *bind.CallOpts, _name uint8) (*big.Int, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "getParamValue", _name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Govern *GovernSession) GetParamValue(_name uint8) (*big.Int, error) {
	return _Govern.Contract.GetParamValue(&_Govern.CallOpts, _name)
}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Govern *GovernCallerSession) GetParamValue(_name uint8) (*big.Int, error) {
	return _Govern.Contract.GetParamValue(&_Govern.CallOpts, _name)
}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Govern *GovernCaller) NextParamProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "nextParamProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Govern *GovernSession) NextParamProposalId() (*big.Int, error) {
	return _Govern.Contract.NextParamProposalId(&_Govern.CallOpts)
}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Govern *GovernCallerSession) NextParamProposalId() (*big.Int, error) {
	return _Govern.Contract.NextParamProposalId(&_Govern.CallOpts)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Govern *GovernCaller) ParamProposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "paramProposals", arg0)

	outstruct := new(struct {
		Proposer     common.Address
		Deposit      *big.Int
		VoteDeadline *big.Int
		Name         uint8
		NewValue     *big.Int
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VoteDeadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.NewValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Govern *GovernSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _Govern.Contract.ParamProposals(&_Govern.CallOpts, arg0)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Govern *GovernCallerSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _Govern.Contract.ParamProposals(&_Govern.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Govern *GovernCaller) Params(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "params", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Govern *GovernSession) Params(arg0 uint8) (*big.Int, error) {
	return _Govern.Contract.Params(&_Govern.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Govern *GovernCallerSession) Params(arg0 uint8) (*big.Int, error) {
	return _Govern.Contract.Params(&_Govern.CallOpts, arg0)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Govern *GovernTransactor) CreateParamProposal(opts *bind.TransactOpts, _name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Govern.contract.Transact(opts, "createParamProposal", _name, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Govern *GovernSession) CreateParamProposal(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Govern.Contract.CreateParamProposal(&_Govern.TransactOpts, _name, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Govern *GovernTransactorSession) CreateParamProposal(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Govern.Contract.CreateParamProposal(&_Govern.TransactOpts, _name, _value)
}

// GovernConfirmParamProposalIterator is returned from FilterConfirmParamProposal and is used to iterate over the raw logs and unpacked data for ConfirmParamProposal events raised by the Govern contract.
type GovernConfirmParamProposalIterator struct {
	Event *GovernConfirmParamProposal // Event containing the contract specifics and raw log

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
func (it *GovernConfirmParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernConfirmParamProposal)
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
		it.Event = new(GovernConfirmParamProposal)
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
func (it *GovernConfirmParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernConfirmParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernConfirmParamProposal represents a ConfirmParamProposal event raised by the Govern contract.
type GovernConfirmParamProposal struct {
	ProposalId *big.Int
	Passed     bool
	Name       uint8
	NewValue   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfirmParamProposal is a free log retrieval operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) FilterConfirmParamProposal(opts *bind.FilterOpts) (*GovernConfirmParamProposalIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return &GovernConfirmParamProposalIterator{contract: _Govern.contract, event: "ConfirmParamProposal", logs: logs, sub: sub}, nil
}

// WatchConfirmParamProposal is a free log subscription operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) WatchConfirmParamProposal(opts *bind.WatchOpts, sink chan<- *GovernConfirmParamProposal) (event.Subscription, error) {

	logs, sub, err := _Govern.contract.WatchLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernConfirmParamProposal)
				if err := _Govern.contract.UnpackLog(event, "ConfirmParamProposal", log); err != nil {
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

// ParseConfirmParamProposal is a log parse operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) ParseConfirmParamProposal(log types.Log) (*GovernConfirmParamProposal, error) {
	event := new(GovernConfirmParamProposal)
	if err := _Govern.contract.UnpackLog(event, "ConfirmParamProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernCreateParamProposalIterator is returned from FilterCreateParamProposal and is used to iterate over the raw logs and unpacked data for CreateParamProposal events raised by the Govern contract.
type GovernCreateParamProposalIterator struct {
	Event *GovernCreateParamProposal // Event containing the contract specifics and raw log

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
func (it *GovernCreateParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernCreateParamProposal)
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
		it.Event = new(GovernCreateParamProposal)
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
func (it *GovernCreateParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernCreateParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernCreateParamProposal represents a CreateParamProposal event raised by the Govern contract.
type GovernCreateParamProposal struct {
	ProposalId   *big.Int
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateParamProposal is a free log retrieval operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) FilterCreateParamProposal(opts *bind.FilterOpts) (*GovernCreateParamProposalIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return &GovernCreateParamProposalIterator{contract: _Govern.contract, event: "CreateParamProposal", logs: logs, sub: sub}, nil
}

// WatchCreateParamProposal is a free log subscription operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) WatchCreateParamProposal(opts *bind.WatchOpts, sink chan<- *GovernCreateParamProposal) (event.Subscription, error) {

	logs, sub, err := _Govern.contract.WatchLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernCreateParamProposal)
				if err := _Govern.contract.UnpackLog(event, "CreateParamProposal", log); err != nil {
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

// ParseCreateParamProposal is a log parse operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) ParseCreateParamProposal(log types.Log) (*GovernCreateParamProposal, error) {
	event := new(GovernCreateParamProposal)
	if err := _Govern.contract.UnpackLog(event, "CreateParamProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernVoteParamIterator is returned from FilterVoteParam and is used to iterate over the raw logs and unpacked data for VoteParam events raised by the Govern contract.
type GovernVoteParamIterator struct {
	Event *GovernVoteParam // Event containing the contract specifics and raw log

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
func (it *GovernVoteParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernVoteParam)
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
		it.Event = new(GovernVoteParam)
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
func (it *GovernVoteParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernVoteParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernVoteParam represents a VoteParam event raised by the Govern contract.
type GovernVoteParam struct {
	ProposalId *big.Int
	Voter      common.Address
	Vote       uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteParam is a free log retrieval operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Govern *GovernFilterer) FilterVoteParam(opts *bind.FilterOpts) (*GovernVoteParamIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return &GovernVoteParamIterator{contract: _Govern.contract, event: "VoteParam", logs: logs, sub: sub}, nil
}

// WatchVoteParam is a free log subscription operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Govern *GovernFilterer) WatchVoteParam(opts *bind.WatchOpts, sink chan<- *GovernVoteParam) (event.Subscription, error) {

	logs, sub, err := _Govern.contract.WatchLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernVoteParam)
				if err := _Govern.contract.UnpackLog(event, "VoteParam", log); err != nil {
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

// ParseVoteParam is a log parse operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Govern *GovernFilterer) ParseVoteParam(log types.Log) (*GovernVoteParam, error) {
	event := new(GovernVoteParam)
	if err := _Govern.contract.UnpackLog(event, "VoteParam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbMetaData contains all meta data concerning the Pb contract.
var PbMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f28e6f2aa01bb7e396f4dd76e55cf5279e7c3564e43d35de33741a643e8c91aa64736f6c63430008070033",
}

// PbABI is the input ABI used to generate the binding from.
// Deprecated: Use PbMetaData.ABI instead.
var PbABI = PbMetaData.ABI

// PbBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PbMetaData.Bin instead.
var PbBin = PbMetaData.Bin

// DeployPb deploys a new Ethereum contract, binding an instance of Pb to it.
func DeployPb(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pb, error) {
	parsed, err := PbMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PbBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pb{PbCaller: PbCaller{contract: contract}, PbTransactor: PbTransactor{contract: contract}, PbFilterer: PbFilterer{contract: contract}}, nil
}

// Pb is an auto generated Go binding around an Ethereum contract.
type Pb struct {
	PbCaller     // Read-only binding to the contract
	PbTransactor // Write-only binding to the contract
	PbFilterer   // Log filterer for contract events
}

// PbCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbSession struct {
	Contract     *Pb               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbCallerSession struct {
	Contract *PbCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbTransactorSession struct {
	Contract     *PbTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRaw struct {
	Contract *Pb // Generic contract binding to access the raw methods on
}

// PbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbCallerRaw struct {
	Contract *PbCaller // Generic read-only contract binding to access the raw methods on
}

// PbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbTransactorRaw struct {
	Contract *PbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPb creates a new instance of Pb, bound to a specific deployed contract.
func NewPb(address common.Address, backend bind.ContractBackend) (*Pb, error) {
	contract, err := bindPb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pb{PbCaller: PbCaller{contract: contract}, PbTransactor: PbTransactor{contract: contract}, PbFilterer: PbFilterer{contract: contract}}, nil
}

// NewPbCaller creates a new read-only instance of Pb, bound to a specific deployed contract.
func NewPbCaller(address common.Address, caller bind.ContractCaller) (*PbCaller, error) {
	contract, err := bindPb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbCaller{contract: contract}, nil
}

// NewPbTransactor creates a new write-only instance of Pb, bound to a specific deployed contract.
func NewPbTransactor(address common.Address, transactor bind.ContractTransactor) (*PbTransactor, error) {
	contract, err := bindPb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbTransactor{contract: contract}, nil
}

// NewPbFilterer creates a new log filterer instance of Pb, bound to a specific deployed contract.
func NewPbFilterer(address common.Address, filterer bind.ContractFilterer) (*PbFilterer, error) {
	contract, err := bindPb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbFilterer{contract: contract}, nil
}

// bindPb binds a generic wrapper to an already deployed contract.
func bindPb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pb *PbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pb.Contract.PbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pb *PbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pb.Contract.PbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pb *PbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pb.Contract.PbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pb *PbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pb *PbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pb *PbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pb.Contract.contract.Transact(opts, method, params...)
}

// PbSgnMetaData contains all meta data concerning the PbSgn contract.
var PbSgnMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d2da8b89abfda8e91392170e0f63fc72528abd3a0e9d4093bf5dbcecb7b04ca164736f6c63430008070033",
}

// PbSgnABI is the input ABI used to generate the binding from.
// Deprecated: Use PbSgnMetaData.ABI instead.
var PbSgnABI = PbSgnMetaData.ABI

// PbSgnBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PbSgnMetaData.Bin instead.
var PbSgnBin = PbSgnMetaData.Bin

// DeployPbSgn deploys a new Ethereum contract, binding an instance of PbSgn to it.
func DeployPbSgn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbSgn, error) {
	parsed, err := PbSgnMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PbSgnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbSgn{PbSgnCaller: PbSgnCaller{contract: contract}, PbSgnTransactor: PbSgnTransactor{contract: contract}, PbSgnFilterer: PbSgnFilterer{contract: contract}}, nil
}

// PbSgn is an auto generated Go binding around an Ethereum contract.
type PbSgn struct {
	PbSgnCaller     // Read-only binding to the contract
	PbSgnTransactor // Write-only binding to the contract
	PbSgnFilterer   // Log filterer for contract events
}

// PbSgnCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbSgnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSgnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbSgnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSgnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbSgnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSgnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbSgnSession struct {
	Contract     *PbSgn            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbSgnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbSgnCallerSession struct {
	Contract *PbSgnCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PbSgnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbSgnTransactorSession struct {
	Contract     *PbSgnTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbSgnRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbSgnRaw struct {
	Contract *PbSgn // Generic contract binding to access the raw methods on
}

// PbSgnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbSgnCallerRaw struct {
	Contract *PbSgnCaller // Generic read-only contract binding to access the raw methods on
}

// PbSgnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbSgnTransactorRaw struct {
	Contract *PbSgnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbSgn creates a new instance of PbSgn, bound to a specific deployed contract.
func NewPbSgn(address common.Address, backend bind.ContractBackend) (*PbSgn, error) {
	contract, err := bindPbSgn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbSgn{PbSgnCaller: PbSgnCaller{contract: contract}, PbSgnTransactor: PbSgnTransactor{contract: contract}, PbSgnFilterer: PbSgnFilterer{contract: contract}}, nil
}

// NewPbSgnCaller creates a new read-only instance of PbSgn, bound to a specific deployed contract.
func NewPbSgnCaller(address common.Address, caller bind.ContractCaller) (*PbSgnCaller, error) {
	contract, err := bindPbSgn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbSgnCaller{contract: contract}, nil
}

// NewPbSgnTransactor creates a new write-only instance of PbSgn, bound to a specific deployed contract.
func NewPbSgnTransactor(address common.Address, transactor bind.ContractTransactor) (*PbSgnTransactor, error) {
	contract, err := bindPbSgn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbSgnTransactor{contract: contract}, nil
}

// NewPbSgnFilterer creates a new log filterer instance of PbSgn, bound to a specific deployed contract.
func NewPbSgnFilterer(address common.Address, filterer bind.ContractFilterer) (*PbSgnFilterer, error) {
	contract, err := bindPbSgn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbSgnFilterer{contract: contract}, nil
}

// bindPbSgn binds a generic wrapper to an already deployed contract.
func bindPbSgn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbSgnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbSgn *PbSgnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbSgn.Contract.PbSgnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbSgn *PbSgnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbSgn.Contract.PbSgnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbSgn *PbSgnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbSgn.Contract.PbSgnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbSgn *PbSgnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbSgn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbSgn *PbSgnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbSgn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbSgn *PbSgnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbSgn.Contract.contract.Transact(opts, method, params...)
}

// PbStakingMetaData contains all meta data concerning the PbStaking contract.
var PbStakingMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e1ced39dfdcc3362ae2d4a2410effc9fa88f704cf2ffc31e13abd9b476871b7d64736f6c63430008070033",
}

// PbStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use PbStakingMetaData.ABI instead.
var PbStakingABI = PbStakingMetaData.ABI

// PbStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PbStakingMetaData.Bin instead.
var PbStakingBin = PbStakingMetaData.Bin

// DeployPbStaking deploys a new Ethereum contract, binding an instance of PbStaking to it.
func DeployPbStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbStaking, error) {
	parsed, err := PbStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PbStakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbStaking{PbStakingCaller: PbStakingCaller{contract: contract}, PbStakingTransactor: PbStakingTransactor{contract: contract}, PbStakingFilterer: PbStakingFilterer{contract: contract}}, nil
}

// PbStaking is an auto generated Go binding around an Ethereum contract.
type PbStaking struct {
	PbStakingCaller     // Read-only binding to the contract
	PbStakingTransactor // Write-only binding to the contract
	PbStakingFilterer   // Log filterer for contract events
}

// PbStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbStakingSession struct {
	Contract     *PbStaking        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbStakingCallerSession struct {
	Contract *PbStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PbStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbStakingTransactorSession struct {
	Contract     *PbStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PbStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbStakingRaw struct {
	Contract *PbStaking // Generic contract binding to access the raw methods on
}

// PbStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbStakingCallerRaw struct {
	Contract *PbStakingCaller // Generic read-only contract binding to access the raw methods on
}

// PbStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbStakingTransactorRaw struct {
	Contract *PbStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbStaking creates a new instance of PbStaking, bound to a specific deployed contract.
func NewPbStaking(address common.Address, backend bind.ContractBackend) (*PbStaking, error) {
	contract, err := bindPbStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbStaking{PbStakingCaller: PbStakingCaller{contract: contract}, PbStakingTransactor: PbStakingTransactor{contract: contract}, PbStakingFilterer: PbStakingFilterer{contract: contract}}, nil
}

// NewPbStakingCaller creates a new read-only instance of PbStaking, bound to a specific deployed contract.
func NewPbStakingCaller(address common.Address, caller bind.ContractCaller) (*PbStakingCaller, error) {
	contract, err := bindPbStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbStakingCaller{contract: contract}, nil
}

// NewPbStakingTransactor creates a new write-only instance of PbStaking, bound to a specific deployed contract.
func NewPbStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*PbStakingTransactor, error) {
	contract, err := bindPbStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbStakingTransactor{contract: contract}, nil
}

// NewPbStakingFilterer creates a new log filterer instance of PbStaking, bound to a specific deployed contract.
func NewPbStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*PbStakingFilterer, error) {
	contract, err := bindPbStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbStakingFilterer{contract: contract}, nil
}

// bindPbStaking binds a generic wrapper to an already deployed contract.
func bindPbStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbStaking *PbStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbStaking.Contract.PbStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbStaking *PbStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbStaking.Contract.PbStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbStaking *PbStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbStaking.Contract.PbStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbStaking *PbStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbStaking *PbStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbStaking *PbStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbStaking.Contract.contract.Transact(opts, method, params...)
}

// SGNMetaData contains all meta data concerning the SGN contract.
var SGNMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStaking\",\"name\":\"_staking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddr\",\"type\":\"bytes\"}],\"name\":\"SgnAddrUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sgnAddrs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sgnAddr\",\"type\":\"bytes\"}],\"name\":\"updateSgnAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_withdrawalRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnAmts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405162001ad438038062001ad4833981016040819052610031916100ac565b61003a3361005c565b6000805460ff60a01b1916905560601b6001600160601b0319166080526100dc565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100be57600080fd5b81516001600160a01b03811681146100d557600080fd5b9392505050565b60805160601c6119b66200011e600039600081816101110152818161056e0152818161060b015281816106b20152818161086a015261096201526119b66000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063c429fe1f11610066578063c429fe1f146101ed578063d0bb93511461020d578063d88ef27114610220578063f2fde38b1461023357600080fd5b80638da5cb5b146101b65780639d4323be146101c7578063b02c43d0146101da57600080fd5b80635c975abb116100c85780635c975abb14610150578063715018a61461016d578063795c2c14146101755780638456cb59146101ae57600080fd5b80633f4ba83a146100ef57806347e7ef24146100f95780634cf088d91461010c575b600080fd5b6100f7610246565b005b6100f76101073660046114c5565b610283565b6101337f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b600054600160a01b900460ff166040519015158152602001610147565b6100f7610398565b6101a061018336600461148c565b600260209081526000928352604080842090915290825290205481565b604051908152602001610147565b6100f76103cc565b6000546001600160a01b0316610133565b6100f76101d53660046114c5565b6103fe565b6101a06101e8366004611610565b610490565b6102006101fb366004611452565b6104b1565b6040516101479190611795565b6100f761021b366004611513565b61054b565b6100f761022e366004611555565b610921565b6100f7610241366004611452565b610b64565b6000546001600160a01b031633146102795760405162461bcd60e51b815260040161027090611802565b60405180910390fd5b610281610bff565b565b600054600160a01b900460ff16156102ad5760405162461bcd60e51b8152600401610270906117d8565b6040516bffffffffffffffffffffffff1933606081811b8316602085015285901b9091166034830152604882018390529060019060680160408051601f198184030181529190528051602091820120825460018101845560009384529190922001556103246001600160a01b038416823085610c9c565b6001805460009161033491611890565b6040805167ffffffffffffffff831681526001600160a01b0385811660208301528716818301526060810186905290519192507f2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8919081900360800190a150505050565b6000546001600160a01b031633146103c25760405162461bcd60e51b815260040161027090611802565b6102816000610d0d565b6000546001600160a01b031633146103f65760405162461bcd60e51b815260040161027090611802565b610281610d5d565b600054600160a01b900460ff1661044e5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610270565b6000546001600160a01b031633146104785760405162461bcd60e51b815260040161027090611802565b61048c6001600160a01b0383163383610dc2565b5050565b600181815481106104a057600080fd5b600091825260209091200154905081565b600360205260009081526040902080546104ca906118d3565b80601f01602080910402602001604051908101604052809291908181526020018280546104f6906118d3565b80156105435780601f1061051857610100808354040283529160200191610543565b820191906000526020600020905b81548152906001019060200180831161052657829003601f168201915b505050505081565b604051636d30878360e01b81523360048201819052906000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636d3087839060240160206040518083038186803b1580156105b057600080fd5b505afa1580156105c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e8919061146f565b6001600160a01b03161461069057604051636d30878360e01b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636d3087839060240160206040518083038186803b15801561065557600080fd5b505afa158015610669573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068d919061146f565b90505b60405163a310624f60e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a310624f9060240160206040518083038186803b1580156106f657600080fd5b505afa15801561070a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072e91906115ef565b905060018160038111156107445761074461193f565b1461078a5760405162461bcd60e51b81526020600482015260166024820152752737ba103ab73137b73232b2103b30b634b230ba37b960511b6044820152606401610270565b6001600160a01b038216600090815260036020526040812080546107ad906118d3565b80601f01602080910402602001604051908101604052809291908181526020018280546107d9906118d3565b80156108265780601f106107fb57610100808354040283529160200191610826565b820191906000526020600020905b81548152906001019060200180831161080957829003601f168201915b505050506001600160a01b03851660009081526003602052604090209192506108529190508686611370565b506040516309146f1160e41b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639146f110906108a39086908990899060040161169a565b600060405180830381600087803b1580156108bd57600080fd5b505af11580156108d1573d6000803e3d6000fd5b50505050826001600160a01b03167f8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4828787604051610912939291906117a8565b60405180910390a25050505050565b600054600160a01b900460ff161561094b5760405162461bcd60e51b8152600401610270906117d8565b60405163453a6aff60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638a74d5fe9061099d9087908790879087906004016116e4565b60206040518083038186803b1580156109b557600080fd5b505afa1580156109c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ed91906114f1565b506000610a2f85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610df792505050565b80516001600160a01b039081166000908152600260209081526040808320828601519094168352929052818120549183015192935091610a6f9190611890565b905060008111610ac15760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f207769746864726177000000000000006044820152606401610270565b60408083015183516001600160a01b0390811660009081526002602090815284822081880180518516845291529390209190915583519151610b069291169083610dc2565b8151602080840151604080516001600160a01b039485168152939091169183019190915281018290527f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9060600160405180910390a1505050505050565b6000546001600160a01b03163314610b8e5760405162461bcd60e51b815260040161027090611802565b6001600160a01b038116610bf35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610270565b610bfc81610d0d565b50565b600054600160a01b900460ff16610c4f5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610270565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6040516001600160a01b0380851660248301528316604482015260648101829052610d079085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610ed2565b50505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff1615610d875760405162461bcd60e51b8152600401610270906117d8565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610c7f3390565b6040516001600160a01b038316602482015260448101829052610df290849063a9059cbb60e01b90606401610cd0565b505050565b604080516060810182526000808252602080830182905282840182905283518085019094528184528301849052909190805b60208301515183511015610eca57610e4083610fa4565b90925090508160011415610e6f57610e5f610e5a84610fde565b61109b565b6001600160a01b03168452610e29565b8160021415610e9757610e84610e5a84610fde565b6001600160a01b03166020850152610e29565b8160031415610ebb57610eb1610eac84610fde565b6110ac565b6040850152610e29565b610ec583826110e3565b610e29565b505050919050565b6000610f27826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111559092919063ffffffff16565b805190915015610df25780806020019051810190610f4591906114f1565b610df25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610270565b6000806000610fb28461116e565b9050610fbf60088261184f565b9250806007166005811115610fd657610fd661193f565b915050915091565b60606000610feb8361116e565b90506000818460000151610fff9190611837565b905083602001515181111561101357600080fd5b8167ffffffffffffffff81111561102c5761102c611955565b6040519080825280601f01601f191660200182016040528015611056576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015611090578181015183820152611089602082611837565b905061106e565b505050935250919050565b60006110a6826111f0565b92915050565b60006020825111156110bd57600080fd5b60208201519050815160206110d29190611890565b6110dd906008611871565b1c919050565b60008160058111156110f7576110f761193f565b141561110657610df28261116e565b600281600581111561111a5761111a61193f565b14156100ea57600061112b8361116e565b9050808360000181815161113f9190611837565b90525060208301515183511115610df257600080fd5b6060611164848460008561120f565b90505b9392505050565b602080820151825181019091015160009182805b600a8110156111ea5783811a915061119b816007611871565b82607f16901b8517945081608016600014156111d8576111bc816001611837565b865187906111cb908390611837565b9052509395945050505050565b806111e28161190e565b915050611182565b50600080fd5b6000815160141461120057600080fd5b5060200151600160601b900490565b6060824710156112705760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610270565b843b6112be5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610270565b600080866001600160a01b031685876040516112da919061167e565b60006040518083038185875af1925050503d8060008114611317576040519150601f19603f3d011682016040523d82523d6000602084013e61131c565b606091505b509150915061132c828286611337565b979650505050505050565b60608315611346575081611167565b8251156113565782518084602001fd5b8160405162461bcd60e51b81526004016102709190611795565b82805461137c906118d3565b90600052602060002090601f01602090048101928261139e57600085556113e4565b82601f106113b75782800160ff198235161785556113e4565b828001600101855582156113e4579182015b828111156113e45782358255916020019190600101906113c9565b506113f09291506113f4565b5090565b5b808211156113f057600081556001016113f5565b60008083601f84011261141b57600080fd5b50813567ffffffffffffffff81111561143357600080fd5b60208301915083602082850101111561144b57600080fd5b9250929050565b60006020828403121561146457600080fd5b81356111678161196b565b60006020828403121561148157600080fd5b81516111678161196b565b6000806040838503121561149f57600080fd5b82356114aa8161196b565b915060208301356114ba8161196b565b809150509250929050565b600080604083850312156114d857600080fd5b82356114e38161196b565b946020939093013593505050565b60006020828403121561150357600080fd5b8151801515811461116757600080fd5b6000806020838503121561152657600080fd5b823567ffffffffffffffff81111561153d57600080fd5b61154985828601611409565b90969095509350505050565b6000806000806040858703121561156b57600080fd5b843567ffffffffffffffff8082111561158357600080fd5b61158f88838901611409565b909650945060208701359150808211156115a857600080fd5b818701915087601f8301126115bc57600080fd5b8135818111156115cb57600080fd5b8860208260051b85010111156115e057600080fd5b95989497505060200194505050565b60006020828403121561160157600080fd5b81516004811061116757600080fd5b60006020828403121561162257600080fd5b5035919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000815180845261166a8160208601602086016118a7565b601f01601f19169290920160200192915050565b600082516116908184602087016118a7565b9190910192915050565b60018060a01b038416815260606020820152600860608201526739b3b716b0b2323960c11b608082015260a0604082015260006116db60a083018486611629565b95945050505050565b6040815260006116f8604083018688611629565b602083820381850152818583528183019050818660051b8401018760005b8881101561178557858303601f190184528135368b9003601e1901811261173c57600080fd5b8a01803567ffffffffffffffff81111561175557600080fd5b8036038c131561176457600080fd5b6117718582898501611629565b958701959450505090840190600101611716565b50909a9950505050505050505050565b6020815260006111676020830184611652565b6040815260006117bb6040830186611652565b82810360208401526117ce818587611629565b9695505050505050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6000821982111561184a5761184a611929565b500190565b60008261186c57634e487b7160e01b600052601260045260246000fd5b500490565b600081600019048311821515161561188b5761188b611929565b500290565b6000828210156118a2576118a2611929565b500390565b60005b838110156118c25781810151838201526020016118aa565b83811115610d075750506000910152565b600181811c908216806118e757607f821691505b6020821081141561190857634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561192257611922611929565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610bfc57600080fdfea2646970667358221220d609e2f9769d3e5cae39993096bf8a4c09e96bbed385ed67fc70ba0762146d9364736f6c63430008070033",
}

// SGNABI is the input ABI used to generate the binding from.
// Deprecated: Use SGNMetaData.ABI instead.
var SGNABI = SGNMetaData.ABI

// SGNBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SGNMetaData.Bin instead.
var SGNBin = SGNMetaData.Bin

// DeploySGN deploys a new Ethereum contract, binding an instance of SGN to it.
func DeploySGN(auth *bind.TransactOpts, backend bind.ContractBackend, _staking common.Address) (common.Address, *types.Transaction, *SGN, error) {
	parsed, err := SGNMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SGNBin), backend, _staking)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SGN{SGNCaller: SGNCaller{contract: contract}, SGNTransactor: SGNTransactor{contract: contract}, SGNFilterer: SGNFilterer{contract: contract}}, nil
}

// SGN is an auto generated Go binding around an Ethereum contract.
type SGN struct {
	SGNCaller     // Read-only binding to the contract
	SGNTransactor // Write-only binding to the contract
	SGNFilterer   // Log filterer for contract events
}

// SGNCaller is an auto generated read-only Go binding around an Ethereum contract.
type SGNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGNTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SGNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGNFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SGNFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGNSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SGNSession struct {
	Contract     *SGN              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SGNCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SGNCallerSession struct {
	Contract *SGNCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SGNTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SGNTransactorSession struct {
	Contract     *SGNTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SGNRaw is an auto generated low-level Go binding around an Ethereum contract.
type SGNRaw struct {
	Contract *SGN // Generic contract binding to access the raw methods on
}

// SGNCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SGNCallerRaw struct {
	Contract *SGNCaller // Generic read-only contract binding to access the raw methods on
}

// SGNTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SGNTransactorRaw struct {
	Contract *SGNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSGN creates a new instance of SGN, bound to a specific deployed contract.
func NewSGN(address common.Address, backend bind.ContractBackend) (*SGN, error) {
	contract, err := bindSGN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SGN{SGNCaller: SGNCaller{contract: contract}, SGNTransactor: SGNTransactor{contract: contract}, SGNFilterer: SGNFilterer{contract: contract}}, nil
}

// NewSGNCaller creates a new read-only instance of SGN, bound to a specific deployed contract.
func NewSGNCaller(address common.Address, caller bind.ContractCaller) (*SGNCaller, error) {
	contract, err := bindSGN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SGNCaller{contract: contract}, nil
}

// NewSGNTransactor creates a new write-only instance of SGN, bound to a specific deployed contract.
func NewSGNTransactor(address common.Address, transactor bind.ContractTransactor) (*SGNTransactor, error) {
	contract, err := bindSGN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SGNTransactor{contract: contract}, nil
}

// NewSGNFilterer creates a new log filterer instance of SGN, bound to a specific deployed contract.
func NewSGNFilterer(address common.Address, filterer bind.ContractFilterer) (*SGNFilterer, error) {
	contract, err := bindSGN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SGNFilterer{contract: contract}, nil
}

// bindSGN binds a generic wrapper to an already deployed contract.
func bindSGN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SGNABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SGN *SGNRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SGN.Contract.SGNCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SGN *SGNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.Contract.SGNTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SGN *SGNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SGN.Contract.SGNTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SGN *SGNCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SGN.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SGN *SGNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SGN *SGNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SGN.Contract.contract.Transact(opts, method, params...)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_SGN *SGNCaller) Deposits(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "deposits", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_SGN *SGNSession) Deposits(arg0 *big.Int) ([32]byte, error) {
	return _SGN.Contract.Deposits(&_SGN.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_SGN *SGNCallerSession) Deposits(arg0 *big.Int) ([32]byte, error) {
	return _SGN.Contract.Deposits(&_SGN.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGN *SGNCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGN *SGNSession) Owner() (common.Address, error) {
	return _SGN.Contract.Owner(&_SGN.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGN *SGNCallerSession) Owner() (common.Address, error) {
	return _SGN.Contract.Owner(&_SGN.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGN *SGNCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGN *SGNSession) Paused() (bool, error) {
	return _SGN.Contract.Paused(&_SGN.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGN *SGNCallerSession) Paused() (bool, error) {
	return _SGN.Contract.Paused(&_SGN.CallOpts)
}

// SgnAddrs is a free data retrieval call binding the contract method 0xc429fe1f.
//
// Solidity: function sgnAddrs(address ) view returns(bytes)
func (_SGN *SGNCaller) SgnAddrs(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "sgnAddrs", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SgnAddrs is a free data retrieval call binding the contract method 0xc429fe1f.
//
// Solidity: function sgnAddrs(address ) view returns(bytes)
func (_SGN *SGNSession) SgnAddrs(arg0 common.Address) ([]byte, error) {
	return _SGN.Contract.SgnAddrs(&_SGN.CallOpts, arg0)
}

// SgnAddrs is a free data retrieval call binding the contract method 0xc429fe1f.
//
// Solidity: function sgnAddrs(address ) view returns(bytes)
func (_SGN *SGNCallerSession) SgnAddrs(arg0 common.Address) ([]byte, error) {
	return _SGN.Contract.SgnAddrs(&_SGN.CallOpts, arg0)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_SGN *SGNCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_SGN *SGNSession) Staking() (common.Address, error) {
	return _SGN.Contract.Staking(&_SGN.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_SGN *SGNCallerSession) Staking() (common.Address, error) {
	return _SGN.Contract.Staking(&_SGN.CallOpts)
}

// WithdrawnAmts is a free data retrieval call binding the contract method 0x795c2c14.
//
// Solidity: function withdrawnAmts(address , address ) view returns(uint256)
func (_SGN *SGNCaller) WithdrawnAmts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "withdrawnAmts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnAmts is a free data retrieval call binding the contract method 0x795c2c14.
//
// Solidity: function withdrawnAmts(address , address ) view returns(uint256)
func (_SGN *SGNSession) WithdrawnAmts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SGN.Contract.WithdrawnAmts(&_SGN.CallOpts, arg0, arg1)
}

// WithdrawnAmts is a free data retrieval call binding the contract method 0x795c2c14.
//
// Solidity: function withdrawnAmts(address , address ) view returns(uint256)
func (_SGN *SGNCallerSession) WithdrawnAmts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SGN.Contract.WithdrawnAmts(&_SGN.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_SGN *SGNTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_SGN *SGNSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.Deposit(&_SGN.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_SGN *SGNTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.Deposit(&_SGN.TransactOpts, _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_SGN *SGNTransactor) DrainToken(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "drainToken", _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_SGN *SGNSession) DrainToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.DrainToken(&_SGN.TransactOpts, _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_SGN *SGNTransactorSession) DrainToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.DrainToken(&_SGN.TransactOpts, _token, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SGN *SGNTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SGN *SGNSession) Pause() (*types.Transaction, error) {
	return _SGN.Contract.Pause(&_SGN.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SGN *SGNTransactorSession) Pause() (*types.Transaction, error) {
	return _SGN.Contract.Pause(&_SGN.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGN *SGNTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGN *SGNSession) RenounceOwnership() (*types.Transaction, error) {
	return _SGN.Contract.RenounceOwnership(&_SGN.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGN *SGNTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SGN.Contract.RenounceOwnership(&_SGN.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGN *SGNTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGN *SGNSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SGN.Contract.TransferOwnership(&_SGN.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGN *SGNTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SGN.Contract.TransferOwnership(&_SGN.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SGN *SGNTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SGN *SGNSession) Unpause() (*types.Transaction, error) {
	return _SGN.Contract.Unpause(&_SGN.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SGN *SGNTransactorSession) Unpause() (*types.Transaction, error) {
	return _SGN.Contract.Unpause(&_SGN.TransactOpts)
}

// UpdateSgnAddr is a paid mutator transaction binding the contract method 0xd0bb9351.
//
// Solidity: function updateSgnAddr(bytes _sgnAddr) returns()
func (_SGN *SGNTransactor) UpdateSgnAddr(opts *bind.TransactOpts, _sgnAddr []byte) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "updateSgnAddr", _sgnAddr)
}

// UpdateSgnAddr is a paid mutator transaction binding the contract method 0xd0bb9351.
//
// Solidity: function updateSgnAddr(bytes _sgnAddr) returns()
func (_SGN *SGNSession) UpdateSgnAddr(_sgnAddr []byte) (*types.Transaction, error) {
	return _SGN.Contract.UpdateSgnAddr(&_SGN.TransactOpts, _sgnAddr)
}

// UpdateSgnAddr is a paid mutator transaction binding the contract method 0xd0bb9351.
//
// Solidity: function updateSgnAddr(bytes _sgnAddr) returns()
func (_SGN *SGNTransactorSession) UpdateSgnAddr(_sgnAddr []byte) (*types.Transaction, error) {
	return _SGN.Contract.UpdateSgnAddr(&_SGN.TransactOpts, _sgnAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd88ef271.
//
// Solidity: function withdraw(bytes _withdrawalRequest, bytes[] _sigs) returns()
func (_SGN *SGNTransactor) Withdraw(opts *bind.TransactOpts, _withdrawalRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "withdraw", _withdrawalRequest, _sigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd88ef271.
//
// Solidity: function withdraw(bytes _withdrawalRequest, bytes[] _sigs) returns()
func (_SGN *SGNSession) Withdraw(_withdrawalRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _SGN.Contract.Withdraw(&_SGN.TransactOpts, _withdrawalRequest, _sigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd88ef271.
//
// Solidity: function withdraw(bytes _withdrawalRequest, bytes[] _sigs) returns()
func (_SGN *SGNTransactorSession) Withdraw(_withdrawalRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _SGN.Contract.Withdraw(&_SGN.TransactOpts, _withdrawalRequest, _sigs)
}

// SGNDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SGN contract.
type SGNDepositIterator struct {
	Event *SGNDeposit // Event containing the contract specifics and raw log

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
func (it *SGNDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNDeposit)
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
		it.Event = new(SGNDeposit)
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
func (it *SGNDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNDeposit represents a Deposit event raised by the SGN contract.
type SGNDeposit struct {
	DepositId *big.Int
	Account   common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8.
//
// Solidity: event Deposit(uint256 depositId, address account, address token, uint256 amount)
func (_SGN *SGNFilterer) FilterDeposit(opts *bind.FilterOpts) (*SGNDepositIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &SGNDepositIterator{contract: _SGN.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8.
//
// Solidity: event Deposit(uint256 depositId, address account, address token, uint256 amount)
func (_SGN *SGNFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SGNDeposit) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNDeposit)
				if err := _SGN.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8.
//
// Solidity: event Deposit(uint256 depositId, address account, address token, uint256 amount)
func (_SGN *SGNFilterer) ParseDeposit(log types.Log) (*SGNDeposit, error) {
	event := new(SGNDeposit)
	if err := _SGN.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SGN contract.
type SGNOwnershipTransferredIterator struct {
	Event *SGNOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SGNOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNOwnershipTransferred)
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
		it.Event = new(SGNOwnershipTransferred)
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
func (it *SGNOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNOwnershipTransferred represents a OwnershipTransferred event raised by the SGN contract.
type SGNOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SGN *SGNFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SGNOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SGN.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SGNOwnershipTransferredIterator{contract: _SGN.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SGN *SGNFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SGNOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SGN.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNOwnershipTransferred)
				if err := _SGN.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SGN *SGNFilterer) ParseOwnershipTransferred(log types.Log) (*SGNOwnershipTransferred, error) {
	event := new(SGNOwnershipTransferred)
	if err := _SGN.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SGN contract.
type SGNPausedIterator struct {
	Event *SGNPaused // Event containing the contract specifics and raw log

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
func (it *SGNPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNPaused)
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
		it.Event = new(SGNPaused)
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
func (it *SGNPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNPaused represents a Paused event raised by the SGN contract.
type SGNPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SGN *SGNFilterer) FilterPaused(opts *bind.FilterOpts) (*SGNPausedIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SGNPausedIterator{contract: _SGN.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SGN *SGNFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SGNPaused) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNPaused)
				if err := _SGN.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SGN *SGNFilterer) ParsePaused(log types.Log) (*SGNPaused, error) {
	event := new(SGNPaused)
	if err := _SGN.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNSgnAddrUpdateIterator is returned from FilterSgnAddrUpdate and is used to iterate over the raw logs and unpacked data for SgnAddrUpdate events raised by the SGN contract.
type SGNSgnAddrUpdateIterator struct {
	Event *SGNSgnAddrUpdate // Event containing the contract specifics and raw log

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
func (it *SGNSgnAddrUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNSgnAddrUpdate)
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
		it.Event = new(SGNSgnAddrUpdate)
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
func (it *SGNSgnAddrUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNSgnAddrUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNSgnAddrUpdate represents a SgnAddrUpdate event raised by the SGN contract.
type SGNSgnAddrUpdate struct {
	ValAddr common.Address
	OldAddr []byte
	NewAddr []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSgnAddrUpdate is a free log retrieval operation binding the contract event 0x8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4.
//
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes oldAddr, bytes newAddr)
func (_SGN *SGNFilterer) FilterSgnAddrUpdate(opts *bind.FilterOpts, valAddr []common.Address) (*SGNSgnAddrUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _SGN.contract.FilterLogs(opts, "SgnAddrUpdate", valAddrRule)
	if err != nil {
		return nil, err
	}
	return &SGNSgnAddrUpdateIterator{contract: _SGN.contract, event: "SgnAddrUpdate", logs: logs, sub: sub}, nil
}

// WatchSgnAddrUpdate is a free log subscription operation binding the contract event 0x8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4.
//
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes oldAddr, bytes newAddr)
func (_SGN *SGNFilterer) WatchSgnAddrUpdate(opts *bind.WatchOpts, sink chan<- *SGNSgnAddrUpdate, valAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _SGN.contract.WatchLogs(opts, "SgnAddrUpdate", valAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNSgnAddrUpdate)
				if err := _SGN.contract.UnpackLog(event, "SgnAddrUpdate", log); err != nil {
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

// ParseSgnAddrUpdate is a log parse operation binding the contract event 0x8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4.
//
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes oldAddr, bytes newAddr)
func (_SGN *SGNFilterer) ParseSgnAddrUpdate(log types.Log) (*SGNSgnAddrUpdate, error) {
	event := new(SGNSgnAddrUpdate)
	if err := _SGN.contract.UnpackLog(event, "SgnAddrUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SGN contract.
type SGNUnpausedIterator struct {
	Event *SGNUnpaused // Event containing the contract specifics and raw log

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
func (it *SGNUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNUnpaused)
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
		it.Event = new(SGNUnpaused)
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
func (it *SGNUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNUnpaused represents a Unpaused event raised by the SGN contract.
type SGNUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SGN *SGNFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SGNUnpausedIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SGNUnpausedIterator{contract: _SGN.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SGN *SGNFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SGNUnpaused) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNUnpaused)
				if err := _SGN.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SGN *SGNFilterer) ParseUnpaused(log types.Log) (*SGNUnpaused, error) {
	event := new(SGNUnpaused)
	if err := _SGN.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SGN contract.
type SGNWithdrawIterator struct {
	Event *SGNWithdraw // Event containing the contract specifics and raw log

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
func (it *SGNWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNWithdraw)
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
		it.Event = new(SGNWithdraw)
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
func (it *SGNWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNWithdraw represents a Withdraw event raised by the SGN contract.
type SGNWithdraw struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address account, address token, uint256 amount)
func (_SGN *SGNFilterer) FilterWithdraw(opts *bind.FilterOpts) (*SGNWithdrawIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &SGNWithdrawIterator{contract: _SGN.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address account, address token, uint256 amount)
func (_SGN *SGNFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SGNWithdraw) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNWithdraw)
				if err := _SGN.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address account, address token, uint256 amount)
func (_SGN *SGNFilterer) ParseWithdraw(log types.Log) (*SGNWithdraw, error) {
	event := new(SGNWithdraw)
	if err := _SGN.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_governProposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_governVoteTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorBondInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ConfirmParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CreateParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"tokenDiff\",\"type\":\"int256\"}],\"name\":\"DelegationUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPool\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPoolSize\",\"type\":\"uint256\"}],\"name\":\"RewardPoolContribution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmt\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashAmtCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"ValidatorNotice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGovern.VoteOption\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"VoteParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bondedValAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"celerToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rewardRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"completeUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"confirmParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"confirmUnbondedValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"createParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structStaking.DelegatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getParamProposalVote\",\"outputs\":[{\"internalType\":\"enumGovern.VoteOption\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"}],\"name\":\"getParamValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"enumStaking.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_checkSelfDelegation\",\"type\":\"bool\"}],\"name\":\"hasMinRequiredTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_commissionRate\",\"type\":\"uint64\"}],\"name\":\"initializeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBondedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBondBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextParamProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paramProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"name\":\"setMaxSlashFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_whitelistEnabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerVals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_slashRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newRate\",\"type\":\"uint64\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"}],\"name\":\"updateMinSelfDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"updateValidatorSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"valAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"validatorNotice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"enumStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"bondBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"unbondBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.VoteOption\",\"name\":\"_vote\",\"type\":\"uint8\"}],\"name\":\"voteParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162005d4e38038062005d4e833981016040819052620000349162000227565b898989898989898989896200004933620001d7565b6000805460ff60a01b19168155600280546001600160a01b039c909c1661010002610100600160a81b0319909c169b909b17909a5560036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff989098557fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c969096557fc3a24b0501bd2c13a7e57f2db4369ec4c223447539fc0724a9d55ac4a06ebd4d949094557fcbc4e5fb02c3d1de23a9f1e014b4d2ee5aeaea9505df5e855c9210bf472495af929092557f83ec6a1f0257b830b5e016457c9cf1435391bf56cc98f369a58a54fe93772465557f405aad32e1adbac89bb7f176e338b8fc6e994ca210c9bb7bdca249b465942250557fc69056f16cbaa3c616b828e333ab7d3a32310765507f8f58359e99ebb7a885f3557ff2c49132ed1cee2a7e75bde50d332a2f81f1d01e5456d8a19d1df09bd561dbd25560089091527f85aaa47b6dc46495bb8824fad4583769726fea36efd831a35556690b830a8fbe5550620002b498505050505050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806000806000806000806000806101408b8d0312156200024857600080fd5b8a516001600160a01b03811681146200026057600080fd5b809a505060208b0151985060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935060e08b015192506101008b015191506101208b015190509295989b9194979a5092959850565b615a8a80620002c46000396000f3fe6080604052600436106103545760003560e01c806366666aa9116101c657806392bb243c116100f7578063c8f9f98411610095578063eecefef81161006f578063eecefef814610a9a578063f2fde38b14610ac7578063f8df0dc514610ae7578063fa52c7d814610b0757600080fd5b8063c8f9f98414610a14578063e478ed9d14610a4d578063eb505dd514610a6d57600080fd5b8063a310624f116100d1578063a310624f14610976578063acc62ccf146109af578063b4f7fa34146109cf578063c6c21e9d146109ef57600080fd5b806392bb243c14610906578063934a18ec146109265780639b19251a1461094657600080fd5b80638456cb59116101645780638da5cb5b1161013e5780638da5cb5b146108835780638dc2336d146108a157806390e360f8146108b65780639146f110146108e657600080fd5b80638456cb591461083957806389f9aab51461084e5780638a74d5fe1461086357600080fd5b806371bc0216116101a057806371bc0216146107705780637a50dbd2146107905780637e5fb8f3146107b057806383cfb3181461082357600080fd5b806366666aa9146106f75780636d3087831461070d578063715018a61461075b57600080fd5b80633985c4e6116102a05780634d99dd161161023e578063581c53c511610218578063581c53c51461064b5780635c975abb146106a25780635e593eff146106c157806365d5d420146106e157600080fd5b80634d99dd16146105f157806351fb012d14610611578063525eba211461062b57600080fd5b8063410ae02c1161027a578063410ae02c14610571578063473849bd1461059157806347abfdbf146105b157806349955e39146105d157600080fd5b80633985c4e6146104f35780633af32abf146105135780633f4ba83a1461055c57600080fd5b80631cfe4f0b1161030d57806325ed6b35116102e757806325ed6b3514610489578063291d9549146104a957806336f1635f146104c9578063386c024a146104de57600080fd5b80631cfe4f0b146104225780631e6f3d8a1461044657806322da79271461047357600080fd5b8063026e402b14610360578063052d9e7e146103825780630a300b09146103a257806310154bad146103c2578063145aa116146103e25780631a2032571461040257600080fd5b3661035b57005b600080fd5b34801561036c57600080fd5b5061038061037b3660046151a9565b610b9f565b005b34801561038e57600080fd5b5061038061039d36600461520f565b610db9565b3480156103ae57600080fd5b506103806103bd366004615390565b610df7565b3480156103ce57600080fd5b506103806103dd3660046150a4565b610eac565b3480156103ee57600080fd5b506103806103fd366004615390565b610edf565b34801561040e57600080fd5b5061038061041d366004615390565b610f75565b34801561042e57600080fd5b506009545b6040519081526020015b60405180910390f35b34801561045257600080fd5b506104336104613660046150a4565b600d6020526000908152604090205481565b34801561047f57600080fd5b5061043360055481565b34801561049557600080fd5b506103806104a43660046153cc565b610fcd565b3480156104b557600080fd5b506103806104c43660046150a4565b611050565b3480156104d557600080fd5b50610380611083565b3480156104ea57600080fd5b506104336113e3565b3480156104ff57600080fd5b5061038061050e366004615249565b611410565b34801561051f57600080fd5b5061054c61052e3660046150a4565b6001600160a01b031660009081526001602052604090205460ff1690565b604051901515815260200161043d565b34801561056857600080fd5b50610380611aa0565b34801561057d57600080fd5b5061043361058c366004615359565b611ad4565b34801561059d57600080fd5b506103806105ac3660046150a4565b611b13565b3480156105bd57600080fd5b5061054c6105cc3660046150f2565b611dbe565b3480156105dd57600080fd5b506103806105ec3660046153f5565b611e78565b3480156105fd57600080fd5b5061038061060c3660046151a9565b611f89565b34801561061d57600080fd5b5060025461054c9060ff1681565b34801561063757600080fd5b506103806106463660046151d3565b612329565b34801561065757600080fd5b506106956106663660046153a9565b60008281526004602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b60405161043d91906154e7565b3480156106ae57600080fd5b50600054600160a01b900460ff1661054c565b3480156106cd57600080fd5b506103806106dc366004615390565b61276e565b3480156106ed57600080fd5b5061043360075481565b34801561070357600080fd5b5061043360065481565b34801561071957600080fd5b506107436107283660046150a4565b600c602052600090815260409020546001600160a01b031681565b6040516001600160a01b03909116815260200161043d565b34801561076757600080fd5b5061038061293a565b34801561077c57600080fd5b5061038061078b3660046150a4565b61296e565b34801561079c57600080fd5b506103806107ab3660046150a4565b612aab565b3480156107bc57600080fd5b506108116107cb366004615390565b60046020819052600091825260409091208054600182015460028301546003840154948401546005909401546001600160a01b03909316949193909260ff928316921686565b60405161043d96959493929190615495565b34801561082f57600080fd5b5061043360085481565b34801561084557600080fd5b50610380612ca6565b34801561085a57600080fd5b50600a54610433565b34801561086f57600080fd5b5061054c61087e3660046152e2565b612cd8565b34801561088f57600080fd5b506000546001600160a01b0316610743565b3480156108ad57600080fd5b50610433612f7e565b3480156108c257600080fd5b5061054c6108d1366004615390565b600e6020526000908152604090205460ff1681565b3480156108f257600080fd5b50610380610901366004615129565b61307d565b34801561091257600080fd5b50610743610921366004615390565b613109565b34801561093257600080fd5b50610380610941366004615390565b613133565b34801561095257600080fd5b5061054c6109613660046150a4565b60016020526000908152604090205460ff1681565b34801561098257600080fd5b506106956109913660046150a4565b6001600160a01b03166000908152600b602052604090205460ff1690565b3480156109bb57600080fd5b506107436109ca366004615390565b613261565b3480156109db57600080fd5b5061054c6109ea3660046150a4565b613271565b3480156109fb57600080fd5b506002546107439061010090046001600160a01b031681565b348015610a2057600080fd5b50610433610a2f3660046150a4565b6001600160a01b03166000908152600b602052604090206001015490565b348015610a5957600080fd5b50610380610a68366004615374565b6132a9565b348015610a7957600080fd5b50610433610a88366004615359565b60036020526000908152604090205481565b348015610aa657600080fd5b50610aba610ab53660046150bf565b6133f5565b60405161043d9190615736565b348015610ad357600080fd5b50610380610ae23660046150a4565b613601565b348015610af357600080fd5b50610380610b02366004615249565b613699565b348015610b1357600080fd5b50610b89610b223660046150a4565b600b60205260009081526040902080546001820154600283015460038401546004850154600686015460079096015460ff8616966101009096046001600160a01b031695906001600160401b0380821691600160401b8104821691600160801b909104168a565b60405161043d9a999897969594939291906154fa565b600054600160a01b900460ff1615610bd25760405162461bcd60e51b8152600401610bc99061560e565b60405180910390fd5b33670de0b6b3a7640000821015610c2b5760405162461bcd60e51b815260206004820152601860248201527f4d696e696d616c20616d6f756e7420697320312043454c5200000000000000006044820152606401610bc9565b6001600160a01b0383166000908152600b6020526040812090815460ff166003811115610c5a57610c5a6159be565b1415610c785760405162461bcd60e51b8152600401610bc9906156cb565b6000610c8d84836001015484600201546138ae565b6001600160a01b0384166000908152600584016020526040812080549293509183918391610cbc90849061587e565b9250508190555081836002016000828254610cd7919061587e565b9250508190555084836001016000828254610cf2919061587e565b9091555060039050835460ff166003811115610d1057610d106159be565b1415610d3b578460076000828254610d28919061587e565b90915550506001830154610d3b906138db565b600254610d589061010090046001600160a01b03168530886139dd565b6001830154815460408051928352602083019190915281018690526001600160a01b0380861691908816907f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea906060015b60405180910390a3505050505050565b6000546001600160a01b03163314610de35760405162461bcd60e51b8152600401610bc990615696565b6002805460ff191682151517905550565b50565b600054600160a01b900460ff1615610e215760405162461bcd60e51b8152600401610bc99061560e565b60003390508160066000828254610e38919061587e565b9091555050600254610e5a9061010090046001600160a01b03168230856139dd565b806001600160a01b03167f574fb579a16e67a50f3626feb14c7c23cb5d3c7ff880e0d3e2f792fbe8d22fca83600654604051610ea0929190918252602082015260400190565b60405180910390a25050565b6000546001600160a01b03163314610ed65760405162461bcd60e51b8152600401610bc990615696565b610df481613a4e565b600054600160a01b900460ff16610f2f5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610bc9565b6000546001600160a01b03163314610f595760405162461bcd60e51b8152600401610bc990615696565b600254610df49061010090046001600160a01b03163383613b09565b6000546001600160a01b03163314610f9f5760405162461bcd60e51b8152600401610bc990615696565b600860005260036020527f85aaa47b6dc46495bb8824fad4583769726fea36efd831a35556690b830a8fbe55565b336000818152600b602052604090205460039060ff1681811115610ff357610ff36159be565b146110405760405162461bcd60e51b815260206004820152601f60248201527f566f746572206973206e6f74206120626f6e6465642076616c696461746f72006044820152606401610bc9565b61104b838284613b39565b505050565b6000546001600160a01b0316331461107a5760405162461bcd60e51b8152600401610bc990615696565b610df481613cde565b336000818152600c60205260409020546001600160a01b0316156110bc5750336000908152600c60205260409020546001600160a01b03165b6001600160a01b0381166000908152600b602052604090206001815460ff1660038111156110ec576110ec6159be565b148061110d57506002815460ff16600381111561110b5761110b6159be565b145b6111595760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076616c696461746f722073746174757300000000000000006044820152606401610bc9565b60078101546001600160401b03164310156111af5760405162461bcd60e51b8152602060048201526016602482015275109bdb9908189b1bd8dac81b9bdd081c995858da195960521b6044820152606401610bc9565b6008544310156112015760405162461bcd60e51b815260206004820152601b60248201527f546f6f206672657175656e742076616c696461746f7220626f6e6400000000006044820152606401610bc9565b600760005260036020527ff2c49132ed1cee2a7e75bde50d332a2f81f1d01e5456d8a19d1df09bd561dbd254611237904361587e565b600855611245826001611dbe565b6112875760405162461bcd60e51b81526020600482015260136024820152724e6f742068617665206d696e20746f6b656e7360681b6044820152606401610bc9565b600360008190526020527fcbc4e5fb02c3d1de23a9f1e014b4d2ee5aeaea9505df5e855c9210bf472495af54600a548111156112c65761104b83613d89565b6000196000805b8381101561137b5782600b6000600a84815481106112ed576112ed6159ea565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154101561136957809150600b6000600a8381548110611335576113356159ea565b60009182526020808320909101546001600160a01b031683528201929092526040019020600101549250826113695761137b565b806113738161594c565b9150506112cd565b50818460010154116113c55760405162461bcd60e51b8152602060048201526013602482015272496e73756666696369656e7420746f6b656e7360681b6044820152606401610bc9565b6113cf8582613ddd565b6113dc84600101546138db565b5050505050565b6000600360075460026113f691906158b8565b6114009190615896565b61140b90600161587e565b905090565b600054600160a01b900460ff161561143a5760405162461bcd60e51b8152600401610bc99061560e565b600061147b85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613e5892505050565b90506114c385858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061087e9250869150879050615913565b5080606001516001600160401b031643106115105760405162461bcd60e51b815260206004820152600d60248201526c14db185cda08195e1c1a5c9959609a1b6044820152606401610bc9565b620f424081604001516001600160401b031611156115675760405162461bcd60e51b815260206004820152601460248201527324b73b30b634b21039b630b9b4103330b1ba37b960611b6044820152606401610bc9565b600860005260036020527f85aaa47b6dc46495bb8824fad4583769726fea36efd831a35556690b830a8fbe5460408201516001600160401b031611156115ef5760405162461bcd60e51b815260206004820152601760248201527f457863656564206d617820736c61736820666163746f720000000000000000006044820152606401610bc9565b6020808201516001600160401b03166000908152600e909152604090205460ff16156116505760405162461bcd60e51b815260206004820152601060248201526f5573656420736c617368206e6f6e636560801b6044820152606401610bc9565b6020808201516001600160401b03166000908152600e82526040808220805460ff1916600190811790915584516001600160a01b0381168452600b909452912090815460ff1660038111156116a7576116a76159be565b14156116eb5760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881d5b989bdd5b991959606a1b6044820152606401610bc9565b6000620f424084604001516001600160401b0316836001015461170e91906158b8565b6117189190615896565b90508082600101600082825461172e91906158d7565b9091555060039050825460ff16600381111561174c5761174c6159be565b14156117e957806007600082825461176491906158d7565b909155505060808401516001600160401b031615158061178c575061178a836001611dbe565b155b156117e95761179a836140c5565b60808401516001600160401b0316156117e95760808401516117c5906001600160401b03164361587e565b60078301805467ffffffffffffffff19166001600160401b03929092169190911790555b60006001600160a01b0316836001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea84600101546000856118309061598b565b6040805193845260208401929092529082015260600160405180910390a36000620f424085604001516001600160401b0316846003015461187191906158b8565b61187b9190615896565b90508083600301600082825461189191906158d7565b909155506118a19050818361587e565b91506000805b8660a00151518110156119db5760008760a0015182815181106118cc576118cc6159ea565b602002602001015190508060200151836118e6919061587e565b81519093506001600160a01b031661195957602081015160025461191b916101009091046001600160a01b0316903390613b09565b60208082015160405190815233917fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3910160405180910390a26119c8565b8051602082015160025461197c926101009091046001600160a01b031691613b09565b80600001516001600160a01b03167fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a382602001516040516119bf91815260200190565b60405180910390a25b50806119d38161594c565b9150506118a7565b5080831015611a215760405162461bcd60e51b8152602060048201526012602482015271496e76616c696420636f6c6c6563746f727360701b6044820152606401610bc9565b611a2b81846158d7565b60066000828254611a3c919061587e565b9091555050602080870151604080516001600160401b0390921682529181018590526001600160a01b038716917f10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008910160405180910390a250505050505050505050565b6000546001600160a01b03163314611aca5760405162461bcd60e51b8152600401610bc990615696565b611ad2614222565b565b600060036000836008811115611aec57611aec6159be565b6008811115611afd57611afd6159be565b8152602001908152602001600020549050919050565b6001600160a01b0381166000908152600b602052604081203391815460ff166003811115611b4357611b436159be565b1415611b615760405162461bcd60e51b8152600401610bc9906156cb565b6001600160a01b038216600090815260058201602090815260408220600283526003918290527fc3a24b0501bd2c13a7e57f2db4369ec4c223447539fc0724a9d55ac4a06ebd4d54845491939092909160019160ff90911690811115611bc957611bc96159be565b60028501549114915063ffffffff1660005b600285015463ffffffff64010000000090910481169083161015611c8d578280611c2b575063ffffffff821660009081526001808701602052604090912001544390611c2890869061587e565b11155b15611c765763ffffffff82166000908152600186016020526040902054611c52908261587e565b63ffffffff8316600090815260018088016020526040822082815501559050611c7b565b611c8d565b81611c8581615967565b925050611bdb565b60028501805463ffffffff191663ffffffff841617905580611cff5760405162461bcd60e51b815260206004820152602560248201527f4e6f20756e64656c65676174696f6e20726561647920746f20626520636f6d706044820152641b195d195960da1b6064820152608401610bc9565b6000611d1482886003015489600401546142bf565b905081876004016000828254611d2a91906158d7565b9250508190555080876003016000828254611d4591906158d7565b9091555050600254611d669061010090046001600160a01b03168983613b09565b876001600160a01b0316896001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c83604051611dab91815260200190565b60405180910390a3505050505050505050565b6001600160a01b0382166000908152600b60209081526040822060018101546004845260039092527f83ec6a1f0257b830b5e016457c9cf1435391bf56cc98f369a58a54fe9377246554909190811015611e1d57600092505050611e72565b8315611e6b576001600160a01b03851660009081526005830160205260408120546002840154611e4f919084906142bf565b90508260060154811015611e695760009350505050611e72565b505b6001925050505b92915050565b336000818152600b6020526040812090815460ff166003811115611e9e57611e9e6159be565b1415611ebc5760405162461bcd60e51b8152600401610bc9906156cb565b612710836001600160401b03161115611f0a5760405162461bcd60e51b815260206004820152601060248201526f496e76616c6964206e6577207261746560801b6044820152606401610bc9565b60078101805467ffffffffffffffff60801b1916600160801b6001600160401b038616908102919091179091556040805160208101929092526001600160a01b03841691600080516020615a35833981519152910160408051601f1981840301815290829052611f7c91600090615702565b60405180910390a2505050565b33670de0b6b3a7640000821015611fe25760405162461bcd60e51b815260206004820152601960248201527f4d696e696d616c20616d6f756e742069732031207368617265000000000000006044820152606401610bc9565b6001600160a01b0383166000908152600b6020526040812090815460ff166003811115612011576120116159be565b141561202f5760405162461bcd60e51b8152600401610bc9906156cb565b600061204484836001015484600201546142bf565b6001600160a01b03841660009081526005840160205260408120805492935091869183916120739084906158d7565b925050819055508483600201600082825461208e91906158d7565b92505081905550818360010160008282546120a991906158d7565b9091555060019050835460ff1660038111156120c7576120c76159be565b141561212e576002546120e99061010090046001600160a01b03168584613b09565b836001600160a01b0316866001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c84604051610da991815260200190565b6003835460ff166003811115612146576121466159be565b141561219057816007600082825461215e91906158d7565b9250508190555061218386876001600160a01b0316866001600160a01b031614611dbe565b61219057612190866140c5565b6002810154600a906121b39063ffffffff808216916401000000009004166158ee565b63ffffffff16106122065760405162461bcd60e51b815260206004820152601f60248201527f457863656564206d617820756e64656c65676174696f6e20656e7472696573006044820152606401610bc9565b600061221b83856003015486600401546138ae565b905080846004016000828254612231919061587e565b925050819055508284600301600082825461224c919061587e565b909155505060028201805463ffffffff6401000000009182900481166000908152600180870160205260409091208581554391810191909155835490939290041690600461229983615967565b91906101000a81548163ffffffff021916908363ffffffff16021790555050856001600160a01b0316886001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea87600101548660000154886123019061598b565b6040805193845260208401929092529082015260600160405180910390a35050505050505050565b600054600160a01b900460ff16156123535760405162461bcd60e51b8152600401610bc99061560e565b60025460ff16156123bd573360009081526001602052604090205460ff166123bd5760405162461bcd60e51b815260206004820152601960248201527f43616c6c6572206973206e6f742077686974656c6973746564000000000000006044820152606401610bc9565b336000818152600b6020526040812090815460ff1660038111156123e3576123e36159be565b146124305760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f7220697320696e697469616c697a656400000000000000006044820152606401610bc9565b6001600160a01b0385166000908152600b602052604081205460ff16600381111561245d5761245d6159be565b146124a65760405162461bcd60e51b815260206004820152601960248201527829b4b3b732b91034b99037ba3432b9103b30b634b230ba37b960391b6044820152606401610bc9565b6001600160a01b038281166000908152600c6020526040902054161561250e5760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206973206f74686572207369676e6572000000000000006044820152606401610bc9565b6001600160a01b038581166000908152600c6020526040902054161561256c5760405162461bcd60e51b815260206004820152601360248201527214da59db995c88185b1c9958591e481d5cd959606a1b6044820152606401610bc9565b612710836001600160401b031611156125c75760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636f6d6d697373696f6e20726174650000000000000000006044820152606401610bc9565b600560005260036020527f405aad32e1adbac89bb7f176e338b8fc6e994ca210c9bb7bdca249b465942250548410156126425760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610bc9565b80546001600160a81b03191660ff196101006001600160a01b038881169182029290921692909217600190811784556006840187905560078401805467ffffffffffffffff60801b1916600160801b6001600160401b038916021790556009805491820190557f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b03199081169286169283179091556000928352600c602052604090922080549092161790556127028285610b9f565b604080516001600160a01b0387811660208301529181018690526001600160401b038516606082015290831690600080516020615a358339815191529060800160408051601f198184030181529082905261275f91600090615668565b60405180910390a25050505050565b336000818152600b6020526040812090815460ff166003811115612794576127946159be565b14156127b25760405162461bcd60e51b8152600401610bc9906156cb565b600560005260036020527f405aad32e1adbac89bb7f176e338b8fc6e994ca210c9bb7bdca249b4659422505483101561282d5760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610bc9565b80600601548310156128ef576003815460ff166003811115612851576128516159be565b14156128955760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881a5cc8189bdb991959606a1b6044820152606401610bc9565b600660005260036020527fc69056f16cbaa3c616b828e333ab7d3a32310765507f8f58359e99ebb7a885f3546128cb904361587e565b60078201805467ffffffffffffffff19166001600160401b03929092169190911790555b6006810183905560408051602081018590526001600160a01b03841691600080516020615a35833981519152910160408051601f1981840301815290829052611f7c916000906155ba565b6000546001600160a01b031633146129645760405162461bcd60e51b8152600401610bc990615696565b611ad260006142d8565b6001600160a01b0381166000908152600b602052604090206002815460ff16600381111561299e5761299e6159be565b146129eb5760405162461bcd60e51b815260206004820152601760248201527f56616c696461746f72206e6f7420756e626f6e64696e670000000000000000006044820152606401610bc9565b6007810154600160401b90046001600160401b0316431015612a4f5760405162461bcd60e51b815260206004820152601860248201527f556e626f6e6420626c6f636b206e6f74207265616368656400000000000000006044820152606401610bc9565b805460ff19166001908117825560078201805467ffffffffffffffff60401b191690555b6040516001600160a01b038416907fd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a90600090a35050565b336000818152600b6020526040812090815460ff166003811115612ad157612ad16159be565b1415612b1f5760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206e6f7420696e697469616c697a6564000000000000006044820152606401610bc9565b6001600160a01b038381166000908152600c60205260409020541615612b7d5760405162461bcd60e51b815260206004820152601360248201527214da59db995c88185b1c9958591e481d5cd959606a1b6044820152606401610bc9565b816001600160a01b0316836001600160a01b031614612c0c576001600160a01b0383166000908152600b602052604081205460ff166003811115612bc357612bc36159be565b14612c0c5760405162461bcd60e51b815260206004820152601960248201527829b4b3b732b91034b99037ba3432b9103b30b634b230ba37b960391b6044820152606401610bc9565b8054610100908190046001600160a01b039081166000908152600c6020908152604080832080546001600160a01b03199081169091558654610100600160a81b0319168986169687021787558584529281902080549093169387169384179092558151908101939093529091600080516020615a35833981519152910160408051601f1981840301815290829052611f7c91600090615638565b6000546001600160a01b03163314612cd05760405162461bcd60e51b8152600401610bc990615696565b611ad2614328565b600080612d3984805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050600083516001600160401b03811115612d5657612d56615a00565b604051908082528060200260200182016040528015612d7f578160200160208202803683370190505b509050600080805b8651811015612f2457612dbc878281518110612da557612da56159ea565b60200260200101518661438d90919063ffffffff16565b848281518110612dce57612dce6159ea565b60200260200101906001600160a01b031690816001600160a01b031681525050816001600160a01b0316848281518110612e0a57612e0a6159ea565b60200260200101516001600160a01b031611612e685760405162461bcd60e51b815260206004820152601e60248201527f5369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610bc9565b838181518110612e7a57612e7a6159ea565b602002602001015191506000600b6000600c6000888681518110612ea057612ea06159ea565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081205490931684528301939093529101902090506003815460ff166003811115612ef457612ef46159be565b14612eff5750612f12565b6001810154612f0e908561587e565b9350505b80612f1c8161594c565b915050612d87565b50612f2d6113e3565b821015612f715760405162461bcd60e51b8152602060048201526012602482015271145d5bdc9d5b481b9bdd081c995858da195960721b6044820152606401610bc9565b5060019695505050505050565b600080600b6000600a600081548110612f9957612f996159ea565b60009182526020808320909101546001600160a01b03168352820192909252604001902060019081015491505b600a548110156130775781600b6000600a8481548110612fe857612fe86159ea565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154101561306557600b6000600a838154811061302d5761302d6159ea565b60009182526020808320909101546001600160a01b031683528201929092526040019020600101549150816130655760009250505090565b8061306f8161594c565b915050612fc6565b50919050565b6001600160a01b0385166000908152600b6020526040812090815460ff1660038111156130ac576130ac6159be565b14156130ca5760405162461bcd60e51b8152600401610bc9906156cb565b856001600160a01b0316600080516020615a3583398151915286868686336040516130f9959493929190615565565b60405180910390a2505050505050565b6009818154811061311957600080fd5b6000918252602090912001546001600160a01b0316905081565b6000805b600a5463ffffffff8216101561321757600161319c84600a8463ffffffff1681548110613166576131666159ea565b60009182526020808320909101549282526004815260408083206001600160a01b03909416835260069093019052205460ff1690565b60038111156131ad576131ad6159be565b141561320557600b6000600a8363ffffffff16815481106131d0576131d06159ea565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154613202908361587e565b91505b8061320f81615967565b915050613137565b5060006132226113e3565b82101590508061325757600083815260046020526040812060010154600680549192909161325190849061587e565b90915550505b61104b8382614431565b600a818154811061311957600080fd5b600060036001600160a01b0383166000908152600b602052604090205460ff1660038111156132a2576132a26159be565b1492915050565b6005546000818152600460205260409020906132c690600161587e565b60055560036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff5481546001600160a01b03191633908117835560018084018390556000527fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c5490919061333c904361587e565b600284015560038301805486919060ff19166001836008811115613362576133626159be565b02179055506004830184905560058301805460ff1916600117905560025461339b906001600160a01b03610100909104168330846139dd565b7f4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f16060016005546133cb91906158d7565b8383866002015489896040516133e6969594939291906157f2565b60405180910390a15050505050565b6134306040518060a0016040528060006001600160a01b03168152602001600081526020016000815260200160008152602001606081525090565b6001600160a01b038084166000908152600b6020908152604080832093861683526005840190915281208054600184015460028501549293926134749291906142bf565b6002830154909150600090819061349c9063ffffffff808216916401000000009004166158ee565b63ffffffff1690506000816001600160401b038111156134be576134be615a00565b60405190808252806020026020018201604052801561350357816020015b60408051808201909152600080825260208201528152602001906001900390816134dc5790505b50905060005b828110156135ab576002860154600187019060009061352e9063ffffffff168461587e565b81526020019081526020016000206040518060400160405290816000820154815260200160018201548152505082828151811061356d5761356d6159ea565b6020908102919091018101919091526000828152600188019091526040902054613597908561587e565b9350806135a38161594c565b915050613509565b5060006135c184886003015489600401546142bf565b6040805160a0810182526001600160a01b03909c168c5260208c01969096529554948a019490945260608901949094525050506080850152509192915050565b6000546001600160a01b0316331461362b5760405162461bcd60e51b8152600401610bc990615696565b6001600160a01b0381166136905760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610bc9565b610df4816142d8565b600054600160a01b900460ff16156136c35760405162461bcd60e51b8152600401610bc99061560e565b61370984848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061087e9250859150869050615913565b50600061374b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506145c492505050565b80516001600160a01b03166000908152600d6020908152604082205490830151929350909161377a91906158d7565b9050600081116137bc5760405162461bcd60e51b815260206004820152600d60248201526c139bc81b995dc81c995dd85c99609a1b6044820152606401610bc9565b80600654101561380e5760405162461bcd60e51b815260206004820152601860248201527f496e73756666696369656e742072657761726420706f6f6c00000000000000006044820152606401610bc9565b60208083015183516001600160a01b03166000908152600d9092526040822055600680548392906138409084906158d7565b90915550508151600254613864916101009091046001600160a01b03169083613b09565b81600001516001600160a01b03167ff01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743826006546040516130f9929190918252602082015260400190565b6000826138bc5750826138d4565b826138c783866158b8565b6138d19190615896565b90505b9392505050565b600a5460028114806138ed5750806003145b15613963576138fa6113e3565b821061395f5760405162461bcd60e51b815260206004820152602e60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201526d2071756f72756d20746f6b656e7360901b6064820152608401610bc9565b5050565b600381111561395f57600360075461397b9190615896565b821061395f5760405162461bcd60e51b815260206004820152602b60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201526a20312f3320746f6b656e7360a81b6064820152608401610bc9565b6040516001600160a01b0380851660248301528316604482015260648101829052613a489085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261466b565b50505050565b6001600160a01b03811660009081526001602052604090205460ff1615613aad5760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481dda1a5d195b1a5cdd1959606a1b6044820152606401610bc9565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f91015b60405180910390a150565b6040516001600160a01b03831660248201526044810182905261104b90849063a9059cbb60e01b90606401613a11565b60008381526004602052604090206001600582015460ff166002811115613b6257613b626159be565b14613ba95760405162461bcd60e51b8152602060048201526017602482015276496e76616c69642070726f706f73616c2073746174757360481b6044820152606401610bc9565b80600201544310613bf35760405162461bcd60e51b8152602060048201526014602482015273159bdd1948191958591b1a5b99481c185cdcd95960621b6044820152606401610bc9565b6001600160a01b038316600090815260068201602052604081205460ff166003811115613c2257613c226159be565b14613c615760405162461bcd60e51b815260206004820152600f60248201526e159bdd195c881a185cc81d9bdd1959608a1b6044820152606401610bc9565b6001600160a01b03831660009081526006820160205260409020805483919060ff19166001836003811115613c9857613c986159be565b02179055507f06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65848484604051613cd0939291906157c5565b60405180910390a150505050565b6001600160a01b03811660009081526001602052604090205460ff16613d385760405162461bcd60e51b815260206004820152600f60248201526e139bdd081dda1a5d195b1a5cdd1959608a1b6044820152606401610bc9565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b69101613afe565b600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180546001600160a01b0319166001600160a01b038316179055610df48161473d565b613e0d600a8281548110613df357613df36159ea565b6000918252602090912001546001600160a01b031661479a565b81600a8281548110613e2157613e216159ea565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061395f8261473d565b6040805160c0810182526000808252602080830182905282840182905260608084018390526080840183905260a084015283518085019094528184528301849052909190613ea782600661483f565b905080600681518110613ebc57613ebc6159ea565b60200260200101516001600160401b03811115613edb57613edb615a00565b604051908082528060200260200182016040528015613f2057816020015b6040805180820190915260008082526020820152815260200190600190039081613ef95790505b508360a00181905250600081600681518110613f3e57613f3e6159ea565b6020026020010181815250506000805b602084015151845110156140bc57613f65846148f8565b90925090508160011415613f9457613f84613f7f85614932565b6149ee565b6001600160a01b03168552613f4e565b8160021415613fb957613fa6846149f9565b6001600160401b03166020860152613f4e565b8160031415613fde57613fcb846149f9565b6001600160401b03166040860152613f4e565b816004141561400357613ff0846149f9565b6001600160401b03166060860152613f4e565b816005141561402857614015846149f9565b6001600160401b03166080860152613f4e565b81600614156140ad5761404261403d85614932565b614a7b565b8560a001518460068151811061405a5761405a6159ea565b602002602001015181518110614072576140726159ea565b602002602001018190525082600681518110614090576140906159ea565b6020026020010180518091906140a59061594c565b905250613f4e565b6140b78482614b15565b613f4e565b50505050919050565b600a546000906140d7906001906158d7565b905060005b600a548110156141e257826001600160a01b0316600a8281548110614103576141036159ea565b6000918252602090912001546001600160a01b031614156141d0578181101561419457600a8281548110614139576141396159ea565b600091825260209091200154600a80546001600160a01b039092169183908110614165576141656159ea565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b600a8054806141a5576141a56159d4565b600082815260209020810160001990810180546001600160a01b031916905501905561104b8361479a565b806141da8161594c565b9150506140dc565b5060405162461bcd60e51b81526020600482015260146024820152732737ba103137b73232b2103b30b634b230ba37b960611b6044820152606401610bc9565b600054600160a01b900460ff166142725760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610bc9565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000816142cd5750826138d4565b816138c784866158b8565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff16156143525760405162461bcd60e51b8152600401610bc99061560e565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586142a23390565b60008151604114156143c15760208201516040830151606084015160001a6143b786828585614b87565b9350505050611e72565b8151604014156143e957602082015160408301516143e0858383614d30565b92505050611e72565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610bc9565b60008281526004602052604090206001600582015460ff16600281111561445a5761445a6159be565b146144a15760405162461bcd60e51b8152602060048201526017602482015276496e76616c69642070726f706f73616c2073746174757360481b6044820152606401610bc9565b80600201544310156144f55760405162461bcd60e51b815260206004820152601960248201527f566f746520646561646c696e65206e6f742072656163686564000000000000006044820152606401610bc9565b60058101805460ff1916600217905581156145765780546001820154600254614531926001600160a01b03610100909204821692911690613b09565b600481015460038083015460009060ff166008811115614553576145536159be565b6008811115614564576145646159be565b81526020810191909152604001600020555b600381015460048201546040517fd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9926145b7928792879260ff169190615823565b60405180910390a1505050565b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b6020830151518351101561466357614606836148f8565b9092509050816001141561463057614620613f7f84614932565b6001600160a01b031684526145ef565b81600214156146545761464a61464584614932565b614d5a565b60208501526145ef565b61465e8382614b15565b6145ef565b505050919050565b60006146c0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614d919092919063ffffffff16565b80519091501561104b57808060200190518101906146de919061522c565b61104b5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610bc9565b6001600160a01b0381166000908152600b60205260408120805460ff191660031781556007808201805467ffffffffffffffff60401b19169055600182015481549293909261478d90849061587e565b9091555060039050612a73565b6001600160a01b0381166000908152600b602090815260408220805460ff191660029081178255909252600390527fc3a24b0501bd2c13a7e57f2db4369ec4c223447539fc0724a9d55ac4a06ebd4d546147f4904361587e565b8160070160086101000a8154816001600160401b0302191690836001600160401b0316021790555080600101546007600082825461483291906158d7565b9091555060029050612a73565b815160609061484f83600161587e565b6001600160401b0381111561486657614866615a00565b60405190808252806020026020018201604052801561488f578160200160208202803683370190505b5091506000805b602086015151865110156148ef576148ad866148f8565b809250819350505060018483815181106148c9576148c96159ea565b602002602001018181516148dd919061587e565b9052506148ea8682614b15565b614896565b50509092525090565b6000806000614906846149f9565b9050614913600882615896565b925080600716600581111561492a5761492a6159be565b915050915091565b6060600061493f836149f9565b90506000818460000151614953919061587e565b905083602001515181111561496757600080fd5b816001600160401b0381111561497f5761497f615a00565b6040519080825280601f01601f1916602001820160405280156149a9576020820181803683370190505b50602080860151865192955091818601919083010160005b858110156149e35781810151838201526149dc60208261587e565b90506149c1565b505050935250919050565b6000611e7282614da0565b602080820151825181019091015160009182805b600a811015614a755783811a9150614a268160076158b8565b82607f16901b851794508160801660001415614a6357614a4781600161587e565b86518790614a5690839061587e565b9052509395945050505050565b80614a6d8161594c565b915050614a0d565b50600080fd5b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b6020830151518351101561466357614abd836148f8565b90925090508160011415614ae757614ad7613f7f84614932565b6001600160a01b03168452614aa6565b8160021415614b0657614afc61464584614932565b6020850152614aa6565b614b108382614b15565b614aa6565b6000816005811115614b2957614b296159be565b1415614b385761104b826149f9565b6002816005811115614b4c57614b4c6159be565b141561035b576000614b5d836149f9565b90508083600001818151614b71919061587e565b9052506020830151518351111561104b57600080fd5b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115614c045760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610bc9565b8360ff16601b1480614c1957508360ff16601c145b614c705760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610bc9565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015614cc4573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614d275760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610bc9565b95945050505050565b60006001600160ff1b03821660ff83901c601b01614d5086828785614b87565b9695505050505050565b6000602082511115614d6b57600080fd5b6020820151905081516020614d8091906158d7565b614d8b9060086158b8565b1c919050565b60606138d18484600085614dbf565b60008151601414614db057600080fd5b5060200151600160601b900490565b606082471015614e205760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610bc9565b843b614e6e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610bc9565b600080866001600160a01b03168587604051614e8a9190615479565b60006040518083038185875af1925050503d8060008114614ec7576040519150601f19603f3d011682016040523d82523d6000602084013e614ecc565b606091505b5091509150614edc828286614ee7565b979650505050505050565b60608315614ef65750816138d4565b825115614f065782518084602001fd5b8160405162461bcd60e51b8152600401610bc991906155a7565b60006001600160401b0380841115614f3a57614f3a615a00565b8360051b6020614f4b81830161584e565b868152935080840185838101891015614f6357600080fd5b60009350835b88811015614f9e57813586811115614f7f578586fd5b614f8b8b828b0161500f565b8452509183019190830190600101614f69565b5050505050509392505050565b80356001600160a01b0381168114614fc257600080fd5b919050565b60008083601f840112614fd957600080fd5b5081356001600160401b03811115614ff057600080fd5b60208301915083602082850101111561500857600080fd5b9250929050565b600082601f83011261502057600080fd5b81356001600160401b0381111561503957615039615a00565b61504c601f8201601f191660200161584e565b81815284602083860101111561506157600080fd5b816020850160208301376000918101602001919091529392505050565b803560098110614fc257600080fd5b80356001600160401b0381168114614fc257600080fd5b6000602082840312156150b657600080fd5b6138d482614fab565b600080604083850312156150d257600080fd5b6150db83614fab565b91506150e960208401614fab565b90509250929050565b6000806040838503121561510557600080fd5b61510e83614fab565b9150602083013561511e81615a26565b809150509250929050565b60008060008060006060868803121561514157600080fd5b61514a86614fab565b945060208601356001600160401b038082111561516657600080fd5b61517289838a01614fc7565b9096509450604088013591508082111561518b57600080fd5b5061519888828901614fc7565b969995985093965092949392505050565b600080604083850312156151bc57600080fd5b6151c583614fab565b946020939093013593505050565b6000806000606084860312156151e857600080fd5b6151f184614fab565b9250602084013591506152066040850161508d565b90509250925092565b60006020828403121561522157600080fd5b81356138d481615a26565b60006020828403121561523e57600080fd5b81516138d481615a26565b6000806000806040858703121561525f57600080fd5b84356001600160401b038082111561527657600080fd5b61528288838901614fc7565b9096509450602087013591508082111561529b57600080fd5b818701915087601f8301126152af57600080fd5b8135818111156152be57600080fd5b8860208260051b85010111156152d357600080fd5b95989497505060200194505050565b600080604083850312156152f557600080fd5b82356001600160401b038082111561530c57600080fd5b6153188683870161500f565b9350602085013591508082111561532e57600080fd5b508301601f8101851361534057600080fd5b61534f85823560208401614f20565b9150509250929050565b60006020828403121561536b57600080fd5b6138d48261507e565b6000806040838503121561538757600080fd5b6151c58361507e565b6000602082840312156153a257600080fd5b5035919050565b600080604083850312156153bc57600080fd5b823591506150e960208401614fab565b600080604083850312156153df57600080fd5b8235915060208301356004811061511e57600080fd5b60006020828403121561540757600080fd5b6138d48261508d565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008151808452615451816020860160208601615920565b601f01601f19169290920160200192915050565b60098110615475576154756159be565b9052565b6000825161548b818460208701615920565b9190910192915050565b6001600160a01b0387168152602081018690526040810185905260c081016154c06060830186615465565b836080830152600383106154d6576154d66159be565b8260a0830152979650505050505050565b602081016154f483615a16565b91905290565b61014081016155088c615a16565b9a81526001600160a01b039990991660208a015260408901979097526060880195909552608087019390935260a086019190915260c08501526001600160401b0390811660e0850152908116610100840152166101209091015290565b606081526000615579606083018789615410565b828103602084015261558c818688615410565b91505060018060a01b03831660408301529695505050505050565b6020815260006138d46020830184615439565b60608152601360608201527236b4b716b9b2b63316b232b632b3b0ba34b7b760691b608082015260a0602082015260006155f760a0830185615439565b905060018060a01b03831660408301529392505050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b60608152600660608201526539b4b3b732b960d11b608082015260a0602082015260006155f760a0830185615439565b6060815260046060820152631a5b9a5d60e21b608082015260a0602082015260006155f760a0830185615439565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601c908201527f56616c696461746f72206973206e6f7420696e697469616c697a656400000000604082015260600190565b60608152600a60608201526931b7b6b6b4b9b9b4b7b760b11b608082015260a0602082015260006155f760a0830185615439565b6000602080835260c0830160018060a01b03855116828501528185015160408181870152808701516060870152606087015160808701526080870151915060a08087015282825180855260e0880191508584019450600093505b808410156157b95784518051835286015186830152938501936001939093019290820190615790565b50979650505050505050565b8381526001600160a01b0383166020820152606081016157e483615a16565b826040830152949350505050565b8681526001600160a01b0386166020820152604081018590526060810184905260c081016154d66080830185615465565b84815283151560208201526080810161583f6040830185615465565b82606083015295945050505050565b604051601f8201601f191681016001600160401b038111828210171561587657615876615a00565b604052919050565b60008219821115615891576158916159a8565b500190565b6000826158b357634e487b7160e01b600052601260045260246000fd5b500490565b60008160001904831182151516156158d2576158d26159a8565b500290565b6000828210156158e9576158e96159a8565b500390565b600063ffffffff8381169083168181101561590b5761590b6159a8565b039392505050565b60006138d4368484614f20565b60005b8381101561593b578181015183820152602001615923565b83811115613a485750506000910152565b6000600019821415615960576159606159a8565b5060010190565b600063ffffffff80831681811415615981576159816159a8565b6001019392505050565b6000600160ff1b8214156159a1576159a16159a8565b5060000390565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60048110610df457610df46159be565b8015158114610df457600080fdfe3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512ca2646970667358221220e3040f4c3f4857c9f11c00508c4d1331cea68400917e70a05b4321ae89b9822364736f6c63430008070033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _celerTokenAddress common.Address, _governProposalDeposit *big.Int, _governVoteTimeout *big.Int, _slashTimeout *big.Int, _maxBondedValidators *big.Int, _minValidatorTokens *big.Int, _minSelfDelegation *big.Int, _advanceNoticePeriod *big.Int, _validatorBondInterval *big.Int, _maxSlashFactor *big.Int) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, _celerTokenAddress, _governProposalDeposit, _governVoteTimeout, _slashTimeout, _maxBondedValidators, _minValidatorTokens, _minSelfDelegation, _advanceNoticePeriod, _validatorBondInterval, _maxSlashFactor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_Staking *StakingCaller) BondedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "bondedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_Staking *StakingSession) BondedTokens() (*big.Int, error) {
	return _Staking.Contract.BondedTokens(&_Staking.CallOpts)
}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_Staking *StakingCallerSession) BondedTokens() (*big.Int, error) {
	return _Staking.Contract.BondedTokens(&_Staking.CallOpts)
}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_Staking *StakingCaller) BondedValAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "bondedValAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_Staking *StakingSession) BondedValAddrs(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.BondedValAddrs(&_Staking.CallOpts, arg0)
}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) BondedValAddrs(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.BondedValAddrs(&_Staking.CallOpts, arg0)
}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Staking *StakingCaller) CelerToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "celerToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Staking *StakingSession) CelerToken() (common.Address, error) {
	return _Staking.Contract.CelerToken(&_Staking.CallOpts)
}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Staking *StakingCallerSession) CelerToken() (common.Address, error) {
	return _Staking.Contract.CelerToken(&_Staking.CallOpts)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x1e6f3d8a.
//
// Solidity: function claimedReward(address ) view returns(uint256)
func (_Staking *StakingCaller) ClaimedReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "claimedReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedReward is a free data retrieval call binding the contract method 0x1e6f3d8a.
//
// Solidity: function claimedReward(address ) view returns(uint256)
func (_Staking *StakingSession) ClaimedReward(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.ClaimedReward(&_Staking.CallOpts, arg0)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x1e6f3d8a.
//
// Solidity: function claimedReward(address ) view returns(uint256)
func (_Staking *StakingCallerSession) ClaimedReward(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.ClaimedReward(&_Staking.CallOpts, arg0)
}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_Staking *StakingCaller) GetBondedValidatorNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getBondedValidatorNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_Staking *StakingSession) GetBondedValidatorNum() (*big.Int, error) {
	return _Staking.Contract.GetBondedValidatorNum(&_Staking.CallOpts)
}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_Staking *StakingCallerSession) GetBondedValidatorNum() (*big.Int, error) {
	return _Staking.Contract.GetBondedValidatorNum(&_Staking.CallOpts)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
func (_Staking *StakingCaller) GetDelegatorInfo(opts *bind.CallOpts, _valAddr common.Address, _delAddr common.Address) (StakingDelegatorInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDelegatorInfo", _valAddr, _delAddr)

	if err != nil {
		return *new(StakingDelegatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingDelegatorInfo)).(*StakingDelegatorInfo)

	return out0, err

}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
func (_Staking *StakingSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (StakingDelegatorInfo, error) {
	return _Staking.Contract.GetDelegatorInfo(&_Staking.CallOpts, _valAddr, _delAddr)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
func (_Staking *StakingCallerSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (StakingDelegatorInfo, error) {
	return _Staking.Contract.GetDelegatorInfo(&_Staking.CallOpts, _valAddr, _delAddr)
}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_Staking *StakingCaller) GetMinValidatorTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getMinValidatorTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_Staking *StakingSession) GetMinValidatorTokens() (*big.Int, error) {
	return _Staking.Contract.GetMinValidatorTokens(&_Staking.CallOpts)
}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinValidatorTokens() (*big.Int, error) {
	return _Staking.Contract.GetMinValidatorTokens(&_Staking.CallOpts)
}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Staking *StakingCaller) GetParamProposalVote(opts *bind.CallOpts, _proposalId *big.Int, _voter common.Address) (uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getParamProposalVote", _proposalId, _voter)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Staking *StakingSession) GetParamProposalVote(_proposalId *big.Int, _voter common.Address) (uint8, error) {
	return _Staking.Contract.GetParamProposalVote(&_Staking.CallOpts, _proposalId, _voter)
}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Staking *StakingCallerSession) GetParamProposalVote(_proposalId *big.Int, _voter common.Address) (uint8, error) {
	return _Staking.Contract.GetParamProposalVote(&_Staking.CallOpts, _proposalId, _voter)
}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Staking *StakingCaller) GetParamValue(opts *bind.CallOpts, _name uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getParamValue", _name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Staking *StakingSession) GetParamValue(_name uint8) (*big.Int, error) {
	return _Staking.Contract.GetParamValue(&_Staking.CallOpts, _name)
}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Staking *StakingCallerSession) GetParamValue(_name uint8) (*big.Int, error) {
	return _Staking.Contract.GetParamValue(&_Staking.CallOpts, _name)
}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_Staking *StakingCaller) GetQuorumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getQuorumTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_Staking *StakingSession) GetQuorumTokens() (*big.Int, error) {
	return _Staking.Contract.GetQuorumTokens(&_Staking.CallOpts)
}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_Staking *StakingCallerSession) GetQuorumTokens() (*big.Int, error) {
	return _Staking.Contract.GetQuorumTokens(&_Staking.CallOpts)
}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_Staking *StakingCaller) GetValidatorNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_Staking *StakingSession) GetValidatorNum() (*big.Int, error) {
	return _Staking.Contract.GetValidatorNum(&_Staking.CallOpts)
}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_Staking *StakingCallerSession) GetValidatorNum() (*big.Int, error) {
	return _Staking.Contract.GetValidatorNum(&_Staking.CallOpts)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_Staking *StakingCaller) GetValidatorStatus(opts *bind.CallOpts, _valAddr common.Address) (uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorStatus", _valAddr)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_Staking *StakingSession) GetValidatorStatus(_valAddr common.Address) (uint8, error) {
	return _Staking.Contract.GetValidatorStatus(&_Staking.CallOpts, _valAddr)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_Staking *StakingCallerSession) GetValidatorStatus(_valAddr common.Address) (uint8, error) {
	return _Staking.Contract.GetValidatorStatus(&_Staking.CallOpts, _valAddr)
}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_Staking *StakingCaller) GetValidatorTokens(opts *bind.CallOpts, _valAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorTokens", _valAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_Staking *StakingSession) GetValidatorTokens(_valAddr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetValidatorTokens(&_Staking.CallOpts, _valAddr)
}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_Staking *StakingCallerSession) GetValidatorTokens(_valAddr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetValidatorTokens(&_Staking.CallOpts, _valAddr)
}

// HasMinRequiredTokens is a free data retrieval call binding the contract method 0x47abfdbf.
//
// Solidity: function hasMinRequiredTokens(address _valAddr, bool _checkSelfDelegation) view returns(bool)
func (_Staking *StakingCaller) HasMinRequiredTokens(opts *bind.CallOpts, _valAddr common.Address, _checkSelfDelegation bool) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "hasMinRequiredTokens", _valAddr, _checkSelfDelegation)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMinRequiredTokens is a free data retrieval call binding the contract method 0x47abfdbf.
//
// Solidity: function hasMinRequiredTokens(address _valAddr, bool _checkSelfDelegation) view returns(bool)
func (_Staking *StakingSession) HasMinRequiredTokens(_valAddr common.Address, _checkSelfDelegation bool) (bool, error) {
	return _Staking.Contract.HasMinRequiredTokens(&_Staking.CallOpts, _valAddr, _checkSelfDelegation)
}

// HasMinRequiredTokens is a free data retrieval call binding the contract method 0x47abfdbf.
//
// Solidity: function hasMinRequiredTokens(address _valAddr, bool _checkSelfDelegation) view returns(bool)
func (_Staking *StakingCallerSession) HasMinRequiredTokens(_valAddr common.Address, _checkSelfDelegation bool) (bool, error) {
	return _Staking.Contract.HasMinRequiredTokens(&_Staking.CallOpts, _valAddr, _checkSelfDelegation)
}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_Staking *StakingCaller) IsBondedValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isBondedValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_Staking *StakingSession) IsBondedValidator(_addr common.Address) (bool, error) {
	return _Staking.Contract.IsBondedValidator(&_Staking.CallOpts, _addr)
}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_Staking *StakingCallerSession) IsBondedValidator(_addr common.Address) (bool, error) {
	return _Staking.Contract.IsBondedValidator(&_Staking.CallOpts, _addr)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Staking *StakingCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Staking *StakingSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Staking.Contract.IsWhitelisted(&_Staking.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Staking *StakingCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Staking.Contract.IsWhitelisted(&_Staking.CallOpts, account)
}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_Staking *StakingCaller) NextBondBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "nextBondBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_Staking *StakingSession) NextBondBlock() (*big.Int, error) {
	return _Staking.Contract.NextBondBlock(&_Staking.CallOpts)
}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_Staking *StakingCallerSession) NextBondBlock() (*big.Int, error) {
	return _Staking.Contract.NextBondBlock(&_Staking.CallOpts)
}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Staking *StakingCaller) NextParamProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "nextParamProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Staking *StakingSession) NextParamProposalId() (*big.Int, error) {
	return _Staking.Contract.NextParamProposalId(&_Staking.CallOpts)
}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Staking *StakingCallerSession) NextParamProposalId() (*big.Int, error) {
	return _Staking.Contract.NextParamProposalId(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Staking *StakingCaller) ParamProposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "paramProposals", arg0)

	outstruct := new(struct {
		Proposer     common.Address
		Deposit      *big.Int
		VoteDeadline *big.Int
		Name         uint8
		NewValue     *big.Int
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VoteDeadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.NewValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Staking *StakingSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _Staking.Contract.ParamProposals(&_Staking.CallOpts, arg0)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Staking *StakingCallerSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _Staking.Contract.ParamProposals(&_Staking.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Staking *StakingCaller) Params(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "params", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Staking *StakingSession) Params(arg0 uint8) (*big.Int, error) {
	return _Staking.Contract.Params(&_Staking.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Staking *StakingCallerSession) Params(arg0 uint8) (*big.Int, error) {
	return _Staking.Contract.Params(&_Staking.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCallerSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_Staking *StakingCaller) RewardPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_Staking *StakingSession) RewardPool() (*big.Int, error) {
	return _Staking.Contract.RewardPool(&_Staking.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_Staking *StakingCallerSession) RewardPool() (*big.Int, error) {
	return _Staking.Contract.RewardPool(&_Staking.CallOpts)
}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_Staking *StakingCaller) SignerVals(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "signerVals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_Staking *StakingSession) SignerVals(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.SignerVals(&_Staking.CallOpts, arg0)
}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_Staking *StakingCallerSession) SignerVals(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.SignerVals(&_Staking.CallOpts, arg0)
}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_Staking *StakingCaller) SlashNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "slashNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_Staking *StakingSession) SlashNonces(arg0 *big.Int) (bool, error) {
	return _Staking.Contract.SlashNonces(&_Staking.CallOpts, arg0)
}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_Staking *StakingCallerSession) SlashNonces(arg0 *big.Int) (bool, error) {
	return _Staking.Contract.SlashNonces(&_Staking.CallOpts, arg0)
}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_Staking *StakingCaller) ValAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "valAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_Staking *StakingSession) ValAddrs(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.ValAddrs(&_Staking.CallOpts, arg0)
}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) ValAddrs(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.ValAddrs(&_Staking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 minSelfDelegation, uint64 bondBlock, uint64 unbondBlock, uint64 commissionRate)
func (_Staking *StakingCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	MinSelfDelegation  *big.Int
	BondBlock          uint64
	UnbondBlock        uint64
	CommissionRate     uint64
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Status             uint8
		Signer             common.Address
		Tokens             *big.Int
		Shares             *big.Int
		UndelegationTokens *big.Int
		UndelegationShares *big.Int
		MinSelfDelegation  *big.Int
		BondBlock          uint64
		UnbondBlock        uint64
		CommissionRate     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Signer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Tokens = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Shares = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UndelegationTokens = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UndelegationShares = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MinSelfDelegation = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BondBlock = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.UnbondBlock = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.CommissionRate = *abi.ConvertType(out[9], new(uint64)).(*uint64)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 minSelfDelegation, uint64 bondBlock, uint64 unbondBlock, uint64 commissionRate)
func (_Staking *StakingSession) Validators(arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	MinSelfDelegation  *big.Int
	BondBlock          uint64
	UnbondBlock        uint64
	CommissionRate     uint64
}, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 minSelfDelegation, uint64 bondBlock, uint64 unbondBlock, uint64 commissionRate)
func (_Staking *StakingCallerSession) Validators(arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	MinSelfDelegation  *big.Int
	BondBlock          uint64
	UnbondBlock        uint64
	CommissionRate     uint64
}, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, arg0)
}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_Staking *StakingCaller) VerifySignatures(opts *bind.CallOpts, _msg []byte, _sigs [][]byte) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "verifySignatures", _msg, _sigs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_Staking *StakingSession) VerifySignatures(_msg []byte, _sigs [][]byte) (bool, error) {
	return _Staking.Contract.VerifySignatures(&_Staking.CallOpts, _msg, _sigs)
}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_Staking *StakingCallerSession) VerifySignatures(_msg []byte, _sigs [][]byte) (bool, error) {
	return _Staking.Contract.VerifySignatures(&_Staking.CallOpts, _msg, _sigs)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Whitelist(&_Staking.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Whitelist(&_Staking.CallOpts, arg0)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Staking *StakingCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Staking *StakingSession) WhitelistEnabled() (bool, error) {
	return _Staking.Contract.WhitelistEnabled(&_Staking.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Staking *StakingCallerSession) WhitelistEnabled() (bool, error) {
	return _Staking.Contract.WhitelistEnabled(&_Staking.CallOpts)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Staking *StakingTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Staking *StakingSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddWhitelisted(&_Staking.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Staking *StakingTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddWhitelisted(&_Staking.TransactOpts, account)
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_Staking *StakingTransactor) BondValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "bondValidator")
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_Staking *StakingSession) BondValidator() (*types.Transaction, error) {
	return _Staking.Contract.BondValidator(&_Staking.TransactOpts)
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_Staking *StakingTransactorSession) BondValidator() (*types.Transaction, error) {
	return _Staking.Contract.BondValidator(&_Staking.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_Staking *StakingTransactor) ClaimReward(opts *bind.TransactOpts, _rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claimReward", _rewardRequest, _sigs)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_Staking *StakingSession) ClaimReward(_rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.ClaimReward(&_Staking.TransactOpts, _rewardRequest, _sigs)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_Staking *StakingTransactorSession) ClaimReward(_rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.ClaimReward(&_Staking.TransactOpts, _rewardRequest, _sigs)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_Staking *StakingTransactor) CompleteUndelegate(opts *bind.TransactOpts, _valAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "completeUndelegate", _valAddr)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_Staking *StakingSession) CompleteUndelegate(_valAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.CompleteUndelegate(&_Staking.TransactOpts, _valAddr)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_Staking *StakingTransactorSession) CompleteUndelegate(_valAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.CompleteUndelegate(&_Staking.TransactOpts, _valAddr)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_Staking *StakingTransactor) ConfirmParamProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "confirmParamProposal", _proposalId)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_Staking *StakingSession) ConfirmParamProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ConfirmParamProposal(&_Staking.TransactOpts, _proposalId)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_Staking *StakingTransactorSession) ConfirmParamProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ConfirmParamProposal(&_Staking.TransactOpts, _proposalId)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_Staking *StakingTransactor) ConfirmUnbondedValidator(opts *bind.TransactOpts, _valAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "confirmUnbondedValidator", _valAddr)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_Staking *StakingSession) ConfirmUnbondedValidator(_valAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ConfirmUnbondedValidator(&_Staking.TransactOpts, _valAddr)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_Staking *StakingTransactorSession) ConfirmUnbondedValidator(_valAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ConfirmUnbondedValidator(&_Staking.TransactOpts, _valAddr)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_Staking *StakingTransactor) ContributeToRewardPool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "contributeToRewardPool", _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_Staking *StakingSession) ContributeToRewardPool(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ContributeToRewardPool(&_Staking.TransactOpts, _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_Staking *StakingTransactorSession) ContributeToRewardPool(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ContributeToRewardPool(&_Staking.TransactOpts, _amount)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Staking *StakingTransactor) CreateParamProposal(opts *bind.TransactOpts, _name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "createParamProposal", _name, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Staking *StakingSession) CreateParamProposal(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CreateParamProposal(&_Staking.TransactOpts, _name, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Staking *StakingTransactorSession) CreateParamProposal(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CreateParamProposal(&_Staking.TransactOpts, _name, _value)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, _valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", _valAddr, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingSession) Delegate(_valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, _valAddr, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingTransactorSession) Delegate(_valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, _valAddr, _tokens)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_Staking *StakingTransactor) DrainToken(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "drainToken", _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_Staking *StakingSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DrainToken(&_Staking.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_Staking *StakingTransactorSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DrainToken(&_Staking.TransactOpts, _amount)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x525eba21.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint64 _commissionRate) returns()
func (_Staking *StakingTransactor) InitializeValidator(opts *bind.TransactOpts, _signer common.Address, _minSelfDelegation *big.Int, _commissionRate uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initializeValidator", _signer, _minSelfDelegation, _commissionRate)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x525eba21.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint64 _commissionRate) returns()
func (_Staking *StakingSession) InitializeValidator(_signer common.Address, _minSelfDelegation *big.Int, _commissionRate uint64) (*types.Transaction, error) {
	return _Staking.Contract.InitializeValidator(&_Staking.TransactOpts, _signer, _minSelfDelegation, _commissionRate)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x525eba21.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint64 _commissionRate) returns()
func (_Staking *StakingTransactorSession) InitializeValidator(_signer common.Address, _minSelfDelegation *big.Int, _commissionRate uint64) (*types.Transaction, error) {
	return _Staking.Contract.InitializeValidator(&_Staking.TransactOpts, _signer, _minSelfDelegation, _commissionRate)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactorSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Staking *StakingTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Staking *StakingSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveWhitelisted(&_Staking.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Staking *StakingTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveWhitelisted(&_Staking.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// SetMaxSlashFactor is a paid mutator transaction binding the contract method 0x1a203257.
//
// Solidity: function setMaxSlashFactor(uint256 _maxSlashFactor) returns()
func (_Staking *StakingTransactor) SetMaxSlashFactor(opts *bind.TransactOpts, _maxSlashFactor *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMaxSlashFactor", _maxSlashFactor)
}

// SetMaxSlashFactor is a paid mutator transaction binding the contract method 0x1a203257.
//
// Solidity: function setMaxSlashFactor(uint256 _maxSlashFactor) returns()
func (_Staking *StakingSession) SetMaxSlashFactor(_maxSlashFactor *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxSlashFactor(&_Staking.TransactOpts, _maxSlashFactor)
}

// SetMaxSlashFactor is a paid mutator transaction binding the contract method 0x1a203257.
//
// Solidity: function setMaxSlashFactor(uint256 _maxSlashFactor) returns()
func (_Staking *StakingTransactorSession) SetMaxSlashFactor(_maxSlashFactor *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxSlashFactor(&_Staking.TransactOpts, _maxSlashFactor)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Staking *StakingTransactor) SetWhitelistEnabled(opts *bind.TransactOpts, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setWhitelistEnabled", _whitelistEnabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Staking *StakingSession) SetWhitelistEnabled(_whitelistEnabled bool) (*types.Transaction, error) {
	return _Staking.Contract.SetWhitelistEnabled(&_Staking.TransactOpts, _whitelistEnabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Staking *StakingTransactorSession) SetWhitelistEnabled(_whitelistEnabled bool) (*types.Transaction, error) {
	return _Staking.Contract.SetWhitelistEnabled(&_Staking.TransactOpts, _whitelistEnabled)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts, _slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash", _slashRequest, _sigs)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_Staking *StakingSession) Slash(_slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _slashRequest, _sigs)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_Staking *StakingTransactorSession) Slash(_slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _slashRequest, _sigs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingTransactor) Undelegate(opts *bind.TransactOpts, _valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegate", _valAddr, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingSession) Undelegate(_valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, _valAddr, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingTransactorSession) Undelegate(_valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, _valAddr, _shares)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x49955e39.
//
// Solidity: function updateCommissionRate(uint64 _newRate) returns()
func (_Staking *StakingTransactor) UpdateCommissionRate(opts *bind.TransactOpts, _newRate uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateCommissionRate", _newRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x49955e39.
//
// Solidity: function updateCommissionRate(uint64 _newRate) returns()
func (_Staking *StakingSession) UpdateCommissionRate(_newRate uint64) (*types.Transaction, error) {
	return _Staking.Contract.UpdateCommissionRate(&_Staking.TransactOpts, _newRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x49955e39.
//
// Solidity: function updateCommissionRate(uint64 _newRate) returns()
func (_Staking *StakingTransactorSession) UpdateCommissionRate(_newRate uint64) (*types.Transaction, error) {
	return _Staking.Contract.UpdateCommissionRate(&_Staking.TransactOpts, _newRate)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_Staking *StakingTransactor) UpdateMinSelfDelegation(opts *bind.TransactOpts, _minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateMinSelfDelegation", _minSelfDelegation)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_Staking *StakingSession) UpdateMinSelfDelegation(_minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateMinSelfDelegation(&_Staking.TransactOpts, _minSelfDelegation)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_Staking *StakingTransactorSession) UpdateMinSelfDelegation(_minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateMinSelfDelegation(&_Staking.TransactOpts, _minSelfDelegation)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_Staking *StakingTransactor) UpdateValidatorSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateValidatorSigner", _signer)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_Staking *StakingSession) UpdateValidatorSigner(_signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateValidatorSigner(&_Staking.TransactOpts, _signer)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_Staking *StakingTransactorSession) UpdateValidatorSigner(_signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateValidatorSigner(&_Staking.TransactOpts, _signer)
}

// ValidatorNotice is a paid mutator transaction binding the contract method 0x9146f110.
//
// Solidity: function validatorNotice(address _valAddr, string _key, bytes _data) returns()
func (_Staking *StakingTransactor) ValidatorNotice(opts *bind.TransactOpts, _valAddr common.Address, _key string, _data []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "validatorNotice", _valAddr, _key, _data)
}

// ValidatorNotice is a paid mutator transaction binding the contract method 0x9146f110.
//
// Solidity: function validatorNotice(address _valAddr, string _key, bytes _data) returns()
func (_Staking *StakingSession) ValidatorNotice(_valAddr common.Address, _key string, _data []byte) (*types.Transaction, error) {
	return _Staking.Contract.ValidatorNotice(&_Staking.TransactOpts, _valAddr, _key, _data)
}

// ValidatorNotice is a paid mutator transaction binding the contract method 0x9146f110.
//
// Solidity: function validatorNotice(address _valAddr, string _key, bytes _data) returns()
func (_Staking *StakingTransactorSession) ValidatorNotice(_valAddr common.Address, _key string, _data []byte) (*types.Transaction, error) {
	return _Staking.Contract.ValidatorNotice(&_Staking.TransactOpts, _valAddr, _key, _data)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_Staking *StakingTransactor) VoteParam(opts *bind.TransactOpts, _proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "voteParam", _proposalId, _vote)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_Staking *StakingSession) VoteParam(_proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _Staking.Contract.VoteParam(&_Staking.TransactOpts, _proposalId, _vote)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_Staking *StakingTransactorSession) VoteParam(_proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _Staking.Contract.VoteParam(&_Staking.TransactOpts, _proposalId, _vote)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// StakingConfirmParamProposalIterator is returned from FilterConfirmParamProposal and is used to iterate over the raw logs and unpacked data for ConfirmParamProposal events raised by the Staking contract.
type StakingConfirmParamProposalIterator struct {
	Event *StakingConfirmParamProposal // Event containing the contract specifics and raw log

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
func (it *StakingConfirmParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingConfirmParamProposal)
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
		it.Event = new(StakingConfirmParamProposal)
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
func (it *StakingConfirmParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingConfirmParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingConfirmParamProposal represents a ConfirmParamProposal event raised by the Staking contract.
type StakingConfirmParamProposal struct {
	ProposalId *big.Int
	Passed     bool
	Name       uint8
	NewValue   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfirmParamProposal is a free log retrieval operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Staking *StakingFilterer) FilterConfirmParamProposal(opts *bind.FilterOpts) (*StakingConfirmParamProposalIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return &StakingConfirmParamProposalIterator{contract: _Staking.contract, event: "ConfirmParamProposal", logs: logs, sub: sub}, nil
}

// WatchConfirmParamProposal is a free log subscription operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Staking *StakingFilterer) WatchConfirmParamProposal(opts *bind.WatchOpts, sink chan<- *StakingConfirmParamProposal) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingConfirmParamProposal)
				if err := _Staking.contract.UnpackLog(event, "ConfirmParamProposal", log); err != nil {
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

// ParseConfirmParamProposal is a log parse operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Staking *StakingFilterer) ParseConfirmParamProposal(log types.Log) (*StakingConfirmParamProposal, error) {
	event := new(StakingConfirmParamProposal)
	if err := _Staking.contract.UnpackLog(event, "ConfirmParamProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCreateParamProposalIterator is returned from FilterCreateParamProposal and is used to iterate over the raw logs and unpacked data for CreateParamProposal events raised by the Staking contract.
type StakingCreateParamProposalIterator struct {
	Event *StakingCreateParamProposal // Event containing the contract specifics and raw log

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
func (it *StakingCreateParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCreateParamProposal)
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
		it.Event = new(StakingCreateParamProposal)
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
func (it *StakingCreateParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCreateParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCreateParamProposal represents a CreateParamProposal event raised by the Staking contract.
type StakingCreateParamProposal struct {
	ProposalId   *big.Int
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateParamProposal is a free log retrieval operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Staking *StakingFilterer) FilterCreateParamProposal(opts *bind.FilterOpts) (*StakingCreateParamProposalIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return &StakingCreateParamProposalIterator{contract: _Staking.contract, event: "CreateParamProposal", logs: logs, sub: sub}, nil
}

// WatchCreateParamProposal is a free log subscription operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Staking *StakingFilterer) WatchCreateParamProposal(opts *bind.WatchOpts, sink chan<- *StakingCreateParamProposal) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCreateParamProposal)
				if err := _Staking.contract.UnpackLog(event, "CreateParamProposal", log); err != nil {
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

// ParseCreateParamProposal is a log parse operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Staking *StakingFilterer) ParseCreateParamProposal(log types.Log) (*StakingCreateParamProposal, error) {
	event := new(StakingCreateParamProposal)
	if err := _Staking.contract.UnpackLog(event, "CreateParamProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegationUpdateIterator is returned from FilterDelegationUpdate and is used to iterate over the raw logs and unpacked data for DelegationUpdate events raised by the Staking contract.
type StakingDelegationUpdateIterator struct {
	Event *StakingDelegationUpdate // Event containing the contract specifics and raw log

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
func (it *StakingDelegationUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegationUpdate)
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
		it.Event = new(StakingDelegationUpdate)
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
func (it *StakingDelegationUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegationUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegationUpdate represents a DelegationUpdate event raised by the Staking contract.
type StakingDelegationUpdate struct {
	ValAddr   common.Address
	DelAddr   common.Address
	ValTokens *big.Int
	DelShares *big.Int
	TokenDiff *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationUpdate is a free log retrieval operation binding the contract event 0x2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea.
//
// Solidity: event DelegationUpdate(address indexed valAddr, address indexed delAddr, uint256 valTokens, uint256 delShares, int256 tokenDiff)
func (_Staking *StakingFilterer) FilterDelegationUpdate(opts *bind.FilterOpts, valAddr []common.Address, delAddr []common.Address) (*StakingDelegationUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegationUpdate", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegationUpdateIterator{contract: _Staking.contract, event: "DelegationUpdate", logs: logs, sub: sub}, nil
}

// WatchDelegationUpdate is a free log subscription operation binding the contract event 0x2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea.
//
// Solidity: event DelegationUpdate(address indexed valAddr, address indexed delAddr, uint256 valTokens, uint256 delShares, int256 tokenDiff)
func (_Staking *StakingFilterer) WatchDelegationUpdate(opts *bind.WatchOpts, sink chan<- *StakingDelegationUpdate, valAddr []common.Address, delAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegationUpdate", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegationUpdate)
				if err := _Staking.contract.UnpackLog(event, "DelegationUpdate", log); err != nil {
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

// ParseDelegationUpdate is a log parse operation binding the contract event 0x2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea.
//
// Solidity: event DelegationUpdate(address indexed valAddr, address indexed delAddr, uint256 valTokens, uint256 delShares, int256 tokenDiff)
func (_Staking *StakingFilterer) ParseDelegationUpdate(log types.Log) (*StakingDelegationUpdate, error) {
	event := new(StakingDelegationUpdate)
	if err := _Staking.contract.UnpackLog(event, "DelegationUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Staking contract.
type StakingPausedIterator struct {
	Event *StakingPaused // Event containing the contract specifics and raw log

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
func (it *StakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPaused)
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
		it.Event = new(StakingPaused)
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
func (it *StakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPaused represents a Paused event raised by the Staking contract.
type StakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingPausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingPausedIterator{contract: _Staking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingPaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPaused)
				if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Staking *StakingFilterer) ParsePaused(log types.Log) (*StakingPaused, error) {
	event := new(StakingPaused)
	if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Staking contract.
type StakingRewardClaimedIterator struct {
	Event *StakingRewardClaimed // Event containing the contract specifics and raw log

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
func (it *StakingRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardClaimed)
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
		it.Event = new(StakingRewardClaimed)
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
func (it *StakingRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardClaimed represents a RewardClaimed event raised by the Staking contract.
type StakingRewardClaimed struct {
	Recipient  common.Address
	Reward     *big.Int
	RewardPool *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0xf01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743.
//
// Solidity: event RewardClaimed(address indexed recipient, uint256 reward, uint256 rewardPool)
func (_Staking *StakingFilterer) FilterRewardClaimed(opts *bind.FilterOpts, recipient []common.Address) (*StakingRewardClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardClaimedIterator{contract: _Staking.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0xf01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743.
//
// Solidity: event RewardClaimed(address indexed recipient, uint256 reward, uint256 rewardPool)
func (_Staking *StakingFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *StakingRewardClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardClaimed)
				if err := _Staking.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0xf01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743.
//
// Solidity: event RewardClaimed(address indexed recipient, uint256 reward, uint256 rewardPool)
func (_Staking *StakingFilterer) ParseRewardClaimed(log types.Log) (*StakingRewardClaimed, error) {
	event := new(StakingRewardClaimed)
	if err := _Staking.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardPoolContributionIterator is returned from FilterRewardPoolContribution and is used to iterate over the raw logs and unpacked data for RewardPoolContribution events raised by the Staking contract.
type StakingRewardPoolContributionIterator struct {
	Event *StakingRewardPoolContribution // Event containing the contract specifics and raw log

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
func (it *StakingRewardPoolContributionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardPoolContribution)
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
		it.Event = new(StakingRewardPoolContribution)
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
func (it *StakingRewardPoolContributionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardPoolContributionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardPoolContribution represents a RewardPoolContribution event raised by the Staking contract.
type StakingRewardPoolContribution struct {
	Contributor    common.Address
	Contribution   *big.Int
	RewardPoolSize *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardPoolContribution is a free log retrieval operation binding the contract event 0x574fb579a16e67a50f3626feb14c7c23cb5d3c7ff880e0d3e2f792fbe8d22fca.
//
// Solidity: event RewardPoolContribution(address indexed contributor, uint256 contribution, uint256 rewardPoolSize)
func (_Staking *StakingFilterer) FilterRewardPoolContribution(opts *bind.FilterOpts, contributor []common.Address) (*StakingRewardPoolContributionIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardPoolContribution", contributorRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardPoolContributionIterator{contract: _Staking.contract, event: "RewardPoolContribution", logs: logs, sub: sub}, nil
}

// WatchRewardPoolContribution is a free log subscription operation binding the contract event 0x574fb579a16e67a50f3626feb14c7c23cb5d3c7ff880e0d3e2f792fbe8d22fca.
//
// Solidity: event RewardPoolContribution(address indexed contributor, uint256 contribution, uint256 rewardPoolSize)
func (_Staking *StakingFilterer) WatchRewardPoolContribution(opts *bind.WatchOpts, sink chan<- *StakingRewardPoolContribution, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardPoolContribution", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardPoolContribution)
				if err := _Staking.contract.UnpackLog(event, "RewardPoolContribution", log); err != nil {
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

// ParseRewardPoolContribution is a log parse operation binding the contract event 0x574fb579a16e67a50f3626feb14c7c23cb5d3c7ff880e0d3e2f792fbe8d22fca.
//
// Solidity: event RewardPoolContribution(address indexed contributor, uint256 contribution, uint256 rewardPoolSize)
func (_Staking *StakingFilterer) ParseRewardPoolContribution(log types.Log) (*StakingRewardPoolContribution, error) {
	event := new(StakingRewardPoolContribution)
	if err := _Staking.contract.UnpackLog(event, "RewardPoolContribution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlashIterator is returned from FilterSlash and is used to iterate over the raw logs and unpacked data for Slash events raised by the Staking contract.
type StakingSlashIterator struct {
	Event *StakingSlash // Event containing the contract specifics and raw log

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
func (it *StakingSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlash)
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
		it.Event = new(StakingSlash)
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
func (it *StakingSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlash represents a Slash event raised by the Staking contract.
type StakingSlash struct {
	ValAddr  common.Address
	Nonce    uint64
	SlashAmt *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlash is a free log retrieval operation binding the contract event 0x10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008.
//
// Solidity: event Slash(address indexed valAddr, uint64 nonce, uint256 slashAmt)
func (_Staking *StakingFilterer) FilterSlash(opts *bind.FilterOpts, valAddr []common.Address) (*StakingSlashIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Slash", valAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlashIterator{contract: _Staking.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0x10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008.
//
// Solidity: event Slash(address indexed valAddr, uint64 nonce, uint256 slashAmt)
func (_Staking *StakingFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *StakingSlash, valAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Slash", valAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlash)
				if err := _Staking.contract.UnpackLog(event, "Slash", log); err != nil {
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

// ParseSlash is a log parse operation binding the contract event 0x10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008.
//
// Solidity: event Slash(address indexed valAddr, uint64 nonce, uint256 slashAmt)
func (_Staking *StakingFilterer) ParseSlash(log types.Log) (*StakingSlash, error) {
	event := new(StakingSlash)
	if err := _Staking.contract.UnpackLog(event, "Slash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlashAmtCollectedIterator is returned from FilterSlashAmtCollected and is used to iterate over the raw logs and unpacked data for SlashAmtCollected events raised by the Staking contract.
type StakingSlashAmtCollectedIterator struct {
	Event *StakingSlashAmtCollected // Event containing the contract specifics and raw log

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
func (it *StakingSlashAmtCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlashAmtCollected)
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
		it.Event = new(StakingSlashAmtCollected)
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
func (it *StakingSlashAmtCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlashAmtCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlashAmtCollected represents a SlashAmtCollected event raised by the Staking contract.
type StakingSlashAmtCollected struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashAmtCollected is a free log retrieval operation binding the contract event 0xb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3.
//
// Solidity: event SlashAmtCollected(address indexed recipient, uint256 amount)
func (_Staking *StakingFilterer) FilterSlashAmtCollected(opts *bind.FilterOpts, recipient []common.Address) (*StakingSlashAmtCollectedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SlashAmtCollected", recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlashAmtCollectedIterator{contract: _Staking.contract, event: "SlashAmtCollected", logs: logs, sub: sub}, nil
}

// WatchSlashAmtCollected is a free log subscription operation binding the contract event 0xb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3.
//
// Solidity: event SlashAmtCollected(address indexed recipient, uint256 amount)
func (_Staking *StakingFilterer) WatchSlashAmtCollected(opts *bind.WatchOpts, sink chan<- *StakingSlashAmtCollected, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SlashAmtCollected", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlashAmtCollected)
				if err := _Staking.contract.UnpackLog(event, "SlashAmtCollected", log); err != nil {
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

// ParseSlashAmtCollected is a log parse operation binding the contract event 0xb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3.
//
// Solidity: event SlashAmtCollected(address indexed recipient, uint256 amount)
func (_Staking *StakingFilterer) ParseSlashAmtCollected(log types.Log) (*StakingSlashAmtCollected, error) {
	event := new(StakingSlashAmtCollected)
	if err := _Staking.contract.UnpackLog(event, "SlashAmtCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Staking contract.
type StakingUndelegatedIterator struct {
	Event *StakingUndelegated // Event containing the contract specifics and raw log

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
func (it *StakingUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUndelegated)
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
		it.Event = new(StakingUndelegated)
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
func (it *StakingUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUndelegated represents a Undelegated event raised by the Staking contract.
type StakingUndelegated struct {
	ValAddr common.Address
	DelAddr common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed valAddr, address indexed delAddr, uint256 amount)
func (_Staking *StakingFilterer) FilterUndelegated(opts *bind.FilterOpts, valAddr []common.Address, delAddr []common.Address) (*StakingUndelegatedIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Undelegated", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingUndelegatedIterator{contract: _Staking.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed valAddr, address indexed delAddr, uint256 amount)
func (_Staking *StakingFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *StakingUndelegated, valAddr []common.Address, delAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Undelegated", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUndelegated)
				if err := _Staking.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed valAddr, address indexed delAddr, uint256 amount)
func (_Staking *StakingFilterer) ParseUndelegated(log types.Log) (*StakingUndelegated, error) {
	event := new(StakingUndelegated)
	if err := _Staking.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Staking contract.
type StakingUnpausedIterator struct {
	Event *StakingUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnpaused)
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
		it.Event = new(StakingUnpaused)
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
func (it *StakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnpaused represents a Unpaused event raised by the Staking contract.
type StakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingUnpausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingUnpausedIterator{contract: _Staking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnpaused)
				if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Staking *StakingFilterer) ParseUnpaused(log types.Log) (*StakingUnpaused, error) {
	event := new(StakingUnpaused)
	if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorNoticeIterator is returned from FilterValidatorNotice and is used to iterate over the raw logs and unpacked data for ValidatorNotice events raised by the Staking contract.
type StakingValidatorNoticeIterator struct {
	Event *StakingValidatorNotice // Event containing the contract specifics and raw log

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
func (it *StakingValidatorNoticeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorNotice)
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
		it.Event = new(StakingValidatorNotice)
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
func (it *StakingValidatorNoticeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorNoticeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorNotice represents a ValidatorNotice event raised by the Staking contract.
type StakingValidatorNotice struct {
	ValAddr common.Address
	Key     string
	Data    []byte
	From    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorNotice is a free log retrieval operation binding the contract event 0x3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c.
//
// Solidity: event ValidatorNotice(address indexed valAddr, string key, bytes data, address from)
func (_Staking *StakingFilterer) FilterValidatorNotice(opts *bind.FilterOpts, valAddr []common.Address) (*StakingValidatorNoticeIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorNotice", valAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorNoticeIterator{contract: _Staking.contract, event: "ValidatorNotice", logs: logs, sub: sub}, nil
}

// WatchValidatorNotice is a free log subscription operation binding the contract event 0x3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c.
//
// Solidity: event ValidatorNotice(address indexed valAddr, string key, bytes data, address from)
func (_Staking *StakingFilterer) WatchValidatorNotice(opts *bind.WatchOpts, sink chan<- *StakingValidatorNotice, valAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorNotice", valAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorNotice)
				if err := _Staking.contract.UnpackLog(event, "ValidatorNotice", log); err != nil {
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

// ParseValidatorNotice is a log parse operation binding the contract event 0x3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c.
//
// Solidity: event ValidatorNotice(address indexed valAddr, string key, bytes data, address from)
func (_Staking *StakingFilterer) ParseValidatorNotice(log types.Log) (*StakingValidatorNotice, error) {
	event := new(StakingValidatorNotice)
	if err := _Staking.contract.UnpackLog(event, "ValidatorNotice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorStatusUpdateIterator is returned from FilterValidatorStatusUpdate and is used to iterate over the raw logs and unpacked data for ValidatorStatusUpdate events raised by the Staking contract.
type StakingValidatorStatusUpdateIterator struct {
	Event *StakingValidatorStatusUpdate // Event containing the contract specifics and raw log

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
func (it *StakingValidatorStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorStatusUpdate)
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
		it.Event = new(StakingValidatorStatusUpdate)
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
func (it *StakingValidatorStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorStatusUpdate represents a ValidatorStatusUpdate event raised by the Staking contract.
type StakingValidatorStatusUpdate struct {
	ValAddr common.Address
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusUpdate is a free log retrieval operation binding the contract event 0xd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a.
//
// Solidity: event ValidatorStatusUpdate(address indexed valAddr, uint8 indexed status)
func (_Staking *StakingFilterer) FilterValidatorStatusUpdate(opts *bind.FilterOpts, valAddr []common.Address, status []uint8) (*StakingValidatorStatusUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorStatusUpdate", valAddrRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorStatusUpdateIterator{contract: _Staking.contract, event: "ValidatorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusUpdate is a free log subscription operation binding the contract event 0xd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a.
//
// Solidity: event ValidatorStatusUpdate(address indexed valAddr, uint8 indexed status)
func (_Staking *StakingFilterer) WatchValidatorStatusUpdate(opts *bind.WatchOpts, sink chan<- *StakingValidatorStatusUpdate, valAddr []common.Address, status []uint8) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorStatusUpdate", valAddrRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorStatusUpdate)
				if err := _Staking.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
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

// ParseValidatorStatusUpdate is a log parse operation binding the contract event 0xd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a.
//
// Solidity: event ValidatorStatusUpdate(address indexed valAddr, uint8 indexed status)
func (_Staking *StakingFilterer) ParseValidatorStatusUpdate(log types.Log) (*StakingValidatorStatusUpdate, error) {
	event := new(StakingValidatorStatusUpdate)
	if err := _Staking.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVoteParamIterator is returned from FilterVoteParam and is used to iterate over the raw logs and unpacked data for VoteParam events raised by the Staking contract.
type StakingVoteParamIterator struct {
	Event *StakingVoteParam // Event containing the contract specifics and raw log

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
func (it *StakingVoteParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVoteParam)
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
		it.Event = new(StakingVoteParam)
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
func (it *StakingVoteParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVoteParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVoteParam represents a VoteParam event raised by the Staking contract.
type StakingVoteParam struct {
	ProposalId *big.Int
	Voter      common.Address
	Vote       uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteParam is a free log retrieval operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Staking *StakingFilterer) FilterVoteParam(opts *bind.FilterOpts) (*StakingVoteParamIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return &StakingVoteParamIterator{contract: _Staking.contract, event: "VoteParam", logs: logs, sub: sub}, nil
}

// WatchVoteParam is a free log subscription operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Staking *StakingFilterer) WatchVoteParam(opts *bind.WatchOpts, sink chan<- *StakingVoteParam) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVoteParam)
				if err := _Staking.contract.UnpackLog(event, "VoteParam", log); err != nil {
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

// ParseVoteParam is a log parse operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Staking *StakingFilterer) ParseVoteParam(log types.Log) (*StakingVoteParam, error) {
	event := new(StakingVoteParam)
	if err := _Staking.contract.UnpackLog(event, "VoteParam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the Staking contract.
type StakingWhitelistedAddedIterator struct {
	Event *StakingWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *StakingWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWhitelistedAdded)
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
		it.Event = new(StakingWhitelistedAdded)
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
func (it *StakingWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWhitelistedAdded represents a WhitelistedAdded event raised by the Staking contract.
type StakingWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Staking *StakingFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts) (*StakingWhitelistedAddedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return &StakingWhitelistedAddedIterator{contract: _Staking.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Staking *StakingFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *StakingWhitelistedAdded) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWhitelistedAdded)
				if err := _Staking.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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

// ParseWhitelistedAdded is a log parse operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Staking *StakingFilterer) ParseWhitelistedAdded(log types.Log) (*StakingWhitelistedAdded, error) {
	event := new(StakingWhitelistedAdded)
	if err := _Staking.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the Staking contract.
type StakingWhitelistedRemovedIterator struct {
	Event *StakingWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *StakingWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWhitelistedRemoved)
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
		it.Event = new(StakingWhitelistedRemoved)
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
func (it *StakingWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWhitelistedRemoved represents a WhitelistedRemoved event raised by the Staking contract.
type StakingWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Staking *StakingFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts) (*StakingWhitelistedRemovedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return &StakingWhitelistedRemovedIterator{contract: _Staking.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Staking *StakingFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *StakingWhitelistedRemoved) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWhitelistedRemoved)
				if err := _Staking.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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

// ParseWhitelistedRemoved is a log parse operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Staking *StakingFilterer) ParseWhitelistedRemoved(log types.Log) (*StakingWhitelistedRemoved, error) {
	event := new(StakingWhitelistedRemoved)
	if err := _Staking.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistMetaData contains all meta data concerning the Whitelist contract.
var WhitelistMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WhitelistABI is the input ABI used to generate the binding from.
// Deprecated: Use WhitelistMetaData.ABI instead.
var WhitelistABI = WhitelistMetaData.ABI

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
	WhitelistFilterer   // Log filterer for contract events
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// NewWhitelistFilterer creates a new log filterer instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistFilterer, error) {
	contract, err := bindWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistFilterer{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Whitelist *WhitelistCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Whitelist *WhitelistSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Whitelist.Contract.IsWhitelisted(&_Whitelist.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Whitelist *WhitelistCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Whitelist.Contract.IsWhitelisted(&_Whitelist.CallOpts, account)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Whitelist *WhitelistCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Whitelist *WhitelistSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Whitelist.Contract.Whitelist(&_Whitelist.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Whitelist *WhitelistCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Whitelist.Contract.Whitelist(&_Whitelist.CallOpts, arg0)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Whitelist *WhitelistCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Whitelist *WhitelistSession) WhitelistEnabled() (bool, error) {
	return _Whitelist.Contract.WhitelistEnabled(&_Whitelist.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Whitelist *WhitelistCallerSession) WhitelistEnabled() (bool, error) {
	return _Whitelist.Contract.WhitelistEnabled(&_Whitelist.CallOpts)
}

// WhitelistWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the Whitelist contract.
type WhitelistWhitelistedAddedIterator struct {
	Event *WhitelistWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *WhitelistWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistWhitelistedAdded)
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
		it.Event = new(WhitelistWhitelistedAdded)
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
func (it *WhitelistWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistWhitelistedAdded represents a WhitelistedAdded event raised by the Whitelist contract.
type WhitelistWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Whitelist *WhitelistFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts) (*WhitelistWhitelistedAddedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return &WhitelistWhitelistedAddedIterator{contract: _Whitelist.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Whitelist *WhitelistFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *WhitelistWhitelistedAdded) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistWhitelistedAdded)
				if err := _Whitelist.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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

// ParseWhitelistedAdded is a log parse operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Whitelist *WhitelistFilterer) ParseWhitelistedAdded(log types.Log) (*WhitelistWhitelistedAdded, error) {
	event := new(WhitelistWhitelistedAdded)
	if err := _Whitelist.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the Whitelist contract.
type WhitelistWhitelistedRemovedIterator struct {
	Event *WhitelistWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *WhitelistWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistWhitelistedRemoved)
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
		it.Event = new(WhitelistWhitelistedRemoved)
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
func (it *WhitelistWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistWhitelistedRemoved represents a WhitelistedRemoved event raised by the Whitelist contract.
type WhitelistWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Whitelist *WhitelistFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts) (*WhitelistWhitelistedRemovedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return &WhitelistWhitelistedRemovedIterator{contract: _Whitelist.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Whitelist *WhitelistFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistWhitelistedRemoved) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistWhitelistedRemoved)
				if err := _Whitelist.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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

// ParseWhitelistedRemoved is a log parse operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Whitelist *WhitelistFilterer) ParseWhitelistedRemoved(log types.Log) (*WhitelistWhitelistedRemoved, error) {
	event := new(WhitelistWhitelistedRemoved)
	if err := _Whitelist.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
