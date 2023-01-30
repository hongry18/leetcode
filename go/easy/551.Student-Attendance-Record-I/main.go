package main

import "strings"

func checkRecord(s string) bool {
	return strings.Index(s, "LLL") == -1 && strings.Count(s, "A") < 2
}
