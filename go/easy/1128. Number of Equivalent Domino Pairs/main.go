package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {2, 1}, {2, 2}}))
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 2}, {2, 1}, {2, 1}}))
}

func numEquivDominoPairs(dominoes [][]int) int {
	var sum int
	var m = make(map[string]int)

	for _, d := range dominoes {
		var buf bytes.Buffer

		if d[0] < d[1] {
			buf.WriteByte(byte(d[0]))
			buf.WriteByte(',')
			buf.WriteByte(byte(d[1]))
		} else {
			buf.WriteByte(byte(d[1]))
			buf.WriteByte(',')
			buf.WriteByte(byte(d[0]))
		}

		m[buf.String()]++
	}

	for _, v := range m {
		if v < 2 {
			continue
		}
		sum += v * (v - 1) / 2
	}

	return sum
}
