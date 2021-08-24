package common

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
)

type ParamChange struct {
	Record   sdk.Int `json:"record"`
	NewValue sdk.Int `json:"new_value"`
}

func NewParamChange(record, newValue sdk.Int) ParamChange {
	return ParamChange{
		Record:   record,
		NewValue: newValue,
	}
}

// implement fmt.Stringer
func (p ParamChange) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Record: %v, NewValue: %v`, p.Record, p.NewValue))
}

type Sig struct {
	Signer string `json:"signer"`
	Sig    []byte `json:"sig"`
}

func NewSig(signer string, sig []byte) Sig {
	return Sig{
		Signer: signer,
		Sig:    sig,
	}
}

// implement fmt.Stringer
func (r Sig) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Signer: %s, Sig: %x,`, r.Signer, r.Sig))
}

func AddSig(sigs []Sig, msg []byte, sig []byte, expectedSigner string) ([]Sig, error) {
	signer, err := ethutils.RecoverSigner(msg, sig)
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
			if bytes.Compare(s.Sig, sig) == 0 {
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

func NewEthClientFromConfig() (*eth.EthClient, error) {
	return eth.NewEthClient(
		viper.GetString(FlagEthGateway),
		viper.GetString(FlagEthKeystore),
		viper.GetString(FlagEthPassphrase),
		&eth.TransactorConfig{
			BlockDelay:           viper.GetUint64(FlagEthBlockDelay),
			BlockPollingInterval: viper.GetUint64(FlagEthPollInterval),
			ChainId:              big.NewInt(viper.GetInt64(FlagEthChainID)),
			AddGasPriceGwei:      viper.GetUint64(FlagEthAddGasPriceGwei),
			MinGasPriceGwei:      viper.GetUint64(FlagEthMinGasPriceGwei),
		},
		viper.GetString(FlagEthStakingAddress),
		viper.GetString(FlagEthSGNAddress),
	)
}
