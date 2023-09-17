package main

/*
26개 배열을 만든다.
반복을 하면서 26개의 원소의 인덱스가 1이면  cnt를 재정의하고 26배열을 0으로 초기화한다.
*/
func lengthOfLongestSubstring(s string) int {
	var max int
	var n int = len(s)

	for i := 0; i < n; i++ {
		var cnt int
		exists := make([]int, 256)

		for j := i; j < n; j++ {
			if exists[s[j]] != 0 {
				break
			}

			exists[s[j]]++
			cnt++
		}

		if cnt > max {
			max = cnt
		}
	}

	return max
}
