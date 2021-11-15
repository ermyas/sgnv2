package keeper

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/celer-network/sgn-v2/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// HandleAdjustProvisionsProposal is a handler for executing a passed AdjustProvisionsProposal
func HandleAdjustProvisionsProposal(ctx sdk.Context, k Keeper, p *types.AdjustProvisionsProposal) error {
	if err := k.CheckAdjustProvisionsProposal(ctx, p); err != nil {
		return err
	}
	k.SetMinter(ctx, types.NewMinter(p.NewAnnualProvisions))
	return nil
}

// CheckAdjustProvisionsProposal checks the validity of an AdjustProvisionsProposal
func (k Keeper) CheckAdjustProvisionsProposal(ctx sdk.Context, p *types.AdjustProvisionsProposal) error {
	if p.NewAnnualProvisions.IsNegative() {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "annual provisions cannot be negative")
	}
	return nil
}
