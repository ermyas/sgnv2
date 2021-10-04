
## Local Multi-Node Manual Testing

Follow instructions to start a local testnet with multiple validators and test different scenarios on your local machine.

### Test Cases

- [Staking](./docs/staking.md)
- [Governance](./docs/governance.md)
- [cBridge](./docs/cbridge.md)

## Local Single-Node Manual Testing

#### Reset, Build, Start

Run following commands in sgn repo root folder
```sh
make update-test-data
WITH_CLEVELDB=yes make install
make localnet-start-geth
sgnd ops deploy
echo 12341234 | sgnd start 2> ~/.sgnd/app.log
tail -f ~/.sgnd/app.log # optional, in another terminal
```

#### Test Staking

Note: wait for a few seconds between steps
```sh
sgnd ops init-validator --commission-rate 0.15 --min-self-delegation 1000 --keystore ~/.sgnd/keys/vethks0.json
sgnd query staking validator 00078b31fa8b29a76bce074b5ea0d515a6aeaee7

echo 12341234 | sgnd tx staking edit-description --website "celer.network"
sgnd query staking validator 00078b31fa8b29a76bce074b5ea0d515a6aeaee7

sgnd ops delegate --validator 00078b31fa8b29a76bce074b5ea0d515a6aeaee7 --amount 100 --keystore ~/.sgnd/keys/dethks0.json
sgnd query staking validator 00078b31fa8b29a76bce074b5ea0d515a6aeaee7
sgnd query staking delegations 00078b31fa8b29a76bce074b5ea0d515a6aeaee7
sgnd query staking delegation 00078b31fa8b29a76bce074b5ea0d515a6aeaee7 d0f2596d700c9bd4d605c938e586ec67b01c7364

echo 12341234 | sgnd tx staking set-transactors
sgnd query staking transactors 00078b31fa8b29a76bce074b5ea0d515a6aeaee7
```

#### Stop Geth After Test
```sh
docker-compose stop geth
```
