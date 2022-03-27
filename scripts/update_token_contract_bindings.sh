sgnroot=../
contractroot=../../sgn-v2-contracts
out=$contractroot/out
contractdir=$contractroot/contracts

echo "gen bridge test token bindings"
./solc --base-path $contractdir --allow-paths $contractdir --overwrite --optimize --optimize-runs 800 --abi --bin -o $out '@openzeppelin/'=$contractroot/node_modules/@openzeppelin/ $contractdir/pegged-bridge/tokens/MintSwapCanonicalToken.sol
abigen --abi $out/MintSwapCanonicalToken.abi --bin $out/MintSwapCanonicalToken.bin --pkg contracts --type BridgeTestToken > $sgnroot/test/contracts/bindings_test_token.go

echo "clean up"
rm -rf $out
echo "done"
