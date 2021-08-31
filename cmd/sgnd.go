package cmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/app/params"
	sgndimpl "github.com/celer-network/sgn-v2/cmd/sgnd/impl"
	"github.com/celer-network/sgn-v2/common"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			if err != nil {
				return err
			}

			// TODO: Use customAppConfig
			sgnConfigPath := viper.GetString(common.FlagConfig)
			_, err = os.Stat(sgnConfigPath)
			if err != nil {
				return err
			}
			viper.SetConfigFile(sgnConfigPath)
			err = viper.ReadInConfig()
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
