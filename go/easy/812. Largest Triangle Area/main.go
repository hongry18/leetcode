package main

import "fmt"

func main() {
	// largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}})
	// largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}})
	// largestTriangleArea([][]int{{4, 6}, {6, 5}, {3, 1}})
	permu([][]int{{0, 0}, {0, 1}, {1, 0}}, 0, 3, 3)
}

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	var res float64
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				t := float64(abs(points[i][0]*points[j][1]+points[j][0]*points[k][1]+points[k][0]*points[i][1]-points[i][1]*points[j][0]-points[j][1]*points[k][0]-points[k][1]*points[i][0])) / 2
				if res < t {
					res = t
				}
			}
		}
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func permu(p [][]int, d, n, k int) {
	if d == k {
		for i := 0; i < k; i++ {
			fmt.Print(p[i])
		}
		fmt.Println()
		return
	}

	for i := d; i < n; i++ {
		p[i], p[d] = p[d], p[i]
		permu(p, d+1, n, k)
		p[i], p[d] = p[d], p[i]
	}
}
