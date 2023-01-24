package main

func findMaxConsecutiveOnes(nums []int) int {
	var cnt, res int
	for i := range nums {
		if nums[i] == 0 {
			cnt = 0
		} else {
			cnt++
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}
