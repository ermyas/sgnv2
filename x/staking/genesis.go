package validator

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/keeper"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) []abci.ValidatorUpdate {
	keeper.SetParams(ctx, data.Params)
	if data.GetSyncer().GetEthAddress() != "" {
		keeper.SetSyncer(ctx, data.Syncer)
	}

	for _, validator := range data.Validators {
		keeper.SetValidatorParams(ctx, &validator, true)
		keeper.SetValidatorStates(ctx, &validator)
	}

	for _, delegation := range data.Delegations {
		keeper.SetDelegation(ctx, delegation)
	}

	return keeper.TmValidatorUpdates(ctx)
}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	syncer := keeper.GetSyncer(ctx)
	validators := keeper.GetAllValidators(ctx)
	delegations := []types.Delegation{}

	for _, validator := range validators {
		delegations = append(
			delegations,
			keeper.GetAllDelegations(ctx, eth.Hex2Addr(validator.EthAddress))...)
	}

	vals := make([]types.Validator, 0)
	for i := range validators {
		vals = append(vals, validators[i])
	}
	return &types.GenesisState{
		Params:      params,
		Syncer:      syncer,
		Validators:  vals,
		Delegations: delegations,
	}
}

func ValidateGenesis(data *types.GenesisState) error {
	return data.Params.Validate()
}
