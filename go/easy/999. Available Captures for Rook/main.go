package main

func numRookCaptures(board [][]byte) int {
	var x, y int

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				x, y = i, j
			}
		}
	}

	var cnt int
	m := [][]int{{-1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, -1, 0}, {0, 0, 0, 1}}
	for _, i := range m {
		if move(board, i, x, y) {
			cnt++
		}
	}

	return cnt
}

func move(b [][]byte, m []int, x, y int) bool {
	for {
		x += m[0] + m[1]
		y += m[2] + m[3]
		if x < 0 || x > 7 || y < 0 || y > 7 || b[x][y] == 'B' {
			break
		}

		if b[x][y] == 'p' {
			return true
		}
	}

	return false
}
