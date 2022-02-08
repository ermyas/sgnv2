package validator

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/keeper"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// EndBlocker called every block, update validator set, distribute rewards
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) (updates []abci.ValidatorUpdate) {
	setSyncer(ctx, keeper)

	return keeper.TmValidatorUpdates(ctx)
}

// Update syncer for every syncerDuration
func setSyncer(ctx sdk.Context, keeper keeper.Keeper) {
	syncerDuration := keeper.SyncerDuration(ctx)
	if uint64(ctx.BlockHeight())%syncerDuration != 0 {
		return
	}
	syncer := keeper.GetSyncer(ctx)
	candidates := keeper.SyncerCandidates(ctx)
	if len(candidates) != 0 {
		candidates = filterNotBondedCandidates(ctx, keeper, candidates)
		if len(candidates) == 0 {
			log.Error("No valid bonded candidates configured!!!")
			return
		}
	} else {
		validators := keeper.GetBondedValidators(ctx)
		for _, val := range validators {
			candidates = append(candidates, val.EthAddress)
		}
	}

	vIdx := uint64(ctx.BlockHeight()) / syncerDuration % uint64(len(candidates))
	if syncer.ValIndex != vIdx || syncer.EthAddress == "" || syncer.EthAddress != candidates[vIdx] {
		syncer = types.NewSyncer(vIdx, candidates[vIdx])
		keeper.SetSyncer(ctx, syncer)
		log.Infof("set syncer to %s", syncer.EthAddress)
	}
}

func filterNotBondedCandidates(ctx sdk.Context, keeper keeper.Keeper, candidates []string) (filtered []string) {
	for _, candidate := range candidates {
		val, found := keeper.GetValidator(ctx, eth.Hex2Addr(candidate))
		if found && val.IsBonded() {
			filtered = append(filtered, eth.FormatAddrHex(candidate))
		}
	}
	return
}
