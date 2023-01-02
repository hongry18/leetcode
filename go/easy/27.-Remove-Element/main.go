package mian

func removeElement(nums []int, val int) int {
	var i int
	for j := range nums {
		if nums[j] == val {
			continue
		}
		nums[i] = nums[j]
		i++
	}

	return i
}
