package main

func main() {
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
}

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != 0 {
			continue
		}

		for j := len(arr) - 1; j > i; j-- {
			arr[j] = arr[j-1]
		}
		i++
	}
}
