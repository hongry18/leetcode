package main

import "strconv"

func fizzBuzz(n int) []string {
	var res = make([]string, n)
	for i := range res {
		switch (i + 1) % 15 {
		case 0:
			res[i] = "FizzBuzz"
		case 3, 6, 9, 12:
			res[i] = "Fizz"
		case 5, 10:
			res[i] = "Buzz"
		default:
			res[i] = strconv.Itoa(i + 1)
		}
	}
	return res
}
