package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	split := strings.Split(text, " ")
	res := []string{}

	for i := 0; i < len(split)-2; i++ {
		if split[i] == first && split[i+1] == second {
			res = append(res, split[i+2])
		}
	}

	// fmt.Println(res)
	return res
}
