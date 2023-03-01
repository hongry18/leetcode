package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("LR"))
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("UU"))
}

func judgeCircle(moves string) bool {
	var ar = [4]int{}

	for _, r := range moves {
		switch r {
		case 'L':
			ar[0]++
		case 'R':
			ar[1]++
		case 'U':
			ar[2]++
		case 'D':
			ar[3]++
		}
	}

	return ar[0]-ar[1] == 0 && ar[2]-ar[3] == 0
}
