package types

import (
	"fmt"
	"math/big"
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
	// todo: make sure assets are multi-chain correctly?
	// todo: also check all chains in assets are in chainpairs
	return nil
}

func (ast *ChainAsset) Validate() error {
	_, good := new(big.Int).SetString(ast.MaxFeeAmount, 10)
	if !good {
		return fmt.Errorf("max_fee_amount %s bad format", ast.MaxFeeAmount)
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
	return nil
}

func (r *RelayOnChain) String() string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("sender %x, receiver %x, token %x, amount %s, src_chain_id %d, dst_chain_id %d, src_xfer_id %x",
		r.Sender, r.Receiver, r.Token, big.NewInt(0).SetBytes(r.Amount), r.SrcChainId, r.DstChainId, r.SrcTransferId)
}

func (w *WithdrawOnchain) String() string {
	if w == nil {
		return ""
	}
	return fmt.Sprintf("chainid %d, seqnum %d, receiver %x, token %x, amount %d",
		w.Chainid, w.Seqnum, w.Receiver, w.Token, big.NewInt(0).SetBytes(w.Amount))
}
