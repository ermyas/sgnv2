## Staking

### Setup

Run `go run localnet.go -start` to set up the docker test environment with three sgn nodes.

### Add nodes to become validators

#### Node 0
```sh
sgnd ops init-validator --commission-rate 0.15 --min-self-delegation 1000 --keystore data/node0/keys/vethks0.json --home data/node0/sgnd
sgnd query staking validator 00078b31fa8b29a76bce074b5ea0d515a6aeaee7 --home data/node0/sgnd
```

#### Node 1
```sh
sgnd ops init-validator --commission-rate 0.2 --min-self-delegation 1100 --keystore data/node1/keys/vethks1.json --home data/node1/sgnd
sgnd query staking validator 0015f5863ddc59ab6610d7b6d73b2eacd43e6b7e --home data/node1/sgnd
```

#### Node 2
```sh
sgnd ops init-validator --commission-rate 0.2 --min-self-delegation 1200 --keystore data/node2/keys/vethks2.json --home data/node2/sgnd
sgnd query staking validator 00290a43e5b2b151d530845b2d5a818240bc7c70 --home data/node2/sgnd
```

### Query all validators

```sh
sgnd query staking validators --home data/node0/sgnd
sgnd query tendermint-validator-set --home data/node0/sgnd
```