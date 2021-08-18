package common

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/contracts"
	"github.com/celer-network/sgn-v2/transactor"
	sgnval "github.com/celer-network/sgn-v2/x/validator"
	vtypes "github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type SGNParams struct {
	CelrAddr               contracts.Addr
	GovernProposalDeposit  *big.Int
	GovernVoteTimeout      *big.Int
	SlashTimeout           *big.Int
	MaxBondedValidators    *big.Int
	MinValidatorTokens     *big.Int
	MinStakingPool         *big.Int
	AdvanceNoticePeriod    *big.Int
	SidechainGoLiveTimeout *big.Int
	StartGateway           bool
}

func SetupSidechain() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
	config.Seal()
}

func NewTestTransactor(t *testing.T, sgnCLIHome, sgnChainID, sgnNodeURI, sgnValAcct, sgnPassphrase string) *transactor.Transactor {

	// TODO
	return nil
}

func CheckCandidate(t *testing.T, transactor *transactor.Transactor, ethAddr contracts.Addr, valacct string, expAmt *big.Int) {
	var candidate vtypes.Candidate
	var err error
	expectedRes := fmt.Sprintf(`ValAccount: %s, EthAddress: %x, StakingPool: %s`, valacct, ethAddr, expAmt) // defined in Candidate.String()
	for retry := 0; retry < RetryLimit; retry++ {
		candidate, err = sgnval.CLIQueryCandidate(transactor.CliCtx, sgnval.RouterKey, ethAddr.Hex())
		if err != nil {
			log.Debugln("retry due to err:", err)
		}
		if err == nil && expectedRes == candidate.String() {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to queryCandidate", err)
	log.Infoln("Query sgn about the validator candidate:", candidate)
	assert.Equal(t, expectedRes, candidate.String(), "The expected result should be: "+expectedRes)
}

func CheckDelegator(t *testing.T, transactor *transactor.Transactor, validatorAddr, delegatorAddr contracts.Addr, expAmt *big.Int) {
	var delegator vtypes.Delegator
	var err error
	expectedRes := fmt.Sprintf(`CandidateAddr: %s, DelegatorAddr: %s, DelegatedStake: %s`,
		contracts.Addr2Hex(validatorAddr), contracts.Addr2Hex(delegatorAddr), expAmt) // defined in Delegator.String()
	for retry := 0; retry < RetryLimit; retry++ {
		delegator, err = sgnval.CLIQueryDelegator(transactor.CliCtx, sgnval.RouterKey, validatorAddr.Hex(), delegatorAddr.Hex())
		if err == nil && expectedRes == delegator.String() {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to queryDelegator")
	log.Infoln("Query sgn about the validator's delegator:", delegator)
	assert.Equal(t, expectedRes, delegator.String(), "The expected result should be: "+expectedRes)
}

func CheckValidator(t *testing.T, transactor *transactor.Transactor, valacct string, expAmt *big.Int, expStatus stypes.BondStatus) {
	var validator stypes.Validator
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		validator, err = sgnval.CLIQueryValidator(transactor.CliCtx, stypes.RouterKey, valacct)
		if err == nil &&
			validator.Status == expStatus {
			expToken := sdk.NewIntFromBigInt(expAmt).QuoRaw(common.TokenDec).String()
			if expToken == validator.Tokens.String() {
				break
			}
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to queryValidator")
	expToken := sdk.NewIntFromBigInt(expAmt).QuoRaw(common.TokenDec).String()
	assert.Equal(t, expToken, validator.Tokens.String(), "validator token should be "+expToken)
	//assert.Equal(t, expStatus, validator.Status, "validator should be "+sdkStatusName(validator.Status))
}
