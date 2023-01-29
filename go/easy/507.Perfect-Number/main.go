package main

func checkPerfectNumber(num int) bool {
	if num < 2 {
		return false
	}
	var sum int
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i + (num / i)
		}
	}
	return sum-num == num
}
