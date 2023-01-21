package main

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	var res []string
	if turnedOn > 8 {
		return res
	}
	var ref = [60]int{}
	for i := 1; i < 60; i++ {
		ref[i] = ref[i>>1] + (i & 1)
	}

	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			b := ref[i] + ref[j]

			if b != turnedOn {
				continue
			}

			res = append(res, fmt.Sprintf("%d:%02d", i, j))
		}
	}
	return res
}
