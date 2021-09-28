package types

import (
	bytes "bytes"
	"fmt"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
)

func NewSignature(signer string, sigBytes []byte) Signature {
	return Signature{
		Signer:   signer,
		SigBytes: sigBytes,
	}
}

// TODO: Remove duplicate functionality
func AddSig(sigs []Signature, msg []byte, sigBytes []byte, expectedSigner string) ([]Signature, error) {
	// make sure sig won't be changed by callee
	tmpSig := make([]byte, len(sigBytes))
	copy(tmpSig, sigBytes)
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
			if bytes.Equal(s.SigBytes, sigBytes) {
				// already signed with the same sig
				return sigs, nil
			}
			log.Debugf("repeated signer %s overwite existing sig", signerAddr)
			sigs[i] = NewSignature(signerAddr, sigBytes)
			return sigs, nil
		}
	}

	return append(sigs, NewSignature(signerAddr, sigBytes)), nil
}

// AddSig adds a signature to a reward claim details
func (d *RewardClaimDetails) AddSig(sig []byte, expectedSigner string) error {
	sigs, err := AddSig(d.Signatures, d.RewardProtoBytes, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}
	d.Signatures = sigs
	return nil
}
