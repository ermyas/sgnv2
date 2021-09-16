package types

func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}

func ValidateGenesis(data *GenesisState) error {
	return data.Params.ValidateBasic()
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParams())
}
