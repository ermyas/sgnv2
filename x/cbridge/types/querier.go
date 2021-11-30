package types

const (
	QueryParams                  = "params"
	QueryConfig                  = "config"
	QueryRelay                   = "relay"
	QueryChainTokensConfig       = "chain-tokens-config"
	QueryFee                     = "fee"
	QueryFeePerc                 = "fee-perc"
	QueryTransferStatus          = "xfer-status"
	QueryLiquidityDetailList     = "liquidity-detail-list"
	QueryTotalLiquidity          = "total-liquidity"
	QueryAddLiquidityStatus      = "add-liquidity-status"
	QueryWithdrawLiquidityStatus = "withdraw-liquidity-status"
	QueryChainSigners            = "chain-signers"
	QueryLatestSigners           = "latest-signers"
	QueryDebugAny                = "debug-anykey"
	QueryCheckChainTokenValid    = "check-chain-token-valid"
	QueryChkLiqSum               = "chk-liqsum"
)

func NewQueryRelayParams(xrefId []byte) *QueryRelayParams {
	return &QueryRelayParams{
		XrefId: xrefId,
	}
}

type QueryRelayParams struct {
	XrefId []byte `json:"xref_id,omitempty"`
}

type QueryChainParams struct {
	ChainId uint64
}

func NewQueryChainParams(chainId uint64) *QueryChainParams {
	return &QueryChainParams{ChainId: chainId}
}
