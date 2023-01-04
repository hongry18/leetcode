package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j, k int = m - 1, n - 1, m + n - 1

	for ; i >= 0 && j >= 0; k-- {
		if nums2[j] > nums1[i] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}

	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
}
