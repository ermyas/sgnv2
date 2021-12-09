## Bridge

### Setup

Run `echo 12341234 | go run localnet.go -start -cbr -auto` to start testnet and auto config all nodes as validators.

### Query signers

```sh
sgnd query cbridge latest-signers --home data/node0/sgnd
sgnd query cbridge chain-signers 883 --home data/node0/sgnd
sgnd query cbridge chain-signers 884 --home data/node0/sgnd
```

### Query config

```sh
sgnd query cbridge config --home data/node0/sgnd
sgnd query pegbridge config --home data/node0/sgnd
```