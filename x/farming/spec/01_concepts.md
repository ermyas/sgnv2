<!--
order: 1
-->

# Concepts

This module distributes farming rewards fairly to the constituent users of each pool.

Rewards are calculated per period. The period is updated each time a pool's stake changes, for example, when a new user enters the pool.
The rewards for a single pool can then be calculated by taking the total rewards for the period before the staking event happened, minus the current total rewards.
To learn more, see the [F1 Fee Distribution paper](https://drops.dagstuhl.de/opus/volltexte/2020/11974/pdf/OASIcs-Tokenomics-2019-10.pdf).

The rewards to a user are distributed whenever their stake is changed.

## Reference Counting in F1 Fee Distribution

In F1 fee distribution, the rewards a user receives are calculated when their stake is changed. This calculation must read the terms of the summation of rewards divided by the share of tokens from the period which they ended when they staked, and the final period that was created for the withdrawal.

All stored historical rewards records for periods which are no longer referenced by any `StakeInfo` can thus be safely removed, as they will never be read (future staking operations will always reference future periods). This is implemented by tracking a `ReferenceCount`
along with each historical reward storage entry. Each time a new `StakeInfo`
is created which might need to reference the historical record, the reference count is incremented.
Each time one object which previously needed to reference the historical record is deleted, the reference
count is decremented. If the reference count hits zero, the historical record is deleted.
