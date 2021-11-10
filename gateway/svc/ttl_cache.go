package gatewaysvc

import "time"

const cacheTTL = 10 * time.Second

type farmingApy struct {
	data      map[uint64]map[string]float64
	expiredAt time.Time
}

type tx24h struct {
	data      map[uint64]map[string]*txData
	expiredAt time.Time
}

var farmingApyCache *farmingApy
var tx24hCache *tx24h

func SetFarmingApyCache(cache map[uint64]map[string]float64) {
	if farmingApyCache == nil {
		farmingApyCache = &farmingApy{
			data:      cache,
			expiredAt: time.Now().Add(cacheTTL),
		}
	} else {
		farmingApyCache.data = cache
		farmingApyCache.expiredAt = time.Now().Add(cacheTTL)
	}
}

func GetFarmingApyCache() map[uint64]map[string]float64 {
	if farmingApyCache != nil && farmingApyCache.expiredAt.After(time.Now()) {
		return farmingApyCache.data
	} else {
		return nil
	}
}

func SetTx24hCache(cache map[uint64]map[string]*txData) {
	if tx24hCache == nil {
		tx24hCache = &tx24h{
			data:      cache,
			expiredAt: time.Now().Add(3600 * time.Second),
		}
	} else {
		tx24hCache.data = cache
		tx24hCache.expiredAt = time.Now().Add(3600 * time.Second)
	}
}

func GetTx24hCache() map[uint64]map[string]*txData {
	if tx24hCache != nil && tx24hCache.expiredAt.After(time.Now()) {
		return tx24hCache.data
	} else {
		return nil
	}
}
