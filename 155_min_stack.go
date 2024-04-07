package main

type MinStack struct {
	stack []int
	min   int
}

func ConstructorMinStack() MinStack {
	return MinStack{stack: []int{}, min: 1<<63 - 1}
}

func (this *MinStack) Push(val int) {
	if this.min > val {
		this.min = val
	}

	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	popValue := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if popValue == this.min && len(this.stack) > 0 {
		newMinVal := this.stack[0]

		for _, num := range this.stack {
			if num < newMinVal {
				newMinVal = num
			}
		}

		this.min = newMinVal
	}

	if len(this.stack) == 0 {
		this.min = 1<<63 - 1
	}

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
