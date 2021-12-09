package keeper

import (
	"errors"
	"math/big"
	"strings"

	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	ErrNoChainPair = errors.New("chain pair not found")
)

// We don't use paramstore at all as the configs are complicated

// set the params, delete all param kvs first
func (k Keeper) SetCbrConfig(ctx sdk.Context, cfg types.CbrConfig) {
	kv := ctx.KVStore(k.storeKey)
	setUint32(kv, types.CfgKeyFeePerc, cfg.LpFeePerc)
	setUint32(kv, types.CfgKeyPickLpSize, cfg.PickLpSize)
	setUint32(kv, types.CfgKeyMaxGainPerc, cfg.MaxGainPerc)
	// Note: we don't iter and del all cfg-xxx key/val if we're removing asset
	// because this is VERY unlikely, also need to take care of past xfers
	// now we have xfer_disabled in asset, so there should be no need to delete
	// asset

	// lp hack for scalability test
	// chidTokenMap := make(map[uint64]eth.Addr) // only support one asset
	// go over asset and set ch2sym and sym2info
	for _, asset := range cfg.Assets {
		addr := eth.Hex2Addr(asset.Addr)
		// chidTokenMap[asset.ChainId] = addr
		sym := strings.ToUpper(asset.Symbol)
		kv.Set(types.CfgKeyChain2Sym(asset.ChainId, addr), []byte(sym))
		raw, _ := asset.Marshal()
		kv.Set(types.CfgKeySym2Info(sym, asset.ChainId), raw)
	}
	for _, chpair := range cfg.ChainPairs {
		raw, _ := chpair.Marshal()
		kv.Set(types.CfgKeyChainPair(chpair.Chid1, chpair.Chid2), raw)
		// SetLPs(kv, chpair.Chid1, chidTokenMap[chpair.Chid1])
		// SetLPs(kv, chpair.Chid2, chidTokenMap[chpair.Chid2])
	}
	for _, ov := range cfg.Override {
		raw, _ := ov.Chpair.Marshal()
		kv.Set(types.CfgKeyChainPairAssetOverride(ov.Symbol, ov.Chpair.Chid1, ov.Chpair.Chid2), raw)
	}
	for _, relayGasCost := range cfg.GetRelayGasCost() {
		raw, _ := relayGasCost.Marshal()
		kv.Set(types.CfgKeyChain2RelayGasCostParam(relayGasCost.GetChainId()), raw)
	}
	for _, chainContract := range cfg.GetCbrContracts() {
		raw, _ := chainContract.Marshal()
		kv.Set(types.CfgKeyCbrContract(chainContract.GetChainId()), raw)
	}
}

/*
func SetLPs(kv sdk.KVStore, chid uint64, token eth.Addr) {
	lpCnt := 65535
	perLpAmt := big.NewInt(1e8)
	lmKeyStr := fmt.Sprintf("lm-%d-%x-%x", chid, token, eth.ZeroAddr)
	amt := perLpAmt.Bytes() // $100 per lp
	lpAddrBegin := len(lmKeyStr) - 40
	for i := 1; i < lpCnt; i++ {
		lmKey := []byte(lmKeyStr)
		copy(lmKey[lpAddrBegin:], []byte(fmt.Sprintf("%02x", i)))
		kv.Set(lmKey, amt)
	}
	kv.Set(types.LiqSumKey(chid, token), new(big.Int).Mul(perLpAmt, big.NewInt(int64(lpCnt))).Bytes())
}
*/

func (k Keeper) SetCbrPrice(ctx sdk.Context, cfg *types.CbrPrice) {
	kv := ctx.KVStore(k.storeKey)
	SetGasPrice(kv, cfg.GetGasPrice())
	SetAssetPrice(kv, cfg.GetAssetPrice())
}

// return 0 if key not foud
func getUint32(kv sdk.KVStore, key []byte) uint32 {
	return uint32(new(big.Int).SetBytes(kv.Get(key)).Int64())
}

func setUint32(kv sdk.KVStore, key []byte, val uint32) {
	kv.Set(key, big.NewInt(int64(val)).Bytes())
}

func (k Keeper) GetCbrConfig(ctx sdk.Context) types.CbrConfig {
	var cbrConfig types.CbrConfig
	kv := ctx.KVStore(k.storeKey)
	cbrConfig.LpFeePerc = getUint32(kv, types.CfgKeyFeePerc)
	cbrConfig.PickLpSize = getUint32(kv, types.CfgKeyPickLpSize)
	cbrConfig.MaxGainPerc = getUint32(kv, types.CfgKeyMaxGainPerc)
	cbrConfig.Assets = make([]*types.ChainAsset, 0)
	cbrConfig.ChainPairs = make([]*types.ChainPair, 0)

	iter := sdk.KVStorePrefixIterator(kv, []byte("cfg-sym2info-"))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		assetRaw := iter.Value()
		asset := new(types.ChainAsset)
		asset.Unmarshal(assetRaw)
		cbrConfig.Assets = append(cbrConfig.Assets, asset)
	}

	iter2 := sdk.KVStorePrefixIterator(kv, []byte("cfg-chpair-"))
	defer iter2.Close()
	for ; iter2.Valid(); iter2.Next() {
		pairRaw := iter2.Value()
		pair := new(types.ChainPair)
		pair.Unmarshal(pairRaw)
		cbrConfig.ChainPairs = append(cbrConfig.ChainPairs, pair)
	}

	iter3 := sdk.KVStorePrefixIterator(kv, []byte("cfg-ch2relaygascostparam-"))
	defer iter3.Close()
	for ; iter3.Valid(); iter3.Next() {
		relayGasCostRaw := iter3.Value()
		relayGasCostParam := new(types.RelayGasCostParam)
		relayGasCostParam.Unmarshal(relayGasCostRaw)
		cbrConfig.RelayGasCost = append(cbrConfig.RelayGasCost, relayGasCostParam)
	}

	iter4 := sdk.KVStorePrefixIterator(kv, []byte("cfg-cbrcontract-"))
	defer iter4.Close()
	for ; iter4.Valid(); iter4.Next() {
		chainContractRaw := iter4.Value()
		chainContract := new(commontypes.ContractInfo)
		chainContract.Unmarshal(chainContractRaw)
		cbrConfig.CbrContracts = append(cbrConfig.CbrContracts, chainContract)
	}

	iter5 := sdk.KVStorePrefixIterator(kv, []byte("cfg-override-"))
	defer iter5.Close()
	for ; iter5.Valid(); iter5.Next() {
		pairRaw := iter5.Value()
		pair := new(types.ChainPair)
		pair.Unmarshal(pairRaw)
		cbrConfig.Override = append(cbrConfig.Override, &types.PerChainPairAssetOverride{
			Symbol: strings.Split(string(iter5.Key()), "-")[2], // key is cfg-override-%s-%d-%d
			Chpair: pair,
		})
	}
	return cbrConfig
}

