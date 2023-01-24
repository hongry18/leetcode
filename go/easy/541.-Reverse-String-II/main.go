package main

func reverseStr(s string, k int) string {
	var b = []byte(s)
	for i := 0; i < k/2; i++ {
		b[i], b[k-1-i] = b[k-1-i], b[i]
	}
	return string(b)
}
