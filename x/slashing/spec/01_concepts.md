<!--
order: 1
-->

# Concepts

## States

At any given time, there are any number of validators registered in the state
machine. Each block, the top `MaxValidators` (defined by `x/staking`) validators
who are not jailed become _bonded_, meaning that they may propose and vote on
blocks. Validators who are _bonded_ are _at stake_, meaning that part or all of
their stake and their delegators' stake is at risk if they commit a protocol fault.

For each of these validators we keep a `ValidatorSigningInfo` record that contains
information pertaining to validator's liveness and other infraction related
attributes.

## Infraction Timelines

To illustrate how the `x/slashing` module handles submitted evidence through
Tendermint consensus, consider the following examples:

**Definitions**:

_[_ : timeline start
_]_ : timeline end
_C<sub>n</sub>_ : infraction `n` committed
_D<sub>n</sub>_ : infraction `n` discovered
_V<sub>b</sub>_ : validator bonded
_V<sub>u</sub>_ : validator unbonded

### Single Double Sign Infraction

<----------------->
[----------C<sub>1</sub>----D<sub>1</sub>,V<sub>u</sub>-----]

A single infraction is committed then later discovered, at which point the
validator is unbonded and slashed at the full amount for the infraction.

### Multiple Double Sign Infractions

<--------------------------->
[----------C<sub>1</sub>--C<sub>2</sub>---C<sub>3</sub>---D<sub>1</sub>,D<sub>2</sub>,D<sub>3</sub>V<sub>u</sub>-----]

Multiple infractions are committed and then later discovered, at which point the
validator is jailed and slashed for only one infraction. Because the validator
is also tombstoned, they can not rejoin the validator set.
