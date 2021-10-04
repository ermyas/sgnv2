package types

import (
	"bytes"
	"fmt"
	"math/big"
	"sort"

	"github.com/celer-network/sgn-v2/eth"
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
	for _, s := range ls.GetSortedSigs() {
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
		for i, s := range ls.GetSortedSigs() {
			sigs[i] = s.Sig
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

func ValidateSigs(sortedSigs []*AddrSig, curss *SortedSigners) (pass bool, sigsBytes [][]byte) {
	if len(curss.GetSigners()) == 0 {
		return false, nil
	}
	totalPower := big.NewInt(0)
	curssMap := make(map[eth.Addr]*AddrAmt)
	for _, s := range curss.GetSigners() {
		power := big.NewInt(0).SetBytes(s.Amt)
		totalPower.Add(totalPower, power)
		curssMap[eth.Bytes2Addr(s.Addr)] = s
	}
	quorumStake := big.NewInt(0).Mul(totalPower, big.NewInt(2))
	quorumStake = quorumStake.Quo(quorumStake, big.NewInt(3))

	signedPower := big.NewInt(0)
	for _, s := range sortedSigs {
		if addrAmt, ok := curssMap[eth.Bytes2Addr(s.Addr)]; ok {
			power := big.NewInt(0).SetBytes(addrAmt.Amt)
			signedPower.Add(signedPower, power)
			sigsBytes = append(sigsBytes, s.Sig)
			if signedPower.Cmp(quorumStake) > 0 {
				return true, sigsBytes
			}
			delete(curssMap, eth.Bytes2Addr(s.Addr))
		}
	}

	return false, nil
}
