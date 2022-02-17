package multinode

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/transactor"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupBridgeTest() {
	log.Infoln("Set up new sgn env")
	p := &tc.ContractParams{
		CelrAddr:              tc.CelrAddr,
		ProposalDeposit:       big.NewInt(1e18),
		VotePeriod:            big.NewInt(5),
		UnbondingPeriod:       big.NewInt(5),
		MaxBondedValidators:   big.NewInt(3),
		MinValidatorTokens:    big.NewInt(1e18),
		MinSelfDelegation:     big.NewInt(1e18),
		AdvanceNoticePeriod:   big.NewInt(1),
		ValidatorBondInterval: big.NewInt(0),
		MaxSlashFactor:        big.NewInt(1e5),
	}
	SetupNewSgnEnv(p, true, false, false, false)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestBridge(t *testing.T) {
	t.Run("e2e-bridge", func(t *testing.T) {
		t.Run("bridgeTest", bridgeTest)
	})
}

// Test pegbridge
func bridgeTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("============ Test Bridge (Both cBridge and pegBridge) =============")
	setupBridgeTest()

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	prepareValidators(t, transactor)
	govSyncerCandidates(t)

	tc.RunAllAndWait(
		func() {
			cbrTest(t, transactor)
		},
		func() {
			pbrTest1(t, transactor)
		},
		func() {
			pbrTest2(t, transactor)
		},
	)
}

func prepareValidators(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("================== Prepare validators start =================")

	log.Infoln("================== Setup validators ======================")
	// Make the stake amounts more realistic to test precision handling when distributing fee share
	vAmt := new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18))
	vAmts := []*big.Int{vAmt, vAmt, vAmt}
	tc.SetupValidators(t, transactor, vAmts)

	log.Infoln("================== Setup bridge signers ======================")
	tc.CbrChain1.SetInitSigners(vAmts)
	tc.CbrChain2.SetInitSigners(vAmts)
	tc.CbrChain3.SetInitSigners(vAmts)

	log.Infoln("================== Delegate from delegator 0 to all validators ======================")
	valAddrs := []eth.Addr{tc.ValEthAddrs[0], tc.ValEthAddrs[1], tc.ValEthAddrs[2]}
	dAmt := new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18))
	dAmts := []*big.Int{dAmt, dAmt, dAmt}
	tc.MultiDelegate(tc.DelAuths[0], valAddrs, dAmts)
	for i := 0; i < 3; i++ {
		expDel := &stakingtypes.Delegation{
			DelegatorAddress: eth.Addr2Hex(tc.DelEthAddrs[0]),
			ValidatorAddress: eth.Addr2Hex(tc.ValEthAddrs[i]),
			Shares:           sdk.NewIntFromBigInt(dAmts[i]),
		}
		tc.CheckDelegation(t, transactor, expDel)
	}

	expSigners := genSortedSigners([]eth.Addr{tc.ValSignerAddrs[0], tc.ValSignerAddrs[1], tc.ValSignerAddrs[2]}, vAmts)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain3.ChainId, expSigners)

	log.Infoln("================== Prepare validators done =================")
}

func govSyncerCandidates(t *testing.T) {
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

	paramChanges := []govtypes.ParamChange{govtypes.NewParamChange("staking", "SyncerCandidates", fmt.Sprintf("[\"%s\", \"%s\"]", tc.ValEthAddrs[0], tc.ValEthAddrs[1]))}
	content := govtypes.NewParameterProposal("Subscribe Param Change", "Update SyncerCandidates", paramChanges)
	submitProposalmsg, _ := govtypes.NewMsgSubmitProposal(content, sdk.NewInt(1e18), transactor0.Key.GetAddress())
	transactor0.AddTxMsg(submitProposalmsg)

	proposalID := uint64(1)
	proposal, err := tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusVotingPeriod)
	require.NoError(t, err, "failed to query proposal 1 with voting status")

	byteVoteOption, _ := govtypes.VoteOptionFromString("Yes")
	voteMsg := govtypes.NewMsgVote(transactor0.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor0.AddTxMsg(voteMsg)
	voteMsg = govtypes.NewMsgVote(transactor1.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor1.AddTxMsg(voteMsg)
	voteMsg = govtypes.NewMsgVote(transactor2.Key.GetAddress(), proposal.ProposalId, byteVoteOption)
	transactor2.AddTxMsg(voteMsg)

	time.Sleep(10 * time.Second)
	proposal, err = tc.QueryProposal(transactor0.CliCtx, proposalID, govtypes.StatusPassed)
	require.NoError(t, err, "failed to query proposal 1 with passed status")

	stakingParams, err := stakingcli.QueryParams(transactor0.CliCtx)
	require.NoError(t, err, "failed to query staking params")
	candidates, _ := json.Marshal(stakingParams.SyncerCandidates)
	log.Infoln("SyncerCandidates:", string(candidates))
}
