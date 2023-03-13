package main

func main() {
	numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz")
	numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa")
}

func numberOfLines(widths []int, s string) []int {
	var ar []int
	var cnt int
	for i := 0; i < len(s); i++ {
		if cnt+widths[s[i]-'a'] < 101 {
			cnt += widths[s[i]-'a']
		} else {
			ar = append(ar, cnt)
			cnt = widths[s[i]-'a']
		}
	}

	return []int{len(ar) + 1, cnt}
}
