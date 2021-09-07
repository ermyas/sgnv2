package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
)

// ProposalRESTHandler returns a ProposalRESTHandler that exposes the param
// change REST handler with a given sub-route.
func UpgradeProposalRESTHandler(cliCtx client.Context) ProposalRESTHandler {
	return ProposalRESTHandler{
		SubRoute: "upgrade",
		Handler:  postUpgradeProposalHandlerFn(cliCtx),
	}
}
func postUpgradeProposalHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
