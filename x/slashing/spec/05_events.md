<!--
order: 5
-->

# Tags

The slashing module emits the following events/tags if slashing is enabled:

## BeginBlocker: HandleValidatorSignature & HandleDoubleSign

| Type  | Attribute Key | Attribute Value             |
| ----- | ------------- | --------------------------- |
| slash | nonce         | {slash.SlashOnChain.Nonce}  |
| slash | reason        | {slashReason}               |
