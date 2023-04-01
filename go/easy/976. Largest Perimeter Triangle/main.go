package main

import (
	"sort"
)

func main() {
	largestPerimeter([]int{2, 1, 2})
	largestPerimeter([]int{1, 2, 1, 10})
}

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1]+nums[i+2] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}
