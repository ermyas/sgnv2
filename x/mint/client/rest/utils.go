package rest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type (
	// AdjustProvisionsProposalReq defines an AdjustProvisionsProposal request body.
	AdjustProvisionsProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title               string         `json:"title" yaml:"title"`
		Description         string         `json:"description" yaml:"description"`
		NewAnnualProvisions sdk.Dec        `json:"new_annual_provisions" yaml:"new_annual_provisions"`
		Proposer            sdk.AccAddress `json:"proposer" yaml:"proposer"`
		Deposit             sdk.Coins      `json:"deposit" yaml:"deposit"`
	}
)
