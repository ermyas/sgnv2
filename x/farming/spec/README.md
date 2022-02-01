<!--
order: 0
title: Farming Overview
parent:
  title: "farming"
-->

# `farming`

## Overview

The farming module implements a liquidity mining mechanism using the he mechanism is similar to the F1 Fee Distribution Algorithm in `x/distribution`, which is also similar to the one used in the popular Sushiswap MasterChef contract. Users stake LP tokens in each farming pool to earn multiple reward tokens that can be claimed and withdrawn.

## Contents

1. **[Concepts](01_concepts.md)**
    - [Reference Counting in F1 Fee Distribution](01_concepts.md#reference-counting-in-f1-fee-distribution)
2. **[State](02_state.md)**
3. **[Messages](03_messages.md)**
    - [MsgClaimRewards](03_messages.md#msgclaimrewards)
    - [MsgClaimAllRewards](03_messages.md#msgclaimallrewards)
    - [MsgSignRewards](03_messages.md#msgsignrewards)
    - [Common Operations](03_messages.md#common-operations)
4. **[Events](04_events.md)**
    - [Handlers](04_events.md#handlers)
5. **[Parameters](05_params.md)**
