package main

func checkRecord(s string) bool {
	var a, l int

	for i := range s {

		if s[i] == 'L' {
			l++
			if l > 2 {
				return false
			}
			continue
		}

		l = 0
		if s[i] == 'P' {
			continue
		}
		a++
		if a > 1 {
			return false
		}
	}

	return true
}
