package keeper

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
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

	err = k.accumulateRewards(ctx, addr, claimInfo)
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

	err = k.accumulateRewards(ctx, addr, claimInfo)
	if err != nil {
		return nil, err
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
		return types.WrapErrPoolNotExist(poolName)
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
	origAccumulatedRewards := updatedPool.TotalAccumulatedRewards
	var hasNeg bool
	updatedPool.TotalAccumulatedRewards, hasNeg = origAccumulatedRewards.SafeSub(rewards)
	if hasNeg {
		return types.WrapErrInsufficientAmount(origAccumulatedRewards.String(), rewards.String())
	}
	k.SetFarmingPool(ctx, updatedPool)

	return nil
}

// accumulateRewards updates RewardClaimInfo
func (k msgServer) accumulateRewards(ctx sdk.Context, addr eth.Addr, claimInfo *types.RewardClaimInfo) error {
	// 1. Collect chainIds
	chainIdToDetails := make(map[uint64]*types.RewardClaimDetails)
	for _, detail := range claimInfo.RewardClaimDetailsList {
		chainIdToDetails[detail.ChainId] = &detail
	}
	// 2. Update CumulativeRewardAmounts
	derivedRewardAccount := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, addr)
	rewards := k.bankKeeper.GetAllBalances(ctx, derivedRewardAccount)
	for _, reward := range rewards {
		denom := reward.Denom
		cumulativeReward := k.bankKeeper.GetBalance(ctx, derivedRewardAccount, denom)
		chainId, _, parseErr := ParseERC20TokenDenom(denom)
		if parseErr != nil {
			return parseErr
		}
		details, found := chainIdToDetails[chainId]
		if !found {
			// Create details if not existent
			details = &types.RewardClaimDetails{
				ChainId:                 chainId,
				CumulativeRewardAmounts: sdk.NewDecCoins(),
			}
			chainIdToDetails[chainId] = details
		}
		cumulativeRewardAmount := cumulativeReward.Amount
		existing := sdk.NewDecCoinFromDec(denom, details.CumulativeRewardAmounts.AmountOf(denom))
		updated := sdk.NewDecCoin(denom, cumulativeRewardAmount)
		if !existing.Amount.Equal(updated.Amount) {
			details.CumulativeRewardAmounts =
				details.CumulativeRewardAmounts.Sub(sdk.NewDecCoins(existing)).Add(updated)
		}
	}
	// 3. Update TokenAddresses and CumulativeRewardAmounts, Reconstruct RewardProtoBytes
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
			cumulativeRewardAmounts = append(cumulativeRewardAmounts, coin.Amount.RoundInt().BigInt().Bytes())
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

	// 4.1.  Append RewardClaimDetails and set RewardClaimInfo
	// TODO: 1. Avoid copying 2. Sort by ascending chain IDs?
	claimInfo.RewardClaimDetailsList = []types.RewardClaimDetails{}
	for _, details := range chainIdToDetails {
		claimInfo.RewardClaimDetailsList = append(claimInfo.RewardClaimDetailsList, *details)
	}
	// 4.2. Clear stale signatures
	for i := 0; i < len(claimInfo.RewardClaimDetailsList); i++ {
		detail := &claimInfo.RewardClaimDetailsList[i]
		detail.Signatures = []commontypes.Signature{}
	}
	k.SetRewardClaimInfo(ctx, *claimInfo)
	return nil
}

func (k msgServer) SignRewards(
	goCtx context.Context, msg *types.MsgSignRewards) (*types.MsgSignRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAcct, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, types.WrapErrInvalidAddress(msg.Sender)
	}
	validator, found := k.stakingKeeper.GetValidatorBySgnAddr(ctx, senderAcct)
	if !found {
		return nil, fmt.Errorf("sender is not a validator")
	}
	if !validator.IsBonded() {
		return nil, fmt.Errorf("validator is not bonded")
	}

	claimInfo, found := k.GetRewardClaimInfo(ctx, eth.Hex2Addr(msg.Address))
	if !found {
		return nil, types.WrapErrNoClaimInfoFound(msg.Address)
	}
	if len(claimInfo.RewardClaimDetailsList) == 0 {
		return nil, types.WrapErrInvalidInput("empty reward claim details list")
	}

	chainIdToRewardClaimDetails := make(map[uint64]*types.RewardClaimDetails)
	for i := 0; i < len(claimInfo.RewardClaimDetailsList); i++ {
		detail := &claimInfo.RewardClaimDetailsList[i]
		chainIdToRewardClaimDetails[detail.ChainId] = detail
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
