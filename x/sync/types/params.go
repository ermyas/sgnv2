package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultVotingPeriod uint64 = 15 // 15 seconds
)

var (
	KeyVotingPeriod   = []byte("VotingPeriod")
	KeyTallyThreshold = []byte("TallyThreshold")

	DefaultTallyThreshold sdk.Dec = sdk.NewDecWithPrec(667, 3)
)

var _ sdk_params.ParamSet = (*Params)(nil)

func ParamKeyTable() sdk_params.KeyTable {
	return sdk_params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(votingPeriod uint64, tallyThreshold sdk.Dec) Params {
	return Params{votingPeriod, tallyThreshold}
}

func (p *Params) ParamSetPairs() sdk_params.ParamSetPairs {
	return sdk_params.ParamSetPairs{
		sdk_params.NewParamSetPair(KeyVotingPeriod, &p.VotingPeriod, validateVotingPeriod),
		sdk_params.NewParamSetPair(KeyTallyThreshold, &p.TallyThreshold, validateTallyThreshold),
	}
}

func DefaultParams() Params {
	return NewParams(DefaultVotingPeriod, DefaultTallyThreshold)
}

func (p *Params) Validate() error {
	if p.VotingPeriod == 0 {
		return fmt.Errorf("validator parameter VotingPeriod must be a positive integer")
	}

	if p.TallyThreshold.LTE(sdk.ZeroDec()) {
		return fmt.Errorf("validator parameter TallyThreshold must be a positive integer")
	}

	return nil
}

func validateVotingPeriod(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("sync parameter VotingPeriod must be positive: %d", v)
	}

	return nil
}

func validateTallyThreshold(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LTE(sdk.ZeroDec()) {
		return fmt.Errorf("sync parameter TallyThreshold must be positive: %s", v)
	}
	if v.GTE(sdk.OneDec()) {
		return fmt.Errorf("sync parameter TallyThreshold must be less than 1: %s", v)
	}

	return nil
}
