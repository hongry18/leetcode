package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	var res [][]int
	var p int

	for _, row := range mat {
		for _, col := range row {

			if p%c == 0 {
				res = append(res, []int{col})
			} else {
				res[p/c] = append(res[p/c], col)
			}

			p++
		}
	}

	return res
}
