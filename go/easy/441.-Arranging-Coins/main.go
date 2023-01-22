package main

func arrangeCoins(n int) int {
	var i int = 1
	for n > i {
		n -= i
		i++
	}

	if n == i {
		return i
	} else {
		return i - 1
	}
}
