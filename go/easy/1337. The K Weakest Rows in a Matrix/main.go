package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kWeakestRows([][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3))
}

func kWeakestRows(mat [][]int, k int) []int {
	var res []int
	var tmp [][]int

	for i := 0; i < len(mat); i++ {
		var cnt int
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				cnt++
			}
		}
		tmp = append(tmp, []int{i, cnt})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return (tmp[i][1] == tmp[j][1] && tmp[i][0] < tmp[j][0]) || tmp[i][1] < tmp[j][1]
	})
	for i := 0; i < k; i++ {
		res = append(res, tmp[i][0])
	}
	return res
}
