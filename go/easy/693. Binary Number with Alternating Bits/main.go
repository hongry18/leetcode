package main

func main() {
	hasAlternatingBits(5)
	hasAlternatingBits(7)
	hasAlternatingBits(11)
}

func hasAlternatingBits(n int) bool {
	ar := []int{}
	for n != 0 {
		ar = append(ar, n%2)
		n >>= 1
	}
	for i := 1; i < len(ar); i++ {
		if ar[i] == ar[i-1] {
			return false
		}
	}
	return true
}
