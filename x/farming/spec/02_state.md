<!--
order: 2
-->

# State

## FarmingPool

All parameters for a farming pool are stored within the
`FarmingPool` struct. Rewards are collected and added to the reward pool and distributed to users.

Note that the reward pool holds decimal coins ([`DecCoins`](https://github.com/cosmos/cosmos-sdk/blob/59d7dc4679b1a3bb2a992b1e3da374743f3c2c9c/proto/cosmos/base/v1beta1/coin.proto#L25)) to allow
for fractions of coins to be received from operations like inflation.
When coins are distributed from the pool they are truncated back to
[`sdk.Coins`](https://github.com/cosmos/cosmos-sdk/blob/59d7dc4679b1a3bb2a992b1e3da374743f3c2c9c/proto/cosmos/base/v1beta1/coin.proto#L14) which are non-decimal.

- FarmingPool: `0x01 | PoolName -> ProtocolBuffer(FarmingPool)`

[Protobuf reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/farming/v1/farming.proto#L27)

## AddressInFarmingPool

This key is used to track the users in each farming pool.

- AddressInFarmingPool: `0x02 | PoolName | UserAddr -> ""`

## StakeInfo

This key tracks the stake of a user in a farming pool.

- StakeInfo: `0x03 | UserAddr | PoolName -> ProtocolBuffer(StakeInfo)`

[Protobuf reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/farming/v1/farming.proto#L63)

## PoolHistoricalRewards

This key records the reward ratio of one user in a pool.

- PoolHistoricalRewards: `0x04 | PoolName -> ProtocolBuffer(PoolHistoricalRewards)`

[Protobuf reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/farming/v1/farming.proto#L72)

## PoolCurrentRewards

This key records the rewards in the current period.

- PoolCurrentRewards: `0x05 | PoolName -> ProtocolBuffer(PoolCurrentRewards)`

[Protobuf reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/farming/v1/farming.proto#L82)

## RewardClaimInfo

This key describes the reward claim metadata and details for a recipient.

- RewardClaimInfo: `0x06 | UserAddr -> ProtocolBuffer(RewardClaimInfo)`

[Protobuf reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/farming/v1/farming.proto#L139)

## ERC20Token

This key describes an ERC20 token on a specific EVM-compatible chain.

- ERC20Token: `0x07 | ChainId | Symbol -> ProtocolBuffer(ERC20Token)`

[Protobuf reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/common/v1/common.proto#L257)

## RewardContract

This key is used to store the info about the farming reward contracts.

- RewardContract: `0x08 | ChainId -> ProtocolBuffer(ContractInfo)`

[Protobuf reference](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/common/v1/common.proto#L17)
