package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SdkAccAddrFromSgnBech32(sgnAddr string) (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(sgnAddr)
}

func SdkValAddrFromSgnBech32(sgnAddr string) (sdk.ValAddress, error) {
	acct, err := sdk.AccAddressFromBech32(sgnAddr)
	if err != nil {
		return sdk.ValAddress{}, err
	}
	return sdk.ValAddress(acct), nil
}
