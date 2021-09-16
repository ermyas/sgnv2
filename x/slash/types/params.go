package types

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultEnableSlash           = true
	DefaultSignedBlocksWindow    = 100
	DefaultSlashTimeout          = 10000 * 15 // in seconds
	DefaultSlashFactorDoubleSign = 5e4        // Base is 1e6
	DefaultSlashFactorDowntime   = 1e4
	DefaultJailPeriod            = 0 // blk number in jail
)

// slash params default values
var (
	DefaultMinSignedPerWindow = sdk.NewDecWithPrec(5, 1)
)

// nolint - Keys for parameter access
var (
	KeyEnableSlash           = []byte("EnableSlash")
	KeySignedBlocksWindow    = []byte("SignedBlocksWindow")
	KeySlashTimeout          = []byte("SlashTimeout")
	KeySlashFactorDoubleSign = []byte("SlashFactorDoubleSign")
	KeySlashFactorDowntime   = []byte("SlashFactorDowntime")
	KeyJailPeriod            = []byte("SlashJailPeriod")
	KeyMinSignedPerWindow    = []byte("MinSignedPerWindow")
)

// NewParams creates a new Params instance
func NewParams(enableSlash bool, signedBlocksWindow int64, slashTimeout, slashFactorDoubleSign, slashFactorDowntime, jailPeriod uint64,
	minSignedPerWindow sdk.Dec) Params {
	return Params{
		EnableSlash:           enableSlash,
		SignedBlocksWindow:    signedBlocksWindow,
		SlashTimeout:          slashTimeout,
		SlashFactorDoubleSign: slashFactorDoubleSign,
		SlashFactorDowntime:   slashFactorDowntime,
		JailPeriod:            jailPeriod,
		MinSignedPerWindow:    minSignedPerWindow,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultEnableSlash, DefaultSignedBlocksWindow, DefaultSlashTimeout,
		DefaultSlashFactorDoubleSign, DefaultSlashFactorDowntime, DefaultJailPeriod, DefaultMinSignedPerWindow)
}

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyEnableSlash, &p.EnableSlash, validateEnableSlash),
		paramtypes.NewParamSetPair(KeySignedBlocksWindow, &p.SignedBlocksWindow, validateSignedBlocksWindow),
		paramtypes.NewParamSetPair(KeySlashTimeout, &p.SlashTimeout, validateSlashTimeout),
		paramtypes.NewParamSetPair(KeySlashFactorDoubleSign, &p.SlashFactorDoubleSign, validateSlashFactorDoubleSign),
		paramtypes.NewParamSetPair(KeySlashFactorDowntime, &p.SlashFactorDowntime, validateSlashFactorDowntime),
		paramtypes.NewParamSetPair(KeyJailPeriod, &p.JailPeriod, validateJailPeriod),
		paramtypes.NewParamSetPair(KeyMinSignedPerWindow, &p.MinSignedPerWindow, validateMinSignedPerWindow),
	}
}

// Equal returns a boolean determining if two Param types are identical.
func (p Params) Equal(p2 Params) bool {
	bz1 := ModuleCdc.MustMarshalLengthPrefixed(&p)
	bz2 := ModuleCdc.MustMarshalLengthPrefixed(&p2)
	return bytes.Equal(bz1, bz2)
}

// validate a set of params
func (p Params) ValidateBasic() error {
	if err := validateEnableSlash(p.EnableSlash); err != nil {
		return err
	}
	if err := validateSignedBlocksWindow(p.SignedBlocksWindow); err != nil {
		return err
	}
	if err := validateSlashTimeout(p.SlashTimeout); err != nil {
		return err
	}
	if err := validateSlashFactorDoubleSign(p.SlashFactorDoubleSign); err != nil {
		return err
	}
	if err := validateSlashFactorDowntime(p.SlashFactorDowntime); err != nil {
		return err
	}
	if err := validateJailPeriod(p.JailPeriod); err != nil {
		return err
	}
	if err := validateMinSignedPerWindow(p.MinSignedPerWindow); err != nil {
		return err
	}
	return nil
}

func validateEnableSlash(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateSignedBlocksWindow(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("slash parameter SignedBlocksWindow must be positive: %d", v)
	}

	return nil
}

func validateSlashTimeout(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("slash parameter SlashTimeout must be positive: %d", v)
	}

	return nil
}

func validateSlashFactorDoubleSign(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("slash parameter SlashFactorDoubleSign must be positive: %d", v)
	}

	return nil
}

func validateSlashFactorDowntime(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("slash parameter SlashFactorDowntime must be positive: %d", v)
	}

	return nil
}

func validateJailPeriod(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateMinSignedPerWindow(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("slash parameter MinSignedPerWindow cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("slash parameter MinSignedPerWindow must be less or equal than 1: %s", v)
	}

	return nil
}
