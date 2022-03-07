package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractType int

const (
	LiquidityBridge ContractType = iota
	PegVault
	PegBridge
	WdInbox
	MsgBridge
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
func GetEventSignature(eventSigStr string) Hash {
	return crypto.Keccak256Hash([]byte(eventSigStr))
}

// GetTxSender returns the sender address of the given transaction
func GetTxSender(ec *ethclient.Client, txHashStr string) (string, error) {
	tx, _, err := ec.TransactionByHash(context.Background(), Hex2Hash(txHashStr))
	if err != nil {
		return "", fmt.Errorf("failed to get tx: %w", err)
	}
	msg, err := tx.AsMessage(ethtypes.NewLondonSigner(tx.ChainId()), nil) // TODO: base fee
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

func SignerBytes(addrs []Addr, powers []*big.Int) []byte {
	var packed []byte
	for _, addr := range addrs {
		packed = append(packed, Pad32Bytes(addr.Bytes())...)
	}
	for _, power := range powers {
		packed = append(packed, Pad32Bytes(power.Bytes())...)
	}
	return packed
}

// given evname like LiquidityAdded, return its event ID, aka. topics[0]
// if evname not found, all 0 hash (default value) will be returned
// as this func parse abi internally, caller should call once and save the return
// instead of keep calling it.
func GetContractEventID(ctype ContractType, evname string) Hash {
	var contractAbi abi.ABI
	switch ctype {
	case LiquidityBridge:
		contractAbi, _ = abi.JSON(strings.NewReader(BridgeABI))
	case PegVault:
		contractAbi, _ = abi.JSON(strings.NewReader(OriginalTokenVaultABI))
	case PegBridge:
		contractAbi, _ = abi.JSON(strings.NewReader(PeggedTokenBridgeABI))
	case WdInbox:
		contractAbi, _ = abi.JSON(strings.NewReader(WithdrawInboxABI))
	case MsgBridge:
		contractAbi, _ = abi.JSON(strings.NewReader(MessageBusABI))
	default:
		return ZeroHash
	}
	return contractAbi.Events[evname].ID
}

// given list of logs, find matching event (log.topics[0] == GetBridgeEventID(cbrEvName) && log.Address == expAddr)
// from last to first, return first matched log. if evname not found in GetBridgeEventID or no match, return nil
// if found, return pointer from logs directly so be careful not changing logs after this call
// per eth design, event ID must match event topics[0]
// We MUST be extra careful dealing with log as attacker could generate same topics using their own contract
// why search backwards in logs: we were assuming our event is the last so just do receipt.Logs[len(receipt.Logs)-1],
// but Polygon adds its own event and breaks this assumption. So now we go backwards and search for first matched event.
// WARNING: must check log Address!!! other projects have been hacked by missing the check
func FindMatchContractEvent(ctype ContractType, evName string, expAddr Addr, logs []*ethtypes.Log) *ethtypes.Log {
	evID := GetContractEventID(ctype, evName)
	if evID == ZeroHash {
		return nil
	}
	for idx := len(logs) - 1; idx >= 0; idx-- {
		if logs[idx].Topics[0] == evID {
			// event ID matches and from expected contract
			if logs[idx].Address == expAddr {
				return logs[idx]
			} else {
				log.Warnln("topic match but contract addr mismatch, hack or misconfig. log has:", logs[idx].Address, "expect:", expAddr)
				return nil
			}
		}
	}
	// go over all logs, no match
	return nil
}

// return human friendly string for logging
func (ev *BridgeSend) PrettyLog(srcChid uint64) string {
	// max slippage uint is float * 1e6 so percentage needs to divide by 1e4
	return fmt.Sprintf("send-%x src: %d-%x dstchid: %d sender: %x receiver: %x amt: %s maxslip: %f%%", ev.TransferId, srcChid, ev.Token, ev.DstChainId, ev.Sender, ev.Receiver, ev.Amount, float64(ev.MaxSlippage)/10000)
}

func (ev *BridgeSend) CalcXferId(srcChid uint64) Hash {
	var b []byte
	b = append(b, ev.Sender[:]...)
	b = append(b, ev.Receiver[:]...)
	b = append(b, ev.Token[:]...)
	b = append(b, ToPadBytes(ev.Amount)...)
	b = append(b, ToPadBytes(ev.DstChainId)...)
	b = append(b, ToPadBytes(ev.Nonce)...)
	// old contract uses uint256, new contract uses uint64
	// b = append(b, ToPadBytes(big.NewInt(int64(srcChid)))...)
	b = append(b, ToPadBytes(srcChid)...)
	return Bytes2Hash(crypto.Keccak256(b))
}

func GetRelayTransferId(sender, receiver, token Addr, amount *big.Int, srcChainId, destChainId uint64, srcTransferId Hash) Hash {
	var b []byte
	b = append(b, sender[:]...)
	b = append(b, receiver[:]...)
	b = append(b, token[:]...)
	b = append(b, ToPadBytes(amount)...)
	b = append(b, ToPadBytes(srcChainId)...)
	b = append(b, ToPadBytes(destChainId)...)
	b = append(b, srcTransferId[:]...)
	return Bytes2Hash(crypto.Keccak256(b))
}

/*
bytes32 wdId = keccak256(
  abi.encodePacked(wdmsg.chainid, wdmsg.seqnum, wdmsg.receiver, wdmsg.token, wdmsg.amount)
);
*/
func (ev *BridgeWithdrawDone) CalcWdID(chid uint64) Hash {
	var b []byte
	b = append(b, ToPadBytes(chid)...)
	b = append(b, ToPadBytes(ev.Seqnum)...)
	b = append(b, ev.Receiver[:]...)
	b = append(b, ev.Token[:]...)
	b = append(b, ToPadBytes(ev.Amount)...)
	return Bytes2Hash(crypto.Keccak256(b))
}

// chid is the chain id we saw this event
/*
bytes32 transferId = keccak256(abi.encodePacked(
    request.sender, request.receiver, request.token, request.amount,
    request.srcChainId, request.dstChainId, request.srcTransferId));
*/
func (ev *BridgeRelay) CalcXferId(chid uint64) Hash {
	var b []byte
	b = append(b, ev.Sender[:]...)
	b = append(b, ev.Receiver[:]...)
	b = append(b, ev.Token[:]...)
	b = append(b, ToPadBytes(ev.Amount)...)
	b = append(b, ToPadBytes(ev.SrcChainId)...)
	b = append(b, ToPadBytes(chid)...)
	b = append(b, ev.SrcTransferId[:]...)
	return Bytes2Hash(crypto.Keccak256(b))
}

func (ev *BridgeLiquidityAdded) Equal(b *BridgeLiquidityAdded) bool {
	if ev.Seqnum != b.Seqnum {
		return false
	}
	if ev.Provider != b.Provider {
		return false
	}
	if ev.Token != b.Token {
		return false
	}
	if ev.Amount.Cmp(b.Amount) != 0 {
		return false
	}
	return true
}

func (ev *MessageBusMessage) Equal(b *MessageBusMessage) bool {
	if ev.Sender != b.Sender {
		return false
	}
	if ev.DstChainId.Cmp(b.DstChainId) != 0 {
		return false
	}
	if ev.Receiver != b.Receiver {
		return false
	}

	if bytes.Compare(ev.Message, b.Message) != 0 {
		return false
	}
	if ev.Fee.Cmp(b.Fee) != 0 {
		return false
	}
	return true
}

func (ev *MessageBusMessageWithTransfer) Equal(b *MessageBusMessageWithTransfer) bool {
	if ev.Bridge != b.Bridge {
		return false
	}
	if ev.SrcTransferId != b.SrcTransferId {
		return false
	}
	if ev.Sender != b.Sender {
		return false
	}
	if ev.DstChainId.Cmp(b.DstChainId) != 0 {
		return false
	}
	if ev.Receiver != b.Receiver {
		return false
	}
	if bytes.Compare(ev.Message, b.Message) != 0 {
		return false
	}
	if ev.Fee.Cmp(b.Fee) != 0 {
		return false
	}
	return true
}

func (ev *MessageBusExecuted) Equal(b *MessageBusExecuted) bool {
	if ev.Id != b.Id {
		return false
	}
	if ev.Status != b.Status {
		return false
	}
	if ev.MsgType != b.MsgType {
		return false
	}
	return true
}

// onchid is the chainid this event happen
func (ev *BridgeLiquidityAdded) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("cbr-liqadd-%d-%d token: %x lp: %x amt: %s", onchid, ev.Seqnum, ev.Token, ev.Provider, ev.Amount)
}

func (ev *BridgeWithdrawDone) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("cbr-withdraw-%d chid: %d token: %x receiver: %x amt: %s", ev.Seqnum, onchid, ev.Token, ev.Receiver, ev.Amount)
}

