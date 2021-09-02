package multinode

import (
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/stretchr/testify/require"
)

func setupStaking() {
	log.Infoln("set up new sgn env")
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

func TestStaking(t *testing.T) {
	t.Run("e2e-staking", func(t *testing.T) {
		t.Run("stakingTest", stakingTest)
	})
}

func stakingTest(t *testing.T) {
	log.Info("===================================================================")
	log.Info("======================== Test staking ===========================")
	setupStaking()

	transactor := tc.NewTestTransactor(
		t,
		tc.SgnCLIHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.SgnCLIAddr,
		tc.SgnPassphrase,
	)

	amts := []*big.Int{
		big.NewInt(8e18),
		big.NewInt(5e18),
		big.NewInt(4e18),
		big.NewInt(6e18),
	}

	expVals := make([]*types.Validator, 0)
	log.Infoln("---------- It should add two validators successfully ----------")
	for i := 0; i < 2; i++ {
		log.Infoln("Adding validator", i)
		ethAddr, auth, err := tc.GetAuth(tc.ValEthKs[i])
		log.Infof("validator eth address %x", ethAddr)
		require.NoError(t, err, "failed to get auth")
		sgnAddr, err := types.SdkAccAddrFromSgnBech32(tc.ValAccounts[i])
		require.NoError(t, err, "failed to get sgnAddr")
		err = tc.InitializeValidator(auth, sgnAddr, amts[i], eth.CommissionRate(0.02))
		require.NoError(t, err, "failed to initialize validator")
		expVal := &types.Validator{
			EthAddress:     eth.Addr2Hex(ethAddr),
			EthSigner:      eth.Addr2Hex(ethAddr),
			Status:         eth.Bonded,
			SgnAddress:     sgnAddr.String(),
			Tokens:         amts[i].String(),
			Shares:         amts[i].String(),
			CommissionRate: 200,
		}
		expVals = append(expVals, expVal)
		tc.CheckValidators(t, transactor, expVals)
	}

	log.Infoln("---------- It should fail to add validator 2 without enough self delegation ----------")
	ethAddr, auth, err := tc.GetAuth(tc.ValEthKs[2])
	require.NoError(t, err, "failed to get auth")
	sgnAddr, err := types.SdkAccAddrFromSgnBech32(tc.ValAccounts[2])
	require.NoError(t, err, "failed to get sgnAddr")
	initialDelegation := big.NewInt(1e17)
	err = tc.InitializeValidator(auth, sgnAddr, initialDelegation, eth.CommissionRate(0.02))
	require.NoError(t, err, "failed to initialize validator")
	time.Sleep(10 * time.Second)               // wait for processing
	tc.CheckValidators(t, transactor, expVals) // still the previous two

	log.Infoln("---------- It should correctly add validator 2 with enough delegation ----------")
	err = tc.Delegate(auth, ethAddr, big.NewInt(0).Sub(amts[2], initialDelegation))
	require.NoError(t, err, "failed to delegate stake")
	expVal := &types.Validator{
		EthAddress:     eth.Addr2Hex(ethAddr),
		EthSigner:      eth.Addr2Hex(ethAddr),
		Status:         eth.Bonded,
		SgnAddress:     sgnAddr.String(),
		Tokens:         amts[2].String(),
		Shares:         amts[2].String(),
		CommissionRate: 200,
	}
	expVals = append(expVals, expVal)
	tc.CheckValidators(t, transactor, expVals)

	log.Infoln("---------- It should correctly replace validator 2 with validator 3 ----------")
	ethAddr, auth, err = tc.GetAuth(tc.ValEthKs[3])
	require.NoError(t, err, "failed to get auth")
	sgnAddr, err = types.SdkAccAddrFromSgnBech32(tc.ValAccounts[3])
	require.NoError(t, err, "failed to get sgnAddr")
	err = tc.InitializeValidator(auth, sgnAddr, amts[3], eth.CommissionRate(0.02))
	require.NoError(t, err, "failed to initialize validator")
	expVals = expVals[:2]
	expVal = &types.Validator{
		EthAddress:     eth.Addr2Hex(ethAddr),
		EthSigner:      eth.Addr2Hex(ethAddr),
		Status:         eth.Bonded,
		SgnAddress:     sgnAddr.String(),
		Tokens:         amts[3].String(),
		Shares:         amts[3].String(),
		CommissionRate: 200,
	}
	expVals = append(expVals, expVal)
	tc.CheckValidators(t, transactor, expVals)
}
