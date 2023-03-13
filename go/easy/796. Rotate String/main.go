package main

import (
	"strings"
)

func main() {
	rotateString("abcde", "cdeab")
	rotateString("abcde", "abced")
	rotateString("clrwmpkwru", "wmpkwruclr")
	rotateString("bbbacddceeb", "ceebbbbacdd")
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}
