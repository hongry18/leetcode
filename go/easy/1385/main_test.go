package main

import "testing"

func TestXxx(t *testing.T) {
	t.Log(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 18}, 2))
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int

	for _, v1 := range arr1 {
		var ok bool = true
		for _, v2 := range arr2 {
			if abs(v1-v2) <= d {
				ok = false
				break
			}
		}
		if ok {
			res++
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
