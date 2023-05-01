package main

func findNumbers(nums []int) int {
	var res int
	for _, n := range nums {
		if Even(n) {
			res++
		}
	}
	return res
}

func IsEven(a int) bool {
	if a > 99999 || (a < 10000 && a > 999) || (a > 9 && a < 100) {
		return true
	}
	return false
}

func Even(a int) bool {
	var cnt int
	for a > 0 {
		a /= 10
		cnt++
	}

	return cnt%2 == 0
}
