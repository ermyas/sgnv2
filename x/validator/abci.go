package validator

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/validator/keeper"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// EndBlocker called every block, update validator set, distribute rewards
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) (updates []abci.ValidatorUpdate) {
	setSyncer(ctx, keeper)

	return keeper.BlockValidatorUpdates(ctx)
}

// Update syncer for every syncerDuration
func setSyncer(ctx sdk.Context, keeper keeper.Keeper) {
	syncer := keeper.GetSyncer(ctx)
	validators := keeper.GetBondedSgnValidators(ctx)
	syncerDuration := keeper.SyncerDuration(ctx)
	vIdx := uint64(ctx.BlockHeight()) / syncerDuration % uint64(len(validators))

	if syncer.ValIndex != vIdx || syncer.SgnAddress == "" {
		syncer = types.NewSyncer(vIdx, validators[vIdx].OperatorAddress)
		keeper.SetSyncer(ctx, syncer)
		log.Infof("set syncer to %s", syncer.SgnAddress)
	}
}
