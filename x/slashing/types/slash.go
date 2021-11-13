package types

import (
	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
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
	sigs, err := commontypes.AddSig(s.Signatures, s.EthSlashBytes, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}

	s.Signatures = sigs
	return nil
}

func (s *Slash) GetSigsBytes() [][]byte {
	if s != nil {
		sigs := make([][]byte, 0)
		for i := range s.Signatures {
			sigs = append(sigs, s.Signatures[i].SigBytes)
		}
		return sigs
	}
	return nil
}
