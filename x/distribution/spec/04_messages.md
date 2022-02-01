<!--
order: 4
-->

# Messages

## MsgSetWithdrawAddress

By default, the withdraw address is the delegator address. To change its withdraw address, a delegator must send a `MsgSetWithdrawAddress` message.
Changing the withdraw address is possible only if the parameter `WithdrawAddrEnabled` is set to `true`.

The withdraw address cannot be any of the module accounts. These accounts are blocked from being withdraw addresses by being added to the distribution keeper's `blockedAddrs` array at initialization.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/tx.proto#L38)

## MsgWithdrawDelegatorReward

A delegator can withdraw its rewards.
Internally in the distribution module, this transaction simultaneously removes the previous delegation with associated rewards, the same as if the delegator simply started a new delegation of the same value.
The rewards are sent immediately from the distribution `ModuleAccount` to the withdraw address.
Any remainder (truncated decimals) are sent to the community pool.
The starting height of the delegation is set to the current validator period, and the reference count for the previous period is decremented.
The amount withdrawn is deducted from the `ValidatorOutstandingRewards` variable for the validator.

In the F1 distribution, the total rewards are calculated per validator period, and a delegator receives a piece of those rewards in proportion to their stake in the validator.
In basic F1, the total rewards that all the delegators are entitled to between to periods is calculated the following way.
Let `R(X)` be the total accumulated rewards up to period `X` divided by the tokens staked at that time. The delegator allocation is `R(X) * delegator_stake`.
Then the rewards for all the delegators for staking between periods `A` and `B` are `(R(B) - R(A)) * total stake`.

The final calculated stake is equivalent to the actual staked coins in the delegation with a margin of error due to rounding errors.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/tx.proto#L55)

## MsgWithdrawValidatorCommission

The validator can send the WithdrawValidatorCommission message to withdraw their accumulated commission.
The commission is calculated in every block during `BeginBlock`, so no iteration is required to withdraw.
The amount withdrawn is deducted from the `ValidatorOutstandingRewards` variable for the validator.
Only integer amounts can be sent. If the accumulated awards have decimals, the amount is truncated before the withdrawal is sent, and the remainder is left to be withdrawn later.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/tx.proto#L72)

## MsgFundCommunityPool

This message sends coins directly from the sender to the community pool.

The message fails if the amount cannot be transferred from the sender to the distribution module account.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/tx.proto#L87)

## MsgClaimAllStakingReward

This message claims rewards for a delegator from all the validators they have delegated to.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/tx.proto#L102)

## MsgSignStakingReward

This message is sent by the validators to co-sign staking rewards upon seeing a valid claim from a user.

[Msg reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/tx.proto#L116)

## Common Operations

These operations take place during many different messages.

### Initialize Delegation

Each time a delegation is changed, the rewards are withdrawn and the delegation is reinitialized.
Initializing a delegation increments the validator period and keeps track of the starting period of the delegation.

[Code reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/x/distribution/keeper/delegation.go#L15)
