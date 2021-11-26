package keeper

import (
	"errors"

	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandleSetRewardContractsProposal is a handler for executing a passed SetRewardContractsProposal
func HandleSetRewardContractsProposal(ctx sdk.Context, k Keeper, p *types.SetRewardContractsProposal) error {
	if err := k.CheckSetRewardContractsProposal(ctx, p); err != nil {
		return err
	}
	for _, info := range p.RewardContracts {
		k.SetRewardContract(ctx, info)

		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeSetRewardContract,
			sdk.NewAttribute(types.AttributeKeyRewardContract, info.FormatStr()),
		))
	}
	return nil
}

// CheckSetRewardContractsProposal checks the validity of an SetRewardContractsProposal
func (k Keeper) CheckSetRewardContractsProposal(ctx sdk.Context, p *types.SetRewardContractsProposal) error {
	for _, info := range p.RewardContracts {
		// Sanity check
		if info.Address == "" {
			return errors.New("empty reward contract address")
		}
	}
	return nil
}
