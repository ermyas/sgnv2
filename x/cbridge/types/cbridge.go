package types

import "fmt"

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

// basic check of config
func (c *CbrConfig) Validate() error {
	if c.LpFee > 100 {
		return fmt.Errorf("lp_fee %d > 100", c.LpFee)
	}
	// todo: make sure assets are multi-chain correctly?
	for _, chpair := range c.ChainPairs {
		if chpair.Chid1 > chpair.Chid2 {
			return fmt.Errorf("chid1 %d > chid2 %d", chpair.Chid1, chpair.Chid2)
		}
	}
	return nil
}
