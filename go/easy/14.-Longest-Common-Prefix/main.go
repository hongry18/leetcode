package main

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	min := math.MaxInt

	for i := range strs {
		size := len(strs[i])
		if size < min {
			min = size
		}
	}

	if min == 0 {
		return ""
	}

	for i := 0; i < min; i++ {
		var tmp byte
		for j := range strs {
			if j == 0 {
				tmp = strs[j][i]
				continue
			}

			if tmp != strs[j][i] {
				return strs[0][0:i]
			}

			tmp = strs[j][i]
		}
	}

	return strs[0][0:min]
}
