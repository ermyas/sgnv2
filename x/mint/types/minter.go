package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultAnnualProvision = sdk.NewDecFromBigInt(big.NewInt(30000000)).Mul(sdk.NewDecFromBigInt(big.NewInt(1e18)))
)

// NewMinter returns a new Minter object with the given annual
// provisions values.
func NewMinter(annualProvisions sdk.Dec) Minter {
	return Minter{
		AnnualProvisions: annualProvisions,
	}
}

// DefaultMinter returns a default Minter object for a new chain.
func DefaultMinter() Minter {
	return NewMinter(DefaultAnnualProvision)
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params) sdk.Coin {
	provisionAmt := m.AnnualProvisions.QuoInt(sdk.NewInt(int64(params.BlocksPerYear)))
	return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
}
