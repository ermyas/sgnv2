package cli

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
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
		GetCmdUpdateLatestSigners(),
		GetCmdSignAgainRelay(),
	)...)

	return cmd
}

func GetCmdSignAgainRelay() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-again-relay",
		Short: "Trigger sign again of relay msg",
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
			msg := &types.MsgSignAgain{
				DataType: types.SignDataType_RELAY,
				XferId:   eth.Hex2Bytes(args[0]),
				Creator:  txr.Key.GetAddress().String(),
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

func GetCmdUpdateLatestSigners() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-latest-signers",
		Short: "Trigger update of latest signers",
		Args:  cobra.ExactArgs(0),
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
			msg := types.NewMsgUpdateLatestSigners(txr.Key.GetAddress().String())
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

// if err not nil, should return immediately when estimate gas
func InitWithdraw(t *transactor.Transactor, req *types.MsgInitWithdraw) (resp *types.MsgInitWithdrawResp, err error) {
	req.Creator = t.Key.GetAddress().String() // make sure the msg creator is the transactor
	_, err = t.LockSendTx(req)
	return
}

// if err not nil, should return immediately when estimate gas
func SignAgain(t *transactor.Transactor, req *types.MsgSignAgain) (resp *types.MsgSignAgainResp, err error) {
	req.Creator = t.Key.GetAddress().String() // make sure the msg creator is the transactor
	_, err = t.LockSendTx(req)
	return
}
