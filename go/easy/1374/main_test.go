package main

import (
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Log(generateTheString(7))
	t.Log(generateTheString(4))
}

func generateTheString(n int) string {
	if n%2 == 1 {
		return strings.Repeat("a", n)
	}
	return strings.Repeat("a", n-1) + "b"
}
