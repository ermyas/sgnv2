package eth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrPeersNotMatch = errors.New("channel peers not match")
)

func ParseValStatus(valStatus uint8) string {
	switch valStatus {
	case Bonded:
		return "bonded"
	case Unbonding:
		return "unbonding"
	case Unbonded:
		return "unbonded"
	}

	return "null"
}

// GetEventSignature accepts the string of an event signature and return the hex
func GetEventSignature(eventSigStr string) HashType {
	return crypto.Keccak256Hash([]byte(eventSigStr))
}

// GetTxSender returns the sender address of the given transaction
func GetTxSender(ec *ethclient.Client, txHashStr string) (string, error) {
	tx, _, err := ec.TransactionByHash(context.Background(), Hex2Hash(txHashStr))
	if err != nil {
		return "", fmt.Errorf("failed to get tx: %w", err)
	}
	msg, err := tx.AsMessage(ethtypes.NewLondonSigner(tx.ChainId()), nil) //TODO: base fee
	if err != nil {
		return "", fmt.Errorf("failed to get msg: %w", err)
	}
	return Addr2Hex(msg.From()), nil
}

func GetAddressFromKeystore(ksBytes []byte) (string, error) {
	type ksStruct struct {
		Address string
	}
	var ks ksStruct
	if err := json.Unmarshal(ksBytes, &ks); err != nil {
		return "", err
	}
	return ks.Address, nil
}

func CommissionRate(rate float64) uint64 {
	return uint64(rate * CommissionRateBase)
}

// return human friendly string for logging
func (ev *BridgeSend) PrettyLog(srcChid uint64) string {
	// max slippage uint is float * 1e6 so percentage needs to divide by 1e4
	return fmt.Sprintf("send-%x src: %d-%x dstchid: %d sender: %x receiver: %x amt: %s maxslip: %f%%", ev.TransferId, srcChid, ev.Token, ev.DstChainId, ev.Sender, ev.Receiver, ev.Amount, float64(ev.MaxSlippage)/10000)
}

// onchid is the chainid this event happen
func (ev *BridgeLiquidityAdded) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("liqadd-%d-%d token: %x lp: %x amt: %s", onchid, ev.Seqnum, ev.Token, ev.Provider, ev.Amount)
}

func (ev *BridgeWithdrawDone) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("withdraw-%d chid: %d token: %x receiver: %x amt: %s", ev.Seqnum, onchid, ev.Token, ev.Receiver, ev.Amount)
}

// relay-%x is src transfer id!!! so we can easily correlate with send log
func (ev *BridgeRelay) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("relay-%x srcchid: %d dst: %d-%x sender: %x receiver: %x amt: %s thisXferId: %x", ev.SrcTransferId, ev.SrcChainId, onchid, ev.Token, ev.Sender, ev.Receiver, ev.Amount, ev.TransferId)
}

func (r *BridgeRelay) String() string {
	return fmt.Sprintf("transferId %x, sender %x, receiver %x, token %x, amount %s, srcChainId %d, srcTransferId %x",
		r.TransferId, r.Sender, r.Receiver, r.Token, r.Amount, r.SrcChainId, r.SrcTransferId)
}

func (s *BridgeSend) String() string {
	return fmt.Sprintf("transferId %x, sender %x, receiver %x, token %x, amount %s, dstChainId %d, nonce %d, maxSlippage %d",
		s.TransferId, s.Sender, s.Receiver, s.Token, s.Amount, s.DstChainId, s.Nonce, s.MaxSlippage)
}

func (l *BridgeLiquidityAdded) String() string {
	return fmt.Sprintf("seqNum %d, provider %x, token %x, amount %s", l.Seqnum, l.Provider, l.Token, l.Amount)
}

func (w *BridgeWithdrawDone) String() string {
	return fmt.Sprintf("withdrawId %x, seqNum %d, receiver %x, token %x, amount %s",
		w.WithdrawId, w.Seqnum, w.Receiver, w.Token, w.Amount)

}
