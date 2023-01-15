package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ar := make([]int, 26)

	for _, v := range s {
		ar[v-'a']++
	}

	for _, v := range t {
		ar[v-'a']--
		if ar[v-'a'] < 0 {
			return false
		}
	}

	return true
}
