package main

func main() {
	findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10})
}

func findSpecialInteger(arr []int) int {
	_m := map[int]int{}
	for _, n := range arr {
		_m[n]++
	}

	for k, v := range _m {
		if v*100/len(arr) > 25 {
			return k
		}
	}
	return 0
}
