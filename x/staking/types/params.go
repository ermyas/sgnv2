package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// validator params default values
const (
	DefaultSyncerDuration uint64 = 10
)

var DefaultPowerReduction = sdk.NewIntFromUint64(1000000000000)

// nolint - Keys for parameter access
var (
	KeySyncerDuration = []byte("SyncerDuration")
)

var _ sdk_params.ParamSet = (*Params)(nil)

func ParamKeyTable() sdk_params.KeyTable {
	return sdk_params.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(syncerDuration uint64) Params {

	return Params{
		SyncerDuration: syncerDuration,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() sdk_params.ParamSetPairs {
	return sdk_params.ParamSetPairs{
		sdk_params.NewParamSetPair(KeySyncerDuration, &p.SyncerDuration, validateSyncerDuration),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultSyncerDuration)
}

// validate a set of params
func (p *Params) Validate() error {
	if p.SyncerDuration == 0 {
		return fmt.Errorf("validator parameter SyncerDuration must be a positive integer")
	}

	return nil
}

func validateSyncerDuration(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("validator parameter SyncerDuration must be positive: %d", v)
	}

	return nil
}
