package keeper

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryValidator:
			return queryValidator(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidatorBySgnAddr:
			return queryValidatorBySgnAddr(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidatorByConsAddr:
			return queryValidatorByConsAddr(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidators:
			return queryValidators(ctx, k, legacyQuerierCdc)
		case types.QueryTransactors:
			return queryTransactors(ctx, req, k, legacyQuerierCdc)
		case types.QueryDelegation:
			return queryDelegation(ctx, req, k, legacyQuerierCdc)
		case types.QueryDelegations:
			return queryDelegations(ctx, req, k, legacyQuerierCdc)
		case types.QuerySgnAccount:
			return querySgnAccountExist(ctx, req, k, legacyQuerierCdc)
		case types.QuerySyncer:
			return querySyncer(ctx, k, legacyQuerierCdc)
		case types.QueryParams:
			return queryParams(ctx, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Unknown validator query endpoint")
		}
	}
}

func queryValidator(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	validator, found := k.GetValidator(ctx, eth.Hex2Addr(params.EthAddress))
	if !found {
		return nil, types.ErrValidatorNotFound
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validator)
	if err != nil {
		log.Error(err)
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorBySgnAddr(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorBySgnAddrParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	sgnAddr, err := sdk.AccAddressFromBech32(params.SgnAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAddress, err.Error())
	}
	validator, found := k.GetValidatorBySgnAddr(ctx, sgnAddr)
	if !found {
		return nil, types.ErrValidatorNotFound
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validator)
	if err != nil {
		log.Error(err)
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorByConsAddr(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorByConsAddrParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	consAddr, err := sdk.ConsAddressFromBech32(params.ConsAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAddress, err.Error())
	}
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		return nil, types.ErrValidatorNotFound
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validator)
	if err != nil {
		log.Error(err)
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidators(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	validators := k.GetAllValidators(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validators)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryTransactors(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryTransactorsParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	transactors := k.GetTransactors(ctx, eth.Hex2Addr(params.ValAddress))

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, transactors)
	if err != nil {
		log.Error(err)
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryDelegation(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDelegationParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	delegation, found :=
		k.GetDelegation(
			ctx,
			eth.Hex2Addr(params.DelAddress),
			eth.Hex2Addr(params.ValAddress))
	if !found {
		return nil, types.ErrDelegationNotFound
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, delegation)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// query all delegations on a given validator
func queryDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDelegationsParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	delegations := k.GetAllDelegations(ctx, eth.Hex2Addr(params.ValAddress))
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, delegations)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func querySgnAccountExist(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QuerySgnAccountParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	acctAddr, err := sdk.AccAddressFromBech32(params.SgnAddress)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidAddress, "%s", params.SgnAddress)
	}
	account := k.sdkAccountKeeper.GetAccount(ctx, acctAddr)
	if account == nil {
		return nil, sdkerrors.Wrapf(types.ErrSgnAccounNotFound, "%s", params.SgnAddress)
	}

	return []byte{1}, nil
}

func querySyncer(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	syncer := k.GetSyncer(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, syncer)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
