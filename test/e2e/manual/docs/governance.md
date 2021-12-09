## Governance

### Setup

Run `go run localnet.go -start -cbr -auto` to start testnet and auto config all nodes as validators.

### Param Change - Example: update block reward

1. Query current block mining reward and submit change proposal:

```sh
sgnd query staking params --home data/node0/sgnd
sgnd tx gov submit-proposal param-change data/param_change_proposal.json --home data/node0/sgnd
sgnd query gov proposal 1 --home data/node0/sgnd
```

2. All nodes vote yes:

```sh
sgnd tx gov vote 1 yes --home data/node0/sgnd
sgnd tx gov vote 1 yes --home data/node1/sgnd
sgnd tx gov vote 1 yes --home data/node2/sgnd
sgnd tx gov vote 1 yes --home data/node3/sgnd
```

3. Query proposal status and updated block mining reward after voting timeout (2 mins):

```sh
sgnd query gov proposal 1 --home data/node0/sgnd
sgnd query staking params --home data/node0/sgnd
```

### Cbridge Cbr Config Update
note - the placeholder {proposal_id} must be replaced with the real proposal id, based on how much proposals you have submitted
1. Query current cbr config and submit change proposal:

```sh
sgnd query cbridge config --home data/node0/sgnd
sgnd tx gov submit-proposal cbridge-change data/cbridge_cbr_proposal.json --home data/node0/sgnd
sgnd query gov proposal {proposal_id} --home data/node0/sgnd
```

2. All nodes vote yes:

```sh
sgnd tx gov vote {proposal_id} yes --home data/node0/sgnd
sgnd tx gov vote {proposal_id} yes --home data/node1/sgnd
sgnd tx gov vote {proposal_id} yes --home data/node2/sgnd
sgnd tx gov vote {proposal_id} yes --home data/node3/sgnd
```

3. Query proposal status and updated cbr config after voting timeout (2 mins):

```sh
sgnd query gov proposal {proposal_id} --home data/node0/sgnd
sgnd query cbridge config --home data/node0/sgnd
```
