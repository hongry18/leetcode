package main

import (
	"strings"
)

func main() {
	removeOuterParentheses("(()())(())")
}

func removeOuterParentheses(s string) string {
	var j, k int
	var res []string

	for i, r := range s {
		if r == '(' {
			j++
		} else {
			j--
		}

		if j == 0 {
			res = append(res, s[k+1:i])
			k = i + 1
		}
	}
	return strings.Join(res, "")
}
