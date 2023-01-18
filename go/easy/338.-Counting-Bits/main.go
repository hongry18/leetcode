package main

func countBits(n int) []int {
	var res []int

	for i := 0; i <= n; i++ {
		j := i
		cnt := 0
		for j > 0 {
			if j&1 == 1 {
				cnt++
			}
			j >>= 1
		}
		res = append(res, cnt)
	}

	return res
}
