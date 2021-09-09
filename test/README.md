# Local E2E Testing

Requirements for multi-node or manual testing: Install [docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/), start Docker daemon.

## Automated Testing

### Single-Node Automated Testing

Run the following command from the sgn-v2 repo root folder

```sh
go test -failfast -v -timeout 30m github.com/celer-network/sgn-v2/test/e2e/singlenode
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

The passphrase for the test keyring should be set in [sgn_template.toml](data/.sgnd/config/sgn_template.toml).

From the project root directory, run:

```sh
rm -rf ~/.sgnd
cp -a test/data/.sgnd ~/.sgnd
rm ~/.sgnd/config/genesis.json
sgnd init node0 --chain-id sgn-localnet-1000
sgnd add-genesis-account $(sgnd keys show alice -a --keyring-backend file --keyring-dir ~/.sgnd) 100000000stake
sgnd add-genesis-validator alice 100000000
cp ~/.sgnd/config/genesis.json test/data/.sgnd/config/genesis.json
```
