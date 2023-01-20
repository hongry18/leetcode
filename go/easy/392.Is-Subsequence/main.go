package main

func isSubsequence(s string, t string) bool {
	var i int

	for j := 0; j < len(t); j++ {
		if i < len(s) && s[i] == t[j] {
			i++
		}
	}

	return i == len(s)
}
