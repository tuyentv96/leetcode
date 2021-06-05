type MinStack struct {
	values, mins []int
	length       int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		values: make([]int, 0),
		mins:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
	if this.length > 0 && this.mins[this.length-1] <= val {
		this.mins = append(this.mins, this.mins[this.length-1])
	} else {
		this.mins = append(this.mins, val)
	}

	this.length++
}

func (this *MinStack) Pop() {
	this.values = this.values[:this.length-1]
	this.mins = this.mins[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.values[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[this.length-1]
}
