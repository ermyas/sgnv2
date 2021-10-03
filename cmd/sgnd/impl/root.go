package impl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	tmcfg "github.com/tendermint/tendermint/config"

	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/app/params"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/ops"
)

// NewRootCmd creates a new root command for sgnd. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := app.MakeEncodingConfig()

	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
	sdkConfig.SetAddressVerifier(common.VerifyAddressFormat)
	sdkConfig.Seal()

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithViper("SGN")

	rootCmd := &cobra.Command{
		Use:   "sgnd",
		Short: "SGN App",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// Set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx = client.ReadHomeFlag(initClientCtx, cmd)
			initClientCtx, err := config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			serverCtx := server.GetServerContextFromCmd(cmd)

			sgnAppTemplate, sgnAppConfig := initAppConfig()

			err = server.InterceptConfigsPreRunHandler(cmd, sgnAppTemplate, sgnAppConfig)
			if err != nil {
				return err
			}

			// Merge in SGN-specific configs, must be run after InterceptConfigsPreRunHandler
			// TODO: Maybe use proper AppConfig and avoid global viper
			rootViper := serverCtx.Viper
			rootDir := rootViper.GetString(flags.FlagHome)
			cfgFile := filepath.Join(rootDir, "config", "config.toml")
			viper.SetConfigFile(cfgFile)
			if err := viper.ReadInConfig(); err != nil {
				return fmt.Errorf("failed to read in tendermint configuration: %w", err)
			}
			cbrCfgFile := filepath.Join(rootDir, "config", "cbridge.toml")
			viper.SetConfigFile(cbrCfgFile)
			if err := viper.MergeInConfig(); err != nil {
				return fmt.Errorf("failed to read in cbridge configuration: %w", err)
			}
			sgnCfgFile := filepath.Join(rootDir, "config", "sgn.toml")
			viper.SetConfigFile(sgnCfgFile)
			if err := viper.MergeInConfig(); err != nil {
				return fmt.Errorf("failed to read in SGN configuration: %w", err)
			}

			// TODO: Check if necessary
			// Reset logger TimeFormat
			var logWriter io.Writer
			if strings.ToLower(serverCtx.Viper.GetString(flags.FlagLogFormat)) == tmcfg.LogFormatPlain {
				logWriter = zerolog.ConsoleWriter{
					Out:         os.Stdout,
					TimeFormat:  "2006-01-02 15:04:05 UTC",
					NoColor:     !viper.GetBool(common.FlagLogColor),
					FormatLevel: logFormatLevel(viper.GetBool(common.FlagLogColor)),
				}
			} else {
				logWriter = os.Stderr
			}
			logLvlStr := serverCtx.Viper.GetString(flags.FlagLogLevel)
			logLvl, err := zerolog.ParseLevel(logLvlStr)
			if err != nil {
				return fmt.Errorf("failed to parse log level (%s): %w", logLvlStr, err)
			}
			zerolog.TimestampFunc = func() time.Time { return time.Now().UTC().Round(time.Second) }
			serverCtx.Logger = ZeroLogWrapper{zerolog.New(logWriter).Level(logLvl).With().Timestamp().Logger()}
			server.SetCmdServerContext(cmd, serverCtx)

			return nil
		},
	}

	initRootCmd(rootCmd, encodingConfig)

	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(app.DefaultNodeHome),
		AddGenesisValidatorCmd(app.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		// TODO: Add cmd to generate a testnet?
		debug.Cmd(),
	)

	a := appCreator{rootCmd, encodingConfig}
	server.AddCommands(rootCmd, app.DefaultNodeHome, a.newApp, a.appExport, addModuleInitFlags)

	// Add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		ops.OpsCommand(),
		keys.Commands(app.DefaultNodeHome),
	)

	// Add rosetta commands
	rootCmd.AddCommand(server.RosettaCommand(encodingConfig.InterfaceRegistry, encodingConfig.Codec))
}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		flags.LineBreak,
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

type appCreator struct {
	rootCmd        *cobra.Command // To pass along serverCtx
	encodingConfig params.EncodingConfig
}

// newApp is an AppCreator
func (a appCreator) newApp(logger log.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := sdk.NewLevelDB("metadata", snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}
	serverCtx := server.GetServerContextFromCmd(a.rootCmd)

	return app.NewSgnApp(
		logger, db, traceStore, true /* loadLatest */, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		a.encodingConfig,
		appOpts,
		serverCtx.Config,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshotStore(snapshotStore),
		baseapp.SetSnapshotInterval(cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval))),
		baseapp.SetSnapshotKeepRecent(cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent))),
	)
}

func (a appCreator) appExport(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailAllowedAddrs []string,
	appOpts servertypes.AppOptions) (servertypes.ExportedApp, error) {
	serverCtx := server.GetServerContextFromCmd(a.rootCmd)

	var sgnApp *app.SgnApp
	if height != -1 {
		sgnApp =
			app.NewSgnApp(
				logger, db, traceStore,
				false,            /* loadLatest */
				map[int64]bool{}, /* skipUpgradeHeights */
				"",               /* homePath */
				cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
				a.encodingConfig,
				appOpts,
				serverCtx.Config,
			)

		if err := sgnApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		sgnApp =
			app.NewSgnApp(
				logger, db, traceStore,
				true,             /* loadLatest */
				map[int64]bool{}, /* skipUpgradeHeights */
				"",               /* homePath */
				cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
				a.encodingConfig,
				appOpts,
				serverCtx.Config,
			)
	}

	return sgnApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
