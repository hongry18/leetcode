package main

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	lo, hi := 0, n

	for lo < hi {
		m := lo + (hi-lo)/2

		if isBadVersion(m) {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return lo
}
