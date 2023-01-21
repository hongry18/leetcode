package main

func longestPalindrome(s string) int {
	var res int
	var center bool
	var ar = [26]int{}
	for _, v := range s {
		if v >= 'a' {
			ar[v-'a']++
		} else {
			ar[v-'A'+26]++
		}
	}

	for _, v := range ar {
		res += v
		if v%2 == 1 {
			if center {
				res--
			} else {
				center = true
			}
		}
	}

	return res
}
