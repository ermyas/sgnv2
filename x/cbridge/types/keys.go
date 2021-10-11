package types

import (
	"fmt"

	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	// DefaultParamspace default name for parameter store
	DefaultParamspace = ModuleName

	// this line is used by starport scaffolding # ibc/keys/name
)

var (
	// value is big.NewInt(int).Bytes
	WithdrawSeqNumKey = []byte("withdrawSeqNum")

	ChainSignersKey  = []byte("cs")
	LatestSignersKey = []byte("ls")
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetChainSignersKey(chid uint64) []byte {
	return append(ChainSignersKey, sdk.Uint64ToBigEndian(chid)...)
}

/* states owned by cbridge module
1. liquidity map, lm-chid-token-lp -> amount big.Int.Bytes
2. processed add liquidity event, evliqadd-chid-seq -> true, to avoid process same event again
3. send event, evsend-%x transferid, module has seen this event, value is enum status
4. relay event, evrelay-%x relay transferid -> srcTransferid
5. xfer relay: xferRelay-%x, src transfer id, relay msg and sigs
6. withdraw seq num: withdrawSeqNum, value is big.Int bytes
7. withdraw detail, wdDetail-%d seqnum, onchain msg and sigs
8. xfer refund, xferRefund-%x src xfer id -> withdrawonchain, only for failed xfer. first set when apply send, but no seqnum, later when user InitWithdraw, set seqnum
9. lp fee, lpfee-chid-token-lp -> fee big.Int bytes on this (chain,token)
10. sgn fee, sgnfee-chid-token -> big.Int bytes
*/

// key for liquidity map, chainid-tokenaddr-lpaddr
// value is big.Int.Bytes()
func LiqMapKey(chid uint64, token, lp eth.Addr) []byte {
	return []byte(fmt.Sprintf("lm-%d-%x-%x", chid, token, lp))
}

// value is 0x01 to indicate has applied event
func EvLiqAddKey(chid, seq uint64) []byte {
	return []byte(fmt.Sprintf("evliqadd-%d-%d", chid, seq))
}

// tid is user's transfer if. value is enum xfer status
func EvSendKey(tid [32]byte) []byte {
	return []byte(fmt.Sprintf("evsend-%x", tid))
}

// key is relay transfer id, value is ev.srcTransferId. only for debugging to connect full path
func EvRelayKey(tid [32]byte) []byte {
	return []byte(fmt.Sprintf("evrelay-%x", tid))
}

// serialized relay msg and sigs, add sig when receive msg
func XferRelayKey(tid [32]byte) []byte {
	return []byte(fmt.Sprintf("xferRelay-%x", tid))
}

// so we know which withdraw seqnum is for this xfer, only for debugging
func XferRefundKey(tid [32]byte) []byte {
	return []byte(fmt.Sprintf("xferRefund-%x", tid))
}

func WdDetailKey(seqnum uint64) []byte {
	return []byte(fmt.Sprintf("wdDetail-%d", seqnum))
}

// for chid, token, how much fee this lp has earned
func LpFeeKey(chid uint64, token, lp eth.Addr) []byte {
	return []byte(fmt.Sprintf("lpfee-%d-%x-%x", chid, token, lp))
}

func SgnFeeKey(chid uint64, token eth.Addr) []byte {
	return []byte(fmt.Sprintf("sgnfee-%d-%x", chid, token))
}

/* ================ config kv, all governable
1. fee percentage goes to cbridge lp, eg. 80 means 80% goes to lp
2. chid-tokenAddr -> asset symbol string eg. "USDT", all uppercase
3. symbol-chid -> ChainAsset, note proto has dup info symbol and chain_id
4. chid1-chid2 -> ChainPair. keys are sorted so chid1 < chid2
*/

var (
	CfgKeyFeePerc = []byte("cfg-feeperc")
)

func CfgKeyChain2Sym(chid uint64, addr eth.Addr) []byte {
	return []byte(fmt.Sprintf("cfg-ch2sym-%d-%x", chid, addr))
}

func CfgKeySym2Info(sym string, chid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-sym2info-%s-%d", sym, chid))
}

func CfgKeyChainPair(chid1, chid2 uint64) []byte {
	if chid1 > chid2 {
		panic(fmt.Sprintf("chid1 %d > chid2 %d", chid2, chid2))
	}
	return []byte(fmt.Sprintf("cfg-chpair-%d-%d", chid1, chid2))
}
