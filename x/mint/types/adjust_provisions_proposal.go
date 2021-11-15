package types

import (
	"fmt"
	"strings"

	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ProposalTypeAdjustProvisions defines the type for an AdjustProvisionsProposal
	ProposalTypeAdjustProvisions = "AdjustProvisions"
)

// Assert AdjustProvisionsProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &AdjustProvisionsProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeAdjustProvisions)
	govtypes.RegisterProposalTypeCodec(AdjustProvisionsProposal{}, "sgn-v2/AdjustProvisionsProposal")
}

// NewAdjustProvisionsProposal creates a new instance of AdjustProvisionsProposal
//nolint:interfacer
func NewAdjustProvisionsProposal(
	title, description string,
	newAnnualProvisions sdk.Dec) *AdjustProvisionsProposal {
	return &AdjustProvisionsProposal{
		Title:               title,
		Description:         description,
		NewAnnualProvisions: newAnnualProvisions,
	}
}

// GetTitle returns title of an AdjustProvisionsProposal object
func (p AdjustProvisionsProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns description of an AdjustProvisionsProposal object
func (p AdjustProvisionsProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns route key of an AdjustProvisionsProposal object
func (p AdjustProvisionsProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns type of an AdjustProvisionsProposal object
func (p AdjustProvisionsProposal) ProposalType() string {
	return ProposalTypeAdjustProvisions
}

// ValidateBasic validates an AdjustProvisionsProposal
func (p AdjustProvisionsProposal) ValidateBasic() error {
	if len(strings.TrimSpace(p.Title)) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "title is required")
	}
	if len(p.Title) > govtypes.MaxTitleLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "title length is longer than the maximum title length")
	}
	if len(p.Description) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "description is required")
	}
	if len(p.Description) > govtypes.MaxDescriptionLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "description length is longer than the maximum description length")
	}
	if p.ProposalType() != ProposalTypeAdjustProvisions {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalType, p.ProposalType())
	}
	if p.NewAnnualProvisions.IsNegative() {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "annual provisions cannot be negative")
	}
	return nil
}

// String returns a human readable string representation of an AdjustProvisionsProposal
func (p AdjustProvisionsProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`AdjustProvisionsProposal:
 Title:						%s
 Description:       		%s
 Type:              		%s
 NewAnnualProvisions:       %v
`, p.Title, p.Description, p.ProposalType(), p.NewAnnualProvisions))
	return b.String()
}
