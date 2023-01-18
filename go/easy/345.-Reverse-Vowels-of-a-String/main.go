package main

func reverseVowels(s string) string {
	b := []byte(s)
	var find bool
	for i, j := 0, len(b)-1; i < j; {
		if find {
			if match(b[j]) {
				b[i], b[j] = b[j], b[i]
				find = false
				i++
			}
			j--
		} else {
			if match(b[i]) {
				find = true
			} else {
				i++
			}
		}
	}
	return string(b)
}

func match(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
