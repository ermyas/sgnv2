package types

import (
	"fmt"

	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyMultiChainAsset     = []byte("MultiChainAsset")
	DefaultMultiChainAsset = MultiChainTokenAsset{}
)

var _ sdk_params.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(
	multiChainAsset MultiChainTokenAsset) Params {

	return Params{
		MultiChainTokenAsset: multiChainAsset,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() sdk_params.ParamSetPairs {
	return sdk_params.ParamSetPairs{
		sdk_params.NewParamSetPair(KeyMultiChainAsset, p.GetMultiChainTokenAsset(), validateMultiChainAsset),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultMultiChainAsset)
}

// validate a set of params
func (p *Params) Validate() error {
	if len(p.GetMultiChainTokenAsset().TokenSymbolAssetMap) == 0 {
		return fmt.Errorf("validator parameter multiChainAsset must be positive")
	}
	return nil
}

func validateMultiChainAsset(i interface{}) error {
	v, ok := i.(MultiChainTokenAsset)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(v.TokenSymbolAssetMap) == 0 {
		return fmt.Errorf("validator parameter multiChainAsset must be positive: %+v", v)
	}
	return nil
}
