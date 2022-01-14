package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

const (
	flagChain = "chain"
	flagToken = "token"
	flagFile  = "file"
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
		GetCmdValidatorClaimFee(),
	)...)

	return cmd
}

func GetCmdSignAgainRelay() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-again-relay [src-transfer-id]",
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
			var wds []*types.WithdrawLq

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

			var reqIds []uint64
			if file != "" {
				f, err := os.Open(file)
				if err != nil {
					return err
				}
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					line := scanner.Text()
					fields := strings.Fields(line)
					if len(fields) != 3 {
						return fmt.Errorf("invalid file input: %s", line)
					}
					reqId, err := strconv.Atoi(fields[0])
					if err != nil {
						return err
					}
					reqIds = append(reqIds, uint64(reqId))

					chainId, err := strconv.Atoi(fields[1])
					if err != nil {
						return err
					}
					token = fields[2]
					wd := &types.WithdrawLq{
						FromChainId: uint64(chainId),
						TokenAddr:   token,
					}
					wds = append(wds, wd)
				}
			} else if chainId != 0 && token != "" {
				wd := &types.WithdrawLq{
					FromChainId: chainId,
					TokenAddr:   token,
				}
				wds = append(wds, wd)
				reqId := uint64(time.Now().Unix())
				reqIds = append(reqIds, reqId)
				fmt.Println("Withdraw request Id:", reqId)
			} else {
				return fmt.Errorf("invalid flag inputs")
			}

			var msgs []sdk.Msg
			for i, wd := range wds {
				withdrawReq := &types.WithdrawReq{
					Withdraws:    []*types.WithdrawLq{wd},
					ExitChainId:  wd.FromChainId,
					ReqId:        reqIds[i],
					WithdrawType: types.ValidatorClaimFeeShare,
				}
				wdBytes, _ := withdrawReq.Marshal()
				msg := &types.MsgInitWithdraw{
					WithdrawReq: wdBytes,
					Creator:     txr.Key.GetAddress().String(),
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
