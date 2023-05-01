package main

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}

func replaceElements(arr []int) []int {
	var t int = -1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > t {
			t, arr[i] = arr[i], t
		} else {
			arr[i] = t
		}
	}

	return arr
}
