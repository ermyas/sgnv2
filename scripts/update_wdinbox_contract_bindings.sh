sgnroot=../
contractroot=../../sgn-v2-contracts
out=$contractroot/out
contractdir=$contractroot/contracts

echo "gen withdraw inbox bindings"
./solc --base-path $contractdir --allow-paths $contractdir --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi,bin -o $out '@openzeppelin/'=$contractroot/node_modules/@openzeppelin/ $contractdir/liquidity-bridge/WithdrawInbox.sol $contractdir/integration-examples/ContractAsLP.sol
jq '."contracts"|=with_entries(select(.key| test("^openzeppelin") or test("^interfaces") or test("^libraries") or test("^safeguard/Pauser") or test("^safeguard/Ownable") | not))' $out/combined.json > $out/wdinbox.json
abigen -combined-json $out/wdinbox.json --pkg eth -out $sgnroot/eth/bindings_wdinbox.go

echo "clean up"
rm -rf $out
echo "done"
