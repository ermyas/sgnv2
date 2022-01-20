package cli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/transactor"
	govcli "github.com/celer-network/sgn-v2/x/gov/client"
	govrest "github.com/celer-network/sgn-v2/x/gov/client/rest"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/celer-network/sgn-v2/x/message/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetCmdSubmitMsgProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "message-update [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a message buses update proposal",
		Long: `
proposal file is path to json like below
{
	"title": "message buses update",
	"description": "Update message param to v1.1",
	"message_buses": [
		{
			"address": "58712219a4bdbb0e581dcaf6f5c4c2b2d2f42158",
			"chain_id": 884
		}
	],
	"deposit": "0"
}
`,
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
			msgProp := parseMsgProposalJson(args[0])
			msg, _ := govtypes.NewMsgSubmitProposal(msgProp /*msgProp.Deposit*/, sdk.NewInt(0), txr.Key.GetAddress())
			if err := msg.ValidateBasic(); err != nil {
				log.Error(err)
				return err
			}

			txr.CliSendTxMsgWaitMined(msg)
			return nil
		},
	}
	return cmd
}

// parse json at fpath and return MsgProposal
func parseMsgProposalJson(fpath string) *types.MsgProposal {
	ret := new(types.MsgProposal)
	raw, _ := ioutil.ReadFile(fpath)
	json.Unmarshal(raw, ret)
	return ret
}

// ProposalRESTHandler returns a ProposalRESTHandler that exposes the REST handler with a given sub-route.
func MsgProposalRESTHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "message-update",
		Handler:  postMsgProposalHandlerFn(cliCtx),
	}
}
func postMsgProposalHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

var MsgUpdateProposalHandler = govcli.NewProposalHandler(GetCmdSubmitMsgProposal, MsgProposalRESTHandler)
