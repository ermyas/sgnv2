package keeper

import (
	"fmt"

	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandleAddPoolProposal is a handler for executing a passed AddPoolProposal
func HandleAddPoolProposal(ctx sdk.Context, k Keeper, p *types.AddPoolProposal) error {
	if err := k.CheckAddPoolProposal(ctx, p); err != nil {
		return err
	}
	// 1. Create stake token if not existent
	stakeToken := p.StakeToken
	found := k.HasERC20Token(ctx, stakeToken.ChainId, stakeToken.Symbol)
	if !found {
		k.SetERC20Token(ctx, stakeToken)
	}
	// 2. Create reward tokens if not existent
	for _, rewardToken := range p.RewardTokens {
		found = k.HasERC20Token(ctx, rewardToken.ChainId, rewardToken.Symbol)
		if !found {
			k.SetERC20Token(ctx, rewardToken)
		}
	}
	// 3.1. Add pool
	stakeTokenDenom := DeriveERC20TokenDenom(stakeToken.ChainId, stakeToken.Symbol)
	totalStakedAmount :=
		sdk.NewDecCoin(
			stakeTokenDenom,
			sdk.ZeroInt(),
		)
	var rewardTokenInfos types.RewardTokenInfos
	var totalAccumulatedRewards sdk.DecCoins
	for _, initialRewardInput := range p.InitialRewardInputs {
		truncatedAddAmount, _ := initialRewardInput.AddAmount.TruncateDecimal()
		// Mint reward
		k.bankKeeper.MintCoins(ctx, types.RewardModuleAccountName, sdk.NewCoins(truncatedAddAmount))
		rewardTokenInfo := types.RewardTokenInfo{
			RemainingAmount:        sdk.NewDecCoinFromCoin(truncatedAddAmount),
			RewardStartBlockHeight: ctx.BlockHeight() + initialRewardInput.RewardStartBlockDelay,
			RewardAmountPerBlock:   initialRewardInput.NewRewardAmountPerBlock,
		}
		rewardTokenInfos = append(rewardTokenInfos, rewardTokenInfo)
		totalAccumulatedRewards =
			append(totalAccumulatedRewards, sdk.NewDecCoin(truncatedAddAmount.Denom, sdk.ZeroInt()))
	}
	pool :=
		types.NewFarmingPool(
			p.PoolName,
			p.StakeToken,
			p.RewardTokens,
			totalStakedAmount,
			rewardTokenInfos,
			totalAccumulatedRewards,
		)
	k.SetFarmingPool(ctx, pool)

	// 3.2. Set initial pool period
	poolHistoricalRewards := types.NewPoolHistoricalRewards(sdk.DecCoins{}, 1)
	k.SetPoolHistoricalRewards(ctx, p.PoolName, 0, poolHistoricalRewards)
	poolCurrentRewards := types.NewPoolCurrentRewards(ctx.BlockHeight(), 1, sdk.DecCoins{})
	k.SetPoolCurrentRewards(ctx, p.PoolName, poolCurrentRewards)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddPool,
		sdk.NewAttribute(types.AttributeKeyPool, p.PoolName),
		sdk.NewAttribute(types.AttributeKeyStakeToken, stakeTokenDenom),
	))
	return nil
}

// CheckAddPoolProposal checks the validity of an AddPoolProposal
func (k Keeper) CheckAddPoolProposal(ctx sdk.Context, p *types.AddPoolProposal) error {
	// 1. Check pool existence
	_, found := k.GetFarmingPool(ctx, p.PoolName)
	if found {
		return types.WrapErrPoolAlreadyExist(p.PoolName)
	}
	// 2.1. Check reward tokens and initial reward inputs lengths match
	if len(p.RewardTokens) != len(p.InitialRewardInputs) {
		return types.WrapErrInvalidInput(
			fmt.Sprintf(
				"reward token info array lengths mismatch: %d vs %d",
				len(p.RewardTokens),
				len(p.InitialRewardInputs),
			))
	}
	// 2.2. Check reward token denoms match
	for i, rewardToken := range p.RewardTokens {
		denom1 := DeriveERC20TokenDenom(rewardToken.ChainId, rewardToken.Symbol)
		denom2 := p.InitialRewardInputs[i].AddAmount.Denom
		if denom1 != denom2 {
			return types.WrapErrInvalidInput(
				fmt.Sprintf("reward token denoms mismatch: %s vs %s", denom1, denom2))
		}
	}
	return nil
}
