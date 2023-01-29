package main

func findWords(words []string) []string {
	var res []string
	for i := range words {
		ar := [3]int{}
		for j := range words[i] {
			b := words[i][j]
			if b >= 'A' && b <= 'Z' {
				b += 'a' - 'A'
			}

			switch b {
			case 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p':
				ar[0]++
			case 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l':
				ar[1]++
			case 'z', 'x', 'c', 'v', 'b', 'n', 'm':
				ar[2]++
			}
		}
		for j := range ar {
			if len(words[i]) == ar[j] {
				res = append(res, words[i])
			}
		}
	}

	return res
}
