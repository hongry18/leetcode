package main

func titleToNumber(columnTitle string) int {
	var res int
	for i := 0; i < len(columnTitle); i++ {
		res = res*26 + int(columnTitle[i]-'A'+1)
	}
	return res
}
