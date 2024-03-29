package common

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Bech32MainPrefix = "sgn"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32MainPrefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32MainPrefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic

	// MaxAddrLen is the maximum allowed length (in bytes) for an address.
	//
	// NOTE: In the SDK, the default value is 255.
	MaxAddrLen = 20

	QuotaCoinName = "quota"

	CelrDenom     = "CELR"
	CelrPrecision = 1e18
	GrpcTimeOut   = 10 * time.Second
)

var (
	StakingScaleFactor = sdk.NewDec(1e12)
)
