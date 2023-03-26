package main

import "fmt"

func main() {
	sortArrayByParityII([]int{4, 7, 5, 2})
	sortArrayByParityII([]int{4, 2, 5, 7})
	sortArrayByParityII([]int{2, 3})
	sortArrayByParityII([]int{648, 831, 560, 986, 192, 424, 997, 829, 897, 843})
	sortArrayByParityII([]int{8, 1, 60, 86, 92, 24, 7, 9, 7, 3})
}

func sortArrayByParityII(nums []int) []int {
	for i, num := range nums {
		if (i&1 == 0 && num&1 == 0) || (i&1 == 1 && num&1 == 1) {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if (i&1 == 0 && nums[j]&1 == 0) || (i&1 == 1 && nums[j]&1 == 1) {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	fmt.Println(nums)
	return nums
}
