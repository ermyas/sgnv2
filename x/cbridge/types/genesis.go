package types

// this line is used by starport scaffolding # genesis/types/import
// this line is used by starport scaffolding # ibc/genesistype/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return NewGenesisState(DefaultParams())
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}
