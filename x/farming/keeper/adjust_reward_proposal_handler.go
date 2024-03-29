package keeper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandleAdjustRewardProposal is a handler for executing a passed AdjustRewardProposal
func HandleAdjustRewardProposal(ctx sdk.Context, k Keeper, p *types.AdjustRewardProposal) error {
	adjustRewardInfo := p.GetAdjustRewardInfo()
	return handleAdjustRewardProposalByAdjustRewardInfo(ctx, k, &adjustRewardInfo)
}

func handleAdjustRewardProposalByAdjustRewardInfo(ctx sdk.Context, k Keeper, p *types.AdjustRewardInfo) error {
	pool, tokensFromInputs, err := k.CheckAdjustRewardProposal(ctx, p)
	if err != nil {
		return err
	}

	// 1 Calculate how many reward tokens have been earned between start block height and current height
	updatedPool, earnedTokens := k.CalculateAmountEarnedBetween(ctx, *pool)

	// 2. Increment pool period
	k.IncrementPoolPeriod(ctx, pool.Name, pool.TotalStakedAmount, earnedTokens)

	// 3. Update existing infos with new inputs, assuming both are sorted by ascending denoms
	// Mint AddAmount, update RemainingAmount, set RewardAmountPerBlock. For new rewards, also set RewardStartBlockHeight

	// Remove possible duplicates due to a historical bug, only keeping the first one.
	if p.RemoveDuplicates {
		denoms := make(map[string]bool)
		fixedTokens := []commontypes.ERC20Token{}
		fixedInfos := []types.RewardTokenInfo{}
		for i, info := range pool.RewardTokenInfos {
			denom := info.RemainingAmount.Denom
			if _, exists := denoms[denom]; !exists {
				denoms[denom] = true
				fixedTokens = append(fixedTokens, pool.RewardTokens[i])
				fixedInfos = append(fixedInfos, info)
			}
		}
		pool.RewardTokens = fixedTokens
		pool.RewardTokenInfos = fixedInfos
	}

	newInfos := []types.RewardTokenInfo{}
	newTokens := []commontypes.ERC20Token{}
	newAccumulatedRewards := []sdk.DecCoin{}
	indexInfos, indexInputs := 0, 0
	lenInfos, lenInputs := len(pool.RewardTokenInfos), len(p.RewardAdjustmentInputs)
	for {
		if indexInfos == lenInfos {
			if indexInputs == lenInputs {
				// break if both arrays are empty
				break
			}

			// append the rest of the inputs if infos are exhausted
			for i := indexInputs; i < lenInputs; i++ {
				input := p.RewardAdjustmentInputs[i]
				newInfo, err := k.processNewInput(ctx, pool.Name, &input)
				if err != nil {
					return err
				}
				newInfos = append(newInfos, *newInfo)
				newTokens = append(newTokens, tokensFromInputs[i])
				newAccumulatedRewards =
					append(newAccumulatedRewards, sdk.NewDecCoin(input.AddAmount.Denom, sdk.ZeroInt()))
			}
			break
		} else if indexInputs == lenInputs {
			// append the rest of the infos if inputs are exhausted
			for i := indexInfos; i < lenInfos; i++ {
				info := pool.RewardTokenInfos[i]
				newInfos = append(newInfos, info)
				newTokens = append(newTokens, pool.RewardTokens[i])
				newAccumulatedRewards =
					append(newAccumulatedRewards, updatedPool.TotalAccumulatedRewards[i])
			}
			break
		}

		info := pool.RewardTokenInfos[indexInfos]
		input := p.RewardAdjustmentInputs[indexInputs]
		infoAmount, inputAmount := info.RemainingAmount, input.AddAmount

		switch strings.Compare(infoAmount.Denom, inputAmount.Denom) {
		case -1: // info denom < input denom, add existing info
			newInfos = append(newInfos, info)
			newTokens = append(newTokens, pool.RewardTokens[indexInfos])
			newAccumulatedRewards = append(newAccumulatedRewards, updatedPool.TotalAccumulatedRewards[indexInfos])
			indexInfos++

		case 0: // info denom == input denom: adjust existing reward, mint AddAmount, adjust RewardAmountPerBlock
			newInfo, err := k.processInputForExistingInfo(ctx, pool.Name, &input, &info)
			if err != nil {
				return err
			}
			newInfos = append(newInfos, *newInfo)
			newTokens = append(newTokens, pool.RewardTokens[indexInfos])
			newAccumulatedRewards = append(newAccumulatedRewards, updatedPool.TotalAccumulatedRewards[indexInfos])

			indexInfos++
			indexInputs++

		case 1: // info denom > input denom: add new reward token, mint AddAmount, set RewardAmountPerBlock, set RewardStartBlockHeight
			newInfo, err := k.processNewInput(ctx, pool.Name, &input)
			if err != nil {
				return err
			}
			newInfos = append(newInfos, *newInfo)
			newTokens = append(newTokens, tokensFromInputs[indexInputs])
			newAccumulatedRewards = append(newAccumulatedRewards, sdk.NewDecCoin(inputAmount.Denom, sdk.ZeroInt()))

			indexInputs++
		}
	}

	// 4. Set fields for updatedPool
	updatedPool.RewardTokenInfos = newInfos
	updatedPool.RewardTokens = newTokens
	updatedPool.TotalAccumulatedRewards = newAccumulatedRewards
	k.SetFarmingPool(ctx, updatedPool)

	return nil
}

