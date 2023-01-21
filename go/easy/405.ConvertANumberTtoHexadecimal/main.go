package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	var n uint32 = uint32(num)
	var hex = "0123456789abcdef"
	var res []byte
	for n > 0 {
		res = append([]byte{hex[n%16]}, res...)
		n /= 16
	}
	return string(res)
}
