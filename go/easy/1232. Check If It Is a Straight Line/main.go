package main

import "fmt"

func main() {
	fmt.Println(checkStraightLine([][]int{{0, 0}, {0, 1}, {0, -1}}))
}

func checkStraightLine(coordinates [][]int) bool {
	t := float64(coordinates[1][1]-coordinates[0][1]) / float64(coordinates[1][0]-coordinates[0][0])

	fmt.Println(t)
	for i := 2; i < len(coordinates); i++ {
		fmt.Println(float64(coordinates[i][1]-coordinates[i-1][1]) / float64(coordinates[i][0]-coordinates[i-1][0]))
		if t != float64(coordinates[i][1]-coordinates[i-1][1])/float64(coordinates[i][0]-coordinates[i-1][0]) {
			return false
		}
	}
	return true
}
