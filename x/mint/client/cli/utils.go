package cli

import (
	"io/ioutil"

	"github.com/celer-network/sgn-v2/x/mint/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// ParseAdjustProvisionsProposalWithDeposit reads and parses an AdjustProvisionsProposalWithDeposit from a JSON file.
func ParseAdjustProvisionsProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (
	types.AdjustProvisionsProposalWithDeposit, error) {
	proposal := types.AdjustProvisionsProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
