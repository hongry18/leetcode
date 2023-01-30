package main

func reverseWords(s string) string {
	var ar []int
	ar = append(ar, 0)
	for i := range s {
		if s[i] == ' ' {
			ar = append(ar, i)
			ar = append(ar, i+1)
		}
	}
	ar = append(ar, len(s))

	var b = []byte(s)

	for i := 0; i < len(ar); i += 2 {

		e := (ar[i+1] + ar[i])
		for j := ar[i]; j < e/2; j++ {
			b[j], b[e-1-j] = b[e-1-j], b[j]
		}
	}

	return string(b)
}
