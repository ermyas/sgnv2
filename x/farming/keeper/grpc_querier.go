package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

// Params queries params of distribution module
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Pools queries the current state of all the pools.
func (k Keeper) Pools(c context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	pools := k.GetFarmingPools(ctx)
	var updatedPools types.FarmingPools
	for _, pool := range pools {
		updatedPool, _ := k.CalculateAmountEarnedBetween(ctx, pool)
		updatedPools = append(updatedPools, updatedPool)
	}
	return &types.QueryPoolsResponse{Pools: updatedPools}, nil
}

// Pool queries the current state of a single pool.
func (k Keeper) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.PoolName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty pool name")
	}
	ctx := sdk.UnwrapSDKContext(c)

	poolName := req.PoolName
	pool, found := k.GetFarmingPool(ctx, poolName)
	if !found {
		return nil, status.Errorf(codes.NotFound, "pool %s not found", poolName)
	}
	updatedPool, _ := k.CalculateAmountEarnedBetween(ctx, pool)
	return &types.QueryPoolResponse{Pool: updatedPool}, nil
}

// Tokens queries all the tokens.
func (k Keeper) Tokens(c context.Context, req *types.QueryTokensRequest) (*types.QueryTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	tokens := k.GetERC20Tokens(ctx)
	return &types.QueryTokensResponse{Tokens: tokens}, nil
}

// Token queries a single token.
func (k Keeper) Token(c context.Context, req *types.QueryTokenRequest) (*types.QueryTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.ChainId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid chain ID")
	}
	if req.Symbol == "" {
		return nil, status.Error(codes.InvalidArgument, "empty symbol")
	}
	ctx := sdk.UnwrapSDKContext(c)

	token, found := k.GetERC20Token(ctx, req.ChainId, req.Symbol)
	if !found {
		return nil, status.Errorf(codes.NotFound, "token %s/%d not found", req.Symbol, req.ChainId)
	}
	return &types.QueryTokenResponse{Token: token}, nil
}

// RewardContracts queries all the reward contracts.
func (k Keeper) RewardContracts(c context.Context, req *types.QueryRewardContractsRequest) (*types.QueryRewardContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	var infos []commontypes.ContractInfo
	k.IterateAllRewardContracts(ctx, func(info commontypes.ContractInfo) bool {
		infos = append(infos, info)
		return false
	})
	return &types.QueryRewardContractsResponse{RewardContracts: infos}, nil

}

// RewardContract queries a single reward contract.
func (k Keeper) RewardContract(c context.Context, req *types.QueryRewardContractRequest) (*types.QueryRewardContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.ChainId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid chain ID")
	}
	ctx := sdk.UnwrapSDKContext(c)

	info, found := k.GetRewardContract(ctx, req.ChainId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "reward contract for chain ID %d not found", req.ChainId)
	}
	return &types.QueryRewardContractResponse{RewardContract: info}, nil
}

// Earnings queries the current earnings of an account in a pool.
func (k Keeper) Earnings(c context.Context, req *types.QueryEarningsRequest) (*types.QueryEarningsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.PoolName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty pool name")
	}
	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "empty address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	earnings, err := k.GetEarnings(ctx, req.PoolName, eth.Hex2Addr(req.Address))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryEarningsResponse{Earnings: earnings}, nil
}

// StakeInfo queries the current stake info of an account in a pool.
func (k Keeper) StakeInfo(c context.Context, req *types.QueryStakeInfoRequest) (*types.QueryStakeInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.PoolName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty pool name")
	}
	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "empty address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	stakeInfo, found := k.GetStakeInfo(ctx, eth.Hex2Addr(req.Address), req.PoolName)
	if !found {
		return nil, status.Errorf(codes.NotFound, "stake info not found")
	}
	return &types.QueryStakeInfoResponse{StakeInfo: stakeInfo}, nil
}

// AccountInfo queries the current state of a farming account.
func (k Keeper) AccountInfo(c context.Context, req *types.QueryAccountInfoRequest) (*types.QueryAccountInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "empty address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	addr := eth.Hex2Addr(req.Address)
	// StakedPools
	poolNames := k.GetFarmingPoolNamesForAccount(ctx, addr)
	var updatedPools types.FarmingPools
	for _, poolName := range poolNames {
		pool, found := k.GetFarmingPool(ctx, poolName)
		if !found {
			return nil, status.Errorf(codes.NotFound, "pool %s not found", poolName)
		}
		updatedPool, _ := k.CalculateAmountEarnedBetween(ctx, pool)
		updatedPools = append(updatedPools, updatedPool)
	}

	// EarningsList
	var earningsList []types.Earnings
	for _, poolName := range poolNames {
		earnings, sdkErr := k.GetEarnings(ctx, poolName, addr)
		if sdkErr != nil {
			return nil, sdkErr
		}
		earningsList = append(earningsList, earnings)
	}

	// CumulativeRewards (settled rewards + outstanding earnings)
	derivedRewardAccount := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, addr)
	// TODO: Avoid sorting?
	cumulativeRewards := sdk.NewDecCoinsFromCoins(k.bankKeeper.GetAllBalances(ctx, derivedRewardAccount)...) // sorted
	for _, earnings := range earningsList {
		cumulativeRewards = cumulativeRewards.Add(earnings.RewardAmounts.Sort()...)
	}

	accountInfo := types.AccountInfo{
		StakedPools:             updatedPools,
		EarningsList:            earningsList,
		CumulativeRewardAmounts: cumulativeRewards,
	}
	return &types.QueryAccountInfoResponse{AccountInfo: accountInfo}, nil
}

func (k Keeper) AccountsStakedIn(c context.Context, req *types.QueryAccountsStakedInRequest) (*types.QueryAccountsStakedInResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.PoolName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty pool name")
	}
	ctx := sdk.UnwrapSDKContext(c)

	addrList := k.GetAccountsStakedIn(ctx, req.PoolName)
	var addrs []string
	for _, addr := range addrList {
		addrs = append(addrs, eth.Addr2Hex(addr))
	}
	return &types.QueryAccountsStakedInResponse{Addresses: addrs}, nil
}

func (k Keeper) NumPools(c context.Context, req *types.QueryNumPoolsRequest) (*types.QueryNumPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	numPools := k.GetNumPools(ctx)
	return &types.QueryNumPoolsResponse{NumPools: numPools.NumPools}, nil
}

func (k Keeper) RewardClaimInfo(c context.Context, req *types.QueryRewardClaimInfoRequest) (*types.QueryRewardClaimInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "empty address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	info, found := k.GetRewardClaimInfo(ctx, eth.Hex2Addr(req.Address))
	if !found {
		return nil, status.Errorf(codes.NotFound, "reward claim info not found")
	}
	return &types.QueryRewardClaimInfoResponse{RewardClaimInfo: info}, nil
}
