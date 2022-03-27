package executor

import (
	"github.com/celer-network/sgn-v2/eth"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	"github.com/ethereum/go-ethereum/common"
)

func getMsgBridgeAddr(dstChain *Chain, msg *msgtypes.Message) common.Address {
	var bridgeAddr eth.Addr
	switch msg.GetTransferType() {
	case msgtypes.TRANSFER_TYPE_NULL:
		bridgeAddr = eth.ZeroAddr
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_SEND, msgtypes.TRANSFER_TYPE_LIQUIDITY_WITHDRAW:
		bridgeAddr = dstChain.LiqBridge.Address
	case msgtypes.TRANSFER_TYPE_PEG_MINT:
		bridgeAddr = dstChain.PegBridge.Address
	case msgtypes.TRANSFER_TYPE_PEG_MINT_V2:
		bridgeAddr = dstChain.PegBridgeV2.Address
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		bridgeAddr = dstChain.PegVault.Address
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW_V2:
		bridgeAddr = dstChain.PegVaultV2.Address
	}
	return bridgeAddr
}

func getMessageIdWithTransfer(dstChain *Chain, execCtx *msgtypes.ExecutionContext) []byte {
	msg := execCtx.Message
	dstBridgeAddr := getMsgBridgeAddr(dstChain, &msg)
	return execCtx.ComputeMessageId(dstBridgeAddr)
}
