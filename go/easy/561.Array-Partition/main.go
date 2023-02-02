package main

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	var res int
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}
