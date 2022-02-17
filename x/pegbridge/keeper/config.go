package keeper

import (
	"fmt"

	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetPegConfig(ctx sdk.Context, cfg types.PegConfig) error {
	err := k.CheckOrigPeggedPairsVaultVersion(ctx, cfg.OrigPeggedPairs)
	if err != nil {
		return err
	}
	for _, pair := range cfg.OrigPeggedPairs {
		k.SetOrigPeggedPair(ctx, pair)
	}
	for _, vault := range cfg.OriginalTokenVaults {
		k.SetVersionedVault(ctx, vault.Version, vault.Contract)
	}

	for _, bridge := range cfg.PeggedTokenBridges {
		k.SetVersionedBridge(ctx, bridge.Version, bridge.Contract)
	}
	return nil
}

func (k Keeper) GetConfig(ctx sdk.Context) types.PegConfig {
	vaults := k.GetAllVersionedVaults(ctx)
	k.IterateAllOriginalTokenVaults(ctx,
		func(vault commontypes.ContractInfo) (stop bool) {
			vaults = append(vaults, types.ContractInfo{Contract: vault})
			return false
		},
	)
	bridges := k.GetAllVersionedBridges(ctx)
	k.IterateAllPeggedTokenBridges(ctx,
		func(bridge commontypes.ContractInfo) (stop bool) {
			bridges = append(bridges, types.ContractInfo{Contract: bridge})
			return false
		},
	)
	pairs := k.GetAllOrigPeggedPairs(ctx)

	return types.PegConfig{
		PeggedTokenBridges:  bridges,
		OriginalTokenVaults: vaults,
		OrigPeggedPairs:     pairs,
	}
}

// Due to the limitation of fee distribution, we have to enforce that all pairs of a same original token
// on a same origian chain must use the same vault contract
func (k Keeper) CheckOrigPeggedPairsVaultVersion(ctx sdk.Context, inputPairs []types.OrigPeggedPair) error {
	pairmap := make(map[string]map[uint64]types.OrigPeggedPair) // origChainId-origAddress -> (pegChainId -> pair)

	// construct input map
	for _, p := range inputPairs {
		key := fmt.Sprintf("%d-%x", p.Orig.ChainId, p.Orig.Address)
		_, ok := pairmap[key]
		if !ok {
			pairmap[key] = make(map[uint64]types.OrigPeggedPair)
		}
		pairmap[key][p.Pegged.ChainId] = p
	}

	// add stored value to the map
	storedPairs := k.GetAllOrigPeggedPairs(ctx)
	for _, p := range storedPairs {
		key := fmt.Sprintf("%d-%x", p.Orig.ChainId, p.Orig.Address)
		_, ok := pairmap[key]
		if !ok {
			continue // this origChainId-origAddress set is not modified for the input, so ignore
		}
		_, ok = pairmap[key][p.Pegged.ChainId]
		if ok {
			continue // this origChainId-origAddress-pegChainId has new value set by the input
		} else {
			pairmap[key][p.Pegged.ChainId] = p
		}
	}

	maxVersion := uint32(1<<32 - 1)
	// now check pair map to see if it has inconsistent vault versions
	for _, pairs := range pairmap {
		// check for this origChainId-origAddress
		version := maxVersion
		for _, p := range pairs {
			if version == maxVersion {
				version = p.VaultVersion
			} else if version != p.VaultVersion {
				return fmt.Errorf("inconsistent vault version for origin chain %d token %s", p.Orig.ChainId, p.Orig.Address)
			}
		}
	}
	return nil
}
