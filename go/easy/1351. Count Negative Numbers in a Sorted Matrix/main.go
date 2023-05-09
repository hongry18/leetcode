package main

import "fmt"

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
	fmt.Println(countNegatives([][]int{{3, 2}, {1, 0}}))
}

func countNegatives(grid [][]int) int {
	var cnt int
	var m, n = len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		l, r := 0, n-1

		if grid[i][l] < 0 {
			cnt += n
			continue
		}

		if grid[i][r] > -1 {
			continue
		}

		for l < r {
			mid := l + (r-l)/2
			if grid[i][mid] < 0 {
				r = mid
			} else {
				l = mid + 1
			}

		}
		cnt += n - l
	}

	return cnt
}
