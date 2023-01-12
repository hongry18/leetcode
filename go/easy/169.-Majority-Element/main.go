package main

func majorityElement(nums []int) int {
	his := map[int]int{}

	var max int
	var res int
	for _, v := range nums {
		his[v]++
		if his[v] > max {
			max = his[v]
			res = v
		}
	}

	return res
}
