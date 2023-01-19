package main

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	var res []int

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	var i int
	m := make(map[int]int)
	f1 := make(map[int]bool)
	f2 := make(map[int]bool)
	for j := 0; j < len(nums1); j++ {
		if !f1[nums1[j]] {
			f1[nums1[j]] = true
			m[nums1[j]]++
		}
		if !f2[nums2[j]] {
			f2[nums2[j]] = true
			m[nums2[j]]++
		}
		i++
	}

	for i < len(nums2) {
		if !f2[nums2[i]] {
			f2[nums2[i]] = true
			m[nums2[i]]++
		}
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
