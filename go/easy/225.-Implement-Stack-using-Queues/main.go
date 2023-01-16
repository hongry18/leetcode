package main

type MyStack struct {
	ar  []int
	pos int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.ar = append(this.ar, x)
	this.pos++
}

func (this *MyStack) Pop() int {
	return 0
}

func (this *MyStack) Top() int {
	return 0
}

func (this *MyStack) Empty() bool {
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
