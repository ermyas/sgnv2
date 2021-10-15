package keeper

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Stake(
	ctx sdk.Context,
	poolName string,
	address eth.Addr,
	amount sdk.Coin,
	mintStakes bool,
) error {
	// 1.1 Get farming pool
	pool, found := k.GetFarmingPool(ctx, poolName)
	if !found {
		return types.WrapErrPoolNotExist(poolName)
	}
	// 1.2. Check stake token denom
	stakeTokenDenom := DeriveERC20TokenDenom(pool.StakeToken.ChainId, pool.StakeToken.Symbol)
	if stakeTokenDenom != amount.Denom {
		return types.WrapErrInvalidDenom(stakeTokenDenom, amount.Denom)
	}

	hasStake := k.HasStakeInfo(ctx, address, poolName)

	// 2. Calculate how many reward tokens are earned in the current period
	updatedPool, earnedTokens := k.CalculateAmountEarnedBetween(ctx, pool)

	// 3. Check stake info
	var rewards sdk.DecCoins
	if hasStake {
		// If it exists, withdraw money
		var err error
		rewards, err = k.WithdrawRewards(ctx, pool.Name, pool.TotalStakedAmount, earnedTokens, address)
		if err != nil {
			return err
		}
		// NOTE: Will panic if TotalAccumulatedRewards < rewards
		updatedPool.TotalAccumulatedRewards = updatedPool.TotalAccumulatedRewards.Sub(rewards)

	} else {
		// If it doesn't exist, only increase period
		k.IncrementPoolPeriod(ctx, pool.Name, pool.TotalStakedAmount, earnedTokens)

		// Create new stake info
		stakeInfo := types.NewStakeInfo(
			address, pool.Name, sdk.NewDecCoinFromDec(stakeTokenDenom, sdk.ZeroDec()),
			ctx.BlockHeight(), 0,
		)
		k.SetStakeInfo(ctx, stakeInfo)
		k.SetAddressInFarmingPool(ctx, poolName, address)
	}

	// 4. Update stake info
	k.UpdateStakeInfo(ctx, address, poolName, sdk.NewDecFromInt(amount.Amount))

	// 5. Mint new stakes or send the staked tokens from its own account to farming module account
	if mintStakes {
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amount)); err != nil {
			return types.WrapErrMintCoinsFailed(err.Error())
		}
	} else {
		derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, address)
		if err := k.bankKeeper.SendCoinsFromAccountToModule(
			ctx, derivedAccAddress, types.ModuleName, sdk.NewCoins(amount),
		); err != nil {
			return types.WrapErrSendCoinsFromAccountToModuleFailed(err.Error())
		}
	}

	// 6. Update farming pool
	updatedPool.TotalStakedAmount = updatedPool.TotalStakedAmount.Add(sdk.NewDecCoinFromCoin(amount))
	k.SetFarmingPool(ctx, updatedPool)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeStake,
		sdk.NewAttribute(types.AttributeKeyAddress, address.String()),
		sdk.NewAttribute(types.AttributeKeyPool, poolName),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	))
	return nil
}

func (k Keeper) Unstake(
	ctx sdk.Context,
	poolName string,
	address eth.Addr,
	amount sdk.Coin,
	burnStakes bool,
) error {
	// 1.1 Check if there are enough tokens to unstake
	stakeInfo, found := k.GetStakeInfo(ctx, address, poolName)
	if !found {
		return types.WrapErrNoStakeInfoFound(address.String(), poolName)
	}

	if stakeInfo.Amount.Denom != amount.Denom {
		return types.WrapErrInvalidDenom(stakeInfo.Amount.Denom, amount.Denom)
	}

	if stakeInfo.Amount.IsLT(sdk.NewDecCoinFromCoin(amount)) {
		return types.WrapErrInsufficientAmount(stakeInfo.Amount.String(), amount.String())
	}

	// 1.2. Get the pool info
	pool, poolFound := k.GetFarmingPool(ctx, poolName)
	if !poolFound {
		return types.WrapErrPoolNotExist(poolName)
	}
	stakeTokenDenom := DeriveERC20TokenDenom(pool.StakeToken.ChainId, pool.StakeToken.Symbol)
	if stakeTokenDenom != amount.Denom {
		return types.WrapErrInvalidDenom(stakeTokenDenom, amount.Denom)
	}

	// 2. Calculate how many reward tokens are earned in the current period
	updatedPool, earnedTokens := k.CalculateAmountEarnedBetween(ctx, pool)

	// 3. Withdraw rewards
	rewards, err := k.WithdrawRewards(ctx, pool.Name, pool.TotalStakedAmount, earnedTokens, address)
	if err != nil {
		return err
	}

	// 4. Update the stake info
	k.UpdateStakeInfo(ctx, address, poolName, sdk.NewDecFromInt(amount.Amount.Neg()))

	// 5. Burn stakes or send the staked tokens from farming module account to its own account
	if burnStakes {
		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amount)); err != nil {
			return types.WrapErrBurnCoinsFailed(err.Error())
		}
	} else {
		derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, address)
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, derivedAccAddress, sdk.NewCoins(amount)); err != nil {
			return types.WrapErrSendCoinsFromModuleToAccountFailed(err.Error())
		}
	}

	// 6. Update farming pool
	updatedPool.TotalStakedAmount = updatedPool.TotalStakedAmount.Sub(sdk.NewDecCoinFromCoin(amount))
	// NOTE: Will panic if TotalAccumulatedRewards < rewards
	updatedPool.TotalAccumulatedRewards = updatedPool.TotalAccumulatedRewards.Sub(rewards)
	k.SetFarmingPool(ctx, updatedPool)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnstake,
		sdk.NewAttribute(types.AttributeKeyAddress, address.String()),
		sdk.NewAttribute(types.AttributeKeyPool, poolName),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	))
	return nil
}
