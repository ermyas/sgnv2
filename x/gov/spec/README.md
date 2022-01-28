<!--
order: 0
title: Gov Overview
parent:
  title: "gov"
-->

# `gov`

## Abstract

This paper specifies the Governance module of the Celer SGN network.

In Celer SGN network, bonded validators can vote
on proposals on a 1 staking token 1 vote basis. Next is a list of features the module
currently supports:

- **Proposal submission:** Users can submit proposals with a deposit. Once the
minimum deposit is reached, proposal enters voting period
- **Vote:** Participants can vote on proposals that reached MinDeposit

## Contents

1. **[Concepts](01_concepts.md)**
    - [Proposal submission](01_concepts.md#proposal-submission)
    - [Vote](01_concepts.md#vote)
2. **[State](02_state.md)**
    - [Parameters and base types](02_state.md#parameters-and-base-types)
    - [Proposals](02_state.md#proposals)
    - [Stores](02_state.md#stores)
    - [Proposal Processing Queue](02_state.md#proposal-processing-queue)
3. **[Messages](03_messages.md)**
    - [Proposal Submission](03_messages.md#proposal-submission)
    - [Deposit](03_messages.md#deposit)
    - [Vote](03_messages.md#vote)
4. **[Events](04_events.md)**
    - [EndBlocker](04_events.md#endblocker)
    - [Handlers](04_events.md#handlers)
5. **[Parameters](05_params.md)**
