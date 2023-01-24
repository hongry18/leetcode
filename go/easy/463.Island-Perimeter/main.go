package main

func islandPerimeter(grid [][]int) int {
	var cnt int

	var ar [][]int
	for gy := range grid {
		for gx := range grid[gy] {
			if grid[gy][gx] != 1 {
				continue
			}
			ar = append(ar, []int{gy, gx})

			for len(ar) != 0 {
				el := ar[len(ar)-1]
				ar = ar[:len(ar)-1]
				x, y := el[1], el[0]
				if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || grid[y][x] == 0 || grid[y][x] == 2 {
					continue
				}

				cnt += 4
				grid[y][x] = 2
				if y > 0 && grid[y-1][x] > 0 {
					cnt -= 2
				}

				if x > 0 && grid[y][x-1] > 0 {
					cnt -= 2
				}

				ar = append(ar,
					[]int{y + 1, x},
					[]int{y - 1, x},
					[]int{y, x + 1},
					[]int{y, x - 1},
				)
			}
			break
		}
	}

	return cnt
}
