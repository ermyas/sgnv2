package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/message/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
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

	return cmd
}

// if err not nil, should return immediately when estimate gas
func ClaimAllFees(t *transactor.Transactor, req *types.MsgClaimAllFees) (resp *types.MsgClaimAllFeesResponse, err error) {
	req.Sender = t.Key.GetAddress().String()
	_, err = t.LockSendTx(req)
	return
}
