package types

import (
	"fmt"
	"strings"

	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ProposalTypeRemovePool defines the type for a RemovePoolProposal
	ProposalTypeRemovePool = "RemovePool"
)

// Assert RemovePoolProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &RemovePoolProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeRemovePool)
	govtypes.RegisterProposalTypeCodec(RemovePoolProposal{}, "sgn-v2/farming/RemovePoolProposal")
}

// NewRemovePoolProposal creates a new instance of RemovePoolProposal
//nolint:interfacer
func NewRemovePoolProposal(
	title, description, poolName string) *RemovePoolProposal {
	return &RemovePoolProposal{
		Title:       title,
		Description: description,
		PoolName:    poolName,
	}
}

// GetTitle returns title of a RemovePoolProposal object
func (rp RemovePoolProposal) GetTitle() string {
	return rp.Title
}

// GetDescription returns description of a RemovePoolProposal object
func (rp RemovePoolProposal) GetDescription() string {
	return rp.Description
}

// ProposalRoute returns route key of a RemovePoolProposal object
func (rp RemovePoolProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns type of a RemovePoolProposal object
func (rp RemovePoolProposal) ProposalType() string {
	return ProposalTypeRemovePool
}

// ValidateBasic validates a RemovePoolProposal
func (rp RemovePoolProposal) ValidateBasic() error {
	if len(strings.TrimSpace(rp.Title)) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "title is required")
	}
	if len(rp.Title) > govtypes.MaxTitleLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "title length is longer than the maximum title length")
	}
	if len(rp.Description) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "description is required")
	}
	if len(rp.Description) > govtypes.MaxDescriptionLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "description length is longer than the maximum description length")
	}
	if rp.ProposalType() != ProposalTypeRemovePool {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalType, rp.ProposalType())
	}
	if len(rp.PoolName) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "pool name is required")
	}
	return nil
}

// String returns a human readable string representation of a RemovePoolProposal
func (rp RemovePoolProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`RemovePoolProposal:
 Title:					%s
 Description:        	%s
 Type:                	%s
 PoolName:				%s
`, rp.Title, rp.Description, rp.ProposalType(), rp.PoolName))
	return b.String()
}
