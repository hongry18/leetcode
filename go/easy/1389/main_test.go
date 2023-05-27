package main

import (
	"testing"
)

func TestXxx(t *testing.T) {
	t.Log(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	t.Log(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
}

func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i := 0; i < len(index); i++ {
		res = append(res, nums[i])
		if i == index[i] {
			continue
		}
		copy(res[index[i]+1:], res[index[i]:])
		res[index[i]] = nums[i]
	}
	return res
}
