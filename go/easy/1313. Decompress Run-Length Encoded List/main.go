package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4, 6, 5}))
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(decompressRLElist([]int{2, 1, 3, 4}))
}

func decompressRLElist(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
