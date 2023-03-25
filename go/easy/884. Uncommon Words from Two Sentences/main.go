package main

import (
	"strings"
)

func main() {

	uncommonFromSentences("this apple is sweet", "sour")
}

func uncommonFromSentences(s1 string, s2 string) []string {
	m := map[string]int{}

	for _, w := range strings.Split(strings.Join([]string{s1, s2}, " "), " ") {
		m[w]++
	}

	res := []string{}
	for i, v := range m {
		if v == 1 {
			res = append(res, i)
		}
	}

	return res
}
