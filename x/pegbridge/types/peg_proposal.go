package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
)

const (
	ProposalTypePeg = "PegConfigChange"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypePeg)
}

var _ govtypes.Content = &PegProposal{}

// ProposalRoute returns the routing key of a cbr proposal.
func (cp *PegProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (cp *PegProposal) ProposalType() string { return ProposalTypePeg }

// ValidateBasic validates the parameter change proposal
func (cp *PegProposal) ValidateBasic() error {
	return nil
}
