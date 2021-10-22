# Command tools
You can use this tools to send tx to onchain contract.
Key store is required.

# Example use in manual test env
For this example, we use a current dir as the root path.
Build the binary first
```
go build
```
## All Command depends on a .toml file.
Here we use manual_test_cbridge.toml
Also, you can get this file from your sgn config path.
--chain according to you .toml file.
## Add Liq
```
./cbridge --cfg ./manual_test_cbridge.toml --chain geth1 --ks ../../test/keys/cethks0.json addLiq usdt 100
```
## Send
```
./cbridge --cfg ./manual_test_cbridge.toml --chain geth1 --ks ../../test/keys/cethks0.json send usdt 884 100 500000
```