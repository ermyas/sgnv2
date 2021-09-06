package singlenode

import (
	"math/big"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func setupStaking() []tc.Killable {
	res := setupNewSgnEnv(nil, "staking")
	tc.SleepWithLog(10, "sgn being ready")

	return res
}

func TestStaking(t *testing.T) {
	toKill := setupStaking()
	defer tc.TearDown(toKill)

	t.Run("e2e-staking", func(t *testing.T) {
		t.Run("stakingTest", stakingTest)
	})
}

func stakingTest(t *testing.T) {
	log.Info("===================================================================")
	log.Info("======================== Test staking ===========================")

	transactor := tc.NewTestTransactor(
		t,
		CLIHome,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetStringSlice(common.FlagSgnTransactors)[0],
		viper.GetString(common.FlagSgnPassphrase),
	)
	vAmt := big.NewInt(2e18)
	dAmts := []*big.Int{
		big.NewInt(2e18),
		big.NewInt(2e18),
		big.NewInt(4e18),
		big.NewInt(1e18),
	}
	totalAmts := tc.NewBigInt(11, 18) // vAmt + dAmts

	err := tc.InitializeValidator(tc.ValAuths[0], tc.ValSgnAddrs[0], vAmt, eth.CommissionRate(0.02))
	require.NoError(t, err, "failed to initialize validator")
	tc.Sleep(5)
	expVal := &types.Validator{
		EthAddress:     eth.Addr2Hex(tc.ValEthAddrs[0]),
		EthSigner:      eth.Addr2Hex(tc.ValEthAddrs[0]),
		Status:         eth.Bonded,
		SgnAddress:     tc.ValSgnAddrs[0].String(),
		Tokens:         vAmt.String(),
		Shares:         vAmt.String(),
		CommissionRate: eth.CommissionRate(0.02),
	}
	tc.CheckValidator(t, transactor, expVal)
	expDel := &types.Delegator{
		ValAddress: eth.Addr2Hex(tc.ValEthAddrs[0]),
		DelAddress: eth.Addr2Hex(tc.ValEthAddrs[0]),
		Shares:     vAmt.String(),
	}
	tc.CheckDelegator(t, transactor, expDel)
	expSdkVal := &sdk_staking.Validator{
		OperatorAddress: sdk.ValAddress(tc.ValSgnAddrs[0]).String(),
		Status:          sdk_staking.Bonded,
		Tokens:          sdk.NewIntFromBigInt(vAmt),
	}
	tc.CheckSdkValidator(t, transactor, expSdkVal)
	tc.CheckBondedSdkValidatorNum(t, transactor, 1)
	tc.PrintTendermintValidators(t, transactor)

	log.Info("add delegators ...")
	for i := 0; i < len(tc.DelEthKs); i++ {
		go tc.Delegate(tc.DelAuths[i], tc.ValEthAddrs[0], dAmts[i])
	}
	tc.Sleep(5)
	for i := 0; i < len(tc.DelEthKs); i++ {
		expDel := &types.Delegator{
			ValAddress: eth.Addr2Hex(tc.ValEthAddrs[0]),
			DelAddress: eth.Addr2Hex(tc.DelEthAddrs[i]),
			Shares:     dAmts[i].String(),
		}
		tc.CheckDelegator(t, transactor, expDel)
	}

	expVal.Tokens = totalAmts.String()
	expVal.Shares = totalAmts.String()
	tc.CheckValidator(t, transactor, expVal)
	expSdkVal.Tokens = sdk.NewIntFromBigInt(totalAmts)
	tc.CheckSdkValidator(t, transactor, expSdkVal)
}
