package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	var ar [101]int

	for _, n := range nums {
		ar[n]++
	}

	var cnt int
	for i := 0; i < 101; i++ {
		ar[i], cnt = cnt, cnt+ar[i]
	}

	var res = make([]int, len(nums))
	for i := range nums {
		res[i] = ar[nums[i]]
	}

	return res
}
