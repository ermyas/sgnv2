package executor

// flags
const (
	FlagExecutorDbUrl     = "db.url"
	FlagExecutorContracts = "executor.contracts"
	FlagGatewayUrl        = "sgnd.gateway"
)

// monitor event names
const (
	MessageBusEventExecuted = "Executed"
	LiqBridgeEventRelay     = "Relay"
	PegBridgeEventMint      = "Mint"
	PegVaultEventWithdrawn  = "Withdrawn"
)
