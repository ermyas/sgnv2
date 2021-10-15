package types

// farming module event types
const (
	EventTypeAddPool      = "add-pool"
	EventTypeRemovePool   = "remove-pool"
	EventTypeAddToken     = "add-token"
	EventTypeAdjustReward = "adjust-reward"
	EventTypeStake        = "stake"
	EventTypeUnstake      = "unstake"
	EventTypeClaim        = "claim"
	EventTypeClaimAll     = "claim_all"

	AttributeKeyAddress              = "address"
	AttributeKeyPool                 = "pool"
	AttributeKeyToken                = "token"
	AttributeKeyStakeToken           = "stake_token"
	AttributeKeyAddAmount            = "add_amount"
	AttributeKeyRewardStartHeight    = "reward_start_height"
	AttributeKeyRewardAmountPerBlock = "reward_amount_per_block"
	AttributeKeyRewardToken          = "reward_token"
	AttributeKeyClaimed              = "claimed"

	AttributeValueCategory = ModuleName
)
