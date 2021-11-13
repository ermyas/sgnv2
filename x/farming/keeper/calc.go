package keeper

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CalculateAmountEarnedBetween is used for calculating how many reward tokens haven been earned from
// startBlockHeight to endBlockHeight, and return the amount.
func (k Keeper) CalculateAmountEarnedBetween(ctx sdk.Context, pool types.FarmingPool) (types.FarmingPool, sdk.DecCoins) {
	currentPeriod := k.GetPoolCurrentRewards(ctx, pool.Name)
	endBlockHeight := ctx.BlockHeight()

	totalEarnedTokens := sdk.DecCoins{}
	for i := 0; i < len(pool.RewardTokenInfos); i++ {
		rewardStartBlockHeight := pool.RewardTokenInfos[i].RewardStartBlockHeight
		var startBlockHeight int64
		if currentPeriod.StartBlockHeight <= rewardStartBlockHeight {
			startBlockHeight = rewardStartBlockHeight
		} else {
			startBlockHeight = currentPeriod.StartBlockHeight
		}

		// no tokens earned
		if startBlockHeight == 0 || startBlockHeight >= endBlockHeight {
			continue
		}

		var earnedTokens sdk.DecCoins
		// calculate how many tokens to be earned between startBlockHeight and endBlockHeight
		blockInterval := sdk.NewDec(endBlockHeight - startBlockHeight)
		amount := blockInterval.MulTruncate(pool.RewardTokenInfos[i].RewardAmountPerBlock)
		remaining := pool.RewardTokenInfos[i].RemainingAmount
		if amount.LT(remaining.Amount) {
			pool.RewardTokenInfos[i].RemainingAmount.Amount = remaining.Amount.Sub(amount)
			earnedTokens = sdk.DecCoins{sdk.NewDecCoinFromDec(remaining.Denom, amount)}
		} else {
			pool.RewardTokenInfos[i] = types.NewRewardTokenInfo(
				sdk.NewDecCoin(remaining.Denom, sdk.ZeroInt()), 0, sdk.ZeroDec(),
			)
			earnedTokens = sdk.DecCoins{sdk.NewDecCoinFromDec(remaining.Denom, remaining.Amount)}
		}
		pool.TotalAccumulatedRewards = pool.TotalAccumulatedRewards.Add(earnedTokens...)
		totalEarnedTokens = totalEarnedTokens.Add(earnedTokens...)
	}
	return pool, totalEarnedTokens
}

func (k Keeper) WithdrawRewards(
	ctx sdk.Context, poolName string, totalStakedAmount sdk.DecCoin, earnedTokens sdk.DecCoins, addr eth.Addr,
) (sdk.DecCoins, error) {
	// 0. Check existence of stake info
	stakeInfo, found := k.GetStakeInfo(ctx, addr, poolName)
	if !found {
		return nil, types.WrapErrNoStakeInfoFound(addr.String(), poolName)
	}

	// 1. End current period and calculate rewards
	endingPeriod := k.IncrementPoolPeriod(ctx, poolName, totalStakedAmount, earnedTokens)
	rewards := k.calculateRewards(ctx, poolName, addr, endingPeriod, stakeInfo)

	// 2. Transfer rewards to user account
	// truncate coins
	truncatedRewards, _ := rewards.TruncateDecimal()
	derivedAddr := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, addr)
	if !truncatedRewards.IsZero() {
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.RewardModuleAccountName, derivedAddr, truncatedRewards)
		if err != nil {
			return nil, err
		}
	}

	// 3. Decrement reference count of stake info
	k.decrementReferenceCount(ctx, poolName, stakeInfo.ReferencePeriod)

	return sdk.NewDecCoinsFromCoins(truncatedRewards...), nil
}

// IncrementPoolPeriod increments pool period, returning the period just ended
func (k Keeper) IncrementPoolPeriod(
	ctx sdk.Context, poolName string, totalStakedAmount sdk.DecCoin, earnedTokens sdk.DecCoins,
) uint64 {
	// 1. Fetch current period rewards
	rewards := k.GetPoolCurrentRewards(ctx, poolName)
	// 2. Calculate current reward ratio
	rewards.Rewards = rewards.Rewards.Add(earnedTokens...)
	var currentRatio sdk.DecCoins
	if totalStakedAmount.IsZero() {
		currentRatio = sdk.DecCoins{}
	} else {
		// NOTE: first multiply by StakingScaleFactor (1e12) to support reward tokens with much smaller decimals than
		// the stake token. Necessary to truncate so we don't allow withdrawing more rewards than owed.
		currentRatio = rewards.Rewards.MulDec(common.StakingScaleFactor).QuoDecTruncate(totalStakedAmount.Amount)
	}

	// 3.1. Get the previous pool historical rewards
	historical := k.GetPoolHistoricalRewards(ctx, poolName, rewards.Period-1).CumulativeRewardRatio
	// 3.2. Decrement reference count
	k.decrementReferenceCount(ctx, poolName, rewards.Period-1)
	// 3.3. Create new pool historical rewards with reference count of 1, then set it into store
	newHistoricalRewards := types.NewPoolHistoricalRewards(historical.Add(currentRatio...), 1)
	k.SetPoolHistoricalRewards(ctx, poolName, rewards.Period, newHistoricalRewards)

	// 4. Set new current rewards into store, incrementing period by 1
	newCurRewards := types.NewPoolCurrentRewards(ctx.BlockHeight(), rewards.Period+1, sdk.DecCoins{})
	k.SetPoolCurrentRewards(ctx, poolName, newCurRewards)

	return rewards.Period
}

