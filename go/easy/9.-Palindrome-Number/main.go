package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	t := x
	var n int

	for t > 0 {
		n = n*10 + t%10
		t /= 10
	}

	return n == x
}
