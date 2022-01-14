package farming

import (
	"github.com/celer-network/sgn-v2/x/farming/keeper"
	"github.com/celer-network/sgn-v2/x/farming/types"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewProposalHandler returns a gov handler for x/farming
func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) (err error) {
		switch c := content.(type) {
		case *types.AddPoolProposal:
			err = keeper.HandleAddPoolProposal(ctx, k, c)
			return err
		case *types.BatchAddPoolProposal:
			err = keeper.HandleBatchAddPoolProposal(ctx, k, c)
			return err
		case *types.AdjustRewardProposal:
			err = keeper.HandleAdjustRewardProposal(ctx, k, c)
			return err
		case *types.BatchAdjustRewardProposal:
			err = keeper.HandleBatchAdjustRewardProposal(ctx, k, c)
			return err
		case *types.AddTokensProposal:
			err = keeper.HandleAddTokensProposal(ctx, k, c)
			return err
		case *types.SetRewardContractsProposal:
			err = keeper.HandleSetRewardContractsProposal(ctx, k, c)
			return err
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized farming proposal content type: %T", c)
		}
	}
}
