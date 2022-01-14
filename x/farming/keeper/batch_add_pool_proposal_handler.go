package keeper

import (
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandleBatchAddPoolProposal is a handler for executing a passed BatchAddPoolProposal
func HandleBatchAddPoolProposal(ctx sdk.Context, k Keeper, p *types.BatchAddPoolProposal) error {
	if err := k.CheckBatchAddPoolProposal(ctx, p); err != nil {
		return err
	}
	for _, info := range p.AddPoolInfos {
		// 1. Create stake token if not existent
		err := handleAddPoolProposalByAddPoolInfo(ctx, k, &info)
		if err != nil {
			return err
		}
	}
	return nil
}

// CheckAddPoolProposal checks the validity of an BatchAddPoolProposal
func (k Keeper) CheckBatchAddPoolProposal(ctx sdk.Context, p *types.BatchAddPoolProposal) error {
	for _, info := range p.AddPoolInfos {
		err := k.CheckAddPoolProposal(ctx, &info)
		if err != nil {
			return err
		}
	}
	return nil
}
