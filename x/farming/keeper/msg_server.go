package keeper

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	ethproto "github.com/celer-network/sgn-v2/proto/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the farming MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) ClaimRewards(goCtx context.Context, msg *types.MsgClaimRewards) (*types.MsgClaimRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	claimInfo, err := k.checkCooldownAndUpdateClaimTime(ctx, msg.Address)
	if err != nil {
		return nil, err
	}

	addr := eth.Hex2Addr(msg.Address)
	err = k.claimOnePool(ctx, msg.PoolName, addr, claimInfo)
	if err != nil {
		return nil, err
	}

	// Emit claim event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaim,
		sdk.NewAttribute(types.AttributeKeyAddress, addr.String()),
		sdk.NewAttribute(types.AttributeKeyPool, msg.PoolName),
	))

	return &types.MsgClaimRewardsResponse{}, nil
}

func (k msgServer) ClaimAllRewards(
	goCtx context.Context, msg *types.MsgClaimAllRewards) (*types.MsgClaimAllRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	claimInfo, err := k.checkCooldownAndUpdateClaimTime(ctx, msg.Address)
	if err != nil {
		return nil, err
	}

	addr := eth.Hex2Addr(msg.Address)
	poolNames := k.GetFarmingPoolNamesForAccount(ctx, addr)
	for _, poolName := range poolNames {
		claimErr := k.claimOnePool(ctx, poolName, addr, claimInfo)
		if claimErr != nil {
			return nil, claimErr
		}
	}

	// Emit claim_all event to trigger validators signing
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaimAll,
		sdk.NewAttribute(types.AttributeKeyAddress, addr.String()),
	))

	return &types.MsgClaimAllRewardsResponse{}, nil
}

func (k msgServer) checkCooldownAndUpdateClaimTime(ctx sdk.Context, addr string) (*types.RewardClaimInfo, error) {
	// Reject if last claim is too recent
	blockTime := ctx.BlockTime()
	claimInfo, found := k.GetRewardClaimInfo(ctx, eth.Hex2Addr(addr))
	if found && blockTime.Before(claimInfo.LastClaimTime.Add(k.GetClaimCooldown(ctx))) {
		return nil, types.WrapErrClaimCooldownNotPassed(claimInfo.LastClaimTime)
	}
	// Initialize claimInfo if not present
	if !found {
		claimInfo = types.RewardClaimInfo{
			Recipient: addr,
		}
	}
	claimInfo.LastClaimTime = blockTime
	k.SetRewardClaimInfo(ctx, claimInfo)
	return &claimInfo, nil
}

