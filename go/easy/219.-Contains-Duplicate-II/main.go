package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[v]; ok && abs(i-j) <= k {
			return true
		}
		m[v] = i
	}
	return false
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
