package keeper

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/sync/types"
	vtypes "github.com/celer-network/sgn-v2/x/validator/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) ApplyUpdate(ctx sdk.Context, update *types.PendingUpdate) bool {
	var applied bool
	var err error
	switch update.Type {
	case types.DataType_EthBlkNum:
		applied, err = k.applyEthBlkNum(ctx, update)
	case types.DataType_StakingContractParam:
		applied, err = k.applyStakingContractParam(ctx, update)
	case types.DataType_ValidatorParams:
		applied, err = k.applyValidatorParams(ctx, update)
	case types.DataType_ValidatorStates:
		applied, err = k.applyValidatorStates(ctx, update)
	case types.DataType_DelegatorShares:
		applied, err = k.applyDelegatorShares(ctx, update)
	}

	if err != nil {
		log.Errorln("Apply update err:", err)
	}
	return applied
}

func (k Keeper) applyEthBlkNum(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	return true, nil
}

func (k Keeper) applyStakingContractParam(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	return true, nil
}

func (k Keeper) applyValidatorParams(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := vtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator %s signer %s sgnaddr %s", v.EthAddress, v.EthSigner, v.SgnAddress)
	val, found := k.valKeeper.GetValidator(ctx, v.EthAddress)
	if found {
		val.EthSigner = v.EthSigner
		val.SgnAddress = v.SgnAddress
	} else {
		if v.Status == vtypes.ValidatorStatus_Bonded {
			if common.HexToAddress(update.Proposer) != common.HexToAddress(v.EthAddress) {
				log.Infof("Bonded validator %s %s not initialized, msg sender: %s", v.SgnAddress, v.EthAddress, update.Proposer)
				return true, nil
			}
			// TODO: Don't use sdkVal
			val = vtypes.NewValidator(v.EthAddress, v.EthSigner, v.SgnAddress)
			val.Description = v.Description
			val.ConsensusPubkey = v.ConsensusPubkey
			sdkValAddr, err := vtypes.SdkValAddrFromSgnBech32(val.SgnAddress)
			if err != nil {
				return false, err
			}
			sdkDescription := sdk_stakingtypes.Description{
				Moniker:         val.Description.Moniker,
				Identity:        val.Description.Identity,
				Website:         val.Description.Website,
				SecurityContact: val.Description.SecurityContact,
				Details:         val.Description.Details,
			}
			sdkVal, err := sdk_stakingtypes.NewValidator(sdkValAddr, v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey), sdkDescription)
			if err != nil {
				return false, err
			}
			k.valKeeper.SetSdkValidator(ctx, sdkVal)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					sdk_stakingtypes.EventTypeCreateValidator,
					sdk.NewAttribute(sdk_stakingtypes.AttributeKeyValidator, val.SgnAddress),
				),
			)
		} else if v.Status == vtypes.ValidatorStatus_Unbonding {
			log.Warnf("Unbonding validator %s %s not found, msg sender: %s", v.SgnAddress, v.EthAddress, update.Proposer)
			return false, nil
		} else {
			log.Debugf("Validator %s %s not bonded", v.SgnAddress, v.EthAddress)
			return true, nil
		}
	}
	val.CommissionRate = v.CommissionRate
	acct, err := vtypes.SdkAccAddrFromSgnBech32(v.SgnAddress)
	if err != nil {
		return false, err
	}
	err = k.valKeeper.InitAccount(ctx, acct)
	if err != nil {
		return false, err
	}
	k.valKeeper.SetValidator(ctx, &v)

	//TODO: gas coins
	return true, nil
}

func (k Keeper) applyValidatorStates(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := vtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator %s status %s tokens %s shares %s", v.EthAddress, v.Status, v.Tokens, v.Shares)
	err = k.valKeeper.SetValidatorStates(ctx, v.EthAddress, v.Status, v.Tokens, v.Shares)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (k Keeper) applyDelegatorShares(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	d, err := vtypes.UnmarshalDelegator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply delegator shares valAddr %s delAddr %s shares %s", d.ValAddress, d.DelAddress, d.Shares)
	err = k.valKeeper.SetDelegatorShares(ctx, d.ValAddress, d.DelAddress, d.Shares)
	if err != nil {
		return false, err
	}
	return true, nil
}
