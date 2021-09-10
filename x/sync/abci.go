package sync

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/celer-network/sgn-v2/x/sync/keeper"
	"github.com/celer-network/sgn-v2/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	vals := keeper.GetBondedValidators(ctx)
	tokens := sdk.ZeroInt()
	valMaps := map[string]stakingtypes.Validator{}

	for _, val := range vals {
		tokens = tokens.Add(val.Tokens)
		valMaps[val.SgnAddress] = val
	}

	// TODO: better float to Dec conversion
	thresholdRatio, _ := sdk.NewDecFromStr(fmt.Sprintf("%f", keeper.GetParams(ctx).TallyThreshold))
	threshold := thresholdRatio.MulInt(tokens).TruncateInt()

	updates := keeper.GetAllPendingUpdates(ctx)
	for _, update := range updates {
		yesVotes := sdk.ZeroInt()
		for _, vote := range update.Votes {
			v, ok := valMaps[vote.Voter]
			if !ok {
				continue
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
