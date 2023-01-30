package main

func reverseWords(s string) string {
	var res, word []byte
	for i := range s {
		if s[i] == ' ' {
			res = append(res, word...)
			res = append(res, s[i])
			word = []byte{}
		} else {
			word = append([]byte{s[i]}, word...)
		}
	}
	res = append(res, word...)
	return string(res)
}
