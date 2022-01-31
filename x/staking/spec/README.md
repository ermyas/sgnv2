<!--
order: 0
title: Staking Overview
parent:
  title: "staking"
-->

# Staking

## Overview

The staking module manages the underlying [tendermint validators](https://docs.tendermint.com/master/nodes/validators.html) of the SGN chain by following the staking changes on the Ethereum [staking contract](https://github.com/celer-network/sgn-v2-contracts/blob/main/contracts/Staking.sol). Whenever staking status is updated on Ethereum, the [sync](../../sync) module would be informed and send instructions to update the stating module's validator and delegation states accordingly.

## Contents

1. **[State](01_state.md)**
    - [Validator](01_state.md#validator)
    - [Delegation](01_state.md#delegation)
    - [Syncer](01_state.md#syncer)
2. **[Messages](02_messages.md)**
    - [MsgSetTransactors](02_messages.md#msgsettransactors)
    - [MsgEditDescription](02_messages.md#msgeditdescription)
3. **[End-Block](03_end_block.md)**
4. **[Parameters](04_params.md)**