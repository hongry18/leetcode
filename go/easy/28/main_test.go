package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	// t.Log(strStr("sadbutsad", "sad"))
	// t.Log(strStr("leetcode", "leeto"))
	t.Log(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
