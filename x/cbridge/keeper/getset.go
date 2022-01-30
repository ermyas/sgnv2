package keeper

import (
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/relayer"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
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
	kv.Set(lqKey, had.Bytes())

	// Run SyncFarming in a separate cached context so farming errors won't affect the main flow
	sym := GetAssetSymbol(kv, &ChainIdTokenAddr{chid, token})
	cacheCtx, writeCache := ctx.CacheContext()
	err := k.SyncFarming(cacheCtx, sym, chid, lp, had)
	if err != nil {
		// Only log error, don't panic
		log.Errorln("Failed to sync farming", err)
	} else {
		ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())
		writeCache()
	}

	return had
}

func (k Keeper) GetXferRelay(ctx sdk.Context, xferId eth.Hash) (*types.XferRelay, bool) {
	kv := ctx.KVStore(k.storeKey)
	raw := kv.Get(types.XferRelayKey(xferId))
	if raw == nil {
		return nil, false
	}
	res := new(types.XferRelay)
	res.Unmarshal(raw)
	return res, true
}

// if not found in liq map, return 0
func GetLPBalance(kv sdk.KVStore, chid uint64, token, lp eth.Addr) *big.Int {
	lqKey := types.LiqMapKey(chid, token, lp)
	value := kv.Get(lqKey)
	return new(big.Int).SetBytes(value)
}

func GetLPsBalanceAtChain(kv sdk.KVStore, chid uint64, token eth.Addr) map[string]string {
	iter := sdk.KVStorePrefixIterator(kv, []byte(fmt.Sprintf("lm-%d-%x", chid, token)))
	m := make(map[string]string)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		key := iter.Key()
		value := kv.Get(key)
		amt := new(big.Int).SetBytes(value)
		if amt.Cmp(new(big.Int)) <= 0 {
			continue
		}
		lpAddr, err := types.GetLpAddrFromLiqMapKey(key)
		if err != nil {
			log.Errorf("failed to get lpAddr from LiqMap key %s: %v", string(key), err)
			continue
		}
		m[lpAddr.Hex()] = amt.String()
	}
	return m
}

// to save storage, we only set a single byte, could be more complicated if need
func SetEvLiqAdd(kv sdk.KVStore, chid, seq uint64) {
	kv.Set(types.EvLiqAddKey(chid, seq), []byte{1})
}

// if get returns non-nil, return true, otherwise false
func HasEvLiqAdd(kv sdk.KVStore, chid, seq uint64) bool {
	return kv.Get(types.EvLiqAddKey(chid, seq)) != nil
}

func HasEvSend(kv sdk.KVStore, xferId eth.Hash) bool {
	return kv.Get(types.EvSendKey(xferId)) != nil
}

// do we want to add protection for status change?
func SetEvSendStatus(kv sdk.KVStore, xferId eth.Hash, status types.XferStatus) {
	log.Infof("Set xfer %x to %s", xferId, status.String())
	kv.Set(types.EvSendKey(xferId), []byte{byte(status)})
}

// if not found, return 0 unknown. xferid is src xfer id
func GetEvSendStatus(kv sdk.KVStore, xferId eth.Hash) types.XferStatus {
	val := kv.Get(types.EvSendKey(xferId))
	if val == nil {
		return types.XferStatus_UNKNOWN
	}
	return types.XferStatus(val[0])
}

// iter over all lm-%d-%x-%x keys and sum liquidity
func GetLiqIterSum(kv sdk.KVStore, chaddr *ChainIdTokenAddr) *big.Int {
	iter := sdk.KVStorePrefixIterator(kv, []byte(fmt.Sprintf("lm-%d-%x-", chaddr.ChId, chaddr.TokenAddr)))
	defer iter.Close()
	totalLiq := new(big.Int)
	for ; iter.Valid(); iter.Next() {
		totalLiq.Add(totalLiq, new(big.Int).SetBytes(iter.Value()))
	}

	return totalLiq
}

