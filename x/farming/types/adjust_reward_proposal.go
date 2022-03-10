package types

import (
	"fmt"
	"strings"

	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ProposalTypeAdjustReward defines the type for an AdjustRewardProposal
	ProposalTypeAdjustReward = "AdjustReward"
)

// Assert AdjustRewardProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &AdjustRewardProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeAdjustReward)
	govtypes.RegisterProposalTypeCodec(AdjustRewardProposal{}, "sgn-v2/AdjustRewardProposal")
}

// NewAdjustRewardProposal creates a new instance of AdjustRewardProposal
//nolint:interfacer
func NewAdjustRewardProposal(
	title, description, poolName string,
	rewardAdjustmentInputs []RewardAdjustmentInput, removeDuplicates bool) *AdjustRewardProposal {
	return &AdjustRewardProposal{
		Title:                  title,
		Description:            description,
		PoolName:               poolName,
		RewardAdjustmentInputs: rewardAdjustmentInputs,
		RemoveDuplicates:       removeDuplicates,
	}
}

// GetTitle returns title of an AdjustRewardProposal object
func (arp AdjustRewardProposal) GetTitle() string {
	return arp.Title
}

// GetDescription returns description of an AdjustRewardProposal object
func (arp AdjustRewardProposal) GetDescription() string {
	return arp.Description
}

// ProposalRoute returns route key of an AdjustRewardProposal object
func (arp AdjustRewardProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns type of an AdjustRewardProposal object
func (arp AdjustRewardProposal) ProposalType() string {
	return ProposalTypeAdjustReward
}

// ValidateBasic validates an AdjustRewardProposal
func (arp AdjustRewardProposal) ValidateBasic() error {
	if len(strings.TrimSpace(arp.Title)) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "title is required")
	}
	if len(arp.Title) > govtypes.MaxTitleLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "title length is longer than the maximum title length")
	}
	if len(arp.Description) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "description is required")
	}
	if len(arp.Description) > govtypes.MaxDescriptionLength {
		return sdkerrors.Wrap(
			govtypes.ErrInvalidProposalContent, "description length is longer than the maximum description length")
	}
	if arp.ProposalType() != ProposalTypeAdjustReward {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalType, arp.ProposalType())
	}
	adjustRewardInfo := arp.GetAdjustRewardInfo()
	return adjustRewardInfo.ValidateBasic()
}

// String returns a human readable string representation of an AdjustRewardProposal
func (arp AdjustRewardProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`AdjustRewardProposal:
 Title:						%s
 Description:       		%s
 Type:              		%s
 PoolName:					%s
 RewardAdjustmentInputs:	%v
 RemoveDuplicates:          %t
`, arp.Title, arp.Description, arp.ProposalType(), arp.PoolName, arp.RewardAdjustmentInputs, arp.RemoveDuplicates))
	return b.String()
}

func (arp AdjustRewardProposal) GetAdjustRewardInfo() AdjustRewardInfo {
	return AdjustRewardInfo{
		PoolName:               arp.PoolName,
		RewardAdjustmentInputs: arp.RewardAdjustmentInputs,
		RemoveDuplicates:       arp.RemoveDuplicates,
	}
}

func (arInfo AdjustRewardInfo) ValidateBasic() error {
	if len(arInfo.PoolName) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "pool name is required")
	}
	for _, rewardAdjustmentInput := range arInfo.RewardAdjustmentInputs {
		if rewardAdjustmentInput.AddAmount.Denom == "" {
			return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "reward token symbol is required")
		}
	}
	return nil
}
