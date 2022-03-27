sgnroot=../
contractroot=../../sgn-v2-contracts
out=$contractroot/out
contractdir=$contractroot/contracts

echo "gen peg v0 bindings"
./solc --base-path $contractdir --allow-paths $contractdir --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi,bin -o $out '@openzeppelin/'=$contractroot/node_modules/@openzeppelin/ $contractdir/pegged-bridge/OriginalTokenVault.sol $contractdir/pegged-bridge/PeggedTokenBridge.sol
jq '."contracts"|=with_entries(select(.key| test("openzeppelin") or test("^interfaces") or test("^libraries") or test("^safeguard") | not))' $out/combined.json >$out/pegged.json
abigen -combined-json $out/pegged.json -pkg eth -out $sgnroot/eth/bindings_pegged.go

echo "gen peg v2 bindings"
./solc --base-path $contractdir --allow-paths $contractdir --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi,bin -o $out '@openzeppelin/'=$contractroot/node_modules/@openzeppelin/ $contractdir/pegged-bridge/OriginalTokenVaultV2.sol $contractdir/pegged-bridge/PeggedTokenBridgeV2.sol
jq '."contracts"|=with_entries(select(.key| test("openzeppelin") or test("^interfaces") or test("^libraries") or test("^safeguard") | not))' $out/combined.json >$out/pegged.json
abigen -combined-json $out/pegged.json -pkg eth -out $sgnroot/eth/bindings_pegged_v2.go

echo "clean up"
rm -rf $out
echo "done"
