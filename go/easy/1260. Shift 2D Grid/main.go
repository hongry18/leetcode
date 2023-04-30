package main

import "fmt"

func main() {
	// fmt.Println(shiftGrid([][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4))
	fmt.Println(shiftGrid([][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}}, 23))
}

func shiftGrid(grid [][]int, k int) [][]int {
	var ar []int
	var m, n, t int = len(grid), len(grid[0]), 0
	if k%(m*n) == 0 {
		return grid
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ar = append(ar, grid[i][j])
		}
	}

	if k > len(ar) {
		k = len(ar) % k
	}
	ar = append(ar[len(ar)-k:], ar[:len(ar)-k]...)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = ar[t]
			t++
		}
	}

	return grid
}
