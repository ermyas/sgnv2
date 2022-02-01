<!--
order: 1
-->

# Concepts

In Proof of Stake (PoS) blockchains, rewards gained from transaction fees are paid to validators. The fee distribution module fairly distributes the rewards to the validators' constituent delegators.

Rewards are calculated per period. The period is updated each time a validator's delegation changes, for example, when the validator receives a new delegation.
The rewards for a single validator can then be calculated by taking the total rewards for the period before the delegation started, minus the current total rewards.
To learn more, see the [F1 Fee Distribution paper](https://drops.dagstuhl.de/opus/volltexte/2020/11974/pdf/OASIcs-Tokenomics-2019-10.pdf).

The commission to the validator is paid when the validator is removed or when the validator requests a withdrawal.
The commission is calculated and incremented at every `BeginBlock` operation to update accumulated fee amounts.

The rewards to a delegator are distributed when the delegation is changed or removed, or a withdrawal is requested.

## Staking reward claiming flow

When a delegator claims their rewards, the coins are first sent via the x/bank module to an `sdk.AccAddress`
derived from the delegator's Ethereum address. Upon seeing a successful claim event in Tendermint, the validators
send `MsgSignStakingReward` to co-sign the reward claim. Once +2/3 voting power has signed the claim, the delegator can submit the claim along with the signatures to the `StakingReward` contract on the Ethereum mainchain to receive rewards in the form of ERC-20 CELR token.

[StakingReward contract reference](https://github.com/celer-network/sgn-v2-contracts/blob/ab87d3060fc697f852a6ba4b30ce37483008bf08/contracts/StakingReward.sol)

## Reference Counting in F1 Fee Distribution

In F1 fee distribution, the rewards a delegator receives are calculated when their delegation is withdrawn. This calculation must read the terms of the summation of rewards divided by the share of tokens from the period which they ended when they delegated, and the final period that was created for the withdrawal.

All stored historical rewards records for periods which are no longer referenced by any delegations can thus be safely removed, as they will never be read (future delegations will always reference future periods). This is implemented by tracking a `ReferenceCount`
along with each historical reward storage entry. Each time a new delegation object is created which might need to reference the historical record, the reference count is incremented.
Each time one object which previously needed to reference the historical record is deleted, the reference count is decremented. If the reference count hits zero, the historical record is deleted.

Note that in the upstream Cosmos SDK implementation, slashes to validators need to be tracked by the reference counting as well. In SGN we sync slashing events from the mainchain and update the validator
tokens directly, so the extra reference counting is not needed.
