package main

func summaryRanges(nums []int) []string {
	var res []string

	v := nums[0]
	for i := 1; i < len(nums); i++ {
		v = nums[i]
	}

	return res
}
