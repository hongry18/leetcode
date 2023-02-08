package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4))
}
func findMaxAverage(nums []int, k int) float64 {
	var res float64
	for i := 0; i < len(nums)-k+1; i++ {
		var sum float64
		for j := i; j < i+k; j++ {
			sum += float64(nums[j])
		}
		t := sum / float64(k)
		if res < t {
			res = t
		}
	}
	return res
}
