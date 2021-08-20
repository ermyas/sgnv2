// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// DPoSDelegatorInfo is an auto generated low-level Go binding around an user-defined struct.
type DPoSDelegatorInfo struct {
	ValAddr            common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	Undelegations      []DPoSUndelegation
}

// DPoSUndelegation is an auto generated low-level Go binding around an user-defined struct.
type DPoSUndelegation struct {
	Shares        *big.Int
	CreationBlock *big.Int
}

// DPoSMetaData contains all meta data concerning the DPoS contract.
var DPoSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_governProposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_governVoteTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorBondInterval\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ConfirmParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CreateParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"tokenDiff\",\"type\":\"int256\"}],\"name\":\"DelegationUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPoolSize\",\"type\":\"uint256\"}],\"name\":\"MiningPoolContribution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPool\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmt\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashAmtCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"name\":\"ValidatorParamsUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDPoS.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGovern.VoteType\",\"name\":\"voteType\",\"type\":\"uint8\"}],\"name\":\"VoteParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bondedValAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"celerToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rewardRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"completeUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"confirmParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"confirmUnbondedValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToMiningPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"createParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDPoS.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structDPoS.DelegatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDPoS.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structDPoS.DelegatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getParamProposalVote\",\"outputs\":[{\"internalType\":\"enumGovern.VoteType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"}],\"name\":\"getParamValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"enumDPoS.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"initializeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBondedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBondBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextParamProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paramProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerVals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_slashRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"}],\"name\":\"updateMinSelfDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"updateValidatorSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"valAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"enumDPoS.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.VoteType\",\"name\":\"_vote\",\"type\":\"uint8\"}],\"name\":\"voteParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162005b6e38038062005b6e833981016040819052620000349162000203565b8888888888888888886200004833620001b3565b6000805460ff60a01b19168155600280546001600160a01b039b909b1661010002610100600160a81b0319909b169a909a1790995560036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff979097557fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c959095557fc3a24b0501bd2c13a7e57f2db4369ec4c223447539fc0724a9d55ac4a06ebd4d939093557fcbc4e5fb02c3d1de23a9f1e014b4d2ee5aeaea9505df5e855c9210bf472495af919091557f83ec6a1f0257b830b5e016457c9cf1435391bf56cc98f369a58a54fe93772465557f405aad32e1adbac89bb7f176e338b8fc6e994ca210c9bb7bdca249b465942250557fc69056f16cbaa3c616b828e333ab7d3a32310765507f8f58359e99ebb7a885f35560079091527ff2c49132ed1cee2a7e75bde50d332a2f81f1d01e5456d8a19d1df09bd561dbd2555062000285975050505050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008060008060008060008060006101208a8c0312156200022357600080fd5b89516001600160a01b03811681146200023b57600080fd5b8099505060208a0151975060408a0151965060608a0151955060808a0151945060a08a0151935060c08a0151925060e08a015191506101008a015190509295985092959850929598565b6158d980620002956000396000f3fe60806040526004361061036e5760003560e01c80636d308783116101c6578063934a18ec116100f7578063d6b0f48411610095578063eecefef81161006f578063eecefef814610a75578063f2fde38b14610aa2578063f8df0dc514610ac2578063fa52c7d814610ae257600080fd5b8063d6b0f48414610a13578063e478ed9d14610a28578063eb505dd514610a4857600080fd5b8063b4f7fa34116100d1578063b4f7fa3414610980578063c6c21e9d146109a0578063c8f9f984146109c5578063cdfb2b4e146109fe57600080fd5b8063934a18ec14610907578063a310624f14610927578063acc62ccf1461096057600080fd5b80638456cb59116101645780638da5cb5b1161013e5780638da5cb5b146108845780638dc2336d146108a257806390e360f8146108b757806392bb243c146108e757600080fd5b80638456cb591461083a57806389f9aab51461084f5780638a74d5fe1461086457600080fd5b80637a3ba4ad116101a05780637a3ba4ad1461077c5780637a50dbd2146107915780637e5fb8f3146107b157806383cfb3181461082457600080fd5b80636d308783146106f9578063715018a61461074757806371bc02161461075c57600080fd5b8063386c024a116102a05780634d99dd161161023e5780635c975abb116102185780635c975abb1461068e5780635e593eff146106ad57806365d5d420146106cd57806366666aa9146106e357600080fd5b80634d99dd161461062757806351fb012d14610647578063581c53c51461066157600080fd5b80633f4ba83a1161027a5780633f4ba83a146105b2578063410ae02c146105c7578063473849bd146105e75780634b7dba6b1461060757600080fd5b8063386c024a146105445780633985c4e6146105595780633af32abf1461057957600080fd5b806324990d7b1161030d578063291d9549116102e7578063291d9549146104b8578063313019bb146104d857806336f1635f146105055780633773e4891461051a57600080fd5b806324990d7b1461046357806324b9bcc01461048357806325ed6b351461049857600080fd5b8063145aa11611610349578063145aa116146103dc5780631cfe4f0b146103fc5780631e6f3d8a1461042057806322da79271461044d57600080fd5b8062fa3d501461037a578063026e402b1461039c57806310154bad146103bc57600080fd5b3661037557005b600080fd5b34801561038657600080fd5b5061039a6103953660046152fc565b610b65565b005b3480156103a857600080fd5b5061039a6103b736600461510b565b610c46565b3480156103c857600080fd5b5061039a6103d73660046150bd565b610e57565b3480156103e857600080fd5b5061039a6103f73660046152fc565b610e8d565b34801561040857600080fd5b506009545b6040519081526020015b60405180910390f35b34801561042c57600080fd5b5061040d61043b3660046150bd565b600d6020526000908152604090205481565b34801561045957600080fd5b5061040d60055481565b34801561046f57600080fd5b5061039a61047e366004615135565b610f23565b34801561048f57600080fd5b5061039a61131d565b3480156104a457600080fd5b5061039a6104b3366004615338565b611356565b3480156104c457600080fd5b5061039a6104d33660046150bd565b6113da565b3480156104e457600080fd5b506104f86104f33660046150bd565b61140d565b6040516104179190615475565b34801561051157600080fd5b5061039a61164e565b34801561052657600080fd5b50600e546105349060ff1681565b6040519015158152602001610417565b34801561055057600080fd5b5061040d6119a4565b34801561056557600080fd5b5061039a61057436600461518a565b6119d1565b34801561058557600080fd5b506105346105943660046150bd565b6001600160a01b031660009081526001602052604090205460ff1690565b3480156105be57600080fd5b5061039a612004565b3480156105d357600080fd5b5061040d6105e23660046152c5565b612038565b3480156105f357600080fd5b5061039a6106023660046150bd565b612077565b34801561061357600080fd5b5061039a6106223660046152fc565b612322565b34801561063357600080fd5b5061039a61064236600461510b565b6123d7565b34801561065357600080fd5b506002546105349060ff1681565b34801561066d57600080fd5b5061068161067c366004615315565b612762565b60405161041791906154d7565b34801561069a57600080fd5b50600054600160a01b900460ff16610534565b3480156106b957600080fd5b5061039a6106c83660046152fc565b612792565b3480156106d957600080fd5b5061040d60075481565b3480156106ef57600080fd5b5061040d60065481565b34801561070557600080fd5b5061072f6107143660046150bd565b600c602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610417565b34801561075357600080fd5b5061039a612919565b34801561076857600080fd5b5061039a6107773660046150bd565b61294d565b34801561078857600080fd5b5061039a612a6b565b34801561079d57600080fd5b5061039a6107ac3660046150bd565b612aa1565b3480156107bd57600080fd5b506108126107cc3660046152fc565b60046020819052600091825260409091208054600182015460028301546003840154948401546005909401546001600160a01b03909316949193909260ff928316921686565b60405161041796959493929190615423565b34801561083057600080fd5b5061040d60085481565b34801561084657600080fd5b5061039a612c8f565b34801561085b57600080fd5b50600a5461040d565b34801561087057600080fd5b5061053461087f36600461524e565b612cc1565b34801561089057600080fd5b506000546001600160a01b031661072f565b3480156108ae57600080fd5b5061040d612f67565b3480156108c357600080fd5b506105346108d23660046152fc565b600f6020526000908152604090205460ff1681565b3480156108f357600080fd5b5061072f6109023660046152fc565b613066565b34801561091357600080fd5b5061039a6109223660046152fc565b613090565b34801561093357600080fd5b506106816109423660046150bd565b6001600160a01b03166000908152600b602052604090205460ff1690565b34801561096c57600080fd5b5061072f61097b3660046152fc565b6131a2565b34801561098c57600080fd5b5061053461099b3660046150bd565b6131b2565b3480156109ac57600080fd5b5060025461072f9061010090046001600160a01b031681565b3480156109d157600080fd5b5061040d6109e03660046150bd565b6001600160a01b03166000908152600b602052604090206001015490565b348015610a0a57600080fd5b5061039a6131ea565b348015610a1f57600080fd5b5061039a613226565b348015610a3457600080fd5b5061039a610a433660046152e0565b61325f565b348015610a5457600080fd5b5061040d610a633660046152c5565b60036020526000908152604090205481565b348015610a8157600080fd5b50610a95610a903660046150d8565b6133ab565b604051610417919061560f565b348015610aae57600080fd5b5061039a610abd3660046150bd565b613584565b348015610ace57600080fd5b5061039a610add36600461518a565b61361c565b348015610aee57600080fd5b50610b4f610afd3660046150bd565b600b6020526000908152604090208054600182015460028301546003840154600485015460068601546007870154600888015460099098015460ff8816986101009098046001600160a01b031697908a565b6040516104179a999897969594939291906154ea565b336000818152600b6020526040812090815460ff166003811115610b8b57610b8b61581b565b1415610bb25760405162461bcd60e51b8152600401610ba9906155d8565b60405180910390fd5b612710831115610bf75760405162461bcd60e51b815260206004820152601060248201526f496e76616c6964206e6577207261746560801b6044820152606401610ba9565b600881018390558054600982015460408051918252602082018690526001600160a01b0361010090930483169285169160008051602061588483398151915291015b60405180910390a3505050565b600054600160a01b900460ff1615610c705760405162461bcd60e51b8152600401610ba990615579565b33670de0b6b3a7640000821015610cc95760405162461bcd60e51b815260206004820152601860248201527f4d696e696d616c20616d6f756e7420697320312043454c5200000000000000006044820152606401610ba9565b6001600160a01b0383166000908152600b6020526040812090815460ff166003811115610cf857610cf861581b565b1415610d165760405162461bcd60e51b8152600401610ba9906155d8565b6000610d2b8483600101548460020154613850565b6001600160a01b0384166000908152600584016020526040812080549293509183918391610d5a9084906156db565b9250508190555081836002016000828254610d7591906156db565b9250508190555084836001016000828254610d9091906156db565b9091555060029050835460ff166003811115610dae57610dae61581b565b1415610dd9578460076000828254610dc691906156db565b90915550506001830154610dd99061387d565b600254610df69061010090046001600160a01b031685308861397f565b6001830154815460408051928352602083019190915281018690526001600160a01b0380861691908816907f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea906060015b60405180910390a3505050505050565b6000546001600160a01b03163314610e815760405162461bcd60e51b8152600401610ba9906155a3565b610e8a816139f0565b50565b600054600160a01b900460ff16610edd5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610ba9565b6000546001600160a01b03163314610f075760405162461bcd60e51b8152600401610ba9906155a3565b600254610e8a9061010090046001600160a01b03163383613aab565b600054600160a01b900460ff1615610f4d5760405162461bcd60e51b8152600401610ba990615579565b60025460ff1615610fb7573360009081526001602052604090205460ff16610fb75760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f742077686974656c6973746564000000000000006044820152606401610ba9565b336000818152600b6020526040812090815460ff166003811115610fdd57610fdd61581b565b1461102a5760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f7220697320696e697469616c697a656400000000000000006044820152606401610ba9565b6001600160a01b0385166000908152600b602052604081205460ff1660038111156110575761105761581b565b146110a05760405162461bcd60e51b815260206004820152601960248201527829b4b3b732b91034b99037ba3432b9103b30b634b230ba37b960391b6044820152606401610ba9565b6001600160a01b038281166000908152600c602052604090205416156111085760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206973206f74686572207369676e6572000000000000006044820152606401610ba9565b6001600160a01b038581166000908152600c602052604090205416156111665760405162461bcd60e51b815260206004820152601360248201527214da59db995c88185b1c9958591e481d5cd959606a1b6044820152606401610ba9565b6127108311156111b85760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636f6d6d697373696f6e20726174650000000000000000006044820152606401610ba9565b600560005260036020527f405aad32e1adbac89bb7f176e338b8fc6e994ca210c9bb7bdca249b465942250548410156112335760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610ba9565b80546001600160a81b03191660ff196101006001600160a01b03888116918202929092169290921760019081178455600980850188905560088501879055805491820190557f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b03199081169286169283179091556000928352600c602052604090922080549092161790556112d28285610c46565b846001600160a01b0316826001600160a01b0316600080516020615884833981519152868660405161130e929190918252602082015260400190565b60405180910390a35050505050565b6000546001600160a01b031633146113475760405162461bcd60e51b8152600401610ba9906155a3565b600e805460ff19166001179055565b336000818152600b602052604090205460029060ff16600381111561137d5761137d61581b565b146113ca5760405162461bcd60e51b815260206004820181905260248201527f43616c6c6572206973206e6f74206120626f6e6465642076616c696461746f726044820152606401610ba9565b6113d5838284613adb565b505050565b6000546001600160a01b031633146114045760405162461bcd60e51b8152600401610ba9906155a3565b610e8a81613c80565b6009546060906000906001600160401b0381111561142d5761142d61585d565b60405190808252806020026020018201604052801561146657816020015b611453614f60565b81526020019060019003908161144b5790505b5090506000805b60095463ffffffff8216101561157b576000600b600060098463ffffffff168154811061149c5761149c615847565b60009182526020808320909101546001600160a01b0390811684528382019490945260409283018220938a1682526005909301909252902080549091501580156114fa5750600281015463ffffffff80821664010000000090920416145b156115685761153660098363ffffffff168154811061151b5761151b615847565b6000918252602090912001546001600160a01b0316876133ab565b848363ffffffff168151811061154e5761154e615847565b60200260200101819052508280611564906157c4565b9350505b5080611573816157c4565b91505061146d565b5060008163ffffffff166001600160401b0381111561159c5761159c61585d565b6040519080825280602002602001820160405280156115d557816020015b6115c2614f60565b8152602001906001900390816115ba5790505b50905060005b8263ffffffff168163ffffffff16101561164557838163ffffffff168151811061160757611607615847565b6020026020010151828263ffffffff168151811061162757611627615847565b6020026020010181905250808061163d906157c4565b9150506115db565b50949350505050565b336000818152600c60205260409020546001600160a01b0316156116875750336000908152600c60205260409020546001600160a01b03165b6001600160a01b0381166000908152600b602052604090206001815460ff1660038111156116b7576116b761581b565b14806116d857506003815460ff1660038111156116d6576116d661581b565b145b6117245760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076616c696461746f722073746174757300000000000000006044820152606401610ba9565b80600601544310156117715760405162461bcd60e51b8152602060048201526016602482015275109bdb9908189b1bd8dac81b9bdd081c995858da195960521b6044820152606401610ba9565b6008544310156117c35760405162461bcd60e51b815260206004820152601b60248201527f546f6f206672657175656e742076616c696461746f7220626f6e6400000000006044820152606401610ba9565b600760005260036020527ff2c49132ed1cee2a7e75bde50d332a2f81f1d01e5456d8a19d1df09bd561dbd2546117f990436156db565b6008556118068280613d2b565b6118485760405162461bcd60e51b81526020600482015260136024820152724e6f742068617665206d696e20746f6b656e7360681b6044820152606401610ba9565b600360008190526020527fcbc4e5fb02c3d1de23a9f1e014b4d2ee5aeaea9505df5e855c9210bf472495af54600a54811115611887576113d583613df7565b6000196000805b8381101561193c5782600b6000600a84815481106118ae576118ae615847565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154101561192a57809150600b6000600a83815481106118f6576118f6615847565b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015492508261192a5761193c565b80611934816157a9565b91505061188e565b50818460010154116119865760405162461bcd60e51b8152602060048201526013602482015272496e73756666696369656e7420746f6b656e7360681b6044820152606401610ba9565b6119908582613e4b565b61199d846001015461387d565b5050505050565b6000600360075460026119b79190615715565b6119c191906156f3565b6119cc9060016156db565b905090565b600054600160a01b900460ff16156119fb5760405162461bcd60e51b8152600401610ba990615579565b600e5460ff1615611a425760405162461bcd60e51b815260206004820152601160248201527014db185cda081a5cc8191a5cd8589b1959607a1b6044820152606401610ba9565b6000611a8385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613ec692505050565b9050611acb85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061087f9250869150879050615770565b5080606001516001600160401b03164310611b185760405162461bcd60e51b815260206004820152600d60248201526c14db185cda08195e1c1a5c9959609a1b6044820152606401610ba9565b6305f5e10081604001516001600160401b03161115611b705760405162461bcd60e51b815260206004820152601460248201527324b73b30b634b21039b630b9b4103330b1ba37b960611b6044820152606401610ba9565b6020808201516001600160401b03166000908152600f909152604090205460ff1615611bd15760405162461bcd60e51b815260206004820152601060248201526f5573656420736c617368206e6f6e636560801b6044820152606401610ba9565b6020808201516001600160401b03166000908152600f82526040808220805460ff1916600190811790915584516001600160a01b0381168452600b909452912090815460ff166003811115611c2857611c2861581b565b1415611c6c5760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881d5b989bdd5b991959606a1b6044820152606401610ba9565b6000633b9aca0084604001516001600160401b03168360010154611c909190615715565b611c9a91906156f3565b905080826001016000828254611cb09190615734565b9091555060029050825460ff166003811115611cce57611cce61581b565b1415611d4c578060076000828254611ce69190615734565b909155505060808401516001600160401b0316151580611d0d5750611d0b8384613d2b565b155b15611d4c57611d1b83614133565b60808401516001600160401b031615611d4c576080840151611d46906001600160401b0316436156db565b60068301555b60006001600160a01b0316836001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea8460010154600085611d93906157e8565b6040805193845260208401929092529082015260600160405180910390a36000633b9aca0085604001516001600160401b03168460030154611dd59190615715565b611ddf91906156f3565b905080836003016000828254611df59190615734565b90915550611e05905081836156db565b91506000805b8660a0015151811015611f3f5760008760a001518281518110611e3057611e30615847565b60200260200101519050806020015183611e4a91906156db565b81519093506001600160a01b0316611ebd576020810151600254611e7f916101009091046001600160a01b0316903390613aab565b60208082015160405190815233917fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3910160405180910390a2611f2c565b80516020820151600254611ee0926101009091046001600160a01b031691613aab565b80600001516001600160a01b03167fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a38260200151604051611f2391815260200190565b60405180910390a25b5080611f37816157a9565b915050611e0b565b5080831015611f855760405162461bcd60e51b8152602060048201526012602482015271496e76616c696420636f6c6c6563746f727360701b6044820152606401610ba9565b611f8f8184615734565b60066000828254611fa091906156db565b9091555050602080870151604080516001600160401b0390921682529181018590526001600160a01b038716917f10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008910160405180910390a250505050505050505050565b6000546001600160a01b0316331461202e5760405162461bcd60e51b8152600401610ba9906155a3565b612036614290565b565b6000600360008360078111156120505761205061581b565b60078111156120615761206161581b565b8152602001908152602001600020549050919050565b6001600160a01b0381166000908152600b602052604081203391815460ff1660038111156120a7576120a761581b565b14156120c55760405162461bcd60e51b8152600401610ba9906155d8565b6001600160a01b038216600090815260058201602090815260408220600283526003918290527fc3a24b0501bd2c13a7e57f2db4369ec4c223447539fc0724a9d55ac4a06ebd4d54845491939092909160019160ff9091169081111561212d5761212d61581b565b60028501549114915063ffffffff1660005b600285015463ffffffff640100000000909104811690831610156121f157828061218f575063ffffffff82166000908152600180870160205260409091200154439061218c9086906156db565b11155b156121da5763ffffffff821660009081526001860160205260409020546121b690826156db565b63ffffffff83166000908152600180880160205260408220828155015590506121df565b6121f1565b816121e9816157c4565b92505061213f565b60028501805463ffffffff191663ffffffff8416179055806122635760405162461bcd60e51b815260206004820152602560248201527f6e6f20756e64656c65676174696f6e20726561647920746f20626520636f6d706044820152641b195d195960da1b6064820152608401610ba9565b6000612278828860030154896004015461432d565b90508187600401600082825461228e9190615734565b92505081905550808760030160008282546122a99190615734565b90915550506002546122ca9061010090046001600160a01b03168983613aab565b876001600160a01b0316896001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c8360405161230f91815260200190565b60405180910390a3505050505050505050565b600054600160a01b900460ff161561234c5760405162461bcd60e51b8152600401610ba990615579565b6000339050816006600082825461236391906156db565b90915550506002546123859061010090046001600160a01b031682308561397f565b806001600160a01b03167f97e19c4040b6c46d4275e0c4fea68f8f92c81138372ffdb089932c211938f765836006546040516123cb929190918252602082015260400190565b60405180910390a25050565b33670de0b6b3a76400008210156124305760405162461bcd60e51b815260206004820152601960248201527f4d696e696d616c20616d6f756e742069732031207368617265000000000000006044820152606401610ba9565b6001600160a01b0383166000908152600b6020526040812090815460ff16600381111561245f5761245f61581b565b141561247d5760405162461bcd60e51b8152600401610ba9906155d8565b6000612492848360010154846002015461432d565b6001600160a01b03841660009081526005840160205260408120805492935091869183916124c1908490615734565b92505081905550848360020160008282546124dc9190615734565b92505081905550818360010160008282546124f79190615734565b9091555060019050835460ff1660038111156125155761251561581b565b141561257c576002546125379061010090046001600160a01b03168584613aab565b836001600160a01b0316866001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c84604051610e4791815260200190565b6002835460ff1660038111156125945761259461581b565b14156125c95781600760008282546125ac9190615734565b909155506125bc90508685613d2b565b6125c9576125c986614133565b6002810154600a906125ec9063ffffffff8082169164010000000090041661574b565b63ffffffff161061263f5760405162461bcd60e51b815260206004820152601f60248201527f457863656564206d617820756e64656c65676174696f6e20656e7472696573006044820152606401610ba9565b60006126548385600301548660040154613850565b90508084600401600082825461266a91906156db565b925050819055508284600301600082825461268591906156db565b909155505060028201805463ffffffff640100000000918290048116600090815260018087016020526040909120858155439181019190915583549093929004169060046126d2836157c4565b91906101000a81548163ffffffff021916908363ffffffff16021790555050856001600160a01b0316886001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea876001015486600001548861273a906157e8565b6040805193845260208401929092529082015260600160405180910390a35050505050505050565b60008281526004602090815260408083206001600160a01b038516845260060190915290205460ff165b92915050565b336000818152600b6020526040812090815460ff1660038111156127b8576127b861581b565b14156127d65760405162461bcd60e51b8152600401610ba9906155d8565b670de0b6b3a764000083101561282e5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206d696e2073656c662064656c65676174696f6e00000000006044820152606401610ba9565b80600901548310156128d2576002815460ff1660038111156128525761285261581b565b14156128965760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881a5cc8189bdb991959606a1b6044820152606401610ba9565b600660005260036020527fc69056f16cbaa3c616b828e333ab7d3a32310765507f8f58359e99ebb7a885f3546128cc90436156db565b60068201555b60098101839055805460088201546040516001600160a01b0361010090930483169285169160008051602061588483398151915291610c3991888252602082015260400190565b6000546001600160a01b031633146129435760405162461bcd60e51b8152600401610ba9906155a3565b6120366000614346565b6001600160a01b0381166000908152600b602052604090206003815460ff16600381111561297d5761297d61581b565b146129ca5760405162461bcd60e51b815260206004820152601760248201527f56616c696461746f72206e6f7420756e626f6e64696e670000000000000000006044820152606401610ba9565b8060070154431015612a1e5760405162461bcd60e51b815260206004820152601860248201527f556e626f6e6420626c6f636b206e6f74207265616368656400000000000000006044820152606401610ba9565b805460ff191660019081178255600060078301555b6040516001600160a01b038416907fd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a90600090a35050565b6000546001600160a01b03163314612a955760405162461bcd60e51b8152600401610ba9906155a3565b600e805460ff19169055565b336000818152600b6020526040812090815460ff166003811115612ac757612ac761581b565b1415612b155760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206e6f7420696e697469616c697a6564000000000000006044820152606401610ba9565b6001600160a01b038381166000908152600c60205260409020541615612b735760405162461bcd60e51b815260206004820152601360248201527214da59db995c88185b1c9958591e481d5cd959606a1b6044820152606401610ba9565b816001600160a01b0316836001600160a01b031614612c02576001600160a01b0383166000908152600b602052604081205460ff166003811115612bb957612bb961581b565b14612c025760405162461bcd60e51b815260206004820152601960248201527829b4b3b732b91034b99037ba3432b9103b30b634b230ba37b960391b6044820152606401610ba9565b80546001600160a01b036101009182900481166000908152600c602052604080822080546001600160a01b03199081169091558554888516958602610100600160a81b031990911617865584835291819020805493871693909216831790915560098401546008850154915160008051602061588483398151915292610c39928252602082015260400190565b6000546001600160a01b03163314612cb95760405162461bcd60e51b8152600401610ba9906155a3565b612036614396565b600080612d2284805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050600083516001600160401b03811115612d3f57612d3f61585d565b604051908082528060200260200182016040528015612d68578160200160208202803683370190505b509050600080805b8651811015612f0d57612da5878281518110612d8e57612d8e615847565b6020026020010151866143fb90919063ffffffff16565b848281518110612db757612db7615847565b60200260200101906001600160a01b031690816001600160a01b031681525050816001600160a01b0316848281518110612df357612df3615847565b60200260200101516001600160a01b031611612e515760405162461bcd60e51b815260206004820152601e60248201527f5369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610ba9565b838181518110612e6357612e63615847565b602002602001015191506000600b6000600c6000888681518110612e8957612e89615847565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081205490931684528301939093529101902090506002815460ff166003811115612edd57612edd61581b565b14612ee85750612efb565b6001810154612ef790856156db565b9350505b80612f05816157a9565b915050612d70565b50612f166119a4565b821015612f5a5760405162461bcd60e51b8152602060048201526012602482015271145d5bdc9d5b481b9bdd081c995858da195960721b6044820152606401610ba9565b5060019695505050505050565b600080600b6000600a600081548110612f8257612f82615847565b60009182526020808320909101546001600160a01b03168352820192909252604001902060019081015491505b600a548110156130605781600b6000600a8481548110612fd157612fd1615847565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154101561304e57600b6000600a838154811061301657613016615847565b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015491508161304e5760009250505090565b80613058816157a9565b915050612faf565b50919050565b6009818154811061307657600080fd5b6000918252602090912001546001600160a01b0316905081565b6000805b600a5463ffffffff821610156131585760016130dd84600a8463ffffffff16815481106130c3576130c3615847565b6000918252602090912001546001600160a01b0316612762565b60038111156130ee576130ee61581b565b141561314657600b6000600a8363ffffffff168154811061311157613111615847565b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015461314390836156db565b91505b80613150816157c4565b915050613094565b5060006131636119a4565b8210159050806131985760008381526004602052604081206001015460068054919290916131929084906156db565b90915550505b6113d5838261449f565b600a818154811061307657600080fd5b600060026001600160a01b0383166000908152600b602052604090205460ff1660038111156131e3576131e361581b565b1492915050565b6000546001600160a01b031633146132145760405162461bcd60e51b8152600401610ba9906155a3565b6120366002805460ff19166001179055565b6000546001600160a01b031633146132505760405162461bcd60e51b8152600401610ba9906155a3565b6120366002805460ff19169055565b60055460008181526004602052604090209061327c9060016156db565b60055560036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff5481546001600160a01b03191633908117835560018084018390556000527fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c549091906132f290436156db565b600284015560038301805486919060ff191660018360078111156133185761331861581b565b02179055506004830184905560058301805460ff19166001179055600254613351906001600160a01b036101009091041683308461397f565b7f4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f16060016005546133819190615734565b83838660020154898960405161339c9695949392919061564f565b60405180910390a15050505050565b6133b3614f60565b6001600160a01b038084166000908152600b6020908152604080832093861683526005840190915281208054600184015460028501549293926133f792919061432d565b6002830154909150600090819061341f9063ffffffff8082169164010000000090041661574b565b63ffffffff1690506000816001600160401b038111156134415761344161585d565b60405190808252806020026020018201604052801561348657816020015b604080518082019091526000808252602082015281526020019060019003908161345f5790505b50905060005b8281101561352e57600286015460018701906000906134b19063ffffffff16846156db565b8152602001908152602001600020604051806040016040529081600082015481526020016001820154815250508282815181106134f0576134f0615847565b602090810291909101810191909152600082815260018801909152604090205461351a90856156db565b935080613526816157a9565b91505061348c565b506000613544848860030154896004015461432d565b6040805160a0810182526001600160a01b03909c168c5260208c01969096529554948a019490945260608901949094525050506080850152509192915050565b6000546001600160a01b031633146135ae5760405162461bcd60e51b8152600401610ba9906155a3565b6001600160a01b0381166136135760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610ba9565b610e8a81614346565b600054600160a01b900460ff16156136465760405162461bcd60e51b8152600401610ba990615579565b61368c84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061087f9250859150869050615770565b5060006136ce85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061463292505050565b80516001600160a01b03166000908152600d602090815260408220549083015192935090916136fd9190615734565b90506000811161373f5760405162461bcd60e51b815260206004820152600d60248201526c139bc81b995dc81c995dd85c99609a1b6044820152606401610ba9565b8060065410156137a05760405162461bcd60e51b815260206004820152602660248201527f52657761726420706f6f6c20697320736d616c6c6572207468616e206e6577206044820152651c995dd85c9960d21b6064820152608401610ba9565b60208083015183516001600160a01b03166000908152600d9092526040822055600680548392906137d2908490615734565b909155505081516002546137f6916101009091046001600160a01b03169083613aab565b81600001516001600160a01b03167ff01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e273174382600654604051613840929190918252602082015260400190565b60405180910390a2505050505050565b60008261385e575082613876565b826138698386615715565b61387391906156f3565b90505b9392505050565b600a54600281148061388f5750806003145b156139055761389c6119a4565b82106139015760405162461bcd60e51b815260206004820152602e60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201526d2071756f72756d20746f6b656e7360901b6064820152608401610ba9565b5050565b600381111561390157600360075461391d91906156f3565b82106139015760405162461bcd60e51b815260206004820152602b60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201526a20312f3320746f6b656e7360a81b6064820152608401610ba9565b6040516001600160a01b03808516602483015283166044820152606481018290526139ea9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526146d9565b50505050565b6001600160a01b03811660009081526001602052604090205460ff1615613a4f5760405162461bcd60e51b8152602060048201526013602482015272185b1c9958591e481dda1a5d195b1a5cdd1959606a1b6044820152606401610ba9565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f91015b60405180910390a150565b6040516001600160a01b0383166024820152604481018290526113d590849063a9059cbb60e01b906064016139b3565b60008381526004602052604090206001600582015460ff166002811115613b0457613b0461581b565b14613b4b5760405162461bcd60e51b8152602060048201526017602482015276496e76616c69642070726f706f73616c2073746174757360481b6044820152606401610ba9565b80600201544310613b955760405162461bcd60e51b8152602060048201526014602482015273159bdd1948191958591b1a5b99481c185cdcd95960621b6044820152606401610ba9565b6001600160a01b038316600090815260068201602052604081205460ff166003811115613bc457613bc461581b565b14613c035760405162461bcd60e51b815260206004820152600f60248201526e159bdd195c881a185cc81d9bdd1959608a1b6044820152606401610ba9565b6001600160a01b03831660009081526006820160205260409020805483919060ff19166001836003811115613c3a57613c3a61581b565b02179055507f06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65848484604051613c7293929190615622565b60405180910390a150505050565b6001600160a01b03811660009081526001602052604090205460ff16613cda5760405162461bcd60e51b815260206004820152600f60248201526e1b9bdd081dda1a5d195b1a5cdd1959608a1b6044820152606401610ba9565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b69101613aa0565b6001600160a01b0382166000908152600b60209081526040822060018101546004845260039092527f83ec6a1f0257b830b5e016457c9cf1435391bf56cc98f369a58a54fe9377246554909190811015613d8a5760009250505061278c565b836001600160a01b0316856001600160a01b03161415613dec576001600160a01b03851660009081526005830160205260408120546002840154613dd09190849061432d565b90508260090154811015613dea576000935050505061278c565b505b506001949350505050565b600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180546001600160a01b0319166001600160a01b038316179055610e8a816147ab565b613e7b600a8281548110613e6157613e61615847565b6000918252602090912001546001600160a01b03166147f9565b81600a8281548110613e8f57613e8f615847565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550613901826147ab565b6040805160c0810182526000808252602080830182905282840182905260608084018390526080840183905260a084015283518085019094528184528301849052909190613f1582600661487f565b905080600681518110613f2a57613f2a615847565b60200260200101516001600160401b03811115613f4957613f4961585d565b604051908082528060200260200182016040528015613f8e57816020015b6040805180820190915260008082526020820152815260200190600190039081613f675790505b508360a00181905250600081600681518110613fac57613fac615847565b6020026020010181815250506000805b6020840151518451101561412a57613fd384614938565b9092509050816001141561400257613ff2613fed85614972565b614a2e565b6001600160a01b03168552613fbc565b81600214156140275761401484614a39565b6001600160401b03166020860152613fbc565b816003141561404c5761403984614a39565b6001600160401b03166040860152613fbc565b81600414156140715761405e84614a39565b6001600160401b03166060860152613fbc565b81600514156140965761408384614a39565b6001600160401b03166080860152613fbc565b816006141561411b576140b06140ab85614972565b614abb565b8560a00151846006815181106140c8576140c8615847565b6020026020010151815181106140e0576140e0615847565b6020026020010181905250826006815181106140fe576140fe615847565b602002602001018051809190614113906157a9565b905250613fbc565b6141258482614b55565b613fbc565b50505050919050565b600a5460009061414590600190615734565b905060005b600a5481101561425057826001600160a01b0316600a828154811061417157614171615847565b6000918252602090912001546001600160a01b0316141561423e578181101561420257600a82815481106141a7576141a7615847565b600091825260209091200154600a80546001600160a01b0390921691839081106141d3576141d3615847565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b600a80548061421357614213615831565b600082815260209020810160001990810180546001600160a01b03191690550190556113d5836147f9565b80614248816157a9565b91505061414a565b5060405162461bcd60e51b81526020600482015260146024820152732737ba103137b73232b2103b30b634b230ba37b960611b6044820152606401610ba9565b600054600160a01b900460ff166142e05760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610ba9565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60008161433b575082613876565b816138698486615715565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff16156143c05760405162461bcd60e51b8152600401610ba990615579565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586143103390565b600081516041141561442f5760208201516040830151606084015160001a61442586828585614bc7565b935050505061278c565b815160401415614457576020820151604083015161444e858383614d70565b9250505061278c565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610ba9565b60008281526004602052604090206001600582015460ff1660028111156144c8576144c861581b565b1461450f5760405162461bcd60e51b8152602060048201526017602482015276496e76616c69642070726f706f73616c2073746174757360481b6044820152606401610ba9565b80600201544310156145635760405162461bcd60e51b815260206004820152601960248201527f566f746520646561646c696e65206e6f742072656163686564000000000000006044820152606401610ba9565b60058101805460ff1916600217905581156145e4578054600182015460025461459f926001600160a01b03610100909204821692911690613aab565b600481015460038083015460009060ff1660078111156145c1576145c161581b565b60078111156145d2576145d261581b565b81526020810191909152604001600020555b600381015460048201546040517fd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a992614625928792879260ff169190615680565b60405180910390a1505050565b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b602083015151835110156146d15761467483614938565b9092509050816001141561469e5761468e613fed84614972565b6001600160a01b0316845261465d565b81600214156146c2576146b86146b384614972565b614d9a565b602085015261465d565b6146cc8382614b55565b61465d565b505050919050565b600061472e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614dd19092919063ffffffff16565b8051909150156113d5578080602001905181019061474c9190615168565b6113d55760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610ba9565b6001600160a01b0381166000908152600b60205260408120805460ff19166002178155600780820183905560018201548154929390926147ec9084906156db565b9091555060029050612a33565b6001600160a01b0381166000908152600b602090815260408220805460ff19166003908117825560029093529190527fc3a24b0501bd2c13a7e57f2db4369ec4c223447539fc0724a9d55ac4a06ebd4d5461485490436156db565b81600701819055508060010154600760008282546148729190615734565b9091555060039050612a33565b815160609061488f8360016156db565b6001600160401b038111156148a6576148a661585d565b6040519080825280602002602001820160405280156148cf578160200160208202803683370190505b5091506000805b6020860151518651101561492f576148ed86614938565b8092508193505050600184838151811061490957614909615847565b6020026020010181815161491d91906156db565b90525061492a8682614b55565b6148d6565b50509092525090565b600080600061494684614a39565b90506149536008826156f3565b925080600716600581111561496a5761496a61581b565b915050915091565b6060600061497f83614a39565b9050600081846000015161499391906156db565b90508360200151518111156149a757600080fd5b816001600160401b038111156149bf576149bf61585d565b6040519080825280601f01601f1916602001820160405280156149e9576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015614a23578181015183820152614a1c6020826156db565b9050614a01565b505050935250919050565b600061278c82614de0565b602080820151825181019091015160009182805b600a811015614ab55783811a9150614a66816007615715565b82607f16901b851794508160801660001415614aa357614a878160016156db565b86518790614a969083906156db565b9052509395945050505050565b80614aad816157a9565b915050614a4d565b50600080fd5b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b602083015151835110156146d157614afd83614938565b90925090508160011415614b2757614b17613fed84614972565b6001600160a01b03168452614ae6565b8160021415614b4657614b3c6146b384614972565b6020850152614ae6565b614b508382614b55565b614ae6565b6000816005811115614b6957614b6961581b565b1415614b78576113d582614a39565b6002816005811115614b8c57614b8c61581b565b1415610375576000614b9d83614a39565b90508083600001818151614bb191906156db565b905250602083015151835111156113d557600080fd5b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115614c445760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610ba9565b8360ff16601b1480614c5957508360ff16601c145b614cb05760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610ba9565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015614d04573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614d675760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610ba9565b95945050505050565b60006001600160ff1b03821660ff83901c601b01614d9086828785614bc7565b9695505050505050565b6000602082511115614dab57600080fd5b6020820151905081516020614dc09190615734565b614dcb906008615715565b1c919050565b60606138738484600085614dff565b60008151601414614df057600080fd5b5060200151600160601b900490565b606082471015614e605760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610ba9565b843b614eae5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610ba9565b600080866001600160a01b03168587604051614eca9190615407565b60006040518083038185875af1925050503d8060008114614f07576040519150601f19603f3d011682016040523d82523d6000602084013e614f0c565b606091505b5091509150614f1c828286614f27565b979650505050505050565b60608315614f36575081613876565b825115614f465782518084602001fd5b8160405162461bcd60e51b8152600401610ba99190615546565b6040518060a0016040528060006001600160a01b03168152602001600081526020016000815260200160008152602001606081525090565b60006001600160401b0380841115614fb257614fb261585d565b8360051b6020614fc38183016156ab565b868152935080840185838101891015614fdb57600080fd5b60009350835b8881101561501657813586811115614ff7578586fd5b6150038b828b0161503f565b8452509183019190830190600101614fe1565b5050505050509392505050565b80356001600160a01b038116811461503a57600080fd5b919050565b600082601f83011261505057600080fd5b81356001600160401b038111156150695761506961585d565b61507c601f8201601f19166020016156ab565b81815284602083860101111561509157600080fd5b816020850160208301376000918101602001919091529392505050565b80356008811061503a57600080fd5b6000602082840312156150cf57600080fd5b61387682615023565b600080604083850312156150eb57600080fd5b6150f483615023565b915061510260208401615023565b90509250929050565b6000806040838503121561511e57600080fd5b61512783615023565b946020939093013593505050565b60008060006060848603121561514a57600080fd5b61515384615023565b95602085013595506040909401359392505050565b60006020828403121561517a57600080fd5b8151801515811461387657600080fd5b600080600080604085870312156151a057600080fd5b84356001600160401b03808211156151b757600080fd5b818701915087601f8301126151cb57600080fd5b8135818111156151da57600080fd5b8860208285010111156151ec57600080fd5b60209283019650945090860135908082111561520757600080fd5b818701915087601f83011261521b57600080fd5b81358181111561522a57600080fd5b8860208260051b850101111561523f57600080fd5b95989497505060200194505050565b6000806040838503121561526157600080fd5b82356001600160401b038082111561527857600080fd5b6152848683870161503f565b9350602085013591508082111561529a57600080fd5b508301601f810185136152ac57600080fd5b6152bb85823560208401614f98565b9150509250929050565b6000602082840312156152d757600080fd5b613876826150ae565b600080604083850312156152f357600080fd5b615127836150ae565b60006020828403121561530e57600080fd5b5035919050565b6000806040838503121561532857600080fd5b8235915061510260208401615023565b6000806040838503121561534b57600080fd5b8235915060208301356004811061536157600080fd5b809150509250929050565b6008811061537c5761537c61581b565b9052565b80516001600160a01b0316825260208082015181840152604080830151818501526060808401519085015260808084015160a09186018290528051918601829052600093908101929091849060c08801905b808310156153fb57855180518352850151858301529484019460019290920191908301906153d2565b50979650505050505050565b6000825161541981846020870161577d565b9190910192915050565b6001600160a01b0387168152602081018690526040810185905260c0810161544e606083018661536c565b836080830152600383106154645761546461581b565b8260a0830152979650505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156154ca57603f198886030184526154b8858351615380565b9450928501929085019060010161549c565b5092979650505050505050565b602081016154e483615873565b91905290565b61014081016154f88c615873565b9a81526001600160a01b039990991660208a015260408901979097526060880195909552608087019390935260a086019190915260c085015260e08401526101008301526101209091015290565b602081526000825180602084015261556581604085016020870161577d565b601f01601f19169190910160400192915050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601c908201527f56616c696461746f72206973206e6f7420696e697469616c697a656400000000604082015260600190565b6020815260006138766020830184615380565b8381526001600160a01b03831660208201526060810161564183615873565b826040830152949350505050565b8681526001600160a01b0386166020820152604081018590526060810184905260c08101615464608083018561536c565b84815283151560208201526080810161569c604083018561536c565b82606083015295945050505050565b604051601f8201601f191681016001600160401b03811182821017156156d3576156d361585d565b604052919050565b600082198211156156ee576156ee615805565b500190565b60008261571057634e487b7160e01b600052601260045260246000fd5b500490565b600081600019048311821515161561572f5761572f615805565b500290565b60008282101561574657615746615805565b500390565b600063ffffffff8381169083168181101561576857615768615805565b039392505050565b6000613876368484614f98565b60005b83811015615798578181015183820152602001615780565b838111156139ea5750506000910152565b60006000198214156157bd576157bd615805565b5060010190565b600063ffffffff808316818114156157de576157de615805565b6001019392505050565b6000600160ff1b8214156157fe576157fe615805565b5060000390565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60048110610e8a57610e8a61581b56fe848523f3f4f7d056ddf7150c34e51564204e16c45e8b98cae2263d507a875346a264697066735822122068938708628060622a88eb6b65d3b5fa2d58147f3aa484150e5fa3860c5fd51364736f6c63430008070033",
}

