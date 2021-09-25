package types

import (
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

var _ ValidatorI = Validator{}

func NewValidator(ethAddress, ethSigner, sgnAddress string) *Validator {
	return &Validator{
		EthAddress: eth.FormatAddrHex(ethAddress),
		EthSigner:  eth.FormatAddrHex(ethSigner),
		SgnAddress: sgnAddress,
		Status:     Unbonded,
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
	if v.ConsensusPubkey != nil {
		if v.ConsensusPubkey.GetCachedValue() != nil {
			consAddr, err := v.GetConsAddr()
			if err != nil {
				pubkey = fmt.Sprintf("consensus_address:%s", err)
			} else {
				pubkey = fmt.Sprintf("consensus_address:\"%s\"", consAddr.String())
			}
		} else {
			pubkey = fmt.Sprintf("consensus_pubkey:\"%x\"", v.ConsensusPubkey.Value)
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
	return v.GetStatus() == Bonded
}

// IsUnbonded checks if the validator status equals Unbonded
func (v Validator) IsUnbonded() bool {
	return v.GetStatus() == Unbonded
}

// IsUnbonding checks if the validator status equals Unbonding
func (v Validator) IsUnbonding() bool {
	return v.GetStatus() == Unbonding
}

// get the bonded tokens which the validator holds
func (v Validator) BondedTokens() sdk.Int {
	if v.IsBonded() {
		return v.Tokens
	}
	return sdk.ZeroInt()
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

func (v Validator) GetSignerAddr() eth.Addr {
	return eth.Hex2Addr(v.EthSigner)
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

func (v Validator) GetOperator() sdk.ValAddress {
	if v.SgnAddress == "" {
		return nil
	}
	addr, err := sdk.ValAddressFromBech32(v.SgnAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

func (v Validator) GetMoniker() string { return v.Description.GetMoniker() }

// Validators is a collection of Validator
type Validators []Validator

// Sort validator array in descending token amount order
func (v Validators) Sort() {
	sort.Sort(v)
}

// Implements sort interface
func (v Validators) Len() int {
	return len(v)
}

// Implements sort interface
func (v Validators) Less(i, j int) bool {
	return v[i].Tokens.GT(v[j].Tokens)
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

func (v Validator) GetCommission() sdk.Dec { return v.CommissionRate }

func (v Validator) GetEthAddress() eth.Addr { return eth.Hex2Addr(v.EthAddress) }

func (v Validator) GetStatus() BondStatus { return v.Status }

func (v Validator) GetTokens() sdk.Int { return v.Tokens }

func (v Validator) GetBondedTokens() sdk.Int { return v.BondedTokens() }

func (v Validator) GetDelegatorShares() sdk.Int { return v.DelegatorShares }

// calculate the token worth of provided shares
func (v Validator) TokensFromShares(shares sdk.Int) sdk.Dec {
	return (shares.Mul(v.Tokens)).ToDec().Quo(v.DelegatorShares.ToDec())
}

// calculate the token worth of provided shares, truncated
func (v Validator) TokensFromSharesTruncated(shares sdk.Int) sdk.Dec {
	return (shares.Mul(v.Tokens)).ToDec().QuoTruncate(v.DelegatorShares.ToDec())
}

// TokensFromSharesRoundUp returns the token worth of provided shares, rounded
// up.
func (v Validator) TokensFromSharesRoundUp(shares sdk.Int) sdk.Dec {
	return (shares.Mul(v.Tokens)).ToDec().QuoRoundUp(v.DelegatorShares.ToDec())
}

func (v Validator) GetConsensusPower(r sdk.Int) int64 {
	return v.ConsensusPower(r)
}

func (v Validators) String() (out string) {
	for _, val := range v {
		out += val.String() + " | "
	}
	return out
}

// ----------------------- CLI print-friendly output --------------------

type ValidatorOutput struct {
	EthAddress       string       `json:"eth_address" yaml:"eth_address"`
	EthSigner        string       `json:"eth_signer" yaml:"signer_address"`
	SgnAddress       string       `json:"sgn_address" yaml:"sgn_address"`
	ConsensusAddress string       `json:"consensus_address" yaml:"consensus_address"`
	Status           string       `json:"status" yaml:"status"`
	Tokens           sdk.Dec      `json:"tokens" yaml:"tokens"`
	Shares           sdk.Dec      `json:"shares" yaml:"shares"`
	CommissionRate   sdk.Dec      `json:"commission_rate" yaml:"commission_rate"`
	Description      *Description `json:"description" yaml:"description"`
}

func newValidatorOutput(v *Validator) *ValidatorOutput {
	output := &ValidatorOutput{
		EthAddress:     v.EthAddress,
		EthSigner:      v.EthSigner,
		SgnAddress:     v.SgnAddress,
		Status:         v.Status.String(),
		Tokens:         sdk.NewDecFromIntWithPrec(v.Tokens, 18),
		Shares:         sdk.NewDecFromIntWithPrec(v.DelegatorShares, 18),
		CommissionRate: v.CommissionRate,
		Description:    v.Description,
	}
	if v.ConsensusPubkey != nil && v.ConsensusPubkey.GetCachedValue() != nil {
		consAddr, err := v.GetConsAddr()
		if err != nil {
			output.ConsensusAddress = err.Error()
		}
		output.ConsensusAddress = consAddr.String()
	}
	return output
}

func (v *Validator) YamlStr() string {
	output := newValidatorOutput(v)
	out, _ := yaml.Marshal(output)
	return string(out)
}
