package types

import "time"

// flags
const (
	FlagExecutorDbUrl     = "db.url"
	FlagExecutorContracts = "executor.contracts"
	FlagEnableAutoRefund  = "executor.enable_auto_refund"
	FlagGatewayGrpcUrl    = "sgnd.gateway_grpc"
	FlagSgnGrpcUrl        = "sgnd.sgn_grpc"
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
	MaxExecuteRetry   = 15
)
