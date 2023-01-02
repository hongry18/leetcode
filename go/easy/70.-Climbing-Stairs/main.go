package main

func climbStairs(n int) int {
	if n < 4 {
		return n
	}

	ar := make([]int, n)
	ar[0] = 1
	ar[1] = 2

	for i := 2; i < n; i++ {
		ar[i] = ar[i-1] + ar[i-2]
	}

	return ar[n-1]
}
