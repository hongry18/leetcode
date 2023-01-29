package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var res = make([]string, len(score))
	var m = map[int]int{}
	for i, v := range score {
		m[v] = i
	}
	sort.Slice(score, func(x, y int) bool { return score[x] > score[y] })
	for i, v := range score {
		switch i {
		case 0:
			res[m[v]] = "Gold Medal"
		case 1:
			res[m[v]] = "Silver Medal"
		case 2:
			res[m[v]] = "Bronze Medal"
		default:
			res[m[v]] = strconv.Itoa(i + 1)
		}
	}
	return res
}
