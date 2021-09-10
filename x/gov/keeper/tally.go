package keeper

import (
	"github.com/celer-network/sgn-v2/x/gov/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: Break into several smaller functions for clarity

// Tally iterates over the votes and updates the tally of a proposal based on the voting power of the
// voters
func (keeper Keeper) Tally(ctx sdk.Context, proposal types.Proposal) (passes bool, burnDeposits bool, tallyResults types.TallyResult) {
	results := make(map[types.VoteOption]sdk.Int)
	results[types.OptionYes] = sdk.ZeroInt()
	results[types.OptionAbstain] = sdk.ZeroInt()
	results[types.OptionNo] = sdk.ZeroInt()
	results[types.OptionNoWithVeto] = sdk.ZeroInt()

	totalVotingPower := sdk.ZeroInt()
	totalBondedTokens := sdk.ZeroInt()
	currValidators := make(map[string]types.ValidatorGovInfo)

	// fetch all the bonded validators, insert them into currValidators
	keeper.stakingKeeper.IterateBondedValidators(ctx, func(validator stakingtypes.Validator) (stop bool) {
		currValidators[validator.SgnAddress] = types.NewValidatorGovInfo(
			validator.GetSgnAddr(),
			validator.Tokens,
			types.OptionEmpty,
		)

		totalBondedTokens = totalBondedTokens.Add(validator.Tokens)
		return false
	})

	keeper.IterateVotes(ctx, proposal.ProposalId, func(vote types.Vote) bool {
		// if validator, just record it in the map
		valAddrStr := sdk.ValAddress(vote.Voter).String()
		if val, ok := currValidators[valAddrStr]; ok {
			val.Vote = vote.Option
			currValidators[valAddrStr] = val
		}

		acc, _ := sdk.AccAddressFromBech32(vote.Voter)
		keeper.deleteVote(ctx, vote.ProposalId, acc)
		return false
	})

	// iterate over the validators again to tally their voting power
	for _, val := range currValidators {
		if val.Vote == types.OptionEmpty {
			continue
		}

		results[val.Vote] = results[val.Vote].Add(val.Tokens)
		totalVotingPower = totalVotingPower.Add(val.Tokens)
	}

	tallyParams := keeper.GetTallyParams(ctx)
	tallyResults = types.NewTallyResultFromMap(results)

	// If there is not enough params.quorum of votes, the proposal fails
	percentVoting := totalVotingPower.ToDec().QuoInt(totalBondedTokens)
	if percentVoting.LT(tallyParams.Quorum) {
		return false, true, tallyResults
	}

	// If no one votes (everyone abstains), proposal fails
	if totalVotingPower.Equal(results[types.OptionAbstain]) {
		return false, false, tallyResults
	}

	// If more than params.veto of voters veto, proposal fails
	if results[types.OptionNoWithVeto].ToDec().QuoInt(totalVotingPower).GT(tallyParams.Veto) {
		return false, true, tallyResults
	}

	// If more than params.threshold of non-abstaining voters vote Yes, proposal passes
	if results[types.OptionYes].ToDec().QuoInt(totalVotingPower.Sub(results[types.OptionAbstain])).GT(tallyParams.Threshold) {
		return true, false, tallyResults
	}

	// If more than params.threshold of non-abstaining voters vote No, proposal fails
	return false, false, tallyResults
}
