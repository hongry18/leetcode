package main

import "fmt"

func main() {
	fmt.Println(sumZero(1))
	fmt.Println(sumZero(2))
	fmt.Println(sumZero(3))
	fmt.Println(sumZero(4))
	fmt.Println(sumZero(5))
	fmt.Println(sumZero(6))
	fmt.Println(sumZero(7))
}

func sumZero(n int) []int {
	var res = make([]int, n)
	var m = n / 2
	if n%2 == 0 {
		res[n-1] = 0
	}

	for i := 1; i <= m; i++ {
		res[i-1] = -i
		res[i-1+m] = i
	}
	return res
}
