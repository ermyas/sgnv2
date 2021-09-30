package keeper

import (
	"fmt"

	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initialize default parameters and the keeper's address to pubkey map
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	k.SetParams(ctx, data.Params)

	var rewardModuleAccHoldings sdk.DecCoins
	var moduleAccHoldings sdk.DecCoins

	for _, pool := range data.Pools {
		// Create stake token if not existent
		stakeToken := pool.StakeToken
		found := k.HasERC20Token(ctx, stakeToken.ChainId, stakeToken.Symbol)
		if !found {
			k.SetERC20Token(ctx, stakeToken)
		}
		// Create reward tokens if not existent
		for _, rewardToken := range pool.RewardTokens {
			found = k.HasERC20Token(ctx, rewardToken.ChainId, rewardToken.Symbol)
			if !found {
				k.SetERC20Token(ctx, rewardToken)
			}
		}

		moduleAccHoldings = moduleAccHoldings.Add(sdk.DecCoins{pool.TotalStakedAmount}...)
		rewardModuleAccHoldings = rewardModuleAccHoldings.Add(pool.TotalAccumulatedRewards...)
		k.SetFarmingPool(ctx, pool)
	}
	for _, stakeInfo := range data.StakeInfos {
		k.SetStakeInfo(ctx, stakeInfo)
	}
	for _, historical := range data.PoolHistoricalRewards {
		k.SetPoolHistoricalRewards(ctx, historical.PoolName, historical.Period, historical.Rewards)
	}
	for _, current := range data.PoolCurrentRewards {
		k.SetPoolCurrentRewards(ctx, current.PoolName, current.Rewards)
	}

	// init module account
	moduleAcc := k.authKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
	moduleAccBalance := k.bankKeeper.GetAllBalances(ctx, moduleAcc.GetAddress())
	if moduleAccBalance.IsZero() {
		truncatedModuleAccHoldings, _ := moduleAccHoldings.TruncateDecimal()
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, truncatedModuleAccHoldings); err != nil {
			panic(err)
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, moduleAcc.GetAddress(), truncatedModuleAccHoldings); err != nil {
			panic(err)
		}
		k.authKeeper.SetModuleAccount(ctx, moduleAcc)
	}

	rewardModuleAcc := k.authKeeper.GetModuleAccount(ctx, types.RewardModuleAccountName)
	if rewardModuleAcc == nil {
		panic(fmt.Sprintf("%s reward module account has not been set", types.RewardModuleAccountName))
	}
	rewardModuleAccBalance := k.bankKeeper.GetAllBalances(ctx, rewardModuleAcc.GetAddress())
	if rewardModuleAccBalance.IsZero() {
		truncatedRewardModuleAccHoldings, _ := rewardModuleAccHoldings.TruncateDecimal()
		if err := k.bankKeeper.MintCoins(ctx, types.RewardModuleAccountName, truncatedRewardModuleAccHoldings); err != nil {
			panic(err)
		}
		if err :=
			k.bankKeeper.SendCoinsFromModuleToAccount(
				ctx,
				types.ModuleName,
				rewardModuleAcc.GetAddress(),
				truncatedRewardModuleAccHoldings,
			); err != nil {
			panic(err)
		}
		k.authKeeper.SetModuleAccount(ctx, rewardModuleAcc)
	}
}

// ExportGenesis writes the current store values to a genesis file, which can be imported again with InitGenesis
func (k Keeper) ExportGenesis(ctx sdk.Context) (data *types.GenesisState) {
	params := k.GetParams(ctx)

	pools := k.GetFarmingPools(ctx)

	stakeInfos := make([]types.StakeInfo, 0)
	k.IterateAllStakeInfos(ctx,
		func(stakeInfo types.StakeInfo) (stop bool) {
			stakeInfos = append(stakeInfos, stakeInfo)
			return false
		},
	)

	allHistoricalRewards := make([]types.PoolHistoricalRewardsRecord, 0)
	k.IterateAllPoolHistoricalRewards(ctx,
		func(poolName string, period uint64, rewards types.PoolHistoricalRewards) (stop bool) {
			allHistoricalRewards = append(allHistoricalRewards, types.PoolHistoricalRewardsRecord{
				PoolName: poolName,
				Period:   period,
				Rewards:  rewards,
			})
			return false
		},
	)

	allCurRewards := make([]types.PoolCurrentRewardsRecord, 0)
	k.IterateAllPoolCurrentRewards(ctx,
		func(poolName string, rewards types.PoolCurrentRewards) (stop bool) {
			allCurRewards = append(allCurRewards, types.PoolCurrentRewardsRecord{
				PoolName: poolName,
				Rewards:  rewards,
			})
			return false
		},
	)

	return types.NewGenesisState(params, pools, stakeInfos, allHistoricalRewards, allCurRewards)
}
