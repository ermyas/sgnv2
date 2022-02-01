<!--
order: 3
-->

# End-Block

End block calls are executed to update tendermint validator sets and validator syncer role

## Validator Set Changes

Validator states could be updated by the syncer module at the end of this block. Then the abci end-block of staking module scans all the validators that has been updated during this block, and return the list of tendermint validator updates.

[Implementation Reference](https://github.com/celer-network/sgn-v2/blob/f9f76fb10d/x/staking/keeper/validator.go#L210-L234)

## Set Syncer

Update syncer if this block height is an integer multiple of the `SyncerDuration`