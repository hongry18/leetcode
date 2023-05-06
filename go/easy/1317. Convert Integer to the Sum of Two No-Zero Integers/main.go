package main

import "fmt"

func main() {
	// for i := 2; i < 10000; i++ {
	// 	fmt.Println(getNoZeroIntegers(i))
	// }
	// getNoZeroIntegers(9904)
	// getNoZeroIntegers(1010)
	// fmt.Println(getNoZeroIntegers(9999))
	fmt.Println(getNoZeroIntegers(9100))
}

func getNoZeroIntegers(n int) []int {
	var i = 1
	for {
		if !hasZero(i) && !hasZero(n-i) {
			return []int{i, n - i}
		}
		i++
	}
}

func hasZero(a int) bool {
	for a > 0 {
		if a%10 == 0 {
			return true
		}
		a /= 10
	}

	return false
}
