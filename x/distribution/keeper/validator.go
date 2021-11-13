package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val stakingtypes.ValidatorI) {
	// set initial historical rewards (period 0) with reference count of 1
	k.SetValidatorHistoricalRewards(ctx, val.GetEthAddr(), 0, types.NewValidatorHistoricalRewards(sdk.DecCoins{}, 1))

	// set current rewards (starting at period 1)
	k.SetValidatorCurrentRewards(ctx, val.GetEthAddr(), types.NewValidatorCurrentRewards(sdk.DecCoins{}, 1))

	// set accumulated commission
	k.SetValidatorAccumulatedCommission(ctx, val.GetEthAddr(), types.InitialValidatorAccumulatedCommission())

	// set outstanding rewards
	k.SetValidatorOutstandingRewards(ctx, val.GetEthAddr(), types.ValidatorOutstandingRewards{Rewards: sdk.DecCoins{}})
}

// increment validator period, returning the period just ended
func (k Keeper) IncrementValidatorPeriod(ctx sdk.Context, val stakingtypes.ValidatorI) uint64 {
	// fetch current rewards
	rewards := k.GetValidatorCurrentRewards(ctx, val.GetEthAddr())

	// calculate current ratio
	var current sdk.DecCoins
	if val.GetTokens().IsZero() {

		// can't calculate ratio for zero-token validators
		// ergo we instead add to the community pool
		feePool := k.GetFeePool(ctx)
		outstanding := k.GetValidatorOutstandingRewards(ctx, val.GetEthAddr())
		feePool.CommunityPool = feePool.CommunityPool.Add(rewards.Rewards...)
		outstanding.Rewards = outstanding.GetRewards().Sub(rewards.Rewards)
		k.SetFeePool(ctx, feePool)
		k.SetValidatorOutstandingRewards(ctx, val.GetEthAddr(), outstanding)

		current = sdk.DecCoins{}
	} else {
		// NOTE: first multiply by StakingScaleFactor (1e12) to support reward tokens with much smaller decimals than
		// the stake token. Necessary to truncate so we don't allow withdrawing more rewards than owed.
		current = rewards.Rewards.MulDec(common.StakingScaleFactor).QuoDecTruncate(val.GetTokens().ToDec())
	}

	// fetch historical rewards for last period
	historical := k.GetValidatorHistoricalRewards(ctx, val.GetEthAddr(), rewards.Period-1).CumulativeRewardRatio

	// decrement reference count
	k.decrementReferenceCount(ctx, val.GetEthAddr(), rewards.Period-1)

	// set new historical rewards with reference count of 1
	k.SetValidatorHistoricalRewards(ctx, val.GetEthAddr(), rewards.Period, types.NewValidatorHistoricalRewards(historical.Add(current...), 1))

	// set current rewards, incrementing period by 1
	k.SetValidatorCurrentRewards(ctx, val.GetEthAddr(), types.NewValidatorCurrentRewards(sdk.DecCoins{}, rewards.Period+1))

	return rewards.Period
}

// increment the reference count for a historical rewards value
func (k Keeper) incrementReferenceCount(ctx sdk.Context, valAddr eth.Addr, period uint64) {
	historical := k.GetValidatorHistoricalRewards(ctx, valAddr, period)
	if historical.ReferenceCount > 1 {
		panic("reference count should never exceed 1")
	}
	historical.ReferenceCount++
	k.SetValidatorHistoricalRewards(ctx, valAddr, period, historical)
}

// decrement the reference count for a historical rewards value, and delete if zero references remain
func (k Keeper) decrementReferenceCount(ctx sdk.Context, valAddr eth.Addr, period uint64) {
	historical := k.GetValidatorHistoricalRewards(ctx, valAddr, period)
	if historical.ReferenceCount == 0 {
		panic("cannot set negative reference count")
	}
	historical.ReferenceCount--
	if historical.ReferenceCount == 0 {
		k.DeleteValidatorHistoricalReward(ctx, valAddr, period)
	} else {
		k.SetValidatorHistoricalRewards(ctx, valAddr, period, historical)
	}
}
