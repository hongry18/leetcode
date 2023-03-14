package main

func main() {
	largeGroupPositions("abbxxxxzzy")
	largeGroupPositions("aaa")
	largeGroupPositions("abcdddeeeeaabbbcd")
}

func largeGroupPositions(s string) [][]int {
	var res [][]int

	i, l := 1, 0
	for ; i < len(s); i++ {
		if s[i] != s[i-1] {
			if i-l > 2 {
				res = append(res, []int{l, i - 1})
			}
			l = i
		}
	}

	if i-l > 2 {
		res = append(res, []int{l, i - 1})
	}
	return res
}
