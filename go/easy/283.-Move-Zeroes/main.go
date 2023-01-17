package main

func moveZeroes(nums []int) {
	var i int

	for _, n := range nums {
		if n != 0 {
			nums[i] = n
			i++
		}
	}

	for i < len(nums) {
		nums[i] = 0
		i++
	}
}