// relay-%x is src transfer id!!! so we can easily correlate with send log
func (ev *BridgeRelay) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("cbr-relay-%x srcchid: %d dst: %d-%x sender: %x receiver: %x amt: %s thisXferId: %x", ev.SrcTransferId, ev.SrcChainId, onchid, ev.Token, ev.Sender, ev.Receiver, ev.Amount, ev.TransferId)
}

func (ev *BridgeSignersUpdated) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("cbr-signersUpdated-%d: %s", onchid, ev.String())
}

func (ev *MessageBusMessage) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("Message-%d: %s", onchid, ev.String())
}

func (ev *MessageBusMessageWithTransfer) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("MessageWithTransfer-%d: %s", onchid, ev.String())
}

func (ev *MessageBusExecuted) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("MessageExecuted-%d: %s", onchid, ev.String())
}

func (ev *MessageBusExecuted) String() string {
	return fmt.Sprintf("id: %s,status: %d, msgType: %d", Hash(ev.Id), ev.Status, ev.MsgType)
}

func (ev *MessageBusMessage) String() string {
	return fmt.Sprintf("sender: %s, receiver: %s, dstChainId: %s, Message: %x", ev.Sender, ev.Receiver, ev.DstChainId, ev.Message)
}

func (ev *MessageBusMessageWithTransfer) String() string {
	return fmt.Sprintf("sender: %s, receiver: %s, dstChainId: %s, bridgeAddr: %s, transferId: %x, Message: %x", ev.Sender, ev.Receiver, ev.DstChainId, ev.Bridge, ev.SrcTransferId, ev.Message)
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

func (s *BridgeSignersUpdated) String() string {
	var out string
	for i, addr := range s.Signers {
		out += fmt.Sprintf("<addr %x power %s> ", addr, s.Powers[i])
	}
	return fmt.Sprintf("< %s>", out)
}

// ToPadBytes return big-endian/network order bytes, left padded to specific length
// if v is uint32: 4 bytes, int64/uint64: 8 bytes, *big.Int: 32 bytes or rlen bytes if set
// return nil if type not supported
func ToPadBytes(v interface{}, rlen ...int) []byte {
	var orig []byte
	var retlen int
	switch k := v.(type) {
	case uint32:
		retlen = 4
		orig = big.NewInt(int64(k)).Bytes()
	case int64:
		retlen = 8
		orig = big.NewInt(k).Bytes()
	case uint64:
		retlen = 8
		orig = new(big.Int).SetUint64(k).Bytes()
	case *big.Int:
		if len(rlen) == 1 {
			retlen = rlen[0]
		} else {
			retlen = 32
		}
		orig = k.Bytes()
	default:
		return nil
	}
	ret := make([]byte, retlen)
	copy(ret[retlen-len(orig):], orig)
	return ret
}

func (ev *OriginalTokenVaultDeposited) Equal(d *OriginalTokenVaultDeposited) bool {
	if ev.DepositId != d.DepositId {
		return false
	}
	if ev.Depositor != d.Depositor {
		return false
	}
	if ev.Token != d.Token {
		return false
	}
	if ev.Amount.Cmp(d.Amount) != 0 {
		return false
	}
	if ev.MintChainId != d.MintChainId {
		return false
	}
	if ev.MintAccount != d.MintAccount {
		return false
	}
	return true
}

func (ev *PeggedTokenBridgeBurn) Equal(b *PeggedTokenBridgeBurn) bool {
	if ev.BurnId != b.BurnId {
		return false
	}
	if ev.Account != b.Account {
		return false
	}
	if ev.Token != b.Token {
		return false
	}
	if ev.Amount.Cmp(b.Amount) != 0 {
		return false
	}
	if ev.WithdrawAccount != b.WithdrawAccount {
		return false
	}
	return true
}

// onchid is the chainid this event happen
func (ev *OriginalTokenVaultDeposited) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-deposited-%d depositId: %x account: %x token: %x amount: %s mintChainId: %d mintAccount: %x",
		onchid, ev.DepositId, ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount)
}

