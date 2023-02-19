package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLHS([]int{1, 2, 3, 4}))
}

func findLHSOptimize(nums []int) int {
	var m = make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var res int
	for k, v := range m {
		if i, ok := m[k+1]; ok {
			if i+v > res {
				res = i + v
			}
		}
	}
	return res
}

func findLHS(nums []int) int {
	var m = make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var ar []int
	for k := range m {
		ar = append(ar, k)
	}
	sort.Ints(ar)
	var res int
	for i := 0; i < len(ar)-1; i++ {
		if ar[i+1]-ar[i] == 1 {
			t := m[ar[i+1]] + m[ar[i]]
			if t > res {
				res = t
			}
		}
	}
	return res
}
