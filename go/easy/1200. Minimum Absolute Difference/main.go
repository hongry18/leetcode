package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	min := 100000
	for i := 1; i < len(arr); i++ {
		t := arr[i] - arr[i-1]
		if t < min {
			min = t
		}
	}

	var res [][]int
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
