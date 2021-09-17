## Governance

### Setup

Run `go run localnet.go -start -auto` to start testnet and auto config all nodes as validators.

### Example: update block reward

1. Query current block mining reward and submit change proposal:

Append args `--home data/node0/sgnd` to the following commands.

```sh
sgnd query staking params
sgnd tx govern submit-proposal param-change data/param_change_proposal.json
sgnd query govern proposal 1
```

2. All nodes vote yes:

```sh
sgnd tx govern vote 1 yes --home data/node0/sgnd
sgnd tx govern vote 1 yes --home data/node1/sgnd
sgnd tx govern vote 1 yes --home data/node2/sgnd
```

3. Query proposal status and updated block mining reward after voting timeout (2 mins):

Append args `--home data/node0/sgnd` to the following commands.

```sh
sgnd query govern proposal 1
sgnd query staking params
```