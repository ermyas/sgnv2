package keeper

import (
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
)

// various algorithms eg. compute dest chain token

// given src chain token amount, calculate how much token on dest chain
// worth the same. pre-fee
// note if decimals are different, extra careful
func (k Keeper) CalcEqualOnDestChain(src, dest *ChainIdTokenAddr, srcAmount *big.Int) *big.Int {
	return new(big.Int).Set(srcAmount)
}

// pick LPs, minus each's destChain liquidity, return how much to add on src chain
func (k Keeper) PickLPsAndAdjustLiquidity(src, dest *ChainIdTokenAddr, srcAmount, destAmount, userGet *big.Int) map[eth.Addr]*big.Int {
	return make(map[eth.Addr]*big.Int)
}

// return how much user receive on dest chain
// note fee and max fee cap
// return usergtt, total-userget=fee
func (k Keeper) CalcUserGet(src, dest *ChainIdTokenAddr, total *big.Int) *big.Int {
	return new(big.Int).Set(total)
}
