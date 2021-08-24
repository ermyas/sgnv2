package monitor

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	PullerKeyPrefix  = []byte{0x01} // Key prefix for puller
	PenaltyKeyPrefix = []byte{0x03} // Key prefix for penalty
)

// get puller key from mainchain txHash
func GetPullerKey(eLog ethtypes.Log) []byte {
	key := strconv.AppendUint(PullerKeyPrefix, eLog.BlockNumber, 10)
	key = strconv.AppendUint(key, uint64(eLog.Index), 10)
	return append(key, eLog.TxHash.Bytes()...)
}

// get penalty key from nonce
func GetPenaltyKey(nonce uint64) []byte {
	return append(PenaltyKeyPrefix, sdk.Uint64ToBigEndian(nonce)...)
}
