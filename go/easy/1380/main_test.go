package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	t.Log(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
	t.Log(luckyNumbers([][]int{{7, 8}, {1, 2}}))
	t.Log(luckyNumbers([][]int{{3, 6}, {7, 1}, {5, 2}, {4, 8}}))
}

func luckyNumbers(matrix [][]int) []int {
	var ar []int
	for _, row := range matrix {
		var min int = 1000000
		for _, val := range row {
			if val < min {
				min = val
			}
		}
		ar = append(ar, min)
	}

	var max int
	for _, v := range ar {
		if v > max {
			max = v
		}
	}
	fmt.Println(ar)
	return []int{max}
}
