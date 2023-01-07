package main

func isPalindrome(s string) bool {
	var ar []rune

	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			ar = append(ar, c)
		} else if c >= 'A' && c <= 'Z' {
			ar = append(ar, c+32)
		}
	}

	size := len(ar)
	for i := 0; i < size/2; i++ {
		if ar[i] != ar[size-i-1] {
			return false
		}
	}
	return true
}
