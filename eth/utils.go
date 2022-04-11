package eth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"

	flowtypes "github.com/celer-network/cbridge-flow/types"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractType int

const (
	ContractTypeLiquidityBridge ContractType = iota
	ContractTypePegVault
	ContractTypePegVaultV2
	ContractTypePegBridge
	ContractTypePegBridgeV2
	ContractTypeWdInbox
	ContractTypeMsgBus
)

var (
	ErrPeersNotMatch = errors.New("channel peers not match")

	EvIdCache     = make(map[string]Hash)
	EvIdCacheLock sync.RWMutex
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

func EvIdCacheKey(ctype ContractType, evname string) string {
	return fmt.Sprintf("%d-%s", ctype, evname)
}

func GetEvIdCache(ctype ContractType, evname string) Hash {
	EvIdCacheLock.RLock()
	defer EvIdCacheLock.RUnlock()
	if evId, ok := EvIdCache[EvIdCacheKey(ctype, evname)]; ok {
		return evId
	}
	return ZeroHash
}

func SetEvIdCache(ctype ContractType, evname string, evId Hash) {
	EvIdCacheLock.Lock()
	defer EvIdCacheLock.Unlock()
	EvIdCache[EvIdCacheKey(ctype, evname)] = evId
}

// given evname like LiquidityAdded, return its event ID, aka. topics[0]
// if evname not found, all 0 hash (default value) will be returned
// as this func parse abi internally, caller should call once and save the return
// instead of keep calling it.
func GetContractEventID(ctype ContractType, evname string) Hash {
	if evId := GetEvIdCache(ctype, evname); evId != ZeroHash {
		return evId
	}
	var contractAbi abi.ABI
	switch ctype {
	case ContractTypeLiquidityBridge:
		contractAbi, _ = abi.JSON(strings.NewReader(BridgeABI))
	case ContractTypePegVault:
		contractAbi, _ = abi.JSON(strings.NewReader(OriginalTokenVaultABI))
	case ContractTypePegBridge:
		contractAbi, _ = abi.JSON(strings.NewReader(PeggedTokenBridgeABI))
	case ContractTypePegVaultV2:
		contractAbi, _ = abi.JSON(strings.NewReader(OriginalTokenVaultV2ABI))
	case ContractTypePegBridgeV2:
		contractAbi, _ = abi.JSON(strings.NewReader(PeggedTokenBridgeV2ABI))
	case ContractTypeWdInbox:
		contractAbi, _ = abi.JSON(strings.NewReader(WithdrawInboxABI))
	case ContractTypeMsgBus:
		contractAbi, _ = abi.JSON(strings.NewReader(MessageBusABI))
	default:
		return ZeroHash
	}
	evId := contractAbi.Events[evname].ID
	SetEvIdCache(ctype, evname, evId)
	return evId
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
		if len(logs[idx].Topics) > 0 && logs[idx].Topics[0] == evID {
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

// set values from Flow event, WARNING ev.Token is ZeroAddr!!! because flowev's token is human string
func (ev *OriginalTokenVaultV2Deposited) SetByFlow(flowev *flowtypes.FlowSafeBoxDeposited) {
	ev.DepositId = flowev.DepositId
	ev.Depositor = flowev.Depositor
	ev.Amount = new(big.Int).Set(flowev.Amount)
	ev.MintChainId = flowev.MintChainId
	ev.MintAccount = flowev.MintAccount
}

// set values from Flow event, WARNING ev.Token is ZeroAddr!!! because flowev's token is human string
// TODO: to chainid isn't supported by flow yet
func (ev *PeggedTokenBridgeV2Burn) SetByFlow(flowev *flowtypes.FlowPegBridgeBurn) {
	ev.BurnId = flowev.BurnId
	ev.Account = flowev.Burner
	ev.Amount = new(big.Int).Set(flowev.Amount)
	ev.ToAccount = flowev.ToAccount
}
