package client

import (
	"github.com/celer-network/sgn-v2/x/farming/client/cli"
	"github.com/celer-network/sgn-v2/x/farming/client/rest"
	govcli "github.com/celer-network/sgn-v2/x/gov/client"
)

var (
	// AddPoolProposalHandler alias gov NewProposalHandler
	AddPoolProposalHandler = govcli.NewProposalHandler(cli.GetCmdSubmitAddPoolProposal, rest.AddPoolProposalRESTHandler)
)
