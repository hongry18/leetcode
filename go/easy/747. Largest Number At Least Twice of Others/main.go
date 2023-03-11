package main

func main() {
	dominantIndex([]int{3, 6, 1, 0})
	dominantIndex([]int{1, 2, 3, 4})
	dominantIndex([]int{4, 3, 2, 1})
}

func dominantIndex(nums []int) int {
	var a, b, c int

	for i := 0; i < len(nums); i++ {
		if nums[i] > a {
			b = a
			a = nums[i]
			c = i
		} else if nums[i] > b {
			b = nums[i]
		}
	}

	if b*2 <= a {
		return c
	}

	return -1
}
