package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
}

func maximumProduct(nums []int) int {
	var sum int = 1
	if len(nums) < 4 {
		for _, v := range nums {
			sum *= v
		}
		return sum
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	a := nums[0] * nums[1] * nums[len(nums)-1]
	b := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	if a > b {
		return a
	}
	return b
}