// DPoSABI is the input ABI used to generate the binding from.
// Deprecated: Use DPoSMetaData.ABI instead.
var DPoSABI = DPoSMetaData.ABI

// DPoSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DPoSMetaData.Bin instead.
var DPoSBin = DPoSMetaData.Bin

// DeployDPoS deploys a new Ethereum contract, binding an instance of DPoS to it.
func DeployDPoS(auth *bind.TransactOpts, backend bind.ContractBackend, _celerTokenAddress common.Address, _governProposalDeposit *big.Int, _governVoteTimeout *big.Int, _slashTimeout *big.Int, _maxBondedValidators *big.Int, _minValidatorTokens *big.Int, _minSelfDelegation *big.Int, _advanceNoticePeriod *big.Int, _validatorBondInterval *big.Int) (common.Address, *types.Transaction, *DPoS, error) {
	parsed, err := DPoSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DPoSBin), backend, _celerTokenAddress, _governProposalDeposit, _governVoteTimeout, _slashTimeout, _maxBondedValidators, _minValidatorTokens, _minSelfDelegation, _advanceNoticePeriod, _validatorBondInterval)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DPoS{DPoSCaller: DPoSCaller{contract: contract}, DPoSTransactor: DPoSTransactor{contract: contract}, DPoSFilterer: DPoSFilterer{contract: contract}}, nil
}

