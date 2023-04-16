package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(relativeSortArray([]int{}, []int{}))
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var res []int

	var ar = make([]int, 1001)
	var tmp []int

	for _, n := range arr2 {
		ar[n]++
	}

	for _, n := range arr1 {
		if ar[n] == 0 {
			tmp = append(tmp, n)
		} else {
			ar[n]++
		}
	}

	for _, n := range arr2 {
		for i := 0; i < ar[n]-1; i++ {
			res = append(res, n)
		}
	}

	sort.Ints(tmp)
	res = append(res, tmp...)

	return res
}
