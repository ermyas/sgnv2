<!--
order: 6
-->

# Client

## CLI

A user can query and interact with the `mint` module using the CLI.

### Query

The `query` commands allow users to query `mint` state.

```sh
sgnd query mint --help
```

#### annual-provisions

The `annual-provisions` command allow users to query the current minting annual provisions value

```sh
sgnd query mint annual-provisions [flags]
```

Example:

```sh
sgnd query mint annual-provisions
```

Example Output:

```
22268504368893.612100895088410693
```

#### params

The `params` command allow users to query the current minting parameters

```sh
sgnd query mint params [flags]
```

Example:

```yaml
blocks_per_year: "6311520"
mint_denom: CELR/reward
```

## gRPC

A user can query the `mint` module using gRPC endpoints.

### AnnualProvisions

The `AnnualProvisions` endpoint allow users to query the current minting annual provisions value

```
/sgn.mint.v1.Query/AnnualProvisions
```

Example:

```sh
grpcurl -plaintext localhost:9090 sgn.mint.v1.Query/AnnualProvisions
```

Example Output:

```
{
  "annualProvisions": "1432452520532626265712995618"
}
```

### Params

The `Params` endpoint allow users to query the current minting parameters

```
/sgn.mint.v1.Query/Params
```

Example:

```sh
grpcurl -plaintext localhost:9090 sgn.mint.v1.Query/Params
```

Example Output:

```json
{
  "params": {
    "mintDenom": "CELR/reward",
    "blocksPerYear": "6311520"
  }
}
```

## REST

A user can query the `mint` module using REST endpoints.

### annual-provisions

```
/sgn/mint/v1/annual_provisions
```

Example:

```sh
curl "localhost:1317/sgn/mint/v1/annual_provisions"
```

Example Output:

```json
{
  "annualProvisions": "1432452520532626265712995618"
}
```

### params

```
/sgn/mint/v1/params
```

Example:

```sh
curl "localhost:1317/sgn/mint/v1/params"
```

Example Output:

```json
{
  "params": {
    "mintDenom": "CELR/reward",
    "blocksPerYear": "6311520"
  }
}
```
