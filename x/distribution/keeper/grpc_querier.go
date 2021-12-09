package keeper

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ types.QueryServer = Keeper{}

// Params queries params of distribution module
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)

	return &types.QueryParamsResponse{Params: params}, nil
}

// ValidatorOutstandingRewards queries rewards of a validator address
func (k Keeper) ValidatorOutstandingRewards(c context.Context, req *types.QueryValidatorOutstandingRewardsRequest) (*types.QueryValidatorOutstandingRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.ValidatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty validator address")
	}

	ctx := sdk.UnwrapSDKContext(c)

	valAdr := eth.Hex2Addr(req.ValidatorAddress)
	rewards := k.GetValidatorOutstandingRewards(ctx, valAdr)

	return &types.QueryValidatorOutstandingRewardsResponse{Rewards: rewards}, nil
}

// ValidatorCommission queries accumulated commission for a validator
func (k Keeper) ValidatorCommission(c context.Context, req *types.QueryValidatorCommissionRequest) (*types.QueryValidatorCommissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.ValidatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty validator address")
	}

	ctx := sdk.UnwrapSDKContext(c)

	valAdr := eth.Hex2Addr(req.ValidatorAddress)
	commission := k.GetValidatorAccumulatedCommission(ctx, valAdr)

	return &types.QueryValidatorCommissionResponse{Commission: commission}, nil
}

// DelegationRewards the total rewards accrued by a delegation
func (k Keeper) DelegationRewards(c context.Context, req *types.QueryDelegationRewardsRequest) (*types.QueryDelegationRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.DelegatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty delegator address")
	}

	if req.ValidatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty validator address")
	}

	ctx := sdk.UnwrapSDKContext(c)

	valAdr := eth.Hex2Addr(req.ValidatorAddress)

	val := k.stakingKeeper.Validator(ctx, valAdr)
	if val == nil {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorExists, req.ValidatorAddress)
	}

	delAdr := eth.Hex2Addr(req.DelegatorAddress)
	del := k.stakingKeeper.Delegation(ctx, delAdr, valAdr)
	if del == nil {
		return nil, types.ErrNoDelegationExists
	}

	endingPeriod := k.IncrementValidatorPeriod(ctx, val)
	rewards := k.CalculateDelegationRewards(ctx, val, del, endingPeriod)

	return &types.QueryDelegationRewardsResponse{Rewards: rewards}, nil
}

// DelegationTotalRewards the total rewards accrued by a each validator
func (k Keeper) DelegationTotalRewards(c context.Context, req *types.QueryDelegationTotalRewardsRequest) (*types.QueryDelegationTotalRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.DelegatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty delegator address")
	}

	ctx := sdk.UnwrapSDKContext(c)

	total := sdk.DecCoins{}
	var delRewards []types.DelegationDelegatorReward

	delAdr := eth.Hex2Addr(req.DelegatorAddress)

	k.stakingKeeper.IterateDelegations(
		ctx, delAdr,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			valAddr := del.GetValidatorAddr()
			val := k.stakingKeeper.Validator(ctx, valAddr)
			endingPeriod := k.IncrementValidatorPeriod(ctx, val)
			delReward := k.CalculateDelegationRewards(ctx, val, del, endingPeriod)

			delRewards = append(delRewards, types.NewDelegationDelegatorReward(valAddr, delReward))
			total = total.Add(delReward...)
			return false
		},
	)

	return &types.QueryDelegationTotalRewardsResponse{Rewards: delRewards, Total: total}, nil
}

// DelegatorValidators queries the validators list of a delegator
func (k Keeper) DelegatorValidators(c context.Context, req *types.QueryDelegatorValidatorsRequest) (*types.QueryDelegatorValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.DelegatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty delegator address")
	}

	ctx := sdk.UnwrapSDKContext(c)
	delAdr := eth.Hex2Addr(req.DelegatorAddress)
	var validators []string

	k.stakingKeeper.IterateDelegations(
		ctx, delAdr,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			validators = append(validators, del.GetValidatorAddr().String())
			return false
		},
	)

	return &types.QueryDelegatorValidatorsResponse{Validators: validators}, nil
}

// DelegatorWithdrawAddress queries Query/delegatorWithdrawAddress
func (k Keeper) DelegatorWithdrawAddress(c context.Context, req *types.QueryDelegatorWithdrawAddressRequest) (*types.QueryDelegatorWithdrawAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.DelegatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty delegator address")
	}
	delAdr := eth.Hex2Addr(req.DelegatorAddress)
	ctx := sdk.UnwrapSDKContext(c)
	withdrawAddr := k.GetDelegatorWithdrawAddr(ctx, delAdr)

	return &types.QueryDelegatorWithdrawAddressResponse{WithdrawAddress: withdrawAddr.String()}, nil
}

// CommunityPool queries the community pool coins
func (k Keeper) CommunityPool(c context.Context, req *types.QueryCommunityPoolRequest) (*types.QueryCommunityPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	pool := k.GetFeePoolCommunityCoins(ctx)

	return &types.QueryCommunityPoolResponse{Pool: pool}, nil
}

