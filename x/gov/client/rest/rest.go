package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// REST Variable names
// nolint
const (
	RestParamsType     = "type"
	RestProposalID     = "proposal-id"
	RestDepositor      = "depositor"
	RestVoter          = "voter"
	RestProposalStatus = "status"
	RestNumLimit       = "limit"
)

// ProposalRESTHandler defines a REST handler implemented in another module. The
// sub-route is mounted on the governance REST handler.
type ProposalRESTHandler struct {
	SubRoute string
	Handler  func(http.ResponseWriter, *http.Request)
}

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx client.Context, r *mux.Router, phs []ProposalRESTHandler) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r, phs)
}

// PostProposalReq defines the properties of a proposal request's body.
type PostProposalReq struct {
	BaseReq        rest.BaseReq   `json:"base_req" yaml:"base_req"`
	Title          string         `json:"title" yaml:"title"`                     // Title of the proposal
	Description    string         `json:"description" yaml:"description"`         // Description of the proposal
	ProposalType   string         `json:"proposal_type" yaml:"proposal_type"`     // Type of proposal. Initial set {PlainTextProposal }
	Proposer       sdk.AccAddress `json:"proposer" yaml:"proposer"`               // Address of the proposer
	InitialDeposit sdk.Int        `json:"initial_deposit" yaml:"initial_deposit"` // Int to add to the proposal's deposit
}

// DepositReq defines the properties of a deposit request's body.
type DepositReq struct {
	BaseReq   rest.BaseReq   `json:"base_req" yaml:"base_req"`
	Depositor sdk.AccAddress `json:"depositor" yaml:"depositor"` // Address of the depositor
	Amount    sdk.Int        `json:"amount" yaml:"amount"`       // Int to add to the proposal's deposit
}

// VoteReq defines the properties of a vote request's body.
type VoteReq struct {
	BaseReq rest.BaseReq   `json:"base_req" yaml:"base_req"`
	Voter   sdk.AccAddress `json:"voter" yaml:"voter"`   // address of the voter
	Option  string         `json:"option" yaml:"option"` // option from OptionSet chosen by the voter
}