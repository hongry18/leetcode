package main

import "fmt"

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}

func isToeplitzMatrix(m [][]int) bool {
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {
			if m[i][j] != m[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

func isToeplitzMatrix2(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	a, b := m, 0
	for i := 0; i < m+n-1; i++ {
		if m-1-i >= 0 {
			a--
		} else {
			b++
		}

		if !cmp(matrix, a, b, m, n) {
			return false
		}
	}
	return true
}

func cmp(matrix [][]int, a, b, m, n int) bool {
	a++
	b++

	for a < m && b < n {
		if matrix[a][b] != matrix[a-1][b-1] {
			return false
		}
		a++
		b++
	}
	return true
}
