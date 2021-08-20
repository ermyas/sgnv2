package validator

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/contracts"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk_staking_keepr "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc           codec.Codec  // The wire codec for binary encoding/decoding.
	storeKey      sdk.StoreKey // Unexposed key to access store from sdk.Context
	accountKeeper ante.AccountKeeper
	stakingKeeper sdk_staking_keepr.Keeper
	paramstore    sdk_params.Subspace
}

// NewKeeper creates new instances of the validator Keeper
func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey,
	accountKeeper ante.AccountKeeper, stakingKeeper sdk_staking_keepr.Keeper, paramstore sdk_params.Subspace) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		accountKeeper: accountKeeper,
		stakingKeeper: stakingKeeper,
		paramstore:    paramstore.WithKeyTable(ParamKeyTable()),
	}
}

// Get validators metadata
func (k Keeper) GetBondedSgnValidators(ctx sdk.Context) []sdk_staking.Validator {
	return k.stakingKeeper.GetBondedValidatorsByPower(ctx)
}

// Get a validator by consencus address
func (k Keeper) GetSgnValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) (sdk_staking.Validator, bool) {
	return k.stakingKeeper.GetValidatorByConsAddr(ctx, addr)
}

// Get a validator by validator account address
func (k Keeper) GetSgnValidator(ctx sdk.Context, addr sdk.ValAddress) (sdk_staking.Validator, bool) {
	return k.stakingKeeper.GetValidator(ctx, addr)
}

// func (k Keeper) IterateBondedValidatorsByPower(ctx sdk.Context, fn func(index int64, validator exported.ValidatorI) (stop bool)) {
// 	k.stakingKeeper.IterateBondedValidatorsByPower(ctx, fn)
// }

// Get the entire Syncer metadata
func (k Keeper) GetSyncer(ctx sdk.Context) types.Syncer {
	store := ctx.KVStore(k.storeKey)

	if !store.Has(SyncerKey) {
		return Syncer{}
	}

	//value := store.Get(SyncerKey)
	var syncer Syncer
	//k.cdc.MustUnmarshalBinaryBare(value, &syncer)
	return syncer
}

// Sets the entire Syncer metadata
func (k Keeper) SetSyncer(ctx sdk.Context, syncer Syncer) {
	//store := ctx.KVStore(k.storeKey)
	//store.Set(SyncerKey, k.cdc.MustMarshalBinaryBare(syncer))
}

// Get the entire Delegator metadata for a validatorAddr and delegatorAddr
func (k Keeper) GetDelegator(ctx sdk.Context, validatorAddr, delegatorAddr string) (delegator Delegator, found bool) {
	store := ctx.KVStore(k.storeKey)

	if !store.Has(GetDelegatorKey(validatorAddr, delegatorAddr)) {
		return delegator, false
	}

	//value := store.Get(GetDelegatorKey(validatorAddr, delegatorAddr))
	//k.cdc.MustUnmarshalBinaryBare(value, &delegator)
	return delegator, true
}

// Get the list of all delegators
func (k Keeper) GetAllDelegators(ctx sdk.Context, validatorAddr string) (delegators []Delegator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, GetDelegatorsKey(validatorAddr))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var delegator Delegator
		//k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &delegator)
		delegators = append(delegators, delegator)
	}
	return delegators
}

// Sets the entire Delegator metadata for a validatorAddr and delegatorAddr
func (k Keeper) SetDelegator(ctx sdk.Context, delegator Delegator) {
	//store := ctx.KVStore(k.storeKey)
	//store.Set(GetDelegatorKey(delegator.ValidatorAddr, delegator.DelegatorAddr), k.cdc.MustMarshalBinaryBare(delegator))
}

func (k Keeper) RemoveDelegator(ctx sdk.Context, delegator Delegator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(GetDelegatorKey(delegator.ValidatorAddr, delegator.DelegatorAddr))
}

// Get the entire Validator metadata
func (k Keeper) GetValidator(ctx sdk.Context, validatorAddr string) (validator Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	validatorKey := GetValidatorKey(validatorAddr)

	if !store.Has(validatorKey) {
		return validator, false
	}

	//value := store.Get(validatorKey)
	//k.cdc.MustUnmarshalBinaryBare(value, &validator)
	return validator, true
}

