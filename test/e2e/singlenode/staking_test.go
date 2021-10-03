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

func setupStaking() []tc.Killable {
	res := setupNewSgnEnv(nil, false)
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
		NodeHome,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
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

	err := tc.InitializeValidator(tc.ValAuths[0], tc.ValSignerAddrs[0], tc.ValSgnAddrs[0], vAmt, eth.CommissionRate(0.02))
	require.NoError(t, err, "failed to initialize validator")
	tc.Sleep(5)
	expVal := &types.Validator{
		EthAddress:      eth.Addr2Hex(tc.ValEthAddrs[0]),
		EthSigner:       eth.Addr2Hex(tc.ValSignerAddrs[0]),
		Status:          eth.Bonded,
		SgnAddress:      tc.ValSgnAddrs[0].String(),
		Tokens:          sdk.NewIntFromBigInt(vAmt),
		DelegatorShares: sdk.NewIntFromBigInt(vAmt),
		CommissionRate:  sdk.NewDecWithPrec(2, 2),
	}
	tc.CheckValidator(t, transactor, expVal)
	tc.CheckValidatorBySgnAddr(t, transactor, expVal)
	expDel := &types.Delegation{
		DelegatorAddress: eth.Addr2Hex(tc.ValEthAddrs[0]),
		ValidatorAddress: eth.Addr2Hex(tc.ValEthAddrs[0]),
		Shares:           sdk.NewIntFromBigInt(vAmt),
	}
	tc.CheckDelegation(t, transactor, expDel)
	tc.PrintTendermintValidators(t, transactor)

	log.Info("add delegators ...")
	for i := 0; i < len(tc.DelEthKs); i++ {
		go tc.Delegate(tc.DelAuths[i], tc.ValEthAddrs[0], dAmts[i])
	}
	tc.Sleep(5)
	tc.PrintTendermintValidators(t, transactor)
	for i := 0; i < len(tc.DelEthKs); i++ {
		expDel := &types.Delegation{
			DelegatorAddress: eth.Addr2Hex(tc.DelEthAddrs[i]),
			ValidatorAddress: eth.Addr2Hex(tc.ValEthAddrs[0]),
			Shares:           sdk.NewIntFromBigInt(dAmts[i]),
		}
		tc.CheckDelegation(t, transactor, expDel)
	}

	expVal.Tokens = sdk.NewIntFromBigInt(totalAmts)
	expVal.DelegatorShares = sdk.NewIntFromBigInt(totalAmts)
	tc.CheckValidator(t, transactor, expVal)
	tc.Sleep(5)
	tc.PrintTendermintValidators(t, transactor)
}
