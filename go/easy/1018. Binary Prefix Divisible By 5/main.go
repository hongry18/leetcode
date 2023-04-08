package main

func prefixesDivBy5(nums []int) []bool {
	var n int
	var res []bool
	for _, num := range nums {
		n = (n*2 + num) % 5
		res = append(res, n == 0)
	}
	return res
}
