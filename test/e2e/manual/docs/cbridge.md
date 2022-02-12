## Bridge

### Setup

Run `echo 12341234 | go run localnet.go -start -cbr -auto -op` to start testnet and auto config all nodes as validators.

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

### Validator withdraw fees

```sh
sgnd query cbridge fee-share 00078b31fa8b29a76bce074b5ea0d515a6aeaee7
echo 12341234 | sgnd tx cbridge validator-claim-fee --file data/claim_fee.txt --home data/node0/sgnd
echo 12341234 | sgnd ops validator withdraw-cbr-fee --file data/claim_fee.txt --home data/node0/sgnd --query
```