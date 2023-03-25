package main

func main() {
	fairCandySwap([]int{1, 1}, []int{2, 2})
	fairCandySwap([]int{1, 2}, []int{2, 3})
	fairCandySwap([]int{2}, []int{1, 3})
}

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	a1, a2 := 0, 0
	m := map[int]bool{}
	for _, v := range aliceSizes {
		a1 += v
		m[v] = true
	}

	for _, v := range bobSizes {
		a2 += v
	}

	diff := (a1 + a2) / 2
	for _, b := range bobSizes {
		t := a1 + b - diff
		if _, ok := m[t]; ok {
			return []int{t, b}
		}
	}
	return nil
}
