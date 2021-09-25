package rest

import (
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type (
	// AddPoolProposalReq defines an AddPoolProposal request body.
	AddPoolProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title               string                        `json:"title" yaml:"title"`
		Description         string                        `json:"description" yaml:"description"`
		PoolName            string                        `json:"pool_name" yaml:"pool_name"`
		StakeToken          types.ERC20Token              `json:"stake_token" yaml:"stake_token"`
		RewardTokens        []types.ERC20Token            `json:"reward_tokens" yaml:"reward_tokens"`
		InitialRewardInputs []types.RewardAdjustmentInput `json:"ini" yaml:"reward_tokens"`
		Proposer            sdk.AccAddress                `json:"proposer" yaml:"proposer"`
		Deposit             sdk.Coins                     `json:"deposit" yaml:"deposit"`
	}
)
