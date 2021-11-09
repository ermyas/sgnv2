package keeper

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Epsilon float64 = 0.00001 // used to replace 0 so no div by 0 error

	DefaultPickLpSize uint32 = 1000 // if pick lp size isn't set in cbrConfig, use this instead
)

// various algorithms eg. compute dest chain token

// given src chain token amount, calculate how much token on dest chain
// worth the same. pre-fee
// note if decimals are different, extra careful
func CalcEqualOnDestChain(kv sdk.KVStore, src, dest *ChainIdTokenDecimal, srcAmount *big.Int, lpAddr eth.Addr) (*big.Int, error) {
	ret := new(big.Int)
	if isNegOrZero(srcAmount) {
		return ret, fmt.Errorf("invalid src amt")
	}
	// A,m,n are from chain pair config
	A, m, n, err := GetAMN(kv, src.ChId, dest.ChId)
	if err != nil {
		return ret, fmt.Errorf("GetAMN err: %w", err) // 0 if not found chain pair
	}
	srcLiqSum := GetLiq(kv, src.ChainIdTokenAddr)
	if lpAddr != eth.ZeroAddr { // internal LP transfer for cross-chain withdrawal
		// caller need to make sure srcAmount <= srcLiqSum
		if srcLiqSum.Cmp(srcAmount) < 0 {
			return ret, fmt.Errorf("insufficient balance, srcLiqSum %s < srcAmount %s", srcLiqSum, srcAmount)
		}
		srcLiqSum.Sub(srcLiqSum, srcAmount)
	}
	x := Epsilon // if srcLiqSum is 0, use Epsilon to avoid div by 0
	if isPos(srcLiqSum) {
		x = amt2float(srcLiqSum, src.Decimal)
	}
	destLiqSum := GetLiq(kv, dest.ChainIdTokenAddr)
	if lpAddr != eth.ZeroAddr { // internal LP transfer for cross-chain withdrawal
		balance := GetLPBalance(kv, dest.ChainIdTokenAddr.ChId, dest.ChainIdTokenAddr.TokenAddr, lpAddr)
		destLiqSum.Sub(destLiqSum, balance)
	}
	if isZero(destLiqSum) {
		return ret, fmt.Errorf("no liqiudity on dest chain") // no liq on dest chain
	}
	log.Infoln("srcLiqSum:", srcLiqSum, "destLiqSum:", destLiqSum)
	y := amt2float(destLiqSum, dest.Decimal) // y can't be 0

	D := solveD(A, x, y, m, n)
	newx := x + amt2float(srcAmount, src.Decimal)
	newy := loopCalcNewY(A, D, newx, y, m, n)
	log.Debugln("chpair:", src.ChId, dest.ChId, "A:", A, "m:", m, "n:", n, "x:", x, "y:", y, "D:", D, "newx:", newx, "newy:", newy)
	if math.IsNaN(newy) {
		// not possible as we already override negative ret in loopCalcNewY
		// keep this check for extra caution
		return ret, fmt.Errorf("newy is NaN")
	}
	if newy < 0 {
		// cloopCalcNewY ould return negative result when
		// ret = ret - fy/fyprime is neg and math.Abs(ret-retPrev) < 0.01
		// if this ever becomes annoying for users, we could just set newy
		// to 0
		return ret, fmt.Errorf("newy %f < 0", newy)
	}
	if newy >= y {
		// newton's method failed
		return ret, fmt.Errorf("newy %f >= y %f, newton method failed", newy, y)
	}
	retFloat := big.NewFloat(y - newy)
	retFloat.Mul(retFloat, big.NewFloat(math.Pow10(int(dest.Decimal))))
	retFloat.Int(ret) // set int in ret, accuracy doesn't matter
	return ret, nil
}

type AddrHexAmtInt struct {
	AddrHex string
	AmtInt  *big.Int
}

