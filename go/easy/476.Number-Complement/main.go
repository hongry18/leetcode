package main

func findComplement(num int) int {
	var sum, tmp = 1, num

	for tmp > 1 {
		sum <<= 1
		sum += 1
		tmp >>= 1
	}
	return sum - num
}
