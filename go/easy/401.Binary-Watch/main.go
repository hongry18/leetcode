package main

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	if turnedOn > 8 {
		return []string{}
	}
	if turnedOn == 0 {
		return []string{"0:00"}
	}
	data := []int{1, 2, 4, 8, 16, 32, 101, 102, 104, 108}
	visit := make([]bool, 10)
	var res []string
	comb(data, visit, 0, 10, turnedOn, &res)

	return res
}

func comb(ar []int, v []bool, s, n, r int, res *[]string) {
	if r == 0 {
		var m, h int
		for i := 0; i < n; i++ {
			if v[i] {
				if ar[i] > 100 {
					h += ar[i] - 100
				} else {
					m += ar[i]
				}
			}
		}

		if m > 59 {
			return
		}
		if h > 11 {
			return
		}
		*res = append(*res, fmt.Sprintf("%d:%02d", h, m))
		return
	}

	for i := s; i < n; i++ {
		v[i] = true
		comb(ar, v, i+1, n, r-1, res)
		v[i] = false
	}
}
