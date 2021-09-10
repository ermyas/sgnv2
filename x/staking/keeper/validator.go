package keeper

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cryptoenc "github.com/tendermint/tendermint/crypto/encoding"
)

func (k Keeper) GetValidator(ctx sdk.Context, ethAddr string) (validator *types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorKey(ethAddr))
	if value == nil {
		return validator, false
	}

	v := types.MustUnmarshalValidator(k.cdc, value)
	validator = &v
	return validator, true
}

func (k Keeper) GetAllValidators(ctx sdk.Context) (validators types.Validators) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		validators = append(validators, validator)
	}

	return validators
}

func (k Keeper) SetValidator(ctx sdk.Context, val *types.Validator) {
	store := ctx.KVStore(k.storeKey)
	validatorKey := types.GetValidatorKey(val.EthAddress)
	store.Set(validatorKey, types.MustMarshalValidator(k.cdc, val))
}

func (k Keeper) SetValidatorParams(ctx sdk.Context, val *types.Validator) {
	k.SetValidator(ctx, val)
	k.SetValidatorBySgnAddr(ctx, val)
	k.SetValidatorByConsAddr(ctx, val)
}

func (k Keeper) SetValidatorStates(ctx sdk.Context, val *types.Validator) {
	k.SetValidator(ctx, val)
	// update validator power
	oldPower := k.GetValidatorPower(ctx, val.EthAddress)
	newPower := val.ConsensusPower(k.PowerReduction(ctx))
	if val.GetStatus() == types.ValidatorStatus_Bonded {
		if newPower != oldPower {
			k.SetValidatorPower(ctx, val.EthAddress, newPower)
		}
	} else if oldPower > 0 {
		k.DeleteValidatorPower(ctx, val)
	}
	if newPower != oldPower {
		k.SetValidatorPowerUpdate(ctx, val.EthAddress, newPower)
	}
}

func (k Keeper) GetValidatorBySgnAddr(ctx sdk.Context, sgnAddr sdk.AccAddress) (validator *types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	ethAddr := store.Get(types.GetValidatorBySgnAddrKey(sgnAddr))
	if ethAddr == nil {
		return validator, false
	}
	return k.GetValidator(ctx, eth.Bytes2AddrHex(ethAddr))
}

func (k Keeper) SetValidatorBySgnAddr(ctx sdk.Context, val *types.Validator) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorBySgnAddrKey(val.GetSgnAddr()), val.GetEthAddr().Bytes())
	return nil
}

func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator *types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	ethAddr := store.Get(types.GetValidatorByConsAddrKey(consAddr))
	if ethAddr == nil {
		return validator, false
	}
	return k.GetValidator(ctx, eth.Bytes2AddrHex(ethAddr))
}

func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, val *types.Validator) error {
	consAddr, err := val.GetConsAddr()
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorByConsAddrKey(consAddr), val.GetEthAddr().Bytes())
	return nil
}

func (k Keeper) GetValidatorPower(ctx sdk.Context, ethAddr string) (power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorPowerKey(ethAddr))
	if bz == nil {
		return 0
	}
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshal(bz, &intV)
	return intV.GetValue()
}

func (k Keeper) SetValidatorPower(ctx sdk.Context, ethAddr string, power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: power})
	store.Set(types.GetValidatorPowerKey(ethAddr), bz)
}

func (k Keeper) DeleteValidatorPower(ctx sdk.Context, val *types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorPowerKey(val.EthAddress))
}

func (k Keeper) GetValidatorPowerUpdate(ctx sdk.Context, ethAddr string) (power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorPowerUpdateKey(ethAddr))
	if bz == nil {
		return 0
	}
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshal(bz, &intV)
	return intV.GetValue()
}

func (k Keeper) SetValidatorPowerUpdate(ctx sdk.Context, ethAddr string, power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: power})
	store.Set(types.GetValidatorPowerUpdateKey(ethAddr), bz)
}

func (k Keeper) DeleteValidatorPowerUpdate(ctx sdk.Context, val *types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorPowerUpdateKey(val.EthAddress))
}

func (k Keeper) GetBondedValidators(ctx sdk.Context) (validators types.Validators) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorPowerKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		addr := types.AddrFromValidatorKey(iterator.Key())
		validator, found := k.GetValidator(ctx, eth.Bytes2AddrHex(addr))
		if !found {
			log.Errorf("validator %x not found", addr)
			continue
		}
		validators = append(validators, *validator)
	}
	return validators
}

func (k Keeper) IterateBondedValidators(ctx sdk.Context, fn func(validator types.Validator) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorPowerKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		addr := types.AddrFromValidatorKey(iterator.Key())
		validator, found := k.GetValidator(ctx, eth.Bytes2AddrHex(addr))
		if !found {
			log.Errorf("validator %x not found", addr)
			continue
		}
		stop := fn(*validator)
		if stop {
			break
		}
	}
}

func (k Keeper) GetUpdatedValidators(ctx sdk.Context) (validators types.Validators) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorPowerUpdateKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		addr := types.AddrFromValidatorKey(iterator.Key())
		validator, found := k.GetValidator(ctx, eth.Bytes2AddrHex(addr))
		if !found {
			log.Errorf("validator %x not found", addr)
			continue
		}
		validators = append(validators, *validator)
	}
	return validators
}

// Tendermint (abci) validator updates
func (k Keeper) TmValidatorUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	powerReduction := k.PowerReduction(ctx)
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
