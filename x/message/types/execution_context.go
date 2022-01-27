package types

import (
	fmt "fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func NewMsgExecutionContext(ev *eth.MessageBusMessage, srcChainId uint64) *ExecutionContext {
	message := Message{
		SrcChainId:      srcChainId,
		Sender:          ev.Sender.String(),
		DstChainId:      ev.DstChainId.Uint64(),
		Receiver:        ev.Receiver.String(),
		Data:            ev.Message,
		Fee:             "0", // ev.Fee.String(),
		ExecutionStatus: EXECUTION_STATUS_PENDING,
	}
	execCtx := &ExecutionContext{
		Message: message,
	}
	execCtx.MessageId = message.ComputeMessageIdNoTransfer()
	return execCtx
}

func NewMsgXferExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	chainId uint64,
	dstToken eth.Addr,
	dstAmt string,
	dstBridge eth.Addr,
	transferType TransferType) *ExecutionContext {

	message := Message{
		SrcChainId:      chainId,
		Sender:          ev.Sender.String(),
		DstChainId:      ev.DstChainId.Uint64(),
		Receiver:        ev.Receiver.String(),
		Data:            ev.Message,
		Fee:             "0", // ev.Fee.String(),
		TransferType:    transferType,
		ExecutionStatus: EXECUTION_STATUS_PENDING,
	}
	transfer := &Transfer{
		Token:  dstToken.Bytes(),
		Amount: dstAmt,
		RefId:  ev.SrcTransferId[:],
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

	transfer := &Transfer{
		Amount: new(big.Int).SetBytes(wdOnchain.Amount).String(),
		Token:  wdOnchain.Token,
		SeqNum: nonce,
		RefId:  ev.SrcTransferId[:],
	}
	message := &Message{
		SrcChainId:      wdOnchain.Chainid,
		DstChainId:      wdOnchain.Chainid,
		Sender:          eth.ZeroAddrHex,
		Receiver:        ev.Sender.String(),
		Data:            ev.Message,
		TransferType:    TRANSFER_TYPE_LIQUIDITY_WITHDRAW,
		ExecutionStatus: EXECUTION_STATUS_PENDING,
		Fee:             "0",
	}
	execCtx := &ExecutionContext{
		Message:  *message,
		Transfer: transfer,
	}
	execCtx.PopulateMessageId(bridge)
	return execCtx
}

func NewMsgPegDepositRefundExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	wdOnChain pegbrtypes.WithdrawOnChain,
	bridge string) *ExecutionContext {

	transfer := &Transfer{
		Amount: new(big.Int).SetBytes(wdOnChain.Amount).String(),
		Token:  wdOnChain.Token,
		RefId:  wdOnChain.RefId,
	}
	message := &Message{
		SrcChainId:      wdOnChain.RefChainId,
		DstChainId:      wdOnChain.RefChainId,
		Sender:          eth.ZeroAddrHex,
		Receiver:        eth.Bytes2AddrHex(wdOnChain.Receiver),
		Data:            ev.Message,
		TransferType:    TRANSFER_TYPE_PEG_WITHDRAW,
		ExecutionStatus: EXECUTION_STATUS_PENDING,
		Fee:             "0",
	}
	execCtx := &ExecutionContext{
		Message:  *message,
		Transfer: transfer,
	}
	execCtx.PopulateMessageId(eth.Hex2Addr(bridge))
	return execCtx
}

func NewMsgPegBurnRefundExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	mintOnChain pegbrtypes.MintOnChain,
	bridge string) *ExecutionContext {

	transfer := &Transfer{
		Amount: new(big.Int).SetBytes(mintOnChain.Amount).String(),
		Token:  mintOnChain.Token,
		RefId:  mintOnChain.RefId,
	}
	message := &Message{
		SrcChainId:      mintOnChain.RefChainId,
		DstChainId:      mintOnChain.RefChainId,
		Sender:          eth.ZeroAddrHex,
		Receiver:        eth.Bytes2AddrHex(mintOnChain.Account),
		Data:            ev.Message,
		TransferType:    TRANSFER_TYPE_PEG_MINT,
		ExecutionStatus: EXECUTION_STATUS_PENDING,
		Fee:             "0",
	}
	execCtx := &ExecutionContext{
		Message:  *message,
		Transfer: transfer,
	}
	execCtx.PopulateMessageId(eth.Hex2Addr(bridge))
	return execCtx
}

