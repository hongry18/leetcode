package main

func hammingDistance(x int, y int) int {
	var cnt int
	if x < y {
		x, y = y, x
	}

	for x > 0 {
		if x&1 != y&1 {
			cnt++
		}
		x >>= 1
		y >>= 1
	}
	return cnt
}
