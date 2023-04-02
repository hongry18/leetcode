package main

func main() {
	commonChars([]string{"bella", "label", "roller"})
}

func commonChars(words []string) []string {
	m := map[rune]int{}

	for i, word := range words {
		t := map[rune]int{}
		for _, r := range word {
			if i == 0 {
				m[r]++
			} else {
				t[r]++
			}
		}

		if i == 0 {
			continue
		}

		for r := range m {
			if v, ok := t[r]; ok {
				if v < m[r] {
					m[r] = v
				}
			} else {
				delete(m, r)
			}
		}
	}

	var res []string
	for i := range m {
		for j := 0; j < m[i]; j++ {
			res = append(res, string(i))
		}
	}

	return res
}
