package main

import "fmt"

func main() {

	fmt.Println(removePalindromeSub("ababa"))
	fmt.Println(removePalindromeSub("abb"))
	fmt.Println(removePalindromeSub("baabb"))

}

func removePalindromeSub(s string) int {
	var l, r = 0, len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return 2
		}
		l++
		r--
	}
	return 1
}
