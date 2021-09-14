package keeper

import (
	"fmt"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// register all distribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "nonnegative-outstanding",
		NonNegativeOutstandingInvariant(k))
	// ir.RegisterRoute(types.ModuleName, "can-withdraw",
	// 	CanWithdrawInvariant(k))
	// ir.RegisterRoute(types.ModuleName, "reference-count",
	// 	ReferenceCountInvariant(k))
	ir.RegisterRoute(types.ModuleName, "module-account",
		ModuleAccountInvariant(k))
}

// AllInvariants runs all invariants of the distribution module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// res, stop := CanWithdrawInvariant(k)(ctx)
		// if stop {
		// 	return res, stop
		// }
		res, stop := NonNegativeOutstandingInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		// res, stop = ReferenceCountInvariant(k)(ctx)
		// if stop {
		// 	return res, stop
		// }
		return ModuleAccountInvariant(k)(ctx)
	}
}

// NonNegativeOutstandingInvariant checks that outstanding unwithdrawn fees are never negative
func NonNegativeOutstandingInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var msg string
		var count int
		var outstanding sdk.DecCoins

		k.IterateValidatorOutstandingRewards(ctx, func(addr eth.Addr, rewards types.ValidatorOutstandingRewards) (stop bool) {
			outstanding = rewards.GetRewards()
			if outstanding.IsAnyNegative() {
				count++
				msg += fmt.Sprintf("\t%v has negative outstanding coins: %v\n", addr, outstanding)
			}
			return false
		})
		broken := count != 0

		return sdk.FormatInvariant(types.ModuleName, "nonnegative outstanding",
			fmt.Sprintf("found %d validators with negative outstanding rewards\n%s", count, msg)), broken
	}
}

// TODO: CanWithdrawInvariant checks that current rewards can be completely withdrawn

// TODO: ReferenceCountInvariant checks that the number of historical rewards records is correct

// ModuleAccountInvariant checks that the coins held by the distr ModuleAccount
// is consistent with the sum of validator outstanding rewards
func ModuleAccountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {

		var expectedCoins sdk.DecCoins
		k.IterateValidatorOutstandingRewards(ctx, func(_ eth.Addr, rewards types.ValidatorOutstandingRewards) (stop bool) {
			expectedCoins = expectedCoins.Add(rewards.Rewards...)
			return false
		})

		communityPool := k.GetFeePoolCommunityCoins(ctx)
		expectedInt, _ := expectedCoins.Add(communityPool...).TruncateDecimal()

		macc := k.GetDistributionAccount(ctx)
		balances := k.bankKeeper.GetAllBalances(ctx, macc.GetAddress())

		broken := !balances.IsEqual(expectedInt)
		return sdk.FormatInvariant(
			types.ModuleName, "ModuleAccount coins",
			fmt.Sprintf("\texpected ModuleAccount coins:     %s\n"+
				"\tdistribution ModuleAccount coins: %s\n",
				expectedInt, balances,
			),
		), broken
	}
}
