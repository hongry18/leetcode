package main

func main() {
	transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	transpose([][]int{{1, 2, 3}, {4, 5, 6}})
}

func transpose(matrix [][]int) [][]int {
	var res = make([][]int, len(matrix[0]))

	for i := range res {
		for j := range matrix {
			res[i] = append(res[i], matrix[j][i])
		}
	}
	return res
}
