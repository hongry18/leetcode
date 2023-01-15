package main

func containsDuplicate(nums []int) bool {
	m := map[int]int{}

	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v]++
		} else {
			return true
		}
	}
	return false
}
