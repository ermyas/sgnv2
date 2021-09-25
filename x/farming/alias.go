package farming

import (
	"github.com/celer-network/sgn-v2/x/farming/keeper"
	"github.com/celer-network/sgn-v2/x/farming/types"
)

const (
	StoreKey                = types.StoreKey
	ModuleName              = types.ModuleName
	RewardModuleAccountName = types.RewardModuleAccountName
	RouterKey               = types.RouterKey
)

var (
	NewKeeper          = keeper.NewKeeper
	RegisterInvariants = keeper.RegisterInvariants
)

type (
	Keeper = keeper.Keeper
)
