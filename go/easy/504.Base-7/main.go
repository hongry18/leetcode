package main

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var b []byte
	var isNegative bool
	if num < 0 {
		isNegative = true
		num *= -1
	}
	for num > 0 {
		b = append(b, byte(num%7)+'0')
		num /= 7
	}
	if isNegative {
		b = append(b, '-')
	}
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b)
}
