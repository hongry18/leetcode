package main

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	out := make([]byte, len(a))
	var carry byte

	for i := 0; i < len(a); i++ {
		sum := carry
		ai := len(a) - i - 1
		bi := len(b) - i - 1

		if len(b)-1 >= i {
			sum += b[bi] - '0'
		}

		sum += a[ai] - '0'

		if sum > 1 {
			carry, sum = 1, sum-2
		} else {
			carry = 0
		}

		if sum == 1 {
			out[ai] = '1'
		} else {
			out[ai] = '0'
		}
	}

	if carry == 1 {
		out = append([]byte{'1'}, out...)
	}

	return string(out)
}
