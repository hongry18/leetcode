package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(0))
	fmt.Println(numberOfSteps(1))
}

func numberOfSteps(num int) int {
	var cnt int
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		cnt++
	}
	return cnt
}
