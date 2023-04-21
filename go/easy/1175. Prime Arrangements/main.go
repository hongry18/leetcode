package main

import "fmt"

func main() {
	fmt.Println(numPrimeArrangements(5))
}

func numPrimeArrangements(n int) int {
	var ar = make([]int, n)
	for i := 0; i < n; i++ {
		ar[i] = i + 1
	}

	permu(ar, 0, 5, 5, 2)
	return 0
}

func permu(ar []int, d, n, r, t int) {
	if d == r {
		if ar[t-1] == t {

			for i := 0; i < r; i++ {
				fmt.Printf("%d", ar[i])
			}
			fmt.Println()
		}
		return
	}

	for i := d; i < n; i++ {
		ar[d], ar[i] = ar[i], ar[d]
		permu(ar, d+1, n, r, t)
		ar[d], ar[i] = ar[i], ar[d]
	}
}
