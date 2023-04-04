<!--
order: 2
-->

# State

## FeePool

All globally tracked parameters for distribution are stored within
`FeePool`. Rewards are collected and added to the reward pool and
distributed to validators/delegators from here.

Note that the reward pool holds decimal coins ([`DecCoins`](https://github.com/cosmos/cosmos-sdk/blob/59d7dc4679b1a3bb2a992b1e3da374743f3c2c9c/proto/cosmos/base/v1beta1/coin.proto#L25)) to allow
for fractions of coins to be received from operations like inflation.
When coins are distributed from the pool they are truncated back to
[`sdk.Coins`](https://github.com/cosmos/cosmos-sdk/blob/59d7dc4679b1a3bb2a992b1e3da374743f3c2c9c/proto/cosmos/base/v1beta1/coin.proto#L14) which are non-decimal.

In Cosmos SDK, the `FeePool` is reserved for distributions from passed community spend proposals. Currently, this functionality is not used in the SGN.

- FeePool: `0x00 | CommunitySpendProposalID -> ProtocolBuffer(FeePool)`

[Protobuf reference](https://github.com/celer-network/sgnv2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/distribution.proto#L88)

## PreviousProposer

This key tracks the previous Tendermint block proposer to receive the fees generated in the block.

- PreviousProposer: `0x01 -> sdk.ConsAddress`

## ValidatorOutstandingRewards

This key records the outstanding rewards for a validator.

- ValidatorOutstandingRewards: `0x02 | ValEthAddrLen (1 byte) | ValEthAddr -> ProtocolBuffer(ValidatorOutstandingRewards)`

[Protobuf reference](https://github.com/celer-network/sgnv2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/distribution.proto#L79)

## DelegatorWithdrawAddress

This key tracks the withdraw address for a delegator. Currently, the SGN does not support setting a custom withdraw address.

- DelegatorWithdrawAddress: `0x03 | DelAddrLen (1 byte) | DelAddr -> WithdrawAddr`

## DelegatorStartingInfo

This key records the delegator starting info.

- DelegatorStartingInfo: `0x04 | ValEthAddrLen (1 byte) | ValEthAddr | DelAddrLen (1 byte) | DelAddr -> ProtocolBuffer(DelegatorStartingInfo)`

[Protobuf reference](https://github.com/celer-network/sgnv2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/distribution.proto#L114)

## ValidatorHistoricalRewards

This key records the historical rewards for a validator.

- ValidatorHistoricalRewards: `0x05 | ValEthAddrLen (1 byte) | ValEthAddr | Period -> ProtocolBuffer(ValidatorHistoricalRewards)`

[Protobuf reference](https://github.com/celer-network/sgnv2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/distribution.proto#L52)

## ValidatorCurrentRewards

This key records the rewards for a validator in the current period.

- ValidatorCurrentRewards: `0x06 | ValEthAddrLen (1 byte) | ValEthAddr -> ProtocolBuffer(ValidatorCurrentRewards)`

[Protobuf reference](https://github.com/celer-network/sgnv2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/distribution.proto#L64)

## ValidatorAccumulatedCommission

This key records the current accumulated commission for a validator.

- ValidatorAccumulatedCommission: `0x07 | ValEthAddrLen (1 byte) | ValEthAddr -> ProtocolBuffer(ValidatorAccumulatedCommission)`

[Protobuf reference](https://github.com/celer-network/sgnv2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/distribution.proto#L72)

## StakingRewardClaimInfo

This key describes the staking reward claim metadata for a delegator.

- StakingRewardClaimInfo: `0x08 | DelAddr -> ProtocolBuffer(StakingRewardClaimInfo)`

[Protobuf reference](https://github.com/celer-network/sgnv2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/distribution/v1/distribution.proto#L147)