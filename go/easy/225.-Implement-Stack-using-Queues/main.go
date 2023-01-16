package main

type MyStack struct {
	ar []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.ar = append(this.ar, x)
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	p := this.ar[len(this.ar)-1]
	this.ar = this.ar[:len(this.ar)-1]
	return p
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.ar[len(this.ar)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.ar) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
