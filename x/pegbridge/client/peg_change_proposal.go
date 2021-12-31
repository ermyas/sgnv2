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
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetCmdSubmitPegProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pegbridge-change [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a pegbridge config change proposal",
		Long: `
proposal file is path to json like below
{
	"title": "peg bridge config change",
	"description": "Update pegbridge param to v1.1",
	"peg_config": {
		"orig_pegged_pairs": [
			{
			"orig": {
				"address": "3ff73bab93c505809c68b0a8e4321a2713d9255c",
				"chain_id": 883,
				"decimals": 18,
				"symbol": "UNI"
			},
			"pegged": {
				"address": "283ab9db53f25d84fa30915816ec53f8affaa86e",
				"chain_id": 884,
				"decimals": 18,
				"symbol": "UNI"
			}
			}
		],
		"original_token_vaults": [
			{
			"address": "14558ead4a122d7fb2e711242500c12963320f20",
			"chain_id": 883
			}
		],
		"pegged_token_bridges": [
			{
			"address": "58712219a4bdbb0e581dcaf6f5c4c2b2d2f42158",
			"chain_id": 884
			}
		]
    },
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
			pegProp := parsePegProposalJson(args[0])
			msg, _ := govtypes.NewMsgSubmitProposal(pegProp /*pegProp.Deposit*/, sdk.NewInt(0), txr.Key.GetAddress())
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

func GetCmdSubmitPairDeleteProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pegged-pair-delete [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a pegged pair delete proposal",
		Long: `
proposal file is path to json like below
{
	"title": "peg bridge pair delete",
	"description": "delete a pair",
	"pair_to_delete": {
		"orig": {
			"address": "3ff73bab93c505809c68b0a8e4321a2713d9255c",
			"chain_id": 883
		},
		"pegged": {
			"address": "283ab9db53f25d84fa30915816ec53f8affaa86e",
			"chain_id": 884
		}
    },
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
			prop := parsePairDeleteProposalJson(args[0])
			msg, _ := govtypes.NewMsgSubmitProposal(prop, prop.Deposit, txr.Key.GetAddress())
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

func GetCmdSubmitTotalSupplyUpdateProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-supply-update [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a pegged total supply update proposal",
		Long: `
proposal file is path to json like below
{
	"title": "peg bridge total supply update",
	"description": "update a pair's total supply'",
	"pair": {
			"orig": {
				"address": "3ff73bab93c505809c68b0a8e4321a2713d9255c",
				"chain_id": 883,
			},
			"pegged": {
				"address": "283ab9db53f25d84fa30915816ec53f8affaa86e",
				"chain_id": 884,
			}
		}
    },
	"total_supply": "100",
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
			prop := parseTotalSupplyUpdateProposalJson(args[0])
			msg, _ := govtypes.NewMsgSubmitProposal(prop, prop.Deposit, txr.Key.GetAddress())
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

// parse json at fpath and return PegProposal
func parsePegProposalJson(fpath string) *types.PegProposal {
	ret := new(types.PegProposal)
	raw, _ := ioutil.ReadFile(fpath)
	json.Unmarshal(raw, ret)
	return ret
}

// parse json at fpath and return PairDeleteProposal
func parsePairDeleteProposalJson(fpath string) *types.PairDeleteProposal {
	ret := new(types.PairDeleteProposal)
	raw, _ := ioutil.ReadFile(fpath)
	json.Unmarshal(raw, ret)
	return ret
}

// parse json at fpath and return TotalSupplyUpdateProposal
func parseTotalSupplyUpdateProposalJson(fpath string) *types.TotalSupplyUpdateProposal {
	ret := new(types.TotalSupplyUpdateProposal)
	raw, _ := ioutil.ReadFile(fpath)
	json.Unmarshal(raw, ret)
	return ret
}

// ProposalRESTHandler returns a ProposalRESTHandler that exposes the REST handler with a given sub-route.
func PegProposalRESTHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "pegbridge-change",
		Handler:  postPegProposalHandlerFn(cliCtx),
	}
}
func postPegProposalHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func PairDeleteProposalRESTHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "pegged-pair-delete",
		Handler:  postPairDeleteProposalHandlerFn(cliCtx),
	}
}
func postPairDeleteProposalHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func TotalSupplyUpdateProposalRESTHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "total-supply-update",
		Handler:  postTotalSupplyUpdateProposalHandlerFn(cliCtx),
	}
}
func postTotalSupplyUpdateProposalHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

var PegConfigProposalHandler = govcli.NewProposalHandler(GetCmdSubmitPegProposal, PegProposalRESTHandler)
var PegPairDeleteProposalHandler = govcli.NewProposalHandler(GetCmdSubmitPairDeleteProposal, PairDeleteProposalRESTHandler)
var TotalSupplyUpdateProposalHandler = govcli.NewProposalHandler(GetCmdSubmitTotalSupplyUpdateProposal, TotalSupplyUpdateProposalRESTHandler)
