package common

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	govcli "github.com/celer-network/sgn-v2/x/gov/client/cli"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/celer-network/sgn-v2/x/validator/client/cli"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func SetupSgnchain() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
	config.Seal()
	for i := 0; i < len(ValSgnAddrStrs); i++ {
		sgnAddr, err := types.SdkAccAddrFromSgnBech32(ValSgnAddrStrs[i])
		if err != nil {
			log.Fatal(err)
		}
		ValSgnAddrs = append(ValSgnAddrs, sgnAddr)
	}
}

func NewTestTransactor(t *testing.T, sgnHomeDir, sgnChainID, sgnNodeURI, sgnValAcct, sgnPassphrase string) *transactor.Transactor {
	encodingConfig := app.MakeEncodingConfig()

	tr, err := transactor.NewTransactor(
		sgnHomeDir,
		sgnChainID,
		sgnNodeURI,
		sgnValAcct,
		sgnPassphrase,
		encodingConfig.Amino,
		encodingConfig.Codec,
		encodingConfig.InterfaceRegistry,
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
	log.Infof("Query sgn and get validator: %s", validator.String())
	assert.True(t, sameValidators(validator, expVal), "The expected validator should be: "+expVal.String())
}

func CheckValidators(t *testing.T, transactor *transactor.Transactor, expVals types.Validators) {
	var validators types.Validators
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		validators, err = cli.QueryValidators(transactor.CliCtx)
		if err != nil {
			log.Debugln("retry due to err:", err)
		}
		if err == nil && sameEachValidators(validators, expVals) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to QueryValidators", err)
	log.Infof("Query sgn and get validators: %s", validators.String())
	assert.True(t, sameEachValidators(validators, expVals), "The expected validators should be: "+expVals.String())
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
	log.Infof("Query sgn and get delegator: %s", delegator.String())
	assert.True(t, sameDelegators(delegator, expDel), "The expected delegator should be: "+expDel.String())
}

func CheckSdkValidator(t *testing.T, transactor *transactor.Transactor, expVal *sdk_staking.Validator) {
	var sdkval *sdk_staking.Validator
	sgnAddr, err := sdk.ValAddressFromBech32(expVal.OperatorAddress)
	require.NoError(t, err, "invalid operator address")
	for retry := 0; retry < RetryLimit; retry++ {
		sdkval, err = cli.QuerySdkValidator(transactor.CliCtx, sdk.AccAddress(sgnAddr).String())
		if err == nil && sameSdkValidators(sdkval, expVal) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to QuerySdkValidator")
	log.Infof("Query sgn and get sdk validator: %s", printSdkVal(*sdkval))
	assert.True(t, sameSdkValidators(sdkval, expVal), "The expected sdk validator should be: "+printSdkVal(*expVal))
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

func PrintTendermintValidators(t *testing.T, transactor *transactor.Transactor) {
	page := 1
	limit := 30
	res, err := rpc.GetValidators(context.Background(), transactor.CliCtx, nil, &page, &limit)
	require.NoError(t, err, "failed to get tendermint validators")
	log.Infof("tendermint validators:\n%s", res)
}

func sameEachValidators(vs []types.Validator, exps []types.Validator) bool {
	same := len(vs) == len(exps)
	if same {
		sort.SliceStable(vs, func(i, j int) bool {
			return vs[i].EthAddress < vs[j].EthAddress
		})
		sort.SliceStable(exps, func(i, j int) bool {
			return exps[i].EthAddress < exps[j].EthAddress
		})

		for i := 0; i < len(vs); i++ {
			same = same && sameValidators(&vs[i], &exps[i])
			if !same {
				break
			}
		}
	}
	return same
}

// TODO: check pubkey, transactors, and description
func sameValidators(v *types.Validator, exp *types.Validator) bool {
	return v.GetEthAddress() == exp.GetEthAddress() &&
		v.GetEthSigner() == exp.GetEthSigner() &&
		v.GetStatus() == exp.GetStatus() &&
		v.GetSgnAddress() == exp.GetSgnAddress() &&
		v.Tokens.Equal(exp.Tokens) &&
		v.Shares.Equal(exp.Shares) &&
		v.GetCommissionRate() == exp.GetCommissionRate()
}

func sameDelegators(d *types.Delegator, exp *types.Delegator) bool {
	return d.GetValAddress() == exp.GetValAddress() &&
		d.GetDelAddress() == exp.GetDelAddress() &&
		d.GetShares() == exp.GetShares()
}

func sameSdkValidators(v *sdk_staking.Validator, exp *sdk_staking.Validator) bool {
	return v.GetOperator().Equals(exp.GetOperator()) &&
		v.GetStatus() == exp.GetStatus() &&
		v.GetTokens().Equal(exp.GetTokens())
}

func printSdkVal(v sdk_staking.Validator) string {
	var pubkey string
	if v.ConsensusPubkey != nil {
		consAddr, err := v.GetConsAddr()
		if err != nil {
			pubkey = fmt.Sprintf("consensus_address:%s", err)
		} else {
			pubkey = fmt.Sprintf("consensus_address:\"%s\"", consAddr.String())
		}
	}
	v.ConsensusPubkey = nil
	out := proto.CompactTextString(&v)
	out += pubkey
	return out
}

func QueryProposal(cliCtx client.Context, proposalID uint64, status govtypes.ProposalStatus) (proposal govtypes.Proposal, err error) {
	for retry := 0; retry < RetryLimit; retry++ {
		proposal, err = govcli.QueryProposal(cliCtx, govtypes.RouterKey, proposalID)
		if err == nil && status == proposal.Status {
			break
		}
		time.Sleep(RetryPeriod)
	}

	if err != nil {
		return
	}

	if status != proposal.Status {
		err = fmt.Errorf("proposal status %s does not match expectation", status)
	}

	return
}
