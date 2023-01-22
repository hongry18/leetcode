package main

func countSegments(s string) int {
	var res int
	var isChar bool

	for _, c := range s {
		if c == ' ' {
			if isChar {
				res++
			}
			isChar = false
		} else {
			isChar = true
		}
	}

	if isChar {
		res++
	}

	return res
}
