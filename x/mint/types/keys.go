package types

// MinterKey is the key to use for the keeper store.
var MinterKey = []byte{0x00}

const (
	// module name
	ModuleName = "mint"

	// StoreKey is the default store key for mint
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the minting store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the minting querier
	QueryParameters       = "parameters"
	QueryAnnualProvisions = "annual_provisions"
)
