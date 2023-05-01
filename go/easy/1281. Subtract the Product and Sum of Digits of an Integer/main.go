package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	var s, m int = 0, 1
	for n > 0 {
		s += n % 10
		m *= n % 10
		n /= 10
	}

	return m - s
}
