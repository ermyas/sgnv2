<!--
order: 3
-->

# Begin-Block

The block reward is paid at the beginning of each block.

## BlockProvision

Calculate the provisions generated for each block based on current annual provisions. The provisions are then minted by the `mint` module's `ModuleMinterAccount` and then transferred to the `auth`'s `FeeCollector` `ModuleAccount`.

[Code reference](https://github.com/celer-network/sgnv2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/x/mint/types/minter.go#L28)
