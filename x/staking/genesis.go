package validator

import (
	"github.com/celer-network/sgn-v2/x/staking/keeper"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) []abci.ValidatorUpdate {
	keeper.SetParams(ctx, data.Params)
	if data.GetSyncer().GetSgnAddress() != "" {
		keeper.SetSyncer(ctx, data.Syncer)
	}

	for _, validator := range data.Validators {
		keeper.SetValidatorStates(ctx, &validator)
	}

	for _, delegator := range data.Delegators {
		keeper.SetDelegator(ctx, &delegator)
	}

	return keeper.GetValidatorPowerUpdates(ctx)
}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	syncer := keeper.GetSyncer(ctx)
	validators := keeper.GetAllValidators(ctx)
	delegators := []types.Delegator{}

	for _, validator := range validators {
		delegators = append(delegators, keeper.GetAllDelegators(ctx, validator.EthAddress)...)
	}

	vals := make([]types.Validator, 0)
	for i := range validators {
		vals = append(vals, validators[i])
	}
	return &types.GenesisState{
		Params:     params,
		Syncer:     syncer,
		Validators: vals,
		Delegators: delegators,
	}
}

func ValidateGenesis(data *types.GenesisState) error {
	return data.Params.Validate()
}
