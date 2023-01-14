package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := map[byte]int{}
	tm := map[byte]int{}

	for i := 0; i < len(s); i++ {
		if _, ok := sm[s[i]]; !ok {
			sm[s[i]] = i
		}

		if _, ok := tm[t[i]]; !ok {
			tm[t[i]] = i
		}

		if sm[s[i]] != tm[t[i]] {
			return false
		}
	}

	return true
}
