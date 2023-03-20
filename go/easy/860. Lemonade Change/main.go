package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}

func lemonadeChange(bills []int) bool {
	ar := []int{0, 0}
	for i := range bills {
		switch bills[i] {
		case 5:
			ar[0]++
		case 10:
			if ar[0] < 1 {
				return false
			}
			ar[0]--
			ar[1]++
		case 20:
			if ar[0] > 0 && ar[1] > 0 {
				ar[0]--
				ar[1]--
			} else if ar[0] > 2 {
				ar[0] -= 3
			} else {
				return false
			}
		}

	}
	return true
}
