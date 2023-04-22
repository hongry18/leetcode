package main

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var f, b int

	for i := start; i != destination; i++ {
		f += distance[i]
		if i == len(distance)-1 {
			i = -1
		}
	}

	for i := destination; i != start; i++ {
		b += distance[i]
		if i == len(distance)-1 {
			i = -1
		}
	}

	return min(f, b)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
