package types

import (
	"fmt"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const (
	ProposalTypeUpgrade string = "Upgrade"
)

// Implements Proposal Interface
var _ Content = &UpgradeProposal{}

func init() {
	RegisterProposalType(ProposalTypeUpgrade)
}

func NewUpgradeProposal(title, description string, plan upgradetypes.Plan) *UpgradeProposal {
	return &UpgradeProposal{title, description, plan}
}

func (sup UpgradeProposal) GetTitle() string { return sup.Title }

func (sup UpgradeProposal) GetDescription() string { return sup.Description }

func (sup UpgradeProposal) ProposalRoute() string { return upgradetypes.RouterKey }

func (sup UpgradeProposal) ProposalType() string { return ProposalTypeUpgrade }

func (sup UpgradeProposal) ValidateBasic() error {
	if err := sup.Plan.ValidateBasic(); err != nil {
		return err
	}
	return ValidateAbstract(sup)
}

func (sup UpgradeProposal) String() string {
	return fmt.Sprintf(`Software Upgrade Proposal:
  Title:       %s
  Description: %s
`, sup.Title, sup.Description)
}