func (k Keeper) StakingRewardInfo(c context.Context, req *types.QueryStakingRewardInfoRequest) (*types.QueryStakingRewardInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.DelegatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// Outstanding reward
	delAddr := eth.Hex2Addr(req.DelegatorAddress)
	totalOutstandingReward := sdk.NewDecCoin(types.StakingRewardDenom, sdk.ZeroInt())
	k.stakingKeeper.IterateDelegations(
		ctx, delAddr,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			valAddr := del.GetValidatorAddr()
			val := k.stakingKeeper.Validator(ctx, valAddr)
			endingPeriod := k.IncrementValidatorPeriod(ctx, val)
			outstandingRewards := k.CalculateDelegationRewards(ctx, val, del, endingPeriod)
			for _, reward := range outstandingRewards {
				if reward.Denom == types.StakingRewardDenom {
					totalOutstandingReward = totalOutstandingReward.Add(reward)
					break
				}
			}
			return false
		},
	)

	// Cumulative reward (settled + outstanding rewards)
	derivedRewardAccount := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, delAddr)
	cumulativeReward := sdk.NewDecCoinFromCoin(k.bankKeeper.GetBalance(ctx, derivedRewardAccount, types.StakingRewardDenom))
	cumulativeReward = cumulativeReward.Add(totalOutstandingReward)

	// Claimed reward
	claimedReward := sdk.NewDecCoin(types.StakingRewardDenom, sdk.ZeroInt())
	info, found := k.GetStakingRewardClaimInfo(ctx, eth.Hex2Addr(req.DelegatorAddress))
	if found {
		claimedReward = info.CumulativeRewardAmount
	}

	rewardInfo := types.StakingRewardInfo{
		CumulativeRewardAmount: cumulativeReward,
		ClaimedRewardAmount:    claimedReward,
	}
	return &types.QueryStakingRewardInfoResponse{RewardInfo: rewardInfo}, nil
}

func (k Keeper) StakingRewardClaimInfo(
	c context.Context, req *types.QueryStakingRewardClaimInfoRequest) (*types.QueryStakingRewardClaimInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.DelegatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty delegator address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	info, found := k.GetStakingRewardClaimInfo(ctx, eth.Hex2Addr(req.DelegatorAddress))
	if !found {
		return nil, status.Errorf(codes.NotFound, "reward claim info not found")
	}
	return &types.QueryStakingRewardClaimInfoResponse{RewardClaimInfo: info}, nil
}

func (k Keeper) CBridgeFeeShareInfo(
	c context.Context, req *types.QueryCBridgeFeeShareInfoRequest) (*types.QueryCBridgeFeeShareInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.DelegatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty delegator address")
	}
	ctx := sdk.UnwrapSDKContext(c)
	feeShareInfo := k.getClaimableFeesInfo(ctx, eth.Hex2Addr(req.DelegatorAddress), cbrtypes.CBridgeFeeDenomPrefix)
	return &types.QueryCBridgeFeeShareInfoResponse{FeeShareInfo: feeShareInfo}, nil
}

func (k Keeper) PegBridgeFeesInfo(
	c context.Context, req *types.QueryPegBridgeFeesInfoRequest) (*types.QueryPegBridgeFeesInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.DelegatorAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "empty delegator address")
	}
	ctx := sdk.UnwrapSDKContext(c)
	feesInfo := k.getClaimableFeesInfo(ctx, eth.Hex2Addr(req.DelegatorAddress), pegbrtypes.PegBridgeFeeDenomPrefix)
	return &types.QueryPegBridgeFeesInfoResponse{FeesInfo: feesInfo}, nil
}

func (k Keeper) getClaimableFeesInfo(
	ctx sdk.Context, delAddr eth.Addr, denomPrefix string) types.ClaimableFeesInfo {
	// Outstanding fees
	totalOutstandingFees := sdk.DecCoins{}
	k.stakingKeeper.IterateDelegations(
		ctx, delAddr,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			valAddr := del.GetValidatorAddr()
			val := k.stakingKeeper.Validator(ctx, valAddr)
			endingPeriod := k.IncrementValidatorPeriod(ctx, val)
			outstandingRewards := k.CalculateDelegationRewards(ctx, val, del, endingPeriod)
			for _, reward := range outstandingRewards {
				if strings.HasPrefix(reward.Denom, denomPrefix) {
					totalOutstandingFees = totalOutstandingFees.Add(reward)
				}
			}
			return false
		},
	)

	// Claimable fees (settled + outstanding fees)
	derivedRewardAccount := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, delAddr)
	balances := k.bankKeeper.GetAllBalances(ctx, derivedRewardAccount) // sorted
	settledFees := sdk.DecCoins{}
	for _, coin := range balances {
		if strings.HasPrefix(coin.Denom, denomPrefix) {
			settledFees = settledFees.Add(sdk.NewDecCoinFromCoin(coin))
		}
	}

	claimableFees := settledFees.Add(totalOutstandingFees.Sort()...)
	claimableFeesInfo := types.ClaimableFeesInfo{
		ClaimableFeeAmounts: claimableFees,
	}
	return claimableFeesInfo
}
