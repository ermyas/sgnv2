package types

import (
	"github.com/celer-network/sgn-v2/eth"
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DelegationI delegation bond for a delegated proof of stake system
type DelegationI interface {
	GetDelegatorAddr() eth.Addr // delegator eth.Addr for the bond
	GetValidatorAddr() eth.Addr // validator ETH address
	GetShares() sdk.Int         // amount of validator's shares held in this delegation
}

// ValidatorI expected validator functions
type ValidatorI interface {
	GetMoniker() string                                // moniker of the validator
	GetStatus() BondStatus                             // status of the validator
	IsBonded() bool                                    // check if has a bonded status
	IsUnbonded() bool                                  // check if has status unbonded
	IsUnbonding() bool                                 // check if has status unbonding
	GetEthAddr() eth.Addr                              // ETH address
	GetSgnAddr() sdk.AccAddress                        // SGN account address
	GetOperator() sdk.ValAddress                       // operator address to receive/return validators coins
	ConsPubKey() (cryptotypes.PubKey, error)           // validation consensus pubkey (cryptotypes.PubKey)
	TmConsPublicKey() (tmprotocrypto.PublicKey, error) // validation consensus pubkey (Tendermint)
	GetConsAddr() (sdk.ConsAddress, error)             // validation consensus address
	GetTokens() sdk.Int                                // validation tokens
	GetBondedTokens() sdk.Int                          // validator bonded tokens
	GetConsensusPower(sdk.Int) int64                   // validation power in tendermint
	GetCommission() sdk.Dec                            // validator commission rate
	GetDelegatorShares() sdk.Int                       // total outstanding delegator shares
	TokensFromShares(sdk.Int) sdk.Dec                  // token worth of provided delegator shares
	TokensFromSharesTruncated(sdk.Int) sdk.Dec         // token worth of provided delegator shares, truncated
	TokensFromSharesRoundUp(sdk.Int) sdk.Dec           // token worth of provided delegator shares, rounded up
}
