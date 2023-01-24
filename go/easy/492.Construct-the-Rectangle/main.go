package main

import (
	"math"
)

func constructRectangle(area int) []int {
	var i = int(math.Sqrt(float64(area)))
	for {
		if i < 2 || area%i == 0 {
			break
		}
		i--
	}
	return []int{area / i, i}
}
