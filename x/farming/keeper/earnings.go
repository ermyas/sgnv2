package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetEarnings gets the earnings info by a given user address and a specific pool name
func (k Keeper) GetEarnings(ctx sdk.Context, poolName string, addr eth.Addr) (types.Earnings, error) {
	var earnings types.Earnings
	stakeInfo, found := k.GetStakeInfo(ctx, addr, poolName)
	if !found {
		return earnings, types.WrapErrNoStakeInfoFound(addr.String(), poolName)
	}

	pool, found := k.GetFarmingPool(ctx, poolName)
	if !found {
		return earnings, types.WrapErrPoolNotExist(poolName)
	}

	// 1.1 Calculate how many reward tokens have been earned between start block height and current height
	updatedPool, earnedTokens := k.CalculateAmountEarnedBetween(ctx, pool)

	endingPeriod := k.IncrementPoolPeriod(ctx, poolName, updatedPool.TotalStakedAmount, earnedTokens)
	rewards := k.calculateRewards(ctx, poolName, addr, endingPeriod, stakeInfo)

	earnings = types.NewEarnings(ctx.BlockHeight(), stakeInfo.Amount, rewards)
	return earnings, nil
}
