package keeper

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cryptoenc "github.com/tendermint/tendermint/crypto/encoding"
)

// Modify based on https://github.com/cosmos/cosmos-sdk/blob/v0.43.0/x/staking/keeper/val_state_change.go

// BlockValidatorUpdates calculates the ValidatorUpdates for the current block
// Called in each EndBlock
func (k Keeper) BlockValidatorUpdates(ctx sdk.Context) []abci.ValidatorUpdate {
	updates := k.GetValidatorPowerUpdates(ctx)
	return updates
}

func (k Keeper) GetValidatorPowerUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	powerReduction := k.PowerReduction(ctx) // TODO: make it a constant
	updatedVals := k.GetUpdatedValidators(ctx)
	for _, val := range updatedVals {
		if val.GetStatus() == types.ValidatorStatus_Bonded {
			updates = append(updates, val.ABCIValidatorUpdate(powerReduction))
		} else {
			updates = append(updates, val.ABCIValidatorUpdateZero())
		}
		k.DeleteValidatorPowerUpdate(ctx, &val)
	}
	if len(updates) > 0 {
		log.Infof("update tendermint validator: %s", printTmUpdates(updates))
	}
	return
}

func printTmUpdates(updates []abci.ValidatorUpdate) string {
	var out string
	for _, v := range updates {
		pub, err := cryptoenc.PubKeyFromProto(v.PubKey)
		if err != nil {
			out += fmt.Sprintf("%s | ", err)
		}

		out += fmt.Sprintf("consaddr %s, power %d | ", sdk.ConsAddress(pub.Address()).String(), v.Power)
	}
	return out
}
