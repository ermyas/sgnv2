package keeper

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
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
		validators = append(validators, &validator)
	}

	return validators
}

func (k Keeper) SetValidator(ctx sdk.Context, val *types.Validator) {
	store := ctx.KVStore(k.storeKey)
	validatorKey := types.GetValidatorKey(val.EthAddress)
	store.Set(validatorKey, types.MustMarshalValidator(k.cdc, val))
}

func (k Keeper) SetValidatorStates(ctx sdk.Context, val *types.Validator) error {
	k.SetValidator(ctx, val)
	return k.UpdateSdkValidator(ctx, val)
}

func (k Keeper) UpdateSdkValidator(ctx sdk.Context, val *types.Validator) error {
	tkInt, ok := sdk.NewIntFromString(val.Tokens)
	if !ok {
		return fmt.Errorf("invalid tokens %s", val.Tokens)
	}
	shInt, ok := sdk.NewIntFromString(val.Shares)
	if !ok {
		return fmt.Errorf("invalid shares %s", val.Shares)
	}

	sdkValAddr, err := types.SdkValAddrFromSgnBech32(val.SgnAddress)
	if err != nil {
		return err
	}
	sdkVal, found := k.sdkStakingKeeper.GetValidator(ctx, sdkValAddr)
	if !found {
		if val.Status == types.ValidatorStatus_Bonded {
			if val.ConsensusPubkey == nil {
				return fmt.Errorf("validator %s consensu pubkey not set", val.EthAddress)
			}
			sdkDescription := sdk_staking.Description{
				Moniker:         val.Description.Moniker,
				Identity:        val.Description.Identity,
				Website:         val.Description.Website,
				SecurityContact: val.Description.SecurityContact,
				Details:         val.Description.Details,
			}
			sdkValAddr, err2 := types.SdkValAddrFromSgnBech32(val.SgnAddress)
			if err2 != nil {
				return fmt.Errorf("invalid sgnAddr %s, err %w", val.SgnAddress, err2)
			}
			sdkVal = sdk_staking.Validator{
				OperatorAddress: sdkValAddr.String(),
				ConsensusPubkey: val.ConsensusPubkey,
				Status:          sdk_staking.Bonded,
				Tokens:          tkInt,
				DelegatorShares: shInt.ToDec(),
				Description:     sdkDescription,
			}
			err = k.sdkStakingKeeper.SetValidatorByConsAddr(ctx, sdkVal)
			if err != nil {
				return fmt.Errorf("SetValidatorByConsAddr %s %s, err %w", val.SgnAddress, sdkVal.OperatorAddress, err)
			}
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					sdk_staking.EventTypeCreateValidator,
					sdk.NewAttribute(sdk_staking.AttributeKeyValidator, val.SgnAddress),
				),
			)
		} else if val.Status == types.ValidatorStatus_Unbonding {
			log.Warnf("Unbonding sdk validator %s %s not found", val.EthAddress, val.SgnAddress)
			return nil
		} else {
			log.Debugf("Validator %s %s %s not bonded, status:", val.EthAddress, val.SgnAddress, val.Status.String())
			return nil
		}
	}

	k.sdkStakingKeeper.DeleteValidatorByPowerIndex(ctx, sdkVal)
	sdkVal.Status = sdk_staking.BondStatus(val.Status)
	if sdkVal.Status == sdk_staking.Unbonded {
		sdkVal.Tokens = sdk.ZeroInt()
	} else {
		sdkVal.Tokens = tkInt
	}
	sdkVal.DelegatorShares = shInt.ToDec()
	k.sdkStakingKeeper.SetValidator(ctx, sdkVal)

	if val.Status == types.ValidatorStatus_Bonded {
		k.sdkStakingKeeper.SetNewValidatorByPowerIndex(ctx, sdkVal)
	} else if val.Status == types.ValidatorStatus_Unbonded {
		log.Infof("remove sdk validator %s %s", val.EthAddress, val.SgnAddress)
		k.sdkStakingKeeper.RemoveValidator(ctx, sdkValAddr)
	}
	return nil
}

// Get sdk validators
func (k Keeper) GetBondedSdkValidators(ctx sdk.Context) []sdk_staking.Validator {
	return k.sdkStakingKeeper.GetBondedValidatorsByPower(ctx)
}

// Get a sdk validator by consensus address
func (k Keeper) GetSdkValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) (sdk_staking.Validator, bool) {
	return k.sdkStakingKeeper.GetValidatorByConsAddr(ctx, addr)
}

// Get a sdk validator by validator account address
func (k Keeper) GetSdkValidator(ctx sdk.Context, addr sdk.ValAddress) (sdk_staking.Validator, bool) {
	return k.sdkStakingKeeper.GetValidator(ctx, addr)
}