func (k msgServer) claimOnePool(
	ctx sdk.Context, poolName string, addr eth.Addr, claimInfo *types.RewardClaimInfo) error {
	// 1. Get the pool info
	pool, poolFound := k.GetFarmingPool(ctx, poolName)
	if !poolFound {
		return types.WrapErrNoFarmingPoolFound(poolName)
	}

	// 2. Calculate how many reward tokens can be earned in the current period
	updatedPool, earnedTokens := k.CalculateAmountEarnedBetween(ctx, pool)

	// 3. Withdraw rewards
	rewards, err := k.WithdrawRewards(ctx, pool.Name, pool.TotalStakedAmount, earnedTokens, addr)
	if err != nil {
		return err
	}

	// 4. Update StakeInfo
	k.UpdateStakeInfo(ctx, addr, pool.Name, sdk.ZeroDec())

	// 5. Update FarmingPool
	// NOTE: Will panic if TotalAccumulatedRewards < rewards
	updatedPool.TotalAccumulatedRewards = updatedPool.TotalAccumulatedRewards.Sub(rewards)
	k.SetFarmingPool(ctx, updatedPool)

	// 6. Update RewardClaimInfo
	// 6.1. Collect chainIds and RewardClaimDetails, create details if not existent
	chainIds := make(map[uint64]bool)
	for _, rewardToken := range pool.RewardTokens {
		chainIds[rewardToken.ChainId] = true
	}
	chainIdToDetails := make(map[uint64]*types.RewardClaimDetails)
	for _, detail := range claimInfo.RewardClaimDetailsList {
		chainIdToDetails[detail.ChainId] = &detail
	}
	for chainId := range chainIds {
		_, found := chainIdToDetails[chainId]
		if !found {
			chainIdToDetails[chainId] = &types.RewardClaimDetails{
				ChainId:                 chainId,
				CumulativeRewardAmounts: sdk.NewDecCoins(),
			}
		}
	}
	// 6.2. Add the amounts from this claim to CumulativeRewardAmounts
	for _, reward := range rewards {
		chainId, _, parseErr := ParseERC20TokenDenom(reward.Denom)
		if parseErr != nil {
			return parseErr
		}
		detail := chainIdToDetails[chainId]
		detail.CumulativeRewardAmounts = detail.CumulativeRewardAmounts.Add(reward)
	}
	// 6.3. Update TokenAddresses and CumulativeRewardAmounts, Reconstruct RewardProtoBytes
	for chainId, details := range chainIdToDetails {
		var tokenAddresses [][]byte
		var cumulativeRewardAmounts [][]byte
		for _, coin := range details.CumulativeRewardAmounts {
			chainId, symbol, parseErr := ParseERC20TokenDenom(coin.Denom)
			if parseErr != nil {
				return parseErr
			}
			token, found := k.GetERC20Token(ctx, chainId, symbol)
			if !found {
				return types.WrapErrTokenNotExist(coin.Denom)
			}
			tokenAddresses = append(tokenAddresses, eth.Hex2Addr(token.Address).Bytes())
			cumulativeRewardAmounts = append(cumulativeRewardAmounts, coin.Amount.BigInt().Bytes())
		}
		// Marshal RewardProtoBytes
		rewardProtoBytes, marshalErr := proto.Marshal(
			&ethproto.FarmingRewards{
				Recipient:               addr.Bytes(),
				ChainId:                 new(big.Int).SetUint64(chainId).Bytes(),
				TokenAddresses:          tokenAddresses,
				CumulativeRewardAmounts: cumulativeRewardAmounts,
			})
		if marshalErr != nil {
			return marshalErr
		}
		details.RewardProtoBytes = rewardProtoBytes
	}

	// 6.3. Append RewardClaimDetails and set RewardClaimInfo
	// TODO: 1. Avoid copying 2. Sort by ascending chain IDs?
	for _, details := range chainIdToDetails {
		claimInfo.RewardClaimDetailsList = append(claimInfo.RewardClaimDetailsList, *details)
	}
	k.SetRewardClaimInfo(ctx, *claimInfo)

	return nil
}

func (k msgServer) SignRewards(
	goCtx context.Context, msg *types.MsgSignRewards) (*types.MsgSignRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validator := k.stakingKeeper.ValidatorByConsAddr(ctx, sdk.ConsAddress(msg.Sender))
	if !validator.IsBonded() {
		return nil, fmt.Errorf("validator not bonded")
	}

	claimInfo, found := k.GetRewardClaimInfo(ctx, eth.Hex2Addr(msg.Address))
	if !found {
		return nil, types.WrapErrNoClaimInfoFound(msg.Address)
	}
	if len(claimInfo.RewardClaimDetailsList) == 0 {
		return nil, types.WrapErrInvalidInput("empty reward claim details list")
	}

	chainIdToRewardClaimDetails := make(map[uint64]*types.RewardClaimDetails)
	for _, detail := range claimInfo.RewardClaimDetailsList {
		chainIdToRewardClaimDetails[detail.ChainId] = &detail
	}
	for _, signatureDetails := range msg.SignatureDetailsList {
		addSigErr := chainIdToRewardClaimDetails[signatureDetails.ChainId].AddSig(
			signatureDetails.Signature,
			validator.GetSignerAddr().String(),
		)
		if addSigErr != nil {
			return nil, fmt.Errorf("failed to add sig: %s", addSigErr)
		}
	}
	k.SetRewardClaimInfo(ctx, claimInfo)
	return &types.MsgSignRewardsResponse{}, nil
}
