package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		t := target - v
		if j, ok := m[t]; ok {
			return []int{j, i}
		}
		m[v] = i
	}

	return nil
}
