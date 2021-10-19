package keeper

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// startLpPre is the lp address prefix iter to start with
func (k Keeper) Transfer(
	ctx sdk.Context, sender, token eth.Addr, amount *big.Int, srcChainId, dstChainId uint64,
	maxSlippage uint32, startLpPre []byte) (status types.XferStatus, destAmount, feeAmt *big.Int, destTokenAddr eth.Addr) {

	kv := ctx.KVStore(k.storeKey)
	src := &ChainIdTokenAddr{
		ChId:      srcChainId,
		TokenAddr: token,
	}
	assetSym := GetAssetSymbol(kv, src)
	if assetSym == "" {
		// unsupported src token, don't allow refund because this must be an attack?
		status = types.XferStatus_BAD_TOKEN
		return
		// SetXferRefund(kv, ev.TransferId, wdOnchain)
	}
	srcToken := GetAssetInfo(kv, assetSym, srcChainId)
	if srcToken == nil {
		// unsupported dest chain
		status = types.XferStatus_BAD_TOKEN
		return
	}
	destToken := GetAssetInfo(kv, assetSym, dstChainId)
	if destToken == nil {
		// unsupported dest chain
		status = types.XferStatus_BAD_TOKEN
		return
	}
	destTokenAddr = eth.Hex2Addr(destToken.Addr)
	dest := &ChainIdTokenAddr{
		ChId:      dstChainId,
		TokenAddr: destTokenAddr,
	}
	// now we need to decide if this send can be completed by sgn, eg. has enough liquidity on dest chain etc
	destAmount = CalcEqualOnDestChain(kv,
		&ChainIdTokenDecimal{
			ChainIdTokenAddr: src,
			Decimal:          srcToken.Decimal,
		},
		&ChainIdTokenDecimal{
			ChainIdTokenAddr: dest,
			Decimal:          destToken.Decimal,
		},
		amount)
	if destAmount.Sign() == 0 { // avoid div by 0
		// define another enum?
		status = types.XferStatus_BAD_LIQUIDITY
		return
	}
	// check has enough liq on dest chain
	if !HasEnoughLiq(kv, dest, destAmount, sender) {
		status = types.XferStatus_BAD_LIQUIDITY
		return
	}
	feeAmt = CalcFee(kv, src, dest, destAmount)
	userReceive := new(big.Int).Sub(destAmount, feeAmt)
	promised := calcPromised(maxSlippage, srcToken.Decimal, destToken.Decimal, amount)
	// actual receive is less than promised
	if userReceive.Cmp(promised) == -1 {
		log.Debugf("bad slippage promised %s userReceive %s", promised, userReceive)
		status = types.XferStatus_BAD_SLIPPAGE
		return
	}

	// pick LPs, minus each's destChain liquidity, add src liquidity
	k.PickLPsAndAdjustLiquidity(ctx, kv, src, dest, amount, destAmount, feeAmt, destToken.Decimal, sender, startLpPre)

	status = types.XferStatus_OK_TO_RELAY
	return
}
