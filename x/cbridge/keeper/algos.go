package keeper

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// various algorithms eg. compute dest chain token

// given src chain token amount, calculate how much token on dest chain
// worth the same. pre-fee
// note if decimals are different, extra careful
func (k Keeper) CalcEqualOnDestChain(src, dest *ChainIdTokenAddr, srcAmount *big.Int) *big.Int {
	return new(big.Int).Set(srcAmount)
}

type AddrHexAmtInt struct {
	AddrHex string
	AmtInt  *big.Int
}

// pick LPs, minus each's destChain liquidity and add srcChain liq
// for each lp, try to use all he has, if enough, we are good, if not, we move on to next LP
// fee and add liq on src are calculated based on ratio this LP contributed into destAmount
func (k Keeper) PickLPsAndAdjustLiquidity(ctx sdk.Context, kv sdk.KVStore, src, dest *ChainIdTokenAddr, srcAmount, destAmount, fee *big.Int, randN uint64) {
	lpFeePerc := new(big.Int).SetBytes(kv.Get(types.CfgKeyFeePerc))
	totalLpFee := new(big.Int).Mul(fee, lpFeePerc)
	totalLpFee.Div(totalLpFee, big.NewInt(100))
	sgnFee := new(big.Int).Sub(fee, totalLpFee)
	if sgnFee.Sign() == 1 {
		AddSgnFee(kv, dest.ChId, dest.TokenAddr, sgnFee)
	}

	// get all LPs for dest chain
	iter := sdk.KVStorePrefixIterator(kv, []byte(fmt.Sprintf("lm-%d-%s-", dest.ChId, eth.Addr2Hex(dest.TokenAddr))))
	defer iter.Close()
	var allLPs []*AddrHexAmtInt
	for ; iter.Valid(); iter.Next() {
		allLPs = append(allLPs, &AddrHexAmtInt{
			AddrHex: getAddr(iter.Key()),
			AmtInt:  new(big.Int).SetBytes(iter.Value()),
		})
	}
	lpCnt := len(allLPs)
	firstLPIdx := int(randN) % lpCnt           // first LP index
	toAllocate := new(big.Int).Set(destAmount) // how much left to allocate to LP
	for cnt := 0; cnt < lpCnt; cnt++ {         // how many LPs we have used
		idx := (cnt + firstLPIdx) % lpCnt
		used := new(big.Int)
		if allLPs[idx].AmtInt.Cmp(toAllocate) >= 0 {
			// this lp has enough for all remaining needed liquidity
			used.Set(toAllocate)
			toAllocate.SetInt64(0)
		} else {
			// not enough, use all this lp has
			used.Set(allLPs[idx].AmtInt)
			toAllocate.Sub(toAllocate, used)
		}
		lpAddr := eth.Hex2Addr(allLPs[idx].AddrHex)
		// fee = totalFee * used/destAmt
		earnedFee := new(big.Int).Mul(used, totalLpFee)
		earnedFee.Div(earnedFee, destAmount)
		// on dest chain, minus used, plus earnedfee
		k.ChangeLiquidity(ctx, kv, dest.ChId, dest.TokenAddr, lpAddr, new(big.Int).Sub(earnedFee, used))
		AddLPFee(kv, dest.ChId, dest.TokenAddr, lpAddr, earnedFee)
		// add LP liquidity on src chain, toadd = srcAmt * used/destAmt
		addOnSrc := new(big.Int).Mul(used, srcAmount)
		addOnSrc.Div(addOnSrc, destAmount)
		if addOnSrc.Sign() == 1 {
			k.ChangeLiquidity(ctx, kv, src.ChId, src.TokenAddr, lpAddr, addOnSrc)
		}
	}
	if toAllocate.Sign() == 1 {
		// if we're here but toAllocate > 0, means we went over all LPs but still have no enough
		// what to do?
		panic(fmt.Sprintf("toAllocate still has %s", toAllocate))
	}

	return
}

// return the lp addr hex part of key, "lm-%d-%s-%s"
func getAddr(lpmapkey []byte) string {
	keystr := string(lpmapkey)
	lastDashIdx := strings.LastIndex(keystr, "-")
	return keystr[lastDashIdx+1:]
}

// total is dest amount, return fee
func CalcFee(kv sdk.KVStore, src, dest *ChainIdTokenAddr, total *big.Int) *big.Int {
	feePerc := GetFeePerc(kv, src.ChId, dest.ChId) // fee percent * 1e6
	if feePerc == 0 {
		return new(big.Int)
	}
	feeAmt := new(big.Int).Mul(total, big.NewInt(int64(feePerc)))
	feeAmt.Div(feeAmt, big.NewInt(1e6))
	// now compare feeAmt to max fee amt for dest chain token
	assetInfo := GetAssetInfo(kv, GetAssetSymbol(kv, dest), dest.ChId)
	if assetInfo == nil {
		return feeAmt
	}
	maxFee, _ := new(big.Int).SetString(assetInfo.MaxFeeAmount, 10)
	if feeAmt.Cmp(maxFee) > 0 {
		return maxFee
	}
	return feeAmt
}
