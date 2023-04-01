package main

func main() {
	addToArrayForm([]int{1, 2, 0, 0}, 34)
}

func addToArrayForm(num []int, k int) []int {
	var res []int

	var s int
	for _, n := range num {
		s = s*10 + n
	}

	s += k

	for s > 0 {
		res = append(res, s%10)
		s /= 10
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
