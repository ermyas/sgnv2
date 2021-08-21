package types

import (
	"fmt"

	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// validator params default values
const (
	DefaultSyncerDuration   uint64 = 10
	DefaultEpochLength      uint64 = 5
	DefaultMaxValidatorDiff uint64 = 10
)

// nolint - Keys for parameter access
var (
	KeySyncerDuration   = []byte("SyncerDuration")
	KeyEpochLength      = []byte("EpochLength")
	KeyMaxValidatorDiff = []byte("KeyMaxValidatorDiff")
)

var _ sdk_params.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(
	syncerDuration, epochLength, maxValidatorDiff uint64) Params {

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
