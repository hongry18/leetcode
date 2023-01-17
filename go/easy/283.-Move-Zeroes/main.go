package main

func moveZeroes(nums []int) {
	var i int

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i] != 0 {
			i++
		}
	}
}
