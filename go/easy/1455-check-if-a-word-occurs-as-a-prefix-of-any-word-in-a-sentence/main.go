package main

import (
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	s1 := len(searchWord)
	ar := strings.Split(sentence, " ")
	for i, s := range ar {
		s2 := len(s)
		if s1 <= s2 && s[:s1] == searchWord {
			return i + 1
		}
	}
	return -1
}
