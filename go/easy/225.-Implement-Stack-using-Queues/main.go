package main

type MyStack struct {
	ar  []int
	pos int
}

func Constructor() MyStack {
	return MyStack{pos: -1}
}

func (this *MyStack) Push(x int) {
	this.ar = append(this.ar, x)
	this.pos++
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	p := this.ar[this.pos]
	this.ar = this.ar[:this.pos]
	this.pos--
	return p
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.ar[this.pos]
}

func (this *MyStack) Empty() bool {
	return this.pos == -1
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
