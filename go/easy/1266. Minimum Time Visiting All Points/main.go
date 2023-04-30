package main

import "math"

func main() {

}

func minTimeToVisitAllPoints(points [][]int) int {
	var step float64
	for i := 1; i < len(points); i++ {
		x := math.Abs(float64(points[i][0] - points[i-1][0]))
		y := math.Abs(float64(points[i][1] - points[i-1][1]))
		step += math.Max(x, y)
	}
	return int(step)
}
