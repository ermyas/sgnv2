package types

import (
	"bytes"
	"fmt"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	ethproto "github.com/celer-network/sgn-v2/proto/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

func NewAcctAmtPair(account string, amount sdk.Int) AcctAmtPair {
	return AcctAmtPair{
		Account: eth.FormatAddrHex(account),
		Amount:  amount,
	}
}

func NewSig(signer string, sig []byte) Sig {
	return Sig{
		Signer: signer,
		Sig:    sig,
	}
}

// TODO: Use commontypes.Signature
func AddSig(sigs []Sig, msg []byte, sig []byte, expectedSigner string) ([]Sig, error) {
	// make sure sig won't be changed by callee
	tmpSig := make([]byte, len(sig))
	copy(tmpSig, sig)
	signer, err := ethutils.RecoverSigner(msg, tmpSig)
	if err != nil {
		return nil, err
	}

	signerAddr := eth.Addr2Hex(signer)
	if signerAddr != eth.FormatAddrHex(expectedSigner) {
		err = fmt.Errorf("invalid signer address %s %s", signerAddr, expectedSigner)
		return nil, err
	}

	for i, s := range sigs {
		if s.Signer == signerAddr {
			if bytes.Equal(s.Sig, sig) {
				// already signed with the same sig
				return sigs, nil
			}
			log.Debugf("repeated signer %s overwite existing sig", signerAddr)
			sigs[i] = NewSig(signerAddr, sig)
			return sigs, nil
		}
	}

	return append(sigs, NewSig(signerAddr, sig)), nil
}

func NewSlash(nonce, slashFactor, jailPeriod, expireTime uint64, reason, valEthAddr string, collectors []AcctAmtPair) Slash {
	return Slash{
		Validator:   eth.FormatAddrHex(valEthAddr),
		Nonce:       nonce,
		SlashFactor: slashFactor,
		ExpireTime:  expireTime,
		JailPeriod:  jailPeriod,
		Collectors:  collectors,
		Reason:      reason,
	}
}

func (s *Slash) GenerateEthSlashBytes() {
	var collectors []*ethproto.AcctAmtPair

	for _, collector := range s.Collectors {
		collectors = append(collectors, &ethproto.AcctAmtPair{
			Account: eth.Hex2Addr(collector.Account).Bytes(),
			Amount:  collector.Amount.BigInt().Bytes(),
		})
	}

	slashBytes, _ := proto.Marshal(&ethproto.Slash{
		Validator:   eth.Hex2Addr(s.Validator).Bytes(),
		Nonce:       s.Nonce,
		SlashFactor: s.SlashFactor,
		ExpireTime:  s.ExpireTime,
		JailPeriod:  s.JailPeriod,
		Collectors:  collectors,
	})

	s.EthSlashBytes = slashBytes
}

// Add signature to slash sigs
func (s *Slash) AddSig(sig []byte, expectedSigner string) error {
	sigs, err := AddSig(s.Sigs, s.EthSlashBytes, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}

	s.Sigs = sigs
	return nil
}

func (s *Slash) GetSigsBytes() [][]byte {
	if s != nil {
		sigs := make([][]byte, 0)
		for i := range s.Sigs {
			sigs = append(sigs, s.Sigs[i].Sig)
		}
		return sigs
	}
	return nil
}
