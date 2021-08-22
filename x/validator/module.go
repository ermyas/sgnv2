package validator

import (
	"encoding/json"

	"github.com/celer-network/sgn-v2/x/validator/keeper"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// app module Basics object
type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	//RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {

}

func (AppModuleBasic) RegisterInterfaces(codectypes.InterfaceRegistry) {

}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {

}

func (AppModuleBasic) DefaultGenesis(c codec.JSONCodec) json.RawMessage {
	return nil
	//ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

// Validation check of the Genesis
func (AppModuleBasic) ValidateGenesis(c codec.JSONCodec, t client.TxEncodingConfig, j json.RawMessage) error {
	var data types.GenesisState
	// err := ModuleCdc.UnmarshalJSON(bz, &data)
	// if err != nil {
	// 	return err
	// }
	// Once json successfully marshalled, passes along to genesis.go
	return ValidateGenesis(&data)
}

// Register rest routes
func (AppModuleBasic) RegisterRESTRoutes(ctx client.Context, rtr *mux.Router) {
}

// Get the root query command of this module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
	//cli.GetQueryCmd(StoreKey, cdc)
}

// Get the root tx command of this module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
	//cli.GetTxCmd(StoreKey, cdc)
}

type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule Object
func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

func (AppModule) ConsensusVersion() uint64 {
	return 0
}

func (AppModule) RegisterServices(module.Configurator) {

}

func (AppModule) Name() string {
	return types.ModuleName
}

func (AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return nil
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}
func (am AppModule) QuerierRoute() string {
	return types.ModuleName
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return keeper.NewQuerier(am.keeper)
}

func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
}

func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return EndBlocker(ctx, am.keeper)
}

func (am AppModule) InitGenesis(ctx sdk.Context, c codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	//var genesisState GenesisState
	//ModuleCdc.MustUnmarshalJSON(data, &genesisState)

	return nil
	//InitGenesis(ctx, am.keeper, genesisState)
}

func (am AppModule) ExportGenesis(ctx sdk.Context, c codec.JSONCodec) json.RawMessage {
	//gs := ExportGenesis(ctx, am.keeper)
	//return ModuleCdc.MustMarshalJSON(gs)
	return nil
}
