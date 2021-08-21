package validator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewGenesisState(params Params) GenesisState {
	return GenesisState{
		Params: &params,
	}
}

func ValidateGenesis(data GenesisState) error {
	return data.Params.Validate()
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState(DefaultParams())
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	keeper.SetParams(ctx, data.Params)
	if data.Syncer.SgnAddress != "" {
		keeper.SetSyncer(ctx, data.Syncer)
	}

	for _, validator := range data.Validators {
		keeper.SetValidator(ctx, validator)
	}

	for _, delegator := range data.Delegators {
		keeper.SetDelegator(ctx, delegator)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	params := keeper.GetParams(ctx)
	syncer := keeper.GetSyncer(ctx)
	validators := keeper.GetAllValidators(ctx)
	delegators := []*Delegator{}

	for _, validator := range validators {
		delegators = append(delegators, keeper.GetAllDelegators(ctx, validator.EthAddress)...)
	}

	return GenesisState{
		Params:     &params,
		Syncer:     &syncer,
		Validators: validators,
		Delegators: delegators,
	}
}
