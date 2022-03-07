package types

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func NewMsgXferExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	chainId uint64,
	dstToken eth.Addr,
	dstAmt string,
	dstBridge eth.Addr,
	transferType TransferType) *ExecutionContext {

	message := Message{
		SrcChainId:      chainId,
		DstChainId:      ev.DstChainId.Uint64(),
		Sender:          eth.Addr2Hex(ev.Sender),
		Receiver:        eth.Addr2Hex(ev.Receiver),
		Data:            ev.Message,
		Fee:             ev.Fee.String(),
		TransferType:    transferType,
		TransferRefId:   ev.SrcTransferId[:],
		ExecutionStatus: EXECUTION_STATUS_PENDING,
	}
	transfer := &Transfer{
		Token:  dstToken.Bytes(),
		Amount: dstAmt,
	}
	execCtx := &ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	execCtx.MessageId = execCtx.ComputeMessageId(dstBridge)
	return execCtx
}

func NewMsgXferRefundExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	wdOnchain *cbrtypes.WithdrawOnchain,
	nonce uint64,
	bridge eth.Addr) *ExecutionContext {

	message := Message{
		SrcChainId:      wdOnchain.Chainid,
		DstChainId:      wdOnchain.Chainid,
		Sender:          eth.ZeroAddrHex,
		Receiver:        eth.Addr2Hex(ev.Sender),
		Data:            ev.Message,
		TransferType:    TRANSFER_TYPE_LIQUIDITY_WITHDRAW,
		TransferRefId:   ev.SrcTransferId[:],
		ExecutionStatus: EXECUTION_STATUS_PENDING,
		Fee:             ev.Fee.String(),
	}
	transfer := &Transfer{
		Amount: new(big.Int).SetBytes(wdOnchain.Amount).String(),
		Token:  wdOnchain.Token,
		SeqNum: nonce,
	}
	execCtx := &ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	execCtx.MessageId = execCtx.ComputeMessageId(bridge)
	return execCtx
}

func NewMsgPegDepositRefundExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	wdOnChain pegbrtypes.WithdrawOnChain,
	bridge eth.Addr) *ExecutionContext {

	message := Message{
		SrcChainId:      wdOnChain.RefChainId,
		DstChainId:      wdOnChain.RefChainId,
		Sender:          eth.ZeroAddrHex,
		Receiver:        eth.Bytes2AddrHex(wdOnChain.Receiver),
		Data:            ev.Message,
		TransferType:    TRANSFER_TYPE_PEG_WITHDRAW,
		TransferRefId:   wdOnChain.RefId,
		ExecutionStatus: EXECUTION_STATUS_PENDING,
		Fee:             ev.Fee.String(),
	}
	transfer := &Transfer{
		Amount: new(big.Int).SetBytes(wdOnChain.Amount).String(),
		Token:  wdOnChain.Token,
	}
	execCtx := &ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	execCtx.MessageId = execCtx.ComputeMessageId(bridge)
	return execCtx
}

func NewMsgPegBurnRefundExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	mintOnChain pegbrtypes.MintOnChain,
	bridge eth.Addr) *ExecutionContext {

	message := Message{
		SrcChainId:      mintOnChain.RefChainId,
		DstChainId:      mintOnChain.RefChainId,
		Sender:          eth.ZeroAddrHex,
		Receiver:        eth.Bytes2AddrHex(mintOnChain.Account),
		Data:            ev.Message,
		TransferType:    TRANSFER_TYPE_PEG_MINT,
		TransferRefId:   mintOnChain.RefId,
		ExecutionStatus: EXECUTION_STATUS_PENDING,
		Fee:             ev.Fee.String(),
	}
	transfer := &Transfer{
		Amount: new(big.Int).SetBytes(mintOnChain.Amount).String(),
		Token:  mintOnChain.Token,
	}
	execCtx := &ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	execCtx.MessageId = execCtx.ComputeMessageId(bridge)
	return execCtx
}

func (c *ExecutionContext) MustMarshal() []byte {
	data, err := c.Marshal()
	if err != nil {
		log.Panicf("failed to marshal execCtx %+v", c)
	}
	return data
}

func (c *ExecutionContext) ComputeMessageId(bridgeAddr eth.Addr) []byte {
	msg := c.Message
	if msg.TransferType == TRANSFER_TYPE_NULL {
		return msg.ComputeMessageIdNoTransfer()
	}
	return c.ComputeMessageIdWithTransfer(bridgeAddr)
}

func (c *ExecutionContext) ComputeMessageIdWithTransfer(dstBridgeAddr common.Address) []byte {
	var dstTransferId []byte
	m := c.Message
	t := c.Transfer
	switch m.TransferType {
	case TRANSFER_TYPE_NULL:
		return nil
	case TRANSFER_TYPE_LIQUIDITY_SEND:
		log.Debugf("TransferType:%s, %s, %s, %x, %s, %d, %d, %x", m.TransferType, m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, m.TransferRefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"address", "address", "address", "uint256", "uint64", "uint64", "bytes32"},
			m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, m.TransferRefId,
		)
	case TRANSFER_TYPE_LIQUIDITY_WITHDRAW:
		log.Debugf("TransferType:%s, %s, %s, %x, %s, %d, %d, %x", m.TransferType, m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, m.TransferRefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"uint64", "uint64", "address", "address", "uint256"},
			m.DstChainId, t.SeqNum, m.Receiver, t.Token, t.Amount,
		)
	case TRANSFER_TYPE_PEG_MINT, TRANSFER_TYPE_PEG_WITHDRAW:
		log.Debugf("TransferType:%s, %s, %x, %s, %s, %d, %x", m.TransferType, m.Receiver, t.Token, t.Amount, m.Sender, m.SrcChainId, m.TransferRefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"address", "address", "uint256", "address", "uint64", "bytes32"},
			m.Receiver, t.Token, t.Amount, m.Sender, m.SrcChainId, m.TransferRefId,
		)
	}
	return ComputeMessageIdFromDstTransfer(dstTransferId, dstBridgeAddr)
}

func ComputeMessageIdFromDstTransfer(dstTransferId []byte, dstBridgeAddr common.Address) []byte {
	// Prepend bridge address and hash again
	msgId := solsha3.SoliditySHA3(
		[]string{"uint8", "address", "bytes32"},
		uint8(MsgType_MSG_TYPE_MESSAGE_WITH_TRANSFER), dstBridgeAddr, dstTransferId,
	)
	log.Debugf("ComputeMessageIdFromDstTransferId, dstTransferId %x, dstBridgeAddr %x, messageId %x",
		dstTransferId, dstBridgeAddr, msgId)
	return msgId
}
