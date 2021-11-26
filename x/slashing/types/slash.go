package types

import (
	fmt "fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/gogo/protobuf/proto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func NewSlash(
	nonce uint64, valEthAddr eth.Addr, slashFactor, jailPeriod,
	expireTime uint64, reason string, collectors []*AcctAmtPair) Slash {

	ethSlashMsg := SlashOnChain{
		Validator:   valEthAddr.Bytes(),
		Nonce:       nonce,
		SlashFactor: slashFactor,
		ExpireTime:  expireTime,
		JailPeriod:  jailPeriod,
		Collectors:  collectors,
	}

	return Slash{
		SlashOnChain: ethSlashMsg,
		Reason:       reason,
	}
}

func (s *Slash) GenerateSlashBytes() {
	slashBytes, err := proto.Marshal(&s.SlashOnChain)
	if err != nil {
		log.Error("generate slash bytes err", err)
		return
	}
	s.SlashBytes = slashBytes
}

// Add signature to slash sigs
func (s *Slash) AddSig(msgToSign []byte, sig []byte, expectedSigner string) error {
	sigs, err := commontypes.AddSig(s.Signatures, msgToSign, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}

	s.Signatures = sigs
	return nil
}

func (s *Slash) EncodeDataToSign(chainId uint64, contractAddr eth.Addr) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(chainId), contractAddr, "Slash"},
	)
	return append(domain, s.SlashBytes...)
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

func (s *Slash) String() string {
	if s == nil {
		return "nil"
	}
	res := fmt.Sprintf("%s reason:%s bytes:0x%x", &s.SlashOnChain, s.Reason, s.SlashBytes)
	if len(s.Signatures) > 0 {
		res += " signers:<"
		for _, s := range s.Signatures {
			res += fmt.Sprintf(" %s", s.Signer)
		}
		res += " >"
	}
	return res
}

func (m *SlashOnChain) String() string {
	if m == nil {
		return "nil"
	}
	res := fmt.Sprintf("validator:%x nonce:%d slash_factor:%d expire_time:%d jail_period:%d",
		m.Validator, m.Nonce, m.SlashFactor, m.ExpireTime, m.JailPeriod)
	if len(m.Collectors) > 0 {
		res += " collectors:<"
		for _, c := range m.Collectors {
			res += fmt.Sprintf(" %s", c)
		}
		res += " >"
	}
	return res
}

func (p *AcctAmtPair) String() string {
	if p == nil {
		return "nil"
	}
	return fmt.Sprintf("account:%x amount:%x", p.Account, new(big.Int).SetBytes(p.Amount))
}
