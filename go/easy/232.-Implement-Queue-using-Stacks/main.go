package main

type MyQueue struct {
	S1 MyStack
	S2 MyStack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.S1.Push(x)
}

func (this *MyQueue) Pop() int {
	var x int
	if this.S2.Empty() {
		for !this.S1.Empty() {
			x = this.S1.Pop()
			this.S2.Push(x)
		}
	}
	x = this.S2.Pop()
	return x
}

func (this *MyQueue) Peek() int {
	var x int
	if this.S2.Empty() {
		for !this.S1.Empty() {
			x = this.S1.Pop()
			this.S2.Push(x)
		}
	}
	x = this.S2.Top()
	return x
}

func (this *MyQueue) Empty() bool {
	return this.S1.Empty() && this.S2.Empty()
}

type MyStack struct {
	Data []int
}

func (s *MyStack) Push(x int) {
	s.Data = append(s.Data, x)
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		return 0
	}

	x := s.Data[s.Size()-1]
	s.Data = s.Data[:s.Size()-1]
	return x
}

func (s *MyStack) Top() int {
	return s.Data[s.Size()-1]
}

func (s *MyStack) Size() int {
	return len(s.Data)
}

func (s *MyStack) Empty() bool {
	return s.Size() == 0
}
