package multinode

import (
	"math/big"
	"testing"

	"github.com/celer-network/goutils/log"
	tc "github.com/celer-network/sgn-v2/test/common"
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

	// transactor := tc.NewTestTransactor(
	// 	t,
	// 	tc.SgnCLIHomes[0],
	// 	tc.SgnChainID,
	// 	tc.SgnNodeURI,
	// 	tc.SgnCLIAddr,
	// 	tc.SgnPassphrase,
	// )

	// amts := []*big.Int{
	// 	big.NewInt(8000000000000000000),
	// 	big.NewInt(5000000000000000000),
	// 	big.NewInt(4000000000000000000),
	// 	big.NewInt(6000000000000000000),
	// }
	// minAmts := []*big.Int{
	// 	big.NewInt(4000000000000000000),
	// 	big.NewInt(2000000000000000000),
	// 	big.NewInt(2000000000000000000),
	// 	big.NewInt(2000000000000000000),
	// }
	// commissionRate := big.NewInt(200)

	// log.Infoln("---------- It should add two validators successfully ----------")
	// for i := 0; i < 2; i++ {
	// 	log.Infoln("Adding validator", i)
	// 	ethAddr, auth, err := tc.GetAuth(tc.ValEthKs[i])
	// 	require.NoError(t, err, "failed to get auth")
	// 	tc.AddValidatorWithStake(
	// 		t, transactor, ethAddr, auth, tc.ValAccounts[i],
	// 		amts[i], minAmts[i], commissionRate, big.NewInt(10000), true)
	// 	tc.CheckValidatorNum(t, transactor, i+1)
	// }

	// log.Infoln("---------- It should fail to add validator 2 without enough delegation ----------")
	// ethAddr, auth, err := tc.GetAuth(tc.ValEthKs[2])
	// require.NoError(t, err, "failed to get auth")
	// initialDelegation := big.NewInt(1000000000000000000)
	// tc.AddValidatorWithStake(
	// 	t, transactor, ethAddr, auth, tc.ValAccounts[2],
	// 	initialDelegation, minAmts[2], commissionRate, big.NewInt(10000), false)
	// log.Info("Query sgn about validators to check if validator 2 is not added...")
	// tc.CheckValidatorNum(t, transactor, 2)

	// log.Infoln("---------- It should correctly add validator 2 with enough delegation ----------")
	// err = tc.Delegate(auth, ethAddr, big.NewInt(0).Sub(amts[2], initialDelegation))
	// require.NoError(t, err, "failed to delegate stake")
	// tc.CheckValidatorNum(t, transactor, 3)
	// tc.CheckValidator(t, transactor, tc.ValAccounts[2], amts[2], sdk.Bonded)

	// log.Infoln("---------- It should successfully remove validator 2 caused by intendWithdraw ----------")
	// err = tc.IntendWithdraw(auth, ethAddr, amts[2])
	// require.NoError(t, err, "failed to intendWithdraw stake")
	// log.Info("Query sgn about the validators to check if it has correct number of validators...")
	// tc.CheckValidatorNum(t, transactor, 2)
	// tc.CheckValidatorStatus(t, transactor, tc.ValAccounts[2], sdk.Unbonding)

	// err = tc.ConfirmUnbondedCandidate(auth, ethAddr)
	// require.NoError(t, err, "failed to confirmUnbondedCandidate")
	// tc.CheckValidator(t, transactor, ethAddr, tc.ValAccounts[2], big.NewInt(0))

	// log.Infoln("---------- It should successfully add back validator 2 with enough delegation ----------")
	// err = tc.Delegate(auth, ethAddr, amts[2])
	// require.NoError(t, err, "failed to delegate stake")
	// tc.CheckValidatorNum(t, transactor, 3)
	// tc.CheckValidator(t, transactor, tc.ValAccounts[2], amts[2], sdk.Bonded)

	// log.Infoln("---------- It should correctly replace validator 2 with validator 3 ----------")
	// ethAddr, auth, err = tc.GetAuth(tc.ValEthKs[3])
	// require.NoError(t, err, "failed to get auth")
	// tc.AddValidatorWithStake(
	// 	t, transactor, ethAddr, auth, tc.ValAccounts[3],
	// 	amts[3], minAmts[3], commissionRate, big.NewInt(10000), true)
	// tc.CheckValidatorNum(t, transactor, 3)
	// tc.CheckValidator(t, transactor, tc.ValAccounts[2], amts[2], sdk.Unbonding)
}