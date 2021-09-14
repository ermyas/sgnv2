package keeper

import (
	"fmt"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets distribution information for genesis
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	var moduleHoldings sdk.DecCoins

	k.SetFeePool(ctx, data.FeePool)
	k.SetParams(ctx, data.Params)

	for _, dwi := range data.DelegatorWithdrawInfos {
		delegatorAddress := eth.Hex2Addr(dwi.DelegatorAddress)
		withdrawAddress := eth.Hex2Addr(dwi.WithdrawAddress)
		k.SetDelegatorWithdrawAddr(ctx, delegatorAddress, withdrawAddress)
	}

	var previousProposer sdk.ConsAddress
	if data.PreviousProposer != "" {
		var err error
		previousProposer, err = sdk.ConsAddressFromBech32(data.PreviousProposer)
		if err != nil {
			panic(err)
		}
	}

	k.SetPreviousProposerConsAddr(ctx, previousProposer)

	for _, rew := range data.OutstandingRewards {
		valAddr := eth.Hex2Addr(rew.ValidatorAddress)
		k.SetValidatorOutstandingRewards(ctx, valAddr, types.ValidatorOutstandingRewards{Rewards: rew.OutstandingRewards})
		moduleHoldings = moduleHoldings.Add(rew.OutstandingRewards...)
	}
	for _, acc := range data.ValidatorAccumulatedCommissions {
		valAddr := eth.Hex2Addr(acc.ValidatorAddress)
		k.SetValidatorAccumulatedCommission(ctx, valAddr, acc.Accumulated)
	}
	for _, his := range data.ValidatorHistoricalRewards {
		valAddr := eth.Hex2Addr(his.ValidatorAddress)
		k.SetValidatorHistoricalRewards(ctx, valAddr, his.Period, his.Rewards)
	}
	for _, cur := range data.ValidatorCurrentRewards {
		valAddr := eth.Hex2Addr(cur.ValidatorAddress)
		k.SetValidatorCurrentRewards(ctx, valAddr, cur.Rewards)
	}
	for _, del := range data.DelegatorStartingInfos {
		valAddr := eth.Hex2Addr(del.ValidatorAddress)
		delegatorAddress := eth.Hex2Addr(del.DelegatorAddress)
		k.SetDelegatorStartingInfo(ctx, valAddr, delegatorAddress, del.StartingInfo)
	}
	for _, evt := range data.ValidatorSlashEvents {
		valAddr := eth.Hex2Addr(evt.ValidatorAddress)
		k.SetValidatorSlashEvent(ctx, valAddr, evt.Height, evt.Period, evt.ValidatorSlashEvent)
	}

	moduleHoldings = moduleHoldings.Add(data.FeePool.CommunityPool...)
	moduleHoldingsInt, _ := moduleHoldings.TruncateDecimal()

	// check if the module account exists
	moduleAcc := k.GetDistributionAccount(ctx)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	balances := k.bankKeeper.GetAllBalances(ctx, moduleAcc.GetAddress())
	if balances.IsZero() {
		k.authKeeper.SetModuleAccount(ctx, moduleAcc)
	}
	if !balances.IsEqual(moduleHoldingsInt) {
		panic(fmt.Sprintf("distribution module balance does not match the module holdings: %s <-> %s", balances, moduleHoldingsInt))
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	feePool := k.GetFeePool(ctx)
	params := k.GetParams(ctx)

	dwi := make([]types.DelegatorWithdrawInfo, 0)
	k.IterateDelegatorWithdrawAddrs(ctx, func(del eth.Addr, addr eth.Addr) (stop bool) {
		dwi = append(dwi, types.DelegatorWithdrawInfo{
			DelegatorAddress: del.String(),
			WithdrawAddress:  addr.String(),
		})
		return false
	})

	pp := k.GetPreviousProposerConsAddr(ctx)
	outstanding := make([]types.ValidatorOutstandingRewardsRecord, 0)

	k.IterateValidatorOutstandingRewards(ctx,
		func(addr eth.Addr, rewards types.ValidatorOutstandingRewards) (stop bool) {
			outstanding = append(outstanding, types.ValidatorOutstandingRewardsRecord{
				ValidatorAddress:   addr.String(),
				OutstandingRewards: rewards.Rewards,
			})
			return false
		},
	)

	acc := make([]types.ValidatorAccumulatedCommissionRecord, 0)
	k.IterateValidatorAccumulatedCommissions(ctx,
		func(addr eth.Addr, commission types.ValidatorAccumulatedCommission) (stop bool) {
			acc = append(acc, types.ValidatorAccumulatedCommissionRecord{
				ValidatorAddress: addr.String(),
				Accumulated:      commission,
			})
			return false
		},
	)

	his := make([]types.ValidatorHistoricalRewardsRecord, 0)
	k.IterateValidatorHistoricalRewards(ctx,
		func(val eth.Addr, period uint64, rewards types.ValidatorHistoricalRewards) (stop bool) {
			his = append(his, types.ValidatorHistoricalRewardsRecord{
				ValidatorAddress: val.String(),
				Period:           period,
				Rewards:          rewards,
			})
			return false
		},
	)

	cur := make([]types.ValidatorCurrentRewardsRecord, 0)
	k.IterateValidatorCurrentRewards(ctx,
		func(val eth.Addr, rewards types.ValidatorCurrentRewards) (stop bool) {
			cur = append(cur, types.ValidatorCurrentRewardsRecord{
				ValidatorAddress: val.String(),
				Rewards:          rewards,
			})
			return false
		},
	)

	dels := make([]types.DelegatorStartingInfoRecord, 0)
	k.IterateDelegatorStartingInfos(ctx,
		func(val eth.Addr, del eth.Addr, info types.DelegatorStartingInfo) (stop bool) {
			dels = append(dels, types.DelegatorStartingInfoRecord{
				ValidatorAddress: val.String(),
				DelegatorAddress: del.String(),
				StartingInfo:     info,
			})
			return false
		},
	)

	slashes := make([]types.ValidatorSlashEventRecord, 0)
	k.IterateValidatorSlashEvents(ctx,
		func(val eth.Addr, height uint64, event types.ValidatorSlashEvent) (stop bool) {
			slashes = append(slashes, types.ValidatorSlashEventRecord{
				ValidatorAddress:    val.String(),
				Height:              height,
				Period:              event.ValidatorPeriod,
				ValidatorSlashEvent: event,
			})
			return false
		},
	)

	return types.NewGenesisState(params, feePool, dwi, pp, outstanding, acc, his, cur, dels, slashes)
}
