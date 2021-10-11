package eth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

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

// return human friendly string for logging
func (ev *BridgeSend) PrettyLog(srcChid uint64) string {
	// max slippage uint is float * 1e6 so percentage needs to divide by 1e4
	return fmt.Sprintf("send-%x src: %d-%x dstchid: %d sender: %x receiver: %x amt: %s maxslip: %f%%", ev.TransferId, srcChid, ev.Token, ev.DstChainId, ev.Sender, ev.Receiver, ev.Amount, float64(ev.MaxSlippage)/10000)
}

// calculate xfer id the same way as Bridge.sol
// todo: change chainid to match latest contract after new deploy
/*
bytes32 transferId = keccak256(
// uint64(block.chainid) for consistency as entire system uses uint64 for chain id
  abi.encodePacked(msg.sender, _receiver, _token, _amount, _dstChainId, _nonce, uint64(block.chainid))
);
*/

func (ev *BridgeSend) CalcXferId(srcChid uint64) HashType {
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

/*
bytes32 wdId = keccak256(
  abi.encodePacked(wdmsg.chainid, wdmsg.seqnum, wdmsg.receiver, wdmsg.token, wdmsg.amount)
);
*/
func (ev *BridgeWithdrawDone) CalcWdID(chid uint64) HashType {
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
func (ev *BridgeRelay) CalcXferId(chid uint64) HashType {
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

func (ev *BridgeSignersUpdated) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("signersUpdated-%d: %s", onchid, ev.String())
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
