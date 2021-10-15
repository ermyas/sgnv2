package client

import (
	"github.com/celer-network/sgn-v2/x/farming/client/cli"
	"github.com/celer-network/sgn-v2/x/farming/client/rest"
	govcli "github.com/celer-network/sgn-v2/x/gov/client"
)

// Aliases for gov NewProposalHandler
var (
	AddPoolProposalHandler      = govcli.NewProposalHandler(cli.GetCmdSubmitAddPoolProposal, rest.AddPoolProposalRESTHandler)
	AddTokensProposalHandler    = govcli.NewProposalHandler(cli.GetCmdSubmitAddTokensProposal, rest.AddTokensProposalRESTHandler)
	AdjustRewardProposalHandler = govcli.NewProposalHandler(cli.GetCmdSubmitAdjustRewardProposal, rest.AdjustRewardProposalRESTHandler)
)
