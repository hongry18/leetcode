package main

func permute(nums []int) [][]int {
	var res [][]int
	permu(nums, &res, 0, len(nums), len(nums))
	return res
}

func permu(ar []int, res *[][]int, d, n, r int) {
	if d == r {
		var item []int
		for i := 0; i < r; i++ {
			item = append(item, ar[i])
		}
		*res = append(*res, item)
		return
	}

	for i := d; i < n; i++ {
		ar[d], ar[i] = ar[i], ar[d]
		permu(ar, res, d+1, n, r)
		ar[d], ar[i] = ar[i], ar[d]
	}
}
