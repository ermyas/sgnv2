package types

import (
	"fmt"
	"strings"

	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ProposalTypeAddPool defines the type for an AddPoolProposal
	ProposalTypeAddPool = "AddPool"
)

// Assert AddPoolProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &AddPoolProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddPool)
	govtypes.RegisterProposalTypeCodec(AddPoolProposal{}, "sgn-v2/farming/AddPoolProposal")
}

// NewAddPoolProposal creates a new instance of AddPoolProposal
//nolint:interfacer
func NewAddPoolProposal(
	title, description,
	poolName string,
	stakeToken ERC20Token,
	rewardTokens []ERC20Token,
	initialRewardInputs []RewardAdjustmentInput,
) *AddPoolProposal {
	return &AddPoolProposal{
		Title:               title,
		Description:         description,
		PoolName:            poolName,
		StakeToken:          stakeToken,
		RewardTokens:        rewardTokens,
		InitialRewardInputs: initialRewardInputs,
	}
}

// GetTitle returns title of an AddPoolProposal object
func (ap AddPoolProposal) GetTitle() string {
	return ap.Title
}

// GetDescription returns description of an AddPoolProposal object
func (ap AddPoolProposal) GetDescription() string {
	return ap.Description
}

// ProposalRoute returns route key of an AddPoolProposal object
func (ap AddPoolProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns type of an AddPoolProposal object
func (ap AddPoolProposal) ProposalType() string {
	return ProposalTypeAddPool
}

// ValidateBasic validates an AddPoolProposal
func (ap AddPoolProposal) ValidateBasic() error {
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
	if ap.ProposalType() != ProposalTypeAddPool {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalType, ap.ProposalType())
	}
	if len(ap.PoolName) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "pool name is required")
	}
	if ap.StakeToken.Symbol == "" {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "stake token symbol is required")
	}
	for _, rewardToken := range ap.RewardTokens {
		if rewardToken.Symbol == "" {
			return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "reward token symbol is required")
		}
	}
	for _, rewardInput := range ap.InitialRewardInputs {
		if rewardInput.AddAmount.Denom == "" {
			return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "reward input denom is required")
		}
	}
	return nil
}

// String returns a human readable string representation of an AddPoolProposal
func (ap AddPoolProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`AddPoolProposal:
 Title:					%s
 Description:        	%s
 Type:                	%s
 PoolName:				%s
 StakeToken:			%v
 RewardTokens:			%v
 InitialRewardInputs:	%v
`, ap.Title, ap.Description, ap.ProposalType(), ap.PoolName, ap.StakeToken, ap.RewardTokens, ap.InitialRewardInputs))
	return b.String()
}
