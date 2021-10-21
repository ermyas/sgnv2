package app

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/celer-network/goutils/log"
	appparams "github.com/celer-network/sgn-v2/app/params"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway"
	"github.com/celer-network/sgn-v2/relayer"
	"github.com/celer-network/sgn-v2/x/cbridge"
	cbrclient "github.com/celer-network/sgn-v2/x/cbridge/client"
	cbridgekeeper "github.com/celer-network/sgn-v2/x/cbridge/keeper"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	distr "github.com/celer-network/sgn-v2/x/distribution"
	distrkeeper "github.com/celer-network/sgn-v2/x/distribution/keeper"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/celer-network/sgn-v2/x/farming"
	farmingclient "github.com/celer-network/sgn-v2/x/farming/client"
	farmingkeeper "github.com/celer-network/sgn-v2/x/farming/keeper"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/celer-network/sgn-v2/x/gov"
	govclient "github.com/celer-network/sgn-v2/x/gov/client"
	govkeeper "github.com/celer-network/sgn-v2/x/gov/keeper"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/celer-network/sgn-v2/x/mint"
	mintkeeper "github.com/celer-network/sgn-v2/x/mint/keeper"
	minttypes "github.com/celer-network/sgn-v2/x/mint/types"
	"github.com/celer-network/sgn-v2/x/slash"
	slashkeeper "github.com/celer-network/sgn-v2/x/slash/keeper"
	slashtypes "github.com/celer-network/sgn-v2/x/slash/types"
	staking "github.com/celer-network/sgn-v2/x/staking"
	stakingkeeper "github.com/celer-network/sgn-v2/x/staking/keeper"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/celer-network/sgn-v2/x/sync"
	synckeeper "github.com/celer-network/sgn-v2/x/sync/keeper"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
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
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	tmcfg "github.com/tendermint/tendermint/config"
	tlog "github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	// unnamed import of statik for swagger UI support
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)

const appName = "sgn"

var (
	// DefaultNodeHome sets the folder where the application data and configuration will be stored
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager that is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		upgrade.AppModuleBasic{},

		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		farming.AppModuleBasic{},
		gov.NewAppModuleBasic(
			govclient.ParamProposalHandler,
			govclient.UpgradeProposalHandler,
			cbrclient.CbrConfigProposalHandler,
			farmingclient.AddPoolProposalHandler,
			farmingclient.AddTokensProposalHandler,
			farmingclient.AdjustRewardProposalHandler,
		),
		slash.AppModule{},
		sync.AppModule{},
		staking.AppModuleBasic{},
		cbridge.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		minttypes.ModuleName:       {authtypes.Minter},
		authtypes.FeeCollectorName: nil,
		distrtypes.ModuleName:      nil,
		// Needed for mint-then-stake & unstake-then-burn
		farming.ModuleName: {authtypes.Minter, authtypes.Burner},
		// Needed for rewards
		farming.RewardModuleAccountName: {authtypes.Minter},
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		farming.ModuleName:              true,
		farming.RewardModuleAccountName: true,
	}
)

// Verify app interface at compile time
var (
	_ simapp.App              = (*SgnApp)(nil)
	_ servertypes.Application = (*SgnApp)(nil)
)

type SgnApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tKeys map[string]*sdk.TransientStoreKey

	// keepers
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	UpgradeKeeper upgradekeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper
	MintKeeper    mintkeeper.Keeper
	DistrKeeper   distrkeeper.Keeper
	FarmingKeeper farmingkeeper.Keeper
	GovKeeper     govkeeper.Keeper
	SlashKeeper   slashkeeper.Keeper
	SyncKeeper    synckeeper.Keeper
	StakingKeeper stakingkeeper.Keeper
	CbridgeKeeper cbridgekeeper.Keeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// the configurator
	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Failed to get home dir %w", err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".sgnd")
}

