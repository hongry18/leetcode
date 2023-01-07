package main

func generate(numRows int) [][]int {
	result := [][]int{{1}}

	for i := 1; i < numRows; i++ {
		v := make([]int, i+1)
		v[0], v[i] = 1, 1
		for j := 1; j < len(v)-1; j++ {
			v[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, v)
	}

	return result
}
