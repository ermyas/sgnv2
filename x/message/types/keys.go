package types

import (
	"encoding/binary"

	"github.com/celer-network/sgn-v2/eth"
	ec "github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName defines the module name
	ModuleName = "message"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_message"

	// Byte length of chain ID / nonce occupied
	Uint64ByteArrayLength = 8

	MessageFeeDenomPrefix = "MBF-"
)

var (
	MessageBusPrefix         = []byte{0x01}
	MessagePrefix            = []byte{0x02}
	TransferPrefix           = []byte{0x03}
	ActiveMessageIdsPrefix   = []byte{0x04}
	MessageRefundNoncePrefix = []byte{0x05}
	MessageRefundPrefix      = []byte{0x06}
	FeeClaimInfoPrefix       = []byte{0x07}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetMessageBusKey(chainId uint64) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return append(MessageBusPrefix, chainIdBytes...)
}

func GetMessageKey(messageId eth.Hash) []byte {
	return append(MessagePrefix, messageId.Bytes()...)
}

func GetRefundNonceKey() []byte {
	return MessageRefundNoncePrefix
}

func GetMessageRefundKey(srcTransferId eth.Hash) []byte {
	return append(MessageRefundPrefix, srcTransferId.Bytes()...)
}

func GetTransferKey(messageId eth.Hash) []byte {
	return append(TransferPrefix, messageId.Bytes()...)
}

func GetChainIdBytes(chainId uint64) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return chainIdBytes
}

func GetActiveMessageIdsPrefixByDstChainId(dstChainId uint64) []byte {
	return append(ActiveMessageIdsPrefix, GetChainIdBytes(dstChainId)...)
}

func GetActiveMessageIdsPrefixByDstChainIdTarget(dstChainId uint64, target eth.Addr) []byte {
	return append(ActiveMessageIdsPrefix, append(GetChainIdBytes(dstChainId), target.Bytes()...)...)
}

func GetActiveMessageIdsKey(dstChainId uint64, target eth.Addr, messageId eth.Hash) []byte {
	return append(ActiveMessageIdsPrefix, append(GetChainIdBytes(dstChainId), append(target.Bytes(), messageId[:]...)...)...)
}

func GetMessageIdFromActiveMessageIdsKey(key []byte) eth.Hash {
	// key is in the format:
	// <ActiveMessageIdsPrefix (1 byte)><dstChainId (8 bytes)><target (20 bytes)><messageId (32 bytes)>
	start := 1 + 8 + ec.AddressLength
	b := key[start : start+ec.HashLength]
	return eth.Bytes2Hash(b)
}

// GetFeeClaimInfoKey gets the fee claim info key from an Ethereum address.
func GetFeeClaimInfoKey(addr eth.Addr) []byte {
	return append(FeeClaimInfoPrefix, addr.Bytes()...)
}
