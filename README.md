# Celer State Guardian Network V2

Official Golang implementation of the SGN V2 node and client based on Cosmos SDK.

## Generating Go bindings for Protobufs

1. Install [Starport](https://docs.starport.network/guide/install.html)
2. From the project root directory, run:

```sh
starport generate proto-go
```

## Generating genesis.json for tests

The passphrase for the test keyring should be set in [sgn_template.toml](test/data/.sgnd/config/sgn_template.toml).

```sh
rm -rf ~/.sgnd
cp -a test/data/.sgnd ~/.sgnd
rm ~/.sgnd/config/genesis.json
rm ~/.sgnd/config/gentx/*.json
sgnd init node0 --chain-id sgn-localnet-1000
sgnd add-genesis-account $(sgnd keys show alice -a --keyring-backend file --keyring-dir ~/.sgnd) 100000000stake # NOTE: Somehow Cosmos SDK requires this to be a large amount
sgnd gentx alice 100000000stake --min-self-delegation 1000 --amount 100000000stake --identity 00078b31fa8b29a76bce074b5ea0d515a6aeaee7 --keyring-backend file --keyring-dir ~/.sgnd --chain-id sgn-localnet-1000
sgnd collect-gentxs
cp ~/.sgnd/config/genesis.json test/data/.sgnd/config/genesis.json
```
