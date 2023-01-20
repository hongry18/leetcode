package main

func findTheDifference(s string, t string) byte {
	var ar = make([]int, 26)
	for i := range s {
		ar[s[i]-'a']++
	}

	for i := range t {
		ar[t[i]-'a']--
		if ar[t[i]-'a'] == -1 {
			return t[i]
		}
	}
	return 0
}
