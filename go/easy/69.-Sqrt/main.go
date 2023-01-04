package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	var lo, hi int = 0, x

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if mid <= x/mid && (mid+1) > x/(mid+1) {
			return mid
		}

		if mid < x/mid {
			lo = mid - 1
		} else {
			hi = mid + 1
		}
	}

	return hi
}