// DPoS is an auto generated Go binding around an Ethereum contract.
type DPoS struct {
	DPoSCaller     // Read-only binding to the contract
	DPoSTransactor // Write-only binding to the contract
	DPoSFilterer   // Log filterer for contract events
}

// DPoSCaller is an auto generated read-only Go binding around an Ethereum contract.
type DPoSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DPoSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DPoSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DPoSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DPoSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DPoSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DPoSSession struct {
	Contract     *DPoS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DPoSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DPoSCallerSession struct {
	Contract *DPoSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DPoSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DPoSTransactorSession struct {
	Contract     *DPoSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DPoSRaw is an auto generated low-level Go binding around an Ethereum contract.
type DPoSRaw struct {
	Contract *DPoS // Generic contract binding to access the raw methods on
}

// DPoSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DPoSCallerRaw struct {
	Contract *DPoSCaller // Generic read-only contract binding to access the raw methods on
}

// DPoSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DPoSTransactorRaw struct {
	Contract *DPoSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDPoS creates a new instance of DPoS, bound to a specific deployed contract.
func NewDPoS(address common.Address, backend bind.ContractBackend) (*DPoS, error) {
	contract, err := bindDPoS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DPoS{DPoSCaller: DPoSCaller{contract: contract}, DPoSTransactor: DPoSTransactor{contract: contract}, DPoSFilterer: DPoSFilterer{contract: contract}}, nil
}

// NewDPoSCaller creates a new read-only instance of DPoS, bound to a specific deployed contract.
func NewDPoSCaller(address common.Address, caller bind.ContractCaller) (*DPoSCaller, error) {
	contract, err := bindDPoS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DPoSCaller{contract: contract}, nil
}

// NewDPoSTransactor creates a new write-only instance of DPoS, bound to a specific deployed contract.
func NewDPoSTransactor(address common.Address, transactor bind.ContractTransactor) (*DPoSTransactor, error) {
	contract, err := bindDPoS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DPoSTransactor{contract: contract}, nil
}

// NewDPoSFilterer creates a new log filterer instance of DPoS, bound to a specific deployed contract.
func NewDPoSFilterer(address common.Address, filterer bind.ContractFilterer) (*DPoSFilterer, error) {
	contract, err := bindDPoS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DPoSFilterer{contract: contract}, nil
}

// bindDPoS binds a generic wrapper to an already deployed contract.
func bindDPoS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DPoSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DPoS *DPoSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DPoS.Contract.DPoSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DPoS *DPoSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.Contract.DPoSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DPoS *DPoSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DPoS.Contract.DPoSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DPoS *DPoSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DPoS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DPoS *DPoSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DPoS *DPoSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DPoS.Contract.contract.Transact(opts, method, params...)
}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_DPoS *DPoSCaller) BondedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "bondedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_DPoS *DPoSSession) BondedTokens() (*big.Int, error) {
	return _DPoS.Contract.BondedTokens(&_DPoS.CallOpts)
}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_DPoS *DPoSCallerSession) BondedTokens() (*big.Int, error) {
	return _DPoS.Contract.BondedTokens(&_DPoS.CallOpts)
}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_DPoS *DPoSCaller) BondedValAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "bondedValAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_DPoS *DPoSSession) BondedValAddrs(arg0 *big.Int) (common.Address, error) {
	return _DPoS.Contract.BondedValAddrs(&_DPoS.CallOpts, arg0)
}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_DPoS *DPoSCallerSession) BondedValAddrs(arg0 *big.Int) (common.Address, error) {
	return _DPoS.Contract.BondedValAddrs(&_DPoS.CallOpts, arg0)
}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_DPoS *DPoSCaller) CelerToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "celerToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_DPoS *DPoSSession) CelerToken() (common.Address, error) {
	return _DPoS.Contract.CelerToken(&_DPoS.CallOpts)
}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_DPoS *DPoSCallerSession) CelerToken() (common.Address, error) {
	return _DPoS.Contract.CelerToken(&_DPoS.CallOpts)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x1e6f3d8a.
//
// Solidity: function claimedReward(address ) view returns(uint256)
func (_DPoS *DPoSCaller) ClaimedReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "claimedReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedReward is a free data retrieval call binding the contract method 0x1e6f3d8a.
//
// Solidity: function claimedReward(address ) view returns(uint256)
func (_DPoS *DPoSSession) ClaimedReward(arg0 common.Address) (*big.Int, error) {
	return _DPoS.Contract.ClaimedReward(&_DPoS.CallOpts, arg0)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x1e6f3d8a.
//
// Solidity: function claimedReward(address ) view returns(uint256)
func (_DPoS *DPoSCallerSession) ClaimedReward(arg0 common.Address) (*big.Int, error) {
	return _DPoS.Contract.ClaimedReward(&_DPoS.CallOpts, arg0)
}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_DPoS *DPoSCaller) GetBondedValidatorNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getBondedValidatorNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_DPoS *DPoSSession) GetBondedValidatorNum() (*big.Int, error) {
	return _DPoS.Contract.GetBondedValidatorNum(&_DPoS.CallOpts)
}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_DPoS *DPoSCallerSession) GetBondedValidatorNum() (*big.Int, error) {
	return _DPoS.Contract.GetBondedValidatorNum(&_DPoS.CallOpts)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
func (_DPoS *DPoSCaller) GetDelegatorInfo(opts *bind.CallOpts, _valAddr common.Address, _delAddr common.Address) (DPoSDelegatorInfo, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getDelegatorInfo", _valAddr, _delAddr)

	if err != nil {
		return *new(DPoSDelegatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DPoSDelegatorInfo)).(*DPoSDelegatorInfo)

	return out0, err

}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
func (_DPoS *DPoSSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (DPoSDelegatorInfo, error) {
	return _DPoS.Contract.GetDelegatorInfo(&_DPoS.CallOpts, _valAddr, _delAddr)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
func (_DPoS *DPoSCallerSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (DPoSDelegatorInfo, error) {
	return _DPoS.Contract.GetDelegatorInfo(&_DPoS.CallOpts, _valAddr, _delAddr)
}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[])[])
func (_DPoS *DPoSCaller) GetDelegatorInfos(opts *bind.CallOpts, _delAddr common.Address) ([]DPoSDelegatorInfo, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getDelegatorInfos", _delAddr)

	if err != nil {
		return *new([]DPoSDelegatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DPoSDelegatorInfo)).(*[]DPoSDelegatorInfo)

	return out0, err

}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[])[])
func (_DPoS *DPoSSession) GetDelegatorInfos(_delAddr common.Address) ([]DPoSDelegatorInfo, error) {
	return _DPoS.Contract.GetDelegatorInfos(&_DPoS.CallOpts, _delAddr)
}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[])[])
func (_DPoS *DPoSCallerSession) GetDelegatorInfos(_delAddr common.Address) ([]DPoSDelegatorInfo, error) {
	return _DPoS.Contract.GetDelegatorInfos(&_DPoS.CallOpts, _delAddr)
}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_DPoS *DPoSCaller) GetMinValidatorTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getMinValidatorTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_DPoS *DPoSSession) GetMinValidatorTokens() (*big.Int, error) {
	return _DPoS.Contract.GetMinValidatorTokens(&_DPoS.CallOpts)
}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_DPoS *DPoSCallerSession) GetMinValidatorTokens() (*big.Int, error) {
	return _DPoS.Contract.GetMinValidatorTokens(&_DPoS.CallOpts)
}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_DPoS *DPoSCaller) GetParamProposalVote(opts *bind.CallOpts, _proposalId *big.Int, _voter common.Address) (uint8, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getParamProposalVote", _proposalId, _voter)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_DPoS *DPoSSession) GetParamProposalVote(_proposalId *big.Int, _voter common.Address) (uint8, error) {
	return _DPoS.Contract.GetParamProposalVote(&_DPoS.CallOpts, _proposalId, _voter)
}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_DPoS *DPoSCallerSession) GetParamProposalVote(_proposalId *big.Int, _voter common.Address) (uint8, error) {
	return _DPoS.Contract.GetParamProposalVote(&_DPoS.CallOpts, _proposalId, _voter)
}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_DPoS *DPoSCaller) GetParamValue(opts *bind.CallOpts, _name uint8) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getParamValue", _name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_DPoS *DPoSSession) GetParamValue(_name uint8) (*big.Int, error) {
	return _DPoS.Contract.GetParamValue(&_DPoS.CallOpts, _name)
}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_DPoS *DPoSCallerSession) GetParamValue(_name uint8) (*big.Int, error) {
	return _DPoS.Contract.GetParamValue(&_DPoS.CallOpts, _name)
}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_DPoS *DPoSCaller) GetQuorumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getQuorumTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_DPoS *DPoSSession) GetQuorumTokens() (*big.Int, error) {
	return _DPoS.Contract.GetQuorumTokens(&_DPoS.CallOpts)
}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_DPoS *DPoSCallerSession) GetQuorumTokens() (*big.Int, error) {
	return _DPoS.Contract.GetQuorumTokens(&_DPoS.CallOpts)
}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_DPoS *DPoSCaller) GetValidatorNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getValidatorNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_DPoS *DPoSSession) GetValidatorNum() (*big.Int, error) {
	return _DPoS.Contract.GetValidatorNum(&_DPoS.CallOpts)
}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_DPoS *DPoSCallerSession) GetValidatorNum() (*big.Int, error) {
	return _DPoS.Contract.GetValidatorNum(&_DPoS.CallOpts)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_DPoS *DPoSCaller) GetValidatorStatus(opts *bind.CallOpts, _valAddr common.Address) (uint8, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getValidatorStatus", _valAddr)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_DPoS *DPoSSession) GetValidatorStatus(_valAddr common.Address) (uint8, error) {
	return _DPoS.Contract.GetValidatorStatus(&_DPoS.CallOpts, _valAddr)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_DPoS *DPoSCallerSession) GetValidatorStatus(_valAddr common.Address) (uint8, error) {
	return _DPoS.Contract.GetValidatorStatus(&_DPoS.CallOpts, _valAddr)
}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_DPoS *DPoSCaller) GetValidatorTokens(opts *bind.CallOpts, _valAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getValidatorTokens", _valAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_DPoS *DPoSSession) GetValidatorTokens(_valAddr common.Address) (*big.Int, error) {
	return _DPoS.Contract.GetValidatorTokens(&_DPoS.CallOpts, _valAddr)
}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_DPoS *DPoSCallerSession) GetValidatorTokens(_valAddr common.Address) (*big.Int, error) {
	return _DPoS.Contract.GetValidatorTokens(&_DPoS.CallOpts, _valAddr)
}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_DPoS *DPoSCaller) IsBondedValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "isBondedValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_DPoS *DPoSSession) IsBondedValidator(_addr common.Address) (bool, error) {
	return _DPoS.Contract.IsBondedValidator(&_DPoS.CallOpts, _addr)
}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_DPoS *DPoSCallerSession) IsBondedValidator(_addr common.Address) (bool, error) {
	return _DPoS.Contract.IsBondedValidator(&_DPoS.CallOpts, _addr)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_DPoS *DPoSCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_DPoS *DPoSSession) IsWhitelisted(account common.Address) (bool, error) {
	return _DPoS.Contract.IsWhitelisted(&_DPoS.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_DPoS *DPoSCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _DPoS.Contract.IsWhitelisted(&_DPoS.CallOpts, account)
}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_DPoS *DPoSCaller) NextBondBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "nextBondBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_DPoS *DPoSSession) NextBondBlock() (*big.Int, error) {
	return _DPoS.Contract.NextBondBlock(&_DPoS.CallOpts)
}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_DPoS *DPoSCallerSession) NextBondBlock() (*big.Int, error) {
	return _DPoS.Contract.NextBondBlock(&_DPoS.CallOpts)
}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_DPoS *DPoSCaller) NextParamProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "nextParamProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_DPoS *DPoSSession) NextParamProposalId() (*big.Int, error) {
	return _DPoS.Contract.NextParamProposalId(&_DPoS.CallOpts)
}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_DPoS *DPoSCallerSession) NextParamProposalId() (*big.Int, error) {
	return _DPoS.Contract.NextParamProposalId(&_DPoS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DPoS *DPoSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DPoS *DPoSSession) Owner() (common.Address, error) {
	return _DPoS.Contract.Owner(&_DPoS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DPoS *DPoSCallerSession) Owner() (common.Address, error) {
	return _DPoS.Contract.Owner(&_DPoS.CallOpts)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_DPoS *DPoSCaller) ParamProposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "paramProposals", arg0)

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
func (_DPoS *DPoSSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _DPoS.Contract.ParamProposals(&_DPoS.CallOpts, arg0)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_DPoS *DPoSCallerSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _DPoS.Contract.ParamProposals(&_DPoS.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_DPoS *DPoSCaller) Params(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "params", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_DPoS *DPoSSession) Params(arg0 uint8) (*big.Int, error) {
	return _DPoS.Contract.Params(&_DPoS.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_DPoS *DPoSCallerSession) Params(arg0 uint8) (*big.Int, error) {
	return _DPoS.Contract.Params(&_DPoS.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DPoS *DPoSCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DPoS *DPoSSession) Paused() (bool, error) {
	return _DPoS.Contract.Paused(&_DPoS.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DPoS *DPoSCallerSession) Paused() (bool, error) {
	return _DPoS.Contract.Paused(&_DPoS.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_DPoS *DPoSCaller) RewardPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_DPoS *DPoSSession) RewardPool() (*big.Int, error) {
	return _DPoS.Contract.RewardPool(&_DPoS.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_DPoS *DPoSCallerSession) RewardPool() (*big.Int, error) {
	return _DPoS.Contract.RewardPool(&_DPoS.CallOpts)
}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_DPoS *DPoSCaller) SignerVals(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "signerVals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_DPoS *DPoSSession) SignerVals(arg0 common.Address) (common.Address, error) {
	return _DPoS.Contract.SignerVals(&_DPoS.CallOpts, arg0)
}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_DPoS *DPoSCallerSession) SignerVals(arg0 common.Address) (common.Address, error) {
	return _DPoS.Contract.SignerVals(&_DPoS.CallOpts, arg0)
}

// SlashDisabled is a free data retrieval call binding the contract method 0x3773e489.
//
// Solidity: function slashDisabled() view returns(bool)
func (_DPoS *DPoSCaller) SlashDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "slashDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlashDisabled is a free data retrieval call binding the contract method 0x3773e489.
//
// Solidity: function slashDisabled() view returns(bool)
func (_DPoS *DPoSSession) SlashDisabled() (bool, error) {
	return _DPoS.Contract.SlashDisabled(&_DPoS.CallOpts)
}

// SlashDisabled is a free data retrieval call binding the contract method 0x3773e489.
//
// Solidity: function slashDisabled() view returns(bool)
func (_DPoS *DPoSCallerSession) SlashDisabled() (bool, error) {
	return _DPoS.Contract.SlashDisabled(&_DPoS.CallOpts)
}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_DPoS *DPoSCaller) SlashNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "slashNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_DPoS *DPoSSession) SlashNonces(arg0 *big.Int) (bool, error) {
	return _DPoS.Contract.SlashNonces(&_DPoS.CallOpts, arg0)
}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_DPoS *DPoSCallerSession) SlashNonces(arg0 *big.Int) (bool, error) {
	return _DPoS.Contract.SlashNonces(&_DPoS.CallOpts, arg0)
}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_DPoS *DPoSCaller) ValAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "valAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_DPoS *DPoSSession) ValAddrs(arg0 *big.Int) (common.Address, error) {
	return _DPoS.Contract.ValAddrs(&_DPoS.CallOpts, arg0)
}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_DPoS *DPoSCallerSession) ValAddrs(arg0 *big.Int) (common.Address, error) {
	return _DPoS.Contract.ValAddrs(&_DPoS.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 bondBlock, uint256 unbondBlock, uint256 commissionRate, uint256 minSelfDelegation)
func (_DPoS *DPoSCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	BondBlock          *big.Int
	UnbondBlock        *big.Int
	CommissionRate     *big.Int
	MinSelfDelegation  *big.Int
}, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Status             uint8
		Signer             common.Address
		Tokens             *big.Int
		Shares             *big.Int
		UndelegationTokens *big.Int
		UndelegationShares *big.Int
		BondBlock          *big.Int
		UnbondBlock        *big.Int
		CommissionRate     *big.Int
		MinSelfDelegation  *big.Int
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
	outstruct.BondBlock = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.UnbondBlock = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.CommissionRate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.MinSelfDelegation = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 bondBlock, uint256 unbondBlock, uint256 commissionRate, uint256 minSelfDelegation)