// Get the list of all validators
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []Validator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, ValidatorKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var validator Validator
		//k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &validator)
		validators = append(validators, validator)
	}
	return validators
}

// Sets the Validator metadata
func (k Keeper) SetValidator(ctx sdk.Context, validator Validator) {
	//store := ctx.KVStore(k.storeKey)
	//validatorKey := GetValidatorKey(validator.EthAddress)
	//store.Set(validatorKey, k.cdc.MustMarshalBinaryBare(validator))
}

// Get the entire Reward metadata for ethAddress
func (k Keeper) GetReward(ctx sdk.Context, ethAddress string) (Reward, bool) {
	store := ctx.KVStore(k.storeKey)
	rewardKey := GetRewardKey(ethAddress)

	if !store.Has(rewardKey) {
		return NewReward(ethAddress), false
	}

	var reward Reward
	//value := store.Get(rewardKey)
	//k.cdc.MustUnmarshalBinaryBare(value, &reward)
	return reward, true
}

// IterateRewards iterates over the stored penalties
func (k Keeper) IterateRewards(ctx sdk.Context,
	handler func(reward Reward) (stop bool)) {

	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, RewardKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var reward Reward
		//k.cdc.MustUnmarshalBinaryBare(iter.Value(), &reward)
		if handler(reward) {
			break
		}
	}
}

// GetRewards returns all the rewards from store
func (keeper Keeper) GetRewards(ctx sdk.Context) (rewards []Reward) {
	keeper.IterateRewards(ctx, func(reward Reward) bool {
		rewards = append(rewards, reward)
		return false
	})
	return
}

// Sets the Reward metadata for ethAddress
func (k Keeper) SetReward(ctx sdk.Context, reward Reward) {
	//store := ctx.KVStore(k.storeKey)
	//store.Set(GetRewardKey(reward.Receiver), k.cdc.MustMarshalBinaryBare(reward))
}

// AddReward add reward to a specific ethAddress
func (k Keeper) AddReward(ctx sdk.Context, ethAddress string, reward sdk.Int) {
}

func (k Keeper) distributeEpochReward(ctx sdk.Context) {

}

func (k Keeper) distributeValidatorReward(ctx sdk.Context) {
	cycleLen := k.EpochLength(ctx) * 2
	validators := k.GetBondedSgnValidators(ctx)
	var idx uint
	if uint(len(validators)) >= cycleLen {
		idx = uint(ctx.BlockHeight()) % uint(len(validators))
	} else {
		skip := cycleLen/uint(len(validators)) + 1
		if uint(ctx.BlockHeight())%skip != 0 {
			return
		}
		idx = uint(ctx.BlockHeight()) / skip % uint(len(validators))
	}
	ethAddr := contracts.FormatAddrHex(validators[idx].Description.Identity)
	k.DistributeValidatorPendingReward(ctx, ethAddr)
}

func (k Keeper) DistributeValidatorPendingReward(ctx sdk.Context, ethAddress string) {
	log.Debugln("Distribute pending reward for validator", ethAddress)
}

// Distribute epoch rewards to all validators and delegators
func (k Keeper) DistributeReward(ctx sdk.Context) {
	k.distributeEpochReward(ctx)
	k.distributeValidatorReward(ctx)
}

func (k Keeper) InitAccount(ctx sdk.Context, accAddress sdk.AccAddress) {
	err := sdk.VerifyAddressFormat(accAddress)
	if err != nil {
		log.Errorf("InitAccount %s err: %s", accAddress, err)
		return
	}
	account := k.accountKeeper.GetAccount(ctx, accAddress)
	if account == nil {
		log.Infof("Set new account %s", accAddress)
		//account = k.accountKeeper.NewAccountWithAddress(ctx, accAddress)
		k.accountKeeper.SetAccount(ctx, account)
	}
}

func (k Keeper) RemoveAccount(ctx sdk.Context, accAddress sdk.AccAddress) {
	account := k.accountKeeper.GetAccount(ctx, accAddress)
	if account != nil {
		log.Infof("Remove account %s", accAddress)
		//k.accountKeeper.RemoveAccount(ctx, account)
	}
}
