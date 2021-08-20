package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// validator params default values
const (
	DefaultSyncerDuration   uint = 10
	DefaultEpochLength      uint = 5
	DefaultMaxValidatorDiff uint = 10
)

// nolint - Keys for parameter access
var (
	KeySyncerDuration   = []byte("SyncerDuration")
	KeyEpochLength      = []byte("EpochLength")
	KeyMaxValidatorDiff = []byte("KeyMaxValidatorDiff")
)

var _ sdk_params.ParamSet = (*Params)(nil)

type Params struct {
	SyncerDuration   uint `json:"syncer_duration" yaml:"syncer_duration"`
	EpochLength      uint `json:"epoch_length" yaml:"epoch_length"`
	MaxValidatorDiff uint `json:"max_validator_diff" yaml:"max_validator_diff"`
}

// NewParams creates a new Params instance
func NewParams(
	syncerDuration, epochLength, maxValidatorDiff uint) Params {

	return Params{
		SyncerDuration: syncerDuration,
		EpochLength:    epochLength,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() sdk_params.ParamSetPairs {
	return sdk_params.ParamSetPairs{
		sdk_params.NewParamSetPair(KeySyncerDuration, &p.SyncerDuration, validateSyncerDuration),
		sdk_params.NewParamSetPair(KeyEpochLength, &p.EpochLength, validateEpochLength),
		sdk_params.NewParamSetPair(KeyMaxValidatorDiff, &p.MaxValidatorDiff, validateMaxValidatorDiff),
	}
}

// Equal returns a boolean determining if two Param types are identical.
func (p Params) Equal(p2 Params) bool {
	// bz1 := ModuleCdc.MustMarshalLengthPrefixed(&p)
	// bz2 := ModuleCdc.MustMarshalLengthPrefixed(&p2)
	// return bytes.Equal(bz1, bz2)
	return false //TODO
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultSyncerDuration, DefaultEpochLength, DefaultMaxValidatorDiff)
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	return fmt.Sprintf(`Params:
  SyncerDuration:   %d,
  EpochLength:      %d,
  MaxValidatorDiff: %d,`,
		p.SyncerDuration, p.EpochLength, p.MaxValidatorDiff)
}

// unmarshal the current validator params value from store key or panic
func MustUnmarshalParams(cdc *codec.Codec, value []byte) Params {
	params, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}
	return params
}

// unmarshal the current validator params value from store key
func UnmarshalParams(cdc *codec.Codec, value []byte) (params Params, err error) {
	// err = cdc.UnmarshalLengthPrefixed(value, &params)
	// if err != nil {
	// 	return
	// }
	return
}

// validate a set of params
func (p Params) Validate() error {
	if p.SyncerDuration == 0 {
		return fmt.Errorf("validator parameter SyncerDuration must be a positive integer")
	}

	if p.EpochLength == 0 {
		return fmt.Errorf("validator parameter EpochLength must be a positive integer")
	}

	return nil
}

func validateSyncerDuration(i interface{}) error {
	v, ok := i.(uint)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("validator parameter SyncerDuration must be positive: %d", v)
	}

	return nil
}

func validateEpochLength(i interface{}) error {
	v, ok := i.(uint)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("validator parameter EpochLength must be positive: %d", v)
	}

	return nil
}

func validateMaxValidatorDiff(i interface{}) error {
	_, ok := i.(uint)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