func GetLPs(kv sdk.KVStore, chaddr *ChainIdTokenAddr) ([]eth.Addr, error) {
	iter := sdk.KVStorePrefixIterator(kv, []byte(fmt.Sprintf("lm-%d-%x-", chaddr.ChId, chaddr.TokenAddr)))
	defer iter.Close()
	addrs := []eth.Addr{}
	for ; iter.Valid(); iter.Next() {
		addr, err := types.GetLpAddrFromLiqMapKey(iter.Key())
		if err != nil {
			return nil, fmt.Errorf("failed to get addr from LiqMap key %s: %v", string(iter.Key()), err)
		}
		addrs = append(addrs, addr)
	}
	return addrs, nil
}

// return liqsum-%d-%x value as big.Int
func GetLiq(kv sdk.KVStore, chaddr *ChainIdTokenAddr) *big.Int {
	return new(big.Int).SetBytes(kv.Get(types.LiqSumKey(chaddr.ChId, chaddr.TokenAddr)))
}

// if delta is negative, means deduct, if result is negative, panic
// return new sum
func ChangeLiqSum(kv sdk.KVStore, chid uint64, token eth.Addr, delta *big.Int) *big.Int {
	has := GetLiq(kv, &ChainIdTokenAddr{
		ChId:      chid,
		TokenAddr: token,
	})
	has.Add(has, delta)
	if has.Sign() == -1 {
		panic("negative liq sum")
	}
	kv.Set(types.LiqSumKey(chid, token), has.Bytes())
	return has
}

func HasEnoughLiq(kv sdk.KVStore, chaddr *ChainIdTokenAddr, needed *big.Int, sender eth.Addr) bool {
	liqsum := GetLiq(kv, chaddr)
	// remove sender's own liquidity
	senderliq := GetLPBalance(kv, chaddr.ChId, chaddr.TokenAddr, sender)
	if isPos(senderliq) {
		liqsum.Sub(liqsum, senderliq)
	}
	if liqsum.Cmp(needed) >= 0 {
		return true
	}
	return false
}

func GetXferRelay(kv sdk.KVStore, xferId eth.Hash) *types.XferRelay {
	raw := kv.Get(types.XferRelayKey(xferId))
	if raw == nil {
		return nil
	}
	res := new(types.XferRelay)
	res.Unmarshal(raw)
	return res
}

