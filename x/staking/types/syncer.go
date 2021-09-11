package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func NewSyncer(valIndex uint64, ethAddress string) *Syncer {
	return &Syncer{
		ValIndex:   valIndex,
		EthAddress: ethAddress,
	}
}

func MustMarshalSyncer(cdc codec.BinaryCodec, syncer *Syncer) []byte {
	return cdc.MustMarshal(syncer)
}

func MustUnmarshalSyncer(cdc codec.BinaryCodec, value []byte) Syncer {
	syncer, err := UnmarshalSyncer(cdc, value)
	if err != nil {
		panic(err)
	}

	return syncer
}

func UnmarshalSyncer(cdc codec.BinaryCodec, value []byte) (d Syncer, err error) {
	err = cdc.Unmarshal(value, &d)
	return d, err
}