func (c *ExecutionContext) MustMarshal() []byte {
	data, err := c.Marshal()
	if err != nil {
		log.Panicf("failed to marshal execCtx %+v", c)
	}
	return data
}

func (c *ExecutionContext) GetSignerPowers() []*big.Int {
	powers := []*big.Int{}
	for _, power := range c.Powers {
		p, _ := new(big.Int).SetString(power, 10)
		powers = append(powers, p)
	}
	return powers
}

func (c *ExecutionContext) GetRefIdBytes32() ([32]byte, error) {
	var ret [32]byte
	t := c.Transfer
	if t == nil {
		return ret, nil
	}
	if len(t.RefId) != 32 {
		return ret, fmt.Errorf("cannot convert []byte to [32]byte: input []byte has length %d", len(t.RefId))
	}
	copy(ret[:], t.RefId)
	return ret, nil
}

func (c *ExecutionContext) computeMessageIdWithTransfer(dstBridgeAddr common.Address) []byte {
	var dstTransferId []byte
	m := c.Message
	t := c.Transfer
	switch m.TransferType {
	case TRANSFER_TYPE_NULL:
		return nil
	case TRANSFER_TYPE_LIQUIDITY_SEND:
		log.Debugf("TransferType:%s, %s, %s, %x, %s, %d, %d, %x", m.TransferType, m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, t.RefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"address", "address", "address", "uint256", "uint64", "uint64", "bytes32"},
			m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, t.RefId,
		)
	case TRANSFER_TYPE_LIQUIDITY_WITHDRAW:
		log.Debugf("TransferType:%s, %s, %s, %x, %s, %d, %d, %x", m.TransferType, m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, t.RefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"uint64", "uint64", "address", "address", "uint256"},
			m.DstChainId, t.SeqNum, m.Receiver, t.Token, t.Amount,
		)
	case TRANSFER_TYPE_PEG_MINT:
		fallthrough
	case TRANSFER_TYPE_PEG_WITHDRAW:
		log.Debugf("TransferType:%s, %s, %x, %s, %s, %d, %x", m.TransferType, m.Receiver, t.Token, t.Amount, m.Sender, m.SrcChainId, t.RefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"address", "address", "uint256", "address", "uint64", "bytes32"},
			m.Receiver, t.Token, t.Amount, m.Sender, m.SrcChainId, t.RefId,
		)
	}
	return ComputeMessageIdFromDstTransfer(dstTransferId, dstBridgeAddr)
}

func (c *ExecutionContext) PopulateMessageId(bridgeAddr eth.Addr) (messageId eth.Hash) {
	c.MessageId = c.ComputeMessageId(bridgeAddr)
	return eth.Bytes2Hash(c.MessageId)
}

func (c *ExecutionContext) ComputeMessageId(bridgeAddr eth.Addr) []byte {
	msg := c.Message
	if msg.TransferType == TRANSFER_TYPE_NULL {
		return msg.ComputeMessageIdNoTransfer()
	}
	return c.computeMessageIdWithTransfer(bridgeAddr)
}

func ComputeMessageIdFromDstTransfer(dstTransferId []byte, dstBridgeAddr common.Address) []byte {
	log.Infof("ComputeMessageIdFromDstTransferId, dstTransferId %x, dstBridgeAddr:%s", dstTransferId, dstBridgeAddr)
	// Prepend bridge address and hash again
	return solsha3.SoliditySHA3(
		[]string{"uint8", "address", "bytes32"},
		uint8(MsgType_MSG_TYPE_MESSAGE_WITH_TRANSFER), dstBridgeAddr, dstTransferId,
	)
}
