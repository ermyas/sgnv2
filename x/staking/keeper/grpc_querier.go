package keeper

import (
	"bytes"
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Validators queries all validators that match the given status
func (k Querier) Validators(c context.Context, req *types.QueryValidatorsRequest) (*types.QueryValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// validate the provided status, return all the validators if the status is empty
	if req.Status != "" && !(req.Status == types.Bonded.String() || req.Status == types.Unbonded.String() || req.Status == types.Unbonding.String()) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid validator status %s", req.Status)
	}

	var validators types.Validators
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	valStore := prefix.NewStore(store, types.ValidatorKey)

	pageRes, err := query.FilteredPaginate(valStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		val, err := types.UnmarshalValidator(k.cdc, value)
		if err != nil {
			return false, err
		}

		if req.Status != "" && !strings.EqualFold(val.GetStatus().String(), req.Status) {
			return false, nil
		}

		if accumulate {
			validators = append(validators, val)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryValidatorsResponse{Validators: validators, Pagination: pageRes}, nil
}

// Validator queries validator info for given validator address
func (k Querier) Validator(c context.Context, req *types.QueryValidatorRequest) (*types.QueryValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	valAddr := eth.Hex2Addr(req.ValidatorAddr)

	ctx := sdk.UnwrapSDKContext(c)
	validator, found := k.GetValidator(ctx, valAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "validator %s not found", req.ValidatorAddr)
	}

	return &types.QueryValidatorResponse{Validator: validator}, nil
}

// ValidatorDelegations queries delegate info for given validator
func (k Querier) ValidatorDelegations(c context.Context, req *types.QueryValidatorDelegationsRequest) (*types.QueryValidatorDelegationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}
	var delegations []types.Delegation
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	valStore := prefix.NewStore(store, types.DelegationKey)
	pageRes, err := query.FilteredPaginate(valStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		delegation, err := types.UnmarshalDelegation(k.cdc, value)
		if err != nil {
			return false, err
		}

		valAddr := eth.Hex2Addr(req.ValidatorAddr)

		if !bytes.Equal(delegation.GetValidatorAddr().Bytes(), valAddr.Bytes()) {
			return false, nil
		}

		if accumulate {
			delegations = append(delegations, delegation)
		}
		return true, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	delResponses, err := DelegationsToDelegationResponses(ctx, k.Keeper, delegations)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryValidatorDelegationsResponse{
		DelegationResponses: delResponses, Pagination: pageRes}, nil
}

// Delegation queries delegate info for given validator delegator pair
func (k Querier) Delegation(c context.Context, req *types.QueryDelegationRequest) (*types.QueryDelegationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}
	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	delAddr := eth.Hex2Addr(req.DelegatorAddr)
	valAddr := eth.Hex2Addr(req.ValidatorAddr)

	delegation, found := k.GetDelegation(ctx, delAddr, valAddr)
	if !found {
		return nil, status.Errorf(
			codes.NotFound,
			"delegation with delegator %s not found for validator %s",
			req.DelegatorAddr, req.ValidatorAddr)
	}

	delResponse, err := DelegationToDelegationResponse(ctx, k.Keeper, delegation)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDelegationResponse{DelegationResponse: &delResponse}, nil
}

// DelegatorDelegations queries all delegations of a give delegator address
func (k Querier) DelegatorDelegations(c context.Context, req *types.QueryDelegatorDelegationsRequest) (*types.QueryDelegatorDelegationsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}
	var delegations types.Delegations
	ctx := sdk.UnwrapSDKContext(c)

	delAddr := eth.Hex2Addr(req.DelegatorAddr)

	store := ctx.KVStore(k.storeKey)
	delStore := prefix.NewStore(store, types.GetDelegationsKey(delAddr))
	pageRes, err := query.Paginate(delStore, req.Pagination, func(key []byte, value []byte) error {
		delegation, err := types.UnmarshalDelegation(k.cdc, value)
		if err != nil {
			return err
		}
		delegations = append(delegations, delegation)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	delegationResps, err := DelegationsToDelegationResponses(ctx, k.Keeper, delegations)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDelegatorDelegationsResponse{DelegationResponses: delegationResps, Pagination: pageRes}, nil

}

// DelegatorValidator queries validator info for given delegator validator pair
func (k Querier) DelegatorValidator(c context.Context, req *types.QueryDelegatorValidatorRequest) (*types.QueryDelegatorValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}
	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	delAddr := eth.Hex2Addr(req.DelegatorAddr)
	valAddr := eth.Hex2Addr(req.ValidatorAddr)

	validator, err := k.GetDelegatorValidator(ctx, delAddr, valAddr)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDelegatorValidatorResponse{Validator: validator}, nil
}

// DelegatorValidators queries all validators info for given delegator address
func (k Querier) DelegatorValidators(c context.Context, req *types.QueryDelegatorValidatorsRequest) (*types.QueryDelegatorValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}
	var validators types.Validators
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	delAddr := eth.Hex2Addr(req.DelegatorAddr)

	delStore := prefix.NewStore(store, types.GetDelegationsKey(delAddr))
	pageRes, err := query.Paginate(delStore, req.Pagination, func(key []byte, value []byte) error {
		delegation, err := types.UnmarshalDelegation(k.cdc, value)
		if err != nil {
			return err
		}

		validator, found := k.GetValidator(ctx, delegation.GetValidatorAddr())
		if !found {
			return types.ErrValidatorNotFound
		}

		validators = append(validators, validator)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDelegatorValidatorsResponse{Validators: validators, Pagination: pageRes}, nil
}

// Params queries the staking parameters
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