// incrementReferenceCount increments the reference count for a historical rewards value
func (k Keeper) incrementReferenceCount(ctx sdk.Context, poolName string, period uint64) {
	historical := k.GetPoolHistoricalRewards(ctx, poolName, period)
	if historical.ReferenceCount > 1 {
		panic("reference count should never exceed 1")
	}
	historical.ReferenceCount++
	k.SetPoolHistoricalRewards(ctx, poolName, period, historical)
}

// decrementReferenceCount decrements the reference count for a historical rewards value,
// and delete if zero references remain.
func (k Keeper) decrementReferenceCount(ctx sdk.Context, poolName string, period uint64) {
	historical := k.GetPoolHistoricalRewards(ctx, poolName, period)
	if historical.ReferenceCount == 0 {
		panic("cannot set negative reference count")
	}
	historical.ReferenceCount--
	if historical.ReferenceCount == 0 {
		k.DeletePoolHistoricalReward(ctx, poolName, period)
	} else {
		k.SetPoolHistoricalRewards(ctx, poolName, period, historical)
	}
}

func (k Keeper) calculateRewards(
	ctx sdk.Context, poolName string, addr eth.Addr, endingPeriod uint64, stakeInfo types.StakeInfo,
) (rewards sdk.DecCoins) {
	if stakeInfo.StartBlockHeight == ctx.BlockHeight() {
		// started this height, no rewards yet
		return
	}

	startingPeriod := stakeInfo.ReferencePeriod
	// calculate rewards for final period
	return k.calculateStakeRewardsBetween(ctx, poolName, startingPeriod, endingPeriod, stakeInfo.Amount)
}

// calculateStakeRewardsBetween calculate the rewards accrued by a pool between two periods
func (k Keeper) calculateStakeRewardsBetween(ctx sdk.Context, poolName string, startingPeriod, endingPeriod uint64,
	amount sdk.DecCoin) (rewards sdk.DecCoins) {

	// sanity check
	if startingPeriod > endingPeriod {
		panic("startingPeriod cannot be greater than endingPeriod")
	}

	if amount.Amount.LT(sdk.ZeroDec()) {
		panic("amount should not be negative")
	}

	// return amount * (ending - starting) / StakingScaleFactor
	starting := k.GetPoolHistoricalRewards(ctx, poolName, startingPeriod)
	ending := k.GetPoolHistoricalRewards(ctx, poolName, endingPeriod)
	difference := ending.CumulativeRewardRatio.Sub(starting.CumulativeRewardRatio)
	// NOTE: necessary to truncate so we don't allow withdrawing more rewards than owed
	rewards = difference.MulDec(amount.Amount).QuoDecTruncate(common.StakingScaleFactor)
	return
}

// UpdateStakeInfo updates stake info for the modified stake info
func (k Keeper) UpdateStakeInfo(ctx sdk.Context, addr eth.Addr, poolName string, changedAmount sdk.Dec) {
	// period has already been incremented - we want to store the period ended by this stake action
	previousPeriod := k.GetPoolCurrentRewards(ctx, poolName).Period - 1

	// get stake info, then set it into store
	stakeInfo, found := k.GetStakeInfo(ctx, addr, poolName)
	if !found {
		panic("the stake info can't be found")
	}
	stakeInfo.StartBlockHeight = ctx.BlockHeight()
	stakeInfo.ReferencePeriod = previousPeriod
	stakeInfo.Amount.Amount = stakeInfo.Amount.Amount.Add(changedAmount)
	if stakeInfo.Amount.IsZero() {
		k.DeleteStakeInfo(ctx, eth.Hex2Addr(stakeInfo.StakerAddress), stakeInfo.PoolName)
		k.DeleteAddressInFarmingPool(ctx, stakeInfo.PoolName, eth.Hex2Addr(stakeInfo.StakerAddress))
	} else {
		// increment reference count for the period we're going to track
		k.incrementReferenceCount(ctx, poolName, previousPeriod)

		// set the updated stake info
		k.SetStakeInfo(ctx, stakeInfo)
		k.SetAddressInFarmingPool(ctx, stakeInfo.PoolName, eth.Hex2Addr(stakeInfo.StakerAddress))
	}
}
