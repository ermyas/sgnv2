package cli

import (
	"io/ioutil"

	"github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// ParseAddPoolProposalWithDeposit reads and parses an AddPoolProposalWithDeposit from a JSON file.
func ParseAddPoolProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (
	types.AddPoolProposalWithDeposit, error) {
	proposal := types.AddPoolProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
