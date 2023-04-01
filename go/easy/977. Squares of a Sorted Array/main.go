package main

func main() {
	sortedSquares([]int{-4, -1, 0, 3, 10})
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[l]) > abs(nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
