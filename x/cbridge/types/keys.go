package types

import (
	"fmt"

	"github.com/celer-network/sgn-v2/eth"
)

const (
	// ModuleName defines the module name
	ModuleName = "cbridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cbridge"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

/* states owned by cbridge module
1. liquidity map, lm-chid-token-lp -> amount
2. processed add liquidity event, evliqadd-chid-seq -> true, to avoid process same event again
3.
*/

// key for liquidity map, chainid-tokenaddr-lpaddr
// value is big.Int.Bytes()
func LiqMapKey(chid uint64, token, lp eth.Addr) []byte {
	return []byte(fmt.Sprintf("lm-%d-%s-%s", chid, eth.Addr2Hex(token), eth.Addr2Hex(lp)))
}

func EvLiqAddKey(chid, seq uint64) []byte {
	return []byte(fmt.Sprintf("evliqadd-%d-%d", chid, seq))
}
