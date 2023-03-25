package main

func main() {
	reverseOnlyLetters("a-bC-dEf-ghIj")
}
func reverseOnlyLetters(s string) string {
	var i, j = 0, len(s) - 1
	var b = []byte(s)

	for i <= j {
		if !isAlpha(s[i]) {
			i++
			continue
		}

		if !isAlpha(s[j]) {
			j--
			continue
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}

func isAlpha(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}
