<!--
order: 0
title: Mint Overview
parent:
  title: "mint"
-->

# `mint`

## Overview

This module controls the minting of CELR block rewards. It is a simplified version of the `x/mint` module in upstream Cosmos SDK without the logic of automatically calculating the target inflation rate.

## Contents

1. **[Concept](01_concepts.md)**
2. **[State](02_state.md)**
    - [Minter](02_state.md#minter)
    - [Params](02_state.md#params)
3. **[Begin-Block](03_begin_block.md)**
    - [BlockProvision](03_begin_block.md#blockprovision)
4. **[Parameters](04_params.md)**
5. **[Events](05_events.md)**
    - [BeginBlocker](05_events.md#beginblocker)
