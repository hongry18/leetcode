package main

func main() {
	sortArrayByParityII([]int{4, 7, 5, 2})
}

func sortArrayByParityII(nums []int) []int {
	var i, j = 0, len(nums) - 1

	for i <= j {
		if (i&1 == 0 && nums[i]&1 == 0) || (i&1 == 1 && nums[i]&1 == 1) {
			i++
			continue
		}

		if (j&1 == 0 && nums[j]&1 == 0) || (j&1 == 1 && nums[j]&1 == 1) {
			j--
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
	return nums
}
