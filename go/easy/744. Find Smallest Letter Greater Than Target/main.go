package main

import "fmt"

func main() {
	nextGreatestLetter([]byte{'c', 'd', 'e', 'f', 'h', 'i', 'j', 'k', 'z'}, 'a')
}

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1

	for l < r {
		m := l + (r-l)/2

		fmt.Printf("l:%d r:%d m:%d\n", l, r, m)
		if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}

	if letters[l] <= target {
		return letters[0]
	}

	return letters[l]
}
