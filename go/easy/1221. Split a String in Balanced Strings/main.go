package main

func main() {
	balancedStringSplit("RLRRLLRLRL")
}

func balancedStringSplit(s string) int {
	var cnt, l, r int = 0, 0, 0
	for _, c := range s {
		if c == 'L' {
			l++
		} else {
			r++
		}

		if l == r {
			cnt++
			l = 0
			r = 0
		}
	}
	return cnt
}
