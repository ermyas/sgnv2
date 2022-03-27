# Steps to generate contract go bindings

In download [solc v0.8.9](https://github.com/ethereum/solidity/releases/tag/v0.8.9) and put the binary `solc` in [scripts](../scripts/) folder.

## Staking and reward contracts bindings

[bindings.go](./bindings.go)

Auto generated along with [sgn-v2-contract](https://github.com/celer-network/sgn-v2-contracts) repo PR by [this script](https://github.com/celer-network/sgn-v2-contracts/blob/main/scripts/solc_abigen.sh).

## Liquidit-Pool Bridge contract bindings

[bindings_cbr.go](./bindings_cbr.go)

In [scripts](../scripts/) folder

```sh
sh update_cbr_contract_bindings.sh
```

## Pegged Bridge contract bindings

[bindings_pegged.go](./bindings_pegged.go)

[bindings_pegged_v2.go](./bindings_pegged_v2.go)

In [scripts](../scripts/) folder

```sh
sh update_peg_contract_bindings.sh
```

## Message contract bindings

[bindings_msg.go](./bindings_msg.go)

[bindings_msg_app.go](../test/contracts/bindings_msg_apps.go)

In [scripts](../scripts/) folder

```sh
sh update_msg_contract_bindings.sh
```

## Bridge test token bindings

[bindings_bridge_test_token.go](./bindings_test_token.go)

In [scripts](../scripts/) folder

```sh
sh update_token_contract_bindings.sh
```

## Withdraw Inbox contract bindings

[bindings_wdinbox.go](./bindings_wdinbox.go)

In [scripts](../scripts/) folder

```sh
sh update_wdinbox_contract_bindings.sh
```

## OVM GasPriceOracle bindings

[bindings_ovm_gas_price_oracle.go](./bindings_ovm_gas_price_oracle.go)

We need to query the OVM_GasPriceOracle to get the right gas price on Optimism. Read [this doc](https://community.optimism.io/docs/developers/l2/new-fees.html#for-frontend-and-wallet-developers) for more information.

First Download [contract OVM_GasPriceOracle.sol](https://github.com/ethereum-optimism/optimism/blob/639e5b13f2ab94b7b49e1f8114ed05a064df8a27/packages/contracts/contracts/L2/predeploys/OVM_GasPriceOracle.sol) and [openzeppelin-contracts-4.3.3](https://github.com/OpenZeppelin/openzeppelin-contracts/releases/tag/v4.3.3)

Then run

```
solc --base-path $PWD --allow-paths . --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi -o . '@openzeppelin/'=openzeppelin-contracts-4.3.3/ OVM_GasPriceOracle.sol
jq '."contracts"|=with_entries(select(.key|test("^openzeppelin")|not))' combined.json > oracle.json
abigen -combined-json ./contracts/oracle.json -pkg eth -out eth/bindings_ovm_gas_price_oracle.go
```
