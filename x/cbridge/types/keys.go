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

var (
	// value is list of src xfer id that are relay haven't been sent
	// to be pulled by relayer and send onchain, xfer will be removed
	// when x/cbridge sees relay event
	ToRelayXfersKey = []byte("torelay")
	// value is big.NewInt(int).Bytes
	WithdrawSeqNumKey = []byte("withdrawSeqNum")
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

/* states owned by cbridge module
1. liquidity map, lm-chid-token-lp -> amount big.Int.Bytes
2. processed add liquidity event, evliqadd-chid-seq -> true, to avoid process same event again
3. send event, evsend-%x transferid, module has seen this event
4. relay event, evrelay-%x transferid -> srcTransferid
5. transferDetail: xfer-%x transferid -> tbd. only src transferid in key!
6. torelaytransferIds: torelay -> [src xfer id]
7. xfer relay: xferRelay-%x, src transfer id, relay msg and sigs
8. withdraw seq num: withdrawSeqNum, value is big.Int bytes
9. withdraw detail, wdDetail-%d seqnum, onchain msg and sigs
*/

// key for liquidity map, chainid-tokenaddr-lpaddr
// value is big.Int.Bytes()
func LiqMapKey(chid uint64, token, lp eth.Addr) []byte {
	return []byte(fmt.Sprintf("lm-%d-%s-%s", chid, eth.Addr2Hex(token), eth.Addr2Hex(lp)))
}

func EvLiqAddKey(chid, seq uint64) []byte {
	return []byte(fmt.Sprintf("evliqadd-%d-%d", chid, seq))
}

func EvSendKey(tid [32]byte) []byte {
	return []byte(fmt.Sprintf("evsend-%x", tid))
}

// relay transfer id, value is ev.srcTransferId
func EvRelayKey(tid [32]byte) []byte {
	return []byte(fmt.Sprintf("evrelay-%x", tid))
}

// only source transfer ID! for relay transfer id, use evrelay-%x to get src transfer id
func XferDetailKey(tid [32]byte) []byte {
	return []byte(fmt.Sprintf("xfer-%x", tid))
}

// serialized relay msg and sigs, add sig when receive msg
func XferRelayKey(tid [32]byte) []byte {
	return []byte(fmt.Sprintf("xferRelay-%x", tid))
}

func WdDetailKey(seqnum uint64) []byte {
	return []byte(fmt.Sprintf("wdDetail-%d", seqnum))
}
