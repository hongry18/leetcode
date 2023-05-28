package main

func maxPower(s string) int {
	var cnt, max int = 1, 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}
