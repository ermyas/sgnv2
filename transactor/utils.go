package transactor

import (
	"fmt"

	"github.com/celer-network/sgn-v2/common"
	"github.com/cosmos/cosmos-sdk/client"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func QueryAccount(cliCtx client.Context, sgnAddr string) (account *authtypes.BaseAccount, err error) {

	route := fmt.Sprintf("custom/%s/%s", authtypes.QuerierRoute, authtypes.QueryAccount)
	params := authtypes.QueryAccountRequest{
		Address: sgnAddr,
	}
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return
	}
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	account = new(authtypes.BaseAccount)
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, account)
	return
}
