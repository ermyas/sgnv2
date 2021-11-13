package slashing

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/slashing/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// Iterate over all the validators which *should* have signed this block
	// store whether or not they have actually signed it and slash/unbond any
	// which have missed too many blocks in a row (downtime slashing)
	for _, voteInfo := range req.LastCommitInfo.GetVotes() {
		k.HandleValidatorSignature(ctx, voteInfo.Validator.Address, voteInfo.SignedLastBlock, req.Header.Time)
	}

	// Iterate through any newly discovered evidence of infraction
	// Slash any validators (and since-unbonded stake within the unbonding period)
	// who contributed to valid infractions
	for _, evidence := range req.ByzantineValidators {
		switch evidence.Type {
		case abci.EvidenceType_DUPLICATE_VOTE:
			k.HandleDoubleSign(ctx, evidence.Validator.Address, req.Header.Time)
		default:
			log.Errorf("ignored unknown evidence type: %s", evidence.Type)
		}
	}
}
