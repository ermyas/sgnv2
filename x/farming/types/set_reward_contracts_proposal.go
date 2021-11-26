package types

import (
	"fmt"
	"strings"

	commontypes "github.com/celer-network/sgn-v2/common/types"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ProposalTypeSetRewardContracts defines the type for an SetRewardContractsProposal
	ProposalTypeSetRewardContracts = "SetRewardContracts"
)

// Assert SetRewardContractsProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &SetRewardContractsProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeSetRewardContracts)
	govtypes.RegisterProposalTypeCodec(SetRewardContractsProposal{}, "sgn-v2/SetRewardContractsProposal")
}

// NewSetRewardContractsProposal creates a new instance of SetRewardContractsProposal
//nolint:interfacer
func NewSetRewardContractsProposal(
	title, description string, rewardContracts []commontypes.ContractInfo) *SetRewardContractsProposal {
	return &SetRewardContractsProposal{
		Title:           title,
		Description:     description,
		RewardContracts: rewardContracts,
	}
}

// GetTitle returns title of an SetRewardContractsProposal object
func (p SetRewardContractsProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns description of an SetRewardContractsProposal object
func (p SetRewardContractsProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns route key of an SetRewardContractsProposal object
func (p SetRewardContractsProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns type of an SetRewardContractsProposal object
func (p SetRewardContractsProposal) ProposalType() string {
	return ProposalTypeSetRewardContracts
}

// ValidateBasic validates an SetRewardContractsProposal
func (p SetRewardContractsProposal) ValidateBasic() error {
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
	if p.ProposalType() != ProposalTypeSetRewardContracts {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalType, p.ProposalType())
	}
	for _, contract := range p.RewardContracts {
		if contract.Address == "" {
			return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "contract address is empty")
		}
	}
	return nil
}

// String returns a human readable string representation of an SetRewardContractsProposal
func (p SetRewardContractsProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`SetRewardContractsProposal:
 Title:				%s
 Description:		%s
 Type:      		%s
 RewardContracts:			%v
`, p.Title, p.Description, p.ProposalType(), p.RewardContracts))
	return b.String()
}
