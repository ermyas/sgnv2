# Pegbridge

## Overview

The pegbridge module implements cross-chain token transfer based on the mint / burn
mode as opposed to the liquidity pool in the cbridge module. To mint pegged tokens, the user deposits original tokens into a vault on the source chain and the SGN mints the same amount of pegged tokens (minus fees) on the destination chain specified. In the reverse transfer, the user burns the tokens on a pegged token bridge and the SGN unlocks the
same amount of original tokens (minus fees) on the original token vault.

## Contents

1. **[Concepts](01_concepts.md)**
   - [Original token](01_concepts.md#original-token)
   - [Pegged token](01_concepts.md#pegged-token)
   - [Original token vault](01_concepts.md#original-token-vault)
   - [Pegged token bridge](01_concepts.md#pegged-token-bridge)
   - [Workflow](01_concepts.md#workflow)
   - [Supply cap](01_concepts.md#supply-cap)
   - [Refund](01_concepts.md#refund)
2. **[State](02_state.md)**
3. **[Messages](03_messages.md)**
   - [MsgSignMint](03_messages.md#msgsignmint)
   - [MsgSignWithdraw](03_messages.md#msgsignwithdraw)
   - [MsgTriggerSignMint](03_messages.md#msgtriggersignmint)
   - [MsgTriggerSignWithdraw](03_messages.md#msgtriggersignwithdraw)
   - [MsgClaimFee](03_messages.md#msgclaimfee)
   - [MsgClaimRefund](03_messages.md#msgclaimrefund)
4. **[Events](04_events.md)**
   - [mint_to_sign](04_events.md#mint_to_sign)
   - [withdraw_to_sign](04_events.md#withdraw_to_sign)
5. **[Parameters](05_params.md)**
