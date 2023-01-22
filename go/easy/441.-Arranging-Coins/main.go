package main

func arrangeCoins(n int) int {
	var i, j int = 0, n

	for i <= j {
		m := i + (j-i)/2
		if m*(m+1)/2 <= n {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return j
}
