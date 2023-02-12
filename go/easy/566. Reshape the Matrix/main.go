package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	var res = make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	row, col := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[row][col] = mat[row][col]
			col++
			if col > c-1 {
				row++
				col = 0
			}
		}
	}

	return res
}
