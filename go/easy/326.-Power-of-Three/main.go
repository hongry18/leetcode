package main

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}

	if n < 1 || n%3 != 0 {
		return false
	}

	for n > 3 {
		if n%3 != 0 {
			return false
		}

		n /= 3
	}

	return true
}
