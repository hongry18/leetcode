package main

func maxScore(s string) int {
	var zeros, ones, max int
	for i := range s {
		if s[i] == '1' {
			ones++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			ones--
		} else {
			zeros++
		}

		if max < ones+zeros {
			max = ones + zeros
		}
	}
	return max
}
