package types

import (
	fmt "fmt"
	"math/big"
	"strings"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/sha3"
)

func NewMessage(ev *eth.MessageBusMessage, srcChainId uint64) (messageId eth.Hash, message *Message) {
	message = &Message{
		SrcChainId: srcChainId,
		Sender:     eth.Addr2Hex(ev.Sender),
		DstChainId: ev.DstChainId.Uint64(),
		Receiver:   eth.Addr2Hex(ev.Receiver),
		Data:       ev.Message,
		Fee:        ev.Fee.String(),
		SrcTxHash:  ev.Raw.TxHash.Hex(),
	}
	messageId = eth.Bytes2Hash(message.ComputeMessageIdNoTransfer())
	return
}

func (m *Message) GetSignerAddrs() []common.Address {
	if m == nil {
		return nil
	}
	signers := []common.Address{}
	for _, sig := range m.Signatures {
		signers = append(signers, eth.Hex2Addr(sig.Signer))
	}
	return signers
}

func (m *Message) GetSigBytes() [][]byte {
	sigBytes := [][]byte{}
	for _, sig := range m.Signatures {
		sigBytes = append(sigBytes, sig.SigBytes)
	}
	return sigBytes
}

func (m *Message) EncodeDataToSign(messageId eth.Hash, messageBusAddr common.Address) []byte {
	// refund msg
	if m.SrcChainId == m.DstChainId {
		domain := solsha3.SoliditySHA3(
			[]string{"uint256", "address", "string"},
			new(big.Int).SetUint64(m.DstChainId), messageBusAddr, "MessageWithTransferRefund",
		)
		data := append(domain, messageId.Bytes()...)
		return append(append(data, m.Data...), eth.Hex2Hash(m.SrcTxHash).Bytes()...)
	}
	// normal msg
	if m.GetTransferType() == TRANSFER_TYPE_NULL {
		domain := solsha3.SoliditySHA3(
			[]string{"uint256", "address", "string"},
			new(big.Int).SetUint64(m.DstChainId), messageBusAddr, "Message",
		)
		return append(domain, messageId.Bytes()...)
	} else {
		domain := solsha3.SoliditySHA3(
			[]string{"uint256", "address", "string"},
			new(big.Int).SetUint64(m.DstChainId), messageBusAddr, "MessageWithTransfer",
		)
		data := append(domain, messageId.Bytes()...)
		return append(append(data, m.Data...), eth.Hex2Hash(m.SrcTxHash).Bytes()...)
	}
}

func (m *Message) AddSig(data []byte, sigBytes []byte, expectedSigner common.Address) error {
	newSigs, err := commontypes.AddSig(m.GetSignatures(), data, sigBytes, expectedSigner.Hex())
	if err != nil {
		return err
	}
	m.Signatures = newSigs
	return nil
}

func (m *Message) ComputeMessageIdNoTransfer() []byte {
	data := solsha3.Pack(
		[]string{"uint8", "address", "address", "uint64", "bytes32", "uint64"},
		[]interface{}{uint8(MsgType_MSG_TYPE_MESSAGE), m.Sender, m.Receiver, m.SrcChainId, eth.Hex2Hash(m.SrcTxHash), m.DstChainId},
	)
	// NOTE: Manual concatenation as solsha3 DOES NOT SUPPORT dynamic "bytes"
	data = append(data, m.Data...)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

func (m *Message) MapOnChainStatus(status eth.MessageReceiverStatus) ExecutionStatus {
	switch status {
	case eth.MessageReceiverTxStatusSuccess:
		return EXECUTION_STATUS_SUCCESS
	case eth.MessageReceiverTxStatusFail:
		return EXECUTION_STATUS_FAILURE
	case eth.MessageReceiverTxStatusFallback:
		return EXECUTION_STATUS_FALLBACK
	default:
		return EXECUTION_STATUS_NULL
	}
}

func (m *Message) PrettyLog() {
	log.Debugln("message:")
	log.Debugf("sender %s", m.Sender)
	log.Debugf("receiver %s", m.Receiver)
	log.Debugf("srcChainId %d", m.SrcChainId)
	log.Debugf("dstChainId %d", m.DstChainId)
	log.Debugf("transferType %v", m.TransferType)
	log.Debugf("executionStatus %v", m.ExecutionStatus)
	log.Debugf("lastSigReqTime %d", m.LastSigReqTime)
	var sigstr []string
	for _, sig := range m.Signatures {
		sigstr = append(sigstr, fmt.Sprintf("%s:%x", sig.Signer, sig.SigBytes))
	}
	log.Debugf("signatures %s", strings.Join(sigstr, ", "))
	log.Debugf("data %x", m.Data)
}
