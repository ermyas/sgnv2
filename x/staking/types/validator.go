package types

import (
	"bytes"
	fmt "fmt"
	"sort"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"
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

func (v Validator) YamlStr() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

func (v Validator) ConsPubKey() (cryptotypes.PubKey, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdkerrors.Wrapf(ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
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
		return nil, sdkerrors.Wrapf(ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	return sdk.ConsAddress(pk.Address()), nil
}

// IsBonded checks if the validator status equals Bonded
func (v Validator) IsBonded() bool {
	return v.GetStatus() == ValidatorStatus_Bonded
}

// IsUnbonded checks if the validator status equals Unbonded
func (v Validator) IsUnbonded() bool {
	return v.GetStatus() == ValidatorStatus_Unbonded
}

// IsUnbonding checks if the validator status equals Unbonding
func (v Validator) IsUnbonding() bool {
	return v.GetStatus() == ValidatorStatus_Unbonding
}

// ConsensusPower gets the consensus-engine power. Aa reduction of 10^6 from
// validator tokens is applied
func (v Validator) ConsensusPower(r sdk.Int) int64 {
	if v.IsBonded() {
		return v.PotentialConsensusPower(r)
	}
	return 0
}

// PotentialConsensusPower returns the potential consensus-engine power.
func (v Validator) PotentialConsensusPower(r sdk.Int) int64 {
	return sdk.TokensToConsensusPower(v.Tokens, r)
}

// ABCIValidatorUpdate returns an abci.ValidatorUpdate from a staking validator type
// with the full validator power
func (v Validator) ABCIValidatorUpdate(r sdk.Int) abci.ValidatorUpdate {
	tmProtoPk, err := v.TmConsPublicKey()
	if err != nil {
		panic(err)
	}

	return abci.ValidatorUpdate{
		PubKey: tmProtoPk,
		Power:  v.ConsensusPower(r),
	}
}

// ABCIValidatorUpdateZero returns an abci.ValidatorUpdate from a staking validator type
// with zero power used for validator updates.
func (v Validator) ABCIValidatorUpdateZero() abci.ValidatorUpdate {
	tmProtoPk, err := v.TmConsPublicKey()
	if err != nil {
		panic(err)
	}

	return abci.ValidatorUpdate{
		PubKey: tmProtoPk,
		Power:  0,
	}
}

func (v Validator) GetEthAddr() eth.Addr {
	return eth.Hex2Addr(v.EthAddress)
}

func (v Validator) GetSgnAddr() sdk.AccAddress {
	if v.SgnAddress == "" {
		return nil
	}
	addr, err := sdk.AccAddressFromBech32(v.SgnAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

func (v Validator) GetMoniker() string { return v.GetDescription().GetMoniker() }

// Validators is a collection of Validator
type Validators []Validator

// Sort Validators sorts validator array in ascending operator address order
func (v Validators) Sort() {
	sort.Sort(v)
}

// Implements sort interface
func (v Validators) Len() int {
	return len(v)
}

// Implements sort interface
func (v Validators) Less(i, j int) bool {
	return v[i].GetEthAddress() < v[j].GetEthAddress()
}

// Implements sort interface
func (v Validators) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

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

// ValidatorsByVotingPower implements sort.Interface for []Validator based on
// the VotingPower and Address fields.
// The validators are sorted first by their voting power (descending). Secondary index - Address (ascending).
// Copied from tendermint/types/validator_set.go
type ValidatorsByVotingPower []Validator

func (valz ValidatorsByVotingPower) Len() int { return len(valz) }

func (valz ValidatorsByVotingPower) Less(i, j int, r sdk.Int) bool {
	if valz[i].ConsensusPower(r) == valz[j].ConsensusPower(r) {
		addrI, errI := valz[i].GetConsAddr()
		addrJ, errJ := valz[j].GetConsAddr()
		// If either returns error, then return false
		if errI != nil || errJ != nil {
			return false
		}
		return bytes.Compare(addrI, addrJ) == -1
	}
	return valz[i].ConsensusPower(r) > valz[j].ConsensusPower(r)
}

func (valz ValidatorsByVotingPower) Swap(i, j int) {
	valz[i], valz[j] = valz[j], valz[i]
}
