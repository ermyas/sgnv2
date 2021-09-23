package cli

import (
	"encoding/json"
	"io/ioutil"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/transactor"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
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
		"lp_fee": 80,
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
			msg, _ := govtypes.NewMsgSubmitProposal(cbrProp, cbrProp.Deposit, txr.Key.GetAddress())
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
func parseJson(fpath string) *govtypes.CbrProposal {
	ret := new(govtypes.CbrProposal)
	raw, _ := ioutil.ReadFile(fpath)
	json.Unmarshal(raw, ret)
	return ret
}
