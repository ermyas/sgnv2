package cli

import (
    "strconv"
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

var _ = strconv.Itoa(0)

func CmdOnchainEvent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "onchain-event [chainid]",
		Short: "Broadcast message onchain-event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsChainid, _ := strconv.ParseUint(args[0], 10, 64)
      
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOnchainEvent(clientCtx.GetFromAddress().String(), uint64(argsChainid))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}