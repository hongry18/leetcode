package main

func checkIfExist(arr []int) bool {
	m := make(map[int]bool)
	z := 0

	for _, n := range arr {
		if n == 0 {
			z++
		}
		m[n] = true
	}

	for k := range m {
		if k != 0 && m[k*2] {
			return true
		}

		if k == 0 && z > 1 {
			return true
		}
	}
	return false
}
