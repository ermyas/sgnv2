package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
)

const (
	ProposalTypeCbridge = "CbrParamChange"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCbridge)
}

var _ govtypes.Content = &CbrProposal{}

// ProposalRoute returns the routing key of a cbr proposal.
func (cp *CbrProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (cp *CbrProposal) ProposalType() string { return ProposalTypeCbridge }

// ValidateBasic validates the parameter change proposal
func (cp *CbrProposal) ValidateBasic() error {
	// todo: validate cp.CbrConfig?
	return nil
}
