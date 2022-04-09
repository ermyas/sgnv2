package impl

import (
	"os"
	"path/filepath"

	"github.com/celer-network/goutils/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Failed to get home dir %w", err)
	}

	DefaultCfgHome = filepath.Join(userHomeDir, ".sgnd")
}

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "aws-kms-tools",
		Short: "Tools to perform operations with an AWS KMS key",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			rootDir := viper.GetString(FlagCfgHome)
			cbrCfgFile := filepath.Join(rootDir, "config", "cbridge.toml")
			viper.SetConfigFile(cbrCfgFile)
			if err := viper.ReadInConfig(); err != nil {
				log.Warnln("failed to read cbridge.toml:", err)
			}
			sgnCfgFile := filepath.Join(rootDir, "config", "sgn.toml")
			viper.SetConfigFile(sgnCfgFile)
			if err := viper.MergeInConfig(); err != nil {
				log.Warnln("failed to read sgn.toml:", err)
			}
			return nil
		},
	}

	rootCmd.AddCommand(
		PrintAddressCommand(),
		PrintNonceCommand(),
		SignMessageCommand(),
		SendTxCommand(),
	)

	rootCmd.PersistentFlags().StringP(FlagRegion, FlagRegionShort, "", "AWS region containing the KMS key")
	rootCmd.PersistentFlags().StringP(FlagAlias, FlagAliasShort, "", "KMS key alias")
	rootCmd.PersistentFlags().String(FlagAwsKey, "", "Optional AWS API key")
	rootCmd.PersistentFlags().String(FlagAwsSec, "", "Optional AWS API secret")
	rootCmd.PersistentFlags().String(FlagCfgHome, DefaultCfgHome, "Home of sgnd config")

	viper.BindPFlag(FlagRegion, rootCmd.PersistentFlags().Lookup(FlagRegion))
	viper.BindPFlag(FlagAlias, rootCmd.PersistentFlags().Lookup(FlagAlias))
	viper.BindPFlag(FlagAwsKey, rootCmd.PersistentFlags().Lookup(FlagAwsKey))
	viper.BindPFlag(FlagAwsSec, rootCmd.PersistentFlags().Lookup(FlagAwsSec))
	viper.BindPFlag(FlagCfgHome, rootCmd.PersistentFlags().Lookup(FlagCfgHome))

	return rootCmd
}
