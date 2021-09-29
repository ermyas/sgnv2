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

func (k Keeper) GetValidator(ctx sdk.Context, ethAddr eth.Addr) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetValidatorKey(ethAddr))
	if value == nil {
		return validator, false
	}
	validator = types.MustUnmarshalValidator(k.cdc, value)
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
	validatorKey := types.GetValidatorKey(eth.Hex2Addr(val.EthAddress))
	store.Set(validatorKey, types.MustMarshalValidator(k.cdc, val))
}

func (k Keeper) SetValidatorParams(ctx sdk.Context, val *types.Validator, newValidator bool) {
	k.SetValidator(ctx, val)
	k.SetValidatorBySgnAddr(ctx, val)
	k.SetValidatorByConsAddr(ctx, val)

	if newValidator {
		valAddr := eth.Hex2Addr(val.EthAddress)
		k.AfterValidatorCreated(ctx, valAddr)
	}
}

func (k Keeper) SetValidatorStates(ctx sdk.Context, val *types.Validator) {
	k.SetValidator(ctx, val)
	// update validator power
	oldPower := k.GetValidatorPower(ctx, eth.Hex2Addr(val.EthAddress))
	newPower := val.ConsensusPower(k.PowerReduction(ctx))
	if val.GetStatus() == types.Bonded {
		if newPower != oldPower {
			k.SetValidatorPower(ctx, eth.Hex2Addr(val.EthAddress), newPower)
		}
	} else if oldPower > 0 {
		k.DeleteValidatorPower(ctx, val)
	}
	if newPower != oldPower {
		k.SetValidatorPowerUpdate(ctx, eth.Hex2Addr(val.EthAddress), newPower)
	}
	if oldPower == 0 && val.GetStatus() == types.Bonded {
		k.AfterValidatorBonded(ctx, val.GetEthAddr())
	} else if newPower == 0 && oldPower > 0 {
		k.AfterValidatorBeginUnbonding(ctx, val.GetEthAddr())
	}
}

func (k Keeper) GetValidatorBySgnAddr(ctx sdk.Context, sgnAddr sdk.AccAddress) (validator types.ValidatorI, found bool) {
	store := ctx.KVStore(k.storeKey)
	ethAddr := store.Get(types.GetValidatorBySgnAddrKey(sgnAddr))
	if ethAddr == nil {
		return validator, false
	}
	return k.GetValidator(ctx, eth.Bytes2Addr(ethAddr))
}

func (k Keeper) SetValidatorBySgnAddr(ctx sdk.Context, val *types.Validator) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorBySgnAddrKey(val.GetSgnAddr()), val.GetEthAddr().Bytes())
	return nil
}

func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	ethAddr := store.Get(types.GetValidatorByConsAddrKey(consAddr))
	if ethAddr == nil {
		return validator, false
	}
	return k.GetValidator(ctx, eth.Bytes2Addr(ethAddr))
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

func (k Keeper) GetValidatorPower(ctx sdk.Context, ethAddr eth.Addr) (power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorPowerKey(ethAddr))
	if bz == nil {
		return 0
	}
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshal(bz, &intV)
	return intV.GetValue()
}

func (k Keeper) SetValidatorPower(ctx sdk.Context, ethAddr eth.Addr, power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: power})
	store.Set(types.GetValidatorPowerKey(ethAddr), bz)
}

func (k Keeper) DeleteValidatorPower(ctx sdk.Context, val *types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorPowerKey(eth.Hex2Addr(val.EthAddress)))
}

func (k Keeper) GetValidatorPowerUpdate(ctx sdk.Context, ethAddr eth.Addr) (power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorPowerUpdateKey(ethAddr))
	if bz == nil {
		return 0
	}
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshal(bz, &intV)
	return intV.GetValue()
}

func (k Keeper) SetValidatorPowerUpdate(ctx sdk.Context, ethAddr eth.Addr, power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: power})
	store.Set(types.GetValidatorPowerUpdateKey(ethAddr), bz)
}

func (k Keeper) DeleteValidatorPowerUpdate(ctx sdk.Context, val *types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorPowerUpdateKey(eth.Hex2Addr(val.EthAddress)))
}

