package main

func lengthOfLastWord(s string) int {
	var cnt int
	var prev byte

	for i := range s {
		if prev == ' ' && s[i] != ' ' {
			cnt = 0
		}

		prev = s[i]
		if s[i] != ' ' {
			cnt++
		}
	}

	return cnt
}
