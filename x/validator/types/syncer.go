package types

func NewSyncer(validatorIdx uint64, sgnAddress string) Syncer {
	return Syncer{
		ValIndex:   validatorIdx,
		SgnAddress: sgnAddress,
	}
}
