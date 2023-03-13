package main

import (
	"bytes"
)

func main() {
	mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})
}

func mostCommonWord(paragraph string, banned []string) string {
	banMap := map[string]bool{}
	for _, v := range banned {
		banMap[v] = true
	}

	m := map[string]int{}
	cnt := 0
	res := ""

	var buf bytes.Buffer
	for i := 0; i < len(paragraph); i++ {
		if paragraph[i] >= 'A' && paragraph[i] <= 'Z' {
			buf.WriteByte(paragraph[i] - 'A' + 'a')
			continue
		}

		if paragraph[i] >= 'a' && paragraph[i] <= 'z' {
			buf.WriteByte(paragraph[i])
			continue
		}

		if buf.Len() == 0 {
			continue
		}

		if !banMap[buf.String()] {
			m[buf.String()]++
			if m[buf.String()] > cnt {
				cnt = m[buf.String()]
				res = buf.String()
			}
		}
		buf.Reset()

	}

	if buf.Len() != 0 {
		if !banMap[buf.String()] {
			m[buf.String()]++
			if m[buf.String()] > cnt {
				cnt = m[buf.String()]
				res = buf.String()
			}
		}
	}

	return res
}
