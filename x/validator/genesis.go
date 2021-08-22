package validator

import (
	"github.com/celer-network/sgn-v2/x/validator/keeper"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewGenesisState(params *types.Params) *types.GenesisState {
	return &types.GenesisState{
		Params: params,
	}
}

func ValidateGenesis(data *types.GenesisState) error {
	return data.Params.Validate()
}

func DefaultGenesisState() *types.GenesisState {
	return NewGenesisState(types.DefaultParams())
}

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) []abci.ValidatorUpdate {
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

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	syncer := keeper.GetSyncer(ctx)
	validators := keeper.GetAllValidators(ctx)
	delegators := []*types.Delegator{}

	for _, validator := range validators {
		delegators = append(delegators, keeper.GetAllDelegators(ctx, validator.EthAddress)...)
	}

	return &types.GenesisState{
		Params:     params,
		Syncer:     syncer,
		Validators: validators,
		Delegators: delegators,
	}
}
