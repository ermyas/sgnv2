package executor

// flags
const (
	FlagExecutorDbUrl = "db.url"
)

// monitor event names
const (
	MessageBusEventExecuted = "Executed"
	LiqBridgeEventRelay     = "Relay"
	PegBridgeEventMint      = "Mint"
	PegVaultEventWithdrawn  = "Withdrawn"
)
