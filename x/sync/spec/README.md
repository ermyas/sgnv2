<!--
order: 0
title: Sync Overview
parent:
  title: "sync"
-->

# Sync

## Overview

The sync module works together with relayer to bridge states from other blockchains to SGN chain. Any validator account can submit an update proposal after observing an external blockchain event (e.g., staking changes, cbridge relay request). After a `PendingUpdate` is proposed, a time window is opened to let all validators vote for the update. If more than 2/3 of the validator voting power have voted `yes`, the `PendingUpdate` would be applied, which will call other related modules' functions to apply the updates accordingly.

## Contents

1. **[State](01_state.md)**
2. **[Messages](02_messages.md)**
3. **[End-Block](03_end_block.md)**
4. **[Parameters](04_params.md)**