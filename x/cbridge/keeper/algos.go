package keeper

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Epsilon float64 = 0.00001 // used to replace 0 so no div by 0 error
)

// various algorithms eg. compute dest chain token

// given src chain token amount, calculate how much token on dest chain
// worth the same. pre-fee
// note if decimals are different, extra careful
func CalcEqualOnDestChain(kv sdk.KVStore, src, dest *ChainIdTokenDecimal, srcAmount *big.Int) *big.Int {
	ret := new(big.Int)
	if srcAmount.Sign() <= 0 {
		return ret
	}
	// A,m,n are from chain pair config
	A, m, n, err := GetAMN(kv, src.ChId, dest.ChId)
	if err != nil {
		log.Errorln("GetAMN err:", err)
		return ret // 0 if not found chain pair
	}
	srcLiqSum := GetLiq(kv, src.ChainIdTokenAddr)
	x := Epsilon // if srcLiqSum is 0, use Epsilon to avoid div by 0
	if srcLiqSum.Sign() == 1 {
		x = amt2float(srcLiqSum, src.Decimal)
	}
	destLiqSum := GetLiq(kv, dest.ChainIdTokenAddr)
	if destLiqSum.Sign() == 0 {
		return ret // no liq on dest chain
	}
	y := amt2float(destLiqSum, dest.Decimal) // y can't be 0

	D := solveD(A, x, y, m, n)
	newx := x + amt2float(srcAmount, src.Decimal)
	newy := loopCalcNewY(A, D, newx, y, m, n)
	log.Debugln("chpair:", src.ChId, dest.ChId, "A:", A, "m:", m, "n:", n, "x:", x, "y:", y, "D:", D, "newx:", newx, "newy:", newy)
	if newy >= y {
		// not possible
		log.Errorf("newy %f > y %f", newy, y)
		return ret
	}
	retFloat := big.NewFloat(y - newy)
	retFloat.Mul(retFloat, big.NewFloat(math.Pow10(int(dest.Decimal))))
	retFloat.Int(ret) // set int in ret, accuracy doesn't matter
	return ret
}

type AddrHexAmtInt struct {
	AddrHex string
	AmtInt  *big.Int
}

// pick LPs, minus each's destChain liquidity and add srcChain liq
// for each lp, try to use all he has, if enough, we are good, if not, we move on to next LP
// fee and add liq on src are calculated based on ratio this LP contributed into destAmount
func (k Keeper) PickLPsAndAdjustLiquidity(
	ctx sdk.Context, kv sdk.KVStore, src, dest *ChainIdTokenAddr, srcAmount, destAmount, fee *big.Int, sender eth.Addr, randN uint64) {
	lpFeePerc := new(big.Int).SetBytes(kv.Get(types.CfgKeyFeePerc))
	totalLpFee := new(big.Int).Mul(fee, lpFeePerc)
	totalLpFee.Div(totalLpFee, big.NewInt(100))
	sgnFee := new(big.Int).Sub(fee, totalLpFee)
	if sgnFee.Sign() == 1 {
		AddSgnFee(kv, dest.ChId, dest.TokenAddr, sgnFee)
	}

	// get all LPs that has non-zero liquidity for dest chain
	iter := sdk.KVStorePrefixIterator(kv, []byte(fmt.Sprintf("lm-%d-%x-", dest.ChId, dest.TokenAddr)))
	defer iter.Close()
	var allLPs []*AddrHexAmtInt
	for ; iter.Valid(); iter.Next() {
		amt := new(big.Int).SetBytes(iter.Value())
		if amt.Sign() == 1 {
			allLPs = append(allLPs, &AddrHexAmtInt{
				AddrHex: getAddr(iter.Key()),
				AmtInt:  amt,
			})
		}
	}
	lpCnt := len(allLPs)
	firstLPIdx := int(randN) % lpCnt           // first LP index
	toAllocate := new(big.Int).Set(destAmount) // how much left to allocate to LP
	for cnt := 0; cnt < lpCnt; cnt++ {         // how many LPs we have used
		idx := (cnt + firstLPIdx) % lpCnt
		lpAddr := eth.Hex2Addr(allLPs[idx].AddrHex)
		if lpAddr == sender {
			// Do not swap sender's liquidity from dst chain to src chain
			continue
		}
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
		log.Errorf("toAllocate still has %s", toAllocate)
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

// ========== below impl price formula

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

// divide amt by 10**(decimal)
func amt2float(amt *big.Int, decimal uint32) float64 {
	if amt.Sign() == 0 {
		return 0
	}
	ret := new(big.Float).SetInt(amt)
	ret.Quo(ret, big.NewFloat(math.Pow10(int(decimal))))
	result, _ := ret.Float64()
	return result
}
