package types

import (
	"bytes"
	"fmt"
	"math/big"
	"sort"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/codec"
)

// ------------------------------------ Signer(s) ------------------------------------

func (s *Signer) String() string {
	return fmt.Sprintf("addr %x power %s", s.Addr, new(big.Int).SetBytes(s.Power))
}

func PrintSigners(ss []*Signer) string {
	var out string
	for _, s := range ss {
		out += fmt.Sprintf("<%s> ", s.String())
	}
	return fmt.Sprintf("< %s>", out)
}

func SignersToEthArrays(ss []*Signer) ([]eth.Addr, []*big.Int) {
	addrs := make([]eth.Addr, len(ss))
	powers := make([]*big.Int, len(ss))
	for i, s := range ss {
		addrs[i] = eth.Bytes2Addr(s.Addr)
		powers[i] = new(big.Int).SetBytes(s.Power)
	}
	return addrs, powers
}

// ------------------------------------ ChainSigners ------------------------------------

func (cs *ChainSigners) String() string {
	return fmt.Sprintf("chainId: %d, signers: %s", cs.ChainId, PrintSigners(cs.GetSortedSigners()))
}

func (ss *ChainSigners) SetByEvent(e *eth.BridgeSignersUpdated) {
	ss.SortedSigners = make([]*Signer, len(e.Powers))
	for i, addr := range e.Signers {
		ss.SortedSigners[i] = &Signer{addr.Bytes(), e.Powers[i].Bytes()}
	}
}

func MustMarshalChainSigners(cdc codec.BinaryCodec, signers *ChainSigners) []byte {
	return cdc.MustMarshal(signers)
}

func MustUnmarshalChainSigners(cdc codec.BinaryCodec, value []byte) ChainSigners {
	signers, err := UnmarshalChainSigners(cdc, value)
	if err != nil {
		panic(err)
	}
	return signers
}

func UnmarshalChainSigners(cdc codec.BinaryCodec, value []byte) (s ChainSigners, err error) {
	err = cdc.Unmarshal(value, &s)
	return s, err
}

// ------------------------------------ LatestSigners ------------------------------------

func (ls *LatestSigners) String() string {
	var sigs string
	for _, s := range ls.GetSortedSigs() {
		sigs += fmt.Sprintf("%x ", s.Addr)
	}
	return fmt.Sprintf("signers: %s, sigs from: < %s>, update time: %s",
		PrintSigners(ls.GetSortedSigners()), sigs, ls.UpdateTime)
}

func (ls *LatestSigners) GenerateSignersBytes() {
	if ls == nil {
		return
	}
	addrs, powers := SignersToEthArrays(ls.SortedSigners)
	ls.SignersBytes = eth.SignerBytes(addrs, powers)
}

func (ls *LatestSigners) GetSortedSigsBytes() [][]byte {
	if ls != nil {
		sigs := make([][]byte, len(ls.SortedSigs))
		for i, s := range ls.GetSortedSigs() {
			sigs[i] = s.Sig
		}
		return sigs
	}
	return nil
}

// Sort signers array in ascending address order
func (ls *LatestSigners) Sort() {
	sort.Sort(ls)
}

// Implements sort interface
func (ls *LatestSigners) Len() int {
	return len(ls.SortedSigners)
}

// Implements sort interface
func (ls *LatestSigners) Less(i, j int) bool {
	return bytes.Compare(eth.Pad20Bytes(ls.SortedSigners[i].Addr), eth.Pad20Bytes(ls.SortedSigners[j].Addr)) == -1
}

// Implements sort interface
func (ls *LatestSigners) Swap(i, j int) {
	ls.SortedSigners[i], ls.SortedSigners[j] = ls.SortedSigners[j], ls.SortedSigners[i]
}

func MustMarshalLatestSigners(cdc codec.BinaryCodec, signers *LatestSigners) []byte {
	return cdc.MustMarshal(signers)
}

func MustUnmarshalLatestSigners(cdc codec.BinaryCodec, value []byte) LatestSigners {
	signers, err := UnmarshalLatestSigners(cdc, value)
	if err != nil {
		panic(err)
	}
	return signers
}

func UnmarshalLatestSigners(cdc codec.BinaryCodec, value []byte) (s LatestSigners, err error) {
	err = cdc.Unmarshal(value, &s)
	return s, err
}

// ------------------------------------ Utils ------------------------------------

func ValidateSigQuorum(sortedSigs []*AddrSig, curss []*Signer) bool {
	if len(curss) == 0 {
		return false
	}
	totalPower := big.NewInt(0)
	curssMap := make(map[eth.Addr]*Signer)
	for _, s := range curss {
		power := big.NewInt(0).SetBytes(s.Power)
		totalPower.Add(totalPower, power)
		curssMap[eth.Bytes2Addr(s.Addr)] = s
	}
	quorumStake := big.NewInt(0).Mul(totalPower, big.NewInt(2))
	quorumStake = quorumStake.Quo(quorumStake, big.NewInt(3))

	signedPower := big.NewInt(0)
	for _, s := range sortedSigs {
		if signer, ok := curssMap[eth.Bytes2Addr(s.Addr)]; ok {
			power := big.NewInt(0).SetBytes(signer.Power)
			signedPower.Add(signedPower, power)
			if signedPower.Cmp(quorumStake) > 0 {
				return true
			}
			delete(curssMap, eth.Bytes2Addr(s.Addr))
		}
	}

	return false
}
