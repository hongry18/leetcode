package main

func intersection(nums1 []int, nums2 []int) []int {
	var res = []int{}
	var find = map[int]bool{}
	for _, n := range nums1 {
		find[n] = true
	}

	for _, n := range nums2 {
		if !find[n] {
			continue
		}
		res = append(res, n)
		find[n] = false
	}

	return res
}
