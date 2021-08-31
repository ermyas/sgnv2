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

func setupValidator() []tc.Killable {
	res := setupNewSgnEnv(nil, "validator")
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
	vAmt := big.NewInt(2e18) // 1 CELR
	dAmts := []*big.Int{
		big.NewInt(2e18), // 2 CELR
		big.NewInt(2e18), // 2 CELR
		big.NewInt(4e18), // 4 CELR
		big.NewInt(1e18), // 1 CELR
	}
	totalAmts := tc.NewBigInt(11, 18) // vAmt + dAmts

	vEthAddr, vAuth, err := tc.GetAuth(tc.ValEthKs[0])
	log.Infof("validator eth address %x", vEthAddr)
	require.NoError(t, err, "failed to get validator auth")

	sgnAddr, err := types.SdkAccAddrFromSgnBech32(tc.ValAccounts[0])
	require.NoError(t, err, "failed to get sgnAddr")
	err = tc.InitializeValidator(vAuth, sgnAddr, vAmt, eth.CommissionRate(0.02))
	require.NoError(t, err, "failed to initialize validator")
	expVal := &types.Validator{
		EthAddress:     eth.Addr2Hex(vEthAddr),
		EthSigner:      eth.Addr2Hex(vEthAddr),
		Status:         eth.Bonded,
		SgnAddress:     sgnAddr.String(),
		Tokens:         vAmt.String(),
		Shares:         vAmt.String(),
		CommissionRate: 200,
	}
	tc.CheckValidator(t, transactor, expVal)
	expDel := &types.Delegator{
		ValAddress: eth.Addr2Hex(vEthAddr),
		DelAddress: eth.Addr2Hex(vEthAddr),
		Shares:     vAmt.String(),
	}
	tc.CheckDelegator(t, transactor, expDel)
	expSdkVal := &sdk_staking.Validator{
		OperatorAddress: sgnAddr.String(),
		Status:          sdk_staking.Bonded,
		Tokens:          sdk.NewIntFromBigInt(vAmt),
	}
	tc.CheckSdkValidator(t, transactor, expSdkVal)
	tc.CheckBondedSdkValidatorNum(t, transactor, 1)

	log.Info("add delegators ...")
	for i := 0; i < len(tc.DelEthKs); i++ {
		_, dAuth, err2 := tc.GetAuth(tc.DelEthKs[i])
		require.NoError(t, err2, "failed to get delegator auth")
		go tc.Delegate(dAuth, vEthAddr, dAmts[i])
	}
	for i := 0; i < len(tc.DelEthKs); i++ {
		expDel := &types.Delegator{
			ValAddress: eth.Addr2Hex(vEthAddr),
			DelAddress: tc.DelEthAddrs[i],
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
