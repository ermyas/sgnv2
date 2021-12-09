package cli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	govcli "github.com/celer-network/sgn-v2/x/gov/client"
	govrest "github.com/celer-network/sgn-v2/x/gov/client/rest"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetCmdSubmitCbrProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cbridge-change [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a cbridge config change proposal",
		Long: `
proposal file is path to json like below
{
	"title": "cBridge config Change",
	"description": "Update cbridge param to v1.1",
	"cbr_config": {
		"lp_fee_perc": 80,
		"assets": [{}],
		"chain_pairs": [{}]
	},
	"deposit": "1000"
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
			cbrProp := parseJson(args[0])
			// enforce deposit to be zero
			msg, _ := govtypes.NewMsgSubmitProposal(cbrProp /*cbrProp.Deposit*/, sdk.NewInt(0), txr.Key.GetAddress())
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

// parse json at fpath and return CbrProposal
func parseJson(fpath string) *types.CbrProposal {
	ret := new(types.CbrProposal)
	raw, _ := ioutil.ReadFile(fpath)
	json.Unmarshal(raw, ret)
	return ret
}

// ProposalRESTHandler returns a ProposalRESTHandler that exposes the REST handler with a given sub-route.
func CbrProposalRESTHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "cbridge-change",
		Handler:  postCbrProposalHandlerFn(cliCtx),
	}
}
func postCbrProposalHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

var CbrConfigProposalHandler = govcli.NewProposalHandler(GetCmdSubmitCbrProposal, CbrProposalRESTHandler)
