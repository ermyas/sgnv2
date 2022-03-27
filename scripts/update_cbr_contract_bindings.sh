sgnroot=../
contractroot=../../sgn-v2-contracts
out=$contractroot/out
contractdir=$contractroot/contracts

echo "gen liquidity-pool bridge bindings"
./solc --base-path $contractdir --allow-paths $contractdir --overwrite --optimize --optimize-runs 800 --abi --bin -o $out '@openzeppelin/'=$contractroot/node_modules/@openzeppelin/ $contractdir/liquidity-bridge/Bridge.sol
abigen --abi $out/Bridge.abi --bin $out/Bridge.bin --pkg eth --type Bridge > $sgnroot/eth/bindings_cbr.go

echo "clean up"
rm -rf $out
echo "done"
