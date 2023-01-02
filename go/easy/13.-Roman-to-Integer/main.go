package main

func romanToInt(s string) int {
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	size := len(s)
	var result, tmp int

	for i := size - 1; i >= 0; i-- {
		v := romans[s[i]]

		if v < tmp {
			result -= v
		} else {
			result += v
		}
		tmp = v
	}

	return result
}
