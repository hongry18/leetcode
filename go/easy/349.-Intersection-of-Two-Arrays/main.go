package main

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	find := make([]int, 1001)
	for _, n := range nums1 {
		if find[n] > 0 {
			continue
		}
		find[n]++
	}

	for _, n := range nums2 {
		if find[n] < 1 {
			continue
		}
		find[n]++
	}

	for i := 0; i < 1001; i++ {
		if find[i] > 1 {
			res = append(res, i)
		}
	}

	return res
}