func (ev *PeggedTokenBridgeMint) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-mint-%d mintId: %x token: %x account: %x amount: %s depositChainId: %d depositId: %x depositor: %x",
		onchid, ev.MintId, ev.Token, ev.Account, ev.Amount, ev.RefChainId, ev.RefId, ev.Depositor)
}

func (ev *PeggedTokenBridgeBurn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-burn-%d burnId: %x token: %x account: %x amount: %s withdrawAccount: %x",
		onchid, ev.BurnId, ev.Token, ev.Account, ev.Amount, ev.WithdrawAccount)
}

func (ev *OriginalTokenVaultWithdrawn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-withdrawn-%d withdrawId: %x receiver: %x token: %x amount: %s burnChainId: %d burnId: %x burnAccount: %x",
		onchid, ev.WithdrawId, ev.Receiver, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.BurnAccount)
}

func (ev *OriginalTokenVaultDeposited) String() string {
	return fmt.Sprintf("depositId %x, account %x, token %x, amount %s, mintChainId %d,  mintAccount %x",
		ev.DepositId, ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount)
}

func (ev *PeggedTokenBridgeMint) String() string {
	return fmt.Sprintf("mintId %x, account %x, token %x, amount %s, depositChainId %d, depositId %x, depositor %x",
		ev.MintId, ev.Account, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.Depositor)
}

