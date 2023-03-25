package main

func main() {
	sortArrayByParity([]int{3, 1, 2, 4})
}

func sortArrayByParity(nums []int) []int {
	j := 0
	for i, v := range nums {
		if v&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
