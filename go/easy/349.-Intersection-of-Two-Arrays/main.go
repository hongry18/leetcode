package main

import (
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	var res []int

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		if m[nums1[i]] {
			continue
		}
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				break
			}
		}
		m[nums1[i]] = true
	}

	sort.Ints(res)
	return res
}
