package main

func maxProduct(nums []int) int {
	var a, b int

	for _, v := range nums {
		if v > a {
			a, b = v, a
		} else if v > b {
			b = v
		}
	}

	return (a - 1) * (b - 1)
}
