package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(common.PostCommands(
		GetCmdSignAgainMint(),
		GetCmdSignAgainWithdraw(),
	)...)

	return cmd
}

func InitClaimFee(t *transactor.Transactor, req *types.MsgClaimFee) (resp *types.MsgClaimFeeResponse, err error) {
	req.Sender = t.Key.GetAddress().String() // make sure the msg sender is the transactor
	_, err = t.LockSendTx(req)
	return
}

func GetCmdSignAgainMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-again-mint",
		Short: "Trigger sign again of mint msg",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			home, err := cmd.Flags().GetString(flags.FlagHome)
			if err != nil {
				return err
			}
			txr, err := transactor.NewCliTransactor(home, clientCtx.LegacyAmino, clientCtx.Codec, clientCtx.InterfaceRegistry)
			if err != nil {
				log.Error(err)
				return err
			}
			msg := &types.MsgTriggerSignMint{
				MintId: args[0],
				Sender: txr.Key.GetAddress().String(),
			}
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			txr.CliSendTxMsgWaitMined(msg)

			return nil
		},
	}
	return cmd
}

func GetCmdSignAgainWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-again-withdraw",
		Short: "Trigger sign again of withdraw msg",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			home, err := cmd.Flags().GetString(flags.FlagHome)
			if err != nil {
				return err
			}
			txr, err := transactor.NewCliTransactor(home, clientCtx.LegacyAmino, clientCtx.Codec, clientCtx.InterfaceRegistry)
			if err != nil {
				log.Error(err)
				return err
			}
			msg := &types.MsgTriggerSignWithdraw{
				WithdrawId: args[0],
				Sender:     txr.Key.GetAddress().String(),
			}
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			txr.CliSendTxMsgWaitMined(msg)

			return nil
		},
	}
	return cmd
}
