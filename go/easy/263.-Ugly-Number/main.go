package main

func isUgly(n int) bool {
	if n < 1 {
		return false
	}

	if n == 1 {
		return true
	}

	for n > 1 {
		switch {
		case n%2 == 0:
			n /= 2
		case n%3 == 0:
			n /= 3
		case n%5 == 0:
			n /= 5
		default:
			return false
		}
	}

	return true
}

func min(a, b, c int) int {
	var min int
	if a < b {
		min = a
	} else {
		min = b
	}

	if min > c {
		min = c
	}

	return min
}
