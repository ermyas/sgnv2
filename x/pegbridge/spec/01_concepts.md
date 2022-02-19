# Concepts

## Original Token 
Original token refers to the existing token that is deposited to and locked in Original Token Vault on one chain, which would triggers the mint of corresponding pegged token on a remote chain.

## Pegged Token
Pegged token refers to the token that is minted by Pegged Token Bridge on one chain, with value 1:1 pegged to original token locked in the original token vault at a remote chain. 

## Original Token Vault
It is a smart contract to deposit (triggers remote pegged mint) and withdraw (triggered by remote pegged burn) original tokens. Each chain can has multiple vault contracts ([version 0](https://github.com/celer-network/sgn-v2-contracts/blob/main/contracts/pegged/OriginalTokenVault.sol), [version 2](https://github.com/celer-network/sgn-v2-contracts/blob/main/contracts/pegged/OriginalTokenVaultV2.sol)). Each original token can only use one vault contract.

### Deposit
A deposit represents a valid transaction to Original Token Vault for depositing Original Token. Deposits differ by their Ids ([version 0](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/OriginalTokenVault.sol#L106-L109), [version 2](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/OriginalTokenVaultV2.sol#L109-L121)).

### Withdraw
A withdraw comes from a valid burn, and represents a transaction to Original Token Vault for withdrawing Original Token. Withdraws differ by their Ids ([version 0](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/OriginalTokenVault.sol#L132-L142), [version 2](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/OriginalTokenVaultV2.sol#L144-L155)).

## Pegged Token Bridge
It is the bridge contract to mint (triggered by remote vault deposit or pegged burn) and burn (triggers remote vault withdraw) pegged token. Each chain can has multiple pegged bridge contracts ([version 0](https://github.com/celer-network/sgn-v2-contracts/blob/main/contracts/pegged/PeggedTokenBridge.sol), [version 2](https://github.com/celer-network/sgn-v2-contracts/blob/main/contracts/pegged/PeggedTokenBridgeV2.sol)).

### Mint
A represents a transaction to Pegged Token Bridge for minting Pegged Token. Mints differ by their Ids ([version 0](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/PeggedTokenBridge.sol#L58-L68), [version 2](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/PeggedTokenBridgeV2.sol#L66-L77))

### Burn
A burn represents a valid transaction to Pegged Token Bridge for burning Pegged Token. Burns differ by their Ids ([version 0](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/PeggedTokenBridge.sol#L106-L109), [version 2](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/PeggedTokenBridgeV2.sol#L118-L130)).

## Work flow

### Vault Deposit -> Pegged Mint

A Deposited event would be emitted when user successfully deposits Original Token to Original Token Vault. Once sgn monitored and finished syncing that Deposited event, if everything goes well, a deposit and a mint would be stored in pegbridge module, and an [event](04_events.md#mint_to_sign) of sgn would be emitted. After this mint is co-signed by the majority of all validators, a mint transaction would be sent to Pegged Token Bridge by one validator without any user intervention. If that transaction succeeds, user would get their expected Pegged Tokens on destination chain.

### Pegged Burn -> Vault Withdraw

Similar to deposit and deposit->mint work flow.

### Pegged Burn -> Pegged Mint

Similar to flow above. One note is that this is only available for [V2 bridge contract](https://github.com/celer-network/sgn-v2-contracts/blob/71a195582a/contracts/pegged/PeggedTokenBridgeV2.sol#L112), where user can specify the destination chain for the burn. If destination chain is not zero or the original token chain, it will trigger a mint at another chain that has the pegged token.

## Supply cap
It is a security number set by 3rd project party, to limit the maximum number of pegged tokens that our Pegged Token Bridge can mint. Once it is hit by any deposit, that deposit would be refunded.

## Refund
Refund generally indicates a cross-chain transferring that failed in the middle, and would be refunded to the user. There would be 3 cases of refund, which are:

1. A too small amount of deposit to cover transaction fee.
2. A too small amount of burn to cover transaction fee.
3. A valid deposit to Original Token Vault, but hits the supply cap of Pegged Token on destination chain.