package main

func countBits(n int) []int {
	res := make([]int, n+1)

	offset := 1
	for i := 1; i <= n; i++ {
		if offset*2 == i {
			offset = i
		}
		res[i] = 1 + res[i-offset]
	}

	return res
}
