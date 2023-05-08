package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("#### permu")
	permu(a, 0, 5, 2)
	// visit := make([]bool, len(a))
	fmt.Println("#### combi")
	// comb(a, visit, 0, 5, 4)
}

func permu(ar []int, d, n, r int) {
	if d == r {
		for i := 0; i < r; i++ {
			fmt.Printf("%d ", ar[i])
		}
		fmt.Println()
		return
	}

	for i := d; i < n; i++ {
		ar[d], ar[i] = ar[i], ar[d]
		permu(ar, d+1, n, r)
		ar[d], ar[i] = ar[i], ar[d]
	}
}

func comb(ar []int, visit []bool, d, n, r int) {
	if r == 0 {
		for i := 0; i < n; i++ {
			if visit[i] {
				fmt.Printf("%d ", ar[i])
			}
		}
		fmt.Println()
		return
	}

	for i := d; i < n; i++ {
		visit[i] = true
		comb(ar, visit, i+1, n, r-1)
		visit[i] = false
	}
}
