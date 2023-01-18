package main

type NumArray struct {
	size        int
	prefixTable []int
}

func Constructor(n []int) NumArray {
	arr := NumArray{size: len(n)}

	arr.prefixTable = make([]int, arr.size)

	arr.prefixTable[0] = n[0]
	for i := 1; i < arr.size; i++ {
		arr.prefixTable[i] = arr.prefixTable[i-1] + n[i]
	}

	return arr
}

func (this *NumArray) SumRange(left, right int) int {
	if left == 0 {
		return this.prefixTable[right]
	}

	return this.prefixTable[right] - this.prefixTable[left-1]
}
