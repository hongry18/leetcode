package main

func main() {
	sortArrayByParityII([]int{4, 7, 5, 2})
	sortArrayByParityII([]int{4, 2, 5, 7})
	sortArrayByParityII([]int{2, 3})
	sortArrayByParityII([]int{648, 831, 560, 986, 192, 424, 997, 829, 897, 843})
	sortArrayByParityII([]int{8, 1, 60, 86, 92, 24, 7, 9, 7, 3})
}

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	res := make([]int, len(nums))
	for _, num := range nums {
		if num&1 == 0 {
			res[even] = num
			even += 2
		} else {
			res[odd] = num
			odd += 2

		}
	}
	return res
}
