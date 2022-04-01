package ops

import (
	"fmt"

	"github.com/celer-network/sgn-v2/eth"
	msgcli "github.com/celer-network/sgn-v2/x/message/client/cli"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func QueryMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-message",
		Short: "query message by chain id and tx hash",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			chainid := viper.GetUint64(FlagChainId)
			txhash := viper.GetString(FlagTxHash)

			cbr, txReceipt, err := setupCbr(chainid, txhash)
			if err != nil {
				return err
			}

			elog := eth.FindMatchContractEvent(
				eth.ContractTypeMsgBus, "Message", cbr.msgContract.Address, txReceipt.Logs)
			if elog == nil {
				return fmt.Errorf("no match event found in tx")
			}
			ev, err := cbr.msgContract.ParseMessage(*elog)
			if err != nil {
				return fmt.Errorf("ParseMessage err: %w", err)
			}
			messageId, _ := msgtypes.NewMessage(ev, chainid)
			fmt.Printf("msgId: %x, ev: %s", messageId, ev.PrettyLog(chainid))

			message, err := msgcli.QueryMessage(cliCtx, messageId.Hex())
			// todo: prettier print output
			return cliCtx.PrintProto(&message)
		},
	}
	cmd.Flags().Uint64(FlagChainId, 0, "which chainid to query tx hash")
	cmd.Flags().String(FlagTxHash, "", "tx hash, will parse message events")
	cmd.MarkFlagRequired(FlagChainId)
	cmd.MarkFlagRequired(FlagTxHash)
	return cmd
}
