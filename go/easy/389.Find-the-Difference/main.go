package main

func findTheDifference(s string, t string) byte {
	var res int
	for i := range s {
		res ^= int(s[i])
	}
	for i := range t {
		res ^= int(t[i])
	}
	return byte(res)
}
