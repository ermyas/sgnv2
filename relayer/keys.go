package relayer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	RelayerDbPrefix = []byte("relay")

	PullerKeyPrefix        = []byte{0x01} // Key prefix for puller
	SlashKeyPrefix         = []byte{0x11} // Key prefix for slash
	CbrXferKeyPrefix       = []byte{0x12} // Key prefix for cbridge transfer
	PegbrMintKeyPrefix     = []byte{0x13}
	PegbrWdKeyPrefix       = []byte{0x14}
	MsgRefundKeyPrefix     = []byte{0x15}
	XferMsgRefundKeyPrefix = []byte{0x16}
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

// get cbridge transfer key from xferId
func GetCbrXferKey(xferId []byte, destChid uint64) []byte {
	return append(GetCbrChainXferPrefix(destChid), xferId...)
}

func GetCbrChainXferPrefix(destChid uint64) []byte {
	return append(CbrXferKeyPrefix, []byte(fmt.Sprintf("-%d-", destChid))...)
}

func GetPegbrMintKey(mintChid uint64, depositChid uint64, depositId []byte) []byte {
	return append(GetPegbrMintPrefix(mintChid), []byte(fmt.Sprintf("%d-%x", depositChid, depositId))...)
}

func GetPegbrMintPrefix(mintChid uint64) []byte {
	return append(PegbrMintKeyPrefix, []byte(fmt.Sprintf("-%d-", mintChid))...)
}

func GetPegbrWdKey(wdChid uint64, burnChid uint64, burnId []byte) []byte {
	return append(GetPegbrWdPrefix(wdChid), []byte(fmt.Sprintf("%d-%x", burnChid, burnId))...)
}

func GetPegbrWdPrefix(wdChid uint64) []byte {
	return append(PegbrWdKeyPrefix, []byte(fmt.Sprintf("-%d-", wdChid))...)
}

func GetMsgRefundKey(srcChainId uint64, xferId eth.Hash) []byte {
	return append(MsgRefundKeyPrefix, []byte(fmt.Sprintf("%d-%s", srcChainId, xferId.Hex()))...)
}

func GetXferMsgRefundKey(xferId eth.Hash) []byte {
	return append(XferMsgRefundKeyPrefix, xferId.Bytes()...)
}

func ParseMsgRefundKey(msgRefundKey []byte) (srcChainId uint64, srcXferId eth.Hash) {
	content := msgRefundKey[len(MsgRefundKeyPrefix):]
	fields := strings.Split(string(content), "-")
	srcChainId, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		srcChainId = 0
	}
	srcXferId = eth.Hex2Hash(fields[1])
	return
}
