package main

import "math"

func thirdMax(nums []int) int {
	var t1, t2, t3 int = math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if t1 == v || t2 == v || t3 == v {
			continue
		}

		if v > t1 {
			t1, t2, t3 = v, t1, t2
		} else if v > t2 {
			t2, t3 = v, t2
		} else if v > t3 {
			t3 = v
		}
	}

	if t3 == math.MinInt64 || t2 == math.MinInt64 {
		return t1
	}

	return t3
}
