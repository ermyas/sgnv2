package cli

import (
	"fmt"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
)

func QueryPendingUpdates(cliCtx client.Context) (pendingUpdates []*types.PendingUpdate, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryPendingUpdates)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, pendingUpdates)
	return
}
