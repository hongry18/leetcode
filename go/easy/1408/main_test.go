package main

import (
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Log(stringMatching([]string{"mass", "as", "hero", "superhero"}))
}

func stringMatching(words []string) []string {
	var res []string
	for i1, s1 := range words {
		for i2, s2 := range words {
			if i1 == i2 {
				continue
			}

			if len(s1) > len(s2) {
				continue
			}

			if s1 == s2 || strings.Contains(s2, s1) {
				res = append(res, s1)
				break
			}
		}
	}
	return res
}
