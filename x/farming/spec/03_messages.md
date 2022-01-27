<!--
order: 3
-->

# Messages

## MsgClaimRewards

This message is for a user to claim their rewards from a single pool. Internally in the farming module, this transaction simultaneously removes the previous `StakeInfo` with associated rewards, the same as if the user simply stakes the same value.
The rewards are sent immediately from the farming `ModuleAccount` to the user address.
The starting height of the `StakeInfo` is set to the current pool period, and the reference count for the previous period is decremented.

In the F1 distribution, the total rewards are calculated per pool period, and a user receives a piece of those rewards in proportion to their stake in the pool.
The total rewards that all the users are entitled to between to periods is calculated the following way:
Let `R(X)` be the total accumulated rewards up to period `X` divided by the tokens staked at that time. The user allocation is `R(X) * user_stake`.
Then the rewards for all the users for staking between periods `A` and `B` are `(R(B) - R(A)) * total stake`.

The final calculated stake is equivalent to the actual staked coins with a margin of error due to rounding errors.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/farming/v1/tx.proto#L23)

## MsgClaimAllRewards

This message claims rewards for a user from all the pools they have stakes in.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/farming/v1/tx.proto#L40)

## MsgSignRewards

This message is sent by the validators to co-sign farming rewards upon seeing a valid claim from a user.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/farming/v1/tx.proto#L62)

## Common operations

These operations take place during different messages.

### Update stake info

Each time a user's stake is changed, the rewards are withdrawn and the `StakeInfo` is updated.
Updating a `StakeInfo` increments the pool period and keeps track of the starting period of the `StakeInfo`.

[Code reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/x/farming/keeper/calc.go#L174)
