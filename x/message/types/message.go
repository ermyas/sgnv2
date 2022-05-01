package types

import (
	fmt "fmt"
	"math/big"
	"strings"

	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
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

func (m *Message) GetSignerAddrs() []eth.Addr {
	if m == nil {
		return nil
	}
	signers := []eth.Addr{}
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

func (m *Message) EncodeDataToSign(messageId eth.Hash, messageBusAddr eth.Addr) []byte {
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

func (m *Message) AddSig(data []byte, sigBytes []byte, expectedSigner eth.Addr) error {
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

func (m *Message) PrettyLog() string {
	var sigstr []string
	for _, sig := range m.Signatures {
		sigstr = append(sigstr, fmt.Sprintf("%s:%x", sig.Signer, sig.SigBytes))
	}
	return fmt.Sprintf("message: chain %d->%d addr %s->%s type %s status %s lastReqTime %s data %x sigs %s",
		m.SrcChainId, m.DstChainId, m.Sender, m.Receiver, m.TransferType, m.ExecutionStatus,
		common.TsSecToTime(uint64(m.LastSigReqTime)), m.Data, strings.Join(sigstr, ", "))
}

func (m *Message) PrettyPrint() {
	if m == nil {
		fmt.Println("nil message")
	}
	fmt.Printf("src_chain_id: %d\n", m.SrcChainId)
	fmt.Printf("dst_chain_id: %d\n", m.DstChainId)
	fmt.Printf("sender: %s\n", m.Sender)
	fmt.Printf("receiver: %s\n", m.Sender)
	fmt.Printf("fee: %s\n", m.Fee)
	fmt.Printf("transfer_type: %s\n", m.TransferType.String())
	fmt.Printf("transfer_ref_id: %x\n", m.TransferRefId)
	fmt.Printf("src_tx_hash: %s\n", m.SrcTxHash)
	fmt.Printf("execution_status: %s\n", m.ExecutionStatus.String())
	fmt.Printf("last_sig_req_time: %s\n", common.TsSecToTime(uint64(m.LastSigReqTime)))
	var signers []string
	for _, sig := range m.Signatures {
		signers = append(signers, sig.Signer)
	}
	fmt.Println("signers:", signers)
	fmt.Printf("data: 0x%x\n", m.Data)
}
