package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
}

func distributeCandies(candies int, num_people int) []int {
	var res = make([]int, num_people)

	for i, c := 0, 1; ; i, c = i+1, c+1 {
		if candies-c < 0 {
			res[i%num_people] += candies
			break
		}

		candies -= c
		res[i%num_people] += c
	}

	return res
}
