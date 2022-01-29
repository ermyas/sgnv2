<!--
order: 3
-->

# Messages

## Proposal Submission

Proposals can be submitted via a `MsgSubmitProposal`
transaction.

The `Content` of a `MsgSubmitProposal` message must have an appropriate router
set in the governance module.

**State modifications:**

- Generate new `proposalID`
- Create new `Proposal`
- Initialize `Proposals` attributes
- Decrease balance of sender by `InitialDeposit`
- If `MinDeposit` is reached:
    - Push `proposalID` in `ActiveProposalQueue`
- Transfer `InitialDeposit` from the `Proposer` to the governance `ModuleAccount`

## Deposit

Once a proposal is submitted, if
`Proposal.TotalDeposit < ActiveParam.MinDeposit`, validators can send
`MsgDeposit` transactions to increase the proposal's deposit.

**State modifications:**

- Decrease balance of sender by `deposit`
- Add `deposit` of sender in `proposal.Deposits`
- Increase `proposal.TotalDeposit` by sender's `deposit`
- If `MinDeposit` is reached:
    - Push `proposalID` in `ActiveProposalQueue`
- Transfer `Deposit` from the `proposer` to the governance `ModuleAccount`

A `MsgDeposit` transaction has to go through a number of checks to be valid.

## Vote

Once `ActiveParam.MinDeposit` is reached, voting period starts. From there,
bonded validators are able to send `MsgVote` transactions to cast their
vote on the proposal.