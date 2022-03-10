package rest

import (
	commontypes "github.com/celer-network/sgn-v2/common/types"
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
		StakeToken          commontypes.ERC20Token        `json:"stake_token" yaml:"stake_token"`
		RewardTokens        []commontypes.ERC20Token      `json:"reward_tokens" yaml:"reward_tokens"`
		InitialRewardInputs []types.RewardAdjustmentInput `json:"initial_reward_inputs" yaml:"initial_reward_inputs"`
		Proposer            sdk.AccAddress                `json:"proposer" yaml:"proposer"`
		Deposit             sdk.Coins                     `json:"deposit" yaml:"deposit"`
	}

	// BatchAddPoolProposalReq defines an BatchAddPoolProposal request body.
	BatchAddPoolProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title        string              `json:"title" yaml:"title"`
		Description  string              `json:"description" yaml:"description"`
		AddPoolInfos []types.AddPoolInfo `json:"add_pool_infos" yaml:"add_pool_infos"`
		Proposer     sdk.AccAddress      `json:"proposer" yaml:"proposer"`
		Deposit      sdk.Coins           `json:"deposit" yaml:"deposit"`
	}

	// AddTokensProposalReq defines an AddTokensProposal request body.
	AddTokensProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title       string                   `json:"title" yaml:"title"`
		Description string                   `json:"description" yaml:"description"`
		Tokens      []commontypes.ERC20Token `json:"tokens" yaml:"tokens"`
		Proposer    sdk.AccAddress           `json:"proposer" yaml:"proposer"`
		Deposit     sdk.Coins                `json:"deposit" yaml:"deposit"`
	}

	// AdjustRewardProposalReq defines an AdjustRewardProposal request body.
	AdjustRewardProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title                  string                        `json:"title" yaml:"title"`
		Description            string                        `json:"description" yaml:"description"`
		PoolName               string                        `json:"pool_name" yaml:"pool_name"`
		RewardAdjustmentInputs []types.RewardAdjustmentInput `json:"reward_adjustment_inputs" yaml:"reward_adjustment_inputs"`
		RemoveDuplicates       bool                          `json:"remove_duplicates" yaml:"remove_duplicates"`
		Proposer               sdk.AccAddress                `json:"proposer" yaml:"proposer"`
		Deposit                sdk.Coins                     `json:"deposit" yaml:"deposit"`
	}

	// BatchAdjustRewardProposalReq defines an BatchAdjustRewardProposal request body.
	BatchAdjustRewardProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title             string                   `json:"title" yaml:"title"`
		Description       string                   `json:"description" yaml:"description"`
		AdjustRewardInfos []types.AdjustRewardInfo `json:"adjust_reward_infos" yaml:"adjust_reward_infos"`
		Proposer          sdk.AccAddress           `json:"proposer" yaml:"proposer"`
		Deposit           sdk.Coins                `json:"deposit" yaml:"deposit"`
	}

	// SetRewardContractsProposalReq defines an SetRewardContractsProposal request body.
	SetRewardContractsProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title           string                     `json:"title" yaml:"title"`
		Description     string                     `json:"description" yaml:"description"`
		RewardContracts []commontypes.ContractInfo `json:"tokens" yaml:"reward_contracts"`
		Proposer        sdk.AccAddress             `json:"proposer" yaml:"proposer"`
		Deposit         sdk.Coins                  `json:"deposit" yaml:"deposit"`
	}
)
