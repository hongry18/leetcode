package main

// best practice
func reverseBits(num uint32) uint32 {
	var res uint32
	n := 31
	for n > -1 {
		res += num & 1 << n
		num >>= 1
		n--
	}
	return res
}

func reverseBits2(num uint32) uint32 {
	var ar []uint32
	for num > 0 {
		ar = append(ar, num&1)
		num >>= 1
	}

	for len(ar) < 32 {
		ar = append(ar, 0)
	}

	for i, v := range ar {
		if v == 1 {
			ar = ar[i:]
			break
		}
	}

	var res uint32
	for i := 0; i < len(ar); i++ {
		res = res*2 + ar[i]
	}

	return res
}
