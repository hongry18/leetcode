package main

func findComplement(num int) int {
	var ar []int
	var res, inc int = 0, 1
	for num > 0 {
		if num&1 == 0 {
			ar = append(ar, 1)
		} else {
			ar = append(ar, 0)
		}
		num >>= 1
	}

	for i := range ar {
		res += (inc) * ar[i]
		inc *= 2
	}
	return res
}
