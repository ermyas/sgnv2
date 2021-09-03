package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/app/params"
	sgndimpl "github.com/celer-network/sgn-v2/cmd/sgnd/impl"
	"github.com/celer-network/sgn-v2/common"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/rs/zerolog"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmcfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tlog "github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func GetSgndExecutor(encodingConfig params.EncodingConfig) cli.Executor {
	rootCmd := &cobra.Command{
		Use:   "sgnd",
		Short: "SGN App",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// Set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx := client.Context{}.
				WithCodec(encodingConfig.Codec).
				WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
				WithTxConfig(encodingConfig.TxConfig).
				WithLegacyAmino(encodingConfig.Amino).
				WithInput(os.Stdin).
				WithAccountRetriever(types.AccountRetriever{}).
				WithHomeDir(app.DefaultNodeHome).
				WithViper("")

			initClientCtx = client.ReadHomeFlag(initClientCtx, cmd)

			initClientCtx, err := config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}
			err = server.InterceptConfigsPreRunHandler(cmd, "", nil)

			// TODO: Use customAppConfig
			sgnConfigPath := viper.GetString(common.FlagConfig)
			_, err = os.Stat(sgnConfigPath)
			if err != nil {
				return err
			}
			viper.SetConfigFile(sgnConfigPath)
			err = viper.ReadInConfig()

			// reset logger TimeFormat
			serverCtx := server.GetServerContextFromCmd(cmd)
			var logWriter io.Writer
			if strings.ToLower(serverCtx.Viper.GetString(flags.FlagLogFormat)) == tmcfg.LogFormatPlain {
				logWriter = zerolog.ConsoleWriter{
					Out:         os.Stderr,
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

			return err
		},
	}
	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		sgndimpl.AddGenesisAccountCmd(app.DefaultNodeHome),
	)

	a := appCreator{rootCmd, encodingConfig}

	server.AddCommands(rootCmd, app.DefaultNodeHome, a.newApp, a.exportAppStateAndTMValidators, addModuleInitFlags)
	rootCmd.PersistentFlags().String(common.FlagCLIHome, app.DefaultCLIHome, "Directory for cli config and data")
	rootCmd.PersistentFlags().String(
		common.FlagConfig, filepath.Join(app.DefaultCLIHome, "config", "sgn.toml"), "Path to SGN-specific configs")

	// prepare and add flags
	return cli.PrepareBaseCmd(rootCmd, "SGN", app.DefaultNodeHome)
}

type appCreator struct {
	rootCmd *cobra.Command
	encCfg  params.EncodingConfig
}

func (a appCreator) newApp(logger tlog.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
	var cache sdk.MultiStorePersistentCache
	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}
	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range viper.GetIntSlice(server.FlagUnsafeSkipUpgrades) {
		skipUpgradeHeights[int64(h)] = true
	}
	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	serverCtx := server.GetServerContextFromCmd(a.rootCmd)
	return app.NewSgnApp(
		logger,
		db,
		true, /* loadLatest */
		skipUpgradeHeights,
		serverCtx.Config,
		a.encCfg,
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetPruning(pruningOpts),
	)
}

func (a appCreator) exportAppStateAndTMValidators(
	logger tlog.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
	appOpts servertypes.AppOptions) (servertypes.ExportedApp, error) {
	var sgnApp *app.SgnApp
	if height != -1 {
		sgnApp = app.NewSgnApp(logger, db, false, map[int64]bool{}, nil, a.encCfg)
		if err := sgnApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		sgnApp = app.NewSgnApp(logger, db, true, map[int64]bool{}, nil, a.encCfg)
	}
	return sgnApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}
