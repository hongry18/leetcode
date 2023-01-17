package main

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	for {
		var x int
		for num > 0 {
			x += num % 10
			num /= 10
		}

		num = x
		if num < 10 {
			break
		}
		x = 0
	}

	return num
}
