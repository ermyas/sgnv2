package types

import (
	"fmt"
	time "time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultSignerUpdateDuration      time.Duration = time.Hour * 24 // 1 day
	DefaultSignAgainCoolDownDuration time.Duration = time.Minute * 10
)

var (
	KeySignerUpdateDuration      = []byte("SignerUpdateDuration")
	KeySignAgainCoolDownDuration = []byte("SignAgainCoolDownDuration")
)

var _ params.ParamSet = (*Params)(nil)

func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(signerUpdateDuration time.Duration, signAgainCoolDownDuration time.Duration) Params {
	return Params{
		SignerUpdateDuration:      signerUpdateDuration,
		SignAgainCoolDownDuration: signAgainCoolDownDuration,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeySignerUpdateDuration, p.GetSignerUpdateDuration(), validateSignerUpdateDuration),
		params.NewParamSetPair(KeySignAgainCoolDownDuration, p.GetSignAgainCoolDownDuration(), validateSignAgainCoolDownDuration),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultSignerUpdateDuration, DefaultSignAgainCoolDownDuration)
}

// validate a set of params
func (p *Params) Validate() error {
	if p.GetSignerUpdateDuration() <= 0 {
		return fmt.Errorf("validator parameter SignerUpdateDuration must be positive")
	}
	if p.GetSignAgainCoolDownDuration() <= 0 {
		return fmt.Errorf("validator parameter SignAgainCoolDownDuration must be positive")
	}
	return nil
}

func validateSignerUpdateDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("validator parameter SignerUpdateDuration must be positive: %+v", v)
	}
	return nil
}

func validateSignAgainCoolDownDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("validator parameter SignAgainCoolDownDuration must be positive: %+v", v)
	}
	return nil
}
