package types

// distribution module event types
const (
	EventTypeSetWithdrawAddress    = "set_withdraw_address"
	EventTypeRewards               = "rewards"
	EventTypeCommission            = "commission"
	EventTypeWithdrawRewards       = "withdraw_rewards"
	EventTypeWithdrawCommission    = "withdraw_commission"
	EventTypeProposerReward        = "proposer_reward"
	EventTypeClaimAllStakingReward = "claim_all_staking_reward"
	EventTypeClaimCBridgeFeeShare  = "claim_cbridge_fee_share"
	EventTypeClaimPegBridgeFees    = "claim_pegbridge_fees"
	EventTypeClaimMessageFees      = "claim_message_fees"

	AttributeKeyWithdrawAddress  = "withdraw_address"
	AttributeKeyValidator        = "validator"
	AttributeKeyDelegatorAddress = "delegator_address"

	AttributeValueCategory = ModuleName
)
