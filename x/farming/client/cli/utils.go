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

// ParseAddTokensProposalWithDeposit reads and parses an AddTokensProposalWithDeposit from a JSON file.
func ParseAddTokensProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (
	types.AddTokensProposalWithDeposit, error) {
	proposal := types.AddTokensProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// ParseAdjustRewardProposalWithDeposit reads and parses an AdjustRewardProposalWithDeposit from a JSON file.
func ParseAdjustRewardProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (
	types.AdjustRewardProposalWithDeposit, error) {
	proposal := types.AdjustRewardProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// ParseSetRewardContractsProposalWithDeposit reads and parses an SetRewardContractsProposalWithDeposit from a JSON file.
func ParseSetRewardContractsProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (
	types.SetRewardContractsProposalWithDeposit, error) {
	proposal := types.SetRewardContractsProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
