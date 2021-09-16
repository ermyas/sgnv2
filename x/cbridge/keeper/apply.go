package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// data is serialized OnChainEvent
func (k Keeper) ApplyEvent(ctx sdk.Context, data []byte) (bool, error) {
	onchev := new(types.OnChainEvent)
	err := onchev.Unmarshal(data)
	if err != nil {
		return false, err
	}
	elog := new(ethtypes.Log)
	err = json.Unmarshal(onchev.Elog, elog)
	if err != nil {
		return false, err
	}
	kv := ctx.KVStore(k.storeKey)
	cbrContract, _ := eth.NewBridgeFilterer(eth.ZeroAddr, nil)
	switch onchev.Evtype {
	case types.CbrEventLiqAdd:
		ev, err := cbrContract.ParseLiquidityAdded(*elog)
		if err != nil {
			return false, err
		}
		if HasEvLiqAdd(kv, onchev.Chainid, ev.Seqnum) {
			// already processed, return error
			return false, fmt.Errorf("already processed liq add event: chainid %d seq %d", onchev.Chainid, ev.Seqnum)
		}
		SetEvLiqAdd(kv, onchev.Chainid, ev.Seqnum)
		newliq := ChangeLiquidity(kv, onchev.Chainid, ev.Token, ev.Provider, ev.Amount)

		k.Logger(ctx).Info("Applied LP add_liquidity", "LQKey", types.LiqMapKey(onchev.Chainid, ev.Token, ev.Provider), "NewAmt", newliq.String())
		return true, nil
	}
	return true, nil
}