func (_DPoS *DPoSSession) Validators(arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	BondBlock          *big.Int
	UnbondBlock        *big.Int
	CommissionRate     *big.Int
	MinSelfDelegation  *big.Int
}, error) {
	return _DPoS.Contract.Validators(&_DPoS.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 bondBlock, uint256 unbondBlock, uint256 commissionRate, uint256 minSelfDelegation)
func (_DPoS *DPoSCallerSession) Validators(arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	BondBlock          *big.Int
	UnbondBlock        *big.Int
	CommissionRate     *big.Int
	MinSelfDelegation  *big.Int
}, error) {
	return _DPoS.Contract.Validators(&_DPoS.CallOpts, arg0)
}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_DPoS *DPoSCaller) VerifySignatures(opts *bind.CallOpts, _msg []byte, _sigs [][]byte) (bool, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "verifySignatures", _msg, _sigs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_DPoS *DPoSSession) VerifySignatures(_msg []byte, _sigs [][]byte) (bool, error) {
	return _DPoS.Contract.VerifySignatures(&_DPoS.CallOpts, _msg, _sigs)
}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_DPoS *DPoSCallerSession) VerifySignatures(_msg []byte, _sigs [][]byte) (bool, error) {
	return _DPoS.Contract.VerifySignatures(&_DPoS.CallOpts, _msg, _sigs)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_DPoS *DPoSCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_DPoS *DPoSSession) WhitelistEnabled() (bool, error) {
	return _DPoS.Contract.WhitelistEnabled(&_DPoS.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_DPoS *DPoSCallerSession) WhitelistEnabled() (bool, error) {
	return _DPoS.Contract.WhitelistEnabled(&_DPoS.CallOpts)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_DPoS *DPoSTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_DPoS *DPoSSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.AddWhitelisted(&_DPoS.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_DPoS *DPoSTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.AddWhitelisted(&_DPoS.TransactOpts, account)
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_DPoS *DPoSTransactor) BondValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "bondValidator")
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_DPoS *DPoSSession) BondValidator() (*types.Transaction, error) {
	return _DPoS.Contract.BondValidator(&_DPoS.TransactOpts)
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_DPoS *DPoSTransactorSession) BondValidator() (*types.Transaction, error) {
	return _DPoS.Contract.BondValidator(&_DPoS.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_DPoS *DPoSTransactor) ClaimReward(opts *bind.TransactOpts, _rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "claimReward", _rewardRequest, _sigs)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_DPoS *DPoSSession) ClaimReward(_rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _DPoS.Contract.ClaimReward(&_DPoS.TransactOpts, _rewardRequest, _sigs)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_DPoS *DPoSTransactorSession) ClaimReward(_rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _DPoS.Contract.ClaimReward(&_DPoS.TransactOpts, _rewardRequest, _sigs)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_DPoS *DPoSTransactor) CompleteUndelegate(opts *bind.TransactOpts, _valAddr common.Address) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "completeUndelegate", _valAddr)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_DPoS *DPoSSession) CompleteUndelegate(_valAddr common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.CompleteUndelegate(&_DPoS.TransactOpts, _valAddr)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_DPoS *DPoSTransactorSession) CompleteUndelegate(_valAddr common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.CompleteUndelegate(&_DPoS.TransactOpts, _valAddr)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_DPoS *DPoSTransactor) ConfirmParamProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "confirmParamProposal", _proposalId)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_DPoS *DPoSSession) ConfirmParamProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.ConfirmParamProposal(&_DPoS.TransactOpts, _proposalId)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_DPoS *DPoSTransactorSession) ConfirmParamProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.ConfirmParamProposal(&_DPoS.TransactOpts, _proposalId)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_DPoS *DPoSTransactor) ConfirmUnbondedValidator(opts *bind.TransactOpts, _valAddr common.Address) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "confirmUnbondedValidator", _valAddr)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_DPoS *DPoSSession) ConfirmUnbondedValidator(_valAddr common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.ConfirmUnbondedValidator(&_DPoS.TransactOpts, _valAddr)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_DPoS *DPoSTransactorSession) ConfirmUnbondedValidator(_valAddr common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.ConfirmUnbondedValidator(&_DPoS.TransactOpts, _valAddr)
}

