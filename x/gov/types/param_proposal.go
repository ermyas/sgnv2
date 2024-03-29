package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)

const (
	// ProposalTypeChange defines the type for a ParameterProposal
	ProposalTypeChange = "ParameterChange"
)

// Assert ParameterProposal implements Content at compile-time
var _ Content = &ParameterProposal{}

func init() {
	RegisterProposalType(ProposalTypeChange)
}

func NewParameterProposal(title, description string, changes []ParamChange) *ParameterProposal {
	return &ParameterProposal{title, description, changes}
}

// GetTitle returns the title of a parameter change proposal.
func (pcp *ParameterProposal) GetTitle() string { return pcp.Title }

// GetDescription returns the description of a parameter change proposal.
func (pcp *ParameterProposal) GetDescription() string { return pcp.Description }

// ProposalRoute returns the routing key of a parameter change proposal.
func (pcp *ParameterProposal) ProposalRoute() string { return proposal.RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (pcp *ParameterProposal) ProposalType() string { return ProposalTypeChange }

// ValidateBasic validates the parameter change proposal
func (pcp *ParameterProposal) ValidateBasic() error {
	err := ValidateAbstract(pcp)
	if err != nil {
		return err
	}

	return ValidateChanges(pcp.Changes)
}

// String implements the Stringer interface.
func (pcp *ParameterProposal) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf(`Parameter Change Proposal:
  Title:       %s
  Description: %s
  Changes:
`, pcp.Title, pcp.Description))

	for _, pc := range pcp.Changes {
		b.WriteString(fmt.Sprintf(`    Param Change:
      Subspace: %s
      Key:      %s
      Value:    %X
`, pc.Subspace, pc.Key, pc.Value))
	}

	return b.String()
}

func NewParamChange(subspace, key, value string) ParamChange {
	return ParamChange{subspace, key, value}
}

// String implements the Stringer interface.
func (pc ParamChange) String() string {
	return fmt.Sprintf(`Param Change:
  Subspace: %s
  Key:      %s
  Value:    %X
`, pc.Subspace, pc.Key, pc.Value)
}

// ValidateChanges performs basic validation checks over a set of ParamChange. It
// returns an error if any ParamChange is invalid.
func ValidateChanges(changes []ParamChange) error {
	if len(changes) == 0 {
		return proposal.ErrEmptyChanges
	}

	for _, pc := range changes {
		if len(pc.Subspace) == 0 {
			return proposal.ErrEmptySubspace
		}
		if len(pc.Key) == 0 {
			return proposal.ErrEmptyKey
		}
		if len(pc.Value) == 0 {
			return proposal.ErrEmptyValue
		}
	}

	return nil
}
