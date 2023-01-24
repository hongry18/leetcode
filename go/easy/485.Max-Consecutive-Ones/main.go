package main

func findMaxConsecutiveOnes(nums []int) int {
	var cnt, res int
	for i := range nums {
		if nums[i] == 0 {
			if cnt > res {
				res = cnt
			}
			cnt = 0
			continue
		}
		cnt++
	}
	if cnt > res {
		res = cnt
	}
	return res
}
