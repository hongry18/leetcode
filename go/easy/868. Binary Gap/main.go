package main

import "fmt"

func main() {
	binaryGap(22)
	fmt.Println("#####")
	binaryGap(12)
	fmt.Println("#####")
	binaryGap(8)
}

func binaryGap(n int) int {
	var cnt, res = 0, 0
	for n > 0 {
		if n%2 == 1 {
			if cnt > res {
				res = cnt
			}
			cnt = 1
		} else {
			if cnt > 0 {
				cnt++
			}
		}
		n >>= 1
	}

	return res
}
