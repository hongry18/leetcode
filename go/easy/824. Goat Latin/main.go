package main

import (
	"strings"
)

func main() {
	toGoatLatin("I speak Goat Latin")
	// fmt.Println(strings.Repeat("a", 4))
}

func toGoatLatin(sentence string) string {
	ar := strings.Split(sentence, " ")
	cnt := 2
	for i := 0; i < len(ar); i++ {
		switch ar[i][0] {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			ar[i] = strings.Join([]string{ar[i], "m", strings.Repeat("a", cnt)}, "")
		default:
			ar[i] = strings.Join([]string{ar[i][1:], ar[i][:1], "m", strings.Repeat("a", cnt)}, "")
		}
		cnt++
	}

	return strings.Join(ar, " ")
}