// NewSgnApp is a constructor function for sgnApp
func NewSgnApp(
	logger tlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig appparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	tmCfg *tmcfg.Config,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SgnApp {
	// TODO: Check if these can be set by config template and remove.
	viper.SetDefault(common.FlagEthPollInterval, 15)
	viper.SetDefault(common.FlagEthBlockDelay, 5)
	viper.SetDefault(common.FlagSgnCheckIntervalSlash, 60)
	viper.SetDefault(common.FlagSgnCheckIntervalCbridge, 15)
	viper.SetDefault(common.FlagSgnCheckIntervalVerifier, 15)

	err := common.SetupUserPassword()
	if err != nil {
		tmos.Exit(err.Error())
	}

	// Celer goutils log configs
	loglevel := viper.GetString(common.FlagLogLevel)
	log.SetLevelByName(loglevel)
	if loglevel == "trace" {
		baseAppOptions = append(baseAppOptions, baseapp.SetTrace(true))
	}
	if viper.GetBool(common.FlagLogColor) {
		log.EnableColor()
	}

	appCodec := encodingConfig.Codec
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey,
		paramstypes.StoreKey, upgradetypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, farmingtypes.StoreKey,
		govtypes.StoreKey, slashtypes.StoreKey, synctypes.StoreKey, stakingtypes.StoreKey,
		cbrtypes.MemStoreKey, cbrtypes.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)

	app := &SgnApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tKeys:             tKeys,
	}

	// Init params keeper and subspaces
	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey])
	// Set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// Add Cosmos SDK keepers
	// The AccountKeeper handles address -> account lookups
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	// The BankKeeper allows you to perform sdk.Coins interactions
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.BlockedAddrs(),
	)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, DefaultNodeHome, app.BaseApp)

	// Initialize SGN-specific keepers
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.GetSubspace(stakingtypes.ModuleName),
	)
	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey], app.GetSubspace(minttypes.ModuleName),
		app.AccountKeeper, app.BankKeeper, authtypes.FeeCollectorName,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName, app.BlockedAddrs(),
	)
	app.FarmingKeeper = farmingkeeper.NewKeeper(
		appCodec, keys[farmingtypes.StoreKey], app.GetSubspace(farmingtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper,
	)
	app.CbridgeKeeper = cbridgekeeper.NewKeeper(
		appCodec,
		keys[cbrtypes.StoreKey],
		app.GetSubspace(cbrtypes.ModuleName),
		&stakingKeeper,
		app.FarmingKeeper,
	)
	app.SyncKeeper = synckeeper.NewKeeper(
		appCodec, keys[synctypes.StoreKey], &stakingKeeper, app.GetSubspace(synctypes.ModuleName), app.CbridgeKeeper,
	)

	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(proposal.RouterKey, gov.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, gov.NewUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(cbrtypes.RouterKey, cbridge.NewCbrProposalHandler(app.CbridgeKeeper)).
		AddRoute(farmingtypes.RouterKey, farming.NewAddPoolProposalHandler(app.FarmingKeeper))
	app.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		stakingKeeper,
		govRouter,
	)
	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.CbridgeKeeper.Hooks()),
	)
	app.SlashKeeper = slashkeeper.NewKeeper(
		keys[slashtypes.StoreKey],
		appCodec,
		app.StakingKeeper,
		app.GetSubspace(slashtypes.ModuleName),
	)

	/****  Module Options ****/

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		params.NewAppModule(app.ParamsKeeper),

		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		farming.NewAppModule(appCodec, app.FarmingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper),
		gov.NewAppModule(app.GovKeeper, app.AccountKeeper),
		slash.NewAppModule(app.SlashKeeper),
		sync.NewAppModule(app.SyncKeeper),
		cbridge.NewAppModule(appCodec, app.CbridgeKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		stakingtypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashtypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		synctypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	// NOTE: Treasury must occur after bank module so that initial supply is properly set
	app.mm.SetOrderInitGenesis(
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		farmingtypes.ModuleName,
		minttypes.ModuleName,
		stakingtypes.ModuleName,
		govtypes.ModuleName,
		slashtypes.ModuleName,
		synctypes.ModuleName,
		cbrtypes.ModuleName,
	)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), legacyAmino)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// Create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		// TODO - uncomment when v0.43.0 fix the simulation bug
		// authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		params.NewAppModule(app.ParamsKeeper),
	)

	app.sm.RegisterStoreDecoders()

	// Initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)

	// Initialize BaseApp
	// The InitChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing
	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)

	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	// Piggy-back starting the relayer
	go app.startRelayer(db, tmCfg, homePath)

	if viper.GetBool(common.FlagToStartGateway) {
		log.Infoln("starting gateway...")
		dbUrl := viper.GetString(common.FlagGatewayDbUrl)
		if dbUrl == "" {
			dbUrl = "127.0.0.1:26257"
		}
		go gateway.InitGateway(
			homePath,
			app.legacyAmino,
			app.appCodec,
			app.interfaceRegistry,
			false,
			dbUrl)
	}

	return app
}

// Name returns the name of the App
func (app *SgnApp) Name() string { return app.BaseApp.Name() }

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
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
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

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive external tokens.
func (app *SgnApp) BlockedAddrs() map[string]bool {
	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blockedAddrs
}

// LegacyAmino returns SgnApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SgnApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns SgnApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SgnApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns SgnApp's InterfaceRegistry
func (app *SgnApp) InterfaceRegistry() codectypes.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SgnApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SgnApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SgnApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *SgnApp) SimulationManager() *module.SimulationManager {
	return app.sm
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

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key sdk.StoreKey, tKey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tKey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)

	paramsKeeper.Subspace(distrtypes.ModuleName).WithKeyTable(distrtypes.ParamKeyTable())
	paramsKeeper.Subspace(farmingtypes.ModuleName).WithKeyTable(farmingtypes.ParamKeyTable())
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(slashtypes.ModuleName).WithKeyTable(slashtypes.ParamKeyTable())
	paramsKeeper.Subspace(stakingtypes.ModuleName).WithKeyTable(stakingtypes.ParamKeyTable())
	paramsKeeper.Subspace(synctypes.ModuleName).WithKeyTable(synctypes.ParamKeyTable())
	paramsKeeper.Subspace(cbrtypes.ModuleName).WithKeyTable(cbrtypes.ParamKeyTable())

	return paramsKeeper
}

func (app *SgnApp) startRelayer(db dbm.DB, tmCfg *tmcfg.Config, homeDir string) {
	operator, err := relayer.NewOperator(homeDir, tmCfg, app.legacyAmino, app.appCodec, app.interfaceRegistry)
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
