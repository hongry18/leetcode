package main

import "fmt"

func main() {
	addToArrayForm([]int{1, 2, 0, 0}, 34)
	addToArrayForm([]int{2, 7, 4}, 181)
	addToArrayForm([]int{2, 1, 5}, 806)
	addToArrayForm([]int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}, 516)
	addToArrayForm([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1)
	addToArrayForm([]int{0}, 23)
	addToArrayForm([]int{6}, 809)
}

func addToArrayForm(num []int, k int) []int {
	ar := []int{}
	for k > 0 {
		ar = append(ar, k%10)
		k /= 10
	}

	for i, j := 0, len(ar)-1; i < j; i, j = i+1, j-1 {
		ar[i], ar[j] = ar[j], ar[i]
	}

	if len(ar) > len(num) {
		ar, num = num, ar
	}

	var j int = len(num) - 1
	var t int
	for i := len(ar) - 1; i >= 0; i-- {
		num[j] += t + ar[i]
		if num[j] > 9 {
			num[j] -= 10
			t = 1
		} else {
			t = 0
		}
		j--
	}

	for j >= 0 {
		num[j] += t
		if num[j] > 9 {
			num[j] -= 10
			t = 1
		} else {
			t = 0
			break
		}
		j--
	}

	if t == 1 {
		num = append([]int{1}, num...)
	}

	fmt.Println(num)
	return num
}
