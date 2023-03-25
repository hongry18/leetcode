package main

func isMonotonic(nums []int) bool {

	cmp := 0
	for i := 1; i < len(nums); i++ {
		t := nums[i-1] - nums[i]

		if cmp == 1 && t < 0 {
			return false
		} else if cmp == -1 && t > 0 {
			return false
		}

		if t > 0 {
			cmp = 1
		} else if t < 0 {
			cmp = -1
		}
	}
	return true
}
