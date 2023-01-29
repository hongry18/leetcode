package main

func findWords(words []string) []string {
	var chars = []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	var res []string

	for i := range words {
		j := 0
		m := true
		for ; j < 3; j++ {
			if contains(chars[j], words[i][0]) {
				break
			}
		}
		for k := 1; k < len(words[i]); k++ {
			if !contains(chars[j], words[i][k]) {
				m = false
				break
			}
		}
		if !m {
			continue
		}
		res = append(res, words[i])
	}

	return res
}

func contains(s string, b byte) bool {
	if b >= 'A' && b <= 'Z' {
		b = b - 'A' + 'a'
	}
	for i := range s {
		if s[i] == b {
			return true
		}
	}
	return false
}
