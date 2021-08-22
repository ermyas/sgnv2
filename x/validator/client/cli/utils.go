package cli

import (
	"fmt"
	"sort"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func QueryValidator(cliCtx client.Context, ethAddress string) (validator *types.Validator, err error) {
	params := types.NewQueryValidatorParams(ethAddress)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidator)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, validator)
	return
}

func QueryValidators(cliCtx client.Context) (validators []*types.Validator, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidators)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &validators)
	sort.SliceStable(validators, func(i, j int) bool {
		return validators[i].Tokens > validators[j].Tokens
	})
	return
}

func QueryDelegator(cliCtx client.Context, valAddr, delAddr string) (delegator *types.Delegator, err error) {
	params := types.NewQueryDelegatorParams(valAddr, delAddr)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegator)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, delegator)
	return
}

func QueryDelegators(cliCtx client.Context, ethAddress string) (delegators []*types.Delegator, err error) {
	params := types.NewQueryValidatorParams(ethAddress)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegators)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, delegators)
	return
}

// addrStr should be bech32 sgn account address with prefix sgn
func QuerySdkValidator(cliCtx client.Context, sgnAddr string) (sdkval *sdk_staking.Validator, err error) {
	addr, err := sdk.AccAddressFromBech32(sgnAddr)
	params := sdk_staking.NewQueryValidatorParams(sdk.ValAddress(addr), 0, 0)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	route := fmt.Sprintf("custom/%s/%s", sdk_staking.QuerierRoute, sdk_staking.QueryValidator)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, sdkval)
	return
}

// Query validators of given status in sdk staking module
func QuerySdkValidators(cliCtx client.Context, status string) (sdkvals sdk_staking.Validators, err error) {
	sdkstatus := sdk_staking.BondStatusUnspecified
	if status == "bonded" {
		sdkstatus = sdk_staking.BondStatusBonded
	} else if status == "unbonding" {
		sdkstatus = sdk_staking.BondStatusUnbonding
	} else if status == "unbonded" {
		sdkstatus = sdk_staking.BondStatusUnbonded
	} else {
		err = fmt.Errorf("Invalid status, need to be bonded || unbonded || unbonding ")
		return
	}
	params := sdk_staking.NewQueryValidatorsParams(1, 100, sdkstatus)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	route := fmt.Sprintf("custom/%s/%s", sdk_staking.QuerierRoute, sdk_staking.QueryValidators)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, sdkvals)
	return
}

// Query syncer info
func QuerySyncer(cliCtx client.Context) (syncer types.Syncer, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySyncer)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &syncer)
	return
}

// Query params info
func QueryParams(cliCtx client.Context) (params types.Params, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParameters)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &params)
	return
}
