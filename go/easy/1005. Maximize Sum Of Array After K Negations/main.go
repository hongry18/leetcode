package main

import (
	"sort"
)

func main() {
	largestSumAfterKNegations([]int{4, 2, 3}, 1)
	largestSumAfterKNegations([]int{3, -1, 0, 2}, 3)
	largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2)
	largestSumAfterKNegations([]int{3, -1, 1, 1, 2}, 3)
}

func largestSumAfterKNegations(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var sum int
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	if nums[0] > 0 {
		if k%2 == 1 {
			nums[0] *= -1
		}
	} else if nums[0] < 0 {
		for i := 0; i < k; i++ {
			nums[0] *= -1
			if nums[0] > nums[1] {
				nums = append(nums, nums[0])
				nums = nums[1:]
			}
		}
	}

	for _, n := range nums {
		sum += n
	}
	return sum
}
