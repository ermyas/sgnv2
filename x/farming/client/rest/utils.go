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
		InitialRewardInputs []types.RewardAdjustmentInput `json:"initial_reward_inputs" yaml:"initial_reward_inputs"`
		Proposer            sdk.AccAddress                `json:"proposer" yaml:"proposer"`
		Deposit             sdk.Coins                     `json:"deposit" yaml:"deposit"`
	}

	// AddTokensProposalReq defines an AddTokensProposal request body.
	AddTokensProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title       string             `json:"title" yaml:"title"`
		Description string             `json:"description" yaml:"description"`
		Tokens      []types.ERC20Token `json:"tokens" yaml:"tokens"`
		Proposer    sdk.AccAddress     `json:"proposer" yaml:"proposer"`
		Deposit     sdk.Coins          `json:"deposit" yaml:"deposit"`
	}

	// AdjustRewardProposalReq defines an AdjustRewardProposal request body.
	AdjustRewardProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title                  string                        `json:"title" yaml:"title"`
		Description            string                        `json:"description" yaml:"description"`
		PoolName               string                        `json:"pool_name" yaml:"pool_name"`
		RewardAdjustmentInputs []types.RewardAdjustmentInput `json:"reward_adjustment_inputs" yaml:"reward_adjustment_inputs"`
		Proposer               sdk.AccAddress                `json:"proposer" yaml:"proposer"`
		Deposit                sdk.Coins                     `json:"deposit" yaml:"deposit"`
	}
)
