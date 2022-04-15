/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/executor"
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
	cmd.AddCommand(GetCmdQueryMsg())
	return cmd
}

func GetCmdQueryMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "message [message-id]",
		Short: "Query message status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			test, err := cmd.Flags().GetBool("test")
			if err != nil {
				return err
			}
			messageId := args[0]
			dal := executor.NewDAL()
			status, err := QueryLocalMessage(dal, eth.Hex2Bytes(messageId))
			if err == nil {
				cmd.Printf("Status of message %s is %s.\n", messageId, status.String())
				return nil
			}
			log.Infof("No message found in local db with id:%s, would query sgn instead.", messageId)
			sgn := sgn.NewSgnClient(viper.GetString(types.FlagSgnGrpcUrl), test)
			status, err = QuerySgnMessage(sgn, messageId)
			if err != nil {
				log.Errorf("QuerySgnMessage err:%v", err)
				return err
			}
			cmd.Printf("Status of message %s is %s.\n", messageId, status.String())
			return nil
		},
	}
	return cmd
}

func QueryLocalMessage(dal *executor.DAL, id []byte) (types.ExecutionStatus, error) {
	request := dal.GetExecuteContext(id)
	if request != nil {
		return request.SS, nil
	}
	return types.ExecutionStatus_Unknown, fmt.Errorf("no execute context found")
}

func QuerySgnMessage(sgn *sgn.SgnClient, id string) (types.ExecutionStatus, error) {
	message, err := sgn.GetMessage(id)
	if err != nil {
		return types.ExecutionStatus_Unknown, err
	}
	status, _ := types.NewExecutionStatus(uint8(message.ExecutionStatus))
	return status, nil
}

func init() {
	rootCmd.AddCommand(GetQueryCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
