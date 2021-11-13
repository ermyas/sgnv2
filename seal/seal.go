package seal

import (
	"time"

	"github.com/celer-network/goutils/log"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	slashingtypes "github.com/celer-network/sgn-v2/x/slashing/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
)

func NewTransactorLog(sender string) *TransactorLog {
	msgtypes := make(map[string]uint32)
	return &TransactorLog{
		MsgType:        msgtypes,
		Sender:         sender,
		StartTimestamp: time.Now().UnixNano(),
	}
}

func CommitTransactorLog(entry *TransactorLog) {
	now := time.Now().UnixNano()
	entry.ExecutionTimeMs = ((float64)(now) - float64(entry.StartTimestamp)) / 1000000
	if len(entry.Error) > 0 {
		log.Errorln("TransactorLog:", entry)
	} else if len(entry.Warn) > 0 {
		log.Warnln("TransactorLog:", entry)
	} else {
		log.Infoln("TransactorLog:", entry)
	}
}

func NewMsgLog(module string) *MsgLog {
	now := time.Now().UnixNano()
	msgLog := &MsgLog{
		ExecutionTimeMs: (float64)(now),
	}
	switch module {
	case stakingtypes.ModuleName:
		msgLog.Staking = &Staking{}
	case slashingtypes.ModuleName:
		msgLog.Slash = &Slash{}
	case synctypes.ModuleName:
		msgLog.Sync = &Sync{}
	case govtypes.ModuleName:
		msgLog.Govern = &Govern{}
	}
	return msgLog
}

func CommitMsgLog(entry *MsgLog) {
	now := time.Now().UnixNano()
	entry.ExecutionTimeMs = ((float64)(now) - entry.ExecutionTimeMs) / 1000000
	if len(entry.Error) > 0 {
		log.Errorln("MsgLog:", entry)
	} else if len(entry.Warn) > 0 {
		log.Warnln("MsgLog:", entry)
	} else {
		log.Infoln("MsgLog:", entry)
	}
}
