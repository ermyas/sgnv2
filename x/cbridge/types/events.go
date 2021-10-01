package types

// must match cbr_monitor.go
const (
	// event names
	CbrEventSend  = "Send"
	CbrEventRelay = "Relay"
	// from pool.sol
	CbrEventLiqAdd   = "LiquidityAdded"
	CbrEventWithdraw = "WithdrawDone" // could be LP or user
	// from signers.sol
	CbrEventSignersUpdated = "SignersUpdated"
)

const (
	// emit cosmos event for nodes to monitor and send sig back
	EventTypeDataToSign = "data_to_sign"

	// event attr for data type, value is relay or withdraw
	AttributeKeyType = "data_type"
	AttributeKeyData = "data" // raw msg to be signed

	// due to async nature, we have to use event to tell caller the msg grpc response
	EventTypeMsgResp    = "msg_resp"
	AttributeKeyMsgType = "msg_type" // string of MsgInitWithdrawResp or MsgSignAgainResp
	AttributeKeyResp    = "resp"     // value is serlized bytes of MsgInitWithdrawResp or MsgSignAgainResp
)
