package types

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params,
	config PegConfig,
) *GenesisState {
	return &GenesisState{
		Params: params,
		Config: config,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		Config: PegConfig{},
	}
}

// ValidateGenesis validates the farming genesis parameters
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

	// TODO: Actually validate?
	return nil
}
