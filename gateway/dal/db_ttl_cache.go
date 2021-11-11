package dal

import (
	"sync"
	"time"

	"github.com/celer-network/sgn-v2/gateway/webapi"
)

const cacheTTL = 1 * time.Minute

type Token struct {
	data      *webapi.TokenInfo
	expiredAt time.Time
}

type Chain struct {
	data      *webapi.Chain
	txUrl     string
	expiredAt time.Time
}

var (
	tokenSymbolCache map[uint64]map[string]*Token
	tokenAddrCache   map[uint64]map[string]*Token
	chainCache       map[uint64]*Chain

	tokenLock sync.RWMutex
	chainLock sync.RWMutex
)

func SetTokenCache(chainId uint64, token *webapi.TokenInfo) {
	if token == nil {
		return
	}
	tokenLock.Lock()
	defer tokenLock.Unlock()
	if tokenSymbolCache == nil {
		tokenSymbolCache = make(map[uint64]map[string]*Token)
	}
	if tokenAddrCache == nil {
		tokenAddrCache = make(map[uint64]map[string]*Token)
	}
	if tokenSymbolCache[chainId] == nil {
		tokenSymbolCache[chainId] = make(map[string]*Token)
	}
	if tokenAddrCache[chainId] == nil {
		tokenAddrCache[chainId] = make(map[string]*Token)
	}
	data := &Token{
		data:      token,
		expiredAt: time.Now().Add(cacheTTL),
	}
	tokenSymbolCache[chainId][token.GetToken().GetSymbol()] = data
	tokenSymbolCache[chainId][token.GetToken().GetAddress()] = data
}

func GetTokenCacheBySymbol(chainId uint64, symbol string) *webapi.TokenInfo {
	if tokenSymbolCache == nil {
		return nil
	}
	tokenLock.RLock()
	defer tokenLock.RUnlock()
	cache, found := tokenSymbolCache[chainId][symbol]
	if !found {
		return nil
	}
	if cache.expiredAt.After(time.Now()) {
		return cache.data
	} else {
		return nil
	}
}

func GetTokenCacheByAddr(chainId uint64, symbol string) *webapi.TokenInfo {
	if tokenAddrCache == nil {
		return nil
	}
	tokenLock.RLock()
	defer tokenLock.RUnlock()
	cache, found := tokenAddrCache[chainId][symbol]
	if !found {
		return nil
	}
	if cache.expiredAt.After(time.Now()) {
		return cache.data
	} else {
		return nil
	}
}

func SetChainCache(chain *webapi.Chain, txUrl string) {
	if chain == nil {
		return
	}
	chainLock.Lock()
	defer chainLock.Unlock()
	if chainCache == nil {
		chainCache = make(map[uint64]*Chain)
	}
	data := &Chain{
		data:      chain,
		txUrl:     txUrl,
		expiredAt: time.Now().Add(cacheTTL),
	}
	chainCache[uint64(chain.GetId())] = data
}

func GetChainCache(chainId uint64) (*webapi.Chain, string) {
	if chainCache == nil {
		return nil, ""
	}
	chainLock.RLock()
	defer chainLock.RUnlock()
	cache, found := chainCache[chainId]
	if !found {
		return nil, ""
	}
	if cache.expiredAt.After(time.Now()) {
		return cache.data, cache.txUrl
	} else {
		return nil, ""
	}
}
