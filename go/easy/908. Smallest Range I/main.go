package main

func smallestRangeI(nums []int, k int) int {
	min, max := 10000, 0
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	t := (max - k) - (min + k)
	if t < 0 {
		return 0
	}
	return t
}
