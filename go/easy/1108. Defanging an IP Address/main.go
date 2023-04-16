package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddr("255.100.50.0"))
}

func defangIPaddr(address string) string {
	var buf bytes.Buffer
	for _, r := range address {
		if r == '.' {
			buf.WriteString("[.]")
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
