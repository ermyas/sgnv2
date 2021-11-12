package gatewaysvc

import (
	"sync"
	"time"
)

const cacheTTL = 5 * time.Minute

type farmingApy struct {
	data      map[uint64]map[string]*FarmingInfo
	expiredAt time.Time
}

type tx24h struct {
	data      map[uint64]map[string]*txData
	expiredAt time.Time
}

var (
	farmingApyCache *farmingApy
	tx24hCache      *tx24h

	farmingApyLock sync.RWMutex
	tx24hLock      sync.RWMutex
)

func SetFarmingApyCache(cache map[uint64]map[string]*FarmingInfo) {
	farmingApyLock.Lock()
	defer farmingApyLock.Unlock()
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

func GetFarmingApyCache() map[uint64]map[string]*FarmingInfo {
	farmingApyLock.RLock()
	defer farmingApyLock.RUnlock()
	if farmingApyCache != nil && farmingApyCache.expiredAt.After(time.Now()) {
		return farmingApyCache.data
	} else {
		return nil
	}
}

func SetTx24hCache(cache map[uint64]map[string]*txData) {
	tx24hLock.Lock()
	defer tx24hLock.Unlock()
	if tx24hCache == nil {
		tx24hCache = &tx24h{
			data:      cache,
			expiredAt: time.Now().Add(cacheTTL),
		}
	} else {
		tx24hCache.data = cache
		tx24hCache.expiredAt = time.Now().Add(cacheTTL)
	}
}

func GetTx24hCache() map[uint64]map[string]*txData {
	tx24hLock.RLock()
	defer tx24hLock.RUnlock()
	if tx24hCache != nil && tx24hCache.expiredAt.After(time.Now()) {
		return tx24hCache.data
	} else {
		return nil
	}
}
