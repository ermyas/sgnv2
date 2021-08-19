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
	ValAddr       common.Address
	Shares        *big.Int
	Undelegations []DPoSUndelegation
}

// DPoSUndelegation is an auto generated low-level Go binding around an user-defined struct.
type DPoSUndelegation struct {
	Shares        *big.Int
	CreationBlock *big.Int
}

// DPoSMetaData contains all meta data concerning the DPoS contract.
var DPoSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_governProposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_governVoteTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"record\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ConfirmParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"record\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CreateParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"tokenDiff\",\"type\":\"int256\"}],\"name\":\"DelegationUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPoolSize\",\"type\":\"uint256\"}],\"name\":\"MiningPoolContribution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPool\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmt\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashAmtCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"name\":\"ValidatorParamsUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDPoS.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGovern.VoteType\",\"name\":\"voteType\",\"type\":\"uint8\"}],\"name\":\"VoteParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"UIntStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bondedValAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondedValTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"celerToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rewardRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"completeUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"confirmParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"confirmUnbondedValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToMiningPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_record\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"createParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDPoS.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structDPoS.DelegatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDPoS.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structDPoS.DelegatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getParamProposalVote\",\"outputs\":[{\"internalType\":\"enumGovern.VoteType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_record\",\"type\":\"uint256\"}],\"name\":\"getUIntValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"enumDPoS.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"initializeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBondedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextParamProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paramProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"record\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_slashRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"}],\"name\":\"updateMinSelfDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"valAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"enumDPoS.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.VoteType\",\"name\":\"_vote\",\"type\":\"uint8\"}],\"name\":\"voteParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162005370380380620053708339810160408190526200003491620001b8565b86868686868686620000463362000168565b6000805460ff60a01b19168155600280546001600160a01b039990991661010002610100600160a81b03199099169890981790975560036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff959095557fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c939093557fc3a24b0501bd2c13a7e57f2db4369ec4c223447539fc0724a9d55ac4a06ebd4d919091557fcbc4e5fb02c3d1de23a9f1e014b4d2ee5aeaea9505df5e855c9210bf472495af557f83ec6a1f0257b830b5e016457c9cf1435391bf56cc98f369a58a54fe937724655560059091527f405aad32e1adbac89bb7f176e338b8fc6e994ca210c9bb7bdca249b46594225055506200022395505050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600080600080600080600060e0888a031215620001d457600080fd5b87516001600160a01b0381168114620001ec57600080fd5b602089015160408a015160608b015160808c015160a08d015160c0909d0151949e939d50919b909a50909850965090945092505050565b61513d80620002336000396000f3fe60806040526004361061031d5760003560e01c806364c66395116101ab57806392bb243c116100f7578063c8f9f98411610095578063eecefef81161006f578063eecefef8146109bd578063f2fde38b146109ea578063f8df0dc514610a0a578063fa52c7d814610a2a57600080fd5b8063c8f9f9841461095a578063cdfb2b4e14610993578063d6b0f484146109a857600080fd5b8063acc62ccf116100d1578063acc62ccf146108df578063af062e6f146108ff578063b4f7fa3414610915578063c6c21e9d1461093557600080fd5b806392bb243c14610866578063934a18ec14610886578063a310624f146108a657600080fd5b80637e5fb8f3116101645780638a74d5fe1161013e5780638a74d5fe146107cf5780638da5cb5b146107ef5780638dc2336d1461082157806390e360f81461083657600080fd5b80637e5fb8f3146107365780638456cb59146107a557806389f9aab5146107ba57600080fd5b806364c663951461067c57806364ed600a146106a957806366666aa9146106d6578063715018a6146106ec57806371bc0216146107015780637a3ba4ad1461072157600080fd5b80633773e4891161026a578063473849bd1161022357806351fb012d116101fd57806351fb012d146105f6578063581c53c5146106105780635c975abb1461063d5780635e593eff1461065c57600080fd5b8063473849bd146105965780634b7dba6b146105b65780634d99dd16146105d657600080fd5b80633773e489146104c9578063386c024a146104f35780633985c4e6146105085780633af32abf146105285780633be88c2a146105615780633f4ba83a1461058157600080fd5b806322da7927116102d7578063291d9549116102b1578063291d9549146104475780633090c0e914610467578063313019bb1461048757806336f1635f146104b457600080fd5b806322da7927146103fc57806324b9bcc01461041257806325ed6b351461042757600080fd5b8062fa3d5014610329578063026e402b1461034b57806310154bad1461036b578063145aa1161461038b5780631cfe4f0b146103ab5780631e6f3d8a146103cf57600080fd5b3661032457005b600080fd5b34801561033557600080fd5b50610349610344366004614bfa565b610aa9565b005b34801561035757600080fd5b50610349610366366004614a73565b610b91565b34801561037757600080fd5b50610349610386366004614a25565b610da2565b34801561039757600080fd5b506103496103a6366004614bfa565b610dd8565b3480156103b757600080fd5b506008545b6040519081526020015b60405180910390f35b3480156103db57600080fd5b506103bc6103ea366004614a25565b600b6020526000908152604090205481565b34801561040857600080fd5b506103bc60055481565b34801561041e57600080fd5b50610349610e6e565b34801561043357600080fd5b50610349610442366004614c36565b610ea7565b34801561045357600080fd5b50610349610462366004614a25565b610f2b565b34801561047357600080fd5b50610349610482366004614c6a565b610f5e565b34801561049357600080fd5b506104a76104a2366004614a25565b6110a8565b6040516103c69190614d67565b3480156104c057600080fd5b506103496112e9565b3480156104d557600080fd5b50600c546104e39060ff1681565b60405190151581526020016103c6565b3480156104ff57600080fd5b506103bc61155e565b34801561051457600080fd5b50610349610523366004614abf565b61158b565b34801561053457600080fd5b506104e3610543366004614a25565b6001600160a01b031660009081526001602052604090205460ff1690565b34801561056d57600080fd5b5061034961057c366004614c6a565b611bbd565b34801561058d57600080fd5b50610349611e12565b3480156105a257600080fd5b506103496105b1366004614a25565b611e46565b3480156105c257600080fd5b506103496105d1366004614bfa565b6120ca565b3480156105e257600080fd5b506103496105f1366004614a73565b61217f565b34801561060257600080fd5b506002546104e39060ff1681565b34801561061c57600080fd5b5061063061062b366004614c13565b612509565b6040516103c69190614dc9565b34801561064957600080fd5b50600054600160a01b900460ff166104e3565b34801561066857600080fd5b50610349610677366004614bfa565b612539565b34801561068857600080fd5b506103bc610697366004614bfa565b60009081526003602052604090205490565b3480156106b557600080fd5b506103bc6106c4366004614bfa565b60036020526000908152604090205481565b3480156106e257600080fd5b506103bc60065481565b3480156106f857600080fd5b506103496126a5565b34801561070d57600080fd5b5061034961071c366004614a25565b6126d9565b34801561072d57600080fd5b506103496127f7565b34801561074257600080fd5b50610793610751366004614bfa565b60046020819052600091825260409091208054600182015460028301546003840154948401546005909401546001600160a01b03909316949193909260ff1686565b6040516103c696959493929190614d1a565b3480156107b157600080fd5b5061034961282d565b3480156107c657600080fd5b506009546103bc565b3480156107db57600080fd5b506104e36107ea366004614b83565b61285f565b3480156107fb57600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016103c6565b34801561082d57600080fd5b506103bc612b32565b34801561084257600080fd5b506104e3610851366004614bfa565b600d6020526000908152604090205460ff1681565b34801561087257600080fd5b50610809610881366004614bfa565b612c31565b34801561089257600080fd5b506103496108a1366004614bfa565b612c5b565b3480156108b257600080fd5b506106306108c1366004614a25565b6001600160a01b03166000908152600a602052604090205460ff1690565b3480156108eb57600080fd5b506108096108fa366004614bfa565b612d6d565b34801561090b57600080fd5b506103bc60075481565b34801561092157600080fd5b506104e3610930366004614a25565b612d7d565b34801561094157600080fd5b506002546108099061010090046001600160a01b031681565b34801561096657600080fd5b506103bc610975366004614a25565b6001600160a01b03166000908152600a602052604090206001015490565b34801561099f57600080fd5b50610349612db5565b3480156109b457600080fd5b50610349612df1565b3480156109c957600080fd5b506109dd6109d8366004614a40565b612e2a565b6040516103c69190614eef565b3480156109f657600080fd5b50610349610a05366004614a25565b612f99565b348015610a1657600080fd5b50610349610a25366004614abf565b613031565b348015610a3657600080fd5b50610a94610a45366004614a25565b600a6020526000908152604090208054600182015460028301546003840154600485015460068601546007870154600888015460099098015460ff909716979596949593949293919290919089565b6040516103c699989796959493929190614ddc565b336000818152600a6020526040812090815460ff166003811115610acf57610acf61509f565b1415610af65760405162461bcd60e51b8152600401610aed90614eb8565b60405180910390fd5b612710831115610b3b5760405162461bcd60e51b815260206004820152601060248201526f496e76616c6964206e6577207261746560801b6044820152606401610aed565b60088101839055600981015460408051918252602082018590526001600160a01b038416917fb7f73ede33f427fed817c4721ec2ae9f56e906564766ad98d79f291f8bf0b79991015b60405180910390a2505050565b600054600160a01b900460ff1615610bbb5760405162461bcd60e51b8152600401610aed90614e59565b33670de0b6b3a7640000821015610c145760405162461bcd60e51b815260206004820152601860248201527f4d696e696d616c20616d6f756e7420697320312043454c5200000000000000006044820152606401610aed565b6001600160a01b0383166000908152600a6020526040812090815460ff166003811115610c4357610c4361509f565b1415610c615760405162461bcd60e51b8152600401610aed90614eb8565b6000610c768483600101548460020154613265565b6001600160a01b0384166000908152600584016020526040812080549293509183918391610ca5908490614f5f565b9250508190555081836002016000828254610cc09190614f5f565b9250508190555084836001016000828254610cdb9190614f5f565b9091555060029050835460ff166003811115610cf957610cf961509f565b1415610d24578460076000828254610d119190614f5f565b90915550506001830154610d2490613292565b600254610d419061010090046001600160a01b0316853088613394565b6001830154815460408051928352602083019190915281018690526001600160a01b0380861691908816907f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea906060015b60405180910390a3505050505050565b6000546001600160a01b03163314610dcc5760405162461bcd60e51b8152600401610aed90614e83565b610dd581613405565b50565b600054600160a01b900460ff16610e285760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610aed565b6000546001600160a01b03163314610e525760405162461bcd60e51b8152600401610aed90614e83565b600254610dd59061010090046001600160a01b031633836134c0565b6000546001600160a01b03163314610e985760405162461bcd60e51b8152600401610aed90614e83565b600c805460ff19166001179055565b336000818152600a602052604090205460029060ff166003811115610ece57610ece61509f565b14610f1b5760405162461bcd60e51b815260206004820181905260248201527f43616c6c6572206973206e6f74206120626f6e6465642076616c696461746f726044820152606401610aed565b610f268382846134f0565b505050565b6000546001600160a01b03163314610f555760405162461bcd60e51b8152600401610aed90614e83565b610dd581613695565b600554600081815260046020526040902090610f7b906001614f5f565b60055560036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff5481546001600160a01b03191633908117835560018084018390556000527fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c54909190610ff19043614f5f565b600284810191909155600384018690556004840185905560058401805460ff1916600117905554611033906001600160a01b0361010090910416833084613394565b7f40109a070319d6004f4e4b31dba4b605c97bd3474d49865158f55fe093e3b33960016005546110639190614fb8565b6002850154604080519283526001600160a01b0386166020840152820184905260608201526080810187905260a0810186905260c00160405180910390a15050505050565b6008546060906000906001600160401b038111156110c8576110c86150e1565b60405190808252806020026020018201604052801561110157816020015b6110ee6148e5565b8152602001906001900390816110e65790505b5090506000805b60085463ffffffff82161015611216576000600a600060088463ffffffff1681548110611137576111376150cb565b60009182526020808320909101546001600160a01b0390811684528382019490945260409283018220938a1682526005909301909252902080549091501580156111955750600281015463ffffffff80821664010000000090920416145b15611203576111d160088363ffffffff16815481106111b6576111b66150cb565b6000918252602090912001546001600160a01b031687612e2a565b848363ffffffff16815181106111e9576111e96150cb565b602002602001018190525082806111ff90615048565b9350505b508061120e81615048565b915050611108565b5060008163ffffffff166001600160401b03811115611237576112376150e1565b60405190808252806020026020018201604052801561127057816020015b61125d6148e5565b8152602001906001900390816112555790505b50905060005b8263ffffffff168163ffffffff1610156112e057838163ffffffff16815181106112a2576112a26150cb565b6020026020010151828263ffffffff16815181106112c2576112c26150cb565b602002602001018190525080806112d890615048565b915050611276565b50949350505050565b336000818152600a602052604090206001815460ff1660038111156113105761131061509f565b148061133157506003815460ff16600381111561132f5761132f61509f565b145b61137d5760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076616c696461746f722073746174757300000000000000006044820152606401610aed565b80600601544310156113ca5760405162461bcd60e51b8152602060048201526016602482015275109bdb9908189b1bd8dac81b9bdd081c995858da195960521b6044820152606401610aed565b6113d382613740565b61141f5760405162461bcd60e51b815260206004820152601f60248201527f4e6f74206d656574206d696e20746f6b656e20726571756972656d656e7473006044820152606401610aed565b600061142b6003610697565b60095490915081111561144157610f26836137c0565b6000196000805b838110156114f65782600a600060098481548110611468576114686150cb565b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015410156114e457809150600a6000600983815481106114b0576114b06150cb565b60009182526020808320909101546001600160a01b031683528201929092526040019020600101549250826114e4576114f6565b806114ee8161502d565b915050611448565b50818460010154116115405760405162461bcd60e51b8152602060048201526013602482015272496e73756666696369656e7420746f6b656e7360681b6044820152606401610aed565b61154a8582613814565b6115578460010154613292565b5050505050565b6000600360075460026115719190614f99565b61157b9190614f77565b611586906001614f5f565b905090565b600054600160a01b900460ff16156115b55760405162461bcd60e51b8152600401610aed90614e59565b600c5460ff16156115fc5760405162461bcd60e51b815260206004820152601160248201527014db185cda081a5cc8191a5cd8589b1959607a1b6044820152606401610aed565b600061163d85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061388f92505050565b905061168585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506107ea9250869150879050614ff4565b5080606001516001600160401b031643106116d25760405162461bcd60e51b815260206004820152600d60248201526c14db185cda08195e1c1a5c9959609a1b6044820152606401610aed565b6305f5e10081604001516001600160401b0316111561172a5760405162461bcd60e51b815260206004820152601460248201527324b73b30b634b21039b630b9b4103330b1ba37b960611b6044820152606401610aed565b6020808201516001600160401b03166000908152600d909152604090205460ff161561178b5760405162461bcd60e51b815260206004820152601060248201526f5573656420736c617368206e6f6e636560801b6044820152606401610aed565b6020808201516001600160401b03166000908152600d82526040808220805460ff1916600190811790915584516001600160a01b0381168452600a909452912090815460ff1660038111156117e2576117e261509f565b14156118265760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881d5b989bdd5b991959606a1b6044820152606401610aed565b6000633b9aca0084604001516001600160401b0316836001015461184a9190614f99565b6118549190614f77565b90508082600101600082825461186a9190614fb8565b9091555060029050825460ff1660038111156118885761188861509f565b14156119055780600760008282546118a09190614fb8565b909155505060808401516001600160401b03161515806118c657506118c483613740565b155b15611905576118d483613afc565b60808401516001600160401b0316156119055760808401516118ff906001600160401b031643614f5f565b60068301555b60006001600160a01b0316836001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea846001015460008561194c9061506c565b6040805193845260208401929092529082015260600160405180910390a36000633b9aca0085604001516001600160401b0316846003015461198e9190614f99565b6119989190614f77565b9050808360030160008282546119ae9190614fb8565b909155506119be90508183614f5f565b91506000805b8660a0015151811015611af85760008760a0015182815181106119e9576119e96150cb565b60200260200101519050806020015183611a039190614f5f565b81519093506001600160a01b0316611a76576020810151600254611a38916101009091046001600160a01b03169033906134c0565b60208082015160405190815233917fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3910160405180910390a2611ae5565b80516020820151600254611a99926101009091046001600160a01b0316916134c0565b80600001516001600160a01b03167fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a38260200151604051611adc91815260200190565b60405180910390a25b5080611af08161502d565b9150506119c4565b5080831015611b3e5760405162461bcd60e51b8152602060048201526012602482015271496e76616c696420636f6c6c6563746f727360701b6044820152606401610aed565b611b488184614fb8565b60066000828254611b599190614f5f565b9091555050602080870151604080516001600160401b0390921682529181018590526001600160a01b038716917f10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008910160405180910390a250505050505050505050565b600054600160a01b900460ff1615611be75760405162461bcd60e51b8152600401610aed90614e59565b60025460ff1615611c51573360009081526001602052604090205460ff16611c515760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f742077686974656c6973746564000000000000006044820152606401610aed565b336000818152600a6020526040812090815460ff166003811115611c7757611c7761509f565b14611cc45760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f7220697320696e697469616c697a656400000000000000006044820152606401610aed565b612710831115611d165760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636f6d6d697373696f6e20726174650000000000000000006044820152606401610aed565b670de0b6b3a7640000841015611d6e5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206d696e2073656c662064656c65676174696f6e00000000006044820152606401610aed565b805460ff191660019081178255600982018590556008808301859055805491820181556000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180546001600160a01b0384166001600160a01b0319909116811790915560408051868152602081018690527fb7f73ede33f427fed817c4721ec2ae9f56e906564766ad98d79f291f8bf0b799910160405180910390a250505050565b6000546001600160a01b03163314611e3c5760405162461bcd60e51b8152600401610aed90614e83565b611e44613c59565b565b6001600160a01b0381166000908152600a602052604081203391815460ff166003811115611e7657611e7661509f565b1415611e945760405162461bcd60e51b8152600401610aed90614eb8565b6001600160a01b0382166000908152600582016020526040812090611eb96002610697565b905060006001845460ff166003811115611ed557611ed561509f565b60028501549114915063ffffffff1660005b600285015463ffffffff64010000000090910481169083161015611f99578280611f37575063ffffffff821660009081526001808701602052604090912001544390611f34908690614f5f565b11155b15611f825763ffffffff82166000908152600186016020526040902054611f5e9082614f5f565b63ffffffff8316600090815260018088016020526040822082815501559050611f87565b611f99565b81611f9181615048565b925050611ee7565b60028501805463ffffffff191663ffffffff84161790558061200b5760405162461bcd60e51b815260206004820152602560248201527f6e6f20756e64656c65676174696f6e20726561647920746f20626520636f6d706044820152641b195d195960da1b6064820152608401610aed565b60006120208288600301548960040154613cf6565b9050818760040160008282546120369190614fb8565b92505081905550808760030160008282546120519190614fb8565b90915550506002546120729061010090046001600160a01b031689836134c0565b876001600160a01b0316896001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c836040516120b791815260200190565b60405180910390a3505050505050505050565b600054600160a01b900460ff16156120f45760405162461bcd60e51b8152600401610aed90614e59565b6000339050816006600082825461210b9190614f5f565b909155505060025461212d9061010090046001600160a01b0316823085613394565b806001600160a01b03167f97e19c4040b6c46d4275e0c4fea68f8f92c81138372ffdb089932c211938f76583600654604051612173929190918252602082015260400190565b60405180910390a25050565b33670de0b6b3a76400008210156121d85760405162461bcd60e51b815260206004820152601960248201527f4d696e696d616c20616d6f756e742069732031207368617265000000000000006044820152606401610aed565b6001600160a01b0383166000908152600a6020526040812090815460ff1660038111156122075761220761509f565b14156122255760405162461bcd60e51b8152600401610aed90614eb8565b600061223a8483600101548460020154613cf6565b6001600160a01b0384166000908152600584016020526040812080549293509186918391612269908490614fb8565b92505081905550848360020160008282546122849190614fb8565b925050819055508183600101600082825461229f9190614fb8565b9091555060019050835460ff1660038111156122bd576122bd61509f565b1415612324576002546122df9061010090046001600160a01b031685846134c0565b836001600160a01b0316866001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c84604051610d9291815260200190565b6002835460ff16600381111561233c5761233c61509f565b14156123705781600760008282546123549190614fb8565b90915550612363905086613740565b6123705761237086613afc565b6002810154600a906123939063ffffffff80821691640100000000900416614fcf565b63ffffffff16106123e65760405162461bcd60e51b815260206004820152601f60248201527f457863656564206d617820756e64656c65676174696f6e20656e7472696573006044820152606401610aed565b60006123fb8385600301548660040154613265565b9050808460040160008282546124119190614f5f565b925050819055508284600301600082825461242c9190614f5f565b909155505060028201805463ffffffff6401000000009182900481166000908152600180870160205260409091208581554391810191909155835490939290041690600461247983615048565b91906101000a81548163ffffffff021916908363ffffffff16021790555050856001600160a01b0316886001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea87600101548660000154886124e19061506c565b6040805193845260208401929092529082015260600160405180910390a35050505050505050565b60008281526004602090815260408083206001600160a01b038516845260060190915290205460ff165b92915050565b336000818152600a6020526040812090815460ff16600381111561255f5761255f61509f565b141561257d5760405162461bcd60e51b8152600401610aed90614eb8565b670de0b6b3a76400008310156125d55760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206d696e2073656c662064656c65676174696f6e00000000006044820152606401610aed565b8060090154831015612657576002815460ff1660038111156125f9576125f961509f565b141561263d5760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881a5cc8189bdb991959606a1b6044820152606401610aed565b6126476005610697565b6126519043614f5f565b60068201555b6009810183905560088101546040516001600160a01b038416917fb7f73ede33f427fed817c4721ec2ae9f56e906564766ad98d79f291f8bf0b79991610b8491878252602082015260400190565b6000546001600160a01b031633146126cf5760405162461bcd60e51b8152600401610aed90614e83565b611e446000613d0f565b6001600160a01b0381166000908152600a602052604090206003815460ff1660038111156127095761270961509f565b146127565760405162461bcd60e51b815260206004820152601760248201527f56616c696461746f72206e6f7420756e626f6e64696e670000000000000000006044820152606401610aed565b80600701544310156127aa5760405162461bcd60e51b815260206004820152601760248201527f556e626f6e642074696d65206e6f7420726561636865640000000000000000006044820152606401610aed565b805460ff191660019081178255600060078301555b6040516001600160a01b038416907fd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a90600090a35050565b6000546001600160a01b031633146128215760405162461bcd60e51b8152600401610aed90614e83565b600c805460ff19169055565b6000546001600160a01b031633146128575760405162461bcd60e51b8152600401610aed90614e83565b611e44613d5f565b6000806128c084805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050600083516001600160401b038111156128dd576128dd6150e1565b604051908082528060200260200182016040528015612906578160200160208202803683370190505b509050600080805b8651811015612ad85761294387828151811061292c5761292c6150cb565b602002602001015186613dc490919063ffffffff16565b848281518110612955576129556150cb565b60200260200101906001600160a01b031690816001600160a01b031681525050816001600160a01b0316848281518110612991576129916150cb565b60200260200101516001600160a01b0316116129ef5760405162461bcd60e51b815260206004820152601e60248201527f5369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610aed565b838181518110612a0157612a016150cb565b6020026020010151915060026003811115612a1e57612a1e61509f565b600a6000868481518110612a3457612a346150cb565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff166003811115612a6c57612a6c61509f565b14612a7657612ac6565b600a6000858381518110612a8c57612a8c6150cb565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206001015483612ac39190614f5f565b92505b80612ad08161502d565b91505061290e565b50612ae161155e565b821015612b255760405162461bcd60e51b8152602060048201526012602482015271145d5bdc9d5b481b9bdd081c995858da195960721b6044820152606401610aed565b5060019695505050505050565b600080600a60006009600081548110612b4d57612b4d6150cb565b60009182526020808320909101546001600160a01b03168352820192909252604001902060019081015491505b600954811015612c2b5781600a600060098481548110612b9c57612b9c6150cb565b60009182526020808320909101546001600160a01b031683528201929092526040019020600101541015612c1957600a600060098381548110612be157612be16150cb565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154915081612c195760009250505090565b80612c238161502d565b915050612b7a565b50919050565b60088181548110612c4157600080fd5b6000918252602090912001546001600160a01b0316905081565b6000805b60095463ffffffff82161015612d23576001612ca88460098463ffffffff1681548110612c8e57612c8e6150cb565b6000918252602090912001546001600160a01b0316612509565b6003811115612cb957612cb961509f565b1415612d1157600a600060098363ffffffff1681548110612cdc57612cdc6150cb565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154612d0e9083614f5f565b91505b80612d1b81615048565b915050612c5f565b506000612d2e61155e565b821015905080612d63576000838152600460205260408120600101546006805491929091612d5d908490614f5f565b90915550505b610f268382613e68565b60098181548110612c4157600080fd5b600060026001600160a01b0383166000908152600a602052604090205460ff166003811115612dae57612dae61509f565b1492915050565b6000546001600160a01b03163314612ddf5760405162461bcd60e51b8152600401610aed90614e83565b611e446002805460ff19166001179055565b6000546001600160a01b03163314612e1b5760405162461bcd60e51b8152600401610aed90614e83565b611e446002805460ff19169055565b612e326148e5565b6001600160a01b038084166000908152600a602090815260408083209386168352600590930190529081206002810154909190612e809063ffffffff80821691640100000000900416614fcf565b63ffffffff1690506000816001600160401b03811115612ea257612ea26150e1565b604051908082528060200260200182016040528015612ee757816020015b6040805180820190915260008082526020820152815260200190600190039081612ec05790505b50905060005b82811015612f6f5760028401546001850190600090612f129063ffffffff1684614f5f565b815260200190815260200160002060405180604001604052908160008201548152602001600182015481525050828281518110612f5157612f516150cb565b60200260200101819052508080612f679061502d565b915050612eed565b50604080516060810182526001600160a01b03881681529354602085015283015250905092915050565b6000546001600160a01b03163314612fc35760405162461bcd60e51b8152600401610aed90614e83565b6001600160a01b0381166130285760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610aed565b610dd581613d0f565b600054600160a01b900460ff161561305b5760405162461bcd60e51b8152600401610aed90614e59565b6130a184848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506107ea9250859150869050614ff4565b5060006130e385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613fdb92505050565b80516001600160a01b03166000908152600b602090815260408220549083015192935090916131129190614fb8565b9050600081116131545760405162461bcd60e51b815260206004820152600d60248201526c139bc81b995dc81c995dd85c99609a1b6044820152606401610aed565b8060065410156131b55760405162461bcd60e51b815260206004820152602660248201527f52657761726420706f6f6c20697320736d616c6c6572207468616e206e6577206044820152651c995dd85c9960d21b6064820152608401610aed565b60208083015183516001600160a01b03166000908152600b9092526040822055600680548392906131e7908490614fb8565b9091555050815160025461320b916101009091046001600160a01b031690836134c0565b81600001516001600160a01b03167ff01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e273174382600654604051613255929190918252602082015260400190565b60405180910390a2505050505050565b60008261327357508261328b565b8261327e8386614f99565b6132889190614f77565b90505b9392505050565b60095460028114806132a45750806003145b1561331a576132b161155e565b82106133165760405162461bcd60e51b815260206004820152602e60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201526d2071756f72756d20746f6b656e7360901b6064820152608401610aed565b5050565b60038111156133165760036007546133329190614f77565b82106133165760405162461bcd60e51b815260206004820152602b60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201526a20312f3320746f6b656e7360a81b6064820152608401610aed565b6040516001600160a01b03808516602483015283166044820152606481018290526133ff9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152614082565b50505050565b6001600160a01b03811660009081526001602052604090205460ff16156134645760405162461bcd60e51b8152602060048201526013602482015272185b1c9958591e481dda1a5d195b1a5cdd1959606a1b6044820152606401610aed565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f91015b60405180910390a150565b6040516001600160a01b038316602482015260448101829052610f2690849063a9059cbb60e01b906064016133c8565b60008381526004602052604090206001600582015460ff1660028111156135195761351961509f565b146135605760405162461bcd60e51b8152602060048201526017602482015276496e76616c69642070726f706f73616c2073746174757360481b6044820152606401610aed565b806002015443106135aa5760405162461bcd60e51b8152602060048201526014602482015273159bdd1948191958591b1a5b99481c185cdcd95960621b6044820152606401610aed565b6001600160a01b038316600090815260068201602052604081205460ff1660038111156135d9576135d961509f565b146136185760405162461bcd60e51b815260206004820152600f60248201526e159bdd195c881a185cc81d9bdd1959608a1b6044820152606401610aed565b6001600160a01b03831660009081526006820160205260409020805483919060ff1916600183600381111561364f5761364f61509f565b02179055507f06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d6584848460405161368793929190614f02565b60405180910390a150505050565b6001600160a01b03811660009081526001602052604090205460ff166136ef5760405162461bcd60e51b815260206004820152600f60248201526e1b9bdd081dda1a5d195b1a5cdd1959608a1b6044820152606401610aed565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b691016134b5565b6001600160a01b0381166000908152600a602090815260408083206001810154600582019093529083205460028201549192918491613780918490613cf6565b905061378c6004610697565b82101561379e57506000949350505050565b82600901548110156137b557506000949350505050565b506001949350505050565b600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b0319166001600160a01b038316179055610dd581614154565b6138446009828154811061382a5761382a6150cb565b6000918252602090912001546001600160a01b03166141a2565b8160098281548110613858576138586150cb565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061331682614154565b6040805160c0810182526000808252602080830182905282840182905260608084018390526080840183905260a0840152835180850190945281845283018490529091906138de826006614204565b9050806006815181106138f3576138f36150cb565b60200260200101516001600160401b03811115613912576139126150e1565b60405190808252806020026020018201604052801561395757816020015b60408051808201909152600080825260208201528152602001906001900390816139305790505b508360a00181905250600081600681518110613975576139756150cb565b6020026020010181815250506000805b60208401515184511015613af35761399c846142bd565b909250905081600114156139cb576139bb6139b6856142f7565b6143b3565b6001600160a01b03168552613985565b81600214156139f0576139dd846143be565b6001600160401b03166020860152613985565b8160031415613a1557613a02846143be565b6001600160401b03166040860152613985565b8160041415613a3a57613a27846143be565b6001600160401b03166060860152613985565b8160051415613a5f57613a4c846143be565b6001600160401b03166080860152613985565b8160061415613ae457613a79613a74856142f7565b614440565b8560a0015184600681518110613a9157613a916150cb565b602002602001015181518110613aa957613aa96150cb565b602002602001018190525082600681518110613ac757613ac76150cb565b602002602001018051809190613adc9061502d565b905250613985565b613aee84826144da565b613985565b50505050919050565b600954600090613b0e90600190614fb8565b905060005b600954811015613c1957826001600160a01b031660098281548110613b3a57613b3a6150cb565b6000918252602090912001546001600160a01b03161415613c075781811015613bcb5760098281548110613b7057613b706150cb565b600091825260209091200154600980546001600160a01b039092169183908110613b9c57613b9c6150cb565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b6009805480613bdc57613bdc6150b5565b600082815260209020810160001990810180546001600160a01b0319169055019055610f26836141a2565b80613c118161502d565b915050613b13565b5060405162461bcd60e51b81526020600482015260146024820152732737ba103137b73232b2103b30b634b230ba37b960611b6044820152606401610aed565b600054600160a01b900460ff16613ca95760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610aed565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600081613d0457508261328b565b8161327e8486614f99565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff1615613d895760405162461bcd60e51b8152600401610aed90614e59565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613cd93390565b6000815160411415613df85760208201516040830151606084015160001a613dee8682858561454c565b9350505050612533565b815160401415613e205760208201516040830151613e178583836146f5565b92505050612533565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610aed565b60008281526004602052604090206001600582015460ff166002811115613e9157613e9161509f565b14613ed85760405162461bcd60e51b8152602060048201526017602482015276496e76616c69642070726f706f73616c2073746174757360481b6044820152606401610aed565b8060020154431015613f2c5760405162461bcd60e51b815260206004820152601960248201527f566f746520646561646c696e65206e6f742072656163686564000000000000006044820152606401610aed565b60058101805460ff191660021790558115613f845780546001820154600254613f68926001600160a01b036101009092048216929116906134c0565b6004810154600380830154600090815260209190915260409020555b600381015460048201546040805186815285151560208201529081019290925260608201527f106f43a560e53395081c0423504b476d1a2cfed9d56ff972bf77ae43ff7d4ba49060800160405180910390a1505050565b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b6020830151518351101561407a5761401d836142bd565b90925090508160011415614047576140376139b6846142f7565b6001600160a01b03168452614006565b816002141561406b5761406161405c846142f7565b61471f565b6020850152614006565b61407583826144da565b614006565b505050919050565b60006140d7826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166147569092919063ffffffff16565b805190915015610f2657808060200190518101906140f59190614a9d565b610f265760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610aed565b6001600160a01b0381166000908152600a60205260408120805460ff1916600217815560078082018390556001820154815492939092614195908490614f5f565b90915550600290506127bf565b6001600160a01b0381166000908152600a60205260409020805460ff191660031781556141cf6002610697565b6141d99043614f5f565b81600701819055508060010154600760008282546141f79190614fb8565b90915550600390506127bf565b8151606090614214836001614f5f565b6001600160401b0381111561422b5761422b6150e1565b604051908082528060200260200182016040528015614254578160200160208202803683370190505b5091506000805b602086015151865110156142b457614272866142bd565b8092508193505050600184838151811061428e5761428e6150cb565b602002602001018181516142a29190614f5f565b9052506142af86826144da565b61425b565b50509092525090565b60008060006142cb846143be565b90506142d8600882614f77565b92508060071660058111156142ef576142ef61509f565b915050915091565b60606000614304836143be565b905060008184600001516143189190614f5f565b905083602001515181111561432c57600080fd5b816001600160401b03811115614344576143446150e1565b6040519080825280601f01601f19166020018201604052801561436e576020820181803683370190505b50602080860151865192955091818601919083010160005b858110156143a85781810151838201526143a1602082614f5f565b9050614386565b505050935250919050565b600061253382614765565b602080820151825181019091015160009182805b600a81101561443a5783811a91506143eb816007614f99565b82607f16901b8517945081608016600014156144285761440c816001614f5f565b8651879061441b908390614f5f565b9052509395945050505050565b806144328161502d565b9150506143d2565b50600080fd5b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b6020830151518351101561407a57614482836142bd565b909250905081600114156144ac5761449c6139b6846142f7565b6001600160a01b0316845261446b565b81600214156144cb576144c161405c846142f7565b602085015261446b565b6144d583826144da565b61446b565b60008160058111156144ee576144ee61509f565b14156144fd57610f26826143be565b60028160058111156145115761451161509f565b1415610324576000614522836143be565b905080836000018181516145369190614f5f565b90525060208301515183511115610f2657600080fd5b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156145c95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610aed565b8360ff16601b14806145de57508360ff16601c145b6146355760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610aed565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015614689573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166146ec5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610aed565b95945050505050565b60006001600160ff1b03821660ff83901c601b016147158682878561454c565b9695505050505050565b600060208251111561473057600080fd5b60208201519050815160206147459190614fb8565b614750906008614f99565b1c919050565b60606132888484600085614784565b6000815160141461477557600080fd5b5060200151600160601b900490565b6060824710156147e55760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610aed565b843b6148335760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610aed565b600080866001600160a01b0316858760405161484f9190614cfe565b60006040518083038185875af1925050503d806000811461488c576040519150601f19603f3d011682016040523d82523d6000602084013e614891565b606091505b50915091506148a18282866148ac565b979650505050505050565b606083156148bb57508161328b565b8251156148cb5782518084602001fd5b8160405162461bcd60e51b8152600401610aed9190614e26565b604051806060016040528060006001600160a01b0316815260200160008152602001606081525090565b60006001600160401b0380841115614929576149296150e1565b8360051b602061493a818301614f2f565b86815293508084018583810189101561495257600080fd5b60009350835b8881101561498d5781358681111561496e578586fd5b61497a8b828b016149b6565b8452509183019190830190600101614958565b5050505050509392505050565b80356001600160a01b03811681146149b157600080fd5b919050565b600082601f8301126149c757600080fd5b81356001600160401b038111156149e0576149e06150e1565b6149f3601f8201601f1916602001614f2f565b818152846020838601011115614a0857600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215614a3757600080fd5b61328b8261499a565b60008060408385031215614a5357600080fd5b614a5c8361499a565b9150614a6a6020840161499a565b90509250929050565b60008060408385031215614a8657600080fd5b614a8f8361499a565b946020939093013593505050565b600060208284031215614aaf57600080fd5b8151801515811461328b57600080fd5b60008060008060408587031215614ad557600080fd5b84356001600160401b0380821115614aec57600080fd5b818701915087601f830112614b0057600080fd5b813581811115614b0f57600080fd5b886020828501011115614b2157600080fd5b602092830196509450908601359080821115614b3c57600080fd5b818701915087601f830112614b5057600080fd5b813581811115614b5f57600080fd5b8860208260051b8501011115614b7457600080fd5b95989497505060200194505050565b60008060408385031215614b9657600080fd5b82356001600160401b0380821115614bad57600080fd5b614bb9868387016149b6565b93506020850135915080821115614bcf57600080fd5b508301601f81018513614be157600080fd5b614bf08582356020840161490f565b9150509250929050565b600060208284031215614c0c57600080fd5b5035919050565b60008060408385031215614c2657600080fd5b82359150614a6a6020840161499a565b60008060408385031215614c4957600080fd5b82359150602083013560048110614c5f57600080fd5b809150509250929050565b60008060408385031215614c7d57600080fd5b50508035926020909101359150565b80516001600160a01b0316825260208082015181840152604080830151606082860181905281519086018190526000939182019290849060808801905b80831015614cf25785518051835285015185830152948401946001929092019190830190614cc9565b50979650505050505050565b60008251614d10818460208701615001565b9190910192915050565b6001600160a01b03871681526020810186905260408101859052606081018490526080810183905260c0810160038310614d5657614d5661509f565b8260a0830152979650505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015614dbc57603f19888603018452614daa858351614c8c565b94509285019290850190600101614d8e565b5092979650505050505050565b60208101614dd6836150f7565b91905290565b6101208101614dea8b6150f7565b998152602081019890985260408801969096526060870194909452608086019290925260a085015260c084015260e08301526101009091015290565b6020815260008251806020840152614e45816040850160208701615001565b601f01601f19169190910160400192915050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601c908201527f56616c696461746f72206973206e6f7420696e697469616c697a656400000000604082015260600190565b60208152600061328b6020830184614c8c565b8381526001600160a01b038316602082015260608101614f21836150f7565b826040830152949350505050565b604051601f8201601f191681016001600160401b0381118282101715614f5757614f576150e1565b604052919050565b60008219821115614f7257614f72615089565b500190565b600082614f9457634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615614fb357614fb3615089565b500290565b600082821015614fca57614fca615089565b500390565b600063ffffffff83811690831681811015614fec57614fec615089565b039392505050565b600061328b36848461490f565b60005b8381101561501c578181015183820152602001615004565b838111156133ff5750506000910152565b600060001982141561504157615041615089565b5060010190565b600063ffffffff8083168181141561506257615062615089565b6001019392505050565b6000600160ff1b82141561508257615082615089565b5060000390565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60048110610dd557610dd561509f56fea26469706673582212209e75a7bd2d6fbb4942d24aa1b9d91846607e934cbd48eb652a4d9a94349df9ff64736f6c63430008070033",
}

