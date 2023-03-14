package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		j, k := 0, len(image[i])-1
		for j <= k {
			image[i][j], image[i][k] = image[i][k], image[i][j]
			if j == k {
				image[i][j] ^= 1
			} else {
				image[i][j] ^= 1
				image[i][k] ^= 1
			}
			j++
			k--
		}
	}
	return image
}
