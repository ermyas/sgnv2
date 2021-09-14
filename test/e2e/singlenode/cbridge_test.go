package singlenode

import (
	"math/big"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func setupCbridge() []tc.Killable {
	res := setupNewSgnEnv(nil, "cbridge")
	tc.SleepWithLog(10, "sgn being ready")

	return res
}

func TestCbridge(t *testing.T) {
	toKill := setupCbridge()
	defer tc.TearDown(toKill)

	t.Run("e2e-cbridge", func(t *testing.T) {
		t.Run("cbridgeTest", cbridgeTest)
	})
}

func cbridgeTest(t *testing.T) {
	log.Info("===================================================================")
	log.Info("======================== Test cbridge ===========================")

	transactor := tc.NewTestTransactor(
		t,
		NodeHome,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetStringSlice(common.FlagSgnTransactors)[0],
		viper.GetString(common.FlagSgnPassphrase),
	)

	amt := big.NewInt(5e18)
	err := tc.InitializeValidator(tc.ValAuths[0], tc.ValSgnAddrs[0], amt, eth.CommissionRate(0.02))
	require.NoError(t, err, "failed to initialize validator")
	tc.Sleep(5)
	expVal := &types.Validator{
		EthAddress:      eth.Addr2Hex(tc.ValEthAddrs[0]),
		EthSigner:       eth.Addr2Hex(tc.ValEthAddrs[0]),
		Status:          eth.Bonded,
		SgnAddress:      tc.ValSgnAddrs[0].String(),
		Tokens:          sdk.NewIntFromBigInt(amt),
		DelegatorShares: sdk.NewIntFromBigInt(amt),
		CommissionRate:  sdk.NewDecWithPrec(2, 2),
	}
	tc.CheckValidator(t, transactor, expVal)

	log.Info("======================== Test Send ===========================")

}
