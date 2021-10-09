package cbridge

import (
	cbrkeeper "github.com/celer-network/sgn-v2/x/cbridge/keeper"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewCbrProposalHandler(k cbrkeeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CbrProposal:
			return handleCbrProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unsupported cbr proposal content type: %T", c)
		}
	}
}

func handleCbrProposal(ctx sdk.Context, k cbrkeeper.Keeper, p *types.CbrProposal) error {
	if err := p.CbrConfig.Validate(); err != nil {
		return err
	}
	k.SetCbrConfig(ctx, *p.CbrConfig)
	return nil
}
