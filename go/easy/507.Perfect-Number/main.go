package main

func checkPerfectNumber(num int) bool {
	var sum int = 1
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return num == sum
}
