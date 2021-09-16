package multinode

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGov() {
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
	SetupNewSgnEnv(p, false)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestGov(t *testing.T) {
	t.Run("e2e-gov", func(t *testing.T) {
		t.Run("sidechainGovTest", sidechainGovTest)
	})
}

func sidechainGovTest(t *testing.T) {
	log.Info("=====================================================================")
	log.Info("======================== Test sidechain gov ===========================")

	setupGov()

	transactor0 := tc.NewTestTransactor(
		t,
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	transactor1 := tc.NewTestTransactor(
		t,
		tc.SgnHomes[1],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[1],
		tc.SgnPassphrase,
	)

	transactor2 := tc.NewTestTransactor(
		t,
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
	var expVals stakingtypes.Validators
	log.Infoln("---------- It should add bonded validators 0, 1 and 2 successfully ----------")
	for i := 0; i < 3; i++ {
		log.Infoln("Adding validator", i, tc.ValEthAddrs[i].Hex())
		err := tc.InitializeValidator(tc.ValAuths[i], tc.ValSignerAddrs[i], tc.ValSgnAddrs[i], amts[i], eth.CommissionRate(0.02))
		require.NoError(t, err, "failed to initialize validator")
		tc.Sleep(5)
		expVal := stakingtypes.Validator{
			EthAddress:      eth.Addr2Hex(tc.ValEthAddrs[i]),
			EthSigner:       eth.Addr2Hex(tc.ValSignerAddrs[i]),
			Status:          eth.Bonded,
			SgnAddress:      tc.ValSgnAddrs[i].String(),
			Tokens:          sdk.NewIntFromBigInt(amts[i]),
			DelegatorShares: sdk.NewIntFromBigInt(amts[i]),
			CommissionRate:  sdk.NewDecWithPrec(2, 2),
		}
		expVals = append(expVals, expVal)
		tc.CheckValidators(t, transactor0, expVals)
	}

	log.Info("======================== Test change epochlength rejected due to small quorum ===========================")
	paramChanges := []govtypes.ParamChange{govtypes.NewParamChange("staking", "EpochLength", "\"2\"")}
	content := govtypes.NewParameterProposal("Guard Param Change", "Update EpochLength", paramChanges)
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
	assert.Equal(t, stakingtypes.DefaultEpochLength, stakingParams.EpochLength, fmt.Sprintf("EpochLength params should stay %d", stakingtypes.DefaultEpochLength))

	log.Info("======================== Test change epochlength passed for reaching quorum ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "EpochLength", "\"2\"")}
	content = govtypes.NewParameterProposal("Guard Param Change", "Update EpochLength", paramChanges)
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
	assert.Equal(t, uint64(2), stakingParams.EpochLength, "EpochLength params should change to 2")

	log.Info("======================== Test change epochlength rejected due to 1/3 veto ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "EpochLength", "\"5\"")}
	content = govtypes.NewParameterProposal("Guard Param Change", "Update EpochLength", paramChanges)
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
	assert.Equal(t, uint64(2), stakingParams.EpochLength, "EpochLength params should stay 2")

	log.Info("======================== Test change epochlength rejected due to 1/2 No ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "EpochLength", "\"5\"")}
	content = govtypes.NewParameterProposal("Guard Param Change", "Update EpochLength", paramChanges)
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
	assert.Equal(t, uint64(2), stakingParams.EpochLength, "EpochLength params should stay 2")

	log.Info("======================== Test change epochlength passed for over 1/2 yes ===========================")
	paramChanges = []govtypes.ParamChange{govtypes.NewParamChange("staking", "EpochLength", "\"5\"")}
	content = govtypes.NewParameterProposal("Gubscribe Param Change", "Update EpochLength", paramChanges)
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
	assert.Equal(t, uint64(5), stakingParams.EpochLength, "EpochLength params should change to 5")
}