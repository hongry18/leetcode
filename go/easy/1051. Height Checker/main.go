package main

func main() {
	heightChecker([]int{1, 1, 4, 2, 1, 3})
}

func heightChecker(heights []int) int {
	ar := make([]int, 101)
	for _, h := range heights {
		ar[h]++
	}

	var res, ptr int

	for height, cnt := range ar {
		for cnt > 0 {
			if heights[ptr] != height {
				res++
			}
			ptr++
			cnt--
		}
	}
	return res
}
