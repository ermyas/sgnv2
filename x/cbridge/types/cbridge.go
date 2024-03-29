package types

import (
	"fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func (m *XferRelay) GetSortedSigsBytes() [][]byte {
	if m != nil {
		sigs := make([][]byte, len(m.SortedSigs))
		for i := range m.SortedSigs {
			sigs[i] = m.SortedSigs[i].Sig
		}
		return sigs
	}
	return nil
}

func (m *XferRelay) SignersStr() string {
	var signers string
	for _, s := range m.SortedSigs {
		signers += fmt.Sprintf("%x ", s.Addr)
	}
	return fmt.Sprintf("signers: < %s>", signers)
}

func (m *WithdrawDetail) GetSortedSigsBytes() [][]byte {
	if m != nil {
		sigs := make([][]byte, len(m.SortedSigs))
		for i := range m.SortedSigs {
			sigs[i] = m.SortedSigs[i].Sig
		}
		return sigs
	}
	return nil
}

func (m *WithdrawDetail) SignersStr() string {
	var signers string
	for _, s := range m.SortedSigs {
		signers += fmt.Sprintf("%x ", s.Addr)
	}
	return fmt.Sprintf("signers: < %s>", signers)
}

// basic check of config
func (c *CbrConfig) Validate() error {
	if c.LpFeePerc > 100 {
		return fmt.Errorf("lp_fee_perc %d > 100", c.LpFeePerc)
	}
	var err error
	for _, ast := range c.Assets {
		err = ast.Validate()
		if err != nil {
			return fmt.Errorf("asset %s invalid: %w", ast.String(), err)
		}
	}
	for _, cp := range c.ChainPairs {
		err = cp.Validate()
		if err != nil {
			return fmt.Errorf("chainpair %s invalid: %w", cp.String(), err)
		}
	}
	for _, ov := range c.Override {
		err = ov.Chpair.Validate()
		if err != nil {
			return fmt.Errorf("%s override %s invalid: %w", ov.Symbol, ov.Chpair.String(), err)
		}
	}
	// todo: make sure assets are multi-chain correctly?
	// also check all chains in assets are in chainpairs
	return nil
}

func (ast *ChainAsset) Validate() error {
	_, good := new(big.Int).SetString(ast.MaxFeeAmount, 10)
	if !good {
		return fmt.Errorf("max_fee_amount %s bad format", ast.MaxFeeAmount)
	}
	if ast.MaxOutAmt != "" {
		maxSend, good := new(big.Int).SetString(ast.MaxOutAmt, 10)
		if !good || maxSend.Sign() == -1 {
			return fmt.Errorf("max_out_amt %s should be a non-negative number", ast.MaxOutAmt)
		}
	}
	return nil
}

func (cp *ChainPair) Validate() error {
	if cp.Chid1 > cp.Chid2 {
		return fmt.Errorf("chid1 %d > chid2 %d", cp.Chid1, cp.Chid2)
	}
	if cp.Fee1To2 > 1e6 {
		return fmt.Errorf("Fee1To2 %d > 1e6", cp.Fee1To2)
	}
	if cp.Fee2To1 > 1e6 {
		return fmt.Errorf("Fee2To1 %d > 1e6", cp.Fee2To1)
	}
	if cp.Weight1 >= 200 {
		return fmt.Errorf("weight1 %d >= 200", cp.Weight1)
	}
	if cp.NoCurve {
		if cp.Weight1 > 0 || cp.ConstA > 0 {
			return fmt.Errorf("both no_curve and (weight1 %d or consta %d) are set", cp.Weight1, cp.ConstA)
		}
	}
	return nil
}

func (r *RelayOnChain) String() string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("sender %x, receiver %x, token %x, amount %s, src_chain_id %d, dst_chain_id %d, src_xfer_id %x",
		r.Sender, r.Receiver, r.Token, big.NewInt(0).SetBytes(r.Amount), r.SrcChainId, r.DstChainId, r.SrcTransferId)
}

// get transfer id of relay (dest transfer id)
func (ev *RelayOnChain) GetRelayOnChainTransferId() eth.Hash {
	return eth.GetRelayTransferId(
		eth.Bytes2Addr(ev.GetSender()),
		eth.Bytes2Addr(ev.GetReceiver()),
		eth.Bytes2Addr(ev.GetToken()),
		new(big.Int).SetBytes(ev.GetAmount()),
		ev.GetSrcChainId(),
		ev.GetDstChainId(),
		eth.Bytes2Hash(ev.GetSrcTransferId()))
}

func (w *WithdrawOnchain) String() string {
	if w == nil {
		return ""
	}
	return fmt.Sprintf("chainid %d, seqnum %d, receiver %x, token %x, amount %d, refid %x",
		w.Chainid, w.Seqnum, w.Receiver, w.Token, big.NewInt(0).SetBytes(w.Amount), w.Refid)
}

func EncodeRelayOnChainToSign(chainId uint64, contractAddr eth.Addr, relayBytes []byte) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(chainId), contractAddr, "Relay"},
	)
	return append(domain, relayBytes...)
}

func EncodeWithdrawOnchainToSign(chainId uint64, contractAddr eth.Addr, withdrawBytes []byte) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(chainId), contractAddr, "WithdrawMsg"},
	)
	return append(domain, withdrawBytes...)
}
