package farming

import (
	"github.com/celer-network/sgn-v2/x/farming/keeper"
	"github.com/celer-network/sgn-v2/x/farming/types"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewAddPoolProposalHandler returns a gov handler for AddPoolProposal
func NewAddPoolProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) (err error) {
		switch c := content.(type) {
		case *types.AddPoolProposal:
			return keeper.HandleAddPoolProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized farming proposal content type: %T", c)
		}
	}
}
