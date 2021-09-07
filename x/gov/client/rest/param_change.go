package rest

import (
	"net/http"

	govutils "github.com/celer-network/sgn-v2/x/gov/client/utils"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// ProposalRESTHandler returns a ProposalRESTHandler that exposes the param
// change REST handler with a given sub-route.
func ParamProposalRESTHandler(cliCtx client.Context) ProposalRESTHandler {
	return ProposalRESTHandler{
		SubRoute: "param_change",
		Handler:  postParamProposalHandlerFn(cliCtx),
	}
}

func postParamProposalHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req govutils.ParamChangeProposalReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		content := govtypes.NewParameterProposal(req.Title, req.Description, req.Changes.ToParamChanges())

		msg, _ := govtypes.NewMsgSubmitProposal(content, req.Deposit, req.Proposer)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, []sdk.Msg{msg})
	}
}
