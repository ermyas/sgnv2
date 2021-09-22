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
)

func NewQueryRelayParams(xrefId []byte) *QueryRelayParams {
	return &QueryRelayParams{
		XrefId: xrefId,
	}
}

type QueryRelayParams struct {
	XrefId []byte `json:"xref_id,omitempty"`
}
