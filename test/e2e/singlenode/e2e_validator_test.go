package singlenode

import (
	"math/big"
	"strings"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func setupValidator() []tc.Killable {
	p := &tc.ContractParams{
		CelrAddr:              tc.E2eProfile.CelrAddr,
		ProposalDeposit:       big.NewInt(1),
		VotePeriod:            big.NewInt(1),
		UnbondingPeriod:       big.NewInt(10),
		MaxBondedValidators:   big.NewInt(11),
		MinValidatorTokens:    big.NewInt(1e18),
		MinSelfDelegation:     big.NewInt(1e18),
		AdvanceNoticePeriod:   big.NewInt(1),
		ValidatorBondInterval: big.NewInt(0),
		MaxSlashFactor:        big.NewInt(1e5),
	}
	res := setupNewSGNEnv(p, "validator")
	tc.SleepWithLog(10, "sgn being ready")

	return res
}

func TestE2EValidator(t *testing.T) {
	toKill := setupValidator()
	defer tc.TearDown(toKill)

	t.Run("e2e-validator", func(t *testing.T) {
		t.Run("validatorTest", validatorTest)
	})
}

func validatorTest(t *testing.T) {
	log.Info("===================================================================")
	log.Info("======================== Test validator ===========================")

	transactor := tc.NewTestTransactor(
		t,
		CLIHome,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetStringSlice(common.FlagSgnTransactors)[0],
		viper.GetString(common.FlagSgnPassphrase),
	)
	//vAmt := big.NewInt(1000000000000000000) // 1 CELR
	dAmts := []*big.Int{
		big.NewInt(2000000000000000000), // 2 CELR
		big.NewInt(2000000000000000000), // 2 CELR
		big.NewInt(4000000000000000000), // 4 CELR
		big.NewInt(1000000000000000000), // 1 CELR
	}
	miningPool := new(big.Int)
	miningPool.SetString("1"+strings.Repeat("0", 20), 10)

	vEthAddr, _, err := tc.GetAuth(tc.ValEthKs[0])
	log.Infof("validator eth address %x", vEthAddr)
	require.NoError(t, err, "failed to get validator auth")

	log.Info("add delegators ...")
	for i := 0; i < len(tc.DelEthKs); i++ {
		_, dAuth, err2 := tc.GetAuth(tc.DelEthKs[i])
		require.NoError(t, err2, "failed to get delegator auth")
		go tc.DelegateStake(dAuth, vEthAddr, dAmts[i])
	}
	for i := 0; i < len(tc.DelEthKs); i++ {
		tc.CheckDelegator(t, transactor, vEthAddr, eth.Hex2Addr(tc.DelEthAddrs[i]), dAmts[i])
	}
}
