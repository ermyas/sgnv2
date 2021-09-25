package keeper

import (
	"fmt"

	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants registers all farm invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "reward-module-account", RewardModuleAccountInvariant(k))
	ir.RegisterRoute(types.ModuleName, "module-account", ModuleAccountInvariant(k))
}

// AllInvariants runs all invariants of the distribution module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := RewardModuleAccountInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return ModuleAccountInvariant(k)(ctx)
	}
}

// ModuleAccountInvariant checks if the farming ModuleAccount balance is consistent with the sum of staked amount.
func ModuleAccountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// iterate all stake infos
		totalStakedAmount := sdk.DecCoins{}
		k.IterateAllStakeInfos(ctx, func(stakeInfo types.StakeInfo) (stop bool) {
			totalStakedAmount = totalStakedAmount.Add(sdk.NewDecCoins(stakeInfo.Amount)...)
			return false
		})

		// get farming module account
		moduleAccount := k.authKeeper.GetModuleAccount(ctx, types.ModuleName)
		moduleAccountAmount := sdk.NewDecCoinsFromCoins(k.bankKeeper.GetAllBalances(ctx, moduleAccount.GetAddress())...)

		// make a comparison
		broken := !(moduleAccountAmount.IsEqual(totalStakedAmount))

		return sdk.FormatInvariant(types.ModuleName, "ModuleAccount coins",
			fmt.Sprintf("\texpected farming ModuleAccount coins: %s\n"+
				"\tactual farming ModuleAccount coins: %s\n",
				totalStakedAmount, moduleAccountAmount)), broken
	}
}

// RewardModuleAccountInvariant checks if the reward-module-account balance is consistent
// with the total accumulated rewards.
func RewardModuleAccountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// iterate all pools, then calculate the total reward amount
		expectedRewardModuleAccAmount := sdk.DecCoins{}
		pools := k.GetFarmingPools(ctx)
		for _, pool := range pools {
			expectedRewardModuleAccAmount = expectedRewardModuleAccAmount.Add(pool.TotalAccumulatedRewards...)
			for _, rewardTokenInfo := range pool.RewardTokenInfos {
				expectedRewardModuleAccAmount = expectedRewardModuleAccAmount.Add(sdk.DecCoins{rewardTokenInfo.RemainingAmount}...)
			}
		}

		// get the reward_module_account
		rewardModuleAccount := k.authKeeper.GetModuleAccount(ctx, types.RewardModuleAccountName)
		actualRewardModuleAccountAmount :=
			sdk.NewDecCoinsFromCoins(
				k.bankKeeper.GetAllBalances(ctx, rewardModuleAccount.GetAddress())...)

		// make a comparison
		broken := !(expectedRewardModuleAccAmount.IsEqual(actualRewardModuleAccountAmount))

		return sdk.FormatInvariant(types.ModuleName, "reward_module_account coins",
			fmt.Sprintf("\texpected reward_module_account coins: %s\n"+
				"\tactual reward_module_account coins: %s\n",
				expectedRewardModuleAccAmount, actualRewardModuleAccountAmount)), broken
	}
}
