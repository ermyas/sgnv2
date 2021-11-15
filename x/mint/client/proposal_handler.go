package client

import (
	govcli "github.com/celer-network/sgn-v2/x/gov/client"
	"github.com/celer-network/sgn-v2/x/mint/client/cli"
	"github.com/celer-network/sgn-v2/x/mint/client/rest"
)

// Aliases for gov NewProposalHandler
var (
	AdjustProvisionsProposalHandler = govcli.NewProposalHandler(cli.GetCmdSubmitAdjustProvisionsProposal, rest.AdjustProvisionsProposalRESTHandler)
)
