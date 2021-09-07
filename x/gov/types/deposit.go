package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewDeposit creates a new Deposit instance
func NewDeposit(proposalID uint64, depositor string, amount sdk.Int) Deposit {
	return Deposit{proposalID, depositor, amount}
}

func (d Deposit) String() string {
	return fmt.Sprintf("deposit by %s on Proposal %d is for the amount %s",
		d.Depositor, d.ProposalId, d.Amount)
}

// Deposits is a collection of Deposit objects
type Deposits []Deposit

func (d Deposits) String() string {
	if len(d) == 0 {
		return "[]"
	}
	out := fmt.Sprintf("Deposits for Proposal %d:", d[0].ProposalId)
	for _, dep := range d {
		out += fmt.Sprintf("\n  %s: %s", dep.Depositor, dep.Amount)
	}
	return out
}

// Equals returns whether two deposits are equal.
func (d Deposit) Equals(comp Deposit) bool {
	return d.Depositor == comp.Depositor && d.ProposalId == comp.ProposalId && d.Amount.BigInt().Cmp(comp.Amount.BigInt()) == 0
}

// Empty returns whether a deposit is empty.
func (d Deposit) Empty() bool {
	return d.Equals(Deposit{})
}

// NewAccTotalDeposit creates a new AccTotalDeposit instance
func NewAccTotalDeposit() AccTotalDeposit {
	return AccTotalDeposit{
		Amount: sdk.ZeroInt(),
	}
}

func (ad AccTotalDeposit) String() string {
	return fmt.Sprintf("Amount: %s", ad.Amount)
}
