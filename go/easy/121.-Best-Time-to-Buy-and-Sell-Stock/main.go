package main

func maxProfit(prices []int) int {
	var lo, price int = 100000, 0

	for _, cur := range prices {
		if cur < lo {
			lo = cur
			continue
		}
		if cur-lo > price {
			price = cur - lo
		}
	}

	return price
}
