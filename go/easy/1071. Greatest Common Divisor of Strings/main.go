package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(gcdOfStrings("CVABED", "ABAB"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABCDEFA", "ABAB"))
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))
}

func gcdOfStrings(str1 string, str2 string) string {
	if strings.Join([]string{str1, str2}, "") != strings.Join([]string{str2, str1}, "") {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
