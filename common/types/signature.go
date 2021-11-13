package types

import (
	bytes "bytes"
	fmt "fmt"

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

func AddSig(sigs []Signature, msg []byte, sigBytes []byte, expectedSigner string) ([]Signature, error) {
	signer, err := ethutils.RecoverSigner(msg, sigBytes)
	if err != nil {
		return nil, err
	}

	signerAddr := eth.Addr2Hex(signer)
	if signerAddr != eth.FormatAddrHex(expectedSigner) {
		err = fmt.Errorf("invalid signer address %s %s", signerAddr, expectedSigner)
		return nil, err
	}

	newSig := NewSignature(signerAddr, sigBytes)
	// Keep sigs sorted in ascending order by signer address and check for duplicates
	for i, sig := range sigs {
		if sig.Signer == signerAddr {
			// Overwriting existing sig
			if bytes.Equal(sig.SigBytes, sigBytes) {
				// no-op, already signed with the same sig
				return sigs, nil
			}
			log.Debugf("repeated signer %s overwite existing sig", signerAddr)
			sigs[i] = newSig
			return sigs, nil
		}
		if signerAddr < sig.Signer {
			// Found the spot, do insertion
			newSigs := append(sigs[:i+1], sigs[i:]...)
			newSigs[i] = newSig
			return newSigs, nil
		}
	}

	// Address larger than all existing signers, append to the end
	return append(sigs, newSig), nil
}
