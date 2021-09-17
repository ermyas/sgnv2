package keeper

import (
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// helper/util to get/set various states, similar to dal

// add delta to liq map, also update per lp info for query, if delta is negative, means deduct
func ChangeLiquidity(kv sdk.KVStore, chid uint64, token, lp eth.Addr, delta *big.Int) *big.Int {
	lqKey := types.LiqMapKey(chid, token, lp)
	value := kv.Get(lqKey)
	had := new(big.Int)
	if value != nil {
		had.SetString(string(value), 10)
	}
	had.Add(had, delta)
	if had.Sign() == -1 { // negative
		panic(string(lqKey) + " negative liquidity: " + had.String())
	}
	kv.Set(lqKey, []byte(had.String()))
	// todo: add to per lp info for query
	return had
}

// to save storage, we only set a single byte, could be more complicated if need
func SetEvLiqAdd(kv sdk.KVStore, chid, seq uint64) {
	kv.Set(types.EvLiqAddKey(chid, seq), []byte{1})
}

// if get returns non-nil, return true, otherwise false
func HasEvLiqAdd(kv sdk.KVStore, chid, seq uint64) bool {
	return kv.Get(types.EvLiqAddKey(chid, seq)) != nil
}

func HasEvSend(kv sdk.KVStore, xferId [32]byte) bool {
	return kv.Get(types.EvSendKey(xferId)) != nil
}

func HasEnoughLiq(kv sdk.KVStore, chaddr *ChainIdTokenAddr, needed *big.Int) bool {
	// sum over all liqmap, if larger than needed, return true
	return true
}
