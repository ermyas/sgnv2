package validator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QuerySyncer:
			return querySyncer(ctx, req, keeper)
		case QueryDelegator:
			return queryDelegator(ctx, req, keeper)
		case QueryValidator:
			return queryValidator(ctx, req, keeper)
		case QueryValidators:
			return queryValidators(ctx, req, keeper)
		case QueryValidatorDelegators:
			return queryValidatorDelegators(ctx, req, keeper)
		case QueryReward:
			return queryReward(ctx, req, keeper)
		case QueryRewardEpoch:
			return queryRewardEpoch(ctx, req, keeper)
		case QueryRewardStats:
			return queryRewardStats(ctx, req, keeper)
		case QueryParameters:
			return queryParameters(ctx, keeper)
		default:
			return nil, sdk_errors.Wrap(sdk_errors.ErrUnknownRequest, "Unknown validator query endpoint")
		}
	}
}

func querySyncer(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	// syncer := keeper.GetSyncer(ctx)
	// res, err := codec.MarshalJSONIndent(keeper.cdc, syncer)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())

	// }

	// return res, nil
	return nil, nil
}

func queryDelegator(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	//var params QueryDelegatorParams
	//err := ModuleCdc.UnmarshalJSON(req.Data, &params)
	//if err != nil {
	//	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONUnmarshal, err.Error())
	//}

	//delegator, found := keeper.GetDelegator(ctx, params.ValidatorAddress, params.DelegatorAddress)
	//if !found {
	//	return nil, fmt.Errorf("%w for delegator %s, candidate %s", common.ErrRecordNotFound, params.DelegatorAddress, params.ValidatorAddress)
	//}

	//res, err := codec.MarshalJSONIndent(keeper.cdc, delegator)
	//if err != nil {
	//	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())

	//}

	return nil, nil
}

func queryValidator(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	// var params QueryValidatorParams
	// err := ModuleCdc.UnmarshalJSON(req.Data, &params)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONUnmarshal, err.Error())
	// }

	// candidate, found := keeper.GetValidator(ctx, params.ValidatorAddress)
	// if !found {
	// 	return nil, fmt.Errorf("%w for candidate %s", common.ErrRecordNotFound, params.ValidatorAddress)
	// }

	// res, err := codec.MarshalJSONIndent(keeper.cdc, candidate)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())

	// }

	return nil, nil
}

func queryValidators(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	// candidates := keeper.GetAllValidators(ctx)
	// res, err := codec.MarshalJSONIndent(keeper.cdc, candidates)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())

	// }

	return nil, nil
}

func queryValidatorDelegators(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	// var params QueryValidatorParams
	// err := ModuleCdc.UnmarshalJSON(req.Data, &params)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONUnmarshal, err.Error())
	// }

	// delegators := keeper.GetAllDelegators(ctx, params.ValidatorAddress)

	// res, err := codec.MarshalJSONIndent(keeper.cdc, delegators)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())

	// }

	return nil, nil
}

func queryReward(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	// var params QueryRewardParams
	// err := ModuleCdc.UnmarshalJSON(req.Data, &params)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONUnmarshal, err.Error())
	// }

	// reward, found := keeper.GetReward(ctx, params.EthAddress)
	// if !found {
	// 	return nil, fmt.Errorf("%w for reward of %s", common.ErrRecordNotFound, params.EthAddress)
	// }

	// res, err := codec.MarshalJSONIndent(keeper.cdc, reward)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())

	// }

	return nil, nil
}

func queryRewardEpoch(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	// epoch := keeper.GetRewardEpoch(ctx)
	// res, err := codec.MarshalJSONIndent(keeper.cdc, epoch)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	// }

	return nil, nil
}

func queryRewardStats(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	// stats := keeper.GetRewardStats(ctx)
	// res, err := codec.MarshalJSONIndent(keeper.cdc, stats)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	// }

	return nil, nil
}

func queryParameters(ctx sdk.Context, k Keeper) ([]byte, error) {
	// params := k.GetParams(ctx)

	// res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	// if err != nil {
	// 	return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	// }

	return nil, nil
}
