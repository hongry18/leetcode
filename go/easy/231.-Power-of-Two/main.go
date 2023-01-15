package main

func isPowerOfTwo(n int) bool {
	if (n != 1 && n%2 != 0) || n < 1 {
		return false
	}

	if n < 3 {
		return true
	}

	for n > 2 {
		if n%2 == 1 {
			return false
		}
		n /= 2
	}

	return true
}
