package app

import (
	"encoding/json"
	"os"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/relayer"
	"github.com/celer-network/sgn-v2/x/sync"
	synckeeper "github.com/celer-network/sgn-v2/x/sync/keeper"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/celer-network/sgn-v2/x/validator"
	valkeeper "github.com/celer-network/sgn-v2/x/validator/keeper"
	valtypes "github.com/celer-network/sgn-v2/x/validator/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	tmcfg "github.com/tendermint/tendermint/config"
	tlog "github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	// unnamed import of statik for swagger UI support
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)

const appName = "sgn"

var (
	// default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.sgncli")

	// DefaultNodeHome sets the folder where the application data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.sgnd")

	// ModuleBasics defines the module BasicManager that is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		params.AppModuleBasic{},
		upgrade.AppModuleBasic{},

		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
		sync.AppModule{},
		validator.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	}

	LegacyAmino       = codec.NewLegacyAmino()
	InterfaceRegistry = types.NewInterfaceRegistry()
	AppCodec          = codec.NewProtoCodec(InterfaceRegistry)
	TxConfig          = tx.NewTxConfig(AppCodec, tx.DefaultSignModes)
)

type SgnApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	txConfig          client.TxConfig

	// keys to access the substores
	tkeyParams   *sdk.TransientStoreKey
	keyAccount   *sdk.KVStoreKey
	keyBank      *sdk.KVStoreKey
	keyStaking   *sdk.KVStoreKey
	keyParams    *sdk.KVStoreKey
	keyUpgrade   *sdk.KVStoreKey
	keySync      *sdk.KVStoreKey
	keyValidator *sdk.KVStoreKey

	// keepers
	accountKeeper   authkeeper.AccountKeeper
	bankKeeper      bankkeeper.Keeper
	stakingKeeper   stakingkeeper.Keeper
	paramsKeeper    paramskeeper.Keeper
	upgradeKeeper   upgradekeeper.Keeper
	syncKeeper      synckeeper.Keeper
	validatorKeeper valkeeper.Keeper

	// the module manager
	mm *module.Manager
}

