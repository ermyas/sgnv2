package mint

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/celer-network/sgn-v2/x/mint/keeper"
	"github.com/celer-network/sgn-v2/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewProposalHandler returns a gov handler for x/farming
func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) (err error) {
		switch c := content.(type) {
		case *types.AdjustProvisionsProposal:
			err = keeper.HandleAdjustProvisionsProposal(ctx, k, c)
			return err
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized mint proposal content type: %T", c)
		}
	}
}
