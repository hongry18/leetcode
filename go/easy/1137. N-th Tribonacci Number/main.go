package main

import "fmt"

func main() {
	fmt.Println(tribonacci(3))
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(5))
	fmt.Println(tribonacci(6))
	fmt.Println(tribonacci(27))
	fmt.Println(tribonacci(37))
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	var ar = make([]int, n)
	ar[0] = 0
	ar[1] = 1
	ar[2] = 1
	for i := 3; i < n; i++ {
		ar[i] = ar[i-1] + ar[i-2] + ar[i-3]
	}
	return ar[n-1] + ar[n-2] + ar[n-3]
}
