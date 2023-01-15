package main

import "fmt"

func summaryRanges(nums []int) []string {
	var res []string

	size := len(nums)
	if size < 1 {
		return res
	}

	if size == 1 {
		res = append(res, fmt.Sprintf("%d", nums[0]))
		return res
	}

	s := nums[0]
	for i := 0; i < size-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			if nums[i] == s {
				res = append(res, fmt.Sprintf("%d", nums[i]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", s, nums[i]))
			}

			s = nums[i+1]
		}
	}

	if s == nums[size-1] {
		res = append(res, fmt.Sprintf("%d", nums[size-1]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", s, nums[size-1]))
	}

	return res
}
