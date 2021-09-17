package relayer

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	RelayerDbPrefix = []byte("relay")

	PullerKeyPrefix = []byte{0x01} // Key prefix for puller
	SlashKeyPrefix  = []byte{0x11} // Key prefix for slash
)

// get puller key from mainchain txHash
func GetPullerKey(eLog ethtypes.Log) []byte {
	key := strconv.AppendUint(PullerKeyPrefix, eLog.BlockNumber, 10)
	key = strconv.AppendUint(key, uint64(eLog.Index), 10)
	return append(key, eLog.TxHash.Bytes()...)
}

// get slash key from nonce
func GetSlashKey(nonce uint64) []byte {
	return append(SlashKeyPrefix, sdk.Uint64ToBigEndian(nonce)...)
}
