package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	var i, j int

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}

	return i == len(s)
}
