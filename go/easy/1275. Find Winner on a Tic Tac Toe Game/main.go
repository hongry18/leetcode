package main

func main() {
	tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}})
}

func tictactoe(moves [][]int) string {
	var ar [][]byte = make([][]byte, 3)
	for i := 0; i < 3; i++ {
		ar[i] = make([]byte, 3)
		ar[i][0] = byte(3*i + 1)
		ar[i][0] = byte(3*i + 2)
		ar[i][0] = byte(3*i + 3)
	}

	for i := 0; i < len(moves); i++ {
		if i%2 == 0 {
			ar[moves[i][0]][moves[i][1]] = 'A'
		} else {
			ar[moves[i][0]][moves[i][1]] = 'B'
		}
	}

	for i := 0; i < 3; i++ {
		if ar[i][0] == ar[i][1] && ar[i][1] == ar[i][2] {
			return string(ar[i][0])
		}

		if ar[0][i] == ar[1][i] && ar[1][i] == ar[2][i] {
			return string(ar[0][i])
		}
	}

	if (ar[0][2] == ar[1][1] && ar[1][1] == ar[2][0]) || (ar[0][0] == ar[1][1] && ar[1][1] == ar[2][2]) {
		return string(ar[1][1])
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
