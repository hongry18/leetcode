package main

func main() {
	maxCount(3, 3, [][]int{{2, 3}, {3, 3}})
	maxCount(3, 3, [][]int{{2, 2}, {3, 3}})
	maxCount(3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}})
}

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) < 1 {
		return m * n
	}

	var x, y int = ops[0][0], ops[0][1]

	for i := 1; i < len(ops); i++ {
		if ops[i][0] < x {
			x = ops[i][0]
		}

		if ops[i][1] < y {
			y = ops[i][1]
		}
	}

	return x * y
}
