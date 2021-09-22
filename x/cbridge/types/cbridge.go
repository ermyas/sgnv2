package types

func (m *XferRelay) GetSortedSigsBytes() [][]byte {
	if m != nil {
		sigs := make([][]byte, len(m.SortedSigs))
		for i := range m.SortedSigs {
			sigs = append(sigs, m.SortedSigs[i].Sig)
		}
		return sigs
	}
	return nil
}
