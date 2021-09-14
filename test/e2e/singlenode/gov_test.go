package singlenode

import (
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGov() []tc.Killable {
	res := setupNewSgnEnv(nil, "gov change parameter")
	tc.SleepWithLog(10, "sgn syncing")

	return res
}

func TestGov(t *testing.T) {
	toKill := setupGov()
	defer tc.TearDown(toKill)

	t.Run("e2e-gov", func(t *testing.T) {
		t.Run("govTest", govTest)
	})
}

func govTest(t *testing.T) {
	log.Info("=====================================================================")
	log.Info("======================== Test gov ===========================")

	transactor := tc.NewTestTransactor(
		t,
		NodeHome,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
	)

	amt := big.NewInt(5e18)
	err := tc.InitializeValidator(tc.ValAuths[0], tc.ValSgnAddrs[0], amt, eth.CommissionRate(0.02))
	require.NoError(t, err, "failed to initialize validator")
	tc.Sleep(5)
	expVal := &stakingtypes.Validator{
		EthAddress:      eth.Addr2Hex(tc.ValEthAddrs[0]),
		EthSigner:       eth.Addr2Hex(tc.ValEthAddrs[0]),
		Status:          eth.Bonded,
		SgnAddress:      tc.ValSgnAddrs[0].String(),
		Tokens:          sdk.NewIntFromBigInt(amt),
		DelegatorShares: sdk.NewIntFromBigInt(amt),
		CommissionRate:  sdk.NewDecWithPrec(2, 2),
	}
	tc.CheckValidator(t, transactor, expVal)

	log.Info("======================== Test change epochlength passed ===========================")
	paramChanges := []govtypes.ParamChange{govtypes.NewParamChange("staking", "EpochLength", "\"2\"")}
	content := govtypes.NewParameterProposal("Guard Param Change", "Update EpochLength", paramChanges)
	submitProposalmsg, err := govtypes.NewMsgSubmitProposal(content, sdk.ZeroInt(), transactor.Key.GetAddress())
	require.NoError(t, err, "failed to create MsgSubmitProposal")
	transactor.AddTxMsg(submitProposalmsg)

	proposalID := uint64(1)
	proposal, err := tc.QueryProposal(transactor.CliCtx, proposalID, govtypes.StatusDepositPeriod)
	require.NoError(t, err, "failed to query proposal 1 with deposit status")
	assert.Equal(t, content.GetTitle(), proposal.GetContent().GetTitle(),
		"The proposal should have same title as submitted proposal")
	assert.Equal(t, content.GetDescription(), proposal.GetContent().GetDescription(),
		"The proposal should have same description as submitted proposal")

	depositMsg := govtypes.NewMsgDeposit(transactor.Key.GetAddress(), proposalID, sdk.NewInt(2e18))
	transactor.AddTxMsg(depositMsg)
	proposal, err = tc.QueryProposal(transactor.CliCtx, proposalID, govtypes.StatusVotingPeriod)
	require.NoError(t, err, "failed to query proposal 1 with voting status")

	byteVoteOption, _ := govtypes.VoteOptionFromString("Yes")
	voteMsg := govtypes.NewMsgVote(transactor.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor.AddTxMsg(voteMsg)
	time.Sleep(10 * time.Second)
	proposal, err = tc.QueryProposal(transactor.CliCtx, proposalID, govtypes.StatusPassed)
	require.NoError(t, err, "failed to query proposal 1 with passed status")

	stakingParams, err := stakingcli.QueryParams(transactor.CliCtx)
	require.NoError(t, err, "failed to query staking params")
	assert.Equal(t, uint64(2), stakingParams.EpochLength, "EpochLength params should be updated to 2")

	log.Info("======================== Test change epochlength rejected ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "EpochLength", "\"5\"")}
	content = govtypes.NewParameterProposal("Guard Param Change", "Update EpochLength", paramChanges)
	submitProposalmsg, err = govtypes.NewMsgSubmitProposal(content, sdk.NewInt(2e18), transactor.Key.GetAddress())
	require.NoError(t, err, "failed to create MsgSubmitProposal")
	transactor.AddTxMsg(submitProposalmsg)

	proposalID = uint64(2)
	proposal, err = tc.QueryProposal(transactor.CliCtx, proposalID, govtypes.StatusVotingPeriod)
	require.NoError(t, err, "failed to query proposal 2 with voting status")

	byteVoteOption, _ = govtypes.VoteOptionFromString("NoWithVeto")
	voteMsg = govtypes.NewMsgVote(transactor.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor.AddTxMsg(voteMsg)
	time.Sleep(10 * time.Second)
	proposal, err = tc.QueryProposal(transactor.CliCtx, proposalID, govtypes.StatusRejected)
	require.NoError(t, err, "failed to query proposal 2 with rejected status")

	stakingParams, err = stakingcli.QueryParams(transactor.CliCtx)
	require.NoError(t, err, "failed to query staking params")
	assert.Equal(t, uint64(2), stakingParams.EpochLength, "EpochLength params should stay 2")
}
