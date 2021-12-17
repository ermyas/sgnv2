package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
)

const (
	ProposalTypePeg        = "PegConfigChange"
	ProposalTypePairDelete = "PegPairDelete"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypePeg)
	govtypes.RegisterProposalType(ProposalTypePairDelete)
}

var _ govtypes.Content = &PegProposal{}
var _ govtypes.Content = &PairDeleteProposal{}

// ProposalRoute returns the routing key of a cbr proposal.
func (cp *PegProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (cp *PegProposal) ProposalType() string { return ProposalTypePeg }

// ValidateBasic validates the parameter change proposal
func (cp *PegProposal) ValidateBasic() error {
	return nil
}

// ProposalRoute returns the routing key of a cbr proposal.
func (cp *PairDeleteProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (cp *PairDeleteProposal) ProposalType() string { return ProposalTypePairDelete }

// ValidateBasic validates the parameter change proposal
func (cp *PairDeleteProposal) ValidateBasic() error {
	return nil
}
