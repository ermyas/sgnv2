package keeper

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// helper/util to get/set various states, similar to dal

// add delta to liq map, also update per lp info for query, if delta is negative, means deduct
// return updated value
// also need to call farm stake/unstake, so we have to keep k and ctx
func (k Keeper) ChangeLiquidity(ctx sdk.Context, kv sdk.KVStore, chid uint64, token, lp eth.Addr, delta *big.Int) *big.Int {
	lqKey := types.LiqMapKey(chid, token, lp)
	value := kv.Get(lqKey)
	had := new(big.Int).SetBytes(value)
	had.Add(had, delta)
	if had.Sign() == -1 { // negative
		panic(string(lqKey) + " negative liquidity: " + had.String())
	}
	kv.Set(lqKey, []byte(had.Bytes()))

	sym := GetAssetSymbol(kv, &ChainIdTokenAddr{chid, token})
	err := k.SyncFarming(ctx, sym, chid, lp, had)
	if err != nil {
		panic("Failed to sync farming" + err.Error())
	}

	return had
}

// if not found in liq map, return 0
func GetLPBalance(kv sdk.KVStore, chid uint64, token, lp eth.Addr) *big.Int {
	lqKey := types.LiqMapKey(chid, token, lp)
	value := kv.Get(lqKey)
	return new(big.Int).SetBytes(value)
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

// do we want to add protection for status change?
func SetEvSendStatus(kv sdk.KVStore, xferId [32]byte, status types.XferStatus) {
	log.Infof("Set xfer %x to %s", xferId, status.String())
	kv.Set(types.EvSendKey(xferId), []byte{byte(status)})
}

// if not found, return 0 unknown. xferid is src xfer id
func GetEvSendStatus(kv sdk.KVStore, xferId [32]byte) types.XferStatus {
	val := kv.Get(types.EvSendKey(xferId))
	if val == nil {
		return types.XferStatus_UNKNOWN
	}
	return types.XferStatus(val[0])
}

// get total liquidity of chid, token
func GetLiq(kv sdk.KVStore, chaddr *ChainIdTokenAddr) *big.Int {
	iter := sdk.KVStorePrefixIterator(kv, []byte(fmt.Sprintf("lm-%d-%x-", chaddr.ChId, chaddr.TokenAddr)))
	defer iter.Close()
	totalLiq := new(big.Int)
	for ; iter.Valid(); iter.Next() {
		totalLiq.Add(totalLiq, new(big.Int).SetBytes(iter.Value()))
	}

	return totalLiq
}

func HasEnoughLiq(kv sdk.KVStore, chaddr *ChainIdTokenAddr, needed *big.Int, sender eth.Addr) bool {
	// sum over all liqmap, if larger than needed, return true
	iter := sdk.KVStorePrefixIterator(kv, []byte(fmt.Sprintf("lm-%d-%x-", chaddr.ChId, chaddr.TokenAddr)))
	defer iter.Close()
	totalLiq := new(big.Int)
	for ; iter.Valid(); iter.Next() {
		lpAddr, err := types.GetLpAddrFromLiqMapKey(iter.Key())
		if err != nil {
			log.Error(err)
		} else if lpAddr == sender {
			// Do not swap sender's liquidity from dst chain to src chain
			continue
		}
		totalLiq.Add(totalLiq, new(big.Int).SetBytes(iter.Value()))
		if totalLiq.Cmp(needed) >= 0 {
			return true
		}
	}
	return false
}

func GetXferRelay(kv sdk.KVStore, xferId [32]byte, cdc codec.BinaryCodec) *types.XferRelay {
	bz := kv.Get(types.XferRelayKey(xferId))
	if bz == nil {
		return nil
	}
	res := new(types.XferRelay)
	cdc.MustUnmarshal(bz, res)
	return res
}

func SetXferRelay(kv sdk.KVStore, xferId [32]byte, xferRelay *types.XferRelay, cdc codec.BinaryCodec) {
	kv.Set(types.XferRelayKey(xferId), cdc.MustMarshal(xferRelay))
}

// only set when apply relay event
func SetEvRelay(kv sdk.KVStore, relayXferId, srcXferId [32]byte) {
	kv.Set(types.EvRelayKey(relayXferId), srcXferId[:])
}

// given relay xfer id, get EvRelayKey and return src xfer id.
// if not found, return nil
func GetSrcXferId(kv sdk.KVStore, relayXferId [32]byte) []byte {
	return kv.Get(types.EvRelayKey(relayXferId))
}

func SaveWithdrawDetail(kv sdk.KVStore, userAddr eth.Addr, reqid uint64, wdd *types.WithdrawDetail) {
	raw, _ := wdd.Marshal()
	kv.Set(types.WdDetailKey(userAddr, reqid), raw)
}

// if not found, return nil
func GetWithdrawDetail(kv sdk.KVStore, userAddr eth.Addr, reqid uint64) *types.WithdrawDetail {
	raw := kv.Get(types.WdDetailKey(userAddr, reqid))
	if raw == nil {
		return nil
	}
	ret := new(types.WithdrawDetail)
	err := ret.Unmarshal(raw)
	if err != nil {
		panic("unmarshal to WithdrawDetail err: " + err.Error())
	}
	return ret
}

// during apply send, if xfer is bad_xxx, set user amount etc so later user can initwithdraw via xferid
// when user call initwithdraw, set wd seqnum value
func SetXferRefund(kv sdk.KVStore, tid [32]byte, wd *types.WithdrawOnchain) {
	raw, _ := wd.Marshal()
	kv.Set(types.XferRefundKey(tid), raw)
}

// return nil if not found
func GetXferRefund(kv sdk.KVStore, tid [32]byte) *types.WithdrawOnchain {
	raw := kv.Get(types.XferRefundKey(tid))
	if raw == nil {
		return nil
	}
	ret := new(types.WithdrawOnchain)
	ret.Unmarshal(raw)
	return ret
}

// if not found, return 0
func GetLPFee(kv sdk.KVStore, chid uint64, token, lp eth.Addr) *big.Int {
	lpfeeKey := types.LpFeeKey(chid, token, lp)
	value := kv.Get(lpfeeKey)
	return new(big.Int).SetBytes(value)
}

// add new fee, return new sum
func AddLPFee(kv sdk.KVStore, chid uint64, token, lp eth.Addr, delta *big.Int) *big.Int {
	lpfeeKey := types.LpFeeKey(chid, token, lp)
	had := new(big.Int).SetBytes(kv.Get(lpfeeKey))
	had.Add(had, delta)
	kv.Set(lpfeeKey, had.Bytes())
	return had
}

func GetSgnFee(kv sdk.KVStore, chid uint64, token eth.Addr) *big.Int {
	feeKey := types.SgnFeeKey(chid, token)
	value := kv.Get(feeKey)
	return new(big.Int).SetBytes(value)
}

// add new fee, return new sum
func AddSgnFee(kv sdk.KVStore, chid uint64, token eth.Addr, delta *big.Int) *big.Int {
	feeKey := types.SgnFeeKey(chid, token)
	had := new(big.Int).SetBytes(kv.Get(feeKey))
	had.Add(had, delta)
	kv.Set(feeKey, had.Bytes())
	return had
}
