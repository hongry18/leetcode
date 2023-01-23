package main

func findDisappearedNumbers(nums []int) []int {
	var ar = make([]int, len(nums))
	var res []int

	for _, v := range nums {
		ar[v-1]++
	}

	for i, v := range ar {
		if v == 0 {
			res = append(res, i+1)
		}
	}

	return res
}
