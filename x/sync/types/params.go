package types

import (
	fmt "fmt"

	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultVotingPeriod   uint64  = 15 // 15 seconds
	DefaultTallyThreshold float32 = 0.667
)

var (
	KeyVotingPeriod   = []byte("VotingPeriod")
	KeyTallyThreshold = []byte("TallyThreshold")
)

var _ sdk_params.ParamSet = (*Params)(nil)

func NewParams(votingPeriod uint64, tallyThreshold float32) Params {
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

	if p.TallyThreshold == 0 {
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
	v, ok := i.(float32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("sync parameter TallyThreshold must be positive: %f", v)
	}
	if v > 1 {
		return fmt.Errorf("sync parameter TallyThreshold must be less than 1: %f", v)
	}

	return nil
}
