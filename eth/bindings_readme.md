# Steps to generate contract go bindings

## Staking and reward contracts bindings

[bindings.go](./bindings.go)

Auto generated along with [sgn-v2-contract](https://github.com/celer-network/sgn-v2-contracts) repo PR by [this script](https://github.com/celer-network/sgn-v2-contracts/blob/main/scripts/solc_abigen.sh).

## cBridge contract bindings

[bindings_cbr.go](./bindings_cbr.go)

In [sgn-v2-contract](https://github.com/celer-network/sgn-v2-contracts) repo, go to `contracts` folder, download `openzeppelin-contracts-4.2.0`

Then run
```
solc --base-path $PWD --allow-paths . --overwrite --optimize --optimize-runs 800 --abi --bin -o . '@openzeppelin/'=openzeppelin-contracts-4.2.0/ Bridge.sol
abigen --abi Bridge.abi --bin Bridge.bin --pkg eth --type Bridge > ../../sgn-v2/eth/bindings_cbr.go
```

## Pegged Bridge contract bindings

[bindings_pegged.go](./bindings_pegged.go)

In [sgn-v2-contract](https://github.com/celer-network/sgn-v2-contracts) repo, go to `contracts` folder, download `openzeppelin-contracts-4.2.0`

Then run
```
solc --base-path $PWD --allow-paths . --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi,bin -o . '@openzeppelin/'=openzeppelin-contracts-4.2.0/ pegged/*.sol
jq '."contracts"|=with_entries(select(.key| test("^openzeppelin") or test("^interfaces") or test("^libraries") or test("^safeguard/Pauser") | not))' combined.json > pegged.json
abigen -combined-json ./pegged.json -pkg eth -out ../../sgn-v2/eth/bindings_pegged.go
```

## OVM GasPriceOracle bindings

[bindings_ovm_gas_price_oracle.go](./bindings_ovm_gas_price_oracle.go)

We need to query the OVM_GasPriceOracle to get the right gas price on Optimism. Read [this doc](https://community.optimism.io/docs/developers/l2/new-fees.html#for-frontend-and-wallet-developers) for more information.

First Download [contract OVM_GasPriceOracle.sol](https://github.com/ethereum-optimism/optimism/blob/639e5b13f2ab94b7b49e1f8114ed05a064df8a27/packages/contracts/contracts/L2/predeploys/OVM_GasPriceOracle.sol).

Then run
```
solc --base-path $PWD --allow-paths . --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi -o . '@openzeppelin/'=openzeppelin-contracts-4.3.3/ OVM_GasPriceOracle.sol
jq '."contracts"|=with_entries(select(.key|test("^openzeppelin")|not))' combined.json > oracle.json
abigen -combined-json ./contracts/oracle.json -pkg eth -out eth/bindings_ovm_gas_price_oracle.go
```