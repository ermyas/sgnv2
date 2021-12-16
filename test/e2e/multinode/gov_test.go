package multinode

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	tc "github.com/celer-network/sgn-v2/test/common"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGov(t *testing.T) {
	log.Infoln("Set up new sgn env")
	p := &tc.ContractParams{
		CelrAddr:              tc.CelrAddr,
		ProposalDeposit:       big.NewInt(1e18),
		VotePeriod:            big.NewInt(5),
		UnbondingPeriod:       big.NewInt(5),
		MaxBondedValidators:   big.NewInt(3),
		MinValidatorTokens:    big.NewInt(2e18),
		MinSelfDelegation:     big.NewInt(1e18),
		AdvanceNoticePeriod:   big.NewInt(1),
		ValidatorBondInterval: big.NewInt(0),
		MaxSlashFactor:        big.NewInt(1e5),
	}
	SetupNewSgnEnv(p, false, false)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestGov(t *testing.T) {
	t.Run("e2e-gov", func(t *testing.T) {
		t.Run("sgnchainGovTest", sgnchainGovTest)
	})
}

func sgnchainGovTest(t *testing.T) {
	log.Info("=====================================================================")
	log.Info("======================== Test sgnchain gov ===========================")

	setupGov(t)

	transactor0 := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	transactor1 := tc.NewTestTransactor(
		tc.SgnHomes[1],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[1],
		tc.SgnPassphrase,
	)

	transactor2 := tc.NewTestTransactor(
		tc.SgnHomes[2],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[2],
		tc.SgnPassphrase,
	)

	amt1 := big.NewInt(3e18)
	amt2 := big.NewInt(2e18)
	amt3 := big.NewInt(2e18)
	amts := []*big.Int{amt1, amt2, amt3}
	tc.SetupValidators(t, transactor0, amts)

	log.Info("======================== Test change syncerDuration rejected due to small quorum ===========================")
	paramChanges := []govtypes.ParamChange{govtypes.NewParamChange("staking", "SyncerDuration", "\"2\"")}
	content := govtypes.NewParameterProposal("Guard Param Change", "Update SyncerDuration", paramChanges)
	submitProposalmsg, _ := govtypes.NewMsgSubmitProposal(content, sdk.NewInt(1e18), transactor1.Key.GetAddress())
	transactor1.AddTxMsg(submitProposalmsg)

	proposalID := uint64(1)
	proposal, err := tc.QueryProposal(transactor1.CliCtx, proposalID, govtypes.StatusVotingPeriod)
	require.NoError(t, err, "failed to query proposal 1 with voting status")

	byteVoteOption, _ := govtypes.VoteOptionFromString("No")
	voteMsg := govtypes.NewMsgVote(transactor1.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor1.AddTxMsg(voteMsg)
	time.Sleep(10 * time.Second)
	proposal, err = tc.QueryProposal(transactor1.CliCtx, proposalID, govtypes.StatusRejected)
	require.NoError(t, err, "failed to query proposal 1 with rejected status")

	stakingParams, err := stakingcli.QueryParams(transactor1.CliCtx)
	require.NoError(t, err, "failed to query staking params")
	assert.Equal(t, stakingtypes.DefaultSyncerDuration, stakingParams.SyncerDuration,
		fmt.Sprintf("SyncerDuration params should stay %d", stakingtypes.DefaultSyncerDuration))

	log.Info("======================== Test change syncerDuration passed for reaching quorum ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "SyncerDuration", "\"2\"")}
	content = govtypes.NewParameterProposal("Guard Param Change", "Update SyncerDuration", paramChanges)
	submitProposalmsg, _ = govtypes.NewMsgSubmitProposal(content, sdk.NewInt(1e18), transactor0.Key.GetAddress())
	transactor0.AddTxMsg(submitProposalmsg)

	proposalID = uint64(2)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusVotingPeriod)
	require.NoError(t, err, "failed to query proposal 2 with voting status")

	byteVoteOption, _ = govtypes.VoteOptionFromString("Yes")
	voteMsg = govtypes.NewMsgVote(transactor0.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor0.AddTxMsg(voteMsg)
	time.Sleep(10 * time.Second)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusPassed)
	require.NoError(t, err, "failed to query proposal 2 with passed status")

	stakingParams, err = stakingcli.QueryParams(transactor0.CliCtx)
	require.NoError(t, err, "failed to query staking params")
	assert.Equal(t, uint64(2), stakingParams.SyncerDuration, "SyncerDuration params should change to 2")

	log.Info("======================== Test change syncerDuration rejected due to 1/3 veto ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "SyncerDuration", "\"5\"")}
	content = govtypes.NewParameterProposal("Guard Param Change", "Update SyncerDuration", paramChanges)
	submitProposalmsg, _ = govtypes.NewMsgSubmitProposal(content, sdk.NewInt(1e18), transactor1.Key.GetAddress())
	transactor1.AddTxMsg(submitProposalmsg)

	proposalID = uint64(3)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusVotingPeriod)
	require.NoError(t, err, "failed to query proposal 3 with voting status")

	byteVoteOption, _ = govtypes.VoteOptionFromString("NoWithVeto")
	voteMsg = govtypes.NewMsgVote(transactor0.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor0.AddTxMsg(voteMsg)
	byteVoteOption, _ = govtypes.VoteOptionFromString("Yes")
	voteMsg = govtypes.NewMsgVote(transactor1.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor1.AddTxMsg(voteMsg)
	voteMsg = govtypes.NewMsgVote(transactor2.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor2.AddTxMsg(voteMsg)

	time.Sleep(10 * time.Second)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusRejected)
	require.NoError(t, err, "failed to query proposal 3 with rejected status")

	stakingParams, err = stakingcli.QueryParams(transactor0.CliCtx)
	require.NoError(t, err, "failed to query staking params")
	assert.Equal(t, uint64(2), stakingParams.SyncerDuration, "SyncerDuration params should stay 2")

	log.Info("======================== Test change syncerDuration rejected due to 1/2 No ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "SyncerDuration", "\"5\"")}
	content = govtypes.NewParameterProposal("Guard Param Change", "Update SyncerDuration", paramChanges)
	submitProposalmsg, _ = govtypes.NewMsgSubmitProposal(content, sdk.NewInt(1e18), transactor2.Key.GetAddress())
	transactor2.AddTxMsg(submitProposalmsg)

	proposalID = uint64(4)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusVotingPeriod)
	require.NoError(t, err, "failed to query proposal 4 with voting status")

	byteVoteOption, _ = govtypes.VoteOptionFromString("No")
	voteMsg = govtypes.NewMsgVote(transactor0.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor0.AddTxMsg(voteMsg)
	byteVoteOption, _ = govtypes.VoteOptionFromString("Yes")
	voteMsg = govtypes.NewMsgVote(transactor1.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor1.AddTxMsg(voteMsg)

	time.Sleep(10 * time.Second)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusRejected)
	require.NoError(t, err, "failed to query proposal 4 with rejected status")

	stakingParams, err = stakingcli.QueryParams(transactor0.CliCtx)
	require.NoError(t, err, "failed to query staking params")
	assert.Equal(t, uint64(2), stakingParams.SyncerDuration, "SyncerDuration params should stay 2")

	log.Info("======================== Test change syncerDuration passed for over 1/2 yes ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "SyncerDuration", "\"5\"")}
	content = govtypes.NewParameterProposal("Subscribe Param Change", "Update SyncerDuration", paramChanges)
	submitProposalmsg, _ = govtypes.NewMsgSubmitProposal(content, sdk.NewInt(1e18), transactor2.Key.GetAddress())
	transactor2.AddTxMsg(submitProposalmsg)

	proposalID = uint64(5)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusVotingPeriod)
	require.NoError(t, err, "failed to query proposal 5 with voting status")

	byteVoteOption, _ = govtypes.VoteOptionFromString("No")
	voteMsg = govtypes.NewMsgVote(transactor2.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor2.AddTxMsg(voteMsg)
	byteVoteOption, _ = govtypes.VoteOptionFromString("Yes")
	voteMsg = govtypes.NewMsgVote(transactor0.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor0.AddTxMsg(voteMsg)

	time.Sleep(10 * time.Second)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusPassed)
	require.NoError(t, err, "failed to query proposal 5 with passed status")

	stakingParams, err = stakingcli.QueryParams(transactor0.CliCtx)
	require.NoError(t, err, "failed to query staking params")
	assert.Equal(t, uint64(5), stakingParams.SyncerDuration, "SyncerDuration params should change to 5")
}
