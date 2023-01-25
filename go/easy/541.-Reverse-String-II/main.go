package main

func reverseStr(s string, k int) string {
	var b = []byte(s)
	for i := 0; i < len(b); i += k * 2 {
		l := i + k
		if l > len(b) {
			l = len(b)
		}
		e := (l-i)/2 + i
		for j := i; j < e; j++ {
			b[j], b[l-1-j+i] = b[l-1-j+i], b[j]
		}
	}
	return string(b)
}
