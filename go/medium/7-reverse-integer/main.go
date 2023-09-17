package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	fmt.Println(int32(x))
	var isEven bool
	if x < 0 {
		isEven = true
		x *= -1
	}

	var result int
	for x > 0 {
		result = (result * 10) + x%10
		x /= 10
	}

	if isEven {
		result *= -1
	}

	if result >= math.MaxInt32 || result <= math.MinInt32 {
		return 0
	}

	return result
}
