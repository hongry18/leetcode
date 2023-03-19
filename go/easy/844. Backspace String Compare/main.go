package main

import "fmt"

func main() {
	// fmt.Println(backspaceCompare("ab#c", "ad#c"))
	// fmt.Println(backspaceCompare("ab##", "c#d#"))
	// fmt.Println(backspaceCompare("a#c", "b"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
}

func backspaceCompare(s string, t string) bool {
	a1, a2 := []byte{}, []byte{}
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			a1 = append(a1, s[i])
		}
	}

	cnt = 0
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] == '#' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			a2 = append(a2, t[i])
		}
	}

	return string(a1) == string(a2)
}
