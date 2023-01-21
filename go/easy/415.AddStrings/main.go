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
	var hasTen bool
	var ar []byte
	var b byte
	for i, j = len(num1)-1, len(num2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		b = num1[i] - '0' + num2[j] - '0'
		if hasTen {
			b++
			hasTen = false
		}

		if b > 9 {
			b -= 10
			hasTen = true
		}

		ar = append(ar, b)
	}

	for ; j >= 0; j-- {
		b = num2[j] - '0'
		if hasTen {
			b++
			hasTen = false
		}
		if b > 9 {
			b -= 10
			hasTen = true
		}
		ar = append(ar, b)
	}

	if hasTen {
		ar = append(ar, 1)
	}

	var res []byte
	for i := len(ar) - 1; i >= 0; i-- {
		res = append(res, ar[i]+'0')
	}

	return string(res)
}
