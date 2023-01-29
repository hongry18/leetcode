package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res []int
	var m = map[int]int{}
	for i, v := range nums2 {
		m[v] = i
	}

	for i := range nums1 {
		var r int = -1
		for j := m[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				r = nums2[j]
				break
			}
		}
		res = append(res, r)
	}
	return res
}
