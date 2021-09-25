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
	EventToSign = ModuleName + "ToSign"
	// event attr for data type, value is relay or withdraw
	EvAttrType = "DataType"
	EvAttrData = "Data" // raw msg to be signed

	// due to async nature, we have to use event to tell caller the msg grpc response
	EventMsgResp  = ModuleName + "MsgResp"
	EvAttrMsgType = "MsgType" // string of MsgInitWithdrawResp or MsgSignAgainResp
	EvAttrResp    = "Resp"    // value is serlized bytes of MsgInitWithdrawResp or MsgSignAgainResp
)
