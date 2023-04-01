package main

func main() {
	// isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")
	isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")
	// isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz")
}

func isAlienSorted(words []string, order string) bool {
	m := map[byte]int{}
	for i := range order {
		m[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if !isOrder(words[i], words[i+1], m) {
			return false
		}
	}
	return true
}

func isOrder(a, b string, m map[byte]int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return m[a[i]] < m[b[i]]
		}
	}
	return len(a) <= len(b)
}
