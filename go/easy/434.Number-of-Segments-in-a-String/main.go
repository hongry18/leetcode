package main

func countSegments(s string) int {
	var res []string

	var i int
	for j := range s {
		if s[i] == ' ' && s[j] != ' ' {
			i = j
		}

		if s[i] != ' ' && s[j] == ' ' {
			res = append(res, s[i:j])
			i = j
		}

		if j == len(s)-1 && s[i] != ' ' && s[j] != ' ' {
			res = append(res, s[i:j])
		}
	}

	return len(res)
}
