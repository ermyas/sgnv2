package types

import (
	"fmt"
	"time"

	yaml "gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	ParamStoreKeyClaimCooldown = []byte("claimcooldown")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		ClaimCooldown: time.Minute, // 1 minute
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyClaimCooldown, &p.ClaimCooldown, validateClaimCooldown),
	}
}

// Validate performs validation on farming parameters.
func (p Params) Validate() error {
	if err := validateClaimCooldown(p.ClaimCooldown); err != nil {
		return err
	}
	return nil
}

func validateClaimCooldown(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("claim cooldown must be positive: %d", v)
	}
	return nil
}
