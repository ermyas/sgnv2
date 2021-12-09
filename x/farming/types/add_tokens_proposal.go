package types

import (
	"fmt"
	"strings"

	commontypes "github.com/celer-network/sgn-v2/common/types"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ProposalTypeAddTokens defines the type for an AddTokensProposal
	ProposalTypeAddTokens = "AddTokens"
)

// Assert AddTokensProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &AddTokensProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddTokens)
	govtypes.RegisterProposalTypeCodec(AddTokensProposal{}, "sgn-v2/AddTokensProposal")
}

// NewAddTokensProposal creates a new instance of AddTokensProposal
//nolint:interfacer
func NewAddTokensProposal(
	title, description string, tokens []commontypes.ERC20Token) *AddTokensProposal {
	return &AddTokensProposal{
		Title:       title,
		Description: description,
		Tokens:      tokens,
	}
}

// GetTitle returns title of an AddTokensProposal object
func (ap AddTokensProposal) GetTitle() string {
	return ap.Title
}

// GetDescription returns description of an AddTokensProposal object
func (ap AddTokensProposal) GetDescription() string {
	return ap.Description
}

// ProposalRoute returns route key of an AddTokensProposal object
func (ap AddTokensProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns type of an AddTokensProposal object
func (ap AddTokensProposal) ProposalType() string {
	return ProposalTypeAddTokens
}

// ValidateBasic validates an AddTokensProposal
func (ap AddTokensProposal) ValidateBasic() error {
	if len(strings.TrimSpace(ap.Title)) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "title is required")
	}
	if len(ap.Title) > govtypes.MaxTitleLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "title length is longer than the maximum title length")
	}
	if len(ap.Description) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "description is required")
	}
	if len(ap.Description) > govtypes.MaxDescriptionLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "description length is longer than the maximum description length")
	}
	if ap.ProposalType() != ProposalTypeAddTokens {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalType, ap.ProposalType())
	}
	for _, token := range ap.Tokens {
		if token.Symbol == "" {
			return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "token symbol is required")
		}
	}
	return nil
}

// String returns a human readable string representation of an AddTokensProposal
func (ap AddTokensProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`AddTokensProposal:
 Title:				%s
 Description:		%s
 Type:      		%s
 Tokens:			%v
`, ap.Title, ap.Description, ap.ProposalType(), ap.Tokens))
	return b.String()
}
