<!--
order: 4
-->

# BeginBlock

## Liveness Tracking

At the beginning of each block, we update the `ValidatorSigningInfo` for each
validator and check if they've crossed below the liveness threshold over a
sliding window. This sliding window is defined by `SignedBlocksWindow` and the
index in this window is determined by `IndexOffset` found in the validator's
`ValidatorSigningInfo`. For each block processed, the `IndexOffset` is incremented
regardless if the validator signed or not. Once the index is determined, the
`MissedBlocksBitArray` and `MissedBlocksCounter` are updated accordingly.

Finally, in order to determine if a validator crosses below the liveness threshold,
we fetch the maximum number of blocks missed, `maxMissed`, which is
`SignedBlocksWindow - (MinSignedPerWindow * SignedBlocksWindow)` and the minimum
height at which we can determine liveness, `minHeight`. If the current block is
greater than `minHeight` and the validator's `MissedBlocksCounter` is greater than
`maxMissed`, they will be slashed by `MissingSignature`, will be jailed
for `DowntimeJailDuration`, and have the following values reset:
`MissedBlocksBitArray`, `MissedBlocksCounter`, and `IndexOffset`.

```go
	height := ctx.BlockHeight()
	consAddr := sdk.ConsAddress(addr)
	validator, found := k.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		log.Errorf("Cannot find validator %s", consAddr)
		return
	}

	signInfo, found := k.GetValidatorSigningInfo(ctx, consAddr)
	if !found {
		signInfo = slashingtypes.NewValidatorSigningInfo(
			consAddr,
			height,
			0,
			time.Unix(0, 0),
			false,
			0,
		)
	}

	// this is a relative index, so it counts blocks the validator *should* have signed
	// will use the 0-value default signing info if not present, except for start height
	signedBlocksWindow := k.SignedBlocksWindow(ctx)
	index := signInfo.IndexOffset % signedBlocksWindow
	signInfo.IndexOffset++

	// Update signed block bit array & counter
	// This counter just tracks the sum of the bit array
	// That way we avoid needing to read/write the whole array each time
	previous := k.GetValidatorMissedBlockBitArray(ctx, consAddr, index)
	missed := !signed
	switch {
	case !previous && missed:
		// Array value has changed from not missed to missed, increment counter
		k.SetValidatorMissedBlockBitArray(ctx, consAddr, index, true)
		signInfo.MissedBlocksCounter++
	case previous && !missed:
		// Array value has changed from missed to not missed, decrement counter
		k.SetValidatorMissedBlockBitArray(ctx, consAddr, index, false)
		signInfo.MissedBlocksCounter--
	default:
		// Array value at this index has not changed, no need to update counter
	}

	minHeight := signInfo.StartHeight + signedBlocksWindow
	maxMissed := signedBlocksWindow - k.MinSignedPerWindow(ctx).MulInt64(signedBlocksWindow).RoundInt64()

	// if we are past the minimum height and the validator has missed too many blocks, slash them
	if height > minHeight && signInfo.MissedBlocksCounter > maxMissed {
		// Downtime confirmed: slash the validator
		log.Infof("Validator %s %s past min height of %d and above max miss threshold of %d",
			validator.EthAddress, consAddr, minHeight, maxMissed)

		// We need to reset the counter & array so that the validator won't be immediately slashed for downtime upon rebonding.
		signInfo.MissedBlocksCounter = 0
		signInfo.IndexOffset = 0
		k.ClearValidatorMissedBlockBitArray(ctx, consAddr)
		k.Slash(ctx, types.AttributeValueMissingSignature, validator, k.SlashFactorDowntime(ctx), nil, blkTime) //collector and syncer reward will be done in next version
	}

	k.SetValidatorSigningInfo(ctx, signInfo)
}
```

## Double Sign Tracking
At the beginning of each block, we will also check if any validator signs two blocks at the same height. And if found any, they will be slashed by `DoubleSign`.

```go
	consAddr := sdk.ConsAddress(addr)
	validator, found := k.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		log.Errorf("Cannot find validator %s", consAddr)
		return
	}

	log.Infof("Confirmed double sign from %s %s", validator.EthAddress, consAddr)
	k.Slash(ctx, types.AttributeValueDoubleSign, validator, k.SlashFactorDoubleSign(ctx), nil, blkTime) //collector and syncer reward will be done in next version
```