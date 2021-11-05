package keeper

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/celer-network/sgn-v2/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ApplyUpdate(ctx sdk.Context, update *types.PendingUpdate) bool {
	var applied bool
	var err error
	switch update.Type {
	case types.DataType_ValidatorSgnAddr:
		applied, err = k.applyValidatorSgnAddr(ctx, update)
	case types.DataType_ValidatorParams:
		applied, err = k.applyValidatorParams(ctx, update)
	case types.DataType_ValidatorStates:
		applied, err = k.applyValidatorStates(ctx, update)
	case types.DataType_DelegatorShares:
		applied, err = k.applyDelegatorShares(ctx, update)
	case types.DataType_CbrOnchainEvent:
		applied, err = k.cbrKeeper.ApplyEvent(ctx, update.Data)
	case types.DataType_CbrUpdateCbrPrice:
		applied, err = k.cbrKeeper.ApplyUpdateCbrPrice(ctx, update.Data)
	}

	if err != nil {
		log.Errorln("Apply update err:", err)
	}
	return applied
}

func (k Keeper) applyValidatorSgnAddr(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator sgn addr %s", v.String())
	// TODO: remove previous sgnaddr account
	acct, err := sdk.AccAddressFromBech32(v.SgnAddress)
	if err != nil {
		return false, err
	}
	err = k.stakingKeeper.InitAccount(ctx, acct)
	if err != nil {
		return false, err
	}
	return true, nil
}

// TODO: allow sgnaddr/consaddr update
func (k Keeper) applyValidatorParams(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	if update.Proposer != v.SgnAddress {
		return false, fmt.Errorf("validator %s not msg sender: %s", v.EthAddress, update.Proposer)
	}
	if v.ConsensusPubkey == nil {
		return false, fmt.Errorf("empty consensus pub key")
	}
	consAddr, err := v.GetConsAddr()
	if err != nil {
		return false, fmt.Errorf("validator %s failed to get consensus addr, err %s", v.EthAddress, err)
	}
	log.Infof("Apply validator params %s", v.String())
	val, found := k.stakingKeeper.GetValidator(ctx, eth.Hex2Addr(v.EthAddress))
	if !found {
		val = *stakingtypes.NewValidator(v.EthAddress, v.EthSigner, v.SgnAddress)
		val.Description = v.Description
	} else {
		if v.SgnAddress != val.SgnAddress {
			return false, fmt.Errorf("update of sgnaddr is not supported: %s %s %s", v.EthAddress, v.SgnAddress, val.SgnAddress)
		}
		storedConsAddr, err := val.GetConsAddr()
		if err != nil {
			return false, fmt.Errorf("validator %s failed to get stored consensus addr, err %s", v.EthAddress, err)
		}
		if !consAddr.Equals(storedConsAddr) {
			return false, fmt.Errorf("update of consaddr is not supported: %s %s %s", v.EthAddress, consAddr, storedConsAddr)
		}
		if val.Status == stakingtypes.Bonded {
			if val.EthSigner != eth.FormatAddrHex(v.EthSigner) {
				log.Infof("Update bonded validator %s signer from %s to %s", val.EthAddress, val.EthSigner, v.EthSigner)
				k.cbrKeeper.UpdateLatestSigners(ctx, true)
			}
		}
		val.EthSigner = eth.FormatAddrHex(v.EthSigner)
		val.SgnAddress = v.SgnAddress
	}
	val.ConsensusPubkey = v.ConsensusPubkey
	val.CommissionRate = v.CommissionRate
	k.stakingKeeper.SetValidatorParams(ctx, &val, !found)
	// TODO: gas coins
	return true, nil
}

func (k Keeper) applyValidatorStates(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator states %s", v.String())
	val, found := k.stakingKeeper.GetValidator(ctx, eth.Hex2Addr(v.EthAddress))
	if !found {
		return false, fmt.Errorf("validator %s not found", val.EthAddress)
	}
	val.Status = v.Status
	val.Tokens = v.Tokens
	val.DelegatorShares = v.DelegatorShares

	k.stakingKeeper.SetValidatorStates(ctx, &val)
	return true, nil
}

func (k Keeper) applyDelegatorShares(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	d, err := stakingtypes.UnmarshalDelegation(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply delegator shares valAddr %s delAddr %s shares %s", d.ValidatorAddress, d.DelegatorAddress, d.Shares)
	valAddr := eth.Hex2Addr(d.ValidatorAddress)
	val, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if found && !val.DelegatorShares.Equal(sdk.ZeroInt()) {
		err = k.stakingKeeper.SetDelegationShares(
			ctx,
			eth.Hex2Addr(d.DelegatorAddress),
			eth.Hex2Addr(d.ValidatorAddress),
			d.Shares)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	// If validator not found or its DelegatorShares is zero, the delegation is processed before we have done the initial sync of the
	// validator, just try again later.
	return false, nil
}
