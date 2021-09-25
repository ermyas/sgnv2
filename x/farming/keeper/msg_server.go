package keeper

import (
	"context"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	err := k.claimOnePool(ctx, msg.PoolName, eth.Hex2Addr(msg.Address))
	if err != nil {
		return nil, err
	}

	return &types.MsgClaimRewardsResponse{}, nil
}

func (k msgServer) ClaimAllRewards(
	goCtx context.Context, msg *types.MsgClaimAllRewards) (*types.MsgClaimAllRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr := eth.Hex2Addr(msg.Address)
	poolNames := k.GetFarmingPoolNamesForAccount(ctx, addr)
	for _, poolName := range poolNames {
		err := k.claimOnePool(ctx, poolName, addr)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgClaimAllRewardsResponse{}, nil
}

func (k msgServer) claimOnePool(ctx sdk.Context, poolName string, addr eth.Addr) error {
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

	// 5. Update farming pool
	// NOTE: Will panic if TotalAccumulatedRewards < rewards
	updatedPool.TotalAccumulatedRewards = updatedPool.TotalAccumulatedRewards.Sub(rewards)
	k.SetFarmingPool(ctx, updatedPool)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaim,
		sdk.NewAttribute(types.AttributeKeyAddress, addr.String()),
		sdk.NewAttribute(types.AttributeKeyPool, poolName),
	))

	return nil
}
