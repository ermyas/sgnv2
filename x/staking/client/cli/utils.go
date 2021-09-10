package cli

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
)

func QueryValidator(cliCtx client.Context, ethAddress string) (validator *types.Validator, err error) {
	params := types.NewQueryValidatorParams(ethAddress)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return
	}
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidator)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	validator = new(types.Validator)
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, validator)
	return
}

func QueryValidatorBySgnAddr(cliCtx client.Context, sgnAddress string) (validator *types.Validator, err error) {
	params := types.NewQueryValidatorBySgnAddrParams(sgnAddress)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return
	}
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidatorBySgnAddr)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	validator = new(types.Validator)
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, validator)
	return
}

func QueryValidatorByConsAddr(cliCtx client.Context, consAddress string) (validator *types.Validator, err error) {
	params := types.NewQueryValidatorByConsAddrParams(consAddress)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return
	}
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidatorByConsAddr)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	validator = new(types.Validator)
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, validator)
	return
}

func QueryValidators(cliCtx client.Context) (validators types.Validators, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidators)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &validators)
	sort.SliceStable(validators, func(i, j int) bool {
		return validators[i].Tokens.GT(validators[j].Tokens)
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
	delegator = new(types.Delegator)
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, delegator)
	return
}

func QueryDelegators(cliCtx client.Context, ethAddress string) (delegators []*types.Delegator, err error) {
	params := types.NewQueryDelegatorsParams(ethAddress)
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

func QuerySgnAccount(cliCtx client.Context, sgnAddr string) (exist bool, err error) {
	params := types.NewQuerySgnAccountParams(sgnAddr)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return false, err
	}
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySgnAccount)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	exist = bytes.Compare(res, []byte{1}) == 0
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
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &params)
	return
}