// ContributeToMiningPool is a paid mutator transaction binding the contract method 0x4b7dba6b.
//
// Solidity: function contributeToMiningPool(uint256 _amount) returns()
func (_DPoS *DPoSTransactor) ContributeToMiningPool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "contributeToMiningPool", _amount)
}

// ContributeToMiningPool is a paid mutator transaction binding the contract method 0x4b7dba6b.
//
// Solidity: function contributeToMiningPool(uint256 _amount) returns()
func (_DPoS *DPoSSession) ContributeToMiningPool(_amount *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.ContributeToMiningPool(&_DPoS.TransactOpts, _amount)
}

// ContributeToMiningPool is a paid mutator transaction binding the contract method 0x4b7dba6b.
//
// Solidity: function contributeToMiningPool(uint256 _amount) returns()
func (_DPoS *DPoSTransactorSession) ContributeToMiningPool(_amount *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.ContributeToMiningPool(&_DPoS.TransactOpts, _amount)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_DPoS *DPoSTransactor) CreateParamProposal(opts *bind.TransactOpts, _name uint8, _value *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "createParamProposal", _name, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_DPoS *DPoSSession) CreateParamProposal(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.CreateParamProposal(&_DPoS.TransactOpts, _name, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_DPoS *DPoSTransactorSession) CreateParamProposal(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.CreateParamProposal(&_DPoS.TransactOpts, _name, _value)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_DPoS *DPoSTransactor) Delegate(opts *bind.TransactOpts, _valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "delegate", _valAddr, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_DPoS *DPoSSession) Delegate(_valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.Delegate(&_DPoS.TransactOpts, _valAddr, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_DPoS *DPoSTransactorSession) Delegate(_valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.Delegate(&_DPoS.TransactOpts, _valAddr, _tokens)
}

// DisableSlash is a paid mutator transaction binding the contract method 0x24b9bcc0.
//
// Solidity: function disableSlash() returns()
func (_DPoS *DPoSTransactor) DisableSlash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "disableSlash")
}

// DisableSlash is a paid mutator transaction binding the contract method 0x24b9bcc0.
//
// Solidity: function disableSlash() returns()
func (_DPoS *DPoSSession) DisableSlash() (*types.Transaction, error) {
	return _DPoS.Contract.DisableSlash(&_DPoS.TransactOpts)
}

// DisableSlash is a paid mutator transaction binding the contract method 0x24b9bcc0.
//
// Solidity: function disableSlash() returns()
func (_DPoS *DPoSTransactorSession) DisableSlash() (*types.Transaction, error) {
	return _DPoS.Contract.DisableSlash(&_DPoS.TransactOpts)
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_DPoS *DPoSTransactor) DisableWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "disableWhitelist")
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_DPoS *DPoSSession) DisableWhitelist() (*types.Transaction, error) {
	return _DPoS.Contract.DisableWhitelist(&_DPoS.TransactOpts)
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_DPoS *DPoSTransactorSession) DisableWhitelist() (*types.Transaction, error) {
	return _DPoS.Contract.DisableWhitelist(&_DPoS.TransactOpts)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_DPoS *DPoSTransactor) DrainToken(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "drainToken", _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_DPoS *DPoSSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.DrainToken(&_DPoS.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_DPoS *DPoSTransactorSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.DrainToken(&_DPoS.TransactOpts, _amount)
}

// EnableSlash is a paid mutator transaction binding the contract method 0x7a3ba4ad.
//
// Solidity: function enableSlash() returns()
func (_DPoS *DPoSTransactor) EnableSlash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "enableSlash")
}

// EnableSlash is a paid mutator transaction binding the contract method 0x7a3ba4ad.
//
// Solidity: function enableSlash() returns()
func (_DPoS *DPoSSession) EnableSlash() (*types.Transaction, error) {
	return _DPoS.Contract.EnableSlash(&_DPoS.TransactOpts)
}

// EnableSlash is a paid mutator transaction binding the contract method 0x7a3ba4ad.
//
// Solidity: function enableSlash() returns()
func (_DPoS *DPoSTransactorSession) EnableSlash() (*types.Transaction, error) {
	return _DPoS.Contract.EnableSlash(&_DPoS.TransactOpts)
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_DPoS *DPoSTransactor) EnableWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "enableWhitelist")
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_DPoS *DPoSSession) EnableWhitelist() (*types.Transaction, error) {
	return _DPoS.Contract.EnableWhitelist(&_DPoS.TransactOpts)
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_DPoS *DPoSTransactorSession) EnableWhitelist() (*types.Transaction, error) {
	return _DPoS.Contract.EnableWhitelist(&_DPoS.TransactOpts)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x24990d7b.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint256 _commissionRate) returns()
