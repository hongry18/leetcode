package main

import (
	"testing"
)

func TestXxx(t *testing.T) {
	t.Log(countLargestGroup(13))
	t.Log(countLargestGroup(19))
	t.Log(countLargestGroup(29))
	t.Log(countLargestGroup(39))
	t.Log(countLargestGroup(49))
	t.Log(countLargestGroup(99))
	t.Log(countLargestGroup(1000))
}

func countLargestGroup(n int) int {
	var ar [100]int
	var tmp, max, res int
	for i := 1; i <= n; i++ {
		if i%10 == 0 {
			tmp = calc(i)
		}
		idx := tmp + (i % 10)
		ar[idx]++
		if ar[idx] > max {
			max = ar[idx]
		}
	}

	for i := 0; i < 100; i++ {
		if ar[i] == max {
			res++
		}
	}
	return res
}

func calc(i int) int {
	var ar []int
	for i > 0 {
		ar = append(ar, i%10)
		i /= 10
	}
	var res int
	for _, v := range ar[1:] {
		res += v
	}
	return res
}
