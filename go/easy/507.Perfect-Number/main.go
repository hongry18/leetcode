package main

func checkPerfectNumber(num int) bool {
	if num < 2 {
		return false
	}
	var sum int
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	sum++
	return sum == num
}
