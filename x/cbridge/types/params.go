package types

import (
	"fmt"

	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyMultiChainAsset     = []byte("MultiChainAssetParam")
	DefaultMultiChainAsset = MultiChainAssetParam{}
)

var _ sdk_params.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(
	multiChainAssetParam MultiChainAssetParam) Params {

	return Params{
		MultiChainAssetParam: multiChainAssetParam,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() sdk_params.ParamSetPairs {
	return sdk_params.ParamSetPairs{
		sdk_params.NewParamSetPair(KeyMultiChainAsset, p.GetMultiChainAssetParam(), validateMultiChainAsset),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultMultiChainAsset)
}

// validate a set of params
func (p *Params) Validate() error {
	if len(p.GetMultiChainAssetParam().ChainAsset) == 0 {
		return fmt.Errorf("validator parameter multiChainAsset must be positive")
	}
	return nil
}

func validateMultiChainAsset(i interface{}) error {
	v, ok := i.(MultiChainAssetParam)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(v.ChainAsset) == 0 {
		return fmt.Errorf("validator parameter multiChainAsset must be positive: %+v", v)
	}
	return nil
}
