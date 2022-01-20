package keeper

import (
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// startLpPre is the lp address prefix iter to start with
func (k Keeper) transfer(
	ctx sdk.Context, token eth.Addr, amount *big.Int, srcChainId, dstChainId uint64,
	maxSlippage uint32, lpSender eth.Addr, startLpPre []byte) (
	status types.XferStatus, userReceive *big.Int, destTokenAddr eth.Addr, percFee, baseFee *big.Int, err error) {

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
	var destAmount *big.Int
	destAmount, err = CalcEqualOnDestChain(kv,
		&ChainIdTokenDecimal{
			ChainIdTokenAddr: src,
			Decimal:          srcToken.Decimal,
		},
		&ChainIdTokenDecimal{
			ChainIdTokenAddr: dest,
			Decimal:          destToken.Decimal,
		},
		amount, lpSender)
	if destAmount.Sign() == 0 { // avoid div by 0
		// define another enum?
		status = types.XferStatus_BAD_LIQUIDITY
		return
	}
	// check has enough liq on dest chain
	if !HasEnoughLiq(kv, dest, destAmount, lpSender) {
		status = types.XferStatus_BAD_LIQUIDITY
		return
	}
	// perc fee is based on total destAmount, before deduct basefee
	percFee = CalcPercFee(kv, src, dest, destAmount)
	userReceive = new(big.Int).Sub(destAmount, percFee)
	baseFee = big.NewInt(0)
	if lpSender == eth.ZeroAddr { // charge base fee if not internal transfer by LP
		baseFee = CalcBaseFee(kv, assetSym, dest.ChId, dest.ChId)
		userReceive.Sub(userReceive, baseFee)
	}
	if isNegOrZero(userReceive) {
		// amount isn't enough to pay fees
		status = types.XferStatus_BAD_SLIPPAGE
		return
	}
	promised := calcPromised(maxSlippage, srcToken.Decimal, destToken.Decimal, amount)
	// actual receive is less than promised
	if userReceive.Cmp(promised) == -1 {
		status = types.XferStatus_BAD_SLIPPAGE
		return
	}

	// rate limit check
	if destToken.GetMaxOutAmt() != "" {
		maxOut, ok := new(big.Int).SetString(destToken.GetMaxOutAmt(), 10)
		if ok && isPos(maxOut) && userReceive.Cmp(maxOut) == 1 {
			status = types.XferStatus_EXCEED_MAX_OUT_AMOUNT
			return
		}
	}

	// pick LPs, minus each's destChain liquidity, add src liquidity
	// this func DOESN'T care baseFee BY DESIGN!
	start := time.Now()
	err = k.PickLPsAndAdjustLiquidity(ctx, kv, src, dest, amount, destAmount, percFee, destToken.Decimal, lpSender, startLpPre)
	if err != nil {
		log.Error(err)
		status = types.XferStatus_BAD_LIQUIDITY
		return
	}
	log.Info("perfxxx pickLPs took: ", time.Since(start))

	// Attribute baseFee to current syncer in SGN. If we ever want to support accurate baseFee attribution,
	// we need to save baseFee in relay detail, and upon seeing the relay event, figure out
	// its sender and add to that address. Note there is no way we can make baseFee equal the actual
	// onchain tx cost because gasprice and token usd prices are all changing constantly.
	k.MintSgnFeeAndSendToSyncer(ctx, kv, dest.ChId, dest.TokenAddr, baseFee)
	status = types.XferStatus_OK_TO_RELAY
	return
}

func (k Keeper) QueryXferStatus(ctx sdk.Context, srcXferId eth.Hash) types.XferStatus {
	return GetEvSendStatus(ctx.KVStore(k.storeKey), srcXferId)
}

func (k Keeper) QueryXferRefund(ctx sdk.Context, srcXferId eth.Hash) *types.WithdrawOnchain {
	kv := ctx.KVStore(k.storeKey)
	return GetXferRefund(kv, srcXferId)
}
