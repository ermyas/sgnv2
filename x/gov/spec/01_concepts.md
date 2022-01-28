<!--
order: 1
-->

# Concepts

The governance process is divided in a few steps that are outlined below:

- **Proposal submission:** Proposal is submitted to the blockchain with a
  deposit.
- **Vote:** Once deposit reaches a certain value (`MinDeposit`), proposal is
  confirmed and vote opens. Bonded validators can then send `TxGovVote`
  transactions to vote on the proposal.

## Proposal submission

### Right to submit a proposal

Any validator, whether bonded or unbonded, can submit proposals by sending a
`TxGovProposal` transaction. Once a proposal is submitted, it is identified by
its unique `proposalID`.

### Proposal types

Proposals below are supported as of the writing of this spec:

- `ParameterChangeProposal` defines a proposal to change one or
  more parameters. If accepted, the requested parameter change is updated
  automatically by the proposal handler upon conclusion of the voting period.
- Software upgrade related proposals:
  - `SoftwareUpgradeProposal`. If accepted, validators are expected to update
  their software in accordance with the proposal. 
  - `CancelSoftwareUpgradeProposal` is a gov Content type for cancelling a software upgrade.
- Cbridge related proposals:
  - `CbrProposal`, to add chains/tokens/contract addresses and update fees etc.
- Pegbridge related proposals:
  - `PegProposal`, to add pagged pairs and update fees etc.
  - `PairDeleteProposal`, to remove a pair.
  - `TotalSupplyUpdateProposal`, to update total supply of a mint token according to the onchain data.
- Farming related proposals: 
  - `AddPoolProposal`
  - `BatchAddPoolProposal`
  - `AdjustRewardProposal`
  - `BatchAdjustRewardProposal`
  - `AddTokensProposal`
  - `SetRewardContractsProposal`
- Message related proposals:
  - `MsgProposal`, to update message bus contract addresses, fees etc.
- Mint related proposals:
  - `AdjustProvisionsProposal`

Other modules expand upon the governance module by implementing their own
proposal types and handlers. These types are registered and processed through the
governance module (eg. `ParamChangeProposal`), which then execute the respective
module's proposal handler when a proposal passes. This custom handler may perform
arbitrary state changes.

## Deposit

To prevent spam, proposals must be submitted with a deposit in the coins defined in the `MinDeposit` param. The voting period will not start until the proposal's deposit equals `MinDeposit`.

When a proposal is submitted, it has to be accompanied by a deposit that must be strictly positive, but can be inferior to `MinDeposit`. The submitter doesn't need to pay for the entire deposit on their own. If a proposal's deposit is inferior to `MinDeposit`, other token holders can increase the proposal's deposit by sending a `Deposit` transaction. The deposit is kept in an escrow in the governance `ModuleAccount` until the proposal is finalized (passed or rejected).

Once the proposal's deposit reaches `MinDeposit`, it enters voting period. If proposal's deposit does not reach `MinDeposit` before `MaxDepositPeriod`, proposal closes and nobody can deposit on it anymore.

### Deposit refund and burn

When a the a proposal finalized, the coins from the deposit are either refunded or burned, according to the final tally of the proposal:

- If the proposal is approved or if it's rejected but _not_ vetoed, deposits will automatically be refunded to their respective depositor (transferred from the governance `ModuleAccount`).
- When the proposal is vetoed with a supermajority, deposits be burned from the governance `ModuleAccount`.

## Vote

### Participants

_Participants_ are users that have the right to vote on proposals. On the
Celer SGN network, participants are bonded validators. Unbonded validators and
other users do not get the right to participate in governance.

### Voting period

Once a proposal reaches `MinDeposit`, it immediately enters `Voting period`. We
define `Voting period` as the interval between the moment the vote opens and
the moment the vote closes. `Voting period` should always be shorter than
`Unbonding period` to prevent double voting. 

### Option set

The option set of a proposal refers to the set of choices a participant can
choose from when casting its vote.

The initial option set includes the following options:

- `Yes`
- `No`
- `NoWithVeto`
- `Abstain`

`NoWithVeto` counts as `No` but also adds a `Veto` vote. `Abstain` option
allows voters to signal that they do not intend to vote in favor or against the
proposal but accept the result of the vote.

### Quorum

Quorum is defined as the minimum percentage of voting power that needs to be
casted on a proposal for the result to be valid.

### Threshold

Threshold is defined as the minimum proportion of `Yes` votes (excluding
`Abstain` votes) for the proposal to be accepted.

Initially, the threshold is set at 50% with a possibility to veto if more than
1/3rd of votes (excluding `Abstain` votes) are `NoWithVeto` votes. This means
that proposals are accepted if the proportion of `Yes` votes (excluding
`Abstain` votes) at the end of the voting period is superior to 50% and if the
proportion of `NoWithVeto` votes is inferior to 1/3 (excluding `Abstain`
votes).

### Validator’s punishment for non-voting

At present, validators are not punished for failing to vote.
