package types

const (
	PegbrEventDeposited = "Deposited"
	PegbrEventMint      = "Mint"
	PegbrEventBurn      = "Burn"
	PegbrEventWithdrawn = "Withdrawn"
)

const (
	EventTypeMintToSign     = "mint_to_sign"
	EventTypeWithdrawToSign = "withdraw_to_sign"

	AttributeKeyMintId          = "mint_id"
	AttributeKeyMintChainId     = "mint_chain_id"
	AttributeKeyWithdrawId      = "withdraw_id"
	AttributeKeyWithdrawChainId = "withdraw_chain_id"

	// due to async nature, we have to use event to tell caller the msg grpc response
	EventTypeMsgResp    = "msg_resp"
	AttributeKeyMsgType = "msg_type" // string of MsgInitWithdrawResp or MsgSignAgainResp
	AttributeKeyResp    = "resp"     // value is serialized bytes of MsgInitWithdrawResp or MsgSignAgainResp
)
