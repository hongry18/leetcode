package main

import (
	"testing"
)

func TestXxx(t *testing.T) {
	t.Log(minStartValue([]int{-3, 2, -3, 4, 2}))
}

func minStartValue(nums []int) int {
	var res, min int = 0, 100
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	if min >= 0 {
		return 1
	}

	res = min*-1 + 1
	for {
		tmp := res
		for _, n := range nums {
			tmp += n
			if tmp < 1 {
				break
			}
		}
		if tmp > 0 {
			break
		}
		res++
	}
	return res
}
