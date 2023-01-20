package main

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	var l, r = 0, num

	for l < r {
		mid := l + (r-l)/2

		if mid <= num/mid && (mid+1) > num/(mid+1) {
			l = mid
			break
		}

		if mid < num/mid {
			l = mid - 1
		} else {
			r = mid + 1
		}
	}

	return float64(l) == float64(num)/float64(l)
}
