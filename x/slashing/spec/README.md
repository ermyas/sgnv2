<!--
order: 20
title: Slashing Overview
parent:
  title: "slashing"
-->

# `x/slashing`

## Abstract

This section specifies the slashing module of the Celer SGN network.

The slashing module enables Celer SGN network to disincentivize any attributable action
by a protocol-recognized actor with value at stake by penalizing them ("slashing").

Penalties may include, but are not limited to:

- Burning some amount of their stake
- Removing their ability to vote on future blocks for a period of time.

## Contents

1. **[Concepts](01_concepts.md)**
   - [States](01_concepts.md#states)
   - [Infraction timelines](01_concepts.md#infraction-timelines)
2. **[State](02_state.md)**
   - [Signing info](02_state.md#signing-info)
3. **[Messages](03_messages.md)**
   - [SignSlash](03_messages.md#signslash)
4. **[Begin-Block](04_begin_block.md)**
   - [Liveness tracking](04_begin_block.md#liveness-tracking)
   - [Double sign tracking](04_begin_block.md#double-sign-tracking)
5. **[Events](05_events.md)**
   - [BeginBlocker](05_events.md#beginblocker)
   - [Handlers](05_events.md#handlers)
6. **[Parameters](06_params.md)**
