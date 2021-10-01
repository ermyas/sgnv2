package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
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

func InitWithdraw(t *transactor.Transactor, req *types.MsgInitWithdraw) (resp *types.MsgInitWithdrawResp, err error) {
	txResponse, err := t.SendTxMsgsWaitMined([]sdk.Msg{req})
	if err != nil {
		return
	}

	for _, log := range txResponse.Logs {
		for _, e := range log.Events {
			if e.Type == types.EventTypeMsgResp && e.Attributes[0].Value == "MsgInitWithdrawResp" {
				resp = new(types.MsgInitWithdrawResp)
				resp.Unmarshal([]byte(e.Attributes[1].Value))
				break
			}
		}

		if resp != nil {
			break
		}
	}

	return
}

func SignAgain(t *transactor.Transactor, req *types.MsgSignAgain) (resp *types.MsgSignAgainResp, err error) {
	txResponse, err := t.SendTxMsgsWaitMined([]sdk.Msg{req})
	if err != nil {
		return
	}

	for _, log := range txResponse.Logs {
		for _, e := range log.Events {
			if e.Type == types.EventTypeMsgResp && e.Attributes[0].Value == "MsgSignAgainResp" {
				resp = new(types.MsgSignAgainResp)
				resp.Unmarshal([]byte(e.Attributes[1].Value))
				break
			}
		}

		if resp != nil {
			break
		}
	}

	return
}
