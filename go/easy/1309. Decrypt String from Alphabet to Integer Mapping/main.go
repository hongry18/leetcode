package main

import "fmt"

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
	fmt.Println(freqAlphabets("1326#"))
}

func freqAlphabets(s string) string {
	var b []byte
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			b = append(b, (s[i-1]-'0')+((s[i-2]-'0')*10)+'a'-1)
			i -= 2
		} else {
			b = append(b, s[i]-'0'+'a'-1)
		}
	}
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	return string(b)
}
