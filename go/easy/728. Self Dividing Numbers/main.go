package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {
	var ar []int
	for i := left; i <= right; i++ {
		if i%10 == 0 {
			continue
		}
		if isDiv(i) {
			ar = append(ar, i)
		}
	}
	return ar
}

func isDiv(a int) bool {
	b := a

	for b > 0 {
		if b%10 == 0 || a%(b%10) != 0 {
			return false
		}
		b /= 10
	}

	return true
}
