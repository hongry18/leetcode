package main

func repeatedNTimes(nums []int) int {
	m := map[int]bool{}

	for i := range nums {
		if m[nums[i]] {
			return nums[i]
		}
		m[nums[i]] = true
	}

	return 0
}
