package main

func isHappy(n int) bool {
	for {
		if n == 1 || n == 7 {
			return true
		} else if n < 10 {
			return false
		}

		var tmp int
		for n > 0 {
			tmp += (n % 10) * (n % 10)
			n /= 10
		}
		n = tmp
	}
}
