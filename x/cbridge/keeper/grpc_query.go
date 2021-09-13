package keeper

import (
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

var _ types.QueryServer = Keeper{}
