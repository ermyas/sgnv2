<!--
order: 5
-->

# Hooks

Available hooks that can be called by and from this module.

## Delegation created or modified

- triggered-by delegation changes in `x/staking`

### Before

- The delegation rewards are withdrawn to the withdraw address of the delegator.
  The rewards include the current period and exclude the starting period.
- The validator period is incremented.
  The validator period is incremented because the validator's power and share distribution might have changed.
- The reference count for the delegator's starting period is decremented.

### After

The starting height of the delegation is set to the previous period.
Because of the `Before`-hook, this period is the last period for which the delegator was rewarded.

## Validator created

- triggered-by validator creation in `x/staking`

When a validator is created, the following validator variables are initialized:

- Historical rewards
- Current accumulated rewards
- Accumulated commission
- Total outstanding rewards
- Period

By default, all values are set to a `0`, except period, which is set to `1`.

## Validator removed

- triggered-by validator removal in `x/staking`

Outstanding commission is sent to the validator's self-delegation withdrawal address.
Remaining delegator rewards get sent to the community fee pool.

Note: The validator gets removed only when it has no remaining delegations.
At that time, all outstanding delegator rewards will have been withdrawn.
Any remaining rewards are dust amounts.
