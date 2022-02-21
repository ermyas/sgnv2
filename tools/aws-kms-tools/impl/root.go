package impl

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "aws-kms-tools",
		Short: "Tools to perform operations with an AWS KMS key",
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

	rootCmd.MarkFlagRequired(FlagRegion)
	rootCmd.MarkFlagRequired(FlagAlias)

	viper.BindPFlag(FlagRegion, rootCmd.PersistentFlags().Lookup(FlagRegion))
	viper.BindPFlag(FlagAlias, rootCmd.PersistentFlags().Lookup(FlagAlias))
	viper.BindPFlag(FlagAwsKey, rootCmd.PersistentFlags().Lookup(FlagAwsKey))
	viper.BindPFlag(FlagAwsSec, rootCmd.PersistentFlags().Lookup(FlagAwsSec))

	return rootCmd
}