// NewSgnApp is a constructor function for sgnApp
func NewSgnApp(
	logger tlog.Logger,
	db dbm.DB,
	height int64,
	skipUpgradeHeights map[int64]bool,
	tmCfg *tmcfg.Config,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SgnApp {
	viper.SetDefault(common.FlagEthPollInterval, 15)
	viper.SetDefault(common.FlagEthBlockDelay, 5)
	viper.SetDefault(common.FlagSgnCheckIntervalSlashQueue, 60)

	err := common.SetupUserPassword()
	if err != nil {
		tmos.Exit(err.Error())
	}

	loglevel := viper.GetString(common.FlagLogLevel)
	log.SetLevelByName(loglevel)
	if loglevel == "trace" {
		baseAppOptions = append(baseAppOptions, baseapp.SetTrace(true))
	}
	if viper.GetBool(common.FlagLogColor) {
		log.EnableColor()
	}

	std.RegisterLegacyAminoCodec(LegacyAmino)
	std.RegisterInterfaces(InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(LegacyAmino)
	ModuleBasics.RegisterInterfaces(InterfaceRegistry)

	bApp := baseapp.NewBaseApp(appName, logger, db, TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetInterfaceRegistry(InterfaceRegistry)

	app := &SgnApp{
		BaseApp:           bApp,
		legacyAmino:       LegacyAmino,
		appCodec:          AppCodec,
		interfaceRegistry: InterfaceRegistry,
		txConfig:          TxConfig,

		tkeyParams:   sdk.NewTransientStoreKey(paramstypes.TStoreKey),
		keyAccount:   sdk.NewKVStoreKey(authtypes.StoreKey),
		keyBank:      sdk.NewKVStoreKey(banktypes.StoreKey),
		keyStaking:   sdk.NewKVStoreKey(stakingtypes.StoreKey),
		keyParams:    sdk.NewKVStoreKey(paramstypes.StoreKey),
		keyUpgrade:   sdk.NewKVStoreKey(upgradetypes.StoreKey),
		keySync:      sdk.NewKVStoreKey(synctypes.StoreKey),
		keyValidator: sdk.NewKVStoreKey(valtypes.StoreKey),
	}

	app.paramsKeeper = paramskeeper.NewKeeper(AppCodec, LegacyAmino, app.keyParams, app.tkeyParams)
	// Set specific subspaces
	authSubspace := app.paramsKeeper.Subspace(authtypes.ModuleName)
	bankSupspace := app.paramsKeeper.Subspace(banktypes.ModuleName)
	stakingSubspace := app.paramsKeeper.Subspace(stakingtypes.ModuleName)
	syncSubspace := app.paramsKeeper.Subspace(synctypes.ModuleName)
	validatorSubspace := app.paramsKeeper.Subspace(valtypes.ModuleName)

	bApp.SetParamStore(app.paramsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = authkeeper.NewAccountKeeper(
		AppCodec, app.keyAccount, authSubspace, authtypes.ProtoBaseAccount, maccPerms,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bankkeeper.NewBaseKeeper(
		AppCodec, app.keyBank, app.accountKeeper, bankSupspace, app.ModuleAccountAddrs(),
	)

	// The staking keeper
	stakingKeeper := stakingkeeper.NewKeeper(
		AppCodec, app.keyStaking, app.accountKeeper, app.bankKeeper, stakingSubspace,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(),
	)

	app.validatorKeeper = valkeeper.NewKeeper(
		AppCodec, app.keyValidator, app.accountKeeper, app.stakingKeeper, validatorSubspace,
	)

	app.upgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, app.keyUpgrade, AppCodec, DefaultNodeHome, app.BaseApp)

	app.syncKeeper = synckeeper.NewKeeper(
		AppCodec, app.keySync, app.validatorKeeper, syncSubspace,
	)

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx,
			TxConfig,
		),
		auth.NewAppModule(AppCodec, app.accountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(AppCodec, app.bankKeeper, app.accountKeeper),
		staking.NewAppModule(AppCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		validator.NewAppModule(app.validatorKeeper),
		sync.NewAppModule(app.syncKeeper),
	)

	app.mm.SetOrderBeginBlockers(upgradetypes.ModuleName)
	app.mm.SetOrderEndBlockers(valtypes.ModuleName, synctypes.ModuleName)

	app.mm.SetOrderInitGenesis(
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		genutiltypes.ModuleName,
		valtypes.ModuleName,
		synctypes.ModuleName,
	)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), LegacyAmino)

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing
	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.accountKeeper,
			BankKeeper:      app.bankKeeper,
			SignModeHandler: TxConfig.SignModeHandler(),
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)

	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	app.MountStores(
		app.tkeyParams,
		app.keyAccount,
		app.keyStaking,
		app.keyParams,
		app.keyUpgrade,
		app.keyValidator,
		app.keySync,
	)

	if height == -1 {
		err = app.LoadLatestVersion()
	} else {
		err = app.LoadHeight(height)
	}
	if err != nil {
		tmos.Exit("Failed to load height:" + err.Error())
	}

	go app.startRelayer(db, tmCfg)

	return app
}

// Name returns the name of the App
func (app *SgnApp) Name() string { return app.BaseApp.Name() }

type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	return ModuleBasics.DefaultGenesis(cdc)
}

// BeginBlocker application updates every begin block
func (app *SgnApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *SgnApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *SgnApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.upgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *SgnApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *SgnApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (app *SgnApp) startRelayer(db dbm.DB, tmCfg *tmcfg.Config) {
	operator, err := relayer.NewOperator(app.appCodec, viper.GetString(common.FlagCLIHome), tmCfg, app.legacyAmino)
	if err != nil {
		tmos.Exit(err.Error())
	}

	_, err = rpc.GetChainHeight(operator.Transactor.CliCtx)
	for err != nil {
		time.Sleep(time.Second)
		_, err = rpc.GetChainHeight(operator.Transactor.CliCtx)
	}

	relayer.NewRelayer(operator, db)
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SgnApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	tx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *SgnApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *SgnApp) RegisterTxService(clientCtx client.Context) {
	tx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// ExportAppStateAndValidators exports the state of the application for a genesis
// file.
func (app *SgnApp) ExportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string,
) (servertypes.ExportedApp, error) {
	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true, tmproto.Header{Height: app.LastBlockHeight()})

	// We export at last height + 1, because that's the height at which
	// Tendermint will start InitChain.
	height := app.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0
	}

	genState := app.mm.ExportGenesis(ctx, app.appCodec)
	appState, err := json.MarshalIndent(genState, "", "  ")
	if err != nil {
		return servertypes.ExportedApp{}, err
	}

	validators, err := staking.WriteValidators(ctx, app.stakingKeeper)
	return servertypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: app.BaseApp.GetConsensusParams(ctx),
	}, err
}
