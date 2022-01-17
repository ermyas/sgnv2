package keeper

import (
	"github.com/armon/go-metrics"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) withdrawDelegatorRewardForOneValidator(ctx sdk.Context, delAddrStr string, valAddrStr string) error {
	valAddr := eth.Hex2Addr(valAddrStr)
	delAddr := eth.Hex2Addr(delAddrStr)
	amount, err := k.WithdrawDelegationRewards(ctx, delAddr, valAddr)
	if err != nil {
		log.Errorln(err)
		return err
	}

	defer func() {
		for _, a := range amount {
			if a.Amount.IsInt64() {
				telemetry.SetGaugeWithLabels(
					[]string{"tx", "msg", "withdraw_reward"},
					float32(a.Amount.Int64()),
					[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
				)
			}
		}
	}()

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, delAddrStr),
		),
	)

	return nil
}

func (k Keeper) withdrawAllDelegatorRewards(ctx sdk.Context, delAddr eth.Addr) error {
	var err error
	k.stakingKeeper.IterateDelegations(
		ctx, delAddr,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			valAddr := del.GetValidatorAddr()
			withdrawErr := k.withdrawDelegatorRewardForOneValidator(ctx, delAddr.String(), valAddr.String())
			if withdrawErr != nil {
				err = withdrawErr
				return true
			}
			return false
		},
	)
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

func (k Keeper) withdrawRewardsAndCommission(ctx sdk.Context, addr eth.Addr) error {
	// 1. Withdraw reward for all validators
	err := k.withdrawAllDelegatorRewards(ctx, addr)
	if err != nil {
		return err
	}

	// 2. If addr is a validator address, withdraw its commission
	accumCommission := k.GetValidatorAccumulatedCommission(ctx, addr)
	if !accumCommission.Commission.IsZero() {
		_, err = k.WithdrawValidatorCommission(ctx, addr)
		if err != nil {
			return err
		}
	}
	return nil
}
