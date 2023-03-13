package main

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = 10001
	}
	last := -1

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			last = i
		}

		if last != -1 {
			res[i] = min(res[i], abs(last-i))
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			last = i
		}

		if last != -1 {
			res[i] = min(res[i], abs(last-i))
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
