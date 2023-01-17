package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}

	m1 := map[byte]int{}
	m2 := map[string]int{}
	for i := 0; i < len(strs); i++ {
		i1, ok1 := m1[pattern[i]]
		if !ok1 {
			m1[pattern[i]] = i + 1
		}

		i2, ok2 := m2[strs[i]]
		if !ok2 {
			m2[strs[i]] = i + 1
		}

		if i1 != i2 {
			return false
		}
	}
	return true
}
