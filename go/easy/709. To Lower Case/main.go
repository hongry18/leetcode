package main

func toLowerCase(s string) string {
	var ar = []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			ar = append(ar, s[i]-'A'+'a')
		} else {
			ar = append(ar, s[i])
		}
	}
	return string(ar)
}
