package types

import (
	"fmt"
	"strings"

	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ProposalTypeBatchAdjustReward defines the type for an BatchAdjustRewardProposal
	ProposalTypeBatchAdjustReward = "BatchAdjustReward"
)

// Assert BatchAdjustRewardProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &BatchAdjustRewardProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeBatchAdjustReward)
}

// NewBatchAdjustRewardProposal creates a new instance of BatchAdjustRewardProposal
//nolint:interfacer
func NewBatchAdjustRewardProposal(
	title, description string,
	adjustRewardInfos []AdjustRewardInfo) *BatchAdjustRewardProposal {
	return &BatchAdjustRewardProposal{
		Title:             title,
		Description:       description,
		AdjustRewardInfos: adjustRewardInfos,
	}
}

// GetTitle returns title of an BatchAdjustRewardProposal object
func (barp BatchAdjustRewardProposal) GetTitle() string {
	return barp.Title
}

// GetDescription returns description of an BatchAdjustRewardProposal object
func (barp BatchAdjustRewardProposal) GetDescription() string {
	return barp.Description
}

// ProposalRoute returns route key of an BatchAdjustRewardProposal object
func (barp BatchAdjustRewardProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns type of an BatchAdjustRewardProposal object
func (barp BatchAdjustRewardProposal) ProposalType() string {
	return ProposalTypeBatchAdjustReward
}

// ValidateBasic validates a BatchAdjustRewardProposal
func (barp BatchAdjustRewardProposal) ValidateBasic() error {
	if len(strings.TrimSpace(barp.Title)) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "title is required")
	}
	if len(barp.Title) > govtypes.MaxTitleLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "title length is longer than the maximum title length")
	}
	if len(barp.Description) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "description is required")
	}
	if len(barp.Description) > govtypes.MaxDescriptionLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "description length is longer than the maximum description length")
	}
	if barp.ProposalType() != ProposalTypeBatchAdjustReward {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalType, barp.ProposalType())
	}
	for _, info := range barp.AdjustRewardInfos {
		err := info.ValidateBasic()
		if err != nil {
			return err
		}
	}
	return nil
}

// String returns a human readable string representation of an BatchAdjustRewardProposal
func (barp BatchAdjustRewardProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`BatchAdjustRewardProposal:
 Title:						%s
 Description:       		%s
 Type:              		%s
 AdjustRewardInfos:	%+v
`, barp.Title, barp.Description, barp.ProposalType(), barp.AdjustRewardInfos))
	return b.String()
}
