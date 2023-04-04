## Live Upgrade

Cosmos SDK [upgrade reference](https://github.com/cosmos/cosmos-sdk/blob/v0.44.3/docs/core/upgrade.md).

### Setup

Run `echo 12341234 | go run localnet.go -start -cbr -auto` to start testnet and auto config all nodes as validators.

### Propose and approve upgrade proposal

1. Query current block height and submit upgrade proposal:

```sh
sgnd query block
echo 12341234 | sgnd tx gov submit-proposal software-upgrade test --title "upgrade test" --description "upgrade test" --upgrade-height [sidechain block height after more than 2 mins] --home data/node0/sgnd
```

2. All nodes vote yes:

```sh
echo 12341234 | sgnd tx gov vote 1 yes --home data/node0/sgnd
echo 12341234 | sgnd tx gov vote 1 yes --home data/node1/sgnd
echo 12341234 | sgnd tx gov vote 1 yes --home data/node2/sgnd
echo 12341234 | sgnd tx gov vote 1 yes --home data/node3/sgnd
```

3. Query proposal status after voting timeout (2 mins):

```sh
sgnd query gov proposal 1 --home data/node0/sgnd
```

4. Wait for the sidechain to halt at the proposed block height.

### Upgrade to the new version

1. Stop the containers:

```sh
go run localnet.go -stopall
```

2. Switch to the new code, add upgrade handler and migration code. [Example: add new staking param](https://github.com/celer-network/sgnv2/commit/1fadc4e3f2c21b449222c24174dd13963ba805ee)

3. Rebuild container images and restart:

```sh
go run localnet.go -rebuild
go run localnet.go -upall
```

4. Confirm upgrade is successful (following example above:)
```sh
sgnd query staking params --home data/node0/sgnd
```