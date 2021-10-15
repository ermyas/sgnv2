package keeper

import (
	"fmt"

	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandleAddTokensProposal is a handler for executing a passed AddTokensProposal
func HandleAddTokensProposal(ctx sdk.Context, k Keeper, p *types.AddTokensProposal) error {
	if err := k.CheckAddTokensProposal(ctx, p); err != nil {
		return err
	}
	// Create tokens
	for _, token := range p.Tokens {
		k.SetERC20Token(ctx, token)

		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeAddToken,
			sdk.NewAttribute(types.AttributeKeyToken, formatTokenName(token.Symbol, token.ChainId)),
		))
	}

	return nil
}

// CheckAddTokensProposal checks the validity of an AddTokensProposal
func (k Keeper) CheckAddTokensProposal(ctx sdk.Context, p *types.AddTokensProposal) error {
	// Check tokens not already existent
	for _, token := range p.Tokens {
		_, found := k.GetERC20Token(ctx, token.ChainId, token.Symbol)
		if found {
			return types.WrapErrTokenAlreadyExist(formatTokenName(token.Symbol, token.ChainId))
		}
	}
	return nil
}

func formatTokenName(symbol string, chainId uint64) string {
	return fmt.Sprintf("%s/%d", symbol, chainId)
}
