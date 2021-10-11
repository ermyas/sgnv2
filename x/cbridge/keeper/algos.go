package keeper

import (
	"fmt"
	"math"
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
func CalcEqualOnDestChain(kv sdk.KVStore, src, dest *ChainIdTokenAddr, srcAmount *big.Int) *big.Int {
	ret := new(big.Int)
	if srcAmount.Sign() <= 0 {
		return ret
	}
	return new(big.Int).Set(srcAmount)
	/*
		// A,m,n are from chain pair config
		A, m, n, err := GetAMN(kv, src.ChId, dest.ChId)
		if err != nil {
			return ret // 0 if not found chain pair
		}

		// x and y are sum of liquidity, divided by corresponding decimal to get int only
		// what if not even 1? or we use big.Float?
		x, y, newx := GetXY(kv, src, dest, srcAmount)

		D := solveD(A, x, y, m, n)
		newy := loopCalcNewY(A, D, newx, y, m, n)

		if newy >= y {
			// not possible
			return ret
		}
		(y - newy) * ydecimal
	*/
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

/* solveD is faster and provides accurate answer
// f(D) = \frac{D^3}{4x_i^{w_i}x_j^{w_j}}+(4A-1)D-4A(x_i+x_j)
//      = D^3 + (4A-1){4x_i^{w_i}x_j^{w_j}}D - 4A(x_i+x_j){4x_i^{w_i}x_j^{w_j}} = 0
func loopCalcD(A, x, y, m, n float64) float64 {
	D := x + y
	for i := 0; i < 100; i++ {
		Dprev := D
		xtimesy := 4 * math.Pow(x, m) * math.Pow(y, n)
		fD := math.Pow(D, 3) + (4*A-1)*(xtimesy)*D - 4*A*(x+y)*xtimesy
		fDprime := 3*math.Pow(D, 2) + 4*(A-1)*(xtimesy)
		D = D - fD/fDprime
		if math.Abs(D-Dprev) < 0.01 {
			return D
		}
	}
	return D
}
*/

// we can solve D directly, p = (4A-1){4x_i^{w_i}x_j^{w_j}}
// q = - 4A(x_i+x_j){4x_i^{w_i}x_j^{w_j}}
func solveD(A, x, y, m, n float64) float64 {
	xtimesy := 4 * math.Pow(x, m) * math.Pow(y, n)
	p := (4*A - 1) * xtimesy
	q := -4 * A * (x + y) * xtimesy
	pqrt := math.Sqrt(math.Pow(q/2, 2) + math.Pow(p/3, 3))
	return math.Cbrt(pqrt-q/2) + math.Cbrt(-pqrt-q/2)
}

// given D and new xi, calculate xj, prev xj - new xj is equal amount
// y is xj for simpler code, m is weight i, n is weight j
// f(y) = 4Ay^(wj+1) + (4AX+D-4AD)y^wj - D^3/4(x^wi)
func loopCalcNewY(A, D, x, y, m, n float64) float64 {
	ret := y
	for i := 0; i < 100; i++ {
		retPrev := ret
		yPowN := math.Pow(ret, n)
		A4 := 4 * A
		fy := A4*yPowN*ret + (A4*x-A4*D+D)*yPowN - math.Pow(D, 3)/(4*math.Pow(x, m))
		fyprime := (n+1)*A4*yPowN + n*(A4*x-A4*D+D)*math.Pow(ret, n-1)
		ret = ret - fy/fyprime
		if math.Abs(ret-retPrev) < 0.01 {
			return ret
		}
	}
	return ret
}

func invarLeft(A, D, x, y float64) float64 {
	return 4*A*(x+y) + D
}

func invarRight(A, D, x, y, m, n float64) float64 {
	return 4*A*D + math.Pow(D, 3)/(4*math.Pow(x, m)*math.Pow(y, n))
}
