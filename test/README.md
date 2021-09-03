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
- geth log path: docker-volumes/geth-env/geth.log
- sgn nodeN log path: docker-volumes/nodeN/sgnd/sgnd.log