// CheckAdjustRewardProposal checks the validity of an AdjustRewardProposal
func (k Keeper) CheckAdjustRewardProposal(ctx sdk.Context, p *types.AdjustRewardInfo) (*types.FarmingPool, []commontypes.ERC20Token, error) {
	// 1.1. Check pool existence
	pool, found := k.GetFarmingPool(ctx, p.PoolName)
	if !found {
		return nil, nil, types.WrapErrPoolNotExist(p.PoolName)
	}
	// 1.1. Check inputs non-empty
	if len(p.RewardAdjustmentInputs) == 0 {
		return nil, nil, types.WrapErrInvalidInput("empty reward token info array")
	}
	// 1.2. Check reward tokens exist and sorted by ascending denom
	tokens := []commontypes.ERC20Token{}
	lastDenom := ""
	for _, input := range p.RewardAdjustmentInputs {
		denom := input.AddAmount.Denom
		chainId, symbol, parseErr := common.ParseERC20TokenDenom(denom)
		if parseErr != nil {
			return nil, nil, parseErr
		}
		token, found := k.GetERC20Token(ctx, chainId, symbol)
		if !found {
			return nil, nil, types.WrapErrTokenNotExist(denom)
		}
		if denom <= lastDenom {
			return nil, nil, types.WrapErrInvalidInput(
				fmt.Sprintf("reward token denoms out-of-order: %s vs %s", lastDenom, denom))
		}
		lastDenom = denom
		tokens = append(tokens, token)
	}
	return &pool, tokens, nil
}

// processNewInput mints AddAmount, sets RewardAmountPerBlock, sets RewardStartBlockHeight
func (k Keeper) processNewInput(ctx sdk.Context, poolName string, input *types.RewardAdjustmentInput) (
	newInfo *types.RewardTokenInfo, err error) {
	truncatedAddAmount, _ := input.AddAmount.TruncateDecimal()
	// 1. Mint reward
	mintErr := k.bankKeeper.MintCoins(ctx, types.RewardModuleAccountName, sdk.NewCoins(truncatedAddAmount))
	if mintErr != nil {
		return nil, types.WrapErrMintCoinsFailed(mintErr.Error())
	}
	// 2. Update info
	startBlockHeight := ctx.BlockHeight() + input.RewardStartBlockDelay
	newInfo = &types.RewardTokenInfo{
		RemainingAmount:        sdk.NewDecCoinFromCoin(truncatedAddAmount),
		RewardStartBlockHeight: startBlockHeight,
		RewardAmountPerBlock:   input.NewRewardAmountPerBlock,
	}
	// 3. Emit event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAdjustReward,
		sdk.NewAttribute(types.AttributeKeyPool, poolName),
		sdk.NewAttribute(types.AttributeKeyAddAmount, input.AddAmount.String()),
		sdk.NewAttribute(types.AttributeKeyRewardStartHeight, strconv.FormatInt(startBlockHeight, 10)),
		sdk.NewAttribute(types.AttributeKeyRewardAmountPerBlock, input.NewRewardAmountPerBlock.String()),
	))
	return newInfo, nil
}

// processInputForExistingInfo adjusts existing reward, mints AddAmount, adjusts RewardAmountPerBlock
// NOTE: only sets RewardStartBlockHeight when existing value is 0. i.e. restarting a reward
func (k Keeper) processInputForExistingInfo(ctx sdk.Context, poolName string, input *types.RewardAdjustmentInput, info *types.RewardTokenInfo) (
	newInfo *types.RewardTokenInfo, err error) {
	truncatedAddAmount, _ := input.AddAmount.TruncateDecimal()
	// 1. Mint reward
	mintErr := k.bankKeeper.MintCoins(ctx, types.RewardModuleAccountName, sdk.NewCoins(truncatedAddAmount))
	if mintErr != nil {
		return nil, types.WrapErrMintCoinsFailed(mintErr.Error())
	}
	// 2. Update info
	startBlockHeight := info.RewardStartBlockHeight
	if startBlockHeight == 0 {
		startBlockHeight = ctx.BlockHeight() + input.RewardStartBlockDelay
	}
	newInfo = &types.RewardTokenInfo{
		RemainingAmount:        info.RemainingAmount.Add(sdk.NewDecCoinFromCoin(truncatedAddAmount)),
		RewardStartBlockHeight: startBlockHeight,
		RewardAmountPerBlock:   input.NewRewardAmountPerBlock,
	}
	// 3. Emit event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAdjustReward,
		sdk.NewAttribute(types.AttributeKeyPool, poolName),
		sdk.NewAttribute(types.AttributeKeyAddAmount, input.AddAmount.String()),
		sdk.NewAttribute(types.AttributeKeyRewardStartHeight, strconv.FormatInt(info.RewardStartBlockHeight, 10)),
		sdk.NewAttribute(types.AttributeKeyRewardAmountPerBlock, input.NewRewardAmountPerBlock.String()),
	))
	return newInfo, nil
}
