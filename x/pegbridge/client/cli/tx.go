package cli

import (
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagChain                  = "chain"
	flagToken                  = "token"
	flagFile                   = "file"
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
		GetCmdValidatorClaimFee(),
	)...)

	return cmd
}

func InitClaimFee(t *transactor.Transactor, req *types.MsgClaimFee) (resp *types.MsgClaimFeeResponse, err error) {
	req.Sender = t.Key.GetAddress().String() // make sure the msg sender is the transactor
	_, err = t.LockSendTx(req)
	return
}

func InitClaimRefund(t *transactor.Transactor, req *types.MsgClaimRefund) (resp *types.MsgClaimRefundResponse, err error) {
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

func GetCmdValidatorClaimFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-claim-fee",
		Short: "Validator claim fee shares",
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
			chainId, err := cmd.Flags().GetUint64(flagChain)
			if err != nil {
				return err
			}
			token, err := cmd.Flags().GetString(flagToken)
			if err != nil {
				return err
			}
			file, err := cmd.Flags().GetString(flagFile)
			if err != nil {
				return err
			}

			reqIds, chainIds, tokens, err := cbrcli.ValidatorClaimFeeHelper(chainId, token, file)
			if err != nil {
				return fmt.Errorf("ValidatorClaimFeeHelper err: %s", err)
			}

			var msgs []sdk.Msg
			for i, chainId := range chainIds {
				msg := &types.MsgClaimFee{
					DelegatorAddress: "",
					ChainId:          chainId,
					TokenAddress:     tokens[i],
					Nonce:            reqIds[i],
					Signature:        []byte{},
					Sender:           txr.Key.GetAddress().String(),
					IsValidator:      true,
				}
				msgs = append(msgs, msg)
			}

			txr.CliSendTxMsgsWaitMined(msgs)
			return nil
		},
	}
	cmd.Flags().Uint64(flagChain, 0, "chain id")
	cmd.Flags().String(flagToken, "", "token address")
	cmd.Flags().String(flagFile, "", "file that contains a list of <reqid, chainid, token> tuples")

	return cmd
}

// if err not nil, should return immediately when estimate gas
func SignAgainWithdraw(txr *transactor.Transactor, msg *types.MsgTriggerSignWithdraw) (err error) {
	err = msg.ValidateBasic()
	if err != nil {
		return
	}
	txr.CliSendTxMsgWaitMined(msg)
	return
}
