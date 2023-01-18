package main

func reverseVowels(s string) string {
	b := []byte(s)

	for i, j := 0, len(b)-1; i < j; {
		if !match(b[i]) {
			i++
			continue
		}
		if !match(b[j]) {
			j--
			continue
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
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
