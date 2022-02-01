<!--
order: 0
title: Distribution Overview
parent:
  title: "distribution"
-->

# `distribution`

## Overview

The distribution module implements a fee distribution mechanism using the efficient F1 Fee Distribution Algorithm. All staking rewards and fees charged by the SGN are first handed out to each validator, then after reserving the validator commission, distributed to the delegators according to their stake on the validator. The module also implements the flow for delegators and validators to claim and withdraw the fees on the EVM chains.

## Contents

1. **[Concepts](01_concepts.md)**
    - [Reference Counting in F1 Fee Distribution](01_concepts.md#reference-counting-in-f1-fee-distribution)
2. **[State](02_state.md)**
3. **[Begin Block](03_begin_block.md)**
4. **[Messages](04_messages.md)**
    - [MsgSetWithdrawAddress](04_messages.md#msgsetwithdrawaddress)
    - [MsgWithdrawDelegatorReward](04_messages.md#msgwithdrawdelegatorreward)
    - [MsgWithdrawValidatorCommission](04_messages.md#msgwithdrawvalidatorcommission)
    - [MsgFundCommunityPool](04_messages.md#msgfundcommunitypool)
    - [MsgClaimAllStakingReward](04_messages.md#msgclaimallstakingreward)
    - [MsgSignStakingReward](04_messages.md#msgsignstakingreward)
    - [Common Operations](04_messages.md#common-operations)
5. **[Hooks](05_hooks.md)**
    - [Delegation created or modified](05_hooks.md#delegation-created-or-modified)
    - [Validator created](05_hooks.md#validator-created)
    - [Validator removed](05_hooks.md#validator-removed)
6. **[Events](06_events.md)**
    - [BeginBlocker](06_events.md#beginblocker)
    - [Handlers](06_events.md#handlers)
    - [Bridge Fee Claims](06_events.md#bridge-fee-claims)
7. **[Parameters](07_params.md)**
