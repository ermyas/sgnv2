package utils

import "sort"

//SelectMid
func SelectMid(s []float64) float64 {
	if s == nil || len(s) == 0 {
		return 0
	}
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s[len(s)/2]
}
