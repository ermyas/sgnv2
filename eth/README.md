## What is OVM_GasPriceOracle_bindings.go and how to get one?

To get the right gas price on Optimistic Ethereum, we need to query special contract provided by Optimistic.

Read [hear](https://community.optimism.io/docs/developers/l2/new-fees.html#for-frontend-and-wallet-developers) for more information.

First Download [contract OVM_GasPriceOracle.sol](https://github.com/ethereum-optimism/optimism/blob/639e5b13f2ab94b7b49e1f8114ed05a064df8a27/packages/contracts/contracts/L2/predeploys/OVM_GasPriceOracle.sol#L26).

Then run 
```
solc --base-path $PWD --allow-paths . --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi -o . '@openzeppelin/'=openzeppelin-contracts-4.3.3/ OVM_GasPriceOracle.sol
jq '."contracts"|=with_entries(select(.key|test("^openzeppelin")|not))' combined.json>tmp.json
mv tmp.json combined.json
abigen -combined-json ./contracts/combined.json -pkg eth -out eth/OVM_GasPriceOracle_bindings.go
```