package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
)

const (
	ProposalTypeMsg = "MessageBusesUpdate"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeMsg)
}

var _ govtypes.Content = &MsgProposal{}

// ProposalRoute returns the routing key of a cbr proposal.
func (mp *MsgProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (mp *MsgProposal) ProposalType() string { return ProposalTypeMsg }

// ValidateBasic validates the parameter change proposal
func (mp *MsgProposal) ValidateBasic() error {
	return nil
}
