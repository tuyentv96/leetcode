type MinStack struct {
	elements, mins []int
	length         int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		elements: make([]int, 0),
		mins:     make([]int, 0),
		length:   0,
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
	if this.length == 0 || this.mins[this.length-1] > val {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, this.mins[this.length-1])
	}

	this.length++
}

func (this *MinStack) Pop() {
	this.elements = this.elements[:this.length-1]
	this.mins = this.mins[:this.length-1]

	this.length--
}

func (this *MinStack) Top() int {
	return this.elements[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[this.length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
