package main

func countBits(n int) []int {
	var res []int

	for i := 0; i <= n; i++ {
		j := i
		cnt := 0
		for j > 0 {
			j &= j - 1
			cnt++
		}
		res = append(res, cnt)
	}

	return res
}
