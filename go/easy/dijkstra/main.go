package main

import "fmt"

func main() {
	dijkstra(10, 10, []int{2, 3}, []int{7, 8})
}

type Pos struct {
	x, y, cost int
}

func dijkstra(m, n int, start, end []int) {
	var ar = make([][]int, m)
	for i := 0; i < m; i++ {
		ar[i] = make([]int, n)
	}

	ar[end[0]][end[1]] = 1

	q := []Pos{{x: start[0], y: start[1]}}

	for len(q) != 0 {

	}

	fmt.Println(ar)
}
