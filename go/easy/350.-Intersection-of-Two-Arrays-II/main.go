package main

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	var i, j int
	var res []int

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
			continue
		}

		if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}

	return res
}
