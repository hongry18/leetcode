package main

func main() {

}

func hasGroupsSizeX(deck []int) bool {
	m := map[int]int{}
	for _, v := range deck {
		m[v]++
	}

	d := 0
	for _, v := range m {
		d = gcd(v, d)
		if d < 2 {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
