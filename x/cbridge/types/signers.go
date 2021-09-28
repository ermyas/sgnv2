package types

import (
	"bytes"
	"fmt"
	"math/big"
	"sort"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gogo/protobuf/proto"
)

func (ss *SortedSigners) String() string {
	var out string
	for _, s := range ss.Signers {
		power := new(big.Int).SetBytes(s.Amt)
		out += fmt.Sprintf("<%x, %s> ", s.Addr, power)
	}
	return fmt.Sprintf("< %s>", out)
}

// Sort signers array in descending token amount order
func (ss SortedSigners) Sort() {
	sort.Sort(ss)
}

// Implements sort interface
func (ss SortedSigners) Len() int {
	return len(ss.Signers)
}

// Implements sort interface
func (ss SortedSigners) Less(i, j int) bool {
	return bytes.Compare(ss.Signers[i].Addr, ss.Signers[j].Addr) == -1
}

// Implements sort interface
func (ss SortedSigners) Swap(i, j int) {
	ss.Signers[i], ss.Signers[j] = ss.Signers[j], ss.Signers[i]
}

func (cs *ChainSigners) String() string {
	return fmt.Sprintf("chainId: %d, signers: %s", cs.ChainId, cs.CurrSigners.String())
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

func (ls *LatestSigners) String() string {
	var sigs string
	for _, s := range ls.SortedSigs {
		sigs += fmt.Sprintf("%x ", s.Addr)
	}
	return fmt.Sprintf("signers: %s, sigs from: < %s>, update time: %s", ls.Signers.String(), sigs, ls.UpdateTime)
}

func (ls *LatestSigners) GenerateSignersBytes() {
	ls.SignersBytes, _ = proto.Marshal(ls.Signers)
}

func (ls *LatestSigners) GetSortedSigsBytes() [][]byte {
	if ls != nil {
		sigs := make([][]byte, len(ls.SortedSigs))
		for i := range ls.SortedSigs {
			sigs[i] = ls.SortedSigs[i].Sig
		}
		return sigs
	}
	return nil
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
