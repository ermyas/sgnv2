# Concepts

## Original Token 
It is an appellation, not a specific token. Original token refers generally to the token that user is willing to transfer to another chain, or in another word, the token that is deposited to and locked in Original Token Vault on source chain.

## Original Token Vault
It is the smart contract that allows users to deposit different Original Tokens and supplies a method to withdraw their Original Tokens. There would be only one Orgiginal Token Vault on each chain.

## Pegged Token
It is an appellation, not a specific token. Pegged token refers generally to the token that user received on a specified chain, or in another word, the token that minted by Pegged Token Bridge on destination chain. Note that a pegged token is always deployed by a 3rd project party. 

## Pegged Token Bridge
It is the smart contract that allows users to burn differents Pegged Tokens and supplies a method to mint Pegged Tokens. There would be only one Pegged Token Bridge on each chain.

## Deposit&Mint

### Deposit
A deposit represents a valid transaction to Original Token Vault for depositing Original Token. Deposits differ by their Id, which is generated in this [form](https://github.com/celer-network/sgn-v2-contracts/blob/ffb014796cca57168773015e150aba1de81d8bf2/contracts/pegged/OriginalTokenVault.sol#L106).

### Mint
A mint comes from a valid deposit, and represents a transaction to Pegged Token Bridge for minting Pegged Token. Mints differ by their Id, which is generated in this [form](https://github.com/celer-network/sgn-v2/blob/5eff428b7fccb0ea11c9c4d69e2094cdc57bb40c/x/pegbridge/types/request_infos.go#L14).

### Work flow

A Deposited event would be emitted when user successfully deposits Original Token to Original Token Vault. Once sgn monitored and finished syncing that Deposited event, if everything goes well, a deposit and a mint would be stored in pegbridge module, and an [event](04_events.md#mint_to_sign) of sgn would be emitted. After this mint is co-signed by the majority of all validators, a mint transaction would be sent to Pegged Token Bridge by one validator without any user intervention. If that transaction succeeds, user would get their expected Pegged Tokens on destination chain.

## Burn&Withdraw

### Burn
A burn represents a valid transaction to Pegged Token Bridge for burning Pegged Token. Burns differ by their Id, which is generated in this [form](https://github.com/celer-network/sgn-v2-contracts/blob/ffb014796cca57168773015e150aba1de81d8bf2/contracts/pegged/PeggedTokenBridge.sol#L106).

### Withdraw
A withdraw comes from a valid burn, and represents a transaction to Original Token Vault for withdrawing Original Token. Withdraws differ by their Id, which is generated in this [form](https://github.com/celer-network/sgn-v2/blob/5eff428b7fccb0ea11c9c4d69e2094cdc57bb40c/x/pegbridge/types/request_infos.go#L25).

### Work flow

Similar to deposit and mint work flow.


## Supply cap
It is a security number set by 3rd project party, to limit the maximum number of pegged tokens that our Pegged Token Bridge can mint. Once it is hit by any deposit, that deposit would be refunded.

## Refund
Refund generally indicates a cross-chain transferring that failed in the middle, and would be refunded to the user. There would be 3 cases of refund, which are:

1. A too small amount of deposit to cover transaction fee.
2. A too small amount of burn to cover transaction fee.
3. A valid deposit to Original Token Vault, but hits the supply cap of Pegged Token on destination chain.