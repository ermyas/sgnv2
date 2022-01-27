package executor

import "time"

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

const (
	MaxPollingRetries = 10
	PollingInterval   = 6 * time.Second
	GatewayTimeout    = 5 * time.Second
)
