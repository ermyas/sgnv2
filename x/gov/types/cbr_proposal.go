package types

import (
	"encoding/json"
	"fmt"

	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
)

const (
	ProposalTypeCbridge = "CbrParamChange"
)

func init() {
	RegisterProposalType(ProposalTypeCbridge)
}

var _ Content = &CbrProposal{}

func (cp *CbrProposal) GetTitle() string { return cp.Title }

func (cp *CbrProposal) GetDescription() string { return cp.Description }

// ProposalRoute returns the routing key of a cbr proposal.
func (cp *CbrProposal) ProposalRoute() string { return cbrtypes.RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (cp *CbrProposal) ProposalType() string { return ProposalTypeCbridge }

// ValidateBasic validates the parameter change proposal
func (cp *CbrProposal) ValidateBasic() error {
	// todo: validate cp.CbrConfig?
	return nil
}

// String implements the Stringer interface.
func (cp *CbrProposal) String() string {
	jsonraw, _ := json.MarshalIndent(cp.CbrConfig, "", "  ")
	return fmt.Sprintf(`
Cbrige Config Proposal:
  Title: %s
  Description: %s
  New Config: %s`, cp.Title, cp.Description, string(jsonraw))
}
