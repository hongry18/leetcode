package main

import "fmt"

func main() {

	fmt.Println(oneCount(6))
	fmt.Println(oneCount(7))
	fmt.Println(oneCount(8))
	fmt.Println(oneCount(10))

}

func countPrimeSetBits(left int, right int) int {
	var cnt int
	prime := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
	}

	for i := left; i <= right; i++ {
		if prime[oneCount(i)] {
			cnt++
		}
	}
	return cnt
}

func oneCount(i int) int {
	var cnt int

	for i > 0 {
		if i&1 == 1 {
			cnt++
		}
		i >>= 1
	}
	return cnt
}
