package multinode

import (
	"math/big"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
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
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.SgnValAcct,
		tc.SgnPassphrase,
	)

	amts := []*big.Int{
		big.NewInt(8e18),
		big.NewInt(5e18),
		big.NewInt(4e18),
		big.NewInt(6e18),
	}

	var expVals types.Validators
	log.Infoln("---------- It should add bonded validators 0 and 1 successfully ----------")
	for i := 0; i < 2; i++ {
		log.Infoln("Adding validator", i, tc.ValEthAddrs[i].Hex())
		err := tc.InitializeValidator(tc.ValAuths[i], tc.ValSgnAddrs[i], amts[i], eth.CommissionRate(0.02))
		require.NoError(t, err, "failed to initialize validator")
		tc.Sleep(5)
		expVal := types.Validator{
			EthAddress:     eth.Addr2Hex(tc.ValEthAddrs[i]),
			EthSigner:      eth.Addr2Hex(tc.ValEthAddrs[i]),
			Status:         eth.Bonded,
			SgnAddress:     tc.ValSgnAddrs[i].String(),
			Tokens:         sdk.NewIntFromBigInt(amts[i]),
			Shares:         sdk.NewIntFromBigInt(amts[i]),
			CommissionRate: eth.CommissionRate(0.02),
		}
		expVals = append(expVals, expVal)
		tc.CheckValidators(t, transactor, expVals)
		expDel := &types.Delegator{
			ValAddress: eth.Addr2Hex(tc.ValEthAddrs[i]),
			DelAddress: eth.Addr2Hex(tc.ValEthAddrs[i]),
			Shares:     amts[i].String(),
		}
		tc.CheckDelegator(t, transactor, expDel)
		expSdkVal := &sdk_staking.Validator{
			OperatorAddress: sdk.ValAddress(tc.ValSgnAddrs[i]).String(),
			Status:          sdk_staking.Bonded,
			Tokens:          sdk.NewIntFromBigInt(amts[i]),
		}
		tc.CheckSdkValidator(t, transactor, expSdkVal)
		tc.CheckBondedSdkValidatorNum(t, transactor, i+1)
	}

	log.Infoln("---------- It should add unbonded validator 2 without enough delegation ----------")
	initialDelegation := big.NewInt(1e18)
	err := tc.InitializeValidator(tc.ValAuths[2], tc.ValSgnAddrs[2], initialDelegation, eth.CommissionRate(0.02))
	require.NoError(t, err, "failed to initialize validator")
	tc.Sleep(5)
	expVal := types.Validator{
		EthAddress:     eth.Addr2Hex(tc.ValEthAddrs[2]),
		EthSigner:      eth.Addr2Hex(tc.ValEthAddrs[2]),
		Status:         eth.Unbonded,
		SgnAddress:     tc.ValSgnAddrs[2].String(),
		Tokens:         sdk.NewIntFromBigInt(initialDelegation),
		Shares:         sdk.NewIntFromBigInt(initialDelegation),
		CommissionRate: eth.CommissionRate(0.02),
	}
	expVals = append(expVals, expVal)
	tc.CheckValidators(t, transactor, expVals)
	tc.CheckBondedSdkValidatorNum(t, transactor, 2)
	tc.PrintTendermintValidators(t, transactor)

	log.Infoln("---------- It should add bonded validator 2 with enough delegation ----------")
	newAmt := big.NewInt(0).Sub(amts[2], initialDelegation)
	err = tc.Delegate(tc.ValAuths[2], tc.ValEthAddrs[2], newAmt)
	require.NoError(t, err, "failed to delegate")
	tc.Sleep(5)
	expVals[2].Status = eth.Bonded
	expVals[2].Tokens = sdk.NewIntFromBigInt(amts[2])
	expVals[2].Shares = sdk.NewIntFromBigInt(amts[2])
	tc.CheckValidators(t, transactor, expVals)
	expSdkVal := &sdk_staking.Validator{
		OperatorAddress: sdk.ValAddress(tc.ValSgnAddrs[2]).String(),
		Status:          sdk_staking.Bonded,
		Tokens:          sdk.NewIntFromBigInt(amts[2]),
	}
	tc.CheckSdkValidator(t, transactor, expSdkVal)
	tc.CheckBondedSdkValidatorNum(t, transactor, 3)
	tc.Sleep(5)
	tc.PrintTendermintValidators(t, transactor)

	log.Infoln("---------- It should unbond validator 2 caused by undelegation ----------")
	err = tc.Undelegate(tc.ValAuths[2], tc.ValEthAddrs[2], newAmt)
	require.NoError(t, err, "failed to undelegate")
	tc.Sleep(5)
	expVals[2].Status = eth.Unbonding
	expVals[2].Tokens = sdk.NewIntFromBigInt(initialDelegation)
	expVals[2].Shares = sdk.NewIntFromBigInt(initialDelegation)
	tc.CheckValidators(t, transactor, expVals)
	tc.CheckBondedSdkValidatorNum(t, transactor, 2)

	tc.ConfirmUnbondedValidator(tc.ValAuths[2], tc.ValEthAddrs[2])
	expVals[2].Status = eth.Unbonded
	tc.CheckValidators(t, transactor, expVals)

	log.Infoln("---------- It should add back bonded validator 2 with enough delegation ----------")
	err = tc.Delegate(tc.ValAuths[2], tc.ValEthAddrs[2], newAmt)
	tc.Sleep(5)
	expVals[2].Status = eth.Bonded
	expVals[2].Tokens = sdk.NewIntFromBigInt(amts[2])
	expVals[2].Shares = sdk.NewIntFromBigInt(amts[2])
	tc.CheckValidators(t, transactor, expVals)
	tc.CheckBondedSdkValidatorNum(t, transactor, 3)

	log.Infoln("---------- It should correctly replace bonded validator 2 with validator 3 ----------")
	err = tc.InitializeValidator(tc.ValAuths[3], tc.ValSgnAddrs[3], amts[3], eth.CommissionRate(0.02))
	tc.Sleep(5)
	require.NoError(t, err, "failed to initialize validator")
	expVals[2].Status = eth.Unbonding
	expVal = types.Validator{
		EthAddress:     eth.Addr2Hex(tc.ValEthAddrs[3]),
		EthSigner:      eth.Addr2Hex(tc.ValEthAddrs[3]),
		Status:         eth.Bonded,
		SgnAddress:     tc.ValSgnAddrs[3].String(),
		Tokens:         sdk.NewIntFromBigInt(amts[3]),
		Shares:         sdk.NewIntFromBigInt(amts[3]),
		CommissionRate: eth.CommissionRate(0.02),
	}
	expVals = append(expVals, expVal)
	tc.CheckValidators(t, transactor, expVals)
	expSdkVal = &sdk_staking.Validator{
		OperatorAddress: sdk.ValAddress(tc.ValSgnAddrs[3]).String(),
		Status:          sdk_staking.Bonded,
		Tokens:          sdk.NewIntFromBigInt(amts[3]),
	}
	tc.CheckSdkValidator(t, transactor, expSdkVal)
	tc.CheckBondedSdkValidatorNum(t, transactor, 3)
	tc.Sleep(5)
	tc.PrintTendermintValidators(t, transactor)
}
