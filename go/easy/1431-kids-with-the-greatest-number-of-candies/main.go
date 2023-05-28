package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res []bool
	var max int
	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	for _, v := range candies {
		res = append(res, v+extraCandies >= max)
	}

	return res
}