// pick LPs, minus each's destChain liquidity and add srcChain liq
// fee and add liq on src are calculated based on ratio this LP contributed into destAmount
// WARNING: this func doesn't care base fee BY DESIGN!!!
func (k Keeper) PickLPsAndAdjustLiquidity(
	ctx sdk.Context, kv sdk.KVStore, src, dest *ChainIdTokenAddr, srcAmount, destAmount, fee *big.Int, destDecimal uint32, sender eth.Addr, lpPre []byte) error {
	// don't write to kv before possible return error because it'll cause wrong state
	start := time.Now()
	pickedLPs, useByRatio := pickLPs(kv, dest.ChId, dest.TokenAddr, sender, destAmount, lpPre)
	log.Infoln("perfxxx picked", len(pickedLPs), "lps, byratio:", useByRatio, "took:", time.Since(start))
	if sumLiq(pickedLPs).Cmp(destAmount) == -1 {
		return fmt.Errorf("sumliq of picked LPs less than needed destAmt. %s < %s", sumLiq(pickedLPs), destAmount)
	}
	for _, lp := range pickedLPs {
		log.Debugln("lp:", lp.AddrHex, "amt:", lp.AmtInt)
	}
	// calc fees
	lpFeePerc := new(big.Int).SetBytes(kv.Get(types.CfgKeyFeePerc))
	totalLpFee := new(big.Int).Mul(fee, lpFeePerc)
	totalLpFee.Div(totalLpFee, big.NewInt(100))
	sgnFee := new(big.Int).Sub(fee, totalLpFee)
	if isPos(sgnFee) {
		k.AddSgnFee(ctx, kv, dest.ChId, dest.TokenAddr, sgnFee)
	}

	// update LP's liquidity
	toAllocate := new(big.Int).Set(destAmount) // how much left to allocate to LP, will be reduced in updateOneLP
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
		randSeed := new(big.Int).SetBytes(lpPre).Int64()
		log.Debugln("seed:", randSeed, "wtList:", wtList)
		rand.Seed(randSeed)
		for isPos(toAllocate) {
			x := rand.Int63n(totalWt) + 1
			lpIdx := searchInts(wtList, x)
			log.Debugln("x:", x, "lpIdx:", lpIdx)
			lpIdx2 := nextNonZeroLp(pickedLPs, lpIdx) // need this because we may choose same LP again due to rand
			if lpIdx2 != lpIdx {
				log.Debugln("new lpIdx:", lpIdx2)
				lpIdx = lpIdx2
			}
			negAmt, srcAdd := k.updateOneLP(ctx, kv, src, dest, pickedLPs[lpIdx], toAllocate, totalLpFee, srcAmount, destAmount)
			log.Infoln("use lp:", pickedLPs[lpIdx].AddrHex, negAmt)
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
			log.Infoln("use lp:", lp.AddrHex, negAmt)
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
		// we can't return err now because kv is already modified, so panic
		panic(fmt.Sprintf("toAllocate still has %s left", toAllocate))
	}
	return nil
}

// return negative big.Int liq delta on dest chain, equals earnedFee - used
// and >=0 big.Int on src chain to add
// will also reduce toAllocate and lp.AmtInt by amount used. Note here it can't be
// negamt because we only consider liquidity in toAllocate
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
	lp.AmtInt.Sub(lp.AmtInt, used) // so this won't be picked again
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
	defer iter.Close()
	defer func() {
		if iter2 != nil {
			iter2.Close()
		}
	}()
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
// note iter can't be closed before return as we may need to resume iter
func pickLpTillSize(kv sdk.KVStore, begin, end []byte, size int, sender string) (picked []*AddrHexAmtInt, iter sdk.Iterator) {
	iter = kv.Iterator(begin, end)
	log.Infoln("pickTillSize:", string(begin), string(end), "size:", size)
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
	log.Infoln("pickTillSum:", expSum)
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

// total is dest amount, return percent fee based on it
func CalcPercFee(kv sdk.KVStore, src, dest *ChainIdTokenAddr, total *big.Int) *big.Int {
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

// base fee only depends on asset price, dest chain gas token price, dest chain gas price and relay gas cost
func CalcBaseFee(kv sdk.KVStore, assetSym string, destChid uint64) (baseFee *big.Int) {
	baseFee = new(big.Int)
	gasTokenUsdPrice := GetGasTokenUsdPrice(kv, destChid)
	assetUsdPrice := GetAssetUsdPrice(kv, assetSym)
	assetInfo := GetAssetInfo(kv, assetSym, destChid)
	gasCost := getUint32(kv, types.CfgKeyChain2EstimateRelayGasCost(destChid))

	gasPrice := GetGasPrice(kv, destChid)
	// formula is gasCost * gasPrice * gasTokenPrice / 1e18 / assetPrice
	if assetUsdPrice == 0 {
		log.Warnln("chainid:", destChid, "asset", assetSym, "usd price is 0")
		return // avoid div by 0
	}
	baseFee.Mul(gasPrice, big.NewInt(int64(gasCost)))
	baseFee.Mul(baseFee, big.NewInt(int64(gasTokenUsdPrice)))
	baseFee.Div(baseFee, big.NewInt(int64(assetUsdPrice)))
	baseFee.Div(baseFee, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18-assetInfo.Decimal)), nil)) // gas token always 18 decimal
	log.Debugf("basefee: %s, chid: %d, gasprice: %s, gascost: %d, gastokenusd: %d, assetusd: %d", baseFee, destChid, gasPrice, gasCost, gasTokenUsdPrice, assetUsdPrice)
	return
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
// f(y) = 4Ay^(n+1) + (4Ax-4AD+D)y^n - D^3/(4x^m)
func loopCalcNewY(A, D, x, y, m, n float64) float64 {
	A4 := 4 * A
	// pick initial guess as y as newy should usually be close to y. however, it's possible
	// when curve are skewed and newy will be close to 0. we'll check ret and if it's negative,
	// we'll choose Epsilon as initial guess and try again
	ret := y
	for i := 0; i < 100; i++ {
		if ret < 0 {
			log.Infoln("loopCalcNewY neg ret, use Epsilon.", A, D, x, y, m, n, ret, i)
			ret = Epsilon
		}
		retPrev := ret
		yPowN := math.Pow(ret, n)
		fy := A4*yPowN*ret + (A4*x-A4*D+D)*yPowN - math.Pow(D, 3)/(4*math.Pow(x, m))
		fyprime := (n+1)*A4*yPowN + n*(A4*x-A4*D+D)*math.Pow(ret, n-1)
		ret = ret - fy/fyprime
		if math.Abs(ret-retPrev) < 0.01 {
			return ret
		}
	}
	// in case newton method doesn't converge after 100 runs, consider it failed and
	// return y so caller will return error
	return y
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
