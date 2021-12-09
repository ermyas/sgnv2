package types

import (
	fmt "fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
)

func (m *MintInfo) GetAddrSigs() []*cbrtypes.AddrSig {
	// NOTE: Already sorted when adding signatures

	addrSigs := make([]*cbrtypes.AddrSig, 0)
	for _, sig := range m.Signatures {
		addrSigs = append(addrSigs, &cbrtypes.AddrSig{
			Addr: eth.Hex2Bytes(sig.Signer),
			Sig:  sig.SigBytes,
		})
	}

	return addrSigs
}

func (m *MintInfo) SignersStr() string {
	var signers string
	for _, s := range m.Signatures {
		signers += fmt.Sprintf("%s ", s.Signer)
	}
	return fmt.Sprintf("signers:< %s>", signers)
}

func (m *MintInfo) String() string {
	if m == nil {
		return "nil"
	}
	mintOnChain := new(MintOnChain)
	mintOnChain.Unmarshal(m.MintProtoBytes)
	return fmt.Sprintf("chain_id:%x mint_on_chain:[ %s ] base_fee:%s perc_fee:%s last_req_time:%d %s success:%t",
		m.ChainId, mintOnChain.String(), m.BaseFee, m.PercentageFee, m.LastReqTime, m.SignersStr(), m.Success)
}

func (w *WithdrawInfo) GetAddrSigs() []*cbrtypes.AddrSig {
	addrSigs := make([]*cbrtypes.AddrSig, 0)
	for _, sig := range w.Signatures {
		addrSigs = append(addrSigs, &cbrtypes.AddrSig{
			Addr: eth.Hex2Bytes(sig.Signer),
			Sig:  sig.SigBytes,
		})
	}

	return addrSigs
}

func (w *WithdrawInfo) SignersStr() string {
	var signers string
	for _, s := range w.Signatures {
		signers += fmt.Sprintf("%s ", s.Signer)
	}
	return fmt.Sprintf("signers:< %s>", signers)
}

func (w *WithdrawInfo) GetSortedSigsBytes() [][]byte {
	if w != nil {
		sigs := make([][]byte, len(w.Signatures))
		for i := range w.Signatures {
			sigs[i] = w.Signatures[i].SigBytes
		}
		return sigs
	}
	return nil
}

func (w *WithdrawInfo) String() string {
	if w == nil {
		return "nil"
	}
	wdOnChain := new(WithdrawOnChain)
	wdOnChain.Unmarshal(w.WithdrawProtoBytes)
	return fmt.Sprintf("chain_id:%x withdraw_on_chain:[ %s ] base_fee:%s perc_fee:%s last_req_time:%d %s success:%t",
		w.ChainId, wdOnChain.String(), w.BaseFee, w.PercentageFee, w.LastReqTime, w.SignersStr(), w.Success)
}

func (d *DepositInfo) String() string {
	if d == nil {
		return "nil"
	}
	return fmt.Sprintf("chain_id:%d deposit_id:%x mint_id:%x", d.ChainId, d.DepositId, d.MintId)
}

func (b *BurnInfo) String() string {
	if b == nil {
		return "nil"
	}
	return fmt.Sprintf("chain_id:%d burn_id:%x withdraw_id:%x", b.ChainId, b.BurnId, b.WithdrawId)
}

func (m *MintOnChain) String() string {
	if m == nil {
		return "nil"
	}
	return fmt.Sprintf("token:%x account:%x amount%s depositor:%x ref_chain_id:%x ref_id:%x",
		m.Token, m.Account, new(big.Int).SetBytes(m.Amount), m.Depositor, m.RefChainId, m.RefId)
}

func (w *WithdrawOnChain) String() string {
	if w == nil {
		return "nil"
	}
	return fmt.Sprintf("token:%x receiver:%x amount%s burn_account:%x ref_chain_id:%x ref_id:%x",
		w.Token, w.Receiver, new(big.Int).SetBytes(w.Amount), w.BurnAccount, w.RefChainId, w.RefId)
}

func (c *PegConfig) Validate() error {
	// TODO
	return nil
}
