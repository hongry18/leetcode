package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	// fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}

func sortByBits(arr []int) []int {
	var m = map[int][]int{}
	for _, n := range arr {
		b := bits.OnesCount(uint(n))
		if _, ok := m[b]; ok {
			m[b] = append(m[b], n)
		} else {
			m[b] = []int{n}
		}
	}

	var keys []int
	for key := range m {
		keys = append(keys, key)
		sort.Ints(m[key])
	}
	sort.Ints(keys)

	var res []int
	for _, key := range keys {
		for _, val := range m[key] {
			res = append(res, val)
		}
	}
	return res
}