// DPoSABI is the input ABI used to generate the binding from.
// Deprecated: Use DPoSMetaData.ABI instead.
var DPoSABI = DPoSMetaData.ABI

// DPoSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DPoSMetaData.Bin instead.
var DPoSBin = DPoSMetaData.Bin

// DeployDPoS deploys a new Ethereum contract, binding an instance of DPoS to it.
func DeployDPoS(auth *bind.TransactOpts, backend bind.ContractBackend, _celerTokenAddress common.Address, _governProposalDeposit *big.Int, _governVoteTimeout *big.Int, _slashTimeout *big.Int, _maxBondedValidators *big.Int, _minValidatorTokens *big.Int, _advanceNoticePeriod *big.Int) (common.Address, *types.Transaction, *DPoS, error) {
	parsed, err := DPoSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DPoSBin), backend, _celerTokenAddress, _governProposalDeposit, _governVoteTimeout, _slashTimeout, _maxBondedValidators, _minValidatorTokens, _advanceNoticePeriod)
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

// UIntStorage is a free data retrieval call binding the contract method 0x64ed600a.
//
// Solidity: function UIntStorage(uint256 ) view returns(uint256)
func (_DPoS *DPoSCaller) UIntStorage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "UIntStorage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UIntStorage is a free data retrieval call binding the contract method 0x64ed600a.
//
// Solidity: function UIntStorage(uint256 ) view returns(uint256)
func (_DPoS *DPoSSession) UIntStorage(arg0 *big.Int) (*big.Int, error) {
	return _DPoS.Contract.UIntStorage(&_DPoS.CallOpts, arg0)
}

