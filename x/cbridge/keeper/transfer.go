package keeper

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// startLpPre is the lp address prefix iter to start with
func (k Keeper) transfer(
	ctx sdk.Context, sender, token eth.Addr, amount *big.Int, srcChainId, dstChainId uint64,
	maxSlippage uint32, chargeBaseFee bool, startLpPre []byte) (
	status types.XferStatus, userReceive *big.Int, destTokenAddr eth.Addr) {

	if srcChainId == dstChainId {
		status = types.XferStatus_BAD_DEST_CHAIN
		return
	}

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

	// check the asset xfer disabled
	if srcToken.XferDisabled || destToken.XferDisabled {
		status = types.XferStatus_BAD_XFER_DISABLED
		return
	}

	// now we need to decide if this send can be completed by sgn, eg. has enough liquidity on dest chain etc
	destAmount := CalcEqualOnDestChain(kv,
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
	// perc fee is based on total destAmount, before deduct basefee
	percFee := CalcPercFee(kv, src, dest, destAmount)
	userReceive = new(big.Int).Sub(destAmount, percFee)
	baseFee := big.NewInt(0)
	if chargeBaseFee {
		baseFee = CalcBaseFee(kv, assetSym, dest.ChId)
		userReceive.Sub(userReceive, baseFee)
	}
	if isNegOrZero(userReceive) {
		// amount isn't enough to pay fees
		log.Debugln(destAmount, "less than fee. base:", baseFee, "perc:", percFee)
		status = types.XferStatus_BAD_SLIPPAGE
		return
	}
	promised := calcPromised(maxSlippage, srcToken.Decimal, destToken.Decimal, amount)
	// actual receive is less than promised
	if userReceive.Cmp(promised) == -1 {
		log.Debugf("bad slippage promised %s userReceive %s", promised, userReceive)
		status = types.XferStatus_BAD_SLIPPAGE
		return
	}

	// pick LPs, minus each's destChain liquidity, add src liquidity
	// this func DOESN'T care baseFee BY DESIGN!
	k.PickLPsAndAdjustLiquidity(ctx, kv, src, dest, amount, destAmount, percFee, destToken.Decimal, sender, startLpPre)
	// baseFee goes to sgn, if we ever want to support accurate baseFee attribution,
	// we need to save baseFee in relay detail, and upon seeing the relay event, figure out
	// its sender and add to that address. Note there is no way we can make baseFee equal the actual
	// onchain tx cost because gasprice and token usd prices are all changing constantly
	AddSgnFee(kv, dest.ChId, dest.TokenAddr, baseFee)
	status = types.XferStatus_OK_TO_RELAY
	return
}
