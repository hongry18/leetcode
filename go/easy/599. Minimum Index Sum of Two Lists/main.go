package main

import "math"

func findRestaurant2(list1 []string, list2 []string) []string {
	var m1, m2 = make(map[string]int), make(map[string]int)
	var min = math.MaxInt
	for i := 0; i < len(list1); i++ {
		m1[list1[i]] = i
	}

	for i := 0; i < len(list2); i++ {
		if v, ok := m1[list2[i]]; ok {
			m2[list2[i]] = i + v

			if m2[list2[i]] < min {
				min = m2[list2[i]]
			}
		}
	}

	var res = make([]string, 0)
	for k, v := range m2 {
		if v == min {
			res = append(res, k)
		}
	}

	return res
}

func findRestaurant(list1 []string, list2 []string) []string {
	var m = make(map[string]int)
	var min = math.MaxInt
	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				m[list1[i]] = i + j
				if m[list1[i]] < min {
					min = m[list1[i]]
				}
				break
			}
		}
	}

	var res = make([]string, 0)
	for k, v := range m {
		if v == min {
			res = append(res, k)
		}
	}

	return res
}
