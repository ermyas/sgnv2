package keeper

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Epsilon float64 = 0.00001 // used to replace 0 so no div by 0 error

	DefaultPickLpSize uint32 = 100 // if pick lp size isn't set in cbrConfig, use this instead
)

// various algorithms eg. compute dest chain token

// given src chain token amount, calculate how much token on dest chain
// worth the same. pre-fee
// note if decimals are different, extra careful
func CalcEqualOnDestChain(kv sdk.KVStore, src, dest *ChainIdTokenDecimal, srcAmount *big.Int) *big.Int {
	ret := new(big.Int)
	if isNegOrZero(srcAmount) {
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
	if isPos(srcLiqSum) {
		x = amt2float(srcLiqSum, src.Decimal)
	}
	destLiqSum := GetLiq(kv, dest.ChainIdTokenAddr)
	if isZero(destLiqSum) {
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
// fee and add liq on src are calculated based on ratio this LP contributed into destAmount
func (k Keeper) PickLPsAndAdjustLiquidity(
	ctx sdk.Context, kv sdk.KVStore, src, dest *ChainIdTokenAddr, srcAmount, destAmount, fee *big.Int, destDecimal uint32, sender eth.Addr, lpPre []byte) {
	lpFeePerc := new(big.Int).SetBytes(kv.Get(types.CfgKeyFeePerc))

	totalLpFee := new(big.Int).Mul(fee, lpFeePerc)
	totalLpFee.Div(totalLpFee, big.NewInt(100))
	sgnFee := new(big.Int).Sub(fee, totalLpFee)
	if isPos(sgnFee) {
		AddSgnFee(kv, dest.ChId, dest.TokenAddr, sgnFee)
	}

	pickedLPs, useByRatio := pickLPs(kv, dest.ChId, dest.TokenAddr, sender, destAmount, lpPre)
	if sumLiq(pickedLPs).Cmp(destAmount) == -1 {
		panic("not enough liq") // todo: return err or set xfer to bad_liq
	}

	toAllocate := new(big.Int).Set(destAmount) // how much left to allocate to LP
	// keep track of total liq changes on dest and src, note diff dest will be negative
	// we can't use destAmount/srcAmount etc due to per LP division rounding error
	totalDestNeg, totalSrcAdd := new(big.Int), new(big.Int)
	if useByRatio {
		// weighted random sample till we have enough liq for destAmount.
		// weight is uint64 each lp's liquidity amount divided by decimal, if 0 due to rounding, use 1
		// each element of weight slice is total weight so far. generate rand number and bisect search in
		// weight slice
		decDivisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(destDecimal)), nil)
		wtList := getWeightSlice(pickedLPs, decDivisor)
		totalWt := wtList[len(wtList)-1]
		rand.Seed(new(big.Int).SetBytes(lpPre).Int64())
		for isPos(toAllocate) {
			lpIdx := searchInts(wtList, rand.Int63n(totalWt)+1)
			lpIdx = nextNonZeroLp(pickedLPs, lpIdx)
			negAmt, srcAdd := k.updateOneLP(ctx, kv, src, dest, pickedLPs[lpIdx], toAllocate, totalLpFee, srcAmount, destAmount)
			totalDestNeg.Add(totalDestNeg, negAmt) // negative!
			totalSrcAdd.Add(totalSrcAdd, srcAdd)
			if isZero(toAllocate) {
				break // we've allocated all
			}
		}
	} else {
		// from first in pickedLPs, one by one
		for _, lp := range pickedLPs {
			negAmt, srcAdd := k.updateOneLP(ctx, kv, src, dest, lp, toAllocate, totalLpFee, srcAmount, destAmount)
			totalDestNeg.Add(totalDestNeg, negAmt) // negative!
			totalSrcAdd.Add(totalSrcAdd, srcAdd)
			if isZero(toAllocate) {
				break // we've allocated all
			}
		}
	}
	// update liqsum to keep it the same as sum over all liq map
	ChangeLiqSum(kv, dest.ChId, dest.TokenAddr, totalDestNeg)
	ChangeLiqSum(kv, src.ChId, src.TokenAddr, totalSrcAdd)
	if isPos(toAllocate) {
		panic("toallocate not 0")
	}
	return
}

// return negative big.Int liq delta on dest chain, and >=0 big.Int on src chain
func (k Keeper) updateOneLP(ctx sdk.Context, kv sdk.KVStore, src, dest *ChainIdTokenAddr, lp *AddrHexAmtInt, toAllocate, totalLpFee, srcAmount, destAmount *big.Int) (*big.Int, *big.Int) {
	used := new(big.Int)
	if lp.AmtInt.Cmp(toAllocate) >= 0 {
		// this lp has enough for all remaining needed liquidity
		used.Set(toAllocate)
		toAllocate.SetInt64(0)
	} else {
		// not enough, use all this lp has
		used.Set(lp.AmtInt)
		toAllocate.Sub(toAllocate, used)
	}
	// fee = totalFee * used/destAmt
	earnedFee := new(big.Int).Mul(used, totalLpFee)
	earnedFee.Div(earnedFee, destAmount)
	// on dest chain, minus used, plus earnedfee
	lpAddr := eth.Hex2Addr(lp.AddrHex)
	negAmt := new(big.Int).Sub(earnedFee, used)
	k.ChangeLiquidity(ctx, kv, dest.ChId, dest.TokenAddr, lpAddr, negAmt)
	AddLPFee(kv, dest.ChId, dest.TokenAddr, lpAddr, earnedFee)
	// add LP liquidity on src chain, toadd = srcAmt * used/destAmt
	addOnSrc := new(big.Int).Mul(used, srcAmount)
	addOnSrc.Div(addOnSrc, destAmount)
	if isPos(addOnSrc) {
		k.ChangeLiquidity(ctx, kv, src.ChId, src.TokenAddr, lpAddr, addOnSrc)
	}
	return negAmt, addOnSrc
}

// return idx for next positive liquidity lp, wrap around if pass last
// if all lps are 0, panic
func nextNonZeroLp(lps []*AddrHexAmtInt, begin int) int {
	lpCnt := len(lps)
	for cnt := 0; cnt < lpCnt; cnt++ {
		idx := (cnt + begin) % lpCnt
		if isPos(lps[idx].AmtInt) {
			return idx
		}
	}
	panic("lps are all zero liquidity")
}

// modified from sort.Search
func searchInts(a []int64, x int64) int {
	i, j := 0, len(a)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		if a[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

// if we have int64 overflow, rand.Int63n will panic
func getWeightSlice(orig []*AddrHexAmtInt, divisor *big.Int) []int64 {
	var ret []int64
	var wtSum int64
	for _, lp := range orig {
		wt := new(big.Int).Div(lp.AmtInt, divisor).Int64()
		if wt == 0 {
			wt = 1
		}
		wtSum += wt
		ret = append(ret, wtSum)
	}
	return ret
}

// pick LPs and return true if should use each LP by their liq ratio , or false if use one by one from first LP
func pickLPs(kv sdk.KVStore, dstchid uint64, dstToken, sender eth.Addr, destAmount *big.Int, lpPre []byte) ([]*AddrHexAmtInt, bool) {
	// select LPs that has non-zero liquidity for dest chain
	// start from the random lpPre eg. if lpPre is 0123, we'll begine w/ lp whose address start with 0123
	// per iterator api doc, if same prefix key not found, next will be chosen eg. 0124. We also need to check
	// if iterator is valid as it may be past last key. if we don't have pickLpSize LPs, we need to wrap around
	// and start another iter from first lp

	/* we have 2 iters
		|      iter2       |       iter          |
		+------------------+---------------------+
		|                  |                     |
	 firstLp            startLp                endLp
	*/
	pickLpSize := int(getUint32(kv, types.CfgKeyPickLpSize))
	if pickLpSize == 0 {
		pickLpSize = int(DefaultPickLpSize)
	}
	startLpKey := []byte(fmt.Sprintf("lm-%d-%x-%x", dstchid, dstToken, lpPre))
	// in ascii table, hyphen - is 45, next one is period . so last Lp key must be before .
	// we could also do %x-ffff...ffff
	endLpKey := []byte(fmt.Sprintf("lm-%d-%x.", dstchid, dstToken))
	firstLpKey := []byte(fmt.Sprintf("lm-%d-%x-", dstchid, dstToken))
	senderHex := eth.Addr2Hex(sender)
	// first iter
	pickedLPs, iter := pickLpTillSize(kv, startLpKey, endLpKey, pickLpSize, senderHex)
	var iter2 sdk.Iterator
	if len(pickedLPs) < pickLpSize {
		// if iter.Valid() {panic()}. iter now must be invalid otherwise pickLpTillSize should return cnt == pickLpSize
		// wrap around to iter from firstLp to pick pickLpSize-lpCnt
		var picked2 []*AddrHexAmtInt
		picked2, iter2 = pickLpTillSize(kv, firstLpKey, startLpKey, pickLpSize-len(pickedLPs), senderHex)
		pickedLPs = append(pickedLPs, picked2...)
	}
	// now we either have pickLpSize or we've picked ALL LPs but still fewer than pickLpSize
	if len(pickedLPs) < pickLpSize { // total LP count < pickLpSize
		return pickedLPs, true // use by ratio
	}
	// enough LPs, now check their sum liq
	liqSum := sumLiq(pickedLPs)
	if liqSum.Cmp(destAmount) >= 0 { // enough
		return pickedLPs, true // use by ratio
	}
	stillNeed := new(big.Int).Sub(destAmount, liqSum)
	picked, actSum := pickLpTillSum(iter, stillNeed, senderHex)
	pickedLPs = append(pickedLPs, picked...)
	if actSum.Cmp(stillNeed) == -1 { // still not enough, need to use iter2
		stillNeed.Sub(stillNeed, actSum)
		picked, actSum = pickLpTillSum(iter2, stillNeed, senderHex)
		pickedLPs = append(pickedLPs, picked...)
	}
	return pickedLPs, false // use one by one
}

// iterator from begin to end, return early if has enough, otherwise reaches end and return (iter will be invalid).
// caller need to check return value to handle 2 cases.
func pickLpTillSize(kv sdk.KVStore, begin, end []byte, size int, sender string) (picked []*AddrHexAmtInt, iter sdk.Iterator) {
	iter = kv.Iterator(begin, end)
	for ; iter.Valid(); iter.Next() {
		amt := new(big.Int).SetBytes(iter.Value())
		if isPos(amt) {
			lpAddr := getAddr(iter.Key())
			if lpAddr == sender {
				continue // don't use sender's own liquidity
			}
			picked = append(picked, &AddrHexAmtInt{
				AddrHex: lpAddr,
				AmtInt:  amt,
			})
			// if has picked enough lps, return early. note iter COULD be invalid
			// if this lp happens to be the last one
			if len(picked) == size {
				return
			}
		}
	}
	// iter is invalid
	return
}

// iter till end and if liqsum >= expSum, return early
func pickLpTillSum(iter sdk.Iterator, expSum *big.Int, sender string) (picked []*AddrHexAmtInt, actualSum *big.Int) {
	actualSum = new(big.Int)
	for ; iter.Valid(); iter.Next() {
		amt := new(big.Int).SetBytes(iter.Value())
		if isPos(amt) {
			lpAddr := getAddr(iter.Key())
			if lpAddr == sender {
				continue // don't use sender's own liquidity
			}
			picked = append(picked, &AddrHexAmtInt{
				AddrHex: lpAddr,
				AmtInt:  amt,
			})
			actualSum.Add(actualSum, amt)
			if actualSum.Cmp(expSum) >= 0 {
				return
			}
		}
	}
	// iter invalid, still not enough
	return
}

func sumLiq(lplist []*AddrHexAmtInt) *big.Int {
	sum := new(big.Int)
	for _, liq := range lplist {
		sum.Add(sum, liq.AmtInt)
	}
	return sum
}

// return the lp addr hex part of key, "lm-%d-%x-%x"
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

func isPos(i *big.Int) bool {
	return i.Sign() == 1
}

func isZero(i *big.Int) bool {
	return i.Sign() == 0
}

func isNegOrZero(i *big.Int) bool {
	return i.Sign() <= 0
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
	if isZero(amt) {
		return 0
	}
	ret := new(big.Float).SetInt(amt)
	ret.Quo(ret, big.NewFloat(math.Pow10(int(decimal))))
	result, _ := ret.Float64()
	return result
}
