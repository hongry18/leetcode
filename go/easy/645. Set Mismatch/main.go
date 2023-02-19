package main

import "fmt"

func main() {
	// fmt.Println(findErrorNums([]int{1, 1}))
	// fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{3, 3, 1}))
}

func findErrorNums(nums []int) []int {
	var ar = make([]int, len(nums))
	var res []int
	for _, v := range nums {
		ar[v-1]++
		if ar[v-1] == 2 {
			res = append(res, v)
			break
		}
	}

	fmt.Println(ar)

	for i, v := range ar {
		if v == 0 {
			res = append(res, i+1)
			break
		}
	}
	return res
}
