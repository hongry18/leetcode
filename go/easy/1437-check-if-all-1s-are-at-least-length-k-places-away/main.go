package main

func kLengthApart(nums []int, k int) bool {
	var i int
	for j, n := range nums {
		if n != 0 {
			i = j
			break
		}
	}
	var cnt int
	for _, n := range nums[i+1:] {
		if n == 1 {
			if cnt < k {
				return false
			}
			cnt = 0
		} else {
			cnt++
		}
	}
	return true
}
