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
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupSlash(t *testing.T) {
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
	SetupNewSgnEnv(p, false, false, false, false)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestSlash(t *testing.T) {
	t.Run("e2e-slash", func(t *testing.T) {
		t.Run("slashTest", slashTest)
		// t.Run("disableSlashTest", disableSlashTest)
		// t.Run("expireSlashTest", expireSlashTest)
	})
}

// Test penalty slash when a validator is offline
func slashTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test slash ===========================")

	setupSlash(t)

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	amt1 := big.NewInt(8e18)
	amt2 := big.NewInt(1e18)
	amts := []*big.Int{amt1, amt2}
	tc.SetupValidators(t, transactor, amts)

	prev, _ := tc.Contracts.Staking.Validators(&bind.CallOpts{}, tc.ValEthAddrs[1])

	ShutdownNode(1)

	log.Infoln("Query sgn about slash info...")
	nonce := uint64(0)
	slash, err := tc.QuerySlash(transactor.CliCtx, nonce, 1)
	require.NoError(t, err, "failed to query slash")
	log.Infoln("Query sgn about slash info:", slash.String())
	assert.Equal(t, slashingtypes.AttributeValueMissingSignature, slash.Reason)
	assert.Equal(t, tc.ValEthAddrs[1], eth.Bytes2Addr(slash.SlashOnChain.Validator))
	assert.Equal(t, eth.Addr2Hex(tc.ValSignerAddrs[0]), slash.Signatures[0].Signer)

	time.Sleep(5 * time.Second) // wait for onchain call mined
	current, _ := tc.Contracts.Staking.Validators(&bind.CallOpts{}, tc.ValEthAddrs[1])
	for retry := 0; retry < 20; retry++ {
		if prev.Tokens.Cmp(current.Tokens) == 1 {
			break
		}
		tc.Sleep(1)
		current, _ = tc.Contracts.Staking.Validators(&bind.CallOpts{}, tc.ValEthAddrs[1])
	}
	log.Infof("Tokens before slash: %d, after slash: %d", prev.Tokens, current.Tokens)
	assert.True(t, prev.Tokens.Cmp(current.Tokens) == 1)
}

// Test disable slash
func disableSlashTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test disableSlash ===========================")

	setupSlash(t)

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	amt1 := big.NewInt(8e18)
	amt2 := big.NewInt(1e18)
	amts := []*big.Int{amt1, amt2}
	tc.SetupValidators(t, transactor, amts)

	paramChanges := []govtypes.ParamChange{govtypes.NewParamChange("slashing", "EnableSlash", "false")}
	content := govtypes.NewParameterProposal("Slash Param Change", "Update EnableSlash", paramChanges)
	submitProposalmsg, _ := govtypes.NewMsgSubmitProposal(content, sdk.NewInt(1e9), transactor.Key.GetAddress())
	transactor.AddTxMsg(submitProposalmsg)

	proposalID := uint64(1)
	byteVoteOption, _ := govtypes.VoteOptionFromString("Yes")
	voteMsg := govtypes.NewMsgVote(transactor.Key.GetAddress(), proposalID, byteVoteOption)
	transactor.AddTxMsg(voteMsg)

	_, err := tc.QueryProposal(transactor.CliCtx, proposalID, govtypes.StatusPassed)
	require.NoError(t, err, "failed to query proposal 1 with passed status")

	ShutdownNode(1)

	log.Infoln("Query sgn about slash info...")
	nonce := uint64(0)
	_, err = tc.QuerySlash(transactor.CliCtx, nonce, 1)
	assert.Error(t, err, "get penalty 0 with 1 sig should fail")
}

// Test expire slash
func expireSlashTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test expireSlash ===========================")

	setupSlash(t)

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	amt1 := big.NewInt(8e18)
	amt2 := big.NewInt(1e18)
	amts := []*big.Int{amt1, amt2}
	tc.SetupValidators(t, transactor, amts)

	prevBalance, _ := tc.CelrContract.BalanceOf(&bind.CallOpts{}, tc.ValEthAddrs[0])

	paramChanges := []govtypes.ParamChange{govtypes.NewParamChange("slashing", "SlashTimeout", string(transactor.CliCtx.LegacyAmino.MustMarshalJSON(1)))}
	content := govtypes.NewParameterProposal("Slash Param Change", "Update SlashTimeout", paramChanges)
	submitProposalmsg, _ := govtypes.NewMsgSubmitProposal(content, sdk.NewInt(1e9), transactor.Key.GetAddress())
	transactor.AddTxMsg(submitProposalmsg)

	proposalID := uint64(1)
	byteVoteOption, _ := govtypes.VoteOptionFromString("Yes")
	voteMsg := govtypes.NewMsgVote(transactor.Key.GetAddress(), proposalID, byteVoteOption)
	transactor.AddTxMsg(voteMsg)

	_, err := tc.QueryProposal(transactor.CliCtx, proposalID, govtypes.StatusPassed)
	require.NoError(t, err, "failed to query proposal 1 with passed status")

	ShutdownNode(1)

	log.Infoln("Query sgn about slash info...")
	nonce := uint64(0)
	slash, err := tc.QuerySlash(transactor.CliCtx, nonce, 1)
	require.NoError(t, err, "failed to query slash")
	log.Infoln("Query sgn about slash info:", slash.String())

	tc.SleepWithLog(30, "wait for submitting slash")
	currentBalance, _ := tc.CelrContract.BalanceOf(&bind.CallOpts{}, tc.ValEthAddrs[0])
	assert.Equal(t, prevBalance, currentBalance, fmt.Sprintf("The expected balance should be %s", prevBalance))
}