// UIntStorage is a free data retrieval call binding the contract method 0x64ed600a.
//
// Solidity: function UIntStorage(uint256 ) view returns(uint256)
func (_DPoS *DPoSCallerSession) UIntStorage(arg0 *big.Int) (*big.Int, error) {
	return _DPoS.Contract.UIntStorage(&_DPoS.CallOpts, arg0)
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

// BondedValTokens is a free data retrieval call binding the contract method 0xaf062e6f.
//
// Solidity: function bondedValTokens() view returns(uint256)
func (_DPoS *DPoSCaller) BondedValTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "bondedValTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedValTokens is a free data retrieval call binding the contract method 0xaf062e6f.
//
// Solidity: function bondedValTokens() view returns(uint256)
func (_DPoS *DPoSSession) BondedValTokens() (*big.Int, error) {
	return _DPoS.Contract.BondedValTokens(&_DPoS.CallOpts)
}

// BondedValTokens is a free data retrieval call binding the contract method 0xaf062e6f.
//
// Solidity: function bondedValTokens() view returns(uint256)
func (_DPoS *DPoSCallerSession) BondedValTokens() (*big.Int, error) {
	return _DPoS.Contract.BondedValTokens(&_DPoS.CallOpts)
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
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,(uint256,uint256)[]))
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
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,(uint256,uint256)[]))
func (_DPoS *DPoSSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (DPoSDelegatorInfo, error) {
	return _DPoS.Contract.GetDelegatorInfo(&_DPoS.CallOpts, _valAddr, _delAddr)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,(uint256,uint256)[]))
