package main

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	var res []int

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	var i int
	m := make(map[int]int)
	for j := 0; j < len(nums1); j++ {
		m[nums1[j]]++
		m[nums2[j]]++
		i++
	}

	for i < len(nums2) {
		m[nums2[i]]++
		i++
	}

	for j, v := range m {
		if v > 1 {
			res = append(res, j)
		}
	}
	sort.Ints(res)
	return res
}
