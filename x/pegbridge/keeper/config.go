package keeper

import (
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetPegConfig(ctx sdk.Context, cfg types.PegConfig) {
	for _, vault := range cfg.OriginalTokenVaults {
		k.SetOriginalTokenVault(ctx, vault)
	}
	for _, bridge := range cfg.PeggedTokenBridges {
		k.SetPeggedTokenBridge(ctx, bridge)
	}
	for _, pair := range cfg.OrigPeggedPairs {
		k.SetOrigPeggedPair(ctx, pair)
	}
}
