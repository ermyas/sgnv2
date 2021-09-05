package types

import (
	fmt "fmt"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	proto "github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"
)

func NewValidator(ethAddress, ethSigner, sgnAddress string) *Validator {
	return &Validator{
		EthAddress: eth.FormatAddrHex(ethAddress),
		EthSigner:  eth.FormatAddrHex(ethSigner),
		SgnAddress: sgnAddress,
		Status:     ValidatorStatus_Unbonded,
	}
}

func MustMarshalValidator(cdc codec.BinaryCodec, validator *Validator) []byte {
	return cdc.MustMarshal(validator)
}

func MustUnmarshalValidator(cdc codec.BinaryCodec, value []byte) Validator {
	validator, err := UnmarshalValidator(cdc, value)
	if err != nil {
		panic(err)
	}

	return validator
}

func UnmarshalValidator(cdc codec.BinaryCodec, value []byte) (v Validator, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (v *Validator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pk cryptotypes.PubKey
	return unpacker.UnpackAny(v.ConsensusPubkey, &pk)
}

func (v Validator) String() string {
	pubkey := v.GetConsensusPubkey()
	v.ConsensusPubkey = nil
	out := proto.CompactTextString(&v)
	if pubkey != nil {
		out += fmt.Sprintf("consensus_pubkey: %x", pubkey.Value)
	}
	return out
}

func (v Validator) YamlStr() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

// Validators is a collection of Validator
type Validators []*Validator

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (v Validators) UnpackInterfaces(c codectypes.AnyUnpacker) error {
	for i := range v {
		if err := v[i].UnpackInterfaces(c); err != nil {
			return err
		}
	}
	return nil
}

func (v Validators) String() (out string) {
	for _, val := range v {
		out += val.String() + " | "
	}
	return out
}
