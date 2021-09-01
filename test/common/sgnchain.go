package common

import (
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/validator/client/cli"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ContractParams struct {
	CelrAddr              eth.Addr
	ProposalDeposit       *big.Int
	VotePeriod            *big.Int
	UnbondingPeriod       *big.Int
	MaxBondedValidators   *big.Int
	MinValidatorTokens    *big.Int
	MinSelfDelegation     *big.Int
	AdvanceNoticePeriod   *big.Int
	ValidatorBondInterval *big.Int
	MaxSlashFactor        *big.Int
	StartGateway          bool
}

func SetupSgnchain() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
	config.Seal()
}

func NewTestTransactor(t *testing.T, sgnCLIHome, sgnChainID, sgnNodeURI, sgnValAcct, sgnPassphrase string) *transactor.Transactor {
	encodingConfig := app.MakeEncodingConfig()

	tr, err := transactor.NewTransactor(
		sgnCLIHome,
		sgnChainID,
		sgnNodeURI,
		sgnValAcct,
		sgnPassphrase,
		encodingConfig.Codec,
		encodingConfig.Amino,
	)
	require.NoError(t, err, "Failed to create new transactor.")
	tr.Run()

	return tr
}

func CheckValidator(t *testing.T, transactor *transactor.Transactor, expVal *types.Validator) {
	var validator *types.Validator
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		validator, err = cli.QueryValidator(transactor.CliCtx, expVal.EthAddress)
		if err != nil {
			log.Debugln("retry due to err:", err)
		}
		if err == nil && sameValidators(validator, expVal) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to QueryValidator", err)
	log.Infof("Query sgn and get validator: %s", validator)
	assert.True(t, sameValidators(validator, expVal), "The expected validator should be: "+expVal.String())
}

func CheckDelegator(t *testing.T, transactor *transactor.Transactor, expDel *types.Delegator) {
	var delegator *types.Delegator
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		delegator, err = cli.QueryDelegator(transactor.CliCtx, expDel.ValAddress, expDel.DelAddress)
		if err == nil && sameDelegators(delegator, expDel) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to queryDelegator")
	log.Infof("Query sgn and get delegator: %s", delegator)
	assert.True(t, sameDelegators(delegator, expDel), "The expected delegator should be: "+expDel.String())
}

func CheckSdkValidator(t *testing.T, transactor *transactor.Transactor, expVal *sdk_staking.Validator) {
	var sdkval *sdk_staking.Validator
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		sdkval, err = cli.QuerySdkValidator(transactor.CliCtx, expVal.OperatorAddress)
		if err == nil && sameSdkValidators(sdkval, expVal) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to QuerySdkValidator")
	log.Infof("Query sgn and get sdk validator: %s", sdkval)
	assert.True(t, sameSdkValidators(sdkval, expVal), "The expected sdk validator should be: "+expVal.String())
}

func CheckBondedSdkValidatorNum(t *testing.T, transactor *transactor.Transactor, expNum int) {
	var sdkvals sdk_staking.Validators
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		sdkvals, err = cli.QuerySdkValidators(transactor.CliCtx, sdk_staking.BondStatusBonded)
		if err == nil && len(sdkvals) == expNum {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to QuerySdkValidators")
	assert.Equal(t, expNum, len(sdkvals), "The length of validators should be: "+strconv.Itoa(expNum))
}

// TODO: check pubkey, transactors, and description
func sameValidators(v *types.Validator, exp *types.Validator) bool {
	return v.GetEthAddress() == exp.GetEthAddress() &&
		v.GetEthSigner() == exp.GetEthSigner() &&
		v.GetStatus() == exp.GetStatus() &&
		v.GetSgnAddress() == exp.GetSgnAddress() &&
		v.GetTokens() == exp.GetTokens() &&
		v.GetShares() == exp.GetShares() &&
		v.GetCommissionRate() == exp.GetCommissionRate()
}

func sameDelegators(d *types.Delegator, exp *types.Delegator) bool {
	return d.GetValAddress() == exp.GetValAddress() &&
		d.GetValAddress() == exp.GetDelAddress() &&
		d.GetShares() == exp.GetShares()
}

func sameSdkValidators(v *sdk_staking.Validator, exp *sdk_staking.Validator) bool {
	return v.GetOperator().Equals(exp.GetOperator()) &&
		v.GetStatus() == exp.GetStatus() &&
		v.GetTokens().Equal(exp.GetTokens())
}