package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	var find bool
	for _, v := range flowerbed {
		if v != 0 && !find {
			find = true
		}

	}
	return true
}

func main() {
	canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)
	canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)
}
