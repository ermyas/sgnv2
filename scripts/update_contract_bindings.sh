sgnroot=../
contractroot=../../sgn-v2-contracts
out=$contractroot/out/
contractdir=$contractroot/contracts
echo "run solc"
solc --base-path $contractdir --allow-paths $contractdir --overwrite --optimize --optimize-runs 800 --pretty-json --combined-json abi,bin -o $contractroot/out '@openzeppelin/'=$contractroot/node_modules/@openzeppelin/ $contractdir/message/apps/*.sol $contractdir/message/messagebus/MessageBus.sol
jq '."contracts"|=with_entries(select(.key| test("^openzeppelin") or test("^interfaces") or test("^libraries") or test("^safeguard/Pauser") | not))' $out/combined.json >$out/message.json
echo "run abigen, output to sgn-v2 repo"
abigen -combined-json $out/message.json -pkg eth -out $sgnroot/eth/bindings_msg.go
echo "clean up"
rm -rf $contractroot/out
echo "done"
