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

// pick LPs, minus each's destChain liquidity, return how much to add on src chain
// destAmount - userGet is total fees
func (k Keeper) PickLPsAndAdjustLiquidity(ctx sdk.Context, src, dest *ChainIdTokenAddr, srcAmount, destAmount, userGet *big.Int) []*types.AddrAmt {
	kv := ctx.KVStore(k.storeKey)
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
	// todo: pick a random LP as first
	firstLPIdx := 0
	// used := make(map[string]*big.Int) // each addrhex, how much is used
	for idx := firstLPIdx; idx < len(allLPs); idx++ {
		// allLPs[idx]
	}
	// todo: logic
	// must sort list due to go map iter
	var ret []*types.AddrAmt
	return ret
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
