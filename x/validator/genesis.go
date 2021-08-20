package validator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Params     Params      `json:"params" yaml:"params"`
	Syncer     Syncer      `json:"syncer" yaml:"syncer"`
	Validators []Validator `json:"validators" yaml:"validators"`
	Delegators []Delegator `json:"delegators" yaml:"delegators"`
	Rewards    []Reward    `json:"rewards" yaml:"rewards"`
}

func NewGenesisState(params Params) GenesisState {
	return GenesisState{
		Params: params,
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
	if !data.Syncer.SgnAddress.Empty() {
		keeper.SetSyncer(ctx, data.Syncer)
	}

	for _, validator := range data.Validators {
		keeper.SetValidator(ctx, validator)
	}

	for _, delegator := range data.Delegators {
		keeper.SetDelegator(ctx, delegator)
	}

	for _, reward := range data.Rewards {
		keeper.SetReward(ctx, reward)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	params := keeper.GetParams(ctx)
	syncer := keeper.GetSyncer(ctx)
	validators := keeper.GetAllValidators(ctx)
	delegators := []Delegator{}
	rewards := keeper.GetRewards(ctx)

	for _, validator := range validators {
		delegators = append(delegators, keeper.GetAllDelegators(ctx, validator.EthAddress)...)
	}

	return GenesisState{
		Params:     params,
		Syncer:     syncer,
		Validators: validators,
		Delegators: delegators,
		Rewards:    rewards,
	}
}
