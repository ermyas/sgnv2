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
func ChangeLiquidity(kv sdk.KVStore, chid uint64, token, lp eth.Addr, delta *big.Int) *big.Int {
	lqKey := types.LiqMapKey(chid, token, lp)
	value := kv.Get(lqKey)
	had := new(big.Int).SetBytes(value)
	had.Add(had, delta)
	if had.Sign() == -1 { // negative
		panic(string(lqKey) + " negative liquidity: " + had.String())
	}
	kv.Set(lqKey, []byte(had.Bytes()))
	// todo: add to per lp info for query
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

func HasEnoughLiq(kv sdk.KVStore, chaddr *ChainIdTokenAddr, needed *big.Int) bool {
	// sum over all liqmap, if larger than needed, return true
	iter := sdk.KVStorePrefixIterator(kv, []byte(fmt.Sprintf("lm-%d-%s-", chaddr.ChId, eth.Addr2Hex(chaddr.TokenAddr))))
	defer iter.Close()
	totalLiq := new(big.Int)
	for ; iter.Valid(); iter.Next() {
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

// increment withdraw seq num by 1 and return new value
// seq num start from 1
func IncrWithdrawSeq(kv sdk.KVStore) uint64 {
	had := GetWithdrawSeq(kv)
	newseq := had + 1
	kv.Set(types.WithdrawSeqNumKey, big.NewInt(int64(newseq)).Bytes())
	return newseq
}

func GetWithdrawSeq(kv sdk.KVStore) uint64 {
	return new(big.Int).SetBytes(kv.Get(types.WithdrawSeqNumKey)).Uint64()
}

func SaveWithdrawDetail(kv sdk.KVStore, seqnum uint64, wdd *types.WithdrawDetail) {
	raw, _ := wdd.Marshal()
	kv.Set(types.WdDetailKey(seqnum), raw)
}

// if not found, return nil
func GetWithdrawDetail(kv sdk.KVStore, seqnum uint64) *types.WithdrawDetail {
	raw := kv.Get(types.WdDetailKey(seqnum))
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
