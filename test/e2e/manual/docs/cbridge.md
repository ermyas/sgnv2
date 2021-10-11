## cBridge

### Setup

Run `go run localnet.go -start -cbr -auto` to start testnet and auto config all nodes as validators.

### Query signers

```sh
sgnd query cbridge latest-signers --home data/node0/sgnd
sgnd query cbridge chain-signers 883 --home data/node0/sgnd
sgnd query cbridge chain-signers 884 --home data/node1/sgnd
```