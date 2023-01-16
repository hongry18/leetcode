package main

type MyStack struct {
	Q MyQueue
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.Q.Push(x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < this.Q.Size()-1; i++ {
		this.Q.Push(this.Q.Pop())
	}
	return this.Q.Pop()
}

func (this *MyStack) Top() int {
	for i := 0; i < this.Q.Size()-1; i++ {
		this.Q.Push(this.Q.Pop())
	}
	x := this.Q.Pop()
	this.Q.Push(x)
	return x
}

func (this *MyStack) Empty() bool {
	return this.Q.Empty()
}

type MyQueue struct {
	Data []int
}

func (q *MyQueue) Push(x int) {
	q.Data = append(q.Data, x)
}

func (q *MyQueue) Peek() int {
	if q.Empty() {
		return 0
	}

	return q.Data[0]
}

func (q *MyQueue) Pop() int {
	if q.Empty() {
		return 0
	}

	t := q.Data[0]
	q.Data = q.Data[1:]
	return t
}

func (q *MyQueue) Size() int {
	return len(q.Data)
}

func (q *MyQueue) Empty() bool {
	return q.Size() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
