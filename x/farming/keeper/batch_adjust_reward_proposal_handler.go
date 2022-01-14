package keeper

import (
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandleBatchAdjustRewardProposal is a handler for executing a passed BatchAdjustRewardProposal
func HandleBatchAdjustRewardProposal(ctx sdk.Context, k Keeper, p *types.BatchAdjustRewardProposal) error {
	for _, info := range p.AdjustRewardInfos {
		err := handleAdjustRewardProposalByAdjustRewardInfo(ctx, k, &info)
		if err != nil {
			return err
		}
	}
	return nil
}