func (_DPoS *DPoSCallerSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (DPoSDelegatorInfo, error) {
	return _DPoS.Contract.GetDelegatorInfo(&_DPoS.CallOpts, _valAddr, _delAddr)
}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,(uint256,uint256)[])[])
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
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,(uint256,uint256)[])[])
func (_DPoS *DPoSSession) GetDelegatorInfos(_delAddr common.Address) ([]DPoSDelegatorInfo, error) {
	return _DPoS.Contract.GetDelegatorInfos(&_DPoS.CallOpts, _delAddr)
}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,(uint256,uint256)[])[])
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

// GetUIntValue is a free data retrieval call binding the contract method 0x64c66395.
//
// Solidity: function getUIntValue(uint256 _record) view returns(uint256)
func (_DPoS *DPoSCaller) GetUIntValue(opts *bind.CallOpts, _record *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "getUIntValue", _record)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUIntValue is a free data retrieval call binding the contract method 0x64c66395.
//
// Solidity: function getUIntValue(uint256 _record) view returns(uint256)
func (_DPoS *DPoSSession) GetUIntValue(_record *big.Int) (*big.Int, error) {
	return _DPoS.Contract.GetUIntValue(&_DPoS.CallOpts, _record)
}

// GetUIntValue is a free data retrieval call binding the contract method 0x64c66395.
//
// Solidity: function getUIntValue(uint256 _record) view returns(uint256)
func (_DPoS *DPoSCallerSession) GetUIntValue(_record *big.Int) (*big.Int, error) {
	return _DPoS.Contract.GetUIntValue(&_DPoS.CallOpts, _record)
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
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue, uint8 status)
func (_DPoS *DPoSCaller) ParamProposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Record       *big.Int
	NewValue     *big.Int
	Status       uint8
}, error) {
	var out []interface{}
	err := _DPoS.contract.Call(opts, &out, "paramProposals", arg0)

	outstruct := new(struct {
		Proposer     common.Address
		Deposit      *big.Int
		VoteDeadline *big.Int
		Record       *big.Int
		NewValue     *big.Int
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VoteDeadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Record = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NewValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue, uint8 status)
func (_DPoS *DPoSSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Record       *big.Int
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _DPoS.Contract.ParamProposals(&_DPoS.CallOpts, arg0)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue, uint8 status)
func (_DPoS *DPoSCallerSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Record       *big.Int
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _DPoS.Contract.ParamProposals(&_DPoS.CallOpts, arg0)
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
// Solidity: function validators(address ) view returns(uint8 status, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 bondBlock, uint256 unbondBlock, uint256 commissionRate, uint256 minSelfDelegation)
func (_DPoS *DPoSCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status             uint8
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
	outstruct.Tokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Shares = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UndelegationTokens = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UndelegationShares = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BondBlock = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UnbondBlock = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CommissionRate = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.MinSelfDelegation = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 bondBlock, uint256 unbondBlock, uint256 commissionRate, uint256 minSelfDelegation)
func (_DPoS *DPoSSession) Validators(arg0 common.Address) (struct {
	Status             uint8
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
// Solidity: function validators(address ) view returns(uint8 status, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 bondBlock, uint256 unbondBlock, uint256 commissionRate, uint256 minSelfDelegation)
func (_DPoS *DPoSCallerSession) Validators(arg0 common.Address) (struct {
	Status             uint8
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

// CreateParamProposal is a paid mutator transaction binding the contract method 0x3090c0e9.
//
// Solidity: function createParamProposal(uint256 _record, uint256 _value) returns()
func (_DPoS *DPoSTransactor) CreateParamProposal(opts *bind.TransactOpts, _record *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "createParamProposal", _record, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0x3090c0e9.
//
// Solidity: function createParamProposal(uint256 _record, uint256 _value) returns()
func (_DPoS *DPoSSession) CreateParamProposal(_record *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.CreateParamProposal(&_DPoS.TransactOpts, _record, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0x3090c0e9.
//
// Solidity: function createParamProposal(uint256 _record, uint256 _value) returns()
func (_DPoS *DPoSTransactorSession) CreateParamProposal(_record *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.CreateParamProposal(&_DPoS.TransactOpts, _record, _value)
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

// InitializeValidator is a paid mutator transaction binding the contract method 0x3be88c2a.
//
// Solidity: function initializeValidator(uint256 _minSelfDelegation, uint256 _commissionRate) returns()
func (_DPoS *DPoSTransactor) InitializeValidator(opts *bind.TransactOpts, _minSelfDelegation *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DPoS.contract.Transact(opts, "initializeValidator", _minSelfDelegation, _commissionRate)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x3be88c2a.
//
// Solidity: function initializeValidator(uint256 _minSelfDelegation, uint256 _commissionRate) returns()
func (_DPoS *DPoSSession) InitializeValidator(_minSelfDelegation *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.InitializeValidator(&_DPoS.TransactOpts, _minSelfDelegation, _commissionRate)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x3be88c2a.
//
// Solidity: function initializeValidator(uint256 _minSelfDelegation, uint256 _commissionRate) returns()
func (_DPoS *DPoSTransactorSession) InitializeValidator(_minSelfDelegation *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DPoS.Contract.InitializeValidator(&_DPoS.TransactOpts, _minSelfDelegation, _commissionRate)
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
	Record     *big.Int
	NewValue   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfirmParamProposal is a free log retrieval operation binding the contract event 0x106f43a560e53395081c0423504b476d1a2cfed9d56ff972bf77ae43ff7d4ba4.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint256 record, uint256 newValue)
func (_DPoS *DPoSFilterer) FilterConfirmParamProposal(opts *bind.FilterOpts) (*DPoSConfirmParamProposalIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return &DPoSConfirmParamProposalIterator{contract: _DPoS.contract, event: "ConfirmParamProposal", logs: logs, sub: sub}, nil
}

// WatchConfirmParamProposal is a free log subscription operation binding the contract event 0x106f43a560e53395081c0423504b476d1a2cfed9d56ff972bf77ae43ff7d4ba4.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint256 record, uint256 newValue)
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

// ParseConfirmParamProposal is a log parse operation binding the contract event 0x106f43a560e53395081c0423504b476d1a2cfed9d56ff972bf77ae43ff7d4ba4.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint256 record, uint256 newValue)
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
	Record       *big.Int
	NewValue     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateParamProposal is a free log retrieval operation binding the contract event 0x40109a070319d6004f4e4b31dba4b605c97bd3474d49865158f55fe093e3b339.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue)
func (_DPoS *DPoSFilterer) FilterCreateParamProposal(opts *bind.FilterOpts) (*DPoSCreateParamProposalIterator, error) {

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return &DPoSCreateParamProposalIterator{contract: _DPoS.contract, event: "CreateParamProposal", logs: logs, sub: sub}, nil
}

// WatchCreateParamProposal is a free log subscription operation binding the contract event 0x40109a070319d6004f4e4b31dba4b605c97bd3474d49865158f55fe093e3b339.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue)
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

// ParseCreateParamProposal is a log parse operation binding the contract event 0x40109a070319d6004f4e4b31dba4b605c97bd3474d49865158f55fe093e3b339.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue)
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
	MinSelfDelegation *big.Int
	CommissionRate    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorParamsUpdate is a free log retrieval operation binding the contract event 0xb7f73ede33f427fed817c4721ec2ae9f56e906564766ad98d79f291f8bf0b799.
//
// Solidity: event ValidatorParamsUpdate(address indexed valAddr, uint256 minSelfDelegation, uint256 commissionRate)
func (_DPoS *DPoSFilterer) FilterValidatorParamsUpdate(opts *bind.FilterOpts, valAddr []common.Address) (*DPoSValidatorParamsUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _DPoS.contract.FilterLogs(opts, "ValidatorParamsUpdate", valAddrRule)
	if err != nil {
		return nil, err
	}
	return &DPoSValidatorParamsUpdateIterator{contract: _DPoS.contract, event: "ValidatorParamsUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorParamsUpdate is a free log subscription operation binding the contract event 0xb7f73ede33f427fed817c4721ec2ae9f56e906564766ad98d79f291f8bf0b799.
//
// Solidity: event ValidatorParamsUpdate(address indexed valAddr, uint256 minSelfDelegation, uint256 commissionRate)
func (_DPoS *DPoSFilterer) WatchValidatorParamsUpdate(opts *bind.WatchOpts, sink chan<- *DPoSValidatorParamsUpdate, valAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _DPoS.contract.WatchLogs(opts, "ValidatorParamsUpdate", valAddrRule)
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

// ParseValidatorParamsUpdate is a log parse operation binding the contract event 0xb7f73ede33f427fed817c4721ec2ae9f56e906564766ad98d79f291f8bf0b799.
//
// Solidity: event ValidatorParamsUpdate(address indexed valAddr, uint256 minSelfDelegation, uint256 commissionRate)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_governProposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_governVoteTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"record\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ConfirmParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"record\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CreateParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGovern.VoteType\",\"name\":\"voteType\",\"type\":\"uint8\"}],\"name\":\"VoteParam\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"UIntStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"celerToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_record\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"createParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getParamProposalVote\",\"outputs\":[{\"internalType\":\"enumGovern.VoteType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_record\",\"type\":\"uint256\"}],\"name\":\"getUIntValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextParamProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paramProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"record\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610bc5380380610bc583398101604081905261002f9161018a565b6100383361013a565b600180546001600160a01b0319166001600160a01b03989098169790971790965560026020527fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b949094557fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e0929092557f679795a0195a1b76cdebb7c51d74e058aee92919b8c3389af86ef24535e8a28c557f88601476d11616a71c5be67555bd1dff4b1cbf21533d2669b768b61518cfe1c3557fee60d0579bcffd98e668647d59fec1ff86a7fb340ce572e844f234ae73a6918f5560056000527fb98b78633099fa36ed8b8680c4f8092689e1e04080eb9cbb077ca38a14d7e384556101f3565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600080600080600080600060e0888a0312156101a557600080fd5b87516001600160a01b03811681146101bc57600080fd5b602089015160408a015160608b015160808c015160a08d015160c0909d0151949e939d50919b909a50909850965090945092505050565b6109c3806102026000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063715018a611610066578063715018a61461015e5780637e5fb8f3146101665780638da5cb5b146101cb578063c6c21e9d146101f0578063f2fde38b1461020357600080fd5b806322da7927146100a35780633090c0e9146100bf578063581c53c5146100d457806364c663951461011e57806364ed600a1461013e575b600080fd5b6100ac60045481565b6040519081526020015b60405180910390f35b6100d26100cd36600461082e565b610216565b005b6101116100e2366004610802565b60008281526003602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b6040516100b691906108b9565b6100ac61012c3660046107e9565b60009081526002602052604090205490565b6100ac61014c3660046107e9565b60026020526000908152604090205481565b6100d2610359565b6101b96101743660046107e9565b60036020819052600091825260409091208054600182015460028301549383015460048401546005909401546001600160a01b03909316949193919290919060ff1686565b6040516100b69695949392919061086c565b6000546001600160a01b03165b6040516001600160a01b0390911681526020016100b6565b6001546101d8906001600160a01b031681565b6100d26102113660046107ac565b6103c4565b600454600081815260036020526040902090610233906001610906565b60045560026020527fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b5481546001600160a01b03191633908117835560018084018390556000527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e0549091906102a99043610906565b6002840155600383018590556004830184905560058301805460ff19166001908117909155546102e4906001600160a01b031683308461048f565b7f40109a070319d6004f4e4b31dba4b605c97bd3474d49865158f55fe093e3b3396001600454610314919061091e565b6002850154604080519283526001600160a01b0386166020840152820184905260608201526080810187905260a0810186905260c00160405180910390a15050505050565b6000546001600160a01b031633146103b85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6103c260006104ef565b565b6000546001600160a01b0316331461041e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103af565b6001600160a01b0381166104835760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103af565b61048c816104ef565b50565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526104e990859061053f565b50505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610594826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166106169092919063ffffffff16565b80519091501561061157808060200190518101906105b291906107c7565b6106115760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016103af565b505050565b6060610625848460008561062f565b90505b9392505050565b6060824710156106905760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016103af565b843b6106de5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103af565b600080866001600160a01b031685876040516106fa9190610850565b60006040518083038185875af1925050503d8060008114610737576040519150601f19603f3d011682016040523d82523d6000602084013e61073c565b606091505b509150915061074c828286610757565b979650505050505050565b60608315610766575081610628565b8251156107765782518084602001fd5b8160405162461bcd60e51b81526004016103af91906108d3565b80356001600160a01b03811681146107a757600080fd5b919050565b6000602082840312156107be57600080fd5b61062882610790565b6000602082840312156107d957600080fd5b8151801515811461062857600080fd5b6000602082840312156107fb57600080fd5b5035919050565b6000806040838503121561081557600080fd5b8235915061082560208401610790565b90509250929050565b6000806040838503121561084157600080fd5b50508035926020909101359150565b60008251610862818460208701610935565b9190910192915050565b6001600160a01b03871681526020810186905260408101859052606081018490526080810183905260c08101600383106108a8576108a8610977565b8260a0830152979650505050505050565b60208101600483106108cd576108cd610977565b91905290565b60208152600082518060208401526108f2816040850160208701610935565b601f01601f19169190910160400192915050565b6000821982111561091957610919610961565b500190565b60008282101561093057610930610961565b500390565b60005b83811015610950578181015183820152602001610938565b838111156104e95750506000910152565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fdfea2646970667358221220b9a36c0cdbded29c40658d6d25bc6ff0b711b88740d5affe40aa27e1c4e4289b64736f6c63430008070033",
}

// GovernABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernMetaData.ABI instead.
var GovernABI = GovernMetaData.ABI

// GovernBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernMetaData.Bin instead.
var GovernBin = GovernMetaData.Bin

// DeployGovern deploys a new Ethereum contract, binding an instance of Govern to it.
func DeployGovern(auth *bind.TransactOpts, backend bind.ContractBackend, _celerTokenAddress common.Address, _governProposalDeposit *big.Int, _governVoteTimeout *big.Int, _slashTimeout *big.Int, _maxBondedValidators *big.Int, _minValidatorTokens *big.Int, _advanceNoticePeriod *big.Int) (common.Address, *types.Transaction, *Govern, error) {
	parsed, err := GovernMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernBin), backend, _celerTokenAddress, _governProposalDeposit, _governVoteTimeout, _slashTimeout, _maxBondedValidators, _minValidatorTokens, _advanceNoticePeriod)
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

// UIntStorage is a free data retrieval call binding the contract method 0x64ed600a.
//
// Solidity: function UIntStorage(uint256 ) view returns(uint256)
func (_Govern *GovernCaller) UIntStorage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "UIntStorage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UIntStorage is a free data retrieval call binding the contract method 0x64ed600a.
//
// Solidity: function UIntStorage(uint256 ) view returns(uint256)
func (_Govern *GovernSession) UIntStorage(arg0 *big.Int) (*big.Int, error) {
	return _Govern.Contract.UIntStorage(&_Govern.CallOpts, arg0)
}

// UIntStorage is a free data retrieval call binding the contract method 0x64ed600a.
//
// Solidity: function UIntStorage(uint256 ) view returns(uint256)
func (_Govern *GovernCallerSession) UIntStorage(arg0 *big.Int) (*big.Int, error) {
	return _Govern.Contract.UIntStorage(&_Govern.CallOpts, arg0)
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

// GetUIntValue is a free data retrieval call binding the contract method 0x64c66395.
//
// Solidity: function getUIntValue(uint256 _record) view returns(uint256)
func (_Govern *GovernCaller) GetUIntValue(opts *bind.CallOpts, _record *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "getUIntValue", _record)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUIntValue is a free data retrieval call binding the contract method 0x64c66395.
//
// Solidity: function getUIntValue(uint256 _record) view returns(uint256)
func (_Govern *GovernSession) GetUIntValue(_record *big.Int) (*big.Int, error) {
	return _Govern.Contract.GetUIntValue(&_Govern.CallOpts, _record)
}

// GetUIntValue is a free data retrieval call binding the contract method 0x64c66395.
//
// Solidity: function getUIntValue(uint256 _record) view returns(uint256)
func (_Govern *GovernCallerSession) GetUIntValue(_record *big.Int) (*big.Int, error) {
	return _Govern.Contract.GetUIntValue(&_Govern.CallOpts, _record)
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
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue, uint8 status)
func (_Govern *GovernCaller) ParamProposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Record       *big.Int
	NewValue     *big.Int
	Status       uint8
}, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "paramProposals", arg0)

	outstruct := new(struct {
		Proposer     common.Address
		Deposit      *big.Int
		VoteDeadline *big.Int
		Record       *big.Int
		NewValue     *big.Int
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VoteDeadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Record = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NewValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue, uint8 status)
func (_Govern *GovernSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Record       *big.Int
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _Govern.Contract.ParamProposals(&_Govern.CallOpts, arg0)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue, uint8 status)
func (_Govern *GovernCallerSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Record       *big.Int
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _Govern.Contract.ParamProposals(&_Govern.CallOpts, arg0)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0x3090c0e9.
//
// Solidity: function createParamProposal(uint256 _record, uint256 _value) returns()
func (_Govern *GovernTransactor) CreateParamProposal(opts *bind.TransactOpts, _record *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Govern.contract.Transact(opts, "createParamProposal", _record, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0x3090c0e9.
//
// Solidity: function createParamProposal(uint256 _record, uint256 _value) returns()
func (_Govern *GovernSession) CreateParamProposal(_record *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Govern.Contract.CreateParamProposal(&_Govern.TransactOpts, _record, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0x3090c0e9.
//
// Solidity: function createParamProposal(uint256 _record, uint256 _value) returns()
func (_Govern *GovernTransactorSession) CreateParamProposal(_record *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Govern.Contract.CreateParamProposal(&_Govern.TransactOpts, _record, _value)
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
	Record     *big.Int
	NewValue   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfirmParamProposal is a free log retrieval operation binding the contract event 0x106f43a560e53395081c0423504b476d1a2cfed9d56ff972bf77ae43ff7d4ba4.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint256 record, uint256 newValue)
func (_Govern *GovernFilterer) FilterConfirmParamProposal(opts *bind.FilterOpts) (*GovernConfirmParamProposalIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return &GovernConfirmParamProposalIterator{contract: _Govern.contract, event: "ConfirmParamProposal", logs: logs, sub: sub}, nil
}

// WatchConfirmParamProposal is a free log subscription operation binding the contract event 0x106f43a560e53395081c0423504b476d1a2cfed9d56ff972bf77ae43ff7d4ba4.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint256 record, uint256 newValue)
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

// ParseConfirmParamProposal is a log parse operation binding the contract event 0x106f43a560e53395081c0423504b476d1a2cfed9d56ff972bf77ae43ff7d4ba4.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint256 record, uint256 newValue)
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
	Record       *big.Int
	NewValue     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateParamProposal is a free log retrieval operation binding the contract event 0x40109a070319d6004f4e4b31dba4b605c97bd3474d49865158f55fe093e3b339.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue)
func (_Govern *GovernFilterer) FilterCreateParamProposal(opts *bind.FilterOpts) (*GovernCreateParamProposalIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return &GovernCreateParamProposalIterator{contract: _Govern.contract, event: "CreateParamProposal", logs: logs, sub: sub}, nil
}

// WatchCreateParamProposal is a free log subscription operation binding the contract event 0x40109a070319d6004f4e4b31dba4b605c97bd3474d49865158f55fe093e3b339.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue)
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

// ParseCreateParamProposal is a log parse operation binding the contract event 0x40109a070319d6004f4e4b31dba4b605c97bd3474d49865158f55fe093e3b339.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint256 record, uint256 newValue)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celrAddr\",\"type\":\"address\"},{\"internalType\":\"contractDPoS\",\"name\":\"_dpos\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"oldAddr\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"newAddr\",\"type\":\"bytes\"}],\"name\":\"SgnAddrUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"celr\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dpos\",\"outputs\":[{\"internalType\":\"contractDPoS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sgnAddrs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sgnAddr\",\"type\":\"bytes\"}],\"name\":\"updateSgnAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161107838038061107883398101604081905261002f916100b3565b61003833610063565b6000805460ff60a01b191690556001600160601b0319606092831b8116608052911b1660a052610105565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600080604083850312156100c657600080fd5b82516100d1816100ed565b60208401519092506100e2816100ed565b809150509250929050565b6001600160a01b038116811461010257600080fd5b50565b60805160601c60a05160601c610f336101456000396000818161014d015261051c0152600081816101fb015281816102a901526103f80152610f336000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638f8451b21161008c578063d0bb935111610066578063d0bb9351146101a2578063f2fde38b146101b5578063fc7e286d146101c8578063fcb02467146101f657600080fd5b80638f8451b214610148578063b6b55f251461016f578063c429fe1f1461018257600080fd5b8063145aa116146100d45780633f4ba83a146100e95780635c975abb146100f1578063715018a6146101135780638456cb591461011b5780638da5cb5b14610123575b600080fd5b6100e76100e2366004610da1565b61021d565b005b6100e76102d3565b600054600160a01b900460ff1660405190151581526020015b60405180910390f35b6100e7610307565b6100e761033b565b6000546001600160a01b03165b6040516001600160a01b03909116815260200161010a565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b6100e761017d366004610da1565b61036d565b610195610190366004610cc3565b610467565b60405161010a9190610e12565b6100e76101b0366004610d0e565b610501565b6100e76101c3366004610cc3565b61072b565b6101e86101d6366004610cc3565b60016020526000908152604090205481565b60405190815260200161010a565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b600054600160a01b900460ff166102725760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064015b60405180910390fd5b6000546001600160a01b0316331461029c5760405162461bcd60e51b815260040161026990610e25565b6102d06001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633836107c3565b50565b6000546001600160a01b031633146102fd5760405162461bcd60e51b815260040161026990610e25565b61030561082b565b565b6000546001600160a01b031633146103315760405162461bcd60e51b815260040161026990610e25565b61030560006108c8565b6000546001600160a01b031633146103655760405162461bcd60e51b815260040161026990610e25565b610305610918565b600054600160a01b900460ff16156103ba5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610269565b336000818152600160205260409020546103d5908390610e5a565b6001600160a01b03808316600090815260016020526040902091909155610420907f0000000000000000000000000000000000000000000000000000000000000000168230856109a0565b806001600160a01b03167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c8360405161045b91815260200190565b60405180910390a25050565b6002602052600090815260409020805461048090610eac565b80601f01602080910402602001604051908101604052809291908181526020018280546104ac90610eac565b80156104f95780601f106104ce576101008083540402835291602001916104f9565b820191906000526020600020905b8154815290600101906020018083116104dc57829003601f168201915b505050505081565b60405163a310624f60e01b81523360048201819052906000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a310624f9060240160206040518083038186803b15801561056657600080fd5b505afa15801561057a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059e9190610d80565b905060018160038111156105b4576105b4610ee7565b146105fa5760405162461bcd60e51b81526020600482015260166024820152752737ba103ab73137b73232b2103b30b634b230ba37b960511b6044820152606401610269565b6001600160a01b0382166000908152600260205260408120805461061d90610eac565b80601f016020809104026020016040519081016040528092919081815260200182805461064990610eac565b80156106965780601f1061066b57610100808354040283529160200191610696565b820191906000526020600020905b81548152906001019060200180831161067957829003601f168201915b505050506001600160a01b03851660009081526002602052604090209192506106c29190508686610c2a565b5084846040516106d3929190610de6565b6040518091039020816040516106e99190610df6565b604051908190038120906001600160a01b038616907f8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb490600090a45050505050565b6000546001600160a01b031633146107555760405162461bcd60e51b815260040161026990610e25565b6001600160a01b0381166107ba5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610269565b6102d0816108c8565b6040516001600160a01b03831660248201526044810182905261082690849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526109de565b505050565b600054600160a01b900460ff1661087b5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610269565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff16156109655760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610269565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586108ab3390565b6040516001600160a01b03808516602483015283166044820152606481018290526109d89085906323b872dd60e01b906084016107ef565b50505050565b6000610a33826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610ab09092919063ffffffff16565b8051909150156108265780806020019051810190610a519190610cec565b6108265760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610269565b6060610abf8484600085610ac9565b90505b9392505050565b606082471015610b2a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610269565b843b610b785760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610269565b600080866001600160a01b03168587604051610b949190610df6565b60006040518083038185875af1925050503d8060008114610bd1576040519150601f19603f3d011682016040523d82523d6000602084013e610bd6565b606091505b5091509150610be6828286610bf1565b979650505050505050565b60608315610c00575081610ac2565b825115610c105782518084602001fd5b8160405162461bcd60e51b81526004016102699190610e12565b828054610c3690610eac565b90600052602060002090601f016020900481019282610c585760008555610c9e565b82601f10610c715782800160ff19823516178555610c9e565b82800160010185558215610c9e579182015b82811115610c9e578235825591602001919060010190610c83565b50610caa929150610cae565b5090565b5b80821115610caa5760008155600101610caf565b600060208284031215610cd557600080fd5b81356001600160a01b0381168114610ac257600080fd5b600060208284031215610cfe57600080fd5b81518015158114610ac257600080fd5b60008060208385031215610d2157600080fd5b823567ffffffffffffffff80821115610d3957600080fd5b818501915085601f830112610d4d57600080fd5b813581811115610d5c57600080fd5b866020828501011115610d6e57600080fd5b60209290920196919550909350505050565b600060208284031215610d9257600080fd5b815160048110610ac257600080fd5b600060208284031215610db357600080fd5b5035919050565b60008151808452610dd2816020860160208601610e80565b601f01601f19169290920160200192915050565b8183823760009101908152919050565b60008251610e08818460208701610e80565b9190910192915050565b602081526000610ac26020830184610dba565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60008219821115610e7b57634e487b7160e01b600052601160045260246000fd5b500190565b60005b83811015610e9b578181015183820152602001610e83565b838111156109d85750506000910152565b600181811c90821680610ec057607f821691505b60208210811415610ee157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052602160045260246000fdfea26469706673582212207f16bcfb3de50d9cade1a8576e2f5337bc88178df044c032b534610e9255273964736f6c63430008070033",
}

// SGNABI is the input ABI used to generate the binding from.
// Deprecated: Use SGNMetaData.ABI instead.
var SGNABI = SGNMetaData.ABI

// SGNBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SGNMetaData.Bin instead.
var SGNBin = SGNMetaData.Bin

// DeploySGN deploys a new Ethereum contract, binding an instance of SGN to it.
func DeploySGN(auth *bind.TransactOpts, backend bind.ContractBackend, _celrAddr common.Address, _dpos common.Address) (common.Address, *types.Transaction, *SGN, error) {
	parsed, err := SGNMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SGNBin), backend, _celrAddr, _dpos)
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

// Celr is a free data retrieval call binding the contract method 0xfcb02467.
//
// Solidity: function celr() view returns(address)
func (_SGN *SGNCaller) Celr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "celr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Celr is a free data retrieval call binding the contract method 0xfcb02467.
//
// Solidity: function celr() view returns(address)
func (_SGN *SGNSession) Celr() (common.Address, error) {
	return _SGN.Contract.Celr(&_SGN.CallOpts)
}

// Celr is a free data retrieval call binding the contract method 0xfcb02467.
//
// Solidity: function celr() view returns(address)
func (_SGN *SGNCallerSession) Celr() (common.Address, error) {
	return _SGN.Contract.Celr(&_SGN.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_SGN *SGNCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "deposits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_SGN *SGNSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _SGN.Contract.Deposits(&_SGN.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_SGN *SGNCallerSession) Deposits(arg0 common.Address) (*big.Int, error) {
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

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SGN *SGNTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SGN *SGNSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.Deposit(&_SGN.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SGN *SGNTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.Deposit(&_SGN.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_SGN *SGNTransactor) DrainToken(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "drainToken", _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_SGN *SGNSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.DrainToken(&_SGN.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_SGN *SGNTransactorSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.DrainToken(&_SGN.TransactOpts, _amount)
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
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_SGN *SGNFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*SGNDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &SGNDepositIterator{contract: _SGN.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_SGN *SGNFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SGNDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Deposit", accountRule)
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
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
	OldAddr common.Hash
	NewAddr common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSgnAddrUpdate is a free log retrieval operation binding the contract event 0x8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4.
//
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes indexed oldAddr, bytes indexed newAddr)
func (_SGN *SGNFilterer) FilterSgnAddrUpdate(opts *bind.FilterOpts, valAddr []common.Address, oldAddr [][]byte, newAddr [][]byte) (*SGNSgnAddrUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}

	logs, sub, err := _SGN.contract.FilterLogs(opts, "SgnAddrUpdate", valAddrRule, oldAddrRule, newAddrRule)
	if err != nil {
		return nil, err
	}
	return &SGNSgnAddrUpdateIterator{contract: _SGN.contract, event: "SgnAddrUpdate", logs: logs, sub: sub}, nil
}

// WatchSgnAddrUpdate is a free log subscription operation binding the contract event 0x8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4.
//
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes indexed oldAddr, bytes indexed newAddr)
func (_SGN *SGNFilterer) WatchSgnAddrUpdate(opts *bind.WatchOpts, sink chan<- *SGNSgnAddrUpdate, valAddr []common.Address, oldAddr [][]byte, newAddr [][]byte) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}

	logs, sub, err := _SGN.contract.WatchLogs(opts, "SgnAddrUpdate", valAddrRule, oldAddrRule, newAddrRule)
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
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes indexed oldAddr, bytes indexed newAddr)
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
