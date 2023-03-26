package main

func main() {
	validMountainArray([]int{0, 2, 3, 4, 5, 2, 1, 0})
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	var i int

	for ; i < len(arr)-1 && arr[i] < arr[i+1]; i++ {
	}

	if i == 0 || i == len(arr)-1 {
		return false
	}

	for ; i < len(arr)-1 && arr[i] > arr[i+1]; i++ {
	}

	return i == len(arr)-1
}
