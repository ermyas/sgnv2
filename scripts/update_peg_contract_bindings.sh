sgnroot=../
contractroot=../../sgn-v2-contracts
out=$contractroot/out
contractdir=$contractroot/contracts
echo "run solc"
solc --base-path $contractdir --allow-paths $contractdir --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi,bin -o $contractroot/out '@openzeppelin/'=$contractroot/node_modules/@openzeppelin/ $contractdir/pegged/OriginalTokenVault.sol $contractdir/pegged/PeggedTokenBridge.sol $contractdir/pegged/OriginalTokenVaultV2.sol $contractdir/pegged/PeggedTokenBridgeV2.sol
jq '."contracts"|=with_entries(select(.key| test("openzeppelin") or test("^interfaces") or test("^libraries") or test("^safeguard") | not))' $out/combined.json >$out/pegged.json
abigen -combined-json $out/pegged.json -pkg eth -out $sgnroot/eth/bindings_pegged.go
# echo "clean up"
# rm -rf $contractroot/out
# echo "done"
