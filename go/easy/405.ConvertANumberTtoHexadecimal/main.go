package main

func toHex(num int) string {
	if num < 0 {
		var tmp uint32
		tmp += uint32(num)
		num = int(tmp)
	}

	var alpha = []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	var res []byte
	for num > 0 {
		mod := num % 16
		if mod < 10 {
			res = append([]byte{byte(mod) + '0'}, res...)
		} else {
			res = append([]byte{alpha[mod-10]}, res...)
		}
		num /= 16
	}
	return string(res)
}
