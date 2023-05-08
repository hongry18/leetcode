package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))
}

func arrayRankTransform(arr []int) []int {
	var t = make([]int, len(arr))
	copy(t, arr)

	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })

	m := map[int]int{}
	var rank = 1
	for _, v := range t {
		if m[v] == 0 {
			m[v] = rank
			rank++
		}
	}

	var res []int
	for _, v := range arr {
		res = append(res, m[v])
	}
	return res
}
