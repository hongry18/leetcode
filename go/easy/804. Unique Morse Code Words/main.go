package main

import "bytes"

func uniqueMorseRepresentations(words []string) int {
	var m = map[string]int{}

	var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	for i := 0; i < len(words); i++ {
		var buf bytes.Buffer
		for j := 0; j < len(words[i]); j++ {
			buf.WriteString(morse[words[i][j]-'a'])
		}

		m[buf.String()]++
	}

	return len(m)
}
