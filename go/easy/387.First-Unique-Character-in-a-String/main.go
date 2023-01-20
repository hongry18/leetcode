package main

func firstUniqChar(s string) int {
	var ar = make([]int, 26)
	for i := range s {
		ar[s[i]-'a']++
	}

	for i := range s {
		if ar[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
