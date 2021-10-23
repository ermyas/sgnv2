# Full Stack Local Testnet

### Start local testnet

Run the following command at [e2e/mnual](../) folder to to start a local testnet with full stack setup, which includes four validators and four delegators for staking tests, and two chains with mock USDT for cBridge tests. Gateway is started in node 0. Nodes IP and port mappings can be found at [docker-compose.yml](../../../../docker-compose.yml).

```sh
echo 12341234 | go run localnet.go -start -full
```

Validator logs are located at

```sh
docker-volumes/nodeN/sgnd/app.log
docker-volumes/nodeN/sgnd/tendermint.log
```

### Query states

```sh
sgnd query staking validators --home data/node0/sgnd
sgnd query staking validator 00078b31fa8b29a76bce074b5ea0d515a6aeaee7 --home data/node0/sgnd
sgnd query staking validator-delegations 003ea363bccfd7d14285a34a6b1deb862df0bc84 --home data/node0/sgnd
sgnd query staking delegator-delegations d0f2596d700c9bd4d605c938e586ec67b01c7364 --home data/node0/sgnd
sgnd query staking transactors 00078b31fa8b29a76bce074b5ea0d515a6aeaee7 --home data/node0/sgnd

sgnd query cbridge chain-signers 883 --home data/node0/sgnd
sgnd query cbridge chain-signers 884 --home data/node0/sgnd
```

### Rebuild local testnet while keeping all states

1. `go run localnet.go -stopall`: stop sgn nodes
2. `go run localnet.go -rebuild`: rebuild sgn images (with code updates)
3. `go run localnet.go -upall`: restart all sgn nodes

### Connect MetaMask

Add two Custom RPCs in MetaMask for staking and cBridge frontend tests:
1. Network Name: `TestChain1`, New RPC URL: `http://127.0.0.1:8545`, Chain ID: `883`
1. Network Name: `TestChain2`, New RPC URL: `http://127.0.0.1:8547`, Chain ID: `884`

### Send test tokens

Run following command to send Mock ETH, CELR, and USDT tokens to a test (e.g., MetaMask account) address.
```sh
go run localnet.go -cbr -fund <eth-addess>
```
The default [test addresses](../../../keys) are already funded.

### Tear down local testnet

```sh
go run localnet.go -down
```
