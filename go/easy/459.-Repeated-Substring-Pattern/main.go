package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	next := make([]int, n)

	for i, j := 1, 0; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	fmt.Println(next)
	return next[n-1] != 0 && n%(n-next[n-1]) == 0
}
