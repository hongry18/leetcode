package main

func missingNumber(nums []int) int {
	ar := make([]int, len(nums)+1)

	for _, v := range nums {
		ar[v]++
	}

	for i, v := range ar {
		if v == 0 {
			return i
		}
	}
	return 0
}