func SetXferRelay(kv sdk.KVStore, xferId eth.Hash, xferRelay *types.XferRelay) {
	raw, _ := xferRelay.Marshal()
	kv.Set(types.XferRelayKey(xferId), raw)
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
func SetXferRefund(kv sdk.KVStore, tid eth.Hash, wd *types.WithdrawOnchain) {
	raw, _ := wd.Marshal()
	kv.Set(types.XferRefundKey(tid), raw)
}

// return nil if not found
func GetXferRefund(kv sdk.KVStore, tid eth.Hash) *types.WithdrawOnchain {
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

func (k Keeper) MintSgnFee(ctx sdk.Context, kv sdk.KVStore, chid uint64, token eth.Addr, delta *big.Int) *big.Int {
	coin, sum := k.addSgnFee(ctx, kv, chid, token, delta)
	err := k.MintFeeShare(ctx, coin)
	if err != nil {
		panic(err)
	}
	return sum
}

func (k Keeper) MintSgnFeeAndSendToSyncer(ctx sdk.Context, kv sdk.KVStore, dstChainId uint64, token eth.Addr, delta *big.Int, start time.Time, srcChainId uint64, sym string, decimal uint32) *big.Int {
	coin, sum := k.addSgnFee(ctx, kv, dstChainId, token, delta)
	syncer := k.stakingKeeper.GetSyncer(ctx)
	syncerAddr := eth.Hex2Addr(syncer.GetEthAddress())
	err := k.MintAndSendFeeShare(ctx, syncerAddr, coin)
	if err != nil {
		panic(err)
	}
	relayer.ReportBaseFeeDistribution(relayer.BridgeType_BRIDGE_TYPE_LIQUIDITY, syncerAddr, start, delta, sym, decimal, srcChainId, dstChainId)
	return sum
}

// add new fee, return new sum
func (k Keeper) addSgnFee(ctx sdk.Context, kv sdk.KVStore, chid uint64, token eth.Addr, delta *big.Int) (coin sdk.Coin, sum *big.Int) {
	feeKey := types.SgnFeeKey(chid, token)
	sum = new(big.Int).SetBytes(kv.Get(feeKey))
	sum.Add(sum, delta)
	kv.Set(feeKey, sum.Bytes())
	// Tell distribution module to add fees
	symbol := GetAssetSymbol(kv, &ChainIdTokenAddr{chid, token})
	denom := fmt.Sprintf("%s%s/%d", types.CBridgeFeeDenomPrefix, symbol, chid)
	coin = sdk.NewCoin(denom, sdk.NewIntFromBigInt(delta))
	return coin, sum
}

// if not found, return 0
func GetGasPrice(kv sdk.KVStore, chainId uint64) *big.Int {
	return new(big.Int).SetBytes(kv.Get(types.CfgKeyChain2GasPrice(chainId)))
}

func SetGasPrice(kv sdk.KVStore, gp []*types.GasPrice) {
	for _, it := range gp {
		price, success := new(big.Int).SetString(it.GetPrice(), 10)
		if !success {
			log.Errorln("SetGasPrice fail, ", gp)
			continue
		}
		kv.Set(types.CfgKeyChain2GasPrice(it.GetChainId()), price.Bytes())
	}
}

// if not found, return empty string
func GetGasTokenSymbol(kv sdk.KVStore, chid uint64) string {
	return string(kv.Get(types.CfgKeyChain2GasTokenSymbol(chid)))
}

// return int(usd price * 1e(4+extraPower10)), extraPower10
// if not found, return 0,0. see GetAssetUsdPrice for details
// this is a helper to make get gas token price a bit simpler
func GetGasTokenUsdPrice(kv sdk.KVStore, chid uint64) (uint32, uint32) {
	sym := GetGasTokenSymbol(kv, chid)
	return GetAssetUsdPrice(kv, sym)
}

// by default we scale up asset USD float by 1e4 to get uint32. however, if float < 0.0001,
// we need extra power10 to make an uint32. eg. 0.000012 will return 12, 2
// we could also add extra power10 if we want more precision of price
// extra power10 should not be set for most assets and 0 is returned
func GetAssetUsdPrice(kv sdk.KVStore, sym string) (uint32, uint32) {
	return getUint32(kv, types.CfgKeySymbol2UsdPrice(sym)), getUint32(kv, types.CfgKeySymbol2ExtraPower10(sym))
}

// set both chid->symbol and symbol->usd price uint32
func SetAssetPrice(kv sdk.KVStore, ap []*types.AssetPrice) {
	for _, it := range ap {
		sym := it.GetSymbol()
		for _, chainId := range it.GetChainIds() {
			kv.Set(types.CfgKeyChain2GasTokenSymbol(chainId), []byte(sym))
		}
		setUint32(kv, types.CfgKeySymbol2UsdPrice(sym), it.Price)
		origExtraPower10 := getUint32(kv, types.CfgKeySymbol2ExtraPower10(sym))
		if it.ExtraPower10 == 0 && origExtraPower10 != 0 {
			kv.Delete(types.CfgKeySymbol2ExtraPower10(sym))
		}
		if it.ExtraPower10 > 0 {
			// only set if ExtraPower10 > 0 to save kv write as non-exist key return 0 too
			setUint32(kv, types.CfgKeySymbol2ExtraPower10(sym), it.ExtraPower10)
		}
	}
}

// if not found, return nil
func GetRelayGasCostParam(kv sdk.KVStore, chid uint64) *types.RelayGasCostParam {
	raw := kv.Get(types.CfgKeyChain2RelayGasCostParam(chid))
	if raw == nil {
		return nil
	}
	ret := new(types.RelayGasCostParam)
	err := ret.Unmarshal(raw)
	if err != nil {
		panic("unmarshal to RelayGasCostParam err: " + err.Error())
	}
	return ret
}

func SetRelayGasCostBySignersAndGasCostParam(store sdk.KVStore, chainId uint64, s []*types.Signer, param *types.RelayGasCostParam) {
	var gasCost uint32
	if param != nil {
		gasCost = param.GetCostBase() + param.GetPerValidator()*uint32(len(s)) +
			param.GetPerSig()*types.MinSigsForQuorum(s)
	}
	setUint32(store, types.CfgKeyChain2EstimateRelayGasCost(chainId), gasCost)
}
