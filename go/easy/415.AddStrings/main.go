package main

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	} else if num2 == "0" {
		return num1
	} else if num1 == "0" && num2 == "0" {
		return num1
	}

	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	var i, j int
	var ar []byte
	var b, carry byte
	for i, j = len(num1)-1, len(num2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		b = num1[i] - '0' + num2[j] - '0' + carry

		if b > 9 {
			b -= 10
			carry = 1
		} else {
			carry = 0
		}

		ar = append(ar, b+'0')
	}

	for ; j >= 0; j-- {
		b = num2[j] - '0' + carry
		if b > 9 {
			b -= 10
			carry = 1
		} else {
			carry = 0
		}
		ar = append(ar, b+'0')
	}

	if carry > 0 {
		ar = append(ar, '1')
	}

	for i := 0; i < len(ar)/2; i++ {
		ar[i], ar[len(ar)-1-i] = ar[len(ar)-1-i], ar[i]
	}

	return string(ar)
}
