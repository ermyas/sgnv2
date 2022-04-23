package types

import (
	"fmt"

	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
)

const (
	ProposalTypePeg               = "PegConfigChange"
	ProposalTypePairDelete        = "PegPairDelete"
	ProposalTypeTotalSupplyUpdate = "PegTotalSupplyUpdate"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypePeg)
	govtypes.RegisterProposalType(ProposalTypePairDelete)
	govtypes.RegisterProposalType(ProposalTypeTotalSupplyUpdate)
}

var _ govtypes.Content = &PegProposal{}
var _ govtypes.Content = &PairDeleteProposal{}
var _ govtypes.Content = &TotalSupplyUpdateProposal{}

// ProposalRoute returns the routing key of a cbr proposal.
func (cp *PegProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (cp *PegProposal) ProposalType() string { return ProposalTypePeg }

// ValidateBasic validates the parameter change proposal
func (cp *PegProposal) ValidateBasic() error {
	return nil
}

func (cp *PegProposal) Validate() error {
	for _, vc := range cp.GetPegConfig().GetOriginalTokenVaults() {
		if vc.GetContract().ChainId == 0 || vc.GetContract().Address == "" {
			return fmt.Errorf("invalid vault contract config")
		}
	}
	for _, pc := range cp.GetPegConfig().GetPeggedTokenBridges() {
		if pc.GetContract().ChainId == 0 || pc.GetContract().Address == "" {
			return fmt.Errorf("invalid pegbridge contract config")
		}
	}
	return nil
}

// ProposalRoute returns the routing key of a pegbr proposal.
func (cp *PairDeleteProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (cp *PairDeleteProposal) ProposalType() string { return ProposalTypePairDelete }

// ValidateBasic validates the parameter change proposal
func (cp *PairDeleteProposal) ValidateBasic() error {
	return nil
}

// ProposalRoute returns the routing key of a pegbr proposal.
func (cp *TotalSupplyUpdateProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (cp *TotalSupplyUpdateProposal) ProposalType() string { return ProposalTypeTotalSupplyUpdate }

// ValidateBasic validates the parameter change proposal
func (cp *TotalSupplyUpdateProposal) ValidateBasic() error {
	return nil
}
