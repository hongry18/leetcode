package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	var ar = []int{0, 1}
	for i := 2; i <= n; i++ {
		ar = append(ar, ar[i-1]+ar[i-2])
	}
	return ar[n-1] + ar[n-2]
}
