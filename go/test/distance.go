package test

func solution(board [][]int, c int) int {
	n := len(board)
	m := len(board[0])

	var robotPos, destPos Pos
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 2 {
				robotPos = Pos{i, j}
			} else if board[i][j] == 3 {
				destPos = Pos{i, j}
			}
		}
	}

	var minEnergy int
	visited := make(map[Pos]bool)
	dfs(board, robotPos, destPos, 0, c, &minEnergy, visited)

	return minEnergy
}

func dfs(board [][]int, pos, dest Pos, energy, c int, minEnergy *int, visited map[Pos]bool) {
	if energy >= *minEnergy { // 현재까지 소비한 전력이 이미 이전에 구한 최소 전력보다 크거나 같으면 중단
		return
	}
	if pos == dest { // 목적지에 도달했으면 최소 전력 업데이트
		if *minEnergy == 0 || energy < *minEnergy {
			*minEnergy = energy
		}
		return
	}

	visited[pos] = true
	defer delete(visited, pos)

	// 상하좌우로 이동
	for _, dir := range []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nextPos := pos.Add(dir)

		if nextPos.x < 0 || nextPos.x >= len(board) || nextPos.y < 0 || nextPos.y >= len(board[0]) {
			continue // 지도 범위를 벗어나는 경우
		}
		if board[nextPos.x][nextPos.y] == 1 {
			if c == 0 || visited[nextPos] {
				continue // 장애물을 제거할 수 없는 경우, 이미 방문한 경우
			}
			// 장애물 제거하고 이동
			board[nextPos.x][nextPos.y] = 0
			dfs(board, nextPos, dest, energy+c, c, minEnergy, visited)
			board[nextPos.x][nextPos.y] = 1 // 장애물 복원
		} else if !visited[nextPos] { // 이전에 방문하지 않은 빈 공간인 경우
			dfs(board, nextPos, dest, energy+1, c, minEnergy, visited)
		}
	}
}

type Pos struct {
	x, y int
}

func (p Pos) Add(q Pos) Pos {
	return Pos{p.x + q.x, p.y + q.y}
}
