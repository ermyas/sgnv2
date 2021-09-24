package common

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	govcli "github.com/celer-network/sgn-v2/x/gov/client/cli"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	slashcli "github.com/celer-network/sgn-v2/x/slash/client/cli"
	slashtypes "github.com/celer-network/sgn-v2/x/slash/types"
	"github.com/celer-network/sgn-v2/x/staking/client/cli"
	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		sgnAddr, err := sdk.AccAddressFromBech32(ValSgnAddrStrs[i])
		if err != nil {
			log.Fatal(err)
		}
		ValSgnAddrs = append(ValSgnAddrs, sgnAddr)
	}
}

func NewTestTransactor(sgnHomeDir, sgnChainID, sgnNodeURI, sgnValAcct, sgnPassphrase string) *transactor.Transactor {
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
	ChkErr(err, "Failed to create new transactor.")
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
	msg := fmt.Sprintf("Expected validator:\n %s\n Actual validator:\n %s\n", expVal.String(), validator.String())
	assert.True(t, sameValidators(validator, expVal), msg)
}

// called after CheckValidator
func CheckValidatorBySgnAddr(t *testing.T, transactor *transactor.Transactor, expVal *types.Validator) {
	validator, err := cli.QueryValidatorBySgnAddr(transactor.CliCtx, expVal.SgnAddress)
	require.NoError(t, err, "failed to QueryValidatorBySgnAddr", err)
	msg := fmt.Sprintf("Expected validator:\n %s\n Actual validator:\n %s\n", expVal.String(), validator.String())
	assert.True(t, sameValidators(validator, expVal), msg)
}

func CheckValidators(transactor *transactor.Transactor, expVals types.Validators) {
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
	ChkErr(err, "failed to QueryValidators")
	log.Infof("Query sgn and get validators: %s", validators.String())
	msg := fmt.Sprintf("Expected validators:\n %s\n Actual validators:\n %s\n", expVals.String(), validators.String())
	if !sameEachValidators(validators, expVals) {
		log.Fatalln(msg)
	}
}

func CheckDelegation(t *testing.T, transactor *transactor.Transactor, expDel *types.Delegation) {
	var delegation *types.Delegation
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		delegation, err = cli.QueryDelegation(transactor.CliCtx, expDel.ValidatorAddress, expDel.DelegatorAddress)
		if err == nil && sameDelegations(delegation, expDel) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	require.NoError(t, err, "failed to queryDelegation")
	log.Infof("Query sgn and get delegator: %s", delegation.String())
	assert.True(t, sameDelegations(delegation, expDel), "The expected delegator should be: "+expDel.String())
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
		v.EthSigner == exp.EthSigner &&
		v.GetStatus() == exp.GetStatus() &&
		v.SgnAddress == exp.SgnAddress &&
		v.Tokens.Equal(exp.Tokens) &&
		v.DelegatorShares.Equal(exp.DelegatorShares) &&
		v.CommissionRate.Equal(exp.CommissionRate)
}

func sameDelegations(d *types.Delegation, exp *types.Delegation) bool {
	return d.GetValidatorAddr() == exp.GetValidatorAddr() &&
		d.GetDelegatorAddr() == exp.GetDelegatorAddr() &&
		d.Shares.Equal(exp.Shares)
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
		err = fmt.Errorf("proposal status %s does not match expectation %s", proposal.Status, status)
	}

	return
}

func QuerySlash(cliCtx client.Context, nonce uint64, sigCount int) (slash slashtypes.Slash, err error) {
	for retry := 0; retry < RetryLimit; retry++ {
		slash, err = slashcli.QuerySlash(cliCtx, slashtypes.StoreKey, nonce)
		if err == nil && len(slash.EthSlashBytes) > 0 && len(slash.Sigs) == sigCount {
			break
		}
		time.Sleep(RetryPeriod)
	}

	if err != nil {
		return
	}

	if len(slash.EthSlashBytes) == 0 {
		err = errors.New("EthSlashBytes cannot be zero")
	}

	if len(slash.Sigs) != sigCount {
		err = errors.New("signature count does not match expectation")
	}

	return
}
