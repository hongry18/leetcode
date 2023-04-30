package main

func uniqueOccurrences(arr []int) bool {
	var m = make(map[int]int)
	var ar = make([]int, 1001)

	for _, n := range arr {
		m[n]++
	}

	for k, v := range m {
		if ar[v] != 0 {
			return false
		}
		ar[v] = k
	}
	return true
}