// utils to deal with asset, chid and address

// given chid and token address, return which asset eg. USDT
// empty string if not found
func GetAssetSymbol(kv sdk.KVStore, chaddr *ChainIdTokenAddr) string {
	return string(kv.Get(types.CfgKeyChain2Sym(chaddr.ChId, chaddr.TokenAddr)))
}

// given asset symbol, return token address for chid, zero address if not found
func GetAssetInfo(kv sdk.KVStore, sym string, chid uint64) *types.ChainAsset {
	raw := kv.Get(types.CfgKeySym2Info(sym, chid))
	if raw == nil {
		return nil
	}
	asset := new(types.ChainAsset)
	asset.Unmarshal(raw)
	return asset
}

// fee percent from src to dest chain, note cfg always save smaller chid as chid1
// return value is actual fee percent * 1e6
// add sym arg to support per chainpair,token override
func GetFeePerc(kv sdk.KVStore, srcChid, destChid uint64, sym string) uint32 {
	chid1, chid2 := srcChid, destChid
	useFee1 := true // srcChain is chid1
	if chid1 > chid2 {
		chid1, chid2 = destChid, srcChid
		useFee1 = false // srcChain is chid2 so we need to use Fee2To1
	}
	pair := new(types.ChainPair)
	raw := kv.Get(types.CfgKeyChainPairAssetOverride(sym, chid1, chid2))
	if raw == nil {
		// if (chainpair,token) override not found, use the chainpair
		raw = kv.Get(types.CfgKeyChainPair(chid1, chid2))
	}
	pair.Unmarshal(raw)
	if useFee1 {
		return pair.Fee1To2
	}
	return pair.Fee2To1
}

// whether per chainpair,token override has no_curve, if no override, return false
// otherwise, return override chain pair.NoCurve
func GetOverrideNotUsingCurve(kv sdk.KVStore, srcChid, destChid uint64, sym string) bool {
	chid1, chid2 := srcChid, destChid
	if chid1 > chid2 {
		chid1, chid2 = destChid, srcChid
	}
	raw := kv.Get(types.CfgKeyChainPairAssetOverride(sym, chid1, chid2))
	if raw == nil {
		return false // no override found
	}
	pair := new(types.ChainPair)
	pair.Unmarshal(raw)
	return pair.NoCurve
}

// chain pair A, src weight as m, dst weight n = 2 - m
// if src,dest not found, return error
// sym is to support per chainpair,token override
func GetAMN(kv sdk.KVStore, srcChid, destChid uint64, sym string) (float64, float64, float64, error) {
	chid1, chid2 := srcChid, destChid
	if chid1 > chid2 {
		chid1, chid2 = destChid, srcChid
	}
	pair := new(types.ChainPair)
	raw := kv.Get(types.CfgKeyChainPairAssetOverride(sym, chid1, chid2))
	if raw == nil {
		raw = kv.Get(types.CfgKeyChainPair(chid1, chid2))
		if raw == nil {
			return 0, 0, 0, ErrNoChainPair
		}
	}
	pair.Unmarshal(raw)
	var A, m, n float64
	if chid1 == srcChid {
		m = float64(pair.Weight1) / 100
		n = 2 - m
	} else {
		// dest is ch1, src is ch2
		// dest weight n is weight1
		n = float64(pair.Weight1) / 100
		m = 2 - n
	}
	if pair.ConstA == 0 {
		A = 100 // default 100
	} else {
		A = float64(pair.ConstA)
	}
	return A, m, n, nil
}

func GetCbrContract(kv sdk.KVStore, chid uint64) (eth.Addr, bool) {
	raw := kv.Get(types.CfgKeyCbrContract(chid))
	if raw == nil {
		return eth.ZeroAddr, false
	}
	chainContract := new(commontypes.ContractInfo)
	chainContract.Unmarshal(raw)
	return eth.Hex2Addr(chainContract.GetAddress()), true
}

func GetCbrContracts(kv sdk.KVStore) map[uint64]eth.Addr {
	cbrContracts := make(map[uint64]eth.Addr)
	iter := sdk.KVStorePrefixIterator(kv, []byte("cfg-cbrcontract-"))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		chainContractRaw := iter.Value()
		chainContract := new(commontypes.ContractInfo)
		chainContract.Unmarshal(chainContractRaw)
		cbrContracts[chainContract.ChainId] = eth.Hex2Addr(chainContract.Address)
	}
	return cbrContracts
}
