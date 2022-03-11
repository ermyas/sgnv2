package pegbridge

import (
	"fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	pegkeeper "github.com/celer-network/sgn-v2/x/pegbridge/keeper"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewPegProposalHandler(k pegkeeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.PegProposal:
			return handlePegProposal(ctx, k, c)
		case *types.PairDeleteProposal:
			return handlePairDeleteProposal(ctx, k, c)
		case *types.TotalSupplyUpdateProposal:
			return handleTotalSupplyUpdateProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unsupported peg proposal content type: %T", c)
		}
	}
}

func handlePegProposal(ctx sdk.Context, k pegkeeper.Keeper, p *types.PegProposal) error {
	if err := p.PegConfig.Validate(); err != nil {
		return err
	}
	return k.SetPegConfig(ctx, *p.PegConfig)
}

func handlePairDeleteProposal(ctx sdk.Context, k pegkeeper.Keeper, p *types.PairDeleteProposal) error {
	pair := p.PairToDelete
	if err := pair.ValidateBasic(); err != nil {
		return err
	}
	k.DeleteOrigPeggedPair(ctx, pair.Orig.ChainId, pair.Orig.Address, pair.Pegged.ChainId, eth.Hex2Addr(pair.Pegged.Address))
	return nil
}

// this proposal is only used for backward compatibility (manually set the total supply for pegged tokens
// that were already supported before this supply tracking feature is launched).
func handleTotalSupplyUpdateProposal(ctx sdk.Context, k pegkeeper.Keeper, p *types.TotalSupplyUpdateProposal) error {
	inputPair := p.Pair
	if inputPair == nil {
		return fmt.Errorf("no pair info in proposal")
	}
	if err := inputPair.ValidateBasic(); err != nil {
		return err
	}
	totalSupply, ok := new(big.Int).SetString(p.TotalSupply, 10)
	if !ok || totalSupply.Sign() == -1 {
		return fmt.Errorf("invalid total supply string")
	}
	expectedPair, found := k.GetOrigPeggedPair(ctx, inputPair.Orig.ChainId, inputPair.Orig.Address, inputPair.Pegged.ChainId)
	if !found {
		return fmt.Errorf("no pair found")
	}
	if expectedPair.SupplyCap != "" {
		supplyCap, _ := new(big.Int).SetString(expectedPair.SupplyCap, 10)
		if supplyCap.Sign() == 1 && totalSupply.Cmp(supplyCap) > 0 {
			return fmt.Errorf("invalid total supply, must be smaller than supply cap")
		}
	}
	k.SetTotalSupply(ctx, expectedPair.Pegged.ChainId, eth.Hex2Addr(expectedPair.Pegged.Address), totalSupply)
	return nil
}
