package main

func findDisappearedNumbers(nums []int) []int {
	var ar = make([]int, len(nums))

	for _, v := range nums {
		ar[v-1]++
	}

	var j int
	for i, v := range ar {
		if v == 0 {
			ar[j] = i + 1
			j++
		}
	}

	return ar[:j]
}
