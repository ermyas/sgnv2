package types

// check if chid is defined in NonEvmChainID enum
// panic if chid is 0
func IsNonEvm(chid uint64) bool {
	if chid == 0 {
		panic("IsNonEvm chid is 0")
	}
	_, ok := NonEvmChainID_name[int32(chid)]
	return ok
}

func IsFlowChain(chid uint64) bool {
	switch chid {
	case uint64(NonEvmChainID_FLOW_MAINNET), uint64(NonEvmChainID_FLOW_TEST), uint64(NonEvmChainID_FLOW_EMULATOR):
		return true
	default:
		return false
	}
}
