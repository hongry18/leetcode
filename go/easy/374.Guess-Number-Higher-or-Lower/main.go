package main

func guess(num int) int {
	p := 2
	if num > p {
		return -1
	} else if num < p {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	var l, r = 1, n

	for l <= r {
		m := l + (r-l)/2
		i := guess(m)

		if i == 0 {
			return m
		} else if i == -1 {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
