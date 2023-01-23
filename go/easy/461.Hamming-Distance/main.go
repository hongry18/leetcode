package main

func hammingDistance(x int, y int) int {
	var cnt int
	x ^= y
	for x > 0 {
		if x&1 != 0 {
			cnt++
		}
		x >>= 1
	}
	return cnt
}
