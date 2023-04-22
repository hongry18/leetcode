package main

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var f, b, i, j, l int = 0, 0, start, destination, len(distance)

	for i != destination {
		f += distance[i]
		i++
		if i == l {
			i = 0
		}
	}

	for j != start {
		b += distance[j]
		j++
		if j == l {
			j = 0
		}
	}

	if f > b {
		return b
	}
	return f
}
