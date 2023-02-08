package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(findMaxAverage([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4))
	fmt.Println(findMaxAverage([]int{-1}, 1))
}
func findMaxAverage(nums []int, k int) float64 {
	ans := math.MinInt
	s := 0
	for i, v := range nums {
		s += v
		if i >= k-1 {
			if i != k-1 {
				s = s - nums[i-k]
			}
			if ans < s {
				ans = s
			}
		}
	}
	return float64(ans) / float64(k)
}
