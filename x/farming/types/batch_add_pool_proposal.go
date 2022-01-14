package types

import (
	"fmt"
	"strings"

	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ProposalTypeBatchAddPool defines the type for an BatchAddPoolProposal
	ProposalTypeBatchAddPool = "BatchAddPool"
)

// Assert BatchAddPoolProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &BatchAddPoolProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeBatchAddPool)
}

// NewBatchAddPoolProposal creates a new instance of BatchAddPoolProposal
//nolint:interfacer
func NewBatchAddPoolProposal(
	title, description string,
	addPoolInfos []AddPoolInfo,
) *BatchAddPoolProposal {
	return &BatchAddPoolProposal{
		Title:        title,
		Description:  description,
		AddPoolInfos: addPoolInfos,
	}
}

// GetTitle returns title of an BatchAddPoolProposal object
func (bap BatchAddPoolProposal) GetTitle() string {
	return bap.Title
}

// GetDescription returns description of an BatchAddPoolProposal object
func (bap BatchAddPoolProposal) GetDescription() string {
	return bap.Description
}

// ProposalRoute returns route key of an BatchAddPoolProposal object
func (bap BatchAddPoolProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns type of an BatchAddPoolProposal object
func (bap BatchAddPoolProposal) ProposalType() string {
	return ProposalTypeBatchAddPool
}

// ValidateBasic validates an BatchAddPoolProposal
func (bap BatchAddPoolProposal) ValidateBasic() error {
	if len(strings.TrimSpace(bap.Title)) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "title is required")
	}
	if len(bap.Title) > govtypes.MaxTitleLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "title length is longer than the maximum title length")
	}
	if len(bap.Description) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "description is required")
	}
	if len(bap.Description) > govtypes.MaxDescriptionLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "description length is longer than the maximum description length")
	}
	if bap.ProposalType() != ProposalTypeBatchAddPool {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalType, bap.ProposalType())
	}
	for _, info := range bap.AddPoolInfos {
		err := info.ValidateBasic()
		if err != nil {
			return err
		}
	}
	return nil
}

// String returns a human readable string representation of an BatchAddPoolProposal
func (bap BatchAddPoolProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`BatchAddPoolProposal:
 Title:					%s
 Description:        	%s
 Type:                	%s
 AddPoolInfos:				%+v
`, bap.Title, bap.Description, bap.ProposalType(), bap.AddPoolInfos))
	return b.String()
}
