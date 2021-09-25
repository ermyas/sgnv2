package types

// farming module event types
const (
	EventTypeAddPool      = "add-pool"
	EventTypeRemovePool   = "remove-pool"
	EventTypeAdjustReward = "adjust-reward"
	EventTypeStake        = "stake"
	EventTypeUnstake      = "unstake"
	EventTypeClaim        = "claim"

	AttributeKeyAddress              = "address"
	AttributeKeyPool                 = "pool"
	AttributeKeyStakeToken           = "stake_token"
	AttributeKeyRewardStartHeight    = "reward_start_height"
	AttributeKeyRewardAmountPerBlock = "reward_amount_per_block"
	AttributeKeyRewardToken          = "reward_token"
	AttributeKeyClaimed              = "claimed"

	AttributeValueCategory = ModuleName
)
