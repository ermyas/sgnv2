package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
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

	// this line is used by starport scaffolding # 1

	return cmd
}

// if err not nil, should return immediately when estimate gas
func InitWithdraw(t *transactor.Transactor, req *types.MsgInitWithdraw) (resp *types.MsgInitWithdrawResp, err error) {
	_, err = t.LockSendTx(req)
	return
}

// if err not nil, should return immediately when estimate gas
func SignAgain(t *transactor.Transactor, req *types.MsgSignAgain) (resp *types.MsgSignAgainResp, err error) {
	_, err = t.LockSendTx(req)
	return
}
