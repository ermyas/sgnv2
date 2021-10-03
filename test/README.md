# Local E2E Testing

Requirements for multi-node or manual testing: Install [docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/), start Docker daemon.

## Manual Testing

Follow instructions in [e2e/manual](./e2e/manual/README.md) for local manual testing.

## Automated Testing

### Single-Node Automated Testing

Run the following command from the sgn-v2 repo root folder

```sh
go test -failfast -v -timeout 30m github.com/celer-network/sgn-v2/test/e2e/singlenode
```

To run a single test (e.g., staking test), run following command in `test/e2e/singlenode` folder

```sh
go test -failfast -v -run ^TestStaking$
```

### Multi-Node Automated Testing

Run the following command from the sgn-v2 repo root folder

```sh
go test -failfast -v -timeout 30m github.com/celer-network/sgn-v2/test/e2e/multinode
```
Logs are located at

```sh
geth: docker-volumes/geth-env/geth.log
node(0-3): docker-volumes/nodeN/sgnd/app.log & docker-volumes/nodeN/sgnd/tendermint.log
```

## Generating genesis.json for tests

The passphrase for the test keyring should be set in [sgn.toml](data/.sgnd/config/sgn.toml).

From the project root directory, run:

```sh
WITH_CLEVELDB=yes make install # Make sure sgnd is updated
rm -rf ~/.sgnd
cp -a test/data/.sgnd ~/.sgnd
rm ~/.sgnd/config/genesis.json
sgnd init node0 --chain-id sgn-localnet-1000
sgnd add-genesis-account $(sgnd keys show alice -a --keyring-backend file --keyring-dir ~/.sgnd) 100stake
sgnd add-genesis-validator alice 1000000000000 # passphrase: 12341234
cp ~/.sgnd/config/genesis.json test/data/.sgnd/config/genesis.json
cp ~/.sgnd/config/genesis.json test/multi-node-data/node0/sgnd/config/genesis.json
cp ~/.sgnd/config/genesis.json test/multi-node-data/node1/sgnd/config/genesis.json
cp ~/.sgnd/config/genesis.json test/multi-node-data/node2/sgnd/config/genesis.json
cp ~/.sgnd/config/genesis.json test/multi-node-data/node3/sgnd/config/genesis.json
```
