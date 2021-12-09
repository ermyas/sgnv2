package pegbridge

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	pegkeeper "github.com/celer-network/sgn-v2/x/pegbridge/keeper"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewPegProposalHandler(k pegkeeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.PegProposal:
			return handlePegProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unsupported cbr proposal content type: %T", c)
		}
	}
}

func handlePegProposal(ctx sdk.Context, k pegkeeper.Keeper, p *types.PegProposal) error {
	if err := p.PegConfig.Validate(); err != nil {
		return err
	}
	k.SetPegConfig(ctx, *p.PegConfig)
	return nil
}
