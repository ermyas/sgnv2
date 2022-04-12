/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/celer-network/sgn-v2/executor/sgn"
	"github.com/celer-network/sgn-v2/executor/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// queryCmd represents the query command
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Short:                      "Querying command to sgn and gateway.",
		SuggestionsMinimumDistance: 2,
	}
	cmd.AddCommand(GetCmdQueryCts())
	return cmd
}

func GetCmdQueryCts() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "start",
		Short:              "Start executor service",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			test, err := cmd.Flags().GetBool("test")
			if err != nil {
				return err
			}
			sgn := sgn.NewSgnClient(viper.GetString(types.FlagSgnGrpcUrl), test)
			_, _ = sgn.GetExecutionContexts(nil)
			return nil
		},
	}
	return cmd
}

func init() {
	// rootCmd.AddCommand(GetQueryCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