func (k Keeper) GetBondedValidators(ctx sdk.Context) (validators types.Validators) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorPowerKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		addr := types.AddrFromValidatorKey(iterator.Key())
		validator, found := k.GetValidator(ctx, eth.Bytes2Addr(addr))
		if !found {
			log.Errorf("validator %x not found", addr)
			continue
		}
		validators = append(validators, validator)
	}
	return validators
}

func (k Keeper) IterateBondedValidators(ctx sdk.Context, fn func(index int64, validator types.ValidatorI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorPowerKey)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		addr := types.AddrFromValidatorKey(iterator.Key())
		validator, found := k.GetValidator(ctx, eth.Bytes2Addr(addr))
		if !found {
			log.Errorf("validator %x not found", addr)
			continue
		}
		stop := fn(i, validator)
		if stop {
			break
		}
		i++
	}
}

func (k Keeper) GetUpdatedValidators(ctx sdk.Context) (validators types.Validators) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorPowerUpdateKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		addr := types.AddrFromValidatorKey(iterator.Key())
		validator, found := k.GetValidator(ctx, eth.Bytes2Addr(addr))
		if !found {
			log.Errorf("validator %x not found", addr)
			continue
		}
		validators = append(validators, validator)
	}
	return validators
}

// get the list of Tendermint abci.ValidatorUpdate
func (k Keeper) TmValidatorUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	powerReduction := k.PowerReduction(ctx)
	updatedVals := k.GetUpdatedValidators(ctx)
	for _, val := range updatedVals {
		if val.GetStatus() == types.Bonded {
			updates = append(updates, val.ABCIValidatorUpdate(powerReduction))
		} else {
			updates = append(updates, val.ABCIValidatorUpdateZero())
		}
		k.DeleteValidatorPowerUpdate(ctx, &val)
	}
	if len(updates) > 0 {
		var out string
		for _, v := range updates {
			pub, err := cryptoenc.PubKeyFromProto(v.PubKey)
			if err != nil {
				out += fmt.Sprintf("%s | ", err)
			}
			out += fmt.Sprintf("consaddr %s, power %d | ", sdk.ConsAddress(pub.Address()).String(), v.Power)
		}
		log.Infof("update tendermint validator: %s", out)
	}
	return
}

func (k Keeper) GetTransactors(ctx sdk.Context, ethAddr eth.Addr) (txs types.ValidatorTransactors) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetValidatorTransactorsKey(ethAddr))
	if value == nil {
		return
	}
	err := k.cdc.Unmarshal(value, &txs)
	if err != nil {
		log.Error(err)
	}
	return
}

func (k Keeper) SetTransactors(
	ctx sdk.Context, op types.SetTransactorsOp, sgnAddr sdk.AccAddress, transactors []string) error {
	// TODO: support other ops
	if op != types.SetTransactorsOp_Overwrite {
		return fmt.Errorf("only support overwrite for now")
	}

	validator, found := k.GetValidatorBySgnAddr(ctx, sgnAddr)
	if !found {
		return fmt.Errorf("validator not found")
	}
	if validator.GetStatus() != types.Bonded {
		return fmt.Errorf("validator not bonded")
	}

	currTransactors := k.GetTransactors(ctx, validator.GetEthAddr())
	txrs := make(map[string]bool)
	for _, transactor := range transactors {
		acct, err := sdk.AccAddressFromBech32(transactor)
		if err != nil {
			return err
		}
		if acct.Equals(sgnAddr) {
			return fmt.Errorf("transactor cannot be validator sgn addr")
		}
		if _, exist := txrs[transactor]; exist {
			return fmt.Errorf("duplicated transactor %s", transactor)
		}
		txrs[transactor] = true
		k.InitAccount(ctx, acct)
		// TODO: set quota coins
	}

	for _, transactor := range currTransactors.Transactors {
		if _, exist := txrs[transactor]; !exist {
			acct, err := sdk.AccAddressFromBech32(transactor)
			if err != nil {
				log.Errorln(transactor, err)
				continue
			}
			k.RemoveAccount(ctx, acct)
		}
	}

	txsproto := &types.ValidatorTransactors{Transactors: transactors}
	store := ctx.KVStore(k.storeKey)
	validatorTransactorsKey := types.GetValidatorTransactorsKey(validator.GetEthAddr())
	store.Set(validatorTransactorsKey, k.cdc.MustMarshal(txsproto))
	return nil
}
