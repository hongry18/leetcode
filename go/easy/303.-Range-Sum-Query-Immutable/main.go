package main

type NumArray struct {
	Data []int
}

func Constructor(n []int) NumArray {
	return NumArray{Data: n}
}

func (this *NumArray) SumRange(left, right int) int {
	var sum int

	for i := left; i <= right; i++ {
		sum += this.Data[i]
	}

	return sum
}
