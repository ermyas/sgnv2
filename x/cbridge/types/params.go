package types

import (
	"fmt"

	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyMultiChainAsset     = []byte("MultiChainAsset")
	DefaultMultiChainAsset = MultiChainAsset{}
)

var _ sdk_params.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(
	multiChainAsset MultiChainAsset) Params {

	return Params{
		MultiChainAsset: multiChainAsset,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() sdk_params.ParamSetPairs {
	return sdk_params.ParamSetPairs{
		sdk_params.NewParamSetPair(KeyMultiChainAsset, p.GetMultiChainAsset(), validateMultiChainAsset),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultMultiChainAsset)
}

// validate a set of params
func (p *Params) Validate() error {
	if len(p.GetMultiChainAsset().SymbolAssetMap) == 0 {
		return fmt.Errorf("validator parameter multiChainAsset must be positive")
	}
	return nil
}

func validateMultiChainAsset(i interface{}) error {
	v, ok := i.(MultiChainAsset)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(v.SymbolAssetMap) == 0 {
		return fmt.Errorf("validator parameter multiChainAsset must be positive: %+v", v)
	}
	return nil
}
