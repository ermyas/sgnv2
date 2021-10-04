package types

const (
	QueryParams                  = "params"
	QueryRelay                   = "relay"
	QueryChainTokensConfig       = "chain-tokens-config"
	QueryFee                     = "fee"
	QueryTransferStatus          = "xfer-status"
	QueryLiquidityDetailList     = "liquidity-detail-list"
	QueryAddLiquidityStatus      = "add-liquidity-status"
	QueryWithdrawLiquidityStatus = "withdraw-liquidity-status"
	QueryChainSigners            = "chain-signers"
	QueryLatestSigners           = "latest-signers"
	QueryDebugAny                = "debug-anykey"
)

func NewQueryRelayParams(xrefId []byte) *QueryRelayParams {
	return &QueryRelayParams{
		XrefId: xrefId,
	}
}

type QueryRelayParams struct {
	XrefId []byte `json:"xref_id,omitempty"`
}

type QueryChainSignersParams struct {
	ChainId uint64
}

func NewQueryChainSignersParams(chainId uint64) *QueryChainSignersParams {
	return &QueryChainSignersParams{ChainId: chainId}
}
