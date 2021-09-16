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
	CbrEventNewSigners = "SignersUpdated"
)