func (_DPoS *DPoSTransactor) InitializeValidator(opts *bind.TransactOpts, _signer common.Address, _minSelfDelegation *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "initializeValidator", _signer, _minSelfDelegation, _commissionRate)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x24990d7b.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint256 _commissionRate) returns()
func (_DPoS *DPoSSession) InitializeValidator(_signer common.Address, _minSelfDelegation *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.InitializeValidator(&_DPoS.TransactOpts, _signer, _minSelfDelegation, _commissionRate)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x24990d7b.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint256 _commissionRate) returns()
func (_DPoS *DPoSTransactorSession) InitializeValidator(_signer common.Address, _minSelfDelegation *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.InitializeValidator(&_DPoS.TransactOpts, _signer, _minSelfDelegation, _commissionRate)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DPoS *DPoSTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DPoS *DPoSSession) Pause() (*types.Transaction, error) {
	return _DPoS.Contract.Pause(&_DPoS.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DPoS *DPoSTransactorSession) Pause() (*types.Transaction, error) {
	return _DPoS.Contract.Pause(&_DPoS.TransactOpts)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_DPoS *DPoSTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_DPoS *DPoSSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.RemoveWhitelisted(&_DPoS.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_DPoS *DPoSTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.RemoveWhitelisted(&_DPoS.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DPoS *DPoSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DPoS *DPoSSession) RenounceOwnership() (*types.Transaction, error) {
	return _DPoS.Contract.RenounceOwnership(&_DPoS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DPoS *DPoSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DPoS.Contract.RenounceOwnership(&_DPoS.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_DPoS *DPoSTransactor) Slash(opts *bind.TransactOpts, _slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "slash", _slashRequest, _sigs)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_DPoS *DPoSSession) Slash(_slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _DPoS.Contract.Slash(&_DPoS.TransactOpts, _slashRequest, _sigs)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_DPoS *DPoSTransactorSession) Slash(_slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _DPoS.Contract.Slash(&_DPoS.TransactOpts, _slashRequest, _sigs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DPoS *DPoSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DPoS *DPoSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.TransferOwnership(&_DPoS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DPoS *DPoSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.TransferOwnership(&_DPoS.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_DPoS *DPoSTransactor) Undelegate(opts *bind.TransactOpts, _valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "undelegate", _valAddr, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_DPoS *DPoSSession) Undelegate(_valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.Undelegate(&_DPoS.TransactOpts, _valAddr, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_DPoS *DPoSTransactorSession) Undelegate(_valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.Undelegate(&_DPoS.TransactOpts, _valAddr, _shares)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DPoS *DPoSTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DPoS *DPoSSession) Unpause() (*types.Transaction, error) {
	return _DPoS.Contract.Unpause(&_DPoS.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DPoS *DPoSTransactorSession) Unpause() (*types.Transaction, error) {
	return _DPoS.Contract.Unpause(&_DPoS.TransactOpts)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 _newRate) returns()
func (_DPoS *DPoSTransactor) UpdateCommissionRate(opts *bind.TransactOpts, _newRate *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "updateCommissionRate", _newRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 _newRate) returns()
func (_DPoS *DPoSSession) UpdateCommissionRate(_newRate *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.UpdateCommissionRate(&_DPoS.TransactOpts, _newRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 _newRate) returns()
func (_DPoS *DPoSTransactorSession) UpdateCommissionRate(_newRate *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.UpdateCommissionRate(&_DPoS.TransactOpts, _newRate)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_DPoS *DPoSTransactor) UpdateMinSelfDelegation(opts *bind.TransactOpts, _minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "updateMinSelfDelegation", _minSelfDelegation)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_DPoS *DPoSSession) UpdateMinSelfDelegation(_minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.UpdateMinSelfDelegation(&_DPoS.TransactOpts, _minSelfDelegation)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_DPoS *DPoSTransactorSession) UpdateMinSelfDelegation(_minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.UpdateMinSelfDelegation(&_DPoS.TransactOpts, _minSelfDelegation)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_DPoS *DPoSTransactor) UpdateValidatorSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "updateValidatorSigner", _signer)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_DPoS *DPoSSession) UpdateValidatorSigner(_signer common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.UpdateValidatorSigner(&_DPoS.TransactOpts, _signer)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_DPoS *DPoSTransactorSession) UpdateValidatorSigner(_signer common.Address) (*types.Transaction, error) {
	return _DPoS.Contract.UpdateValidatorSigner(&_DPoS.TransactOpts, _signer)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_DPoS *DPoSTransactor) VoteParam(opts *bind.TransactOpts, _proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "voteParam", _proposalId, _vote)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_DPoS *DPoSSession) VoteParam(_proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _DPoS.Contract.VoteParam(&_DPoS.TransactOpts, _proposalId, _vote)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_DPoS *DPoSTransactorSession) VoteParam(_proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _DPoS.Contract.VoteParam(&_DPoS.TransactOpts, _proposalId, _vote)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DPoS *DPoSTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DPoS.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DPoS *DPoSSession) Receive() (*types.Transaction, error) {
	return _DPoS.Contract.Receive(&_DPoS.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DPoS *DPoSTransactorSession) Receive() (*types.Transaction, error) {
	return _DPoS.Contract.Receive(&_DPoS.TransactOpts)
}

// DPoSConfirmParamProposalIterator is returned from FilterConfirmParamProposal and is used to iterate over the raw logs and unpacked data for ConfirmParamProposal events raised by the DPoS contract.
type DPoSConfirmParamProposalIterator struct {
	Event *DPoSConfirmParamProposal // Event containing the contract specifics and raw log

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
func (it *DPoSConfirmParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSConfirmParamProposal)
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
		it.Event = new(DPoSConfirmParamProposal)
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
func (it *DPoSConfirmParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSConfirmParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSConfirmParamProposal represents a ConfirmParamProposal event raised by the DPoS contract.
type DPoSConfirmParamProposal struct {
	ProposalId *big.Int
	Passed     bool
	Name       uint8
	NewValue   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfirmParamProposal is a free log retrieval operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_DPoS *DPoSFilterer) FilterConfirmParamProposal(opts *bind.FilterOpts) (*DPoSConfirmParamProposalIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return &DPoSConfirmParamProposalIterator{contract: _DPoS.contract, event: "ConfirmParamProposal", logs: logs, sub: sub}, nil
}

// WatchConfirmParamProposal is a free log subscription operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_DPoS *DPoSFilterer) WatchConfirmParamProposal(opts *bind.WatchOpts, sink chan<- *DPoSConfirmParamProposal) (event.Subscription, error) {

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSConfirmParamProposal)
				if err := _DPoS.contract.UnpackLog(event, "ConfirmParamProposal", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseConfirmParamProposal(log types.Log) (*DPoSConfirmParamProposal, error) {
	event := new(DPoSConfirmParamProposal)
	if err := _DPoS.contract.UnpackLog(event, "ConfirmParamProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSCreateParamProposalIterator is returned from FilterCreateParamProposal and is used to iterate over the raw logs and unpacked data for CreateParamProposal events raised by the DPoS contract.
type DPoSCreateParamProposalIterator struct {
	Event *DPoSCreateParamProposal // Event containing the contract specifics and raw log

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
func (it *DPoSCreateParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSCreateParamProposal)
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
		it.Event = new(DPoSCreateParamProposal)
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
func (it *DPoSCreateParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSCreateParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSCreateParamProposal represents a CreateParamProposal event raised by the DPoS contract.
type DPoSCreateParamProposal struct {
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
func (_DPoS *DPoSFilterer) FilterCreateParamProposal(opts *bind.FilterOpts) (*DPoSCreateParamProposalIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return &DPoSCreateParamProposalIterator{contract: _DPoS.contract, event: "CreateParamProposal", logs: logs, sub: sub}, nil
}

// WatchCreateParamProposal is a free log subscription operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_DPoS *DPoSFilterer) WatchCreateParamProposal(opts *bind.WatchOpts, sink chan<- *DPoSCreateParamProposal) (event.Subscription, error) {

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSCreateParamProposal)
				if err := _DPoS.contract.UnpackLog(event, "CreateParamProposal", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseCreateParamProposal(log types.Log) (*DPoSCreateParamProposal, error) {
	event := new(DPoSCreateParamProposal)
	if err := _DPoS.contract.UnpackLog(event, "CreateParamProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSDelegationUpdateIterator is returned from FilterDelegationUpdate and is used to iterate over the raw logs and unpacked data for DelegationUpdate events raised by the DPoS contract.
type DPoSDelegationUpdateIterator struct {
	Event *DPoSDelegationUpdate // Event containing the contract specifics and raw log

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
func (it *DPoSDelegationUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSDelegationUpdate)
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
		it.Event = new(DPoSDelegationUpdate)
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
func (it *DPoSDelegationUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSDelegationUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSDelegationUpdate represents a DelegationUpdate event raised by the DPoS contract.
type DPoSDelegationUpdate struct {
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
func (_DPoS *DPoSFilterer) FilterDelegationUpdate(opts *bind.FilterOpts, valAddr []common.Address, delAddr []common.Address) (*DPoSDelegationUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "DelegationUpdate", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return &DPoSDelegationUpdateIterator{contract: _DPoS.contract, event: "DelegationUpdate", logs: logs, sub: sub}, nil
}

// WatchDelegationUpdate is a free log subscription operation binding the contract event 0x2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea.
//
// Solidity: event DelegationUpdate(address indexed valAddr, address indexed delAddr, uint256 valTokens, uint256 delShares, int256 tokenDiff)
func (_DPoS *DPoSFilterer) WatchDelegationUpdate(opts *bind.WatchOpts, sink chan<- *DPoSDelegationUpdate, valAddr []common.Address, delAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "DelegationUpdate", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSDelegationUpdate)
				if err := _DPoS.contract.UnpackLog(event, "DelegationUpdate", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseDelegationUpdate(log types.Log) (*DPoSDelegationUpdate, error) {
	event := new(DPoSDelegationUpdate)
	if err := _DPoS.contract.UnpackLog(event, "DelegationUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSMiningPoolContributionIterator is returned from FilterMiningPoolContribution and is used to iterate over the raw logs and unpacked data for MiningPoolContribution events raised by the DPoS contract.
type DPoSMiningPoolContributionIterator struct {
	Event *DPoSMiningPoolContribution // Event containing the contract specifics and raw log

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
func (it *DPoSMiningPoolContributionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSMiningPoolContribution)
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
		it.Event = new(DPoSMiningPoolContribution)
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
func (it *DPoSMiningPoolContributionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSMiningPoolContributionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSMiningPoolContribution represents a MiningPoolContribution event raised by the DPoS contract.
type DPoSMiningPoolContribution struct {
	Contributor    common.Address
	Contribution   *big.Int
	RewardPoolSize *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMiningPoolContribution is a free log retrieval operation binding the contract event 0x97e19c4040b6c46d4275e0c4fea68f8f92c81138372ffdb089932c211938f765.
//
// Solidity: event MiningPoolContribution(address indexed contributor, uint256 contribution, uint256 rewardPoolSize)
func (_DPoS *DPoSFilterer) FilterMiningPoolContribution(opts *bind.FilterOpts, contributor []common.Address) (*DPoSMiningPoolContributionIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "MiningPoolContribution", contributorRule)
	if err != nil {
		return nil, err
	}
	return &DPoSMiningPoolContributionIterator{contract: _DPoS.contract, event: "MiningPoolContribution", logs: logs, sub: sub}, nil
}

// WatchMiningPoolContribution is a free log subscription operation binding the contract event 0x97e19c4040b6c46d4275e0c4fea68f8f92c81138372ffdb089932c211938f765.
//
// Solidity: event MiningPoolContribution(address indexed contributor, uint256 contribution, uint256 rewardPoolSize)
func (_DPoS *DPoSFilterer) WatchMiningPoolContribution(opts *bind.WatchOpts, sink chan<- *DPoSMiningPoolContribution, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "MiningPoolContribution", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSMiningPoolContribution)
				if err := _DPoS.contract.UnpackLog(event, "MiningPoolContribution", log); err != nil {
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

// ParseMiningPoolContribution is a log parse operation binding the contract event 0x97e19c4040b6c46d4275e0c4fea68f8f92c81138372ffdb089932c211938f765.
//
// Solidity: event MiningPoolContribution(address indexed contributor, uint256 contribution, uint256 rewardPoolSize)
func (_DPoS *DPoSFilterer) ParseMiningPoolContribution(log types.Log) (*DPoSMiningPoolContribution, error) {
	event := new(DPoSMiningPoolContribution)
	if err := _DPoS.contract.UnpackLog(event, "MiningPoolContribution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DPoS contract.
type DPoSOwnershipTransferredIterator struct {
	Event *DPoSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DPoSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSOwnershipTransferred)
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
		it.Event = new(DPoSOwnershipTransferred)
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
func (it *DPoSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSOwnershipTransferred represents a OwnershipTransferred event raised by the DPoS contract.
type DPoSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DPoS *DPoSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DPoSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DPoSOwnershipTransferredIterator{contract: _DPoS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DPoS *DPoSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DPoSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSOwnershipTransferred)
				if err := _DPoS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseOwnershipTransferred(log types.Log) (*DPoSOwnershipTransferred, error) {
	event := new(DPoSOwnershipTransferred)
	if err := _DPoS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DPoS contract.
type DPoSPausedIterator struct {
	Event *DPoSPaused // Event containing the contract specifics and raw log

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
func (it *DPoSPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSPaused)
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
		it.Event = new(DPoSPaused)
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
func (it *DPoSPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSPaused represents a Paused event raised by the DPoS contract.
type DPoSPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DPoS *DPoSFilterer) FilterPaused(opts *bind.FilterOpts) (*DPoSPausedIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DPoSPausedIterator{contract: _DPoS.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DPoS *DPoSFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DPoSPaused) (event.Subscription, error) {

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSPaused)
				if err := _DPoS.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParsePaused(log types.Log) (*DPoSPaused, error) {
	event := new(DPoSPaused)
	if err := _DPoS.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the DPoS contract.
type DPoSRewardClaimedIterator struct {
	Event *DPoSRewardClaimed // Event containing the contract specifics and raw log

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
func (it *DPoSRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSRewardClaimed)
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
		it.Event = new(DPoSRewardClaimed)
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
func (it *DPoSRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSRewardClaimed represents a RewardClaimed event raised by the DPoS contract.
type DPoSRewardClaimed struct {
	Recipient  common.Address
	Reward     *big.Int
	RewardPool *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0xf01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743.
//
// Solidity: event RewardClaimed(address indexed recipient, uint256 reward, uint256 rewardPool)
func (_DPoS *DPoSFilterer) FilterRewardClaimed(opts *bind.FilterOpts, recipient []common.Address) (*DPoSRewardClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "RewardClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DPoSRewardClaimedIterator{contract: _DPoS.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0xf01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743.
//
// Solidity: event RewardClaimed(address indexed recipient, uint256 reward, uint256 rewardPool)
func (_DPoS *DPoSFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *DPoSRewardClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "RewardClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSRewardClaimed)
				if err := _DPoS.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseRewardClaimed(log types.Log) (*DPoSRewardClaimed, error) {
	event := new(DPoSRewardClaimed)
	if err := _DPoS.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSSlashIterator is returned from FilterSlash and is used to iterate over the raw logs and unpacked data for Slash events raised by the DPoS contract.
type DPoSSlashIterator struct {
	Event *DPoSSlash // Event containing the contract specifics and raw log

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
func (it *DPoSSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSSlash)
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
		it.Event = new(DPoSSlash)
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
func (it *DPoSSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSSlash represents a Slash event raised by the DPoS contract.
type DPoSSlash struct {
	ValAddr  common.Address
	Nonce    uint64
	SlashAmt *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlash is a free log retrieval operation binding the contract event 0x10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008.
//
// Solidity: event Slash(address indexed valAddr, uint64 nonce, uint256 slashAmt)
func (_DPoS *DPoSFilterer) FilterSlash(opts *bind.FilterOpts, valAddr []common.Address) (*DPoSSlashIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "Slash", valAddrRule)
	if err != nil {
		return nil, err
	}
	return &DPoSSlashIterator{contract: _DPoS.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0x10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008.
//
// Solidity: event Slash(address indexed valAddr, uint64 nonce, uint256 slashAmt)
func (_DPoS *DPoSFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *DPoSSlash, valAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "Slash", valAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSSlash)
				if err := _DPoS.contract.UnpackLog(event, "Slash", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseSlash(log types.Log) (*DPoSSlash, error) {
	event := new(DPoSSlash)
	if err := _DPoS.contract.UnpackLog(event, "Slash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSSlashAmtCollectedIterator is returned from FilterSlashAmtCollected and is used to iterate over the raw logs and unpacked data for SlashAmtCollected events raised by the DPoS contract.
type DPoSSlashAmtCollectedIterator struct {
	Event *DPoSSlashAmtCollected // Event containing the contract specifics and raw log

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
func (it *DPoSSlashAmtCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSSlashAmtCollected)
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
		it.Event = new(DPoSSlashAmtCollected)
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
func (it *DPoSSlashAmtCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSSlashAmtCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSSlashAmtCollected represents a SlashAmtCollected event raised by the DPoS contract.
type DPoSSlashAmtCollected struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashAmtCollected is a free log retrieval operation binding the contract event 0xb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3.
//
// Solidity: event SlashAmtCollected(address indexed recipient, uint256 amount)
func (_DPoS *DPoSFilterer) FilterSlashAmtCollected(opts *bind.FilterOpts, recipient []common.Address) (*DPoSSlashAmtCollectedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "SlashAmtCollected", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DPoSSlashAmtCollectedIterator{contract: _DPoS.contract, event: "SlashAmtCollected", logs: logs, sub: sub}, nil
}

// WatchSlashAmtCollected is a free log subscription operation binding the contract event 0xb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3.
//
// Solidity: event SlashAmtCollected(address indexed recipient, uint256 amount)
func (_DPoS *DPoSFilterer) WatchSlashAmtCollected(opts *bind.WatchOpts, sink chan<- *DPoSSlashAmtCollected, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "SlashAmtCollected", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSSlashAmtCollected)
				if err := _DPoS.contract.UnpackLog(event, "SlashAmtCollected", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseSlashAmtCollected(log types.Log) (*DPoSSlashAmtCollected, error) {
	event := new(DPoSSlashAmtCollected)
	if err := _DPoS.contract.UnpackLog(event, "SlashAmtCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the DPoS contract.
type DPoSUndelegatedIterator struct {
	Event *DPoSUndelegated // Event containing the contract specifics and raw log

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
func (it *DPoSUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSUndelegated)
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
		it.Event = new(DPoSUndelegated)
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
func (it *DPoSUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSUndelegated represents a Undelegated event raised by the DPoS contract.
type DPoSUndelegated struct {
	ValAddr common.Address
	DelAddr common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed valAddr, address indexed delAddr, uint256 amount)
func (_DPoS *DPoSFilterer) FilterUndelegated(opts *bind.FilterOpts, valAddr []common.Address, delAddr []common.Address) (*DPoSUndelegatedIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "Undelegated", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return &DPoSUndelegatedIterator{contract: _DPoS.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed valAddr, address indexed delAddr, uint256 amount)
func (_DPoS *DPoSFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *DPoSUndelegated, valAddr []common.Address, delAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "Undelegated", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSUndelegated)
				if err := _DPoS.contract.UnpackLog(event, "Undelegated", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseUndelegated(log types.Log) (*DPoSUndelegated, error) {
	event := new(DPoSUndelegated)
	if err := _DPoS.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DPoS contract.
type DPoSUnpausedIterator struct {
	Event *DPoSUnpaused // Event containing the contract specifics and raw log

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
func (it *DPoSUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSUnpaused)
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
		it.Event = new(DPoSUnpaused)
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
func (it *DPoSUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSUnpaused represents a Unpaused event raised by the DPoS contract.
type DPoSUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DPoS *DPoSFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DPoSUnpausedIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DPoSUnpausedIterator{contract: _DPoS.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DPoS *DPoSFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DPoSUnpaused) (event.Subscription, error) {

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSUnpaused)
				if err := _DPoS.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseUnpaused(log types.Log) (*DPoSUnpaused, error) {
	event := new(DPoSUnpaused)
	if err := _DPoS.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSValidatorParamsUpdateIterator is returned from FilterValidatorParamsUpdate and is used to iterate over the raw logs and unpacked data for ValidatorParamsUpdate events raised by the DPoS contract.
type DPoSValidatorParamsUpdateIterator struct {
	Event *DPoSValidatorParamsUpdate // Event containing the contract specifics and raw log

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
func (it *DPoSValidatorParamsUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSValidatorParamsUpdate)
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
		it.Event = new(DPoSValidatorParamsUpdate)
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
func (it *DPoSValidatorParamsUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSValidatorParamsUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSValidatorParamsUpdate represents a ValidatorParamsUpdate event raised by the DPoS contract.
type DPoSValidatorParamsUpdate struct {
	ValAddr           common.Address
	Signer            common.Address
	MinSelfDelegation *big.Int
	CommissionRate    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorParamsUpdate is a free log retrieval operation binding the contract event 0x848523f3f4f7d056ddf7150c34e51564204e16c45e8b98cae2263d507a875346.
//
// Solidity: event ValidatorParamsUpdate(address indexed valAddr, address indexed signer, uint256 minSelfDelegation, uint256 commissionRate)
func (_DPoS *DPoSFilterer) FilterValidatorParamsUpdate(opts *bind.FilterOpts, valAddr []common.Address, signer []common.Address) (*DPoSValidatorParamsUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "ValidatorParamsUpdate", valAddrRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &DPoSValidatorParamsUpdateIterator{contract: _DPoS.contract, event: "ValidatorParamsUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorParamsUpdate is a free log subscription operation binding the contract event 0x848523f3f4f7d056ddf7150c34e51564204e16c45e8b98cae2263d507a875346.
//
// Solidity: event ValidatorParamsUpdate(address indexed valAddr, address indexed signer, uint256 minSelfDelegation, uint256 commissionRate)
func (_DPoS *DPoSFilterer) WatchValidatorParamsUpdate(opts *bind.WatchOpts, sink chan<- *DPoSValidatorParamsUpdate, valAddr []common.Address, signer []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "ValidatorParamsUpdate", valAddrRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSValidatorParamsUpdate)
				if err := _DPoS.contract.UnpackLog(event, "ValidatorParamsUpdate", log); err != nil {
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

// ParseValidatorParamsUpdate is a log parse operation binding the contract event 0x848523f3f4f7d056ddf7150c34e51564204e16c45e8b98cae2263d507a875346.
//
// Solidity: event ValidatorParamsUpdate(address indexed valAddr, address indexed signer, uint256 minSelfDelegation, uint256 commissionRate)
func (_DPoS *DPoSFilterer) ParseValidatorParamsUpdate(log types.Log) (*DPoSValidatorParamsUpdate, error) {
	event := new(DPoSValidatorParamsUpdate)
	if err := _DPoS.contract.UnpackLog(event, "ValidatorParamsUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSValidatorStatusUpdateIterator is returned from FilterValidatorStatusUpdate and is used to iterate over the raw logs and unpacked data for ValidatorStatusUpdate events raised by the DPoS contract.
type DPoSValidatorStatusUpdateIterator struct {
	Event *DPoSValidatorStatusUpdate // Event containing the contract specifics and raw log

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
func (it *DPoSValidatorStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSValidatorStatusUpdate)
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
		it.Event = new(DPoSValidatorStatusUpdate)
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
func (it *DPoSValidatorStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSValidatorStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSValidatorStatusUpdate represents a ValidatorStatusUpdate event raised by the DPoS contract.
type DPoSValidatorStatusUpdate struct {
	ValAddr common.Address
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusUpdate is a free log retrieval operation binding the contract event 0xd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a.
//
// Solidity: event ValidatorStatusUpdate(address indexed valAddr, uint8 indexed status)
func (_DPoS *DPoSFilterer) FilterValidatorStatusUpdate(opts *bind.FilterOpts, valAddr []common.Address, status []uint8) (*DPoSValidatorStatusUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "ValidatorStatusUpdate", valAddrRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &DPoSValidatorStatusUpdateIterator{contract: _DPoS.contract, event: "ValidatorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusUpdate is a free log subscription operation binding the contract event 0xd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a.
//
// Solidity: event ValidatorStatusUpdate(address indexed valAddr, uint8 indexed status)
func (_DPoS *DPoSFilterer) WatchValidatorStatusUpdate(opts *bind.WatchOpts, sink chan<- *DPoSValidatorStatusUpdate, valAddr []common.Address, status []uint8) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "ValidatorStatusUpdate", valAddrRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSValidatorStatusUpdate)
				if err := _DPoS.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseValidatorStatusUpdate(log types.Log) (*DPoSValidatorStatusUpdate, error) {
	event := new(DPoSValidatorStatusUpdate)
	if err := _DPoS.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSVoteParamIterator is returned from FilterVoteParam and is used to iterate over the raw logs and unpacked data for VoteParam events raised by the DPoS contract.
type DPoSVoteParamIterator struct {
	Event *DPoSVoteParam // Event containing the contract specifics and raw log

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
func (it *DPoSVoteParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSVoteParam)
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
		it.Event = new(DPoSVoteParam)
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
func (it *DPoSVoteParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSVoteParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSVoteParam represents a VoteParam event raised by the DPoS contract.
type DPoSVoteParam struct {
	ProposalId *big.Int
	Voter      common.Address
	VoteType   uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteParam is a free log retrieval operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 voteType)
func (_DPoS *DPoSFilterer) FilterVoteParam(opts *bind.FilterOpts) (*DPoSVoteParamIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return &DPoSVoteParamIterator{contract: _DPoS.contract, event: "VoteParam", logs: logs, sub: sub}, nil
}

// WatchVoteParam is a free log subscription operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 voteType)
func (_DPoS *DPoSFilterer) WatchVoteParam(opts *bind.WatchOpts, sink chan<- *DPoSVoteParam) (event.Subscription, error) {

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSVoteParam)
				if err := _DPoS.contract.UnpackLog(event, "VoteParam", log); err != nil {
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
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 voteType)
func (_DPoS *DPoSFilterer) ParseVoteParam(log types.Log) (*DPoSVoteParam, error) {
	event := new(DPoSVoteParam)
	if err := _DPoS.contract.UnpackLog(event, "VoteParam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the DPoS contract.
type DPoSWhitelistedAddedIterator struct {
	Event *DPoSWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *DPoSWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSWhitelistedAdded)
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
		it.Event = new(DPoSWhitelistedAdded)
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
func (it *DPoSWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSWhitelistedAdded represents a WhitelistedAdded event raised by the DPoS contract.
type DPoSWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_DPoS *DPoSFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts) (*DPoSWhitelistedAddedIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return &DPoSWhitelistedAddedIterator{contract: _DPoS.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_DPoS *DPoSFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *DPoSWhitelistedAdded) (event.Subscription, error) {

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSWhitelistedAdded)
				if err := _DPoS.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseWhitelistedAdded(log types.Log) (*DPoSWhitelistedAdded, error) {
	event := new(DPoSWhitelistedAdded)
	if err := _DPoS.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DPoSWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the DPoS contract.
type DPoSWhitelistedRemovedIterator struct {
	Event *DPoSWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *DPoSWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DPoSWhitelistedRemoved)
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
		it.Event = new(DPoSWhitelistedRemoved)
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
func (it *DPoSWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DPoSWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DPoSWhitelistedRemoved represents a WhitelistedRemoved event raised by the DPoS contract.
type DPoSWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_DPoS *DPoSFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts) (*DPoSWhitelistedRemovedIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return &DPoSWhitelistedRemovedIterator{contract: _DPoS.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_DPoS *DPoSFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *DPoSWhitelistedRemoved) (event.Subscription, error) {

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DPoSWhitelistedRemoved)
				if err := _DPoS.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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
func (_DPoS *DPoSFilterer) ParseWhitelistedRemoved(log types.Log) (*DPoSWhitelistedRemoved, error) {
	event := new(DPoSWhitelistedRemoved)
	if err := _DPoS.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernMetaData contains all meta data concerning the Govern contract.
var GovernMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_governProposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_governVoteTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorBondInterval\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ConfirmParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CreateParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGovern.VoteType\",\"name\":\"voteType\",\"type\":\"uint8\"}],\"name\":\"VoteParam\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"celerToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"createParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getParamProposalVote\",\"outputs\":[{\"internalType\":\"enumGovern.VoteType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"}],\"name\":\"getParamValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextParamProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paramProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGovern.ParamName\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610cd7380380610cd783398101604081905261002f916101d1565b61003833610181565b600180546001600160a01b0319166001600160a01b039a909a169990991790985560026020527fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b969096557fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e0949094557f679795a0195a1b76cdebb7c51d74e058aee92919b8c3389af86ef24535e8a28c929092557f88601476d11616a71c5be67555bd1dff4b1cbf21533d2669b768b61518cfe1c3557fee60d0579bcffd98e668647d59fec1ff86a7fb340ce572e844f234ae73a6918f557fb98b78633099fa36ed8b8680c4f8092689e1e04080eb9cbb077ca38a14d7e384557f59dd4b18488d12f51eda69757a0ed42a2010c14b564330cc74a06895e60c077b5560076000527facd8ef244210bb6898e73c48bf820ed8ecc857a3bab8d79c10e4fa92b1e9ca6555610251565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008060008060008060008060006101208a8c0312156101f057600080fd5b89516001600160a01b038116811461020757600080fd5b8099505060208a0151975060408a0151965060608a0151955060808a0151945060a08a0151935060c08a0151925060e08a015191506101008a015190509295985092959850929598565b610a77806102606000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638da5cb5b116100665780638da5cb5b1461018d578063c6c21e9d146101b2578063e478ed9d146101c5578063eb505dd5146101d8578063f2fde38b146101f857600080fd5b806322da7927146100a3578063410ae02c146100bf578063581c53c5146100d2578063715018a61461011c5780637e5fb8f314610126575b600080fd5b6100ac60045481565b6040519081526020015b60405180910390f35b6100ac6100cd366004610830565b61020b565b61010f6100e036600461088e565b60008281526003602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b6040516100b6919061093c565b61012461024a565b005b61017b610134366004610875565b60036020819052600091825260409091208054600182015460028301549383015460048401546005909401546001600160a01b03909316949193919260ff91821692911686565b6040516100b6969594939291906108ea565b6000546001600160a01b03165b6040516001600160a01b0390911681526020016100b6565b60015461019a906001600160a01b031681565b6101246101d336600461084b565b6102b5565b6100ac6101e6366004610830565b60026020526000908152604090205481565b6101246102063660046107f3565b6103fc565b60006002600083600781111561022357610223610a2b565b600781111561023457610234610a2b565b8152602001908152602001600020549050919050565b6000546001600160a01b031633146102a95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6102b360006104c7565b565b6004546000818152600360205260409020906102d29060016109ba565b60045560026020527fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b5481546001600160a01b03191633908117835560018084018390556000527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e05490919061034890436109ba565b600284015560038301805486919060ff1916600183600781111561036e5761036e610a2b565b02179055506004830184905560058301805460ff19166001908117909155546103a2906001600160a01b0316833084610517565b7f4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f16060016004546103d291906109d2565b8383866002015489896040516103ed96959493929190610989565b60405180910390a15050505050565b6000546001600160a01b031633146104565760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a0565b6001600160a01b0381166104bb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102a0565b6104c4816104c7565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052610571908590610577565b50505050565b60006105cc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661064e9092919063ffffffff16565b80519091501561064957808060200190518101906105ea919061080e565b6106495760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016102a0565b505050565b606061065d8484600085610667565b90505b9392505050565b6060824710156106c85760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016102a0565b843b6107165760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102a0565b600080866001600160a01b0316858760405161073291906108ce565b60006040518083038185875af1925050503d806000811461076f576040519150601f19603f3d011682016040523d82523d6000602084013e610774565b606091505b509150915061078482828661078f565b979650505050505050565b6060831561079e575081610660565b8251156107ae5782518084602001fd5b8160405162461bcd60e51b81526004016102a09190610956565b80356001600160a01b03811681146107df57600080fd5b919050565b8035600881106107df57600080fd5b60006020828403121561080557600080fd5b610660826107c8565b60006020828403121561082057600080fd5b8151801515811461066057600080fd5b60006020828403121561084257600080fd5b610660826107e4565b6000806040838503121561085e57600080fd5b610867836107e4565b946020939093013593505050565b60006020828403121561088757600080fd5b5035919050565b600080604083850312156108a157600080fd5b823591506108b1602084016107c8565b90509250929050565b600881106108ca576108ca610a2b565b9052565b600082516108e08184602087016109e9565b9190910192915050565b6001600160a01b0387168152602081018690526040810185905260c0810161091560608301866108ba565b8360808301526003831061092b5761092b610a2b565b8260a0830152979650505050505050565b602081016004831061095057610950610a2b565b91905290565b60208152600082518060208401526109758160408501602087016109e9565b601f01601f19169190910160400192915050565b8681526001600160a01b0386166020820152604081018590526060810184905260c0810161092b60808301856108ba565b600082198211156109cd576109cd610a15565b500190565b6000828210156109e4576109e4610a15565b500390565b60005b83811015610a045781810151838201526020016109ec565b838111156105715750506000910152565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fdfea26469706673582212202705513e6ff213f1b500087883f6cb6dfd5841035b9ec73fb119bb8a6b9714ad64736f6c63430008070033",
}

// GovernABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernMetaData.ABI instead.
var GovernABI = GovernMetaData.ABI

// GovernBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernMetaData.Bin instead.
var GovernBin = GovernMetaData.Bin

// DeployGovern deploys a new Ethereum contract, binding an instance of Govern to it.
func DeployGovern(auth *bind.TransactOpts, backend bind.ContractBackend, _celerTokenAddress common.Address, _governProposalDeposit *big.Int, _governVoteTimeout *big.Int, _slashTimeout *big.Int, _maxBondedValidators *big.Int, _minValidatorTokens *big.Int, _minSelfDelegation *big.Int, _advanceNoticePeriod *big.Int, _validatorBondInterval *big.Int) (common.Address, *types.Transaction, *Govern, error) {
	parsed, err := GovernMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernBin), backend, _celerTokenAddress, _governProposalDeposit, _governVoteTimeout, _slashTimeout, _maxBondedValidators, _minValidatorTokens, _minSelfDelegation, _advanceNoticePeriod, _validatorBondInterval)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Govern *GovernCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Govern *GovernSession) Owner() (common.Address, error) {
	return _Govern.Contract.Owner(&_Govern.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Govern *GovernCallerSession) Owner() (common.Address, error) {
	return _Govern.Contract.Owner(&_Govern.CallOpts)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Govern *GovernTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Govern.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Govern *GovernSession) RenounceOwnership() (*types.Transaction, error) {
	return _Govern.Contract.RenounceOwnership(&_Govern.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Govern *GovernTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Govern.Contract.RenounceOwnership(&_Govern.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Govern *GovernTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Govern.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Govern *GovernSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Govern.Contract.TransferOwnership(&_Govern.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Govern *GovernTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Govern.Contract.TransferOwnership(&_Govern.TransactOpts, newOwner)
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

// GovernOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Govern contract.
type GovernOwnershipTransferredIterator struct {
	Event *GovernOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovernOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernOwnershipTransferred)
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
		it.Event = new(GovernOwnershipTransferred)
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
func (it *GovernOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernOwnershipTransferred represents a OwnershipTransferred event raised by the Govern contract.
type GovernOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Govern *GovernFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovernOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Govern.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovernOwnershipTransferredIterator{contract: _Govern.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Govern *GovernFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Govern.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernOwnershipTransferred)
				if err := _Govern.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Govern *GovernFilterer) ParseOwnershipTransferred(log types.Log) (*GovernOwnershipTransferred, error) {
	event := new(GovernOwnershipTransferred)
	if err := _Govern.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
	VoteType   uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteParam is a free log retrieval operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 voteType)
func (_Govern *GovernFilterer) FilterVoteParam(opts *bind.FilterOpts) (*GovernVoteParamIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return &GovernVoteParamIterator{contract: _Govern.contract, event: "VoteParam", logs: logs, sub: sub}, nil
}

// WatchVoteParam is a free log subscription operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 voteType)
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
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 voteType)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDPoS\",\"name\":\"_dpos\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddr\",\"type\":\"bytes\"}],\"name\":\"SgnAddrUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dpos\",\"outputs\":[{\"internalType\":\"contractDPoS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sgnAddrs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sgnAddr\",\"type\":\"bytes\"}],\"name\":\"updateSgnAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_withdrawalRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnAmts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161188638038061188683398101604081905261002f916100aa565b6100383361005a565b6000805460ff60a01b1916905560601b6001600160601b0319166080526100da565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100bc57600080fd5b81516001600160a01b03811681146100d357600080fd5b9392505050565b60805160601c611780610106600039600081816101a101528181610562015261079601526117806000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638f8451b21161008c578063c429fe1f11610066578063c429fe1f146101e9578063d0bb935114610209578063d88ef2711461021c578063f2fde38b1461022f57600080fd5b80638f8451b21461019c5780639d4323be146101c3578063b02c43d0146101d657600080fd5b8063715018a6116100c8578063715018a61461012e578063795c2c14146101365780638456cb591461016f5780638da5cb5b1461017757600080fd5b80633f4ba83a146100ef57806347e7ef24146100f95780635c975abb1461010c575b600080fd5b6100f7610242565b005b6100f76101073660046112f0565b61027f565b600054600160a01b900460ff1660405190151581526020015b60405180910390f35b6100f7610394565b6101616101443660046112bd565b600260209081526000928352604080842090915290825290205481565b604051908152602001610125565b6100f76103c8565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610125565b6101847f000000000000000000000000000000000000000000000000000000000000000081565b6100f76101d13660046112f0565b6103fa565b6101616101e4366004611439565b61048c565b6101fc6101f73660046112a2565b6104ad565b6040516101259190611574565b6100f761021736600461133c565b610547565b6100f761022a36600461137e565b610755565b6100f761023d3660046112a2565b610998565b6000546001600160a01b031633146102755760405162461bcd60e51b815260040161026c906115e1565b60405180910390fd5b61027d610a33565b565b600054600160a01b900460ff16156102a95760405162461bcd60e51b815260040161026c906115b7565b6040516bffffffffffffffffffffffff1933606081811b8316602085015285901b9091166034830152604882018390529060019060680160408051601f198184030181529190528051602091820120825460018101845560009384529190922001556103206001600160a01b038416823085610ad0565b600180546000916103309161166f565b6040805167ffffffffffffffff831681526001600160a01b0385811660208301528716818301526060810186905290519192507f2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8919081900360800190a150505050565b6000546001600160a01b031633146103be5760405162461bcd60e51b815260040161026c906115e1565b61027d6000610b41565b6000546001600160a01b031633146103f25760405162461bcd60e51b815260040161026c906115e1565b61027d610b91565b600054600160a01b900460ff1661044a5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015260640161026c565b6000546001600160a01b031633146104745760405162461bcd60e51b815260040161026c906115e1565b6104886001600160a01b0383163383610bf6565b5050565b6001818154811061049c57600080fd5b600091825260209091200154905081565b600360205260009081526040902080546104c6906116b2565b80601f01602080910402602001604051908101604052809291908181526020018280546104f2906116b2565b801561053f5780601f106105145761010080835404028352916020019161053f565b820191906000526020600020905b81548152906001019060200180831161052257829003601f168201915b505050505081565b60405163a310624f60e01b81523360048201819052906000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a310624f9060240160206040518083038186803b1580156105ac57600080fd5b505afa1580156105c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e49190611418565b905060018160038111156105fa576105fa61171e565b146106405760405162461bcd60e51b81526020600482015260166024820152752737ba103ab73137b73232b2103b30b634b230ba37b960511b604482015260640161026c565b6001600160a01b03821660009081526003602052604081208054610663906116b2565b80601f016020809104026020016040519081016040528092919081815260200182805461068f906116b2565b80156106dc5780601f106106b1576101008083540402835291602001916106dc565b820191906000526020600020905b8154815290600101906020018083116106bf57829003601f168201915b505050506001600160a01b038516600090815260036020526040902091925061070891905086866111a4565b50826001600160a01b03167f8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb482878760405161074693929190611587565b60405180910390a25050505050565b600054600160a01b900460ff161561077f5760405162461bcd60e51b815260040161026c906115b7565b60405163453a6aff60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638a74d5fe906107d19087908790879087906004016114c3565b60206040518083038186803b1580156107e957600080fd5b505afa1580156107fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610821919061131a565b50600061086385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c2b92505050565b80516001600160a01b0390811660009081526002602090815260408083208286015190941683529290528181205491830151929350916108a3919061166f565b9050600081116108f55760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f20776974686472617700000000000000604482015260640161026c565b60408083015183516001600160a01b039081166000908152600260209081528482208188018051851684529152939020919091558351915161093a9291169083610bf6565b8151602080840151604080516001600160a01b039485168152939091169183019190915281018290527f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9060600160405180910390a1505050505050565b6000546001600160a01b031633146109c25760405162461bcd60e51b815260040161026c906115e1565b6001600160a01b038116610a275760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161026c565b610a3081610b41565b50565b600054600160a01b900460ff16610a835760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015260640161026c565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6040516001600160a01b0380851660248301528316604482015260648101829052610b3b9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610d06565b50505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff1615610bbb5760405162461bcd60e51b815260040161026c906115b7565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610ab33390565b6040516001600160a01b038316602482015260448101829052610c2690849063a9059cbb60e01b90606401610b04565b505050565b604080516060810182526000808252602080830182905282840182905283518085019094528184528301849052909190805b60208301515183511015610cfe57610c7483610dd8565b90925090508160011415610ca357610c93610c8e84610e12565b610ecf565b6001600160a01b03168452610c5d565b8160021415610ccb57610cb8610c8e84610e12565b6001600160a01b03166020850152610c5d565b8160031415610cef57610ce5610ce084610e12565b610ee0565b6040850152610c5d565b610cf98382610f17565b610c5d565b505050919050565b6000610d5b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610f899092919063ffffffff16565b805190915015610c265780806020019051810190610d79919061131a565b610c265760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161026c565b6000806000610de684610fa2565b9050610df360088261162e565b9250806007166005811115610e0a57610e0a61171e565b915050915091565b60606000610e1f83610fa2565b90506000818460000151610e339190611616565b9050836020015151811115610e4757600080fd5b8167ffffffffffffffff811115610e6057610e60611734565b6040519080825280601f01601f191660200182016040528015610e8a576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015610ec4578181015183820152610ebd602082611616565b9050610ea2565b505050935250919050565b6000610eda82611024565b92915050565b6000602082511115610ef157600080fd5b6020820151905081516020610f06919061166f565b610f11906008611650565b1c919050565b6000816005811115610f2b57610f2b61171e565b1415610f3a57610c2682610fa2565b6002816005811115610f4e57610f4e61171e565b14156100ea576000610f5f83610fa2565b90508083600001818151610f739190611616565b90525060208301515183511115610c2657600080fd5b6060610f988484600085611043565b90505b9392505050565b602080820151825181019091015160009182805b600a81101561101e5783811a9150610fcf816007611650565b82607f16901b85179450816080166000141561100c57610ff0816001611616565b86518790610fff908390611616565b9052509395945050505050565b80611016816116ed565b915050610fb6565b50600080fd5b6000815160141461103457600080fd5b5060200151600160601b900490565b6060824710156110a45760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161026c565b843b6110f25760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161026c565b600080866001600160a01b0316858760405161110e91906114a7565b60006040518083038185875af1925050503d806000811461114b576040519150601f19603f3d011682016040523d82523d6000602084013e611150565b606091505b509150915061116082828661116b565b979650505050505050565b6060831561117a575081610f9b565b82511561118a5782518084602001fd5b8160405162461bcd60e51b815260040161026c9190611574565b8280546111b0906116b2565b90600052602060002090601f0160209004810192826111d25760008555611218565b82601f106111eb5782800160ff19823516178555611218565b82800160010185558215611218579182015b828111156112185782358255916020019190600101906111fd565b50611224929150611228565b5090565b5b808211156112245760008155600101611229565b80356001600160a01b038116811461125457600080fd5b919050565b60008083601f84011261126b57600080fd5b50813567ffffffffffffffff81111561128357600080fd5b60208301915083602082850101111561129b57600080fd5b9250929050565b6000602082840312156112b457600080fd5b610f9b8261123d565b600080604083850312156112d057600080fd5b6112d98361123d565b91506112e76020840161123d565b90509250929050565b6000806040838503121561130357600080fd5b61130c8361123d565b946020939093013593505050565b60006020828403121561132c57600080fd5b81518015158114610f9b57600080fd5b6000806020838503121561134f57600080fd5b823567ffffffffffffffff81111561136657600080fd5b61137285828601611259565b90969095509350505050565b6000806000806040858703121561139457600080fd5b843567ffffffffffffffff808211156113ac57600080fd5b6113b888838901611259565b909650945060208701359150808211156113d157600080fd5b818701915087601f8301126113e557600080fd5b8135818111156113f457600080fd5b8860208260051b850101111561140957600080fd5b95989497505060200194505050565b60006020828403121561142a57600080fd5b815160048110610f9b57600080fd5b60006020828403121561144b57600080fd5b5035919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008151808452611493816020860160208601611686565b601f01601f19169290920160200192915050565b600082516114b9818460208701611686565b9190910192915050565b6040815260006114d7604083018688611452565b602083820381850152818583528183019050818660051b8401018760005b8881101561156457858303601f190184528135368b9003601e1901811261151b57600080fd5b8a01803567ffffffffffffffff81111561153457600080fd5b8036038c131561154357600080fd5b6115508582898501611452565b9587019594505050908401906001016114f5565b50909a9950505050505050505050565b602081526000610f9b602083018461147b565b60408152600061159a604083018661147b565b82810360208401526115ad818587611452565b9695505050505050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6000821982111561162957611629611708565b500190565b60008261164b57634e487b7160e01b600052601260045260246000fd5b500490565b600081600019048311821515161561166a5761166a611708565b500290565b60008282101561168157611681611708565b500390565b60005b838110156116a1578181015183820152602001611689565b83811115610b3b5750506000910152565b600181811c908216806116c657607f821691505b602082108114156116e757634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561170157611701611708565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212206ef303cc93b4d427d72a3ab740ae09e5ff940eecbdfa2708729c99cc837e0b6964736f6c63430008070033",
}

// SGNABI is the input ABI used to generate the binding from.
// Deprecated: Use SGNMetaData.ABI instead.
var SGNABI = SGNMetaData.ABI

// SGNBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SGNMetaData.Bin instead.
var SGNBin = SGNMetaData.Bin

// DeploySGN deploys a new Ethereum contract, binding an instance of SGN to it.
func DeploySGN(auth *bind.TransactOpts, backend bind.ContractBackend, _dpos common.Address) (common.Address, *types.Transaction, *SGN, error) {
	parsed, err := SGNMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SGNBin), backend, _dpos)
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

// Dpos is a free data retrieval call binding the contract method 0x8f8451b2.
//
// Solidity: function dpos() view returns(address)
func (_SGN *SGNCaller) Dpos(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "dpos")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dpos is a free data retrieval call binding the contract method 0x8f8451b2.
//
// Solidity: function dpos() view returns(address)
func (_SGN *SGNSession) Dpos() (common.Address, error) {
	return _SGN.Contract.Dpos(&_SGN.CallOpts)
}

// Dpos is a free data retrieval call binding the contract method 0x8f8451b2.
//
// Solidity: function dpos() view returns(address)
func (_SGN *SGNCallerSession) Dpos() (common.Address, error) {
	return _SGN.Contract.Dpos(&_SGN.CallOpts)
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

// WhitelistMetaData contains all meta data concerning the Whitelist contract.
var WhitelistMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
