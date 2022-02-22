# State / Config

The keys are defined in [types/keys.go](../types/keys.go).

## State keys
1. liquidity map, lm-chid-token-lp -> amount big.Int.Bytes
2. processed add liquidity event, evliqadd-chid-seq -> true, to avoid process same event again
3. send event, evsend-%x transferid, module has seen this event, value is enum status
4. no longer save evrelay # relay event, evrelay-%x relay transferid -> srcTransferid
5. xfer relay: xferRelay-%x, src transfer id, relay msg and sigs
6. no longer need withdrawSeq
7. withdraw detail, wdDetail-%x-%d user addr and reqid, value is onchain msg and sigs
8. xfer refund, xferRefund-%x src xfer id -> withdrawonchain, only for failed xfer. first set when apply send, but no reqid, later when user InitWithdraw, set reqid in it
9. lp fee, lpfee-chid-token-lp -> fee big.Int bytes on this (chain,token)
10. sgn fee, sgnfee-chid-token -> big.Int bytes
11. liquidity sum of liqsum-chid-token, always equal sum of all lm-chid-token-xxx, we keep sum to avoid iter over all lps

## Config keys
1. fee percentage goes to cbridge lp, eg. 80 means 80% goes to lp
2. chid-tokenAddr -> asset symbol string eg. "USDT", all uppercase
3. symbol-chid -> ChainAsset, note proto has dup info symbol and chain_id
4. chid1-chid2 -> ChainPair. keys are sorted so chid1 < chid2
5. pick lp size, how many LPs on first select. value is big.Int bytes
6. chid -> gas price big.Int.Bytes.
7. chid -> GasTokenSymbol string.
8. symbol -> uint32(USD price * 1e(4+ExtraPower10) if ExtraPower10 isn't set, just 1e4
9. chid -> GasCostParam
10. chid -> GasCost
11. symbol-chid1-chid2 -> ChainPair. per (chainpair, token) info override
12. symbol -> uint32 ExtraPower10, only exist if asset USD float < $0.0001