func (ev *PeggedTokenBridgeBurn) String() string {
	return fmt.Sprintf("burnId %x, account %x, token %x, amount %s, withdrawAccount %x",
		ev.BurnId, ev.Account, ev.Token, ev.Amount, ev.WithdrawAccount)
}

func (ev *OriginalTokenVaultWithdrawn) String() string {
	return fmt.Sprintf("withdrawId %x, receiver %x, token %x, amount %s, burnChainId %d, burnId %x, burnAccount %x",
		ev.WithdrawId, ev.Receiver, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.BurnAccount)
}

func (ev *OriginalTokenVaultV2Deposited) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-deposited2-%d depositId: %x account: %x token: %x amount: %s mintChainId: %d mintAccount: %x",
		onchid, ev.DepositId, ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount)
}

func (ev *PeggedTokenBridgeV2Mint) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-mint2-%d mintId: %x token: %x account: %x amount: %s depositChainId: %d depositId: %x depositor: %x",
		onchid, ev.MintId, ev.Token, ev.Account, ev.Amount, ev.RefChainId, ev.RefId, ev.Depositor)
}

func (ev *PeggedTokenBridgeV2Burn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-burn2-%d burnId: %x token: %x account: %x amount: %s toChainId %d toAccount: %x",
		onchid, ev.BurnId, ev.Token, ev.Account, ev.Amount, ev.ToChainId, ev.ToAccount)
}

func (ev *OriginalTokenVaultV2Withdrawn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-withdrawn2-%d withdrawId: %x receiver: %x token: %x amount: %s mintChainId: %d mintId: %x burnAccount: %x",
		onchid, ev.WithdrawId, ev.Receiver, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.BurnAccount)
}

func (ev *OriginalTokenVaultV2Deposited) String() string {
	return fmt.Sprintf("depositId %x, account %x, token %x, amount %s, mintChainId %d, mintAccount %x, nonce %d",
		ev.DepositId, ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount, ev.Nonce)
}

func (ev *PeggedTokenBridgeV2Mint) String() string {
	return fmt.Sprintf("mintId %x, account %x, token %x, amount %s, depositChainId %d, depositId %x, depositor %x",
		ev.MintId, ev.Account, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.Depositor)
}

func (ev *PeggedTokenBridgeV2Burn) String() string {
	return fmt.Sprintf("burnId %x, account %x, token %x, amount %s, toChainId %d, toAccount %x, nonce %d",
		ev.BurnId, ev.Account, ev.Token, ev.Amount, ev.ToChainId, ev.ToAccount, ev.Nonce)
}

func (ev *OriginalTokenVaultV2Withdrawn) String() string {
	return fmt.Sprintf("withdrawId %x, receiver %x, token %x, amount %s, burnChainId %d, burnId %x, burnAccount %x",
		ev.WithdrawId, ev.Receiver, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.BurnAccount)
}

func (ev *WithdrawInboxWithdrawalRequest) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"wdi-withdrawalrequest-%d sender: %x receiver: %x toChain: %d fromChains: %v tokens: %v ratios: %v slippages: %v deadline: %s",
		onchid, ev.Sender, ev.Receiver, ev.ToChain, ev.FromChains, ev.Tokens, ev.Ratios, ev.Slippages, time.Unix(ev.Deadline.Int64(), 0))
}

func (ev *WithdrawInboxWithdrawalRequest) String() string {
	return fmt.Sprintf(
		"sender: %x receiver: %x toChain: %d fromChains: %v tokens: %v ratios: %v slippages: %v deadline: %s",
		ev.Sender, ev.Receiver, ev.ToChain, ev.FromChains, ev.Tokens, ev.Ratios, ev.Slippages, time.Unix(ev.Deadline.Int64(), 0))
}

func (ev *WithdrawInboxWithdrawalRequest) Equal(b *WithdrawInboxWithdrawalRequest) bool {
	if ev.SeqNum != b.SeqNum {
		return false
	}
	if ev.Sender != b.Sender {
		return false
	}
	if ev.Receiver != b.Receiver {
		return false
	}
	if ev.ToChain != b.ToChain {
		return false
	}
	if ev.Deadline.Cmp(b.Deadline) != 0 {
		return false
	}
	if len(ev.FromChains) != len(b.FromChains) {
		return false
	}
	if !(len(ev.FromChains) == len(ev.Tokens) &&
		len(ev.FromChains) == len(ev.Ratios) &&
		len(ev.FromChains) == len(ev.Slippages)) {
		return false
	}
	for i, fromChain := range ev.FromChains {
		if fromChain != b.FromChains[i] {
			return false
		}
		if ev.Tokens[i] != b.Tokens[i] {
			return false
		}
		if ev.Ratios[i] != b.Ratios[i] {
			return false
		}
		if ev.Slippages[i] != b.Slippages[i] {
			return false
		}
	}
	return true
}
