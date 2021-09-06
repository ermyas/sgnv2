package types

import (
	fmt "fmt"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
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
func (v Validator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pk cryptotypes.PubKey
	return unpacker.UnpackAny(v.ConsensusPubkey, &pk)
}

func (v Validator) String() string {
	var pubkey string
	if v.GetConsensusPubkey() != nil {
		if v.GetConsensusPubkey().GetCachedValue() != nil {
			consAddr, err := v.GetConsAddr()
			if err != nil {
				pubkey = fmt.Sprintf("consensus_address:%s", err)
			} else {
				pubkey = fmt.Sprintf("consensus_address:\"%s\"", consAddr.String())
			}
		} else {
			pubkey = fmt.Sprintf("consensus_pubkey:\"%x\"", v.GetConsensusPubkey().Value)
		}
	}
	v.ConsensusPubkey = nil
	out := proto.CompactTextString(&v)
	out += pubkey
	return out
}

func (v Validator) ConsPubKey() (cryptotypes.PubKey, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdk_errors.Wrapf(ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return pk, nil
}

func (v Validator) TmConsPublicKey() (tmprotocrypto.PublicKey, error) {
	pk, err := v.ConsPubKey()
	if err != nil {
		return tmprotocrypto.PublicKey{}, err
	}

	tmPk, err := cryptocodec.ToTmProtoPublicKey(pk)
	if err != nil {
		return tmprotocrypto.PublicKey{}, err
	}

	return tmPk, nil
}

func (v Validator) GetConsAddr() (sdk.ConsAddress, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdk_errors.Wrapf(ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return sdk.ConsAddress(pk.Address()), nil
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
