package app

import (
	"encoding/json"
	"os"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/monitor"
	"github.com/celer-network/sgn-v2/x/validator"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	tlog "github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
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
		validator.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	}
)

type SgnApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	// keys to access the substores
	tkeyParams   *sdk.TransientStoreKey
	keyAccount   *sdk.KVStoreKey
	keyBank      *sdk.KVStoreKey
	keyStaking   *sdk.KVStoreKey
	keyParams    *sdk.KVStoreKey
	keyUpgrade   *sdk.KVStoreKey
	keyValidator *sdk.KVStoreKey

	// keepers
	accountKeeper   authkeeper.AccountKeeper
	bankKeeper      bankkeeper.Keeper
	stakingKeeper   stakingkeeper.Keeper
	paramsKeeper    paramskeeper.Keeper
	upgradeKeeper   upgradekeeper.Keeper
	validatorKeeper validator.Keeper

	// the module manager
	mm *module.Manager
}

// NewSgnApp is a constructor function for sgnApp
func NewSgnApp(
	logger tlog.Logger, db dbm.DB, skipUpgradeHeights map[int64]bool, homePath string,
	encodingConfig simappparams.EncodingConfig, baseAppOptions ...func(*baseapp.BaseApp),
) *SgnApp {
	viper.SetDefault(common.FlagEthPollInterval, 15)
	viper.SetDefault(common.FlagEthBlockDelay, 5)
	viper.SetDefault(common.FlagSgnCheckIntervalGuardQueue, 60)
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

	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	app := &SgnApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,

		tkeyParams:   sdk.NewTransientStoreKey(paramstypes.TStoreKey),
		keyAccount:   sdk.NewKVStoreKey(authtypes.StoreKey),
		keyBank:      sdk.NewKVStoreKey(banktypes.StoreKey),
		keyStaking:   sdk.NewKVStoreKey(stakingtypes.StoreKey),
		keyParams:    sdk.NewKVStoreKey(paramstypes.StoreKey),
		keyUpgrade:   sdk.NewKVStoreKey(upgradetypes.StoreKey),
		keyValidator: sdk.NewKVStoreKey(validator.StoreKey),
	}

	app.paramsKeeper = paramskeeper.NewKeeper(appCodec, legacyAmino, app.keyParams, app.tkeyParams)
	// Set specific subspaces
	authSubspace := app.paramsKeeper.Subspace(authtypes.ModuleName)
	bankSupspace := app.paramsKeeper.Subspace(banktypes.ModuleName)
	stakingSubspace := app.paramsKeeper.Subspace(stakingtypes.ModuleName)
	validatorSubspace := app.paramsKeeper.Subspace(validator.ModuleName)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = authkeeper.NewAccountKeeper(
		appCodec, app.keyAccount, authSubspace, authtypes.ProtoBaseAccount, maccPerms,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, app.keyBank, app.accountKeeper, bankSupspace, app.ModuleAccountAddrs(),
	)

	// The staking keeper
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, app.keyStaking, app.accountKeeper, app.bankKeeper, stakingSubspace,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(),
	)

	app.validatorKeeper = validator.NewKeeper(
		appCodec, app.keyValidator, app.accountKeeper, app.stakingKeeper, validatorSubspace,
	)

	app.upgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, app.keyUpgrade, appCodec, homePath, app.BaseApp)

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.accountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper),
		staking.NewAppModule(appCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		validator.NewAppModule(app.validatorKeeper),
	)

	app.mm.SetOrderBeginBlockers(upgradetypes.ModuleName)
	app.mm.SetOrderEndBlockers(validator.ModuleName)

	app.mm.SetOrderInitGenesis(
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		genutiltypes.ModuleName,
		validator.ModuleName,
	)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing
	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.accountKeeper,
			BankKeeper:      app.bankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
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
	)

	err = app.LoadLatestVersion()
	if err != nil {
		tmos.Exit(err.Error())
	}

	go app.startMonitor(db)

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

func (app *SgnApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

func (app *SgnApp) AppCodec() codec.Codec {
	return app.appCodec
}

func (app *SgnApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

func (app *SgnApp) startMonitor(db dbm.DB) {
	operator, err := monitor.NewOperator(app.appCodec, viper.GetString(common.FlagCLIHome))
	if err != nil {
		tmos.Exit(err.Error())
	}

	_, err = rpc.GetChainHeight(operator.Transactor.CliCtx)
	for err != nil {
		time.Sleep(time.Second)
		_, err = rpc.GetChainHeight(operator.Transactor.CliCtx)
	}

	monitor.NewMonitor(operator, db)
}
