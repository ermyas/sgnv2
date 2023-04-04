<!--
order: 1
-->

# State

This section lists the core states of the staking module. There are three ways to trigger state transitions
- Updates triggered by the sync module.
- EndBlock operations.
- Message sent my validator account.

## Validator

Validators are elected by the delegators on the Ethereum staking contract. They run SGN chain under Tendermint-based BFT consensus to offer SGN services and distribute rewards.

### Validator Info

The validator info is defined in this [Protobuf reference](https://github.com/celer-network/sgnv2/blob/f9f76fb10d/proto/sgn/staking/v1/staking.proto#L13-L79). The validator info is stored in the following schema:
```
0x11 | ValidatorEthAddr -> ValidatorInfo
0x12 | ValidatorSgnAddr -> ValidatorEthAddr
0x13 | ValidatorConsAddr -> ValidatorEthAddr
```
The primary key to identify a validator is its Ethereum address. The validator's cosmos-sdk account address (SgnAddr) and consensus address (ConsAddr) serve as the secondary keys.

### Validator Power

Each validator has its associated voting power in the Tendermint consensus engine, which is proportional to the total staked token amount on the Ethereum staking contract. Validator power is stored in the following schema:
```
0x21 | ValidatorEthAddr -> ValidatorPower
0x22 | ValidatorEthAddr -> UpdatedValidatorPower
```
`ValidatorPower` records the voting powers for current bonded validators. `UpdatedValidatorPower` records the voting powers of validators that have been updated during the current block.

### Validator Transactors

Each validator can set multiple additional transactor accounts that are able to send transactions to the SGN chain.
```
0x31 | ValidatorEthAddr -> TransactorList
```

## Delegation

Delegations records the delegators' delegation states on the Ethereum staking contract. Each delegation is identified through a combination of the delegator and validator's Ethereum addresses. The delegation info is defined in this [Protobuf reference](https://github.com/celer-network/sgnv2/blob/f9f76fb10d/proto/sgn/staking/v1/staking.proto#L85-L107), and is stored in the following schema:
```
0x31 | DelegatorEthAddr | ValidatorEthAddr -> DelegationInfo
```

## Syncer

Syncer is a validator that is currently responsible for trigger events sync and submit transactions from and to other blockchains. Validators that belongs to the `SyncerCandidates` list switch syncer role in a round-robin manner for each `SyncerDuration`. Syncer is stored in the following schema:
```
0x51 -> SyncerInfo
```