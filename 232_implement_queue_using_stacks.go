package main

type MyQueue struct {
	stack  []int
	cursor int
}

func Constructor() MyQueue {
	return MyQueue{stack: make([]int, 0), cursor: 0}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
	this.cursor++
}

func (this *MyQueue) Pop() int {
	if this.cursor > 0 {
		pop := this.stack[0]
		this.cursor--
		this.stack = this.stack[1:]
		return pop
	} else {
		return 0
	}
}

func (this *MyQueue) Peek() int {
	if this.cursor > 0 {
		return this.stack[0]
	} else {
		return 0
	}
}

func (this *MyQueue) Empty() bool {
	return this.cursor == 0
}
