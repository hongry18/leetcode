package main

import "fmt"

func main() {
	fmt.Println(isPrime(5))
	fmt.Println(isPrime(100))
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	ar := make([]int, n+1)

	for i := 2; i <= n; i++ {
		if ar[i] == 0 {
			for j := i * i; j <= n; j += i {
				ar[j] = 1
			}
		}
	}

	return ar[n] == 0
}
