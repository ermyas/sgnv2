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

func ValidateSigQuorum(sortedSigs []*AddrSig, curss []*Signer) (pass bool, sigsBytes [][]byte) {
	if len(curss) == 0 {
		return false, nil
	}
	totalPower := big.NewInt(0)
	signerPowers := make(map[eth.Addr]*big.Int)
	for _, s := range curss {
		power := big.NewInt(0).SetBytes(s.Power)
		totalPower.Add(totalPower, power)
		signerPowers[eth.Bytes2Addr(s.GetAddr())] = power
	}
	quorumStake := big.NewInt(0).Mul(totalPower, big.NewInt(2))
	quorumStake = quorumStake.Quo(quorumStake, big.NewInt(3))

	var filteredSigs []*AddrSig
	for _, sig := range sortedSigs {
		if _, ok := signerPowers[eth.Bytes2Addr(sig.GetAddr())]; ok {
			filteredSigs = append(filteredSigs, sig)
		}
	}
	// sort signer by power
	sort.Slice(filteredSigs, func(i, j int) bool {
		return signerPowers[eth.Bytes2Addr(filteredSigs[i].GetAddr())].Cmp(signerPowers[eth.Bytes2Addr(filteredSigs[j].GetAddr())]) > 0
	})

	signedPower := big.NewInt(0)
	var quorumSigners []*AddrSig
	for _, s := range filteredSigs {
		if power, ok := signerPowers[eth.Bytes2Addr(s.Addr)]; ok {
			signedPower.Add(signedPower, power)
			quorumSigners = append(quorumSigners, s)
			if signedPower.Cmp(quorumStake) > 0 {
				sort.Slice(quorumSigners, func(i, j int) bool {
					return bytes.Compare(eth.Pad20Bytes(quorumSigners[i].Addr), eth.Pad20Bytes(quorumSigners[j].Addr)) == -1
				})
				for _, signer := range quorumSigners {
					sigsBytes = append(sigsBytes, signer.Sig)
				}
				return true, sigsBytes
			}
			delete(signerPowers, eth.Bytes2Addr(s.Addr))
		}
	}

	return false, nil
}

func MinSigsForQuorum(signers []*Signer) uint32 {
	if len(signers) == 0 {
		return 0
	}

	sigMap := make(map[eth.Addr]*big.Int)
	totalPower := big.NewInt(0)
	for _, signer := range signers {
		power := big.NewInt(0).SetBytes(signer.GetPower())
		sigMap[eth.Bytes2Addr(signer.Addr)] = power
		totalPower.Add(totalPower, power)
	}
	quorumStake := big.NewInt(0).Mul(totalPower, big.NewInt(2))
	quorumStake = quorumStake.Quo(quorumStake, big.NewInt(3))
	sort.Slice(signers, func(i, j int) bool {
		return sigMap[eth.Bytes2Addr(signers[i].Addr)].Cmp(sigMap[eth.Bytes2Addr(signers[j].Addr)]) > 0
	})

	signedPower := big.NewInt(0)
	var count uint32 = 0
	for _, signer := range signers {
		signedPower.Add(signedPower, sigMap[eth.Bytes2Addr(signer.Addr)])
		count++
		if signedPower.Cmp(quorumStake) > 0 {
			return count
		}
	}
	return count
}
