<!--
order: 6
-->

# Parameters

The slashing module contains the following parameters:

| Key                     | Type           | Example                |
| ----------------------- | -------------- | ---------------------- |
| EnableSlash             | bool           | false                  |
| SignedBlocksWindow      | string (int)   | "100"                  |
| SlashTimeout            | string (s)     | "150000"               |
| SlashFactorDoubleSign   | string (int)   | "50000"                |
| SlashFactorDowntime     | string (int)   | "10000"                |
| SlashJailPeriod         | string (s)     | "0"                    |
| MinSignedPerWindow      | string (dec)   | "0.500000000000000000" |
| StakingContract         | object         | {"address":"0x1345c8a6b99536531F1fa3cfe37D8A5B7Fc859aA", "chain_id":5} |