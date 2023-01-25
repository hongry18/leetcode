package main

func reverseStr(s string, k int) string {
	var b = []byte(s)
	for i := 0; i < len(b); i += k * 2 {
		if i+k < len(b) {
			reverse(b[i : i+k])
		} else {
			reverse(b[i:])
		}
	}
	return string(b)
}

func reverse(b []byte) {
	var l, r = 0, len(b) - 1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l, r = l+1, r-1
	}
}
