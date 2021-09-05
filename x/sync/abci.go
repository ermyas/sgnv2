package sync

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/sync/keeper"
	"github.com/celer-network/sgn-v2/x/sync/types"
	valtypes "github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	sdkVals := keeper.GetBondedValidators(ctx)
	tokens := sdk.ZeroInt()
	sdkValMaps := map[string]sdk_staking.Validator{}

	for _, val := range sdkVals {
		tokens = tokens.Add(val.Tokens)
		sdkValMaps[val.OperatorAddress] = val
	}

	// TODO: better float to Dec conversion
	thresholdRatio, _ := sdk.NewDecFromStr(fmt.Sprintf("%f", keeper.GetParams(ctx).TallyThreshold))
	threshold := thresholdRatio.MulInt(tokens).TruncateInt()

	updates := keeper.GetAllPendingUpdates(ctx)
	for _, update := range updates {
		yesVotes := sdk.ZeroInt()
		for _, vote := range update.Votes {
			v, ok := sdkValMaps[vote.Voter]
			if !ok {
				vaddr, _ := valtypes.SdkValAddrFromSgnBech32(vote.Voter)
				v, ok = sdkValMaps[vaddr.String()]
				if !ok {
					continue
				}
			}
			if vote.Option == types.VoteOption_Yes {
				yesVotes = yesVotes.Add(v.Tokens)
			}
		}

		if yesVotes.GT(threshold) {
			log.Infof("Update approved by majority. id: %d, type: %s, votes: %s, threshold %s",
				update.Id, update.Type, yesVotes, threshold)
			keeper.ApplyUpdate(ctx, update)
			keeper.RemovePendingUpdate(ctx, update.Id)
		} else if ctx.BlockTime().Unix() > int64(update.ClosingTs) {
			log.Debugf("Pending update expired, id: %d, type: %s, votes: %s, threshold %s",
				update.Id, update.Type, yesVotes, threshold)
			keeper.RemovePendingUpdate(ctx, update.Id)
		}
	}

}
