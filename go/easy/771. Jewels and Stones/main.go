package main

func numJewelsInStones(jewels string, stones string) int {
	var m = map[byte]bool{}
	for i := 0; i < len(jewels); i++ {
		m[jewels[i]] = true
	}

	var cnt int
	for i := 0; i < len(stones); i++ {
		if m[stones[i]] {
			cnt++
		}
	}

	return cnt
}
