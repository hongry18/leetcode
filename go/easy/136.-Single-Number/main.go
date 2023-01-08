package main

func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
	if len(nums) < 2 {
		return nums[0]
	}

	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for i, v := range m {
		if v == 1 {
			return i
		}
	}
	return 0
}